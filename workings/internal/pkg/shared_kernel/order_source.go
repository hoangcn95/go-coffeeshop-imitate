package sharedkernel

import "fmt"

type OrderSource int8

const (
	OrderSourceCounter OrderSource = iota
	OrderSourceWeb
)

func (e OrderSource) String() string {
	return fmt.Sprintf("%d", int(e))
}
