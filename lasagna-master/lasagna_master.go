package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTime int) int {
	if averageTime > 0 {
		return len(layers)*averageTime
	} else {
		return len(layers)*2
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int,sauce float64){
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "sauce":
			sauce+=0.2
		case "noodles":
			noodles+=50
		default:
			sauce+=0
			noodles+=0
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(originalQuantities []float64, portions int) []float64 {
	neededQuantities := []float64{}
	for i := 0; i < len(originalQuantities); i++ {
		neededQuantities = append(neededQuantities, originalQuantities[i]*(float64(portions)/2))
	}
	return neededQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
