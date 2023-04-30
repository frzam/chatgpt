package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const URL = "https://api.openai.com/v1/chat/completions"
const MODEL = "gpt-3.5-turbo"
const ROLE = "user"

type ChatReq struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	key := os.Getenv("OPENAI_API_KEY")
	token := "Bearer " + key
	fmt.Println("token: ", token)

	msg := Message{
		Role:    ROLE,
		Content: "hello, how are you?",
	}

	chatReq := ChatReq{
		Model:    MODEL,
		Messages: []Message{msg},
	}
	client := &http.Client{}
	fmt.Printf("ChatReq: %s\n", chatReq)

	req, err := json.Marshal(chatReq)
	if err != nil {
		fmt.Println("Error while marshal: ", err)
		return
	}
	fmt.Println("Req: ", string(req))
	r, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(req))
	if err != nil {
		fmt.Println("Error while callig api: ", err)
		return
	}
	r.Header.Add("Authorization", token)
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("resp error: ", err)
		return
	}
	fmt.Printf("Resp: %+v\n", resp)
	defer resp.Body.Close()

	out, _ := io.ReadAll(resp.Body)
	fmt.Printf("Body: %+v", string(out))
}
