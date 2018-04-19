package decode

import (
	"testing"
	"fmt"
	"github.com/chiiia12/golangstudy/ch12/ex03"
	"reflect"
)

func TestDecode(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		IsFamous        bool
		Rating          float64
		Remark          interface{}
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		IsFamous: true,
		Rating:   3.23,
		Remark:   string("aaa"),
	}
	sStatement, _ := sexpr.Marshal(strangelove)
	fmt.Println(string(sStatement))
	answer := &Movie{}
	Unmarshal(sStatement, answer)
	if strangelove.Title != answer.Title {
		t.Errorf("not Equal title")
	}
	if strangelove.Subtitle != answer.Subtitle {
		t.Errorf("not Equal subtitle")
	}
	if strangelove.Year != answer.Year {
		t.Errorf("not Equal Year")
	}
	if !reflect.DeepEqual(strangelove.Actor, answer.Actor) {
		t.Errorf("not Equal Actor")
	}
	if !reflect.DeepEqual(strangelove.Oscars, answer.Oscars) {
		t.Errorf("not Equal Oscars")
	}
	if !reflect.DeepEqual(strangelove.Sequel, answer.Sequel) {
		t.Errorf("not Equal Sequel")
	}
	if !reflect.DeepEqual(strangelove.IsFamous, answer.IsFamous) {
		t.Errorf("not Equal IsFamous")
	}
	if !reflect.DeepEqual(strangelove.Rating, answer.Rating) {
		t.Errorf("not Equal Rating")
	}
	if !reflect.DeepEqual(strangelove.Remark, answer.Remark) {
		t.Errorf("not Equal Remark %v", answer.Remark)
	}
}
