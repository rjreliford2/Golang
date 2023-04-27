package logic

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Z01() {
	fmt.Println("z01")
	url := "http://localhost:1323/"
	data := []byte(`{"param1":"value1","param2":"value2"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating API request:", err)
		return
	}

	//Adds any headers to the reqeust
	req.Header.Set("Content-Type", "application/json")

	//Send API request and gets the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making API call:", err)
		return
	}
	defer resp.Body.Close()
	//Reads the response and prints it to the console
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading API response:", err)
		return
	}
	fmt.Println("API response:", string(body))
}
