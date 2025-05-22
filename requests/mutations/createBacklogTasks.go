package mutations

import (
	"context"
	"log"
	"p4planAPIExamples/models"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/machinebox/graphql"
)

type CreateBacklogTaskInput struct {
	Name        string        `json:"name"`
	Status      models.Status `json:"status"`
	IsUserStory bool          `json:"isUserStory"`
}

func CreateBacklogTasks(client *graphql.Client, token string, projectID string) error {
	input := make([]CreateBacklogTaskInput, 0, 10)
	for i := 0; i < 10; i++ {
		input = append(input, CreateBacklogTaskInput{
			Name:        gofakeit.LoremIpsumSentence(3),
			Status:      models.StatusNotDone,
			IsUserStory: true,
		})
	}

	req := graphql.NewRequest(`
        mutation createBacklogTasks($backlogProjectID: ID! $createBacklogTasksInput: [CreateBacklogTaskInput]!) {
  			createBacklogTasks(projectID: $backlogProjectID, createBacklogTasksInput: $createBacklogTasksInput) {
    			id
  			}
		}`)

	req.Header.Set("Authorization", "Bearer "+token)

	req.Var("createBacklogTasksInput", input)
	req.Var("backlogProjectID", projectID)

	ctx := context.Background()

	if err := client.Run(ctx, req, nil); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
