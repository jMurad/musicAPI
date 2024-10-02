package service_test

import (
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Info(t *testing.T) {
	svc := service.New(service.NewTStore())
	_, err := svc.Info(model.Song{})
	assert.Error(t, err)
}
