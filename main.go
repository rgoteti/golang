package main

import "fmt"

func main(){
	fmt.Println("H W!!")
	
	var a [5]int
	fmt.Println("a is ",a)
	
	a[4]=100
	fmt.Println("emp is ",a)
	
	e := [5]int{10,20,30,40,50}
	
	fmt.Println("new e is ", e)
	fmt.Println("e length", len(e))
	
	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println(" twoD :: ",twoD)
}