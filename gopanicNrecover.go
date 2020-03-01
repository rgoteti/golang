package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cleanup(){
	if r:=recover();r!=nil{
		fmt.Println("Cleanup from panic :: ",r)
	}
}
func print(from string){
	defer wg.Done()
	defer cleanup()
	for i:=0;i<3;i++ {
		fmt.Println(from, "  :: ", i)
		if i==2 {
			panic("oh its a 2!!")
		}
	}
}

func main(){
	wg.Add(1)
	go print("raj")
	wg.Add(1)
	go print("rakku")
	wg.Wait()
}