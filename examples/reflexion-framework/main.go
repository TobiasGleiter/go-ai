package main

import (
	// "fmt"
	// "regexp"
    // "strconv"

	"github.com/TobiasGleiter/go-ai/pkg/ollama"
	//"github.com/TobiasGleiter/go-ai/pkg/agents"
	"github.com/TobiasGleiter/go-ai/types"
)

type longTermMemory struct {
	VerbalReflectionText string
}

func main() {
	prompt := "Marty has 100 centimeters of ribbon that he must cut into 4 equal parts. Each of the cut parts must be divided into 5 equal parts. How long will each final cut be?"
	
	// Streams the response and output in the console.
	// res := ollama.Generate(prompt)
	// fmt.Println(res)

	// actor := agents.NewAgent("I am the actor. My task is to generate content.")
	// evaluator := agents.NewAgent("I am the evaluator. My task is to evaluate the content by the actor.")
	
	initialTrajectory := types.Task{Description: prompt}
	res := ollama.Generate(initialTrajectory.Description)
	answerFromActor := "Annswer from Actor: "+ res

	evaluateInputTask := types.Task{Description: "Evaluate the answer. Output {{TRUE}} for a correct answer and {{FALSE}} for a wrong answer."}
	ollama.Generate(evaluateInputTask.Description + answerFromActor)

	selfRefelctionTask := types.Task{Description: "Create a verbal self-reflection to solve similar tasks. Think step by step."}
	ollama.Generate(selfRefelctionTask.Description + answerFromActor)
	//selfRefectionRes := 
	//memory := []longTermMemory{{VerbalReflectionText: selfRefectionRes}}

	// t := 0
	// maxTrials := 2
	// for !EvaluatorPasses() && t < maxTrials {
	// 	promptWithLongTermMemory := prompt
	// 	trajectory := types.Task{Description: promptWithLongTermMemory}
	// 	res := ollama.Generate(trajectory.Description)
	// 	answerFromActor := "Annswer from Actor: "+ res

	// 	evaluateInputTask := types.Task{Description: "Evaluate the answer. Give a Score from bad to good (0-9) as an answer. Wrap the score in curly braces fr example: score is {{5}} of 10"}
	// 	evaluatorRes := ollama.Generate(evaluateInputTask.Description + answerFromActor)
	// 	answerFromEvaluator := "Answer from the Evaluator: "+ evaluatorRes

	// 	selfRefelctionTask := types.Task{Description: "Create a verbal self-reflection to solve similar tasks:"}
	// 	ollama.Generate(selfRefelctionTask.Description + answerFromEvaluator)

	// 	t++
	// }
}

func EvaluatorPasses() bool {
	// Evaluator pass condition logic here
	// For simplicity, always return false in this example
	return false
}