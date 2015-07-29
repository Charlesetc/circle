// circle_test.go

package circle

import "testing"

func init() {
	Hash = func(bytes []byte) []byte {
		return bytes
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
	c := CircleFromList([]string{"1", "2", "3"})
	val, err := c.KeyAddress([]byte("2"))()
	if err != nil {
		panic(err)
	}
	Equal(t, val[0], "1"[0])
	val, err = c.KeyAddress([]byte("3"))()
	if err != nil {
		panic(err)
	}
	Equal(t, val[0], "2"[0])
}

func TestLargeAddress(t *testing.T) {
	c := CircleFromList([]string{"b", "c", "a", "y"})
	val, err := c.KeyAddress([]byte("z"))()
	if err != nil {
		panic(err)
	}
	Equal(t, val[0], "y"[0])
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
