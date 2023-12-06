package inputs

import "bufio"

type DataInput interface {
	Open() (*bufio.Scanner, error)
	Close()
}
