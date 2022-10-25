// Exercise 1.12: Modify the Lissajous server to read paramter values from the 
// URL. For example, you might arrange it so that a URL like 
// http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead of
// the default 5. Use the strconv.Atoi function to convert the string parameter
// into an integer. You can see its documentation with go doc strconv.Atoi.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
		"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
)
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)


var (
	mu sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler (w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count ++
	mu.Unlock()
	/*
	fmt.Fprintf(w, "Hello web browser. URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
*/

// A small surprise:
// If anything is written to the browser, the browser will prompt for the 
// download instead of displaying the gif.
// Probably due to the context type
	param1 := r.URL.Query().Get("cycles")
	if param1 != "" {
		i, err := strconv.Atoi(param1)
		if err != nil {
			i = 5
		}
//		fmt.Fprintf(w, "cycles=%d\n", i)
		lissajous(w, i)
	} else {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}


func lissajous(out io.Writer, t int) {
	const (
//		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)

	cycles := float64(t)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator

	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t+= res {
			x:= math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
			blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

