package main

import (
	"fmt"
	"lesson10/counters"
	"lesson10/inputs"
	"lesson10/processing"
	"log"
)

func main() {
	cfg, err := processing.LoadConfig()

	if err != nil {
		log.Fatalln("invalid configuration:", err)
	}

	fmt.Printf("config:\n%v\n", cfg)
	fmt.Println()

	processor := processing.NewInputProcessor(cfg,
		&[]inputs.DataInput{
			inputs.NewStdinPipe(),
			inputs.NewSourceFile(cfg.SourceFile),
		},
		&[]counters.Counter{
			counters.NewLogLevelCounter(counters.LogLevel(cfg.LogLevel)),
		})
	processor.Process()
}
