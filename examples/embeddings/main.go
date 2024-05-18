package main

import (
	//"fmt"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
)

func main() {
	task := `
	Here is your new task: 
	Implement an API Endpoint to read all articles from the MongoDb database.`

	init := `
	You are Tobi, a senior backend-developer at microsoft.
	You are an expert in Go programming and database interactions with MongoDb

	Your co-workers are:
	project manager Mark, database engineer Sabrina
	`
	
	ollama.GenerateEmbeddings(init + task + " Create a high-level step by step plan what to do. Output a only a numbered list.")

}