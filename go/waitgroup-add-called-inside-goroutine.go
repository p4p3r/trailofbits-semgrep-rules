// Code from https://go101.org/article/concurrent-common-mistakes.html

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func main() {
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
	        // ruleid: waitgroup-add-called-inside-goroutine
			wg.Add(1)
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}

	fmt.Println("Wait ...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}


func main() {
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
	    // ok: waitgroup-add-called-inside-goroutine
		wg.Add(1)
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}

	fmt.Println("Wait ...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}