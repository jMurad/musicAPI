package app_test

import (
	"MusicApi/musicInfo/internal/app"
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/store/random"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestInfoHandler(t *testing.T) {
	cases := []struct {
		name         string
		song         interface{}
		expectedCode int
	}{
		{
			name:         "Success",
			song:         model.TestSongOK(t),
			expectedCode: 200,
		},
		{
			name:         "No required fields",
			song:         model.TestSongWithoutField(t),
			expectedCode: 400,
		},
		{
			name:         "Empty request",
			song:         model.TestSongNil(t),
			expectedCode: 400,
		},
		{
			name:         "Internal error",
			song:         model.TestSongErr(t),
			expectedCode: 422,
		},
		{
			name:         "No content",
			song:         model.TestSongEmpty(t),
			expectedCode: 204,
		},
	}

	tsvc := app.NewTService(t, random.New())
	logtrash := app.NewTLog(t)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.song)
			req, _ := http.NewRequest(http.MethodGet, "/info", b)

			rec := httptest.NewRecorder()

			r := chi.NewRouter()
			r.Get("/info", app.Info(logtrash, tsvc))
			r.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code, rec.Body)
		})
	}
}
