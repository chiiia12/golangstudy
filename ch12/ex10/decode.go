package decode

import (
	"text/scanner"
	"fmt"
	"reflect"
	"io"
	"bytes"
	"strconv"
	"log"
	"strings"
)

type lexer struct {
	scan  scanner.Scanner
	token rune //current token
}

type Decoder struct {
	reader io.Reader
}

func NewDecoder(reader io.Reader) (*Decoder) {
	return &Decoder{reader: reader}
}

func Unmarshal(data []byte, out interface{}) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	return decoder.Decode(out)
}
func (decoder *Decoder) Decode(out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(decoder.reader)
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)

		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}
func (lex *lexer) text() string {
	return lex.scan.TokenText()
}
func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
		if lex.text() == "t" {
			v.SetBool(true)
			lex.next()
			return
		}
		if lex.text() == "nil" && v.Type().String() == "bool" {
			v.SetBool(false)
			lex.next()
			return
		}

	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.Float:
		switch(v.Kind()) {
		case reflect.Float32:
			value, _ := strconv.ParseFloat(lex.text(), 32)
			v.SetFloat(value)
			lex.next()
			return
		case reflect.Float64:
			value, _ := strconv.ParseFloat(lex.text(), 64)
			v.SetFloat(value)
			lex.next()
			return
		}
	case '(':
		lex.next()
		readList(lex, v)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	case reflect.Interface:
		str, _ := strconv.Unquote(lex.text())
		log.Println("str is ", str)
		value := reflect.New(typeOf(str)).Elem()
		lex.next()
		read(lex, value)
		v.Set(value)
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

var maps = map[string]reflect.Type{
	"int":    reflect.TypeOf(int(0)),
	"int8":   reflect.TypeOf(int8(0)),
	"int16":  reflect.TypeOf(int16(0)),
	"int32":  reflect.TypeOf(int32(0)),
	"int64":  reflect.TypeOf(int64(0)),
	"uint":   reflect.TypeOf(uint(0)),
	"uint8":  reflect.TypeOf(uint8(0)),
	"uint16": reflect.TypeOf(uint16(0)),
	"uint32": reflect.TypeOf(uint32(0)),
	"uint64": reflect.TypeOf(uint64(0)),
	"bool":   reflect.TypeOf(false),
	"string": reflect.TypeOf("aaa"),
}

func typeOf(name string) reflect.Type {
	t, ok := maps[name]
	if ok {
		return t
	}
	if strings.HasPrefix(name, "[]") {
		return reflect.SliceOf(typeOf(name[2:]))
	}
	if name[0] == '[' {
		i := strings.Index(name, "]")
		if i > 0 {
			len, _ := strconv.Atoi(name[1:i])
			return reflect.ArrayOf(len, typeOf(name[i+1:]))
		}
	}
	if strings.HasPrefix(name, "map") {
		i := strings.Index(name, "]")
		if i > 0 {
			return reflect.MapOf(typeOf(name[4:i]), typeOf(name[i+1:]))
		}
	}
	panic(fmt.Sprintf("%s not supported\n", name))
}
func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
