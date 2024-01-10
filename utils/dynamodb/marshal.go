package dynamodb

import (
	"log"
	"main/src/domain/model"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UnmarshalStream(item map[string]types.AttributeValue) (*model.Stream, error) {
	var stream model.Stream
	if err := attributevalue.UnmarshalMap(item, &stream); err != nil {
		log.Printf("GetStreamById: Error unmarshaling result: %s", err)
		return nil, err
	}
	return &stream, nil
}

func MarshalMapStream(stream *model.Stream) (map[string]types.AttributeValue, error) {
	marshalStream, err := attributevalue.MarshalMap(stream)
    if err != nil {
        log.Printf("CreateStream: Error while marshaling stream: %v", err)
        return nil, err
    }
	return marshalStream, nil
}
