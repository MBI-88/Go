package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// encoder
func encode(buf *bytes.Buffer, v reflect.Value, output string) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Array, reflect.Slice:
		if output == "j" {
			buf.WriteByte('[')
		} else {
			buf.WriteByte('(')
		}
		for i := 0; i < v.Len(); i++ {
			if i > 0 && output != "j" {
				buf.WriteByte('\t')
				buf.WriteByte(' ')
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i), output); err != nil {
				return err
			}
			if i < v.Len()-1 {
				if output == "j" {
					buf.WriteByte(',')
				} else {
					buf.WriteByte('\n')
				}
			}
		}
		if output == "j" {
			buf.WriteByte(']')
		} else {
			buf.WriteByte(')')
		}
	case reflect.Map:
		if output == "j" {
			buf.WriteByte('{')
		} else {
			buf.WriteByte('(')
		}
		for i, key := range v.MapKeys() {
			if i > 0 && output != "j" {
				buf.WriteByte('\t')
				buf.WriteByte(' ')
			}
			if output != "j" {
				buf.WriteByte('(')
			}
			if err := encode(buf, key, output); err != nil {
				return err
			}
			if output == "j" {
				buf.WriteByte(':')
			} else {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.MapIndex(key), output); err != nil {
				return err
			}
			if output != "j" {
				buf.WriteByte(')')
			}
			if i < len(v.MapKeys())-1 {
				if output == "j" {
					buf.WriteByte(',')
				} else {
					buf.WriteByte('\n')
				}
			}
		}

		if output == "j" {
			buf.WriteByte('}')
		} else {
			buf.WriteByte(')')
		}

	case reflect.Struct:
		if output == "j" {
			buf.WriteByte('{')
		} else {
			buf.WriteByte('(')
		}
		for i := 0; i < v.NumField(); i++ {
			if i > 0 && output != "j" {
				buf.WriteByte(' ')
			}
			// Ejercicio 12.6
			if v.Field(i).CanInterface() && reflect.Zero(v.Field(i).Type()).CanInterface() &&  
				reflect.DeepEqual(reflect.Zero(v.Field(i).Type()).Interface(), v.Field(i).Interface()) {
				continue
			}
			if output == "j" {
				fmt.Fprintf(buf, "%q:", v.Type().Field(i).Name)
			} else {
				fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			}
			if err := encode(buf, v.Field(i), output); err != nil {
				return err
			}
			if output != "j" {
				buf.WriteByte(')')
			}
			if i < v.NumField()-1 {
				if output == "j" {
					buf.WriteByte(',')
				} else {
					buf.WriteByte('\n')
				}
			}
		}
		if output == "j" {
			buf.WriteByte('}')
		} 
	case reflect.Bool: // Ejercicio 12.3
		if v.Bool() {
			fmt.Fprintf(buf, "true")
		} else {
			fmt.Fprintf(buf, "false")
		}
	case reflect.Ptr: // Ejercicio 12.3
		if v.IsNil() {
			fmt.Fprintf(buf, "%s", "null")
		} else {
			fmt.Fprintf(buf, "%q", v.Type().Elem())
		}

	case reflect.Complex128, reflect.Complex64: // Ejercicio 12.3
		realpart := real(v.Complex())
		imgpart := imag(v.Complex())
		fmt.Fprintf(buf, "#C(%g,%g)", realpart, imgpart)

	case reflect.Interface: // Ejercicio 12.3
		if v.IsNil() {
			fmt.Fprintf(buf, "%v", nil)
		} else {
			buf.WriteByte('(')
			fmt.Fprintf(buf, "%g (1 2 3)", v.Elem().Type())
			buf.WriteByte(')')
		}
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())

	}
	return nil
}

// marsahlString return a string as an ouput
func marshalString(v interface{}) (string, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), "s"); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func marshalJSON(v interface{}) (string, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), "j"); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// wrapMarshal return an encode data
func wrapMarshal(v interface{}, output string) (result string, er error) {
	switch output {
	case "j":
		result, er = marshalJSON(v) // Ejercicio 12.5
	default:
		result, er = marshalString(v) // Ejerciio 12.4
	}
	return result, er
}

// Movie struct
type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr.Strangelove",
		Subtitle: "How I learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Streangelove":           "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merking Muffley":      "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Rpper":   "Sterling Hayden",
			`Maj. T.J "King" Kong`:       "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin. )",
			"Best Picture (Nomin. )",
		},
		Sequel: nil,
	}

	flag := "s"
	if flag == "s" {
		result, err := wrapMarshal(strangelove, "s")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	} else {
		var testMovie Movie
		result, er := wrapMarshal(strangelove, "j")
		fmt.Println(result)
		if er != nil {
			log.Fatal(er)
		}
		if err := json.Unmarshal([]byte(result), &testMovie); err != nil {
			log.Fatal(err)
		}
		fmt.Println(testMovie)
	}

}
