package filesystem

import (
	"os"
    "fmt"
)

func CreateMarkdownFile(input string) error {
    // Create the file
    filePath := "ai_generated/complete_chat" + ".md"
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    // Write Go code to the file
    _, err = file.WriteString(input)
    if err != nil {
        return err
    }

    fmt.Printf("Created file: %s\n", filePath)
    return nil
}