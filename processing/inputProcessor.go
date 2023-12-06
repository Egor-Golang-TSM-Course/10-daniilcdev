package processing

import (
	"fmt"
	"lesson10/counters"
	"lesson10/inputs"
)

type inputProcessor struct {
	sources  *[]inputs.DataInput
	counters *[]counters.Counter
}

func NewInputProcessor(inputSrcs *[]inputs.DataInput, cntrs *[]counters.Counter) Processor {
	return &inputProcessor{
		sources:  inputSrcs,
		counters: cntrs,
	}
}

func (ip *inputProcessor) Process() Summaries {
	println("-> processing inputs:")

	res := make(Summaries, 0, len(*ip.sources))
	for _, src := range *ip.sources {
		fmt.Printf("input - %s\n", src.Name())

		m := make(Metrics)
		nBytes, err := ip.scan(src, m)

		if err != nil {
			fmt.Printf("\t### can't process data, err: %v\n", err)
			continue
		}

		s := &Summary{
			Source: src.Name(),
			Values: m,
		}

		res = append(res, s)

		fmt.Printf("\t### %d bytes processed\n", nBytes)

		if len(m) > 0 {
			fmt.Printf("\t### summary - %v\n", m)
		}
	}

	return res
}

func (ip *inputProcessor) scan(input inputs.DataInput, counts Metrics) (nBytes int, err error) {
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
