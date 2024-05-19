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
	Write a Go function that:
	* Handles incoming HTTP requests.
	* Connects to the MongoDB database using Sabrina's setup.
	* Executes the query defined in step 5.
	* Serializes the result into JSON format.
	* Returns the JSON response.`


	// --- INTELLIGENCE ----
	intelligence01 := " Recognize the problem and the need to make a decision."
	intelligence02 := " What are the most important decision you need to make? Output a numbered list."

	// Memory, co-workers, tools (probably a vector databse?)
	intelligence03 := " Gather information about the problem situation. Output a numbered list."
	intelligence04 := intelligence03 + agents + " Define which agent can help you. For example: 1. Mark can help define requirement."

	intel01 := ollama.Generate(agent01 + todo + intelligence01)
	intel02 := ollama.Generate(agent01 + todo + intel01 + intelligence02)
	intel03 := ollama.Generate(agent01 + todo + intel02 + intelligence03)
	res := ollama.Generate(agent01 + todo + intel03 + intelligence04)
	
	filesystem.CreateMarkdownFile(res, "intelligence_steps.md")
	
	// --- DESIGN ---
	// complex and continuous task, often being partly iterative,
	// where phases can overlap, and the decision-maker might revisit previous stages
	design01 := " Systematically structure the problem. Think in steps."
	design02 := " Establish specific criteria. Output a bullet list."
	design03 := " Identify a range of alternatives aimed at resolving the issue at hand. Output a numbered list."
	// 
	res01 := ollama.Generate(agent01 + todo + res + design01)
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