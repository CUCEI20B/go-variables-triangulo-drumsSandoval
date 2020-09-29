package main

import "fmt"

func main() {
	var base float32
	var height float32
	fmt.Scanln(&base)
	fmt.Scanln(&height)
	fmt.Println(base * height / 2)
}
