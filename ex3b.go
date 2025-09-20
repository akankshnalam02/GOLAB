package main
import "fmt"
func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	sli := []int{11, 12, 13, 14, 15, 16}
	fmt.Println("Fourth element of array =", arr[3])
	fmt.Println("Fourth element of slice =", sli[3])
}
