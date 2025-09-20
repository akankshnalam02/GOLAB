package main
import (
	"fmt"
	"myproject/mathutil"
)
func main() {
	sum := mathutil.Add(5, 3)
	difference := mathutil.Subtract(10, 4)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}
