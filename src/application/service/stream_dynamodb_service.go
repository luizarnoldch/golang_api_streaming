package service

import (
	"main/src/domain/model"
	"main/src/domain/repository"
)

type StreamDynamoDBService struct {
	repo repository.StreamRepository
}

func NewStreamDynamoDBService(repo repository.StreamRepository) *StreamDynamoDBService {
	return &StreamDynamoDBService{
		repo: repo,
	}
}

func (repo *StreamDynamoDBService) GetAllStream() ([]model.Stream, error) {
	return []model.Stream{}, nil
}
