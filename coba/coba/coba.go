package main

import (
	"fmt"
	"strconv"
)

func isPrima(n int) bool {

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}
func faktor(n int) string {
	var res string
	for i := 2; i < n; i++ {
		if n%i == 0 {
			res += fmt.Sprintln(strconv.Itoa(i))
		}

	}
	return res
}

func main() {
	fmt.Println(isPrima(13))
	fmt.Println(faktor(10))

}
