package processing

import (
	"fmt"
	"lesson10/counters"
	"lesson10/inputs"
)

type Processor interface {
	Process()
}

type inputProcessor struct {
	sources  *[]inputs.DataInput
	counters *[]counters.Counter
}

func NewInputProcessor(cgf *Config, inputSrcs *[]inputs.DataInput, cntrs *[]counters.Counter) Processor {
	return &inputProcessor{
		sources:  inputSrcs,
		counters: cntrs,
	}
}

func (ip *inputProcessor) Process() {
	println("-> processing inputs:")

	counts := map[string]int{}
	for _, src := range *ip.sources {
		fmt.Printf("input - %T\n", src)
		nBytes, err := ip.scan(src, counts)

		if err != nil {
			fmt.Printf("\t### can't process data, err: %v\n", err)
			continue
		}

		fmt.Printf("\t### %d bytes processed\n", nBytes)

		if len(counts) > 0 {
			fmt.Printf("\t### summary - %v\n", counts)
			clear(counts)
		}
	}
}

func (ip *inputProcessor) scan(input inputs.DataInput, counts map[string]int) (nBytes int, err error) {
	scanner, err := input.Open()

	if err != nil {
		return 0, fmt.Errorf("can't open scanner for source: %T, err: %v\n", input, err)
	}

	defer input.Close()

	for scanner.Scan() {
		ln := scanner.Text()
		nBytes += len(ln)
		for _, counter := range *ip.counters {
			counter.Count(ln, counts)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("stdin scanner failed - %v", err)
	}

	return nBytes, nil
}