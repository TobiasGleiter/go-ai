package main

import (
	"fmt"
	"strings"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	"github.com/TobiasGleiter/go-ai/pkg/tools/filesystem"
)

func main() {
	var dialogMemory []string

	task := `
	Here is your new task: 
	Implement an API Endpoint to read all articles from the MongoDb database.`

	agent01 := 	`
	You are Tobi, a senior backend-developer at Google.
	You are an expert in Golang programming and database interactions with MongoDb

	Your co-workers are:
	database engineer Sabrina;
	`

	agent02 := 	`
	You are Sabrina, a senior database-engineer at Google.
	You are an expert in MongoDb.

	Your co-workers are:
	backend-developer Tobi;
	`

	// PLAN
	plan01 := ollama.Generate(agent01 + task + " Decide what your next step is. Output a numbered list.")

	// INIT DIALOG
	//thinkTobi := ""//ollama.Generate(agent01 + plan01 + " What are the three key aspects to tell Sabrina? Output a numbered list.")
	
	// goal is defined
	thinkTobi := ollama.Generate(agent01 + plan01 + "What would you say to Sabrina?")
	dialogTobi := ""

	fmt.Println("\n-------------------------------")
	// COMMUNICATION Process

	iterations := 6
    for i := 0; i < iterations; i++ {
        // Sabrina's turn
		fmt.Println("\nSabrina thinks...")
        thinkSabrina := ollama.Generate(agent02 + dialogTobi + " What is the one key aspect to tell Tobi?")
        dialogSabrina := ollama.Generate(agent02 + thinkSabrina + " Format this into: Hi Tobi, I would...")

		dialogMemory = append(dialogMemory, thinkSabrina)

		fmt.Println("\nTobi thinks...")
		thinkTobi = ollama.Generate(agent01 + dialogSabrina + " What is the one key aspect to tell Sabrina?")
		dialogTobi = ollama.Generate(agent01 + thinkTobi + " Format this into: Hi Sabrina, I would...")
        
		dialogMemory = append(dialogMemory, thinkTobi)

        fmt.Println("\nIteration", i+1)
		fmt.Println("-------------------------------")
    }

	// THINKING
	// think01 := ollama.Generate(agent01 + plan01 + " What are the three key aspects you need to know from Sabrina? Output a numbered list.")
	// dialogTobi01 := ollama.Generate(agent01 + think01 + " What would you say to Sabrina?")


	// thinkSabringa01 := ollama.Generate(agent02 + dialogTobi01 + " What are the three key aspects Tobi what to know? Output a numbered list.")
	// dialogSabrina01 := ollama.Generate(agent02 + thinkSabringa01 + " What would you say to Tobi?")


	// thinkTobi02 := ollama.Generate(agent01 + dialogSabrina01 + " What are the three key aspects you need to know from Sabrina? Output a numbered list.")
	// //dialogTobi02 := ollama.Generate(agent01 + thinkTobi02 + " What would you say to Sabrina?")

	//sumTobi01 := ollama.Generate(agent01 + thinkTobi + " Summarize this thoughts. Output a numbered list.")
	memoryString := strings.Join(dialogMemory, "\n")
	filesystem.CreateMarkdownFile(memoryString, "dialog_tobi_sabrina.md")

	
	//ollama.Generate(agent02 + task + " Decide what your next step is. Output a numbered list.")
	// execute plan
	// If question wait for sabrina
}