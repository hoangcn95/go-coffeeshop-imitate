package sharedkernel

import "fmt"

type Location int8

const (
	LocationAtlanta Location = iota
	LocationCharlotte
	LocationRaleigh
)

func (e Location) String() string {
	return fmt.Sprintf("%d", int(e))
}
