package main

import (
	lasagnamaster "exercism/lasagna_master"
	"fmt"
)

func main() {
	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := lasagnamaster.ScaleRecipe(quantities, 4)
	fmt.Println(quantities, scaledQuantities)
}
