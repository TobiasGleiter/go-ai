package main

import (
	//"fmt"
	//"regexp"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	"github.com/TobiasGleiter/go-ai/pkg/tools/filesystem"
)

func main() {	
	agents := "senior backen-developer Tobi; database-engineer Sabrina, project manager Mark;"

	agent01 := `
	You are Tobi, a senior backend-developer at Google.
	You are an expert in Golang programming.
	`

	todo := `	
	Current Todo:
	Decide on a suitable endpoint URL (e.g., "/articles") and HTTP method (e.g., GET). 
	Determine the expected response format (JSON) and any necessary authentication or authorization mechanisms.`


	// --- INTELLIGENCE ----
	intelligence01 := " Recognize the problem and the need to make a decision. Only return high-level steps."
	highLevelPlan := ollama.Generate(agent01 + todo + intelligence01)
	filesystem.CreateMarkdownFile(highLevelPlan, "highLevelPlan.md")
	
	// TODO: Gather information from the memory or a vector database about the task.
	// private todo list:
	// What are my tasks?
	intelligence03 := " Define tasks that You need to do to solve this. Output a numbered list."
	privateTodos := ollama.Generate(agent01 + todo + highLevelPlan + intelligence03)
	filesystem.CreateMarkdownFile(privateTodos, "private_todo.md")
	
	// public todo list:
	intelligence04 := " Define which agent can help You. For example: 1. Mark can help define requirement. Output a numbered list of the agents and the request."
	publicTodos := ollama.Generate(agent01 + todo + highLevelPlan + agents + intelligence04)
	filesystem.CreateMarkdownFile(publicTodos, "public_todo.md")
	
	
	// --- DESIGN ---
	// complex and continuous task, often being partly iterative,
	// where phases can overlap, and the decision-maker might revisit previous stages
	design01 := " Systematically structure the problem. Think in steps."
	design02 := " Establish specific criteria. Output a bullet list."
	design03 := " Identify a range of alternatives aimed at resolving the issue at hand. Output a numbered list."
	// 
	res01 := ollama.Generate(agent01 + todo + privateTodos + design01)
	res02 := ollama.Generate(agent01 + res01 + design02)
	res03 := ollama.Generate(agent01 + res01 + res02 + design03)
	filesystem.CreateMarkdownFile(res03, "design.md")

	// --- CHOICE ---
	// Store as Refelction
	choice01 := " Select the most optimal alternative that aligns with the defined criteria and make the final decision. Output a numbered list."
	//choice02 := " Summarize the final decision short an precise. Output a numbered list with the most important decisions."

	final := ollama.Generate(agent01 + res03 + choice01)
	//final := ollama.Generate(agent01 + choice + choice02)
	filesystem.CreateMarkdownFile(final, "choice.md")
	
	//ollama.Generate(res01 + " Classify the text from above to [ACT] (do something) or [PLAN] (create a plan) or [REFELCT] (reflect on the task). Output only one of these: for example: [REFLECT]")
	// re := regexp.MustCompile(`\[(.*?)\]`)
	// match := re.FindStringSubmatch(res02)
	// if len(match) > 1 {
	// 	keyword := match[1]

	// 	switch keyword {
	// 	case "ACT":
	// 		ollama.Generate(todo + " Implement this todo.")
	// 	case "PLAN":
	// 		ollama.Generate(res01 + " Create a technical todo list what you need to do.")
	// 	case "REFELCT":
	// 		ollama.Generate(res01 + " Summarize.")
	// 	default:
	// 		fmt.Println("No valid action found for the keyword")
	// 	}
	// } else {
	// 	fmt.Println("No match found")
	// }
}