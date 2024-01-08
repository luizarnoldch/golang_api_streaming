package service

import (
	"log"
	"main/src/domain/model"
	"main/src/domain/repository"

	"github.com/google/uuid"
)

type StreamDynamoDBService struct {
	repo repository.StreamRepository
}

func NewStreamDynamoDBService(repo repository.StreamRepository) *StreamDynamoDBService {
	return &StreamDynamoDBService{
		repo: repo,
	}
}

func (service *StreamDynamoDBService) GetAllStream() ([]model.Stream, error) {
	return service.repo.GetAllStream()
}

func (service *StreamDynamoDBService) CreateStream(req *model.Stream) (*model.Stream, error) {
	if err := req.Validate(); err != nil {
		log.Printf("error while validating request from CreateStream Service: %v", err)
		return &model.Stream{}, err
	}

	req.ID = uuid.NewString()

	return service.repo.CreateStream(req)
}
