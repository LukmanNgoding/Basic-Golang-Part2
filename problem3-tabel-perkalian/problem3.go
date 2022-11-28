package main

import (
	"fmt"
	"strconv"
)

func CetakTabelPerkalian(number int) string {
	// your code here
	var res string
	for i := 1; i <= number; i++ {
		for k := 1; k <= number; k++ {
			res += strconv.Itoa(i * k)
			res += "\t"
		}
		res += "\n"
	}
	return res
}

func main() {
	fmt.Println(CetakTabelPerkalian(9))
}
