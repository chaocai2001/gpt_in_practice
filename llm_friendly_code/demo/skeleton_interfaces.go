package demo

// Storage is to persist the data
type Storage interface {
	// RetiveData is to retrive the data by the associated token.
	RetieveData(token string) (string, error)

	// StoreData is to persist the data,
	// input paramters:
	//   token is used to retrive the associated data
	//
	StoreData(token string, data string) error
}

// Processor is to provide the processing service
type Processor interface {
	// Process is to the raw data
	Process(raw string) (string, error)
}

type TokenCreator interface {
	CreateToken(data string) string
}

type ProcessingService interface {
	// Process is to process input string and persist the result.
	// And a token will be returned which can be used for retriving the processed result
	Process(raw string) (string, error)

	// Retrive is to retrive the processed result by given token,
	// which is returned by "Process" method.
	Retrieve(token string) (string, error)
}
