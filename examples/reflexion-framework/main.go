package main

import (
	"fmt"
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
	//prompt := "Your task is to plan a second hand webshop from the technical perspective"
	// Streams the response and output in the console.
	// res := ollama.Generate(prompt)
	// fmt.Println(res)

	// actor := agents.NewAgent("I am the actor. My task is to generate content.")
	// evaluator := agents.NewAgent("I am the evaluator. My task is to evaluate the content by the actor.")
	
	initialTrajectory := types.Task{Description: prompt}
	res := ollama.Generate(initialTrajectory.Description)
	answerFromActor := "\nAnswer from Actor: "+ res

	evaluateInputTask := types.Task{Description: "Evaluate the answer with a score from one to ten (ten is good). No emotions."}
	answerFromEvaluator := ollama.Generate(evaluateInputTask.Description + answerFromActor)

	selfRefelctionTask := types.Task{Description: "Create a precise short verbal self-reflection how to solve the task."}
	selfRefectionRes := ollama.Generate(selfRefelctionTask.Description + answerFromEvaluator)
	//
	var memory []string
	memory = append(memory, selfRefectionRes)

	t := 0
	maxTrials := 2
	for !EvaluatorPasses() && t < maxTrials {
		fmt.Println(t)
		promptWithLongTermMemory := prompt
		trajectory := types.Task{Description: memory[t] + promptWithLongTermMemory }
		res := ollama.Generate(trajectory.Description)
		answerFromActor := "\nAnswer from Actor: "+ res

		evaluatorRes := ollama.Generate(evaluateInputTask.Description + answerFromActor)
		answerFromEvaluator := "Answer from the Evaluator: "+ evaluatorRes

		selfReflectionResLoop := ollama.Generate(selfRefelctionTask.Description + answerFromEvaluator)
		memory = append(memory, selfReflectionResLoop)

		t++
	}
}

func EvaluatorPasses() bool {
	// Evaluator pass condition logic here
	// For simplicity, always return false in this example
	return false
}