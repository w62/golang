package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string `required max: "100"`
	// Tags https://www.youtube.com/watch?v=YS4e4q9oBaU&t=2:42:27
	Origin string 
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

// Embedded https://www.youtube.com/watch?v=YS4e4q9oBaU&t=2:39:11

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	field2, _ := t.FieldByName("Origin")
	fmt.Println(field2.Tag)

	a := []int{}
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 1)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, []int{6, 7, 8, 9, 10, 11, 12, 13}...)

	// spread? 
	// https://www.youtube.com/watch?v=YS4e4q9oBaU%t=2:09:01
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"

	team := map[string]int{
		"alice": 0,
		"bob": 1,
		"carol": 2, 
		"dave": 3,
	}

	if found, ok := team["flora"]; ok {
		// https://www.youtube.com/watch?v=YS4e4q9oBaU%t=2:51:28
		fmt.Println(found)
	}

	fmt.Println(!!!!!!!!!!!!!
	!!!!!!!!!!!!!!!!!!!
	!!!!true)
	// https://www.youtube.com/watch?v=YS4e4q9oBaU%t=3:14:08
	var z interface{} =  20
	switch z.(type) {
	case int:
		fmt.Println("z is int type")
		if z.(int) < 10 {
			fmt.Println("too small")
			break
		}
		fmt.Println(z.(int), "ok go on")
	case float64:
		fmt.Println("z is string type")
	case string:
		if len(z.(string)) < 10 {
			fmt.Println("too short")
			break
		}
		fmt.Println("z is string type")
	case [3]int:
		fmt.Println("z is [3]int{} type")
	case rune:
		fmt.Println("z is rune type")
	default:
		fmt.Println("z is another type")
	}
	//https://www.youtube.com/watch?v=YS4e4q9oBaU%t=3:24:15
	// 如何一次郁兩個變數而不失霸氣
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	
}
