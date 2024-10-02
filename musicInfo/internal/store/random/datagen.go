package random

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func genReleaseDate() string {
	min := time.Date(1980, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min

	return time.Unix(sec, 0).Format("02.01.2006")
}

func genText() string {
	words := []string{"do", "re", "mi", "fa", "so", "la", "si"}
	lyrics := ""
	lines := 12

	for i := 1; i <= lines; i++ {
		for j := 0; j < 4+r.Intn(4); j++ {
			lyrics += words[r.Intn(len(words))] + " "
		}
		if i != lines {
			lyrics += "\n"
			if i%4 == 0 {
				lyrics += "\n"
			}
		}
	}
	return lyrics
}

func getLink() string {
	sources := []string{"youtube.com", "music.apple.com", "vk.com", "music.yandex.ru"}
	return fmt.Sprintf("http://%s/song%d", sources[r.Intn(len(sources))], r.Intn(9999))

}
