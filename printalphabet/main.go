package main

import "github.com/01-edu/z01"

func main () {
	for i:='a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	} 
	z01.PrintRune('\n')

// alternative way
	// for i:=97; i <= 122; i++{
	// 	z01.PrintRune(rune(i))
	// } 
	// z01.PrintRune('\n')
// alternative way
}


