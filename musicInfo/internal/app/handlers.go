package app

import (
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/store"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Group string `json:"group" validate:"required"`
	Song  string `json:"song" validate:"required"`
}

type Service interface {
	Info(song model.Song) (*model.SongDetail, error)
}

func Info(log *slog.Logger, svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const place = "app.handlers.Info"

		log := log.With(
			slog.String("place", place),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")

			rerror(w, r, http.StatusBadRequest, errors.New("empty request"))

			return
		}

		if err != nil {
			log.Error("failed to decode request body", slog.Attr{
				Key:   "error",
				Value: slog.StringValue(err.Error()),
			})

			rerror(w, r, http.StatusBadRequest, err)

			return
		}

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", slog.Attr{
				Key:   "error",
				Value: slog.StringValue(err.Error()),
			})

			rerror(w, r, http.StatusBadRequest, ValidationError(validateErr))

			return
		}

		song := model.Song{
			Group: req.Group,
			Song:  req.Song,
		}

		log.Info("request body decoded", slog.Any("song", song))

		songDetail, err := svc.Info(song)
		if errors.Is(err, store.ErrSongNotFound) {
			log.Info("song detail not found", "song", song)

			rerror(w, r, http.StatusNoContent, errors.New("song detail not found"))

			return
		}
		if err != nil {
			log.Info("failed to get song detail", "song", song)

			rerror(w, r, http.StatusInternalServerError, errors.New("internal error"))

			return
		}

		log.Info("got song detail", slog.Any("song detail", songDetail))

		respond(w, r, http.StatusOK, songDetail)
	}
}

func rerror(w http.ResponseWriter, r *http.Request, code int, err error) {
	respond(w, r, code, map[string]string{"error": err.Error()})
}

func respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		render.JSON(w, r, data)
	}
}
