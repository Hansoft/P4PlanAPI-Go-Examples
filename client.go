package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"p4planAPIExamples/requests"

	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

type ResponseStruct struct {
	AppInfo struct {
		APIVersion string `json:"apiVersion"`
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	p4PlanAPIHost := os.Getenv("P4PLAN_API_HOST")
	url := fmt.Sprintf("%s/graphql", p4PlanAPIHost)
	client := graphql.NewClient(url)

	req := graphql.NewRequest(`
        query appInfo {
		   appInfo {
		   		apiVersion
			}
        }
    `)

	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	log.Printf("Go client connected successfully to P4 Plan API: %+v\n", respData.AppInfo.APIVersion)

	accessToken, err := requests.Login(client)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Access token: %s", accessToken)
}
