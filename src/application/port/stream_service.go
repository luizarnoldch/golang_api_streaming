package port

import (
	"main/src/domain/model"
	appError "main/utils/error"
)

type StreamService interface {
	GetAllStream() ([]model.Stream, *appError.Error)
	CreateStream(*model.Stream) (*model.Stream, *appError.Error)
	GetStreamById(string) (*model.Stream, *appError.Error)
}