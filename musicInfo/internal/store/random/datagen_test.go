package random

import (
	"fmt"
	"testing"
)

func TestGenText(t *testing.T) {
	test1 := genReleaseDate()
	test2 := genText()
	test3 := getLink()

	fmt.Println("|" + test1 + "|\n")
	fmt.Println("|" + test2 + "|\n")
	fmt.Println("|" + test3 + "|")
	t.Error()
}
