package lasagnamaster

const DEFAULT_TIME = 2

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {

	averageTime := DEFAULT_TIME

	if time != 0 {
		averageTime = time
	}

	return len(layers) * averageTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauces float64) {
	noodles = 0
	sauces = 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50

		} else if layers[i] == "sauce" {
			sauces += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {

	scaledRecipe := make([]float64, len(quantities))

	for i := 0; i < len(scaledRecipe); i++ {
		scaledRecipe[i] = (quantities[i] / 2.0) * float64(portion)
	}
	return scaledRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
