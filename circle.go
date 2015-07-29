// circle.go

package circle

import (
	"errors"
	"hash/adler32"

	"github.com/charlesetc/dive"
)

var hash func([]byte) uint32 = adler32.Checksum

type Circle struct {
	address []byte
	hash    uint32
	next    *Circle
}

const (
	ReplicationDepth int = 1
)

func NewCircleHead() *Circle {
	circle := new(Circle)
	circle.hash = 0 // 0 means it's the head
	// circle.address is undefined
	circle.next = circle
	return circle
}

func NewCircle(address []byte) *Circle {
	circle := new(Circle)
	circle.address = address
	circle.hash = hash(address)
	return circle
}

func NewCircleString(address string) *Circle {
	return NewCircle([]byte(address))
}

func (c *Circle) Add(incoming *Circle) *Circle {
	var current *Circle
	for current = c; current.next.hash < incoming.hash; current = current.next {
		if current.next.hash == 0 {
			break
		}
	}
	incoming.next = current.next
	current.next = incoming
	return incoming
}

func (c *Circle) AddString(address string) *Circle {
	return c.Add(NewCircleString(address))
}

func CircleFromList(nodes []string) *Circle {
	circle := NewCircleHead()

	for _, node := range nodes {
		circle.AddString(node)
	}

	return circle
}

func CircleFromNode(node *dive.Node) *Circle {
	circle := NewCircleHead()
	for _, rec := range dive.GetAliveFromMap(node.Members) {
		circle.AddString(rec.Address)
	}
	return circle
}

// Will loop forever with an empty node...
func (c *Circle) KeyAddress(key []byte) func() ([]byte, error) {
	hashed := hash(key)

	var current *Circle
	for current = c; current.next.hash < hashed; current = current.next {
		if current.next.hash == 0 {
			break
		}
	}

	i := 0
	return func() ([]byte, error) {
		output := current.address
		i++

		if i > ReplicationDepth {
			return []byte{}, errors.New("No more replications.")
		}

		current = current.next
		if current.hash == 0 {
			current = current.next
		}

		return output, nil
	}
}
