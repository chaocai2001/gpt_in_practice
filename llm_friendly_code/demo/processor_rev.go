package demo

import (
	"errors"
	//	"strings"
)

type ReverseProcessor struct{}

func (rp *ReverseProcessor) Process(raw string) (string, error) {
	if raw == "" {
		return "", errors.New("input string is empty")
	}

	runes := []rune(raw)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes), nil
}
