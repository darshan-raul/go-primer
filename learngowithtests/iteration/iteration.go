package iteration

import "fmt"

func Repeat(letter string)string{
	result := letter
	for i:=0; i< 4;i++{
		result += letter
	}
	return result
}

func main(){
	fmt.Printf("main func")
}