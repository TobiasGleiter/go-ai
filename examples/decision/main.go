package main

import (
	//"fmt"
	//"regexp"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	"github.com/TobiasGleiter/go-ai/pkg/tools/filesystem"
)

func main() {	
	//agents := "senior backen-developer Tobi; database-engineer Sabrina, project manager Mark;"

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

	design01 := " Establish specific criteria. Output a bullet list."
	specificCriteria := ollama.Generate(agent01 + highLevelPlan + design01)

	choice01 := " Select the most optimal solution that aligns with the defined criteria and make the final decision."
	final := ollama.Generate(agent01 + highLevelPlan +  specificCriteria + choice01)
	filesystem.CreateMarkdownFile(final, "choice.md")
	
}