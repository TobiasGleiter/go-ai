package tasks

import (
	"fmt"

	"github.com/TobiasGleiter/go-ai/types"
)

func ResearchWhyTheSkyIsBlue() types.Task {
	return types.Task{
		Description:    fmt.Sprintf("Why is the sky blue?"),
	}
}