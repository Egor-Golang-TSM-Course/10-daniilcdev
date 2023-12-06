package processing

type Metrics map[string]int

type Summaries []*Summary

type Summary struct {
	Source string
	Values Metrics
}
type Processor interface {
	Process() Summaries
}
