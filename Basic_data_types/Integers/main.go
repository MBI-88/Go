package main

/*
	In the shift operations x<<n and x>>n, the n op erand determines the number of bit position s
	to shif t and must be unsig ned; the x op erand may be unsig ned or sig ned. Arithmetically, a lef t
	shif t x<<n is equivalent to multiplic ation by 2n and a rig ht shif t x>>n is equivalent to the floor
	of div ision by 2n

	Note the use of two fmt tricks. Usually a Printf format string containing multiple % verbs
	would require the same number of extra operands, but the [1] ‘‘adverbs’’ after % tell Printf to
	use the first operand over and over again. Second, the # adverb for %o or %x or %X tells Printf
	to emit a 0 or 0x or 0X prefix respectively.
*/

import "fmt"



func main(){
	var x uint8 = 1 << 1 | 1 << 5
	var y uint8 = 1 << 1 | 1 << 2

	fmt.Printf("%08b\n",x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n",y) // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x & y) // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x | y) // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x ^ y) // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x &^ y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x & (1 << i) != 0 {
			fmt.Println(i) // 1 , 5
		}
	}

	fmt.Printf("%08b\n", x << 1) // "01000100", the set {2, 6}
	fmt.Printf("%0b\n", x >> 1) // "00010001", the set {0, 4}

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n",o) // "438 666 0666"

	z := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", z) // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	ascii := 'a'
	unicode := '@'
	newline := '\n'
	
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // 97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 @ '@'"
	fmt.Printf("%d %[1]q\n", newline) // "10" '\n'


}