package main

import (
	"errors"
	"fmt"
	"image/gif"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const API_URI string = "https://api.gifs.com/media/upload"
const API_KEY string = "API_KEY"
const USER_GIFS_DIR string = "USER_GIFS_DIR"

func setConfig() {
	var userDir string
	for {
		fmt.Print("\nEnter ./path/to/gifs/directory: ")
		fmt.Scanf("%s\n", &userDir)

		if len(userDir) > 3 {
			break
		}
	}
	os.Setenv(USER_GIFS_DIR, userDir)

	os.Setenv(API_KEY, key)
	fmt.Printf("Key set to: %s \n", key)
}

func getGifFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	gifs := []string{}

	for _, file := range files {

		fileName := file.Name()
		ext := filepath.Ext(fileName)

		if ext == ".gif" {
			gifs = append(gifs, dir+fileName)
		}
	}

	if len(gifs) < 1 {
		return nil, errors.New("No .gifs available in dir")
	}

	return gifs, nil
}

func decodeGif(path string) (g *gif.GIF, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	g, err = gif.DecodeAll(file)
	if err != nil {
		return
	}

	return
}

func encodePingPong(g *gif.GIF) *gif.GIF {
	for i := len(g.Image) - 1; i >= 0; i-- {
		g.Image = append(g.Image, g.Image[i])
		g.Delay = append(g.Delay, g.Delay[i])
		g.Disposal = append(g.Disposal, byte(g.Disposal[i]))
	}

	return g
}
