package main

import (
	"bufio"
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
	}

	println("-> processing inputs")
	for i := range sources {
		src := sources[i]
		scanner, err := src.Open()
		fmt.Println()
		if err != nil {
			fmt.Printf("### can't open scanner for source: %T, err: %v\n", src, err)
			continue
		}

		nBytes, err := scan(scanner)
		if err != nil {
			fmt.Printf("### can't process data source: %T, err: %v\n", src, err)
			continue
		}

		fmt.Printf("### %d bytes processed, source: %T\n", nBytes, src)
	}

	if len(counts) > 0 {
		fmt.Println(counts)
	}
}

func scan(scanner *bufio.Scanner) (nBytes int, err error) {
	// handle 'cat sample_logs.txt | go run .' case...

	fmt.Println("### got data via pipe, processing...")

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

	fmt.Println("### piped data processed!")

	return nBytes, nil
}
