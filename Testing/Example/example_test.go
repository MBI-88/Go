package example

import (
	"fmt"
)

/*
	The third kind of function treated specially by go test is an example function, one whose
	name starts with Example. It has neither parameters nor results. Here’s an example function
	for IsPalindrome.

	Based on the suffix of the Example function, the web-based documentation server godoc
	associates example functions with the function or package they exemplify, so ExampleIsPalindrome 
	would be shown with the documentation for the IsPalindrome function, and an
	example function called just Example would be associated with the word package as a whole.

	The second purpose is that examples are executable tests run by gotest. If the example function contains a final 
	// Output: comment like the one above , the test driver will execute the
	function and check that what it printed to its standard output matches the text within the
	comment.
*/

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}