package mutations

import (
	"context"
	"log"
	"p4planAPIExamples/models"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/machinebox/graphql"
)

type CreateProjectInput struct {
	Name           string `json:"name"`
	ArchivedStatus bool   `json:"archivedStatus"`
}

type ResponseStruct struct {
	Project models.Project `json:"createProject"`
}

func CreateProject(client *graphql.Client, token string) (models.Project, error) {
	input := CreateProjectInput{
		Name:           gofakeit.BuzzWord(),
		ArchivedStatus: false,
	}

	req := graphql.NewRequest(`
        mutation createProject($createProjectInput: CreateProjectInput!) {
 			createProject(createProjectInput: $createProjectInput) {
				id
				name
				backlog {
					id
				}
				qa {
					id
				}
			}
		}`)

	req.Header.Set("Authorization", "Bearer "+token)

	req.Var("createProjectInput", input)

	ctx := context.Background()

	var respData ResponseStruct

	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
		return models.Project{}, err
	}

	return respData.Project, nil
}
