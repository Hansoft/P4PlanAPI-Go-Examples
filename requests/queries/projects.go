package queries

import (
	"context"
	"log"
	"p4planAPIExamples/models"

	"github.com/machinebox/graphql"
)

type ResponseStruct struct {
	Projects []models.Project `json:"projects"`
}

func GetProjects(client *graphql.Client, token string) ([]models.Project, error) {
	req := graphql.NewRequest(`
        query projects {
 			projects {
    			id
    			name
    			qa {
      				id
    			}
   	 			backlog {
      				id
    			}
  			}
		}`)

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", "Bearer "+token)

	ctx := context.Background()

	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return respData.Projects, nil
}
