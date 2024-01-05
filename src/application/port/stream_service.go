package port

import "main/src/domain/model"

type StreamService interface {
	GetAllStream() ([]model.Stream, error)
}