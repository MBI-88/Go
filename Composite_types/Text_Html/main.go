package main

import (
	"os"
)

/*
	A template is a string or file containing one or more portions enclosed in double braces,
	{{...}}, called actions. Most of the string is printed literally, but the actions trigger other
	behaviors. Each action contains an expression in the template language, a simple but powerful
	notation for printing values, selecting struct fields, calling functions and methods, expressing
	control flow such as if-else statements and range loops, and instantiating other templates.
	A simple template string is shown below:

	const templ = `{{.TotalCount}} issues:
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}`

	This template first prints the number of matching issues, then prints the number, user, title,
	and age in days of each one. Within an action, there is a notion of the current value, referred
	to as ‘‘dot’’ and written as ‘‘.’’, a period. The dot initially refers to the template’s parameter,
	which will be a github.IssuesSearchResult in this example. The {{.TotalCount}} action
	expands to the value of the TotalCount field, printed in the usual way. The
	{{range .Items}} and {{end}} actions create a loop, so the text between them is expanded
	multiple times, with dot bound to successive elements of Items.
	
	Within an action, the | notation makes the result of one operation the argument of another,
	analogous to a Unix shell pipeline. In the case of Title, the second operation is the printf
	function, which is a built-in synonym for fmt.Sprintf in all templates. For Age, the second
	operation is the following function, daysAgo, which converts the CreatedAt field into an
	elapsed time, using time.Since:

	func daysAgo(t time.Time) int {
		return int(time.Since(t).Hours() / 24)
	}


	
*/


import (
	//"fmt"
	//"templates/github"
	"html/template"
	"time"
	"log"
)

const templ = `{{.TotalCount}} issues:
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

var issueList = template.Must(
	template.New("issuelist").Parse(
		`
		<h1>{{.TotalCount}} issues</h1>
		<table>
		<tr style='text-align: left'>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
			<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
			<td>{{.State}}</td>
			<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
			<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
		</table>
		`))

func main(){
	// Plain Text
	/*	
	// Variante 1
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
		
	if err != nil {
		log.Fatal(err)
	}
	
	// Variante 2
	result, er := github.SearchIssues(os.Args[1:])
	if er != nil {
		log.Fatal(er)
	}
	if errors := report.Execute(os.Stdout, result);
	   errors != nil {
		log.Fatal(errors)
	   }
	
	*/
	
	// HTML
	/*
	response,_ := github.SearchIssues(os.Args[1:])
	if err := issueList.Execute(os.Stdout,response);
	   err != nil {
		log.Fatal(err)
	   }
	*/

	// Diferentes uso string | html
	
	const tm = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("scape").Parse(tm))
	var data struct{
		A string 
		B template.HTML
	}
	data.A = "<p>Hello</p>"
	data.B = "<p>Hello</p>"

	if err := t.Execute(os.Stdout,data); err != nil {
		log.Fatal(err)
	}

	// Ejercicio 4.14



}