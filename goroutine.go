package main

import (
"fmt"
"time"
)

func myfun(from string) {
  for i:=0;i<5;i++{
	fmt.Println(from," :: ",i)
  }
  
}

func main(){
	go myfun("goroutine")

	//myfun("function")
	fmt.Println("done")
	
	time.Sleep(time.Second)
	
}