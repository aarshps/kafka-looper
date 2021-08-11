package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
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
		if index == 0 {
			continue
		}

		sampleSourceConnectorConfigString := string(sampleSourceConnectorConfig)

		tableName := tablesFileRecord[0]
		tableKey := tablesFileRecord[1]
		tableCategory := tablesFileRecord[2]
		tableAuditColumn := tablesFileRecord[3]

		if tableName != "" {
			sampleSourceConnectorConfigString =
				strings.Replace(sampleSourceConnectorConfigString, "$tableName", tableName, -1)
		}

		if tableKey != "" {
			sampleSourceConnectorConfigString =
				strings.Replace(sampleSourceConnectorConfigString, "$tableKey", tableKey, -1)
		}

		if tableCategory != "" {
			sampleSourceConnectorConfigString =
				strings.Replace(sampleSourceConnectorConfigString, "$tableCategory", tableCategory, -1)
		}

		if tableAuditColumn != "" {
			sampleSourceConnectorConfigString =
				strings.Replace(sampleSourceConnectorConfigString, "$tableAuditColumn", tableAuditColumn, -1)
		}

		outputFileName := outputFilePath + "/MW_FACETS-DBO-TS3-" + tableCategory + "-3H-INCR-" + tableName + ".json"

		ioutil.WriteFile(outputFileName, []byte(sampleSourceConnectorConfigString), 0644)

		log.Println("Write to " + outputFileName + " complete")
	}
}
