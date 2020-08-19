package internal

import "fmt"

// Check handles errors quickly
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
