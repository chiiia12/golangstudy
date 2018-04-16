package sexpr

import (
	"testing"
	"log"
	"fmt"
	"strings"
)

func TestMarshal(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "",
		Subtitle: "",
		Year:     0,
		Actor:    nil,
		Oscars:   nil,
	}

	data, err := Marshal(strangelove)
	if err != nil {
		log.Println(err)
	}

	if strings.Contains(string(data), "Title") {
		t.Errorf("data has 'Title'.")
	}
	if strings.Contains(string(data), "Subtitle") {
		t.Errorf("data has 'Subtitle'.")
	}
	if strings.Contains(string(data), "Year") {
		t.Errorf("data has 'Year'.")
	}
	if strings.Contains(string(data), "Actor") {
		t.Errorf("data has 'Actor'.")
	}
	if strings.Contains(string(data), "Oscars") {
		t.Errorf("data has 'Oscars'.")
	}
	fmt.Printf("%v", string(data))
}
