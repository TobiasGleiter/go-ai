package tasks

import (
	"fmt"

	"go-ai/types"
)

func ResearchWhyTheSkyIsBlue() types.Task {
	return types.Task{
		Description:    fmt.Sprintf("Why is the sky blue?"),
	}
}