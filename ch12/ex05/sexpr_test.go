package sexpr

import (
	"testing"
	"log"
	"fmt"
	"encoding/json"
	"reflect"
)

func TestMarshal(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		IsNew           bool
		IsOld           bool
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
		IsNew: true,
		IsOld: false,
	}

	data, err := Marshal(strangelove)
	if err != nil {
		log.Println(err)
	}
	movie := Movie{}
	json.Unmarshal(data, &movie)
	if !reflect.DeepEqual(movie, strangelove) {
		t.Errorf("deepEqual is not true.")
	}
	fmt.Printf("%v", string(data))
}
