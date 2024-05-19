package main

import (
	//"fmt"
	//"regexp"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	"github.com/TobiasGleiter/go-ai/pkg/tools/filesystem"
)

func main() {	
	agent01 := `
	You are Tobi, a senior backend-developer at Google.
	You are an expert in Golang programming.
	`

	todo := `	
	Current Todo:
	Determine the necessary fields and conditions (if any) required to fetch all articles from the MongoDB database.
	`

	intelligence01 := " Recognize the problem and the need to make a decision. Only return high-level steps."
	highLevelPlan := ollama.Generate(agent01 + todo + intelligence01)
	filesystem.CreateMarkdownFile(highLevelPlan, "high-level-plan.md")

	design01 := " Establish specific criteria. Output a bullet list."
	// probably consider old decicions
	specificCriteria := ollama.Generate(agent01 + highLevelPlan + design01)
	filesystem.CreateMarkdownFile(specificCriteria, "specific-criteria.md")

	choice01 := " Summarize an optimal solution that aligns with the defined criteria and make the final decision."
	final := ollama.Generate(agent01 + highLevelPlan +  specificCriteria + choice01)
	filesystem.CreateMarkdownFile(final, "choice.md")

}