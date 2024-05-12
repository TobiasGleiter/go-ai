package main

import (
	"github.com/TobiasGleiter/go-ai/pkg/ollama"
)

func main() {
	prompt := "Marty has 100 centimeters of ribbon that he must cut into 4 equal parts. Each of the cut parts must be divided into 5 equal parts. How long will each final cut be?"

	// Streams the response and output in the console.
	ollama.GenerateSingleStream(prompt)
}