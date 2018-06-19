package main

func HelloWorld() string {
	return "Hello World"
}

func Fact(num uint64) uint64  {
	if num == 0 {
		return 1
	} else if num == 1 {
		return 1
	} else {
		return num * Fact(num-1)
	}
}
