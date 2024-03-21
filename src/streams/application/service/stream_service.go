package service

import (
	"main/src/streams/domain/model"
	appError "main/utils/error"
)

type StreamService interface {
	UpdateStreamById(stream_id string, stream *model.Stream) (*model.Stream, *appError.Error)
	GetStreamById(string) (*model.Stream, *appError.Error)
	GetAllStream() ([]model.Stream, *appError.Error)
	DeleteStream(string) *appError.Error
	CreateStream(*model.Stream) (*model.Stream, *appError.Error)
}
