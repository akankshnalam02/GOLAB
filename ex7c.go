package main
import (
	"fmt"
	"io/ioutil"
	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
)

func Run() error {
	xTrain, yTrain, err := base.LoadDataFromCSV("studentsTrain.csv")
	if err != nil {
		return err
	}
	xTest, yTest, err := base.LoadDataFromCSV("studentsTest.csv")
	if err != nil {
		return err
	}
	model := linear.NewLogistic(base.BatchGA, 0.0001, 0.0, 1000, xTrain, yTrain)
	model.Output = ioutil.Discard
	err = model.Learn()
	if err != nil {
		return err
	}
	accuracy := evaluateAccuracy(model, xTest, yTest)
	fmt.Printf("Test Accuracy: %.2f%%\n", accuracy*100)
	return nil
}

func evaluateAccuracy(model *linear.Logistic, xTest [][]float64, yTest []float64) float64 {
	correctPredictions := 0
	for i := range xTest {
		prediction, err := model.Predict(xTest[i])
		if err != nil {
			fmt.Println("Error during prediction:", err)
			continue
		}
		predictedClass := 0.0
		if prediction[0] >= 0.5 {
			predictedClass = 1.0
		}
		if predictedClass == yTest[i] {
			correctPredictions++
		}
	}
	return float64(correctPredictions) / float64(len(yTest))
}

func main() {
	err := Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
