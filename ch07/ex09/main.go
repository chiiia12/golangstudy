package main

import (
	"fmt"
	"os"
	"sort"
	"time"
	"html/template"
)

//webServerにするの忘れてた
//linkでquery parameterを追加していく
func main() {
	title := func(t1 *Track, t2 *Track) bool {
		return t1.Title < t2.Title
	}

	artist := func(t1 *Track, t2 *Track) bool {
		return t1.Artist < t2.Artist
	}

	fmt.Println("=========NOT SORT=============")
	printTracks(tracks)

	fmt.Println("=========ORIGINAL SORT=============")
	var os originalSort
	os.track = tracks
	os.AddLessFunc(title)
	os.AddLessFunc(artist)
	sort.Sort(&os)
	printTracks(os.track)

	fmt.Println("=========STABLE SORT=============")
	var os2 originalSort
	os2.track = tracks
	os2.AddLessFunc(title)
	os2.AddLessFunc(artist)
	for os.HasNext() {
		sort.Stable(&os2)
	}
	printTracks(os2.track)

}

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	var t = template.Must(template.New("track").Parse(`
<table>
	<tr>
		<td>Title</td>
		<td>Artist</td>
		<td>Album</td>
		<td>Year</td>
		<td>Length</td>
	</tr>
	<tr>
	{{range .}}
		<td>{{.Title}}</td>
		<td>{{.Artist}}</td>
		<td>{{.Album}}</td>
		<td>{{.Year}}</td>
		<td>{{.Length}}</td>
	{{end}}
	</tr>
</table>`))
	if err := t.Execute(os.Stdout, tracks); err != nil {
		fmt.Println(err)
	}
}

type lessFunc func(p1, p2 *Track) bool

type originalSort struct {
	track []*Track
	less  []lessFunc
	index int
}

func (x *originalSort) Len() int { return len(x.track) }

func (x *originalSort) Less(i, j int) bool {
	p, q := x.track[i], x.track[j]
	var k int
	for k = 0; k < len(x.less)-1; k++ {
		less := x.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return x.less[k](p, q)
}
func (x *originalSort) Swap(i, j int) { x.track[i], x.track[j] = x.track[j], x.track[i] }
func (x *originalSort) AddLessFunc(f func(p1, p2 *Track) bool) {
	x.index++
	x.less = append(x.less, f)
}
func (x *originalSort) HasNext() bool {
	x.index--
	return x.index >= 0
}
