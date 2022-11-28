package main

import "fmt"

func UbahHuruf(sentence string) string {
	// your code here
	var result string
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 'A' && sentence[i] <= 'Z' {
			result += string(((sentence[i]-'A')+10)%26 + 'A')
		} else {
			result += (" ")
		}

	}
	return result
}

func main() {
	fmt.Println(UbahHuruf("SEPULSA OKE"))     // COZEVCK YUO
	fmt.Println(UbahHuruf("ALTERRA ACADEMY")) // KVDOBBK KMKNOWI
	fmt.Println(UbahHuruf("INDONESIA"))       // SXNYXOCSK
	fmt.Println(UbahHuruf("GOLANG"))          // QYVKXQ
	fmt.Println(UbahHuruf("PROGRAMMER"))      // ZBYQBKWWOB
}
