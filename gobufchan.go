package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func myFunc(channel chan int, someval int){
	wg.Done()
	channel <- someval *5
}

func main(){
	var channel = make(chan int,10)
	
	for i:=0;i<10;i++{
		wg.Add(1)
		go myFunc(channel, i)		
	}
	
	wg.Wait()
	close(channel)
	
	for item:= range channel{
		fmt.Println(item)
	}
}