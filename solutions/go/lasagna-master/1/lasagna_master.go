package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodleGrams int, sauceLiters float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodleGrams += 50
		} else if layer == "sauce" {
			sauceLiters += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	copy(scaledQuantities, quantities)
	for i := 0; i < len(scaledQuantities); i++ {
		scaledQuantities[i] = scaledQuantities[i] * float64(portions) / 2
	}
	return scaledQuantities
}
