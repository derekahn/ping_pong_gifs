package main

import (
	"log"
	"net/http"
)

const dir string = "gifs/"

func main() {

	setApiKey()

	http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))
}
