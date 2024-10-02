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

func (s *Store) Info(song model.Song) (*StoreSongDetail, error) {

	switch song.Group {
	case "error":
		return nil, errors.New("internal error")
	case "empty":
		return nil, store.ErrSongNotFound
	}

	return &StoreSongDetail{
		ReleaseDate: genReleaseDate(),
		Text:        genText(),
		Link:        getLink(),
	}, nil
}
