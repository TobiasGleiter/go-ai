package main

import (
	"encoding/json"
	"strings"
	"log"
	"os"
	"time"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
)

type Task struct {
	Description string `json:"description"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func main() {
	var memory []string


	task := `
	Here is your new task: 
	Implement an API Endpoint to read all articles from the MongoDb database.`

	init := `
	You are Tobi, a senior backend-developer at microsoft.
	You are an expert in Go programming and database interactions with MongoDb

	Your co-workers are:
	project manager Mark, database engineer Sabrina
	`
	// PLAN -> memory
	plan01 := ollama.Generate(init + task + " Create a high-level plan that you know as senior backend-developer what to do. Output only a numbered list.")
	memory = append(memory, plan01)
	//plan02 := ollama.Generate(plan01 + " Refine this plan and make it realy precise. Output a list in this format: Todo: Review and Planning, Tasks: 1. Review the project require...")
	data := ollama.GenerateJson(plan01 + ` Format this well structured plan. Output json: 
	{
		"tasks":
			[
				{"description": ""}
			]
	}`)

	var tasks Tasks
	err := json.Unmarshal([]byte(data), &tasks)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	for _, task := range tasks.Tasks {
		memory = append(memory, task.Description)
	}

	memoryString := ""
	for _, task := range tasks.Tasks {
		// get memory-stream
		memoryString = strings.Join(memory, "\n")
		// execute tasks step by step.

		genTask := ollama.Generate(task.Description + "\n Decide what to do. Either create a code snippet, talk to project manager Mark, database engineer Sabrina or reflect what you did.")
		//genTask := ollama.Generate(memoryString + task.Description + "\n Output neccessary information.")
		// store information in the memory-stream
		memory = append(memory, genTask)
	}

	currentTime := time.Now()
	err = writeToFile("README" + currentTime.String() + ".md", memoryString)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}

	// extract tasks and ask llms.

	// // Generate Code
	//ollama.Generate(plan01 + " Turn this plan into code. Think in steps.")

	// // Validate the Code
	

	// // Summarize for the memory
	// sum01 := ollama.Generate(code01 + "Explain what you did. Output a numbered list.")
	// sum02 := ollama.Generate(sum01 + "Summarize what you did in one sentence.")

	//fmt.Println("\n" + plan02)
}

func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
