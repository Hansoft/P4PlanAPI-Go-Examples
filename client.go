package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"p4planAPIExamples/requests"
	"p4planAPIExamples/requests/mutations"
	"p4planAPIExamples/requests/queries"
	"p4planAPIExamples/util"

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

	// Get all the projects available to the user
	projects, err := queries.GetProjects(client, accessToken)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Projects received: %v", len(projects))

	// Get all the items of the first project's planning tab, qa tab and backlog tab

	items, err := queries.GetItems(client, accessToken, projects[0].ID)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Planning tab items for %s received: %v", projects[0].Name, len(items))

	newProject, err := mutations.CreateProject(client, accessToken)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New project created: %s", newProject.Name)

	authenticatedUserId, err := util.ExtractSubFromToken(accessToken)

	if err != nil {
		log.Fatal(err)
	}

	err = mutations.AddUserToProject(client, accessToken, newProject.ID, authenticatedUserId)

	if err != nil {
		log.Fatal(err)
	}

	err = mutations.MakeUserMainManager(client, accessToken, newProject.ID, authenticatedUserId)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authenticated user added to new project and is now a main manager")

	err = mutations.CreateBacklogTasks(client, accessToken, newProject.Backlog.ID)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("10 new backlog tasks have been created in the newly created project")

}
