package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mx sync.Mutex
var x int

func main() {
	wg.Add(2)
	go increment("Foo: ")
	go increment("Bar: ")
	wg.Wait()
	fmt.Println("Final counter:", x)
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mx.Lock()
		x++
		fmt.Println(s, i, "Counter", x)
		mx.Unlock()
	}
	wg.Done()
}
