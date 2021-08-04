package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var inputPath = "./input"
var outputFilePath = "./output"

func main() {

	// Read sample-config
	sampleSourceConnectorConfig, err := ioutil.ReadFile(inputPath + "/sample_config.json")

	if err != nil {
		log.Fatal("Unable to read input file "+inputPath+"/sample_config.json", err)
	}

	// Read tables
	tablesFilePath := inputPath + "/tables.csv"

	tablesFile, err := os.Open(tablesFilePath)

	if err != nil {
		log.Fatal("Unable to read input file "+tablesFilePath, err)
	}

	defer tablesFile.Close()

	csvReader := csv.NewReader(tablesFile)

	tablesFileRecords, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+tablesFilePath, err)
	}

	for index, tablesFileRecord := range tablesFileRecords {
		tableName := tablesFileRecord[0]

		sampleSourceConnectorConfig := strings.Replace(string(sampleSourceConnectorConfig), "$tableName", tableName, -1)

		outputFileName := outputFilePath + "/Test_" + tableName + "_" + strconv.Itoa(index) + ".json"

		ioutil.WriteFile(outputFileName, []byte(sampleSourceConnectorConfig), 0644)

		fmt.Println("Write to " + outputFileName + "complete")
	}
}