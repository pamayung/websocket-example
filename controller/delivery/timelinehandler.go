package delivery

import (
	"websocket-example/model"
	"websocket-example/usecase"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TimelineHandler struct {
	TUsecase usecase.TimelineUsecase
}

func (t *TimelineHandler) Store(m string) (err error) {
	var timeline model.Timeline
	timeline.Text = m
	timeline.TimeLine = "me"

	err = t.TUsecase.Store(&timeline)

	return err
}
