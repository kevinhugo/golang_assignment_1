package main
import "fmt"

func main() {
	maxNumber := 10
	for i:=1 ; i<maxNumber ; i++ {
		if i % 2 == 0 {
			fmt.Println(i," Genap")
		} else {
			fmt.Println(i," Ganjil")
		}
	}
}