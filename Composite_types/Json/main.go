package main

/*
	The basic JSON types are numbers (in decimal or scientific notation), booleans (true or
	false), and strings, which are sequences of Unicode code points enclosed in double quotes,
	with backslash escapes using a similar notation to Go, though JSON’s \Uhhhh numeric escapes
	denote UTF-16 codes, not runes.
	
	Data structures like this are an excellent fit for JSON, and it’s easy to convertin both
	directions. Converting a Go data structure like movies to JSON is called marshaling. 
	Marshaling is done by json.Marshal.

	You may have noticed that the name of the Year field changed to released in the output, and
	Color changed to color. That’s because of the field tags. A field tag is a string of metadata
	associated at compile time with the field of a struct:

	Year int `json:"released"`
	Color bool `json:"color,omitempty"

	A field tag maybe any literal string , but it is conventionally interpreted as a space-separated
	list of key:"value" pairs; since they contain double quotation marks, field tags are usually
	written with raw string literals. The json key controls the behavior of the encoding/json
	package, and other encoding/... packages follow this convention. The first part of the json
	field tag specifies an alternative JSON name for the Go field. Field tags are often used to
	specify an idiomatic JSON name like total_count for a Go field named TotalCount. The tag
	for Color has an additional option, omitempty, which indicates that no JSON output should
	be produced if the field has the zero value for its type (false, here) or is otherwise empty.
	Sure enough, the JSON output for Casablanca, a black-and-white movie, has no color field.



*/


import (
	_"os"
	_"log"
	_"fmt"
	_"time"
	//"encoding/json"
	_"jsondata/github"
)

/*Movie struct*/
type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string 
}

var movies = []Movie{
	{
		Title: "Casablanca",Year: 1942, Color: false,
	 	Actors: []string{"Humphrey Bogart","Ingrid Bergman"},
	},
	{
		Title: "Cool Hand Luke",Year: 1967, Color: true,
	 	Actors: []string{"Paul Newman"},
	},
	{
		Title: "Bullitt",Year: 1968, Color: true,
	 	Actors: []string{"Steve McQueen","Jacqueline Bisset"},
	},
}



func main(){
	/*
	//data, err := json.Marshal(movies)
	// Variante para mejor entendimiento
	data, err := json.MarshalIndent(movies,"","  ")
	if err != nil {
		log.Fatalf("JSON marshaling failled: %s", err)
	}
	//fmt.Printf("%s\n", data)
	
	var titles []struct{ Title string}
	if err := json.Unmarshal(data,&titles); err != nil {
		log.Fatalf("JSON unmarshaling failed %s\n", err)
	}
	fmt.Println(titles)
	
	// Github issues
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	*/

	// Ejercicio 4.10
	/*
	timeNow := time.Now() 
	beforeMonth := timeNow.AddDate(0,-1,0)
	beforeYear := timeNow.AddDate(-1,0,0)
	
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Less than a year")
	for _, item := range result.Items {
		if item.CreatedAt.Before(beforeYear) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		} 
	}
	fmt.Println("Less than a month")
	for _, item := range result.Items {
		if item.CreatedAt.Before(beforeMonth) || item.CreatedAt.Equal(beforeYear) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
		
	}
	fmt.Println("More than a year")
	for _, item := range result.Items {
		if item.CreatedAt.After(beforeYear){
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	*/
	// Ejercicio 4.11 done 
	// Ejercicio 4.12 done
	// Ejercicio 4.13 done
	
	
}