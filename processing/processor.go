package processing

type Metrics map[string]int

type Processor interface {
	Process() Metrics
}
