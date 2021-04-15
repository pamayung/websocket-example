package model

type Timeline struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	TimeLine string `json:"timeline"`
}

type TimelineUsecase interface {
	Store(*Timeline) error
}

type TimelineRepository interface {
	Store(t *Timeline) error
}
