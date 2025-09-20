package main
import (
	"fmt"
	"log"
	"os"
	"github.com/go-gota/gota/dataframe"
)
func main() {
	file, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	irisDF := dataframe.ReadCSV(file)
	filter := dataframe.F{Colname: "species", Comparator: "==", Comparando: "Versicolor"}
	versicolorDF := irisDF.Filter(filter)
	fmt.Println(versicolorDF)
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(versicolorDF)
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)
}






