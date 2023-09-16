package demo

import (
	"github.com/google/uuid"
)

type MyTokenCreator struct{}

func (m *MyTokenCreator) CreateToken(data string) string {
	token := uuid.New().String()
	return token
}
