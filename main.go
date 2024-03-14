// package main

// import "fmt"

// func main() {
// 	fmt.Println("hello world")
// }

package main

import "fmt"

func main() {
	a := 5
	b := 5
	fmt.Println("a + b =", a+b)
}

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	onkar := "hello world and bye world "
// 	fmt.Printf("%d %s", len(onkar), onkar)

// 	trimmed := strings.TrimSpace(onkar) //TrimSpace is used to delete extra space from sentence
// 	fmt.Printf("\n %d %s", len(trimmed), trimmed)

// 	hello := onkar[15:] //extracting substring
// 	fmt.Println("\n", hello)

// 	helloMars := strings.Replace(onkar, "world", "mars", -1) //to replace old word to new word
// 	fmt.Println(helloMars)

// 	helloMoon := "Hello \"Moon\""
// 	fmt.Println(helloMoon)

// 	fmt.Println(strings.Title(onkar)) //To capitalieze first letter of each word

// 	fmt.Println(strings.ToUpper(onkar)) //to make sentece upper case
// }
