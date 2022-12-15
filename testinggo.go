package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	x := 2
	y := 3
	z := x + y
	fmt.Println(z)
	myString := "This is a string\nAnother on a new line"
	fmt.Println(myString)
	for k, v := range myString {
		fmt.Printf("k: %d,v: %c == %d\n", k, v, v)
	}
	s := strings.Join([]string{"hello", "world"}, ", ")
	fmt.Println(s)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Your stupid name is " + name)
}
