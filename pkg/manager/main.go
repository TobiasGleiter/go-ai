package manager

import (
	//"fmt"
	"regexp"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	"github.com/TobiasGleiter/go-ai/pkg/agents"
	//"github.com/TobiasGleiter/go-ai/pkg/tools/filesystem"
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

func removeQuotesAndSpecialChars(s string) string {
    // Define the regular expression pattern to match single and double quotation marks, along with any special characters
    reg := regexp.MustCompile(`[‘’'"\p{P}]`)
    // Replace matched characters with an empty string
    cleaned := reg.ReplaceAllString(s, "")
    return cleaned
}

func (m *Manager) plan() {
	ceoDescription := `Your are a CEO. You are talking with your CTO. Output only one sentence. Lets think step by step.`
	ctoDescription := `Your are a CTO. You are talking with your CEO Output only one sentence. Lets think step by step.`

	ceo := agents.NewAgent(ceoDescription)
	cto := agents.NewAgent(ctoDescription)
	ceoInit := types.Task{Description: "You are talking to your CTO about a new product. Lets think step by step." }
	//ctoRequestMessage := types.Task{Description: "Talk to your CTO about a new product. Only one sentence." }

	// ceoMessage := ollama.Talk(ceo, ceoInit)
	// fmt.Println(ceoMessage)

	// ctoMessage := ollama.Talk(cto, ctoRequestMessage)
	// fmt.Println(ctoMessage)
	
	ctoMessage := ollama.TalkAndForget(ceo, ceoInit)
	iterations := 1

    for i := 0; i < iterations; i++ {
		// CEO sends a message to CTO
		cleanedText := removeQuotesAndSpecialChars(ctoMessage)
		ctoTask := types.Task{Description: cleanedText}
		ctoMessage := ollama.TalkAndForget(cto, ctoTask)

		cleanedText = removeQuotesAndSpecialChars(ctoMessage)
		ceoTask := types.Task{Description: cleanedText}
		ceoMessage := ollama.TalkAndForget(ceo, ceoTask)
		
		ctoMessage = ceoMessage

    }

	return
}

func (m *Manager) Go() {
	// Manager Model starts planning and interates X times over the planning process.
	m.plan()

	// for _, agent := range m.Agents {
	// 	var userMessages []types.Message

	// 	for _, task := range agent.Tasks {
	// 		//ollama.Talk(agent, task)
	// 		// llmEndpoint := "http://localhost:11434/api/generate"
	// 		// ollama.TalkFast(llmEndpoint, agent, task)

	// 		var err error
    //         userMessages, err = ollama.Chat(agent, task, userMessages)
    //         if err != nil {
    //             fmt.Println("Error:", err)
    //             return
    //         }

	// 	}

	// 	var assistantContent []string
	// 	for _, msg := range userMessages {
	// 		if msg.Role == "assistant" {
	// 			assistantContent = append(assistantContent, msg.Content)
	// 		}
	// 	}

	// 	assistantContentString := strings.Join(assistantContent, "\n")
	// 	err := filesystem.CreateMarkdownFile(assistantContentString)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// }
}

