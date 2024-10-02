package service

import (
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/store"
)

type TStore struct {
}

func NewTStore() *TStore {
	return &TStore{}
}

func (ts *TStore) Info(_ model.Song) (*store.StoreSongDetail, error) {
	return nil, nil
}
