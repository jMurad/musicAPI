package app

import (
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/store"
	"context"
	"errors"
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
	Info(song model.Song) (*store.StoreSongDetail, error)
}

type TService struct {
	t *testing.T
}

func NewTService(t *testing.T) *TService {
	return &TService{
		t: t,
	}
}

func (ts TService) Info(song model.Song) (*model.SongDetail, error) {
	ts.t.Helper()

	switch song {
	case *model.TestSongEmpty(ts.t):
		return nil, store.ErrSongNotFound
	case *model.TestSongErr(ts.t):
		return nil, errors.New("internal error")
	}

	return &model.SongDetail{}, nil
}
