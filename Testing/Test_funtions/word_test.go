package word

import (
	"strings"
	"fmt"
	"math/rand"
	"testing"
	"time"
	"bytes"
)

/*
	The -v flag prints the name and
	execution time of each test in the package: go test -v

	and the -run flag , whose argument is a regular expression, causes
	go test to run only those
	tests whose function name matches the pattern:
	go test -v -run="French|Canal"

	Test failure messages are usually of the form "f(x) = y, want z", where f(x) explains the
	attempted operation and its input, y is the actual result, and z the expected result. Where convenient,
	as in our palindrome example, actual Go syntax is used for the f(x) part. Displaying
	x is particularly important in a table-driven test, since a given assertion is executed many
	times wit h dif ferent values. Avoid boi ler plate and redundant infor mat ion. When testing a
	boole an function such as IsPalindrome, omit the want z part since it adds no information. If
	x, y, or z is lengthy, print a concise summary of the relevant parts instead. The author of a test
	should strive to help the programmer who must diagnose a test failure


	White-Box Testing

	We’ve already seen examples of both kinds. TestIsPalindrome calls only the exported function IsPalindrome and is thus a black-box test. 
	TestEcho calls the echo function and
	up dates the global variable out, both of which are unexported, making it a white-box test.

	This pattern can be used to temporarily save and restore all kinds of global variables, including
	command-line flags, debugging options, and performance parameters; to install and remove
	hooks that cause the production code to call some test code when something interesting happens; 
	and to coax the production code into rare but important states, such as timeouts, errors,
	and even specific interleavings of con current activities.

	External Test Packages

	Writing Effective Tests

	The assertion function below compares two values, constructs a generic error message , and
	stops the program. It’s easy to use and it’s correct, but when it fails, the error message is almost
	useless. It does not solve the hard problem of providing a good user interface.

	Avoiding Brittle Tests 

	The easiest way to avoid brittle tests is to check only the properties you care about. Test your
	program’s simpler and more stable interfaces in preference to its internal functions. Be selective 
	in your assertions. Don’t check for exact string matches, for example, but look for relevant
	substrings that will remain unchanged as the program evolves. It’s often worth writing a
	substantial function to distill a complex output down to its essence so that assertions will be
	reliable. Even though that may seem like a lot of up-front effort, it can pay for itself quickly in
	time that would otherwise be spent fixing spuriously failing tests.
	
*/

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

func TestCanalPalindrome2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome

	}
	for _, test := range tests {
		if got := IsPalindrome2(test.input); got != test.want {
			t.Errorf("IsPalindrome2(%q) = %v", test.input, got)
		}
	}
}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome2(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestEcho(t *testing.T) {
	tests := []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v, %q, %q)",test.newline,test.sep,test.args)
		out = new(bytes.Buffer)
		if err := echo(test.newline,test.sep,test.args); err != nil {
			t.Errorf("%s failed: %v",descr,err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q",descr,got, test.want)
		}
	}
}


func TestCheckQuotaNotifiesUser(t *testing.T) {
	saved := notifyUser
	defer func(){notifyUser = saved}()
	var notifiedUser, notifiedMsg string 
	notifyUser = func (user, msg string) {
		notifiedUser, notifiedMsg = user,msg
	}
	const user = "joe@example.com"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("Wrong user (%s) notified, want %s",notifiedUser,user)
	}

	const wantSubtring = "98% of your quota"
	if !strings.Contains(notifiedMsg,wantSubtring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q",notifiedMsg,wantSubtring)
	}
}

func assertEqual(x,y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d",x,y))
	}

}

func TestSplit(t *testing.T) {
	words := strings.Split("a:b:c",":")
	assertEqual(len(words),3)
}