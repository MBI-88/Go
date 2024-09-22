package main 

import (
	"time"
	"strings"
	"fmt"
	"reflect"
)

/*

*/

// methods 

// Print function
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n",t)
	for i := 0; i < v.NumMethod(); i++ {
		methThpe := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n",t,t.Method(i).Name,
	strings.TrimPrefix(methThpe.String(),"func"))
	}
}

func main() {
	Print(time.Hour)

}