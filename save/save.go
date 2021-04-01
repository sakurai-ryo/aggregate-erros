package save

import (
	"aggregate_errors/shared"
	"aggregate_errors/typefile"
	"fmt"
	"log"
)

var env string
var desc string
var solution string
var fixed bool

const (
	Environment = "Environment"
	Description = "Description"
	Solution    = "Solution"
	Fixed       = "Fixed"
)

var S = []string{"Environment", "Description", "Solution", "Fixed"}

// Save ... save error and context interactively
func Save() error {
	builder := typefile.NewAggregated("")

	for _, s := range S {
		if err := build(s, &builder); err != nil {
			return err
		}
	}

	result := builder.Build()
	log.Print(result)

	return nil
}

func build(label string, builder *typefile.AggregateBulder) error {
	res, err := readPrompt(label)
	if err != nil {
		return err
	}
	b := *builder
	if label == Environment { // 拡張性低い
		b.WithEnvironment(res)
		return nil
	}
	if label == Description {
		b.WithDescription(res)
		return nil
	}
	if label == Solution {
		b.WithSolution(res)
		return nil
	}
	if label == Fixed {
		b.WithFixed(res)
		return nil
	}
	return fmt.Errorf("Internal ServerError: Label not found")
}

func readPrompt(label string) (string, error) {
	prompt := shared.NewPromptui(label)
	input, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return input, nil
}
