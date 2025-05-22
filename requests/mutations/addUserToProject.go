package mutations

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

func AddUserToProject(client *graphql.Client, token string, projectID string, userID string) error {
	req := graphql.NewRequest(`
       mutation addProjectUser($projectID: ID!, $userID: ID!) {
  	   		addProjectUser(projectID: $projectID, userID: $userID) {	
				id
  			}
		}`)

	req.Header.Set("Authorization", "Bearer "+token)

	req.Var("projectID", projectID)
	req.Var("userID", userID)

	ctx := context.Background()

	if err := client.Run(ctx, req, nil); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
