package service

import (
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/store"
	"errors"
)

type Store interface {
	Info(song model.Song) (*store.StoreSongDetail, error)
}

type Service struct {
	store Store
}

func New(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) Info(song model.Song) (*model.SongDetail, error) {
	storeSD, err := s.store.Info(song)
	if err != nil {
		return nil, err
	}

	if storeSD == nil {
		return nil, errors.New("nil value")
	}

	return &model.SongDetail{
		ReleaseDate: storeSD.ReleaseDate,
		Text:        storeSD.Text,
		Link:        storeSD.Link,
	}, nil
}
