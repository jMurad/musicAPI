package random_test

import (
	"MusicApi/musicInfo/internal/model"
	"MusicApi/musicInfo/internal/store"
	"MusicApi/musicInfo/internal/store/random"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore_Info(t *testing.T) {
	rstore := random.New()

	_, err := rstore.Info(*model.TestSongOK(t))
	assert.NoError(t, err)

	_, err = rstore.Info(*model.TestSongEmpty(t))
	assert.Equal(t, store.ErrSongNotFound, err)

	_, err = rstore.Info(*model.TestSongErr(t))
	assert.Error(t, err)

}
