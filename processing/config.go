package processing

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	SourceFile  string
	OutputFile  string
	LogLevel    string
	WriteToFile bool
}

func LoadConfig() (*Config, error) {
	sourceFilePath := flag.String("source", "", "source file with logs to analyze")

	val, _ := strconv.ParseBool(os.Getenv("REPORT_OUT"))
	writeReportOut := flag.Bool("reportOut", val, "should write report to file")

	outFilePath := flag.String("reportOutPath",
		os.Getenv("REPORT_OUT_PATH"),
		"output report file path")

	logLevel := flag.String("minLogLevel", "INFO", "minimum log level to count")

	flag.Parse()

	if *writeReportOut && *outFilePath == "" {
		*writeReportOut = false
		*outFilePath = ""
		fmt.Println("warn: 'write to file' option not configured, report file will not be written")
	}

	return &Config{
		SourceFile:  *sourceFilePath,
		OutputFile:  *outFilePath,
		LogLevel:    strings.ToUpper(*logLevel),
		WriteToFile: *writeReportOut,
	}, nil
}

func (cfg *Config) String() string {
	return fmt.Sprintf("-SourceFile=%s\n-OutputFle=%s\n-LogLevel=%s\n-ShouldWriteReport=%v",
		cfg.SourceFile, cfg.OutputFile, cfg.LogLevel, cfg.WriteToFile)
}
