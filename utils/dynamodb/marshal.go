package dynamodb

import (
	"log"
	"main/src/streams/domain/model"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UnmarshalStream(item map[string]types.AttributeValue) (*model.Stream, error) {
	var stream model.Stream
	if err := attributevalue.UnmarshalMap(item, &stream); err != nil {
		log.Printf("Error unmarshaling result: %s", err)
		return nil, err
	}

    log.Printf("Unmarshalled Stream: %+v", stream)
	return &stream, nil
}

func MarshalMapStream(stream *model.Stream) (map[string]types.AttributeValue, error) {
	marshalStream, _ := attributevalue.MarshalMap(stream)
    log.Printf("Marshalled Stream: %+v", marshalStream)
	return marshalStream, nil
}
