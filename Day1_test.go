package main

import (
	"testing"
	"strings"
)

func TestHelloWorld(t *testing.T)  {
	const expected = "Hello World"
	actual := HelloWorld()
	result := strings.Compare(expected,actual)
	if result != 0 {
		t.Errorf("Expected the string <%s>. Found <%s>",expected,actual)
		t.Fail()
	}
}

func TestFibonacci(t *testing.T)  {
	const expected1 = 12586269025
	actual := Fib(50,0,1)
	if actual != expected1 {
		t.Errorf("Expected the num <%d>. Found <%d>",expected1,actual)
		t.Fail()
	}
	const expected2  = 0
	actual = Fib(0,0,1)
	if actual != expected2 {
		t.Errorf("Expected the num <%d>. Found <%d>",expected2,actual)
		t.Fail()
	}
}