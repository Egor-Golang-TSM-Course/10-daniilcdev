package main

import (
	"fmt"
	"lesson10/app"
	"lesson10/counters"
	"lesson10/inputs"
	"lesson10/processing"
	"lesson10/reports"
	"log"
)

func main() {
	// cmd to start from
	// cat sample_logs.txt | go run . -source=./sample_logs.txt -reportOut=true -reportOutPath=./logs/report.txt

	cfg, err := app.LoadConfig()

	if err != nil {
		log.Fatalln("invalid configuration:", err)
	}

	fmt.Printf("config:\n%v\n", cfg)
	fmt.Println()

	processor := processing.NewInputProcessor(
		&[]inputs.DataInput{
			inputs.NewStdinPipe(),
			inputs.NewSourceFile(cfg.SourceFile),
		},
		&[]counters.Counter{
			counters.NewLogLevelCounter(counters.LogLevel(cfg.LogLevel)),
		})

	if cfg.WriteToFile {
		processor = reports.NewReport(cfg, processor)
	}

	processor.Process()
}
