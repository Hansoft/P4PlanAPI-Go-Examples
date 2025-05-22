package mutations

import (
	"context"
	"log"
	"p4planAPIExamples/models"

	"github.com/machinebox/graphql"
)

func MakeUserMainManager(client *graphql.Client, token string, projectID string, userID string) error {

	req := graphql.NewRequest(`
       mutation updateProjectUserAccessRights($projectID: ID!, $userID: ID!, $accessRights: ProjectUserAccessRightsInput!) {
  	   		updateProjectUserAccessRights(
    			projectID: $projectID userID: $userID accessRights: $accessRights) {
    			user {
      				id
				}
			}
		}`)
	req.Header.Set("Authorization", "Bearer "+token)

	req.Var("projectID", projectID)
	req.Var("userID", userID)
	req.Var("accessRights", models.AccessRightsInput{IsMainManager: true, CanAccessProjectHistory: true})

	ctx := context.Background()

	if err := client.Run(ctx, req, nil); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
