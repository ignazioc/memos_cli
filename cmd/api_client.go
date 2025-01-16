package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "https://<>/api/v1/memos"
const token = ""

func fetchMemos() []Memo {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(resp.Body)

	var memosResponse MemosResponse
	if err := json.Unmarshal(body, &memosResponse); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	var memos = memosResponse.Memos

	defer resp.Body.Close()
	return memos
}

func post(content string) {

	values := map[string]string{"content": content, "visibility": "PUBLIC"}
	jsonValue, _ := json.Marshal(values)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6InYxIiwidHlwIjoiSldUIn0.eyJuYW1lIjoiaWduYXppb2MiLCJpc3MiOiJtZW1vcyIsInN1YiI6IjEiLCJhdWQiOlsidXNlci5hY2Nlc3MtdG9rZW4iXSwiaWF0IjoxNzM2OTI1ODU1fQ.XUrIfdS6QiTDMNZ2yz5DwtV6_ISpHTagiPk6ztnhkcM")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer res.Body.Close()
}
