package main

/*
	The simplest way to create an error is by calling errors.New, which returns a new error for
	a given error message. The entire errors package is only four lines long:
	package errors
	func New(text string) error { return &errorString{text} }
	type errorString struct { text string }
	func (e *errorString) Error() string { return e.text }

	Although *errorString may be the simplest type of error, it is far from the only one. For
	example, the syscall package provides Go’s low-level system call API. On many platforms, it
	defines a numeric type Errno that satisfies error, and on Unix platforms, Errno’s Error
	method does a lookup in a table of strings.


	Expression Evaluator
*/

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/scanner"
	//"bufio"
)

// ********************* Eval *************************

// Expr evalutor
type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string // Ejercicio 7.13
}

// Var identifies variables
type Var string

// Eval eval float
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Check errors
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) String() string {
	return string(v)
}

type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
}

type binary struct {
	op   rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operation: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (b binary) String() string {
	return fmt.Sprintf("%s %c %s", b.x, b.op, b.y)
}

type call struct {
	fn   string
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (c call) String() string {
	return fmt.Sprintf("%s(%s)", c.fn, c.args)
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

// Env map
type Env map[Var]float64

// ******************* eval packages *************************
// ********************* Lexer ****************************

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

// **************** Parser ********************

// Parse parses the input string as an arithmetic expression.
func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next() // initial lookahead
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) } // Ejercico 7.14 parseBinaryOrTernary

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
	lhs := parseUnary(lex)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec {
			op := lex.token
			lex.next() // consume operator
			rhs := parseBinary(lex, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

// unary = '+' expr | primary
func parseUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next() // consume '+' or '-'
		return unary{op, parseUnary(lex)}
	}
	return parsePrimary(lex)
}

func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next() // consume Ident
		if lex.token != '(' {
			return Var(id)
		}
		lex.next() // consume '('
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next() // consume ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %s, want ')'", lex.describe())
				panic(lexPanic(msg))
			}
		}
		lex.next() // consume ')'
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // consume number
		return literal(f)

	case '(':
		lex.next() // consume '('
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // consume ')'
		return e
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}

// *********************** Surface ********************************

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // x, y axis range (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
)

var sin30, cos30 = 0.5, math.Sqrt(3.0 / 4.0) // sin(30°), cos(30°)

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func corner(f func(x, y float64) float64, i, j int) (float64, float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) // compute surface height z

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(f, i+1, j)
			bx, by := corner(f, i, j)
			cx, cy := corner(f, i, j+1)
			dx, dy := corner(f, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(Env{"x": x, "y": y, "r": r})
	})
}

// ***************** Ejercicios *************************

// Ejercicio 7.13 Agregar un metodo string a Expr

// Ejercicio 7.14 Definir un nuevo tipo de dato cumputa el
// minimo valor de sus operandos

// Ejercicio 7.15 captura por consola 
	

func main() {
	/*
		exp, err := Parse("(x + 3)/(x - 2) + 5")
		if err != nil {
			fmt.Printf("Error %v",err)
		}
		fmt.Printf("Expression => %v\n",exp.Eval(Env{"x":5,}))

		server := http.NewServeMux()
		server.Handle("/surface",http.HandlerFunc(plot))
		log.Fatal(http.ListenAndServe("localhost:8000",server))

		// Ejercicio 7.13
		exp, err := Parse("(x + 5)/2")
		if err != nil {
			fmt.Printf("Error %sv\n",err)
		}
		fmt.Printf("Evaluated: %s\n", exp)

	*/
	input := os.Args[1:]
	var data string 
	for _, elem := range input {
		data += elem
	}
	dataArr := strings.Split(data,"var")
	expre := dataArr[0]
	fmt.Println(dataArr)
	envString := dataArr[1]
	

	eval, err := Parse(expre)
	if err != nil {log.Fatalf("Error %s\n",err)}
	dict := make(map[Var]float64)
	for i, elem := range envString {
		if string(elem) == "x"{
			dict[Var(elem)],_ = strconv.ParseFloat(string(envString[i+2]),64)
		}else if string(elem) == "y" {
			dict[Var(elem)],_ = strconv.ParseFloat(string(envString[i+2]),64)
		}
	}
	
	fmt.Printf("Result %v\n",eval.Eval(Env(dict)))
}
