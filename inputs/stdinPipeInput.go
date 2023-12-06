package inputs

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type stdinPipeInput struct {
}

func NewStdinPipe() DataInput {
	return &stdinPipeInput{}
}

func (spi *stdinPipeInput) Open() (*bufio.Scanner, error) {

	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, fmt.Errorf("stdin error - %v", err)
	}

	nBytes := stat.Size()
	if nBytes == 0 {
		return bufio.NewScanner(&emptyReader{}), nil
	}

	return bufio.NewScanner(os.Stdin), nil
}

type emptyReader struct {
}

func (er *emptyReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}
