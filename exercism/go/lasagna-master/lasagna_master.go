package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime (layers []string, time int) int {
	localtime := 0
	if time <= 0 {
		localtime = 2
	} else {
		localtime = time 
	}

	return len(layers) * localtime

}

// TODO: define the 'Quantities()' function

func Quantities (layers []string)(int, float64) {
	var noodles int
	var sauce float64


	for i := 0; i < len(layers); i ++ {
		switch layers [i] {
	case "sauce":
		sauce += 0.2
	case "noodles":
		noodles += 50
	}
}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient( friendList, myList []string) {

	myList[len(myList)-1] = friendList[len(friendList)-1]

}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe (portion1 []float64, qty int) []float64 {
	var result []float64
	for i := 0; i< len(portion1); i++ {
		result = append(result, portion1[i] *   float64(qty) / 2 )
	}
	return result
}
