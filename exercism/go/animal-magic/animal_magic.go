package chance
import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
	return
	panic("Please implement the SeedWithTime function")
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
	panic("Please implement the RollADie function")
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return 12.0 * rand.Float64() 
	panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	var animalSlice [] string

	//var list map[value]string

	zoo := map[int]string {
		0: "ant",
		1: "beaver",
		2: "cat",
		3: "dog",
		4: "elephant",
		5: "fox",
		6: "giraffe",
		7: "hedgehog",
	}

	for len(zoo) > 0 {
		

		SeedWithTime()
		index := rand.Intn(8) 

		if _, ok := zoo[index]; ok {

		animalSlice = append(animalSlice, zoo[index])
		delete(zoo, index)
		}
	}

	return animalSlice
	panic("Please implement the ShuffleAnimals function")
}
