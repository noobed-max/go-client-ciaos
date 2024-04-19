package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Save(API_URL string, key string, value []string) error {
	data := make([]string, 0)
	if key != "" {
		data = append(data, "key="+key)
	}
	for _, v := range value {
		data = append(data, "encoded_content="+v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/upload/", API_URL), bytes.NewBufferString(strings.Join(data, "&")))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("error: %s", body)
	}

	return nil
}

func main() {
	API_URL := "http://127.0.0.1:8000"
	key := "example_key"
	value := []string{"encoded_content_1", "encoded_content_2"}

	err := Save(API_URL, key, value)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Update successful")
	}
}
