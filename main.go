package main

import (
	"fmt"
	"lesson10/counters"
	"lesson10/inputs"
	"log"
)

var counts map[string]int
var fileCounters []counters.Counter
var sources []inputs.DataInput

func main() {
	cfg, err := LoadConfig()

	if err != nil {
		log.Fatalln("wrong configuration:", err)
	}

	counts = make(map[string]int, 3)
	fileCounters = []counters.Counter{
		counters.NewLogLevelCounter(counters.LogLevel(cfg.LogLevel)),
	}

	sources = []inputs.DataInput{
		inputs.NewStdinPipe(),
		inputs.NewSourceFile(cfg.SourceFile),
	}

	println("-> processing inputs:")

	for i := range sources {
		src := sources[i]
		fmt.Printf("input - %T\n", src)
		nBytes, err := scan(src)

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

func scan(input inputs.DataInput) (nBytes int, err error) {
	scanner, err := input.Open()

	if err != nil {
		return 0, fmt.Errorf("can't open scanner for source: %T, err: %v\n", input, err)
	}

	defer input.Close()

	for scanner.Scan() {
		ln := scanner.Text()
		nBytes += len(ln)
		for _, counter := range fileCounters {
			counter.Count(ln, counts)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("stdin scanner failed - %v", err)
	}

	return nBytes, nil
}
