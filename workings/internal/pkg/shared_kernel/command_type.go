package sharedkernel

import "fmt"

type CommandType int8

const (
	CommandTypePlaceOrder CommandType = iota
)

func (e CommandType) String() string {
	return fmt.Sprintf("%d", int(e))
}
