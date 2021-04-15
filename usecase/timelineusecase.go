package usecase

import (
	"websocket-example/model"
	"websocket-example/repository"
)

type TimelineUsecase struct {
	timelineRepo model.TimelineRepository
}

func (t *TimelineUsecase) Store(m *model.Timeline) (err error) {
	err = repository.Store(m)
	return
}
