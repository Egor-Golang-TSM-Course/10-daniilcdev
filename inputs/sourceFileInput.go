package inputs

import (
	"bufio"
	"os"
)

type sourceFileInput struct {
	fPath string
	f     *os.File
}

func NewSourceFile(srcPath string) DataInput {
	return &sourceFileInput{fPath: srcPath}
}

func (sfi *sourceFileInput) Name() string {
	return sfi.fPath
}

func (sfi *sourceFileInput) Open() (*bufio.Scanner, error) {
	f, err := os.Open(sfi.fPath)
	if err != nil {
		return bufio.NewScanner(&emptyReader{}), nil
	}

	sfi.f = f
	scanner := bufio.NewScanner(f)
	return scanner, nil
}

func (sfi *sourceFileInput) Close() {
	if sfi.f == nil {
		return
	}

	sfi.f.Close()
	sfi.f = nil
}
