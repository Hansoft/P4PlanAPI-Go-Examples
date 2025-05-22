package requests

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

type ResponseStruct struct {
	Login struct {
		AccessToken string `json:"access_token"`
	}
}

func Login(client *graphql.Client) (string, error) {
	p4PlanUsername := os.Getenv("P4PLAN_USERNAME")
	p4PlanPassword := os.Getenv("P4PLAN_PASSWORD")

	req := graphql.NewRequest(`
        	mutation login($username: String!, $password: String!) { 
				login(loginUserInput: {username: $username password: $password}) {
  					access_token 
			}
		}
	`)

	req.Var("username", p4PlanUsername)
	req.Var("password", p4PlanPassword)

	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
		return "", err
	}

	log.Printf("User successfully Authenticated")

	return respData.Login.AccessToken, nil
}
