package unknown

var Umessage1 string

func init () {
	Umessage0 = "Umessage 0"
	Umessage1 = "Umessage 1"
}

func Update0 (newMessage string) string{
	Umessage0 = newMessage
	return "Umessage0 updated to "+Umessage0

}
