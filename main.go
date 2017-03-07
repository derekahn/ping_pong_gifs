package main

import (
	"log"
	"net/http"
)

const dir string = "gifs/"

func main() {

	setApiKey()

	files, err := getGifFiles(dir)
	if err != nil {
		log.Printf("\033[31m[ERROR]\033[0m %s\n", err.Error())
	}

	half := len(files) / 2
	firstHalf := files[:half]
	otherHalf := files[half:]

	go uploadPingPongs(firstHalf)
	go uploadPingPongs(otherHalf)

	http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))
}

func uploadPingPongs(files []string) {
}
