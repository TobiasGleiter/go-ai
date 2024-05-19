package main

import (
	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	"github.com/TobiasGleiter/go-ai/pkg/tools/filesystem"
)

func main() {

	task := `
	Here is your new task: 
	Implement an API Endpoint to read all articles from the MongoDb database.`

	agent01 := 	`
	You are Tobi, a senior backend-developer at Google.
	You are an expert in Golang programming and database interactions with MongoDb

	Your co-workers are:
	database engineer Sabrina;

	Your boss is:
	project manager Mark;
	`

	agent02 := 	`
	You are Sabrina, a senior database-engineer at Google.
	You are an expert in MongoDb.

	Your co-workers are:
	backend-developer Tobi;

	Your boss is:
	project manager Mark;
	`

	// PLAN
	plan01 := ollama.Generate(agent01 + task + " Decide what your next step is. Output a numbered list.")

	// EXECUTE PLAN

	// QUEST
	// THINKING
	think01 := ollama.Generate(agent01 + plan01 + " What are the three key aspects you need to know from Sabrina? Output a numbered list.")
	// memoize what i need from sabrina
	dialogTobi01 := ollama.Generate(agent01 + think01 + " What would you say to Sabrina?")




	thinkSabringa01 := ollama.Generate(agent02 + dialogTobi01 + " What are the three key aspects Tobi what to know? Output a numbered list.")
	// memoize what i need from sabrina
	dialogSabrina01 := ollama.Generate(agent02 + thinkSabringa01 + " What would you say to Tobi?")


	thinkTobi02 := ollama.Generate(agent01 + dialogSabrina01 + " What are the three key aspects you need to know from Sabrina? Output a numbered list.")
	//dialogTobi02 := ollama.Generate(agent01 + thinkTobi02 + " What would you say to Sabrina?")

	sumTobi01 := ollama.Generate(agent01 + thinkTobi02 + " Summarize this thoughts. Output a numbered list.")
	

	filesystem.CreateMarkdownFile(sumTobi01, "sum_tobi_01.md")
	//ollama.Generate(agent02 + task + " Decide what your next step is. Output a numbered list.")
	// execute plan
	// If question wait for sabrina
}