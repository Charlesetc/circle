// circle.go

package circle

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
)

var Zero []byte = make([]byte, 256)

var Hash func([]byte) []byte = func(bytes []byte) []byte {
	hash := sha256.New()
	hash.Write(bytes)
	fmt.Println(hash.Sum(nil))
	return hash.Sum(nil)
}

type Circle struct {
	address []byte
	hash    []byte
	next    *Circle
}

func (c *Circle) String() string {
	var buffer bytes.Buffer
	for current, first := c, true; len(current.hash) != 0 ||
		first; current, first = current.next, false {
		if !first {
			buffer.WriteString(" -> ")
		}
		buffer.Write(current.address)
		buffer.WriteString("/")
		buffer.Write(current.hash)
	}
	return buffer.String()
}

var (
	ReplicationDepth int = 1
)

func NewCircleHead() *Circle {
	circle := new(Circle)
	circle.hash = []byte{} // empty is head.
	// circle.address is undefined
	circle.next = circle
	return circle
}

func NewCircle(address []byte) *Circle {
	circle := new(Circle)
	circle.address = address
	circle.hash = Hash(address)
	return circle
}

func NewCircleString(address string) *Circle {
	return NewCircle([]byte(address))
}

func (c *Circle) Add(incoming *Circle) *Circle {
	var current *Circle
	for current = c; bytes.Compare(current.next.hash, incoming.hash) == -1; current = current.next {
		if bytes.Compare(current.next.hash, nil) == 0 {
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

func CircleFromList(strs []string) *Circle {
	circle := NewCircleHead()
	for _, str := range strs {
		circle.AddString(str)
	}
	return circle
}

// Will loop forever with an empty node...
func (c *Circle) KeyAddress(key []byte) func() ([]byte, error) {
	hashed := Hash(key)

	var current *Circle
	for current = c.next; bytes.Compare(current.next.hash, nil) != 0 &&
		bytes.Compare(current.next.hash, hashed) == -1; current = current.next {
	}

	i := 0
	return func() ([]byte, error) {
		output := current.address
		i++

		if i > ReplicationDepth {
			return []byte{}, errors.New("No more replications.")
		}

		current = current.next
		if bytes.Compare(current.hash, nil) == 0 {
			current = current.next
		}

		return output, nil
	}
}
