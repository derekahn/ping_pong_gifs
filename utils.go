package main

import (
	"os"
	"path/filepath"
	"strings"
)

func rmPathToUserGifs(fileName string) (newName string) {
	newName = strings.TrimPrefix(fileName, pathToUserGifs())
	return
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
