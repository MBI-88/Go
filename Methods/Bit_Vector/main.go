package main

/*


*/


import (
	"bytes"
	"fmt"
)

// instset

// IntSet struct
type IntSet struct {
	words []uint64
}

func (s *IntSet) has(x int) bool {
	word, bit := x/64, uint(x % 64)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

func (s *IntSet) add(x int) {
	word, bit := x / 64, uint(x % 64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) unionwith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = tword
		}else{
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
			}
			fmt.Fprintf(&buf, "%d", 64 * i + j)
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Ejercicio 6.1

func (s *IntSet) len() int {
	var count int
	for _, w := range s.words {
		for w != 0 {
			w &= w - 1
			count++
		}
	}
	return count
}

func (s *IntSet) remove(x int)  {
	word, bit := x / 64 , uint(x % 64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) clear() {
	s.words = []uint64{}

}

func (s *IntSet) copy() *IntSet {
	copyW := &IntSet{}
	for _,w := range s.words {
		copyW.words = append(copyW.words,w)
	}
 return copyW
}

// Ejercicio 6.2

func (s *IntSet) addVariadic(values...int) {
	for _, word := range values {
		w, bit := word / 64, uint(word % 64)
		for w >= len(s.words) {
			s.words = append(s.words,0)
		}
		s.words[w] |= 1 << bit

	}

}

// Ejercicio 6.3

func (s *IntSet) intersectWith(t *IntSet)  {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}else {
			s.words = append(s.words,tword)
		}
	}
}

func (s *IntSet) differenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}else {
			s.words = append(s.words,tword)
		}
	}

}

func (s *IntSet) symetricdiff(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}else {
			s.words = append(s.words,tword)
		}
	}
	

}

// Ejercicio 6.4

func (s *IntSet) elem() []int {
	var output []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				output = append(output, 64*i+j)
			}
		}
	}

	return output

}

// Ejercicio 6.5

// BitVector using uint
type BitVector struct {
	words []uint
}

const size = 32 << (^uint(0) >> 63)

func (b *BitVector) addVariadic(values...int) {
	for _, word := range values {
		w, bit := word / size, uint(word % size)
		for w >= len(b.words) {
			b.words = append(b.words,0)
		}
		b.words[w] |= 1 << bit

	}
}

func (b *BitVector) String() string {
	var buff bytes.Buffer
	buff.WriteByte('{')
	for i, word := range b.words {
		if word == 0 {
			continue
		}
		for j := 0; i < size; i++ {
			fmt.Fprintf(&buff, "%d, ", 64 * i + j)
		}
		
	}
	buff.WriteByte('}')
	return buff.String()

}


func main() {
	/*
	var x, y IntSet
	x.add(1)
	x.add(144)
	x.add(9)
	fmt.Println(x.String())
	y.add(9)
	y.add(42)
	fmt.Println(y.String())
	x.unionwith(&y)
	fmt.Println(x.String())
	fmt.Println(x.has(9),x.has(123))

	var x IntSet
	x.add(144)
	x.add(20)
	fmt.Println(x.len())
	//x.remove(20)
	fmt.Println(x.len())
	x.clear()
	fmt.Println(x.len())
	x.add(12)
	x.add(10)
	fmt.Println(x.copy())

	vector := IntSet{}
	vector.addVariadic(20,30,40,100)
	fmt.Println(vector.len())

	var s,t IntSet
	s.addVariadic(10,50,30)
	t.addVariadic(20,40,30)
	//s.intersectWith(&t)
	//s.differenceWith(&t)
	s.symetricdiff(&t)
	fmt.Println(s.words)

	var s IntSet
	s.addVariadic(10,20,30,45)
	fmt.Println(s.elem())

	*/
	var b BitVector
	b.addVariadic(1,2,3)
	fmt.Println(b.String())
	
	
	

	
	


}