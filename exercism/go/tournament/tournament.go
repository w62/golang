package tournament

import (
	"fmt"
	"io"
		"regexp"
		"sort"
	"strings"
)

	type score struct {
		team string
		score int
	}

func Tally(reader io.Reader, writer io.Writer) error {
	//txt := `^(?P<lteam>(\w|\s)+);(?P<rteam>(\w|\s)+);(?P<result>(\w|\s)+)$`
		txt := `^(?P<lteam>(\w|\s)+);(?P<rteam>(\w|\s)+);(?P<result>(\w|\s)+)$`
	re := regexp.MustCompile(txt)
	buf := ""
	lteam := ""
	rteam := ""
	state := ""
	var MP map[string]int
	var W map[string]int
	var D map[string]int
	var L map[string]int
	var P map[string]int
	MP = make(map[string]int)
	W = make(map[string]int)
	D = make(map[string]int)
	L = make(map[string]int)
	P = make(map[string]int)
//	fmt.Println("reader=", reader)
	if b, err := io.ReadAll(reader); err == nil {
		buf = string(b)
	}

	s := strings.Split(buf, "\n")

	for i := 0; i < len(s); i++ {
		if s[i] == "" {
//			fmt.Printf("s[%d]=is empty\n", i)
		} else {
//			fmt.Printf("s[%d]=%s\n", i, s[i])
			r:= re.FindStringSubmatch(s[i])
			lteam = r[1]
			rteam = r[3]
			state = r[5]

			MP[lteam] ++
			MP[rteam] ++
			switch state {
			case "win":
				W[lteam] += 1
				L[rteam] += 1
				P[lteam] += 3
			case "loss":
				W[rteam] += 1
				L[lteam] += 1
				P[rteam] += 3
			case "draw":
				D[rteam] += 1
				D[lteam] += 1
				P[lteam] += 1
				P[rteam] += 1
			}
//			fmt.Printf("lteam=%s rteam=%s status=%s\n\n", lteam, rteam, state)
		}
	}

	_,_ = writer.Write( []byte(" Team                           | MP |  W |  D |  L |  P\n"))


	var list []score 
	
	for val,_ := range P {
		var note score
		note.team = val
		note.score = P[val]
		list = append(list, note)
	}

	fmt.Println("list=",list)

	list = sortteam(list)

	fmt.Println("list=",list)
/*
	for val,_ := range MP {
		s := fmt.Sprintf("%-*s |  %d |  %d |  %d |  %d |  %d\n", 30, val, MP[val], W[val], D[val], L[val], P[val])
		_, _ = writer.Write([]byte(s))
	}
*/
	for i, _ := range list {
		s := fmt.Sprintf("%-*s |  %d |  %d |  %d |  %d |  %d\n", 30, list[i].team, MP[list[i].team], 
		W[list[i].team], D[list[i].team], L[list[i].team], list[i].score)
		_, _ = writer.Write([]byte(s))
	}

	return nil

	panic("Please implement the Tally function")
}

func sortteam( list []score) []score {

	sort.Slice(list, func(i,j int) bool {
		if list[i].score == list[j].score {
			return list[i].team > list[j].team
		}
		return list[i].score > list[j].score
	})

	/*

		sort.Slice(list, func (i, j int) bool {
			return list[i].team > list[j].team
		})

		*/

		fmt.Println("sorted list=", list)

	return list
}
