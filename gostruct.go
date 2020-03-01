package main

import "fmt"

type car struct{
	gas_pedel uint16
	brake_pedel uint16
	steering int16
	speed    uint16
}

func main(){
	a_car := car{gas_pedel:100, brake_pedel:0, steering:90, speed:120}
	
	b_car := car{100,0, 90, 121}
	
	fmt.Println(a_car.speed)
	fmt.Println(b_car.speed)
}