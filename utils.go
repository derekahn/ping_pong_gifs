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

// Formats "/some/path/to/crazy_funny_cool.gif" to "Crazy Funny Cool"
func createTitle(fileName string) string {
	name := trimFileName(fileName)
	return strings.Title(strings.Join(strings.Split(name, "_"), " "))
}

// Gets the USER_GIFS_DIR env var and appends a "/" if necessary
func pathToUserGifs() (path string) {
	path = os.Getenv(USER_GIFS_DIR)
	if !strings.HasSuffix(path, "/") {
		return path + "/"
	}
	return
}
