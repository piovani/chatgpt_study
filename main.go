package main

/*
Documentation: https://platform.openai.com/docs/overview
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	apiToken := os.Getenv("GPT_TOKEN")

	input := "A NASA já foi quantas vezes para o espaço?"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4.1",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": input,
			},
		},
	})
	if err != nil {
		log.Fatal("Error marshalling request body:", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	fmt.Println(string(body))
}
