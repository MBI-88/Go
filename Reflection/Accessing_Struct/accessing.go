package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// search

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels    []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
	}
	data.MaxResult = 10
	if err := Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

// Unpack function
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

// Ejercicio 12.11
// Input Search: {Labels:[golang programming] MaxResults:10 Exact:true} x->Exact, l->Labels, max->MaxResult
// Ouput 'http://localhost:12345/search?x=true&l=golang&l=programming'

// Pack return an url from Umpack
func Pack(prt interface{}) (string, error) {
	val := reflect.ValueOf(prt)
	url, err := handlePack(val)
	if err != nil {
		return "", err
	}

	return url, nil
}

func handlePack(v reflect.Value) (string, error) {
	url := "http://localhost:12345/search?"
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Slice, reflect.Array:
			for x := 0; x < field.Len(); x++ {
				if i == 0 && x == 0 {
					url += fmt.Sprintf("l=%s", field.Index(x))
				} else {
					url += fmt.Sprintf("&l=%s", field.Index(x))
				}

			}
		case reflect.Int, reflect.Int32, reflect.Int64:
			if i == 0 {
				url += fmt.Sprintf("max=%d", field.Int())
			} else {
				url += fmt.Sprintf("&max=%d", field.Int())
			}

		case reflect.Bool:
			if field.Bool() {
				if i == 0 {
					url += fmt.Sprintf("x=%t", true)
				} else {
					url += fmt.Sprintf("&x=%t", true)
				}
			}else {
				if i == 0 {
					url += fmt.Sprintf("x=%t", true)
				} else {
					url += fmt.Sprintf("&x=%t", true)
				}
			}

		default:
			return "", fmt.Errorf("error format")

		}

	}

	return url, nil
}

func main() {
	data := struct {
		Labels    []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
	}{Labels: []string{"golang", "programming"}, MaxResult: 10, Exact: true}

	result, er := Pack(data)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println(result)

}
