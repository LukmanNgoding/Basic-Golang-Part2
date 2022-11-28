package main

import "fmt"

func PlayWithAsterix(n int) string {
	// your code here

	var res string
	for i := 1; i <= n; i++ {
		for y := 1; y <= n-i; y++ {
			res += " "

		}
		for z := 1; z <= i; z++ {
			res += "* "
		}
		res += "\n"
	}
	return res

}

func main() {
	fmt.Println(PlayWithAsterix(5))
}
