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

	nBytes, err := scanPipe()
	if err != nil {
		fmt.Printf("can't parse piped data: %v\n", err)
	}

	if nBytes > 0 {
		fmt.Printf("N bytes read from stdin: %v\n", nBytes)
	}

	if len(counts) > 0 {
		fmt.Println(counts)
	}
}

func scanPipe() (nBytes int64, err error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return 0, fmt.Errorf("stdin error - %v", err)
	}

	nBytes = stat.Size()
	if nBytes == 0 {
		return 0, nil
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ln := scanner.Text()
		for _, counter := range fileCounters {
			counter.Count(ln, counts)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner failed - %v", err)
	}

	return nBytes, nil
}
