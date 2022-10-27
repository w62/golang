package support

var Greeting string

func init () {
	Greeting = "Greetings from support."
	Message = "A message from support."
}

func Alice (name string) string {
	return "Alice supports " + name
}

func Bob (name string) string {
	return "Bob is at your service, " + name
}
