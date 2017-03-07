package main

import (
	"log"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)

		go func(file string) {
			g, err := decodeGif(file)
			if err != nil {
				log.Printf("\033[31m[ERROR]\033[0m %s\n", err.Error())
			}

			newGif := encodePingPong(g)
			if err := saveGif(file, newGif); err != nil {
				log.Printf("\033[31m[ERROR]\033[0m %s\n", err.Error())
			}
		}(file)

	}
	wg.Wait()
}
