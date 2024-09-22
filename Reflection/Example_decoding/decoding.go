package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

/*
	For each Marshal function provided by the standard library’s encoding/... packages, there is
	a corresponding Unmarshal function that does decoding. For example, as we saw in
	Section 4.5, given a byte slice containing JSON-encoded data for our Movie type (§12.3), we
	can decode it like this:
	data := []byte{...}
	var movie Movie
	err := json.Unmarshal(data, &movie)

*/

// sexpr

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		switch lex.text() {
		case "nil":
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		case "true":
			v.SetBool(true)
			lex.next()
			return
		case "false":
			v.SetBool(false)
			lex.next()
			return
		}
	case scanner.String:
		s, err := strconv.Unquote(lex.text())
		if err != nil {
			panic(fmt.Sprintf("cannot decode interface name %v", v.Type()))
		}
		switch s {
		case "ptr":
			var stringdata *string
			lex.next()
			value, er := strconv.Unquote(lex.text())
			if er != nil {
				panic(fmt.Sprintf("can not decode ptr name %v", v.Type()))
			}
			stringdata = &value
			elem := reflect.ValueOf(stringdata)
			v.Set(elem)
			lex.next()
			return

		case "[]int":
			var interfacedata string
			lex.next()
			value, err := strconv.Unquote(lex.text())
			if err != nil {
				panic(fmt.Sprintf("cannot decode interface name %v", v.Type()))
			}
			interfacedata = value
			elem := reflect.ValueOf(interfacedata)
			v.Set(elem)
			lex.next()
			return

		default:
			v.SetString(s)
			lex.next()
			return

		}

	case scanner.Int:
		i, _ := strconv.ParseInt(lex.text(), 10, 64)
		v.SetInt(i)
		lex.next()
		return
	case scanner.Float: // Ejercicio 12.10
		i, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(i)
		lex.next()
		return

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
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map: // ((key value) ...)
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

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
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

// Unmarshal function
func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}


// Ejercicio 12.8

// Unmarshal2 read from io.Reader
func Unmarshal2(data []byte, ouput interface{}) {
	newDecoder(bytes.NewReader(data)).decode(ouput)
}

type decoder struct {
	r io.Reader
}

func (d *decoder) decode(out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(d.r)
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
			fmt.Println(err)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

func newDecoder(r io.Reader) *decoder {
	return &decoder{r}
}

// Movie struct
type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	TestInter       interface{}
}

func main() {
	expression := `
	((Title "Dr.Strangelove")
 	 (Subtitle "How I learned to Stop Worrying and Love the Bomb")
     (Year 1964)
 	 (Actor (("Pres. Merking Muffley" "Peter Sellers")
         ("Gen. Buck Turgidson" "George C. Scott")
         ("Brig. Gen. Jack D. Rpper" "Sterling Hayden")
         ("Maj. T.J \"King\" Kong" "Slim Pickens")
         ("Dr. Streangelove" "Peter Sellers")
         ("Grp. Capt. Lionel Mandrake" "Peter Sellers")))
 	 (Oscars ("Best Actor (Nomin.)"
          "Best Adapted Screenplay (Nomin.)"
          "Best Director (Nomin. )"
          "Best Picture (Nomin. )"))
	(Sequel "ptr" "test test")
	(TestInter  "[]int" "1,2,3"))  
	`

	flag := "u2"
	var test Movie
	
	if flag == "u1" {
		Unmarshal([]byte(expression), &test)
		fmt.Println(test)
	} else {
		Unmarshal2([]byte(expression), &test)
		fmt.Println(test)
	}

}
