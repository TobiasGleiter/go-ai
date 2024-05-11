package main

import (
	"go-ai/types"
	"go-ai/pkg/agents"
	"go-ai/pkg/manager"
)




func main() {
    jobDescription := `You are a senior backend-developer at microsoft.`

	backendDeveloper := agents.NewAgent(jobDescription)
	task := types.Task{Description: "You are an expert in Go programming. Your task is to implement the backend-functionality for a second hand clothes web shop. I will provide what you should implement in the next prompt."}
	task1 := types.Task{Description: "Read all articles from the MonogDB database."}
	task2 := types.Task{Description: "Implement an API-Endpoint to send the articles to the client, use MongoDB."}
	task3 := types.Task{Description: "Create an API-Documenation for the articles api-endpoint"}

	backendDeveloper.AssignTask(task)
	backendDeveloper.AssignTask(task1)
	backendDeveloper.AssignTask(task2)
	backendDeveloper.AssignTask(task3)

	manager := manager.NewManager()
	manager.AddAgent(backendDeveloper)

	manager.Go()
}
