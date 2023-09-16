package demo

import (
	//"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProcessor struct {
	mock.Mock
}

func (m *MockProcessor) Process(raw string) (string, error) {
	args := m.Called(raw)
	return args.String(0), args.Error(1)
}

type MockTokenCreator struct {
	mock.Mock
}

func (m *MockTokenCreator) CreateToken(data string) string {
	args := m.Called(data)
	return args.String(0)
}

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) RetieveData(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) StoreData(token string, data string) error {
	args := m.Called(token, data)
	return args.Error(0)
}

func TestProcessingServiceImpl_Process(t *testing.T) {
	mockProcessor := new(MockProcessor)
	mockTokenCreator := new(MockTokenCreator)
	mockStorage := new(MockStorage)

	service := NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	raw := "raw data"
	processed := "processed data"
	token := "token"

	mockProcessor.On("Process", raw).Return(processed, nil)
	mockTokenCreator.On("CreateToken", processed).Return(token)
	mockStorage.On("StoreData", token, processed).Return(nil)

	result, err := service.Process(raw)

	assert.NoError(t, err)
	assert.Equal(t, token, result)

	mockProcessor.AssertExpectations(t)
	mockTokenCreator.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}

func TestProcessingServiceImpl_Retrieve(t *testing.T) {
	mockProcessor := new(MockProcessor)
	mockTokenCreator := new(MockTokenCreator)
	mockStorage := new(MockStorage)

	service := NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	token := "token"
	data := "data"

	mockStorage.On("RetieveData", token).Return(data, nil)

	result, err := service.Retrieve(token)

	assert.NoError(t, err)
	assert.Equal(t, data, result)

	mockStorage.AssertExpectations(t)
}
