package main

import (
	"fmt"
)

func DrawXYZ(n int) string {
	// your code here

	var result = ""
	var xyz int = 1
	for i := 1; i <= n; i++ {
		for k := 1; k <= n; k++ {
			if xyz%3 == 0 {
				result += "X "
			} else if xyz%2 == 1 {
				result += "Y "
			} else if xyz%2 == 0 {
				result += "Z "
			}
			xyz++

		}
		result = result + "\n"
	}
	return result
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}
