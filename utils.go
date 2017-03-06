package main

import (
	"fmt"
	"os"
)

const API_URI string = "https://api.gifs.com/media/upload"
const API_KEY string = "API_KEY"

func setApiKey() {
	var key string

	for {
		fmt.Print("\nEnter your gifs.com API key (ie. gifs58xxce10ad223): ")
		fmt.Scanf("%s\n", &key)

		if len(key) > 0 {
			break
		}
	}

	os.Setenv(API_KEY, key)
	fmt.Printf("Key set to: %s \n", key)
}
