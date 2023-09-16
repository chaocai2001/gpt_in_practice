package demo

type ProcessingServiceImpl struct {
	Storage      Storage
	Processor    Processor
	TokenCreator TokenCreator
}

func (ps *ProcessingServiceImpl) Process(raw string) (string, error) {
	processed, err := ps.Processor.Process(raw)
	if err != nil {
		return "", err
	}
	token := ps.TokenCreator.CreateToken(processed)
	err1 := ps.Storage.StoreData(token, processed)
	return token, err1
}

func (ps *ProcessingServiceImpl) Retrieve(token string) (string, error) {
	return ps.Storage.RetieveData(token)
}

// NewProcessingService is the Builder method to
func NewProcessingService(processor Processor,
	tokenCreator TokenCreator, storage Storage) ProcessingService {
	return &ProcessingServiceImpl{
		Storage:      storage,
		Processor:    processor,
		TokenCreator: tokenCreator,
	}
}
