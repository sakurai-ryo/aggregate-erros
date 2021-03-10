package save

import (
	"aggregate_errors/shared"
	"aggregate_errors/typefile"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

// Save ... save error and context interactively
func Save() error {
	// pipeStr, err := readPipe()
	// if err != nil {
	// 	return err
	// }
	// log.Println("Error: ", pipeStr)

	builder := typefile.NewAggregated("")

	if err := build("Environment", &builder); err != nil {
		return err
	}
	if err := build("Description", &builder); err != nil {
		return err
	}
	if err := build("Solution", &builder); err != nil {
		return err
	}
	if err := build("Fixed", &builder); err != nil {
		return err
	}
	result := builder.Build()
	log.Print(result)

	return nil
}

func readPipe() (string, error) {
	read, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", nil
	}
	return string(read), nil
}

func readPrompt(label string) (string, error) {
	prompt := shared.NewPromptui(label)
	input, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return input, nil
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
