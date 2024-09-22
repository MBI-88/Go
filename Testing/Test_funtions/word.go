package word

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"unicode"
)

// IsPalindrome function
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// IsPalindrome2 function
func IsPalindrome2(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// Randomized Testing

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// Testing a Command

// echo
var (
	n             = flag.Bool("n", false, "omit trailing newline")
	s             = flag.String("s", " ", "separator")
	out io.Writer = os.Stdout
)

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}

// storage1

func bytesInUse(username string) int64 { return 0 }

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s",username,err)
	}
}

const (
	sender   = "notification@example.com"
	password = "conrrecthorsebatterystaple"
	hostname = "smt.example.com"
	template = `Warning: you are using %d bytes of storage, %d%% of your quota.`
)

// CheckQuota function
func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username,msg)
	/*
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
	*/
}



func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}
