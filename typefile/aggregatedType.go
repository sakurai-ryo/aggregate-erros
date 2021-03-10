package typefile

import "time"

type AggregateBulder interface {
	Build() *AggreGated
	WithEnvironment(string) AggregateBulder
	WithDescription(string) AggregateBulder
	WithSolution(string) AggregateBulder
	WithFixed(string) AggregateBulder
}

type AggreGated struct {
	Error       string `json:"error"`
	Environment string `json:"environment"`
	Description string `json:"description"`
	Solution    string `json:"solution"`
	CreatedAt   string `json:"created"`
	Fixed       string `json:"fixed"`
}

type aggreGated struct {
	err         string `json:"error"`
	environment string `json:"environment"`
	description string `json:"description"`
	solution    string `json:"solution"`
	createdAt   string `json:"created"`
	fixed       string `json:"fixed"`
}

func (a *aggreGated) WithEnvironment(str string) AggregateBulder {
	a.environment = str
	return a
}

func (a *aggreGated) WithDescription(str string) AggregateBulder {
	a.description = str
	return a
}

func (a *aggreGated) WithSolution(str string) AggregateBulder {
	a.solution = str
	return a
}

func (a *aggreGated) WithFixed(flg string) AggregateBulder {
	a.fixed = flg
	return a
}

func (a *aggreGated) Build() *AggreGated {
	return &AggreGated{
		Error:       a.err,
		Environment: a.environment,
		Description: a.description,
		Solution:    a.solution,
		Fixed:       a.fixed,
		CreatedAt:   time.Now().Format("2006/1/2 15:04:05"),
	}
}

func NewAggregated(err string) AggregateBulder {
	return &aggreGated{
		err: err,
	}
}
