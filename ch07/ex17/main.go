package main

import (
	"encoding/xml"
	"os"
	"io"
	"fmt"
	"strings"
	"log"
	"reflect"
)

type Element struct {
	name string
	attr []xml.Attr
}

type Target struct {
	name string
	attr []Attr
}
type Attr struct {
	name  string
	value string
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	//var stack []string
	var stack2 []Target
	var args = os.Args[1:]
	fmt.Printf("%v\n", args)
	targetList := parseArg(args)
	log.Println(targetList)

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			//stack = append(stack, tok.Name.Local)
			var attr []Attr
			for _, v := range tok.Attr {
				a := Attr{
					name:  v.Name.Local,
					value: v.Value,
				}
				attr = append(attr, a)
			}
			stack2 = append(stack2, Target{name: tok.Name.Local, attr: attr})
		case xml.EndElement:
			//stack = stack[:len(stack)-1]
			stack2 = stack2[:len(stack2)-1]
		case xml.CharData:
			//if containsAll(stack, os.Args[1:]) {
			//	fmt.Printf("%s:%s\n", strings.Join(stack, " "), tok)
			//}
			if containsAll2(stack2, targetList) {
				fmt.Printf("%s:%s\n", strings.Join(toStringArray(stack2), " "), tok)
			}
		}
	}
}
func toStringArray(targetList []Target) []string {
	var list []string
	for _, v := range targetList {
		list = append(list, v.name)
	}
	return list
}
func parseArg(args []string) []Target {
	targetList := []Target{}
	for _, v := range args {
		len := len(targetList)
		if strings.Contains(v, "=") {
			s := strings.Split(v, "=")
			a := Attr{
				name:  s[0],
				value: s[1],
			}
			targetList[len-1].attr = append(targetList[len-1].attr, a)
		} else {
			target := Target{
				name: v,
				attr: []Attr{},
			}
			targetList = append(targetList, target)
		}
	}
	return targetList
}
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
func containsAll2(x, y []Target) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].name == y[0].name && reflect.DeepEqual(x[0].attr, y[0].attr) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
