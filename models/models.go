package models

type TimeSpan struct {
	Start  string `json:"start"`
	Finish string `json:"finish"`
}

type Item struct {
	ID                  string      `json:"id"`
	Name                *string     `json:"name,omitempty"`
	SubprojectPath      *string     `json:"subprojectPath,omitempty"`
	LocalID             string      `json:"localID"`
	Typename            string      `json:"__typename"`
	UserStory           *string     `json:"userStory,omitempty"`
	Status              *string     `json:"status,omitempty"`
	DetailedDescription *string     `json:"detailedDescription,omitempty"`
	TimeSpans           *[]TimeSpan `json:"timeSpans,omitempty"`
	Start               *string     `json:"start,omitempty"`
	Finish              *string     `json:"finish,omitempty"`
	Date                *string     `json:"date,omitempty"`
}

type Project struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	QA      QAOrBacklog `json:"qa"`
	Backlog QAOrBacklog `json:"backlog"`
}

type QAOrBacklog struct {
	ID string `json:"id"`
}
