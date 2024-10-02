package random

import (
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/store"
	"errors"
)

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) Info(song model.Song) (*store.StoreSongDetail, error) {

	switch song.Group {
	case "error":
		return nil, errors.New("internal error")
	case "empty":
		return nil, store.ErrSongNotFound
	}

	return &store.StoreSongDetail{
		ReleaseDate: genReleaseDate(),
		Text:        genText(),
		Link:        getLink(),
	}, nil
}
