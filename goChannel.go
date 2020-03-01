package main

import "fmt"

func myfunc(channel chan int, someval int){
	channel <- someval*5
}

func main(){
	channel := make(chan int)
	go myfunc(channel,5)
	go myfunc(channel,3)
	
	
	//val1 := <- channel
	//val2 := <- channel
	
	val1,val2 := <- channel, <- channel
	fmt.Println(val1,val2)
}