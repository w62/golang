package luhn

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)


func Valid(id string) bool {

	val := 0
	id, err := strip(id)
	if err != nil {

		return false
	}

	for i := 0; i < len(id); i++ {
		parity := len(id) % 2
		j, _ := strconv.Atoi(id[i : i+1])

		if i%2 == parity {

			j *= 2
			if j > 9 {
				j -= 9
			}
		}

		val += j

	}

	return val%10 == 0
	panic("Please implement the Valid function")
}

func strip(id string) (string, error) {
	txt := `^(\d|\s)+$`
	re := regexp.MustCompile(txt)
	val := re.MatchString(id)
	if val {
		id = strings.ReplaceAll(id, " ", "")
	} else {
		return "", errors.New("Invalid code")
	}

	if len(id) <= 1 {
		return "", errors.New("Invalid length")
	}

	return id, nil
}
