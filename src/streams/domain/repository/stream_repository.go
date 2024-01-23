package repository

import (
	"main/src/streams/domain/model"
	appError "main/utils/error"
)

type StreamRepository interface {
	GetAllStream() ([]model.Stream, *appError.Error)
	CreateStream(*model.Stream) (*model.Stream, *appError.Error)
	GetStreamById(string) (*model.Stream, *appError.Error)
}