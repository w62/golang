package support

var Message string

func ReturnMessage() string {
	return Message
}

func ChangeGreeting (newGreeting string) {
	Greeting = newGreeting
}
