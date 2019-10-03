package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	sum int
	v   []int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(a int) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.sum = c.sum + a
	//c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
/*func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}*/

func main() {

	c := SafeCounter{v: make([]int, 5, 6)}
	for i := 0; i < 5; i++ {
		go c.Inc(c.v[i])
	}

	time.Sleep(time.Second)
	fmt.Println(c.sum)
	//fmt.Println(c.Value("somekey"))
}
