package main

func HelloWorld() string {
	return "Hello World"
}

func Fib(num,a,b uint64) uint64  {
	if num == 0 {
		return a
	} else if num == 1 {
		return b
	} else {
		return Fib(num-1,b,a+b)
	}
}
