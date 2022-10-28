package parsinglogfiles

import (
	"fmt"
	"regexp"
)


func IsValidLine(text string) bool {
	txt := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	re := regexp.MustCompile(txt)
	return re.MatchString(text)
	panic("Please implement the IsValidLine function")
}

func SplitLogLine(text string) []string {
	txt:=	`\<(~+|\*+|=+|-+)*\>`
	re := regexp.MustCompile(txt)
	value := re.Split(text, -1)
	return value
	panic("Please implement the SplitLogLine function")
}


func CountQuotedPasswords(lines []string) int {
	txt := `".*[p|P][a|A][s|S][s|S][w|W][|oO][r|R][d|D].*"`
	re := regexp.MustCompile(txt)
	val := 0
	for _, v := range lines {
		if re.MatchString(v) {
			val++
		}
	}
	return val
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	txt := `end-of-line[0-9]*`
	re := regexp.MustCompile(txt)
	tmpslice := re.FindStringSubmatch(text)
	if tmpslice == nil {
		return text
	}
		return re.ReplaceAllString(text, "")
/*
	for _, v := range tmpslice {
	}

	*/
//	return text
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
//	txt := `User +.*`
//	txt := `(User +\w*)`
	 txt := `(?P<tug>User +)(?P<username>\w*)(?P<ignore>.*)`
//	txt :=	`(?P<tug>User +)(?P<username>\w*)(?P<ignore>.*)`
	re := regexp.MustCompile(txt)

			fmt.Printf("lines=%v, type of lines %T\n", lines, lines)
	for _, v := range lines {
		if j := re.FindStringSubmatch(v); j != nil {

			fmt.Printf("j[0]=%v, type of j[0] %T\n", j[0], j[0])
			fmt.Printf("v=%v, type of v %T\n", v, v)
			re.MatchString(j[0])
			key := re.SubexpNames()[2]

			fmt.Printf("s1=%s s2=%s value=%s\n", 
			re.SubexpNames()[1], re.SubexpNames()[2], key)
		}
	}

	return lines
	panic("Please implement the TagWithUserName function")
}
