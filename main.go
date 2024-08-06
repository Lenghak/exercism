package main

import (
	elon "exercism/elons_toys"
	"fmt"
)

func main() {
	car := elon.NewCar(5, 2)
	fmt.Print(car.CanFinish(100))
}
