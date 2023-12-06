package inputs

import "bufio"

type DataInput interface {
	Name() string
	Open() (*bufio.Scanner, error)
	Close()
}
