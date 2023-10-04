package test2

import "fmt"

// before updating the code after running the test
// func Hello() string {
// 	return "Hello, world"
// }

// updated code to pass the test
func Hello(name string) string {
	return "Hello "+ name
}

func main() {
	fmt.Println(Hello("world"))
}