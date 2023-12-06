package reports

import (
	"errors"
	"fmt"
	"lesson10/processing"
	"log"
	"os"
	"path"
)

func NewReport(cfg ReportConfig, base processing.Processor) processing.Processor {
	return &reportProcessor{base: base, cfg: cfg}
}

type reportProcessor struct {
	base processing.Processor
	cfg  ReportConfig
}

func (rp *reportProcessor) Process() processing.Metrics {
	m := rp.base.Process()
	writeMetrics(&m, rp)
	return m
}

func writeMetrics(m *processing.Metrics, rp *reportProcessor) {
	dirPath := path.Dir(rp.cfg.ReportFilePath())
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	file, err := os.OpenFile(rp.cfg.ReportFilePath(), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("can't write report, file can't be opened", err)
		return
	}

	defer file.Close()

	for k, v := range *m {
		file.WriteString(fmt.Sprintf("%s - %d\n", k, v))
	}

	fmt.Println("report written at path", rp.cfg.ReportFilePath())
}
