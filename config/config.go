package config

type (
	CONFIG struct {
		MICRO MICRO
		ENV string
	}

	MICRO struct {
		DB DB
		API API
	}

	DB struct {
		STREAM_DYNAMODB STREAM_DYNAMODB
		USER_DYNAMODB USER_DYNAMODB
	}

	API struct {
		HOST string
		PORT string
	}

	STREAM_DYNAMODB struct {
		TABLE_NAME string
	}
	USER_DYNAMODB struct {
		TABLE_NAME string
	}
)