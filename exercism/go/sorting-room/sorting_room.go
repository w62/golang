package sorting
import (
	"strconv"
	"fmt"
	"reflect"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %2.1f", f)
	panic("Please implement DescribeNumber")
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
	panic("Please implement DescribeNumberBox")
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	var dummpy FancyNumber 

	if reflect.TypeOf(dummpy) != reflect.TypeOf(fnb) {
		return 0
	}
	i, err := strconv.Atoi(fnb.Value())
//	yt := reflect.TypeOf(fnb)

	if err != nil {
		return 0
	} else {
		return i
	}
	panic("Please implement ExtractFancyNumber")
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	var dummy FancyNumber

	if reflect.TypeOf(dummy) == reflect.TypeOf(fnb) {
		 val, err := strconv.ParseFloat(fnb.Value(),32)
	if	 err != nil {
			return ""
		}
	return fmt.Sprintf("This is a fancy box containing the number %.1f", val)
} else {
	return fmt.Sprintf("This is a fancy box containing the number 0.0")
}
	panic("Please implement DescribeFancyNumberBox")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v:=i.(type){
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
	panic("Please implement DescribeAnything")
}
