package app

import (
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/store/random"
	"context"
	"log/slog"
	"testing"
)

func NewTLog(t *testing.T) *slog.Logger {
	return slog.New(NewTHandler(t))
}

type THandler struct {
	T *testing.T
}

func NewTHandler(t *testing.T) *THandler {
	return &THandler{t}
}

func (h *THandler) Handle(_ context.Context, _ slog.Record) error {
	h.T.Helper()
	return nil
}

func (h *THandler) WithAttrs(_ []slog.Attr) slog.Handler {
	h.T.Helper()
	return h
}

func (h *THandler) WithGroup(_ string) slog.Handler {
	h.T.Helper()
	return h
}

func (h *THandler) Enabled(_ context.Context, _ slog.Level) bool {
	h.T.Helper()
	return false
}

type Store interface {
	Info(song model.Song) (*random.StoreSongDetail, error)
}

type TService struct {
	T     *testing.T
	store Store
}

func NewTService(t *testing.T, store Store) *TService {
	return &TService{
		T:     t,
		store: store,
	}
}

func (ts TService) Info(song model.Song) (*model.SongDetail, error) {
	ts.T.Helper()

	sd, err := ts.store.Info(song)
	if err != nil {
		return nil, err
	}

	return &model.SongDetail{
		ReleaseDate: sd.ReleaseDate,
		Text:        sd.Text,
		Link:        sd.Link,
	}, err
}
