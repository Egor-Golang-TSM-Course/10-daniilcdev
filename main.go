package main

import (
	"bufio"
	"fmt"
	"lesson10/counters"
	"log"
	"os"
)

var counts map[string]int
var fileCounters []counters.Counter

func main() {
	cfg, err := LoadConfig()

	if err != nil {
		log.Fatalln("wrong configuration:", err)
	}

	counts = make(map[string]int, 3)
	fileCounters = []counters.Counter{
		counters.NewLogLevelCounter(counters.LogLevel(cfg.LogLevel)),
	}

	_, err = scanPipe()
	if err != nil {
		fmt.Printf("can't process piped data: %v\n", err)
	}

	if len(counts) > 0 {
		fmt.Println(counts)
	}
}

func scanPipe() (nBytes int64, err error) {
	// handle 'cat sample_logs.txt | go run .' case...

	stat, err := os.Stdin.Stat()
	if err != nil {
		return 0, fmt.Errorf("stdin error - %v", err)
	}

	nBytes = stat.Size()
	if nBytes == 0 {
		return 0, nil
	}

	fmt.Println("### got data via pipe, processing...")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ln := scanner.Text()
		for _, counter := range fileCounters {
			counter.Count(ln, counts)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("### stdin scanner failed - %v", err)
	}

	fmt.Println("### piped data processed!")

	return nBytes, nil
}
