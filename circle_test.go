// circle_test.go

package circle

import (
	"fmt"
	"testing"

	"github.com/charlesetc/dive"
)

func init() {
	hash = func(bytes []byte) uint32 {
		return uint32(bytes[0])
	}
}

func Equal(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Error("Not equal:", a, b)
	}
}

func NotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Error("Not equal:", a, b)
	}
}

func NewTestNode() *dive.Node {
	node := &dive.Node{
		Members: make(map[string]*dive.LocalRecord),
	}
	return node
}

func TestAdd(t *testing.T) {
	a := NewCircleHead()
	b := a.Add(NewCircleString("b"))
	c := a.AddString("c")
	Equal(t, a, a)
	NotEqual(t, a, b)
	NotEqual(t, c, b)
	Equal(t, a.next, b)
	Equal(t, b.next, c)
	Equal(t, c.next, a)
}

func TestNode(t *testing.T) {
	n := NewTestNode()
	n.Members["2"] = &dive.LocalRecord{BasicRecord: dive.BasicRecord{Address: "2"}}
	n.Members["1"] = &dive.LocalRecord{BasicRecord: dive.BasicRecord{Address: "1"}}
	n.Members["3"] = &dive.LocalRecord{BasicRecord: dive.BasicRecord{Address: "3"}}
	n.Members["4"] = &dive.LocalRecord{BasicRecord: dive.BasicRecord{Address: "4"}}
	c := CircleFromNode(n)
	val, err := c.keyAddress([]byte("2"))()
	fmt.Println(string(val), err)
}

// func TestAddress(t *testing.T) {
// 	ClusterSize := 5
// 	port := 0
// 	nodes := make([]*dive.Node, ClusterSize)
//
// 	first := dive.NewNode(port, "")
// 	port++
// 	nodes[0] = first
// 	seed := first.Address()
//
// 	time.Sleep(dive.PingInterval)
//
// 	for i := 1; i < ClusterSize; i++ {
// 		nodes[i] = dive.NewNode(port, seed)
// 		port++
// 	}
// 	time.Sleep(dive.PingInterval * time.Duration(ClusterSize*2))
//
// }
