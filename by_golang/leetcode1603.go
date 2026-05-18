package main

import (
	"fmt"
)

type ParkingSystem struct {
	slots [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{slots: [4]int{0, big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slots[carType] <= 0 {
		return false
	}
	this.slots[carType]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func Leetcode() {
	obj := Constructor(1, 1, 1)
	param1 := obj.AddCar(1)
	fmt.Printf("param_1: %d\n", param1)

}
