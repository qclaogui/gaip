// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var contentType = "application/json"

var (
	serverAddr = flag.String("addr", "http://localhost:8080", "HTTP gateway url, e.g. http://localhost:8080")
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	flag.Parse()

	// Call Create
	id := createTodo()

	// Call Get
	todo := getTodo(id)

	// Call Update
	updateTodo(todo)

	// Call List
	listTodo()

	// Call Delete
	deleteTodo(id)
}

func CallHTTP(method string, endpoint string, body io.Reader) []byte {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		log.Fatalf("error constructing HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("HTTP request error: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to call endpoint: %s resp: %v", endpoint, resp)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading HTTP body: %v", err)
	}
	return respData
}

func createTodo() string {
	reqMethod := http.MethodPost

	reqURL := *serverAddr + "/v1/todos"
	reqBody := strings.NewReader(fmt.Sprintf(`{
		"api":"v1",
		"item": {
			"title":"title",
			"description":"description",
			"created_at":"%s"
		}
	}`, time.Now().In(time.UTC).Format(time.RFC3339Nano)))

	respBytes := CallHTTP(reqMethod, reqURL, reqBody)
	log.Printf("Create respBody=%s\n\n", respBytes)

	var createResp struct {
		API string `json:"api"`
		ID  string `json:"id"`
	}
	if err := json.Unmarshal(respBytes, &createResp); err != nil {
		log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
	}

	return createResp.ID
}

func deleteTodo(id string) {
	reqMethod := http.MethodDelete
	reqURL := *serverAddr + "/v1/todos/" + id

	respBytes := CallHTTP(reqMethod, reqURL, nil)
	log.Printf("Delete respBody=%s\n\n", respBytes)
}

func listTodo() {
	reqMethod := http.MethodGet
	reqURL := *serverAddr + "/v1/todos"

	respBytes := CallHTTP(reqMethod, reqURL, nil)
	log.Printf("List respBody=%s\n\n", respBytes)
}

func getTodo(id string) *Todo {
	reqMethod := http.MethodGet
	reqURL := *serverAddr + "/v1/todos/" + id

	respBytes := CallHTTP(reqMethod, reqURL, nil)

	log.Printf("Get respBody=%s\n\n", respBytes)

	var GetResp struct {
		API  string `json:"api"`
		Item *Todo  `json:"item"`
	}
	if err := json.Unmarshal(respBytes, &GetResp); err != nil {
		log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
	}
	return GetResp.Item
}

func updateTodo(todo *Todo) {
	reqMethod := http.MethodPost
	reqURL := *serverAddr + "/v1/todos/" + todo.ID

	reqBody := strings.NewReader(fmt.Sprintf(`
	{
		"api":"v1",
		"item": {
			"id":"%s",
			"title":"%s",
			"description":"%s updated",
			"updated_at":"%s"
		}
	}`, todo.ID, todo.Title, todo.Description, time.Now().In(time.UTC).Format(time.RFC3339Nano)))

	respBytes := CallHTTP(reqMethod, reqURL, reqBody)

	log.Printf("Update respBody=%s\n\n", respBytes)
}
