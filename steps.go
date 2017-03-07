package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image/gif"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type gifsResponse struct {
	Success struct {
		Page string `json:"page"`
	} `json:"success"`
}

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
			gifs = append(gifs, pathToUserGifs()+fileName)
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

func encodePingPong(g *gif.GIF) *gif.GIF {
	for i := len(g.Image) - 1; i >= 0; i-- {
		g.Image = append(g.Image, g.Image[i])
		g.Delay = append(g.Delay, g.Delay[i])
		g.Disposal = append(g.Disposal, byte(g.Disposal[i]))
	}

	return g
}

func saveGif(name string, g *gif.GIF) (fileName string, err error) {
	fileName = GIFS_DIR + rmPathToUserGifs(name)
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	if err = gif.EncodeAll(file, g); err != nil {
		return
	}

	return
}

func createFileUploadRequest(path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Hardcoding param names `file` & `title` for gif file;
	// It's required with gifs.com API `POST /media/upload`
	writer.WriteField("title", createTitle(file.Name()))
	part, err := writer.CreateFormFile("file", path)
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(part, file); err != nil {
		return nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", API_URI, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("gifs-api-key", os.Getenv(API_KEY))

	return req, err
}

func upload(req *http.Request) (success gifsResponse, err error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &success); err != nil {
		return
	}

	return
}
