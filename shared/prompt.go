package shared

import "github.com/manifoldco/promptui"

// TODO: Bulderにする？
func NewPromptui(label string) promptui.Prompt {
	return promptui.Prompt{
		Label: label,
	}
}
