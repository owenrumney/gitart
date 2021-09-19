package main

import (
	"fmt"
	"os"
)

func main() {

	artHome, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		artHome = os.Args[1]
	}

	if err := NewArtist(artHome).GenerateArt(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
