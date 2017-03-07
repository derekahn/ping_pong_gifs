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

	log.Print("files", files)

	http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))
}
