package manager

import (
	"fmt"
	"strings"


	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	"github.com/TobiasGleiter/go-ai/pkg/agents"
	"github.com/TobiasGleiter/go-ai/pkg/tools/filesystem"
	"github.com/TobiasGleiter/go-ai/types"
)

type Manager struct {
	Agents []agents.Agent
}

func NewManager() Manager {
	return Manager{
		Agents: make([]agents.Agent, 0),
	}
}

func (m *Manager) AddAgent(agent agents.Agent) {
	m.Agents = append(m.Agents, agent)
}

func (m *Manager) Go() {
	for _, agent := range m.Agents {
		var userMessages []types.Message

		for _, task := range agent.Tasks {
			//ollama.Talk(agent, task)
			// llmEndpoint := "http://localhost:11434/api/generate"
			// ollama.TalkFast(llmEndpoint, agent, task)

			var err error
            userMessages, err = ollama.Chat(agent, task, userMessages)
            if err != nil {
                fmt.Println("Error:", err)
                return
            }

		}

		var assistantContent []string
		for _, msg := range userMessages {
			if msg.Role == "assistant" {
				assistantContent = append(assistantContent, msg.Content)
			}
		}

		// Join assistant content into a single string
		assistantContentString := strings.Join(assistantContent, "\n")

		// Write assistant content to file
		err := filesystem.CreateMarkdownFile(assistantContentString)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

