package main

import (
	"fmt"
	"example.com/test/sub"
	"whaterever.shit/willdo"

)

type helper func (string) string

func main() {

	var oneArm helper = support.Alice 
	twoArm := support.Bob

	fmt.Println(support.Greeting)
	fmt.Println(support.Message)
	fmt.Println(support.Alice("charlie"))
	fmt.Printf("Alice is %T\n", support.Alice)
	fmt.Println(support.Bob("delta"))
	fmt.Println(Chuck)
	fmt.Println("======")
	fmt.Println(oneArm("echo"))
	fmt.Println(twoArm("foxtrot"))
	fmt.Println("======")

	fmt.Println(support.ReturnMessage())
	support.ChangeGreeting("A new greeting message.")
	fmt.Println(support.Greeting)

	unknown.Um()

	fmt.Println(unknown.Update1("中文啊"))
	unknown.Um()

}

func Chuck (name string) string {
	return "Chuck helps too, " + name 
}



