package hamming
import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len (b) {
		return 0, errors.New("DNA lengths are different.")
	}

	hamming := 0

	for i, _ := range a {
		if a[i] != b[i] {
			hamming ++
		}
	}

	return hamming, nil
	panic("Implement the Distance function")
}
