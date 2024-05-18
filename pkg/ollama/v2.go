package ollama

import (
	"net/http"
	"fmt"
	"strings"
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
        Options: types.OllamaOptions{
            NumCtx: 4096,
        },
        Stream: true,
    }

	fmt.Println("\nPrompt: " + Yellow + prompt + Reset)

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


func Generate(prompt string) string {
	client := &http.Client{}
	llmEndpoint := "http://localhost:11434/api/generate"
	
	request := types.OllamaModel{
        Model:  "llama3:8b",
        Prompt: prompt,
		Options: types.OllamaOptions{
            NumCtx: 4096,
        },
        Stream: false,
    }

	lastWords := lastNWords(prompt, 200)
	fmt.Println("\nPrompt: " + Yellow + lastWords + Reset)

    requestBody, err := json.Marshal(request)
    if err != nil {
        fmt.Println("Error marshaling request:", err)
        return "error"
    }

	req, err := http.NewRequest("POST", llmEndpoint, bytes.NewReader(requestBody))
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

	lastWords = lastNWords(response.Response, 400)
	fmt.Print(Green + lastWords + Reset)
	return response.Response
}

func GenerateJson(prompt string) string {
	client := &http.Client{}
	llmEndpoint := "http://localhost:11434/api/generate"
	
	request := types.OllamaModel{
        Model:  "llama3:8b",
        Prompt: prompt,
		Options: types.OllamaOptions{
            NumCtx: 4096,
        },
		Format: "json",
        Stream: false,
    }

	lastWords := lastNWords(prompt, 200)
	fmt.Println("\nPrompt: " + Yellow + lastWords + Reset)

    requestBody, err := json.Marshal(request)
    if err != nil {
        fmt.Println("Error marshaling request:", err)
        return "error"
    }

	req, err := http.NewRequest("POST", llmEndpoint, bytes.NewReader(requestBody))
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

	lastWords = lastNWords(response.Response, 400)
	fmt.Print(Green + lastWords + Reset)
	return response.Response
}

func GenerateEmbeddings(prompt string) types.OllamaEmbedding {
	client := &http.Client{}
	llmEndpoint := "http://localhost:11434/api/embeddings"
	
	request := types.OllamaModel{
		Model:   "llama3:8b",
		Prompt:  prompt,
		Options: types.OllamaOptions{NumCtx: 4096},
		Stream:  false,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		return types.OllamaEmbedding{} // Return an empty value or handle the error accordingly
	}

	req, err := http.NewRequest("POST", llmEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		fmt.Println("create request failed:", err)
		return types.OllamaEmbedding{} // Return an empty value or handle the error accordingly
	}
	req.Header.Add("Content-Type", "application/json")
		
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("create request failed:", err)
		return types.OllamaEmbedding{} // Return an empty value or handle the error accordingly
	}
	defer resp.Body.Close()

	var embedding types.OllamaEmbedding // Assuming types.OllamaEmbedding is the correct type for embedding
	err = json.NewDecoder(resp.Body).Decode(&embedding)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return types.OllamaEmbedding{} // Return an empty value or handle the error accordingly
	}

	fmt.Print(Green + fmt.Sprintf("%v", embedding.Embedding) + Reset) // Assuming embedding.Embedding is of type []float64
	return embedding
}


func lastNWords(s string, n int) string {
	words := strings.Fields(s) // Split the string into words
	if len(words) > n {
		words = words[len(words)-n:] // Take the last n words
	}
	return strings.Join(words, " ") // Join the words back into a string
}