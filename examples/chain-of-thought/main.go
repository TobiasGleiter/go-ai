package main

import (
	"fmt"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	"github.com/TobiasGleiter/go-ai/types"
	"github.com/TobiasGleiter/go-ai/pkg/agents"
)




func main() {
	var userMessages []types.Message

	// Zero-shot CoT
	agent := agents.NewAgent("Riddle-Solver")
	task := types.Task{Description: "Marty has 100 centimeters of ribbon that he must cut into 4 equal parts. Each of the cut parts must be divided into 5 equal parts. How long will each final cut be?"}

	var err error
	userMessages, err = ollama.Chat(agent, task, userMessages)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
