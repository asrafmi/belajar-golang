package test

import (
	"embed"
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

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

func TestEmbedMultiFiles(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))

	b, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

//go:embed files/*.txt
var path embed.FS

func TestEmbedPathMatcher(t *testing.T) {
	directories, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}
	for _, dir := range directories {
		fmt.Println("File => ", dir.Name())
		if !dir.IsDir() {
			file, err := path.ReadFile("files/" + dir.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println("Isinya => ", string(file))
		}
	}
}
