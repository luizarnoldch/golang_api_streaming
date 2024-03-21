package service

import (
	"log"
	"main/src/streams/domain/model"
	"main/src/streams/domain/repository"
	appError "main/utils/error"

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

func (service *StreamDynamoDBService) UpdateStreamById(stream_id string, stream *model.Stream) (*model.Stream, *appError.Error) {
	return service.repo.UpdateStreamById(stream_id, stream)
}

func (service *StreamDynamoDBService) GetStreamById(stream_id string) (*model.Stream, *appError.Error) {
	return service.repo.GetStreamById(stream_id)
}

func (service *StreamDynamoDBService) GetAllStream() ([]model.Stream, *appError.Error) {
	return service.repo.GetAllStream()
}

func (service *StreamDynamoDBService) DeleteStream(stream_id string) *appError.Error {
	return service.repo.DeleteStream(stream_id)
}

func (service *StreamDynamoDBService) CreateStream(req *model.Stream) (*model.Stream, *appError.Error) {
	req.ID = uuid.NewString()
	if err := req.Validate(); err != nil {
		log.Printf("error while validating request from CreateStream Service: %v", err)
		return nil, err
	}

	return service.repo.CreateStream(req)
}
