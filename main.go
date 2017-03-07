package main

import (
	"log"
	"net/http"
	"sync"
)

const GIFS_DIR string = "gifs/"

func main() {

	setConfig()

	files, err := getGifFiles(pathToUserGifs())
	checkFor(err)

	half := len(files) / 2
	firstHalf := files[:half]
	otherHalf := files[half:]

	go uploadPingPongs(firstHalf)
	go uploadPingPongs(otherHalf)

	http.ListenAndServe(":8080", http.FileServer(http.Dir(GIFS_DIR)))
}

func uploadPingPongs(files []string) {
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)

		go func(file string) {
			g, err := decodeGif(file)
			if err != nil {
				log.Printf("\033[31m[ERROR]\033[0m %s\n", err.Error())
			}

			newGif := encodePingPong(g)
			savedFile, err := saveGif(file, newGif)
			checkFor(err)

			request, err := createFileUploadRequest(savedFile)
			checkFor(err)

		}(file)

	}
	wg.Wait()
}
