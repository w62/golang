package main
import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("about to panic")
	defer func () {
		if err := recover(); err != nil {
			log.Println("error:", err)
			panic("頂唔住啦⋯⋯")
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
// https://www.youtube.com/watch?v=YS4e4q9oBaU%t=3:57:29
func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}
