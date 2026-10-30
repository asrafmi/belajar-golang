package belajar_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed asraf.jpg
var photo []byte

func TestEmbedFile(t *testing.T) {
	err := os.WriteFile("asraf-v2.jpg", photo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
