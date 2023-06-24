package sharedkernel

import "fmt"

type Status int8

const (
	StatusPlaced Status = iota
	StatusInProcess
	StatusFulfilled
)

func (e Status) String() string {
	return fmt.Sprintf("%d", int(e))
}
