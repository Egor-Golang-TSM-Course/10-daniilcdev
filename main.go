package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	config, err := LoadConfig()

	if err != nil {
		log.Fatalln("wrong configuration:", err)
	}
	fmt.Println(config)
	err = scanPipe(config)
	if err != nil {
		fmt.Printf("can't parse piped data: %v", err)
	}
}

func scanPipe(cfg *Config) error {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("stdin error - %v", err)
	}
	if stat.Size() == 0 {
		return fmt.Errorf("stdin is empty")
	}

	counts := make(map[string]int, 3)
	match := []string{}

	if cfg.LogLevel == "ERROR" {
		match = append(match, "ERROR")
	} else if cfg.LogLevel == "WARNING" {
		match = append(match, "ERROR", "WARNING")
	} else if cfg.LogLevel == "INFO" {
		match = append(match, "INFO", "WARNING", "ERROR")
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ln := scanner.Text()
		for _, logLevel := range match {
			if strings.Contains(ln, logLevel) {
				counts[logLevel]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner failed - %v", err)
	}

	if len(counts) > 0 {
		fmt.Println(counts)
	}

	return nil
}
