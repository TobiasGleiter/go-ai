package agents

import (
	"go-ai/types"
)

type Agent struct {
	Role           string
	Tasks []types.Task
}

func NewAgent(role string) Agent {
	return Agent{
		Role: role,
		Tasks: make([]types.Task, 0),
	}
}

func (a *Agent) AssignTask(task types.Task) {
	// Here you can implement the logic to handle the assigned task,
	// such as adding it to a list of tasks the agent is responsible for.
	a.Tasks = append(a.Tasks, task)
}
