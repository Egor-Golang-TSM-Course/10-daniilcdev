package counters

import "strings"

type LogLevel string

type LogLevelCounter struct {
	logLevels *[]LogLevel
}

func (llc *LogLevelCounter) Count(ln string, out map[string]int) {
	for _, logLevel := range *llc.logLevels {
		ll := string(logLevel)
		if strings.Contains(ln, ll) {
			out[ll]++
		}
	}
}

func NewLogLevelCounter(minLogLevel LogLevel) Counter {
	var match []LogLevel = []LogLevel{}

	const err LogLevel = "ERROR"
	const warn LogLevel = "WARNING"
	const info LogLevel = "INFO"

	if minLogLevel == err {
		match = append(match, err)
	} else if minLogLevel == warn {
		match = append(match, err, warn)
	} else if minLogLevel == info {
		match = append(match, info, warn, err)
	}
	return &LogLevelCounter{logLevels: &match}
}
