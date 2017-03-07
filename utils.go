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

	var key string
	fmt.Print("\nEnter your gifs.com API key (ie. gifs58xxce10ad223): ")
	fmt.Scanf("%s\n", &key)

	if strings.Contains(key, "gifs") {
		os.Setenv(API_KEY, key)
		fmt.Printf("Key set to: %s \n", key)
	} else {
		fmt.Print("No key set! Won't be uploading to your dasahboard. \n")
	}
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

// Gets the USER_GIFS_DIR env var and appends a "/" if necessary
func pathToUserGifs() (path string) {
	path = os.Getenv(USER_GIFS_DIR)
	if !strings.HasSuffix(path, "/") {
		return path + "/"
	}
	return
}
