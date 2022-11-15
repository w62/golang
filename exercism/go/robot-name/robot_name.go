// Failed exercise
// Really don't exactly what the expected problem is
// It is a bit more on understand the problem rather than the Go language
package robotname

import (
	"fmt"
	
	"math/rand"
//	"bytes"
	"time"
	"sync"
)

var namelist = map[string]int{}
// Define the Robot type here.
type Robot struct {
	name string
}


func (r *Robot) Name() (string, error) {

//	fmt.Println("len(namelist)=",len(namelist))

	name := randomString(2)+ fmt.Sprintf("%03d",randomInt(0, 999))

	_, chk := namelist[name]
	if  chk {
		return "", fmt.Errorf("Name %s reissued after %d robots.", name, len(namelist))
	} else {
		namelist[name] = 0
	}

	return name, nil
	panic("Please implement the Name function")
}

func (r *Robot) Reset() {
	r.name = ""
	return
	panic("Please implement the Reset function")
}


func randomString(l int) string {
    bytes := make([]byte, l)
    for i := 0; i < l; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}

func randomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())  
    var mu sync.Mutex  
    mu.Lock()
	result := rand.Intn(max-min)
	mu.Unlock()
	return min + result
}



