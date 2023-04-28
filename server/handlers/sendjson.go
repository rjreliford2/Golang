package handlers

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

// will take a url and file and will send as a json

func SendJson(url string, file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
