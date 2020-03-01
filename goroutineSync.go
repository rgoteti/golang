package main

import (
    "fmt"    
	"sync"
)

var wg sync.WaitGroup

func f(from string) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }	
}

func main() {

    wg.Add(1)
    go f("direct")
    wg.Add(1)
    go f("goroutine")

    wg.Wait()
    fmt.Println("done")
}