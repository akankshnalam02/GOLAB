



package main
import (
	"fmt"
	"log"
)
func Max(nums ...int) int {
	if len(nums) == 0 {
		log.Fatal("No numbers provided")
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}
func main() {
	fmt.Println(Max(1, 2, 3, 4, 5))
	fmt.Println(Max(10, 20, 30, 5, 15, 25))
	fmt.Println(Max(7, 2, 9, 3, 8, 1, 6))
	fmt.Println(Max(-1, -5, -3, -10, -2))
	fmt.Println(Max())
}

