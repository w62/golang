package tournament

import (
	"io"
	"fmt"
	"regexp"
)

func Tally(reader io.Reader, writer io.Writer) error {
	//txt := `^(?P<lteam>(\w|\s)+);(?P<rteam>(\w|\s)+);(?P<result>(\w|\s)+)$` 
	txt := `^(?P<lteam>(\w|\s)+);(?P<rteam>(\w|\s)+);(?P<result>(\w|\s)+)$`
	re := regexp.MustCompile(txt)
	buf := ""

	fmt.Println("reader=",reader)
	if b, err := io.ReadAll(reader); err == nil {
		buf = string(b)
	}
	fmt.Println("buf=",buf)
	buf2 := "Allegoric Alaskians;Blithering Badgers;win
	Devastating Donkeys;Courageous Californians;draw"
/*
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win"
*/
	match := re.FindStringSubmatch(buf2)
	fmt.Println("match=",match)
	  for i, name := range re.SubexpNames() {
        fmt.Printf("'%s'\t %d -> %s\n", name, i, match[i])
    }
//	fmt.Println(re)
	return nil
	panic("Please implement the Tally function")
}
