package queries

import (
	"context"
	"log"
	"p4planAPIExamples/models"

	"github.com/machinebox/graphql"
)

type ItemsResponseStruct struct {
	Items []models.Item `json:"items"`
}

func GetItems(client *graphql.Client, token string, projectID string) ([]models.Item, error) {
	req := graphql.NewRequest(`
       query items($projectID: ID!) {
  		items(id: $projectID) {
    		id
    		name
    		subprojectPath
    		localID
    		subprojectPath
    		__typename
    		... on BacklogTask {
				userStory
				status
			}
			... on Bug {
				detailedDescription
				status
			}
			... on ScheduledTask {
				status
				timeSpans {
					start
					finish
				}
			}
			... on Sprint {
				start
				finish
			}
			... on Release {
				date
			}
		}
	}`)

	req.Var("projectID", projectID)

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", "Bearer "+token)

	ctx := context.Background()

	var respData ItemsResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return respData.Items, nil
}
