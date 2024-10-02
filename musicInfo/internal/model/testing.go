package model

import "testing"

// TestSongOK ...
func TestSongOK(t *testing.T) *Song {
	return &Song{
		Group: "Muse",
		Song:  "Supermassive Black Hole",
	}
}

// TestSongNil ...
func TestSongNil(t *testing.T) *Song {
	return nil
}

// TestSongWithout ...
func TestSongWithoutField(t *testing.T) *Song {
	return &Song{
		Song: "Supermassive Black Hole",
	}
}

// TestSongErr ...
func TestSongErr(t *testing.T) *Song {
	return &Song{
		Group: "error",
		Song:  "Supermassive Black Hole",
	}
}

// TestSongEmpty ...
func TestSongEmpty(t *testing.T) *Song {
	return &Song{
		Group: "empty",
		Song:  "Supermassive Black Hole",
	}
}
