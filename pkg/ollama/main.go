package ollama

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
	"io"

	"github.com/TobiasGleiter/go-ai/pkg/agents"
	"github.com/TobiasGleiter/go-ai/types"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func Talk(agent agents.Agent, task types.Task) string {
	llmEndpoint := "http://localhost:11434/api/generate"
	client := &http.Client{}

	prompt := fmt.Sprintf("Agent description: %s; Task: %s;", agent.Role, task.Description)
	fmt.Println("Prompt: " + Yellow + prompt + Reset)

	requestBody := fmt.Sprintf(`{
		"model": "llama3:8b",
		"prompt": "%s",
		"stream": false
	}`, prompt)

	req, err := http.NewRequest("POST", llmEndpoint, strings.NewReader(requestBody))
	if err != nil {
		fmt.Println("create request failed:", err)
		return "Error"
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("create request failed:", err)
		return "Error"
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response types.OllamaResponse
	if err := decoder.Decode(&response); err != nil {	
		fmt.Println("error decoding response:", err)
	}
	fmt.Print(Green + response.Response + Reset)
	return response.Response
}


func TalkFast(llmEndpoint string, agent agents.Agent, task types.Task) {
	client := &http.Client{}
	prompt := fmt.Sprintf("Agent description: %s; Task: %s;", agent.Role, task.Description)
	fmt.Println("Prompt: " + Yellow + prompt + Reset)

	requestBody := fmt.Sprintf(`{
		"model": "llama3:8b",
		"prompt": "%s",
		"options": {"num_ctx": 4096},
		"stream": true
	}`, prompt)

	req, err := http.NewRequest("POST", llmEndpoint, strings.NewReader(requestBody))
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
				fmt.Println(Red + "This is red text" + task.Description + Reset)
				break
			}
			fmt.Println("error decoding response:", err)
			break
		}

		if response.Done {
			fmt.Println(Purple + "\n Done with: "+task.Description + Reset )
            break
		}

		fmt.Print(Green + response.Response + Reset)
	}
}

func Chat(agent agents.Agent, task types.Task, messages []types.Message) ([]types.Message, error) {
	llmEndpoint := "http://localhost:11434/api/chat"
	client := &http.Client{}

	prompt := fmt.Sprintf("Agent description: %s; Task: %s;", agent.Role, task.Description)
	fmt.Println("Prompt: " + Yellow + prompt + Reset)

	messages = append(messages, types.Message{
		Role:    "user",
		Content: prompt,
	})

	var messagesJSON []string
    for _, msg := range messages {
        msgJSON, err := json.Marshal(msg)
        if err != nil {
            fmt.Println("error marshaling message:", err)
            return []types.Message{}, err
        }
        messagesJSON = append(messagesJSON, string(msgJSON))
    }

	messagesJSONString := strings.Join(messagesJSON, ",")
	fmt.Println(messagesJSONString)

	requestBody := fmt.Sprintf(`{
        "model": "llama3:8b",
        "messages": [%s],
        "stream": false
    }`, messagesJSONString)

	req, err := http.NewRequest("POST", llmEndpoint, strings.NewReader(requestBody))
	if err != nil {
		fmt.Println("create request failed:", err)
		return []types.Message{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("create request failed:", err)
		return []types.Message{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response types.OllamaChatResponse
	if err := decoder.Decode(&response); err != nil {	
		fmt.Println("error decoding response:", err)
	}

	assistantMessage := types.Message{
        Role:    "assistant",
        Content: response.Message.Content,
    }
	messages = append(messages, assistantMessage)

	fmt.Print(Green + response.Message.Content + Reset)
	return messages, nil
}
