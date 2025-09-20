package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
type Station struct {
	ID string `json:"station_id"`
	NumBikesAvailable int `json:"num_bikes_available"`
}
type StationData struct {
	Stations []Station `json:"stations"`
}
func main() {
	jsonFile, err := os.Open("sample.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var stationData StationData
	json.Unmarshal(byteValue, &stationData)
	fmt.Printf("Station ID: %s, Bikes: %d\n", stationData.Stations[0].ID, stationData.Stations[0].NumBikesAvailable)
	fmt.Printf("Total Records: %d\n", len(stationData.Stations))
	for _, station := range stationData.Stations {
		fmt.Printf("Station ID: %s, Bikes: %d\n", station.ID, station.NumBikesAvailable)
	}
}
