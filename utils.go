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

// RM dir path to user gifs & file extension
func trimFileName(fileName string) (newFileName string) {
	noPrefix := rmPathToUserGifs(fileName)
	newFileName = strings.TrimSuffix(noPrefix, filepath.Ext(fileName))
	return
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
