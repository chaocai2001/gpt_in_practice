package demo

import (
	"errors"
)

type LocalStorage struct {
	data map[string]string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		data: make(map[string]string),
	}
}

func (ls *LocalStorage) RetieveData(token string) (string, error) {
	data, ok := ls.data[token]
	if !ok {
		return "", errors.New("data not found")
	}
	return data, nil
}

func (ls *LocalStorage) StoreData(token string, data string) error {
	ls.data[token] = data
	return nil
}
