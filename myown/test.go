package main

import (
	"fmt"
	"regexp"
	// "errors"
)

func match(input []string) (result []string)  {

	var output []string

	lookup := make(map[string]string)

	fmt.Println("input", input)
	for _, s := range input {
		// The regular expression means
		// a line delimited by : and ;
		// the first token must be alphabet with hyphen

		cmd := regexp.MustCompile(`^: (?P<marco>[a-zA-Z.\-]+) (?P<details>.*) ;$`)
		matched := cmd.FindString(s)

		if matched == "" {
			output = append(output, s)
			continue
		}

		//removing leading : and trailing ;
		// From ": hello over over ;"
		// To   "hello over over"
		matched = matched[2 : len(matched)-2]

		re := regexp.MustCompile(`(?P<marco>[a-zA-Z.\-]+) (?P<details>.*)`)
		results := re.FindStringSubmatch(matched)
		fmt.Println("results: ", results)
		marco := re.SubexpIndex("marco")
		details := re.SubexpIndex("details")

		fmt.Printf("%s expanded to %s\n", results[marco], results[details])

		lookup[results[marco]] = results[details]
	}

	fmt.Println("maps:", lookup)

	for token, value := range lookup {
		fmt.Println("token", token)
		fmt.Println("value", value)
		temp := regexp.MustCompile(token)
		for i, s := range output {
			fmt.Println("s", s, "i", i)
			output[i] = temp.ReplaceAllString(output[i], value)
		}
	}
	fmt.Println(" after output ", output)

	return output

}

func main() {
	s4 := []string{": foo 5 ;", ": bar foo ;", ": foo 6 ;", "bar foo"}
	s := []string{": + * ;", " 3 4 +" }
	s3 := []string{"1 2 3 4 5"}
	s2 := []string{"dup over ;", ":dup over "}
	t := []string{": dup-twice dup dup ;", "3 dupice"}
	t2 := []string{": dup-and-twice dup dup ;", "5 dup-and-twice"}
	t4 := []string{": dup-and-twice dup dup ;", "5 dup-and-twice",
		": dup-again over swap ;", "2 dup-again 2 3 5 dup-again 4 5",
		"dup-and-twice dup-again"}
	result := match(s)
		fmt.Println("is ok", result)
	match(s2)
	fmt.Println("===")
	match(t)
	fmt.Println("===")
	match(t2)
	fmt.Println("===")
	match(t4)
	match(s3)
	fmt.Println("===")
	match(s4)

}
