package main

import "fmt"

type car struct{
	gas_pedel uint16
	brake_pedel uint16
	steering int16
	speed    uint16
}

func (c car) kmph() uint16{
	return c.gas_pedel*c.speed
}

func main(){
	b_car := car{100,0, 90, 121}
	
	fmt.Println(b_car.speed)
	fmt.Println(b_car.kmph())
}