package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf("stack after push: %v\n", s)
	s.pop()
	s.pop()
	fmt.Printf("stack after pop: %v\n", s)
	s.pop()
	s.pop()
	fmt.Printf("stack after pop: %v\n", s)
}

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	if s.i-1 < 0 {
		return -1
	}
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
