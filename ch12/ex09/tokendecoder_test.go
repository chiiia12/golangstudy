package ex09

import (
	"testing"
	"strings"
	"reflect"
)

func TestTokenDecorder(t *testing.T) {
	for _, test := range []struct {
		sexpr  string
		tokens []Token
	}{
		{"( )", []Token{StartList{}, EndList{}}},
		{"(ABC)", []Token{StartList{},
			Symbol{"ABC"},
			EndList{}}},
		{`(ABC "DEF")`, []Token{StartList{},
			Symbol{"ABC"},
			String{"DEF"},
			EndList{}}},
		{`(ABC nil)`, []Token{StartList{},
			Symbol{"ABC"},
			Symbol{"nil"},
			EndList{}}},
		{`("ABC" nil)`, []Token{StartList{},
			String{"ABC"},
			Symbol{"nil"},
			EndList{}}},
		{`("ABC" 10)`, []Token{StartList{},
			String{"ABC"},
			Int{10},
			EndList{}}},
		{`(ABC (x 10))`, []Token{StartList{},
			Symbol{"ABC"},
			StartList{},
			Symbol{"x"},
			Int{10},
			EndList{},
			EndList{}}},
	} {
		//スペースが入るとだめそう。
		t.Run(test.sexpr, func(t *testing.T) {
			//t.Log(test.sexpr)
			//t.Log(test.tokens)
			d := NewDecoder(strings.NewReader(test.sexpr))
			for _, token := range test.tokens {
				result, _ := d.Token()
				if !reflect.DeepEqual(token, result) {
					t.Errorf("token : %v, result : %v", token, result)
				}
			}

		})
	}
}
