

package main
import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)
func main() {
	file, err := os.Open("Employee.csv")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV")
		return
	}
	header := records[0]
	colIndex := make(map[string]int)
	for i, col := range header {
		colIndex[col] = i
	}
	max := 0
	for _, record := range records[1:] {
		value, _ := strconv.Atoi(record[colIndex["Salary"]])
		if value > max {
			max = value
		}
	}
	fmt.Printf("Max Salary: %d\n", max)
}

