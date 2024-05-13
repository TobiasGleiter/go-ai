package ollama

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"io"

	"github.com/TobiasGleiter/go-ai/types"
)

func GenerateSingleStream(prompt string) {
	client := &http.Client{}
	llmEndpoint := "http://localhost:11434/api/generate"

	request := types.OllamaModel{
        Model:  "llama3:8b",
        Prompt: prompt,
        Options: struct {
            NumCtx int `json:"num_ctx,omitempty"`
        }{
            NumCtx: 4096,
        },
        Stream: true,
    }

	fmt.Println("Prompt: " + Yellow + prompt + Reset)

    requestBody, err := json.Marshal(request)
    if err != nil {
        fmt.Println("Error marshaling request:", err)
        return
    }

	req, err := http.NewRequest("POST", llmEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		fmt.Println("create request failed:", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("create request failed:", err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	for {
		var response types.OllamaStreamResponse
		if err := decoder.Decode(&response); err != nil {
			if err == io.EOF {
				fmt.Println(Red + "EOF" + Reset)
				break
			}
			fmt.Println("error decoding response:", err)
			break
		}

		if response.Done {
			fmt.Println(Purple + "\n Done!" + Reset )
            break
		}

		fmt.Print(Green + response.Response + Reset)
	}
}
