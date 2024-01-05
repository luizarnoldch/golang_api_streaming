package repository

import "main/src/domain/model"

type StreamRepository interface {
	GetAllStream() ([]model.Stream, error)
}