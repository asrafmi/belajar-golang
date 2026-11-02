package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed asraf.jpg
var photo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("asraf-v2.jpg", photo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

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
