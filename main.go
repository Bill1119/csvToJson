package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Employee struct {
	ID          string `json:"id"`
	FirstName   string `json:"first name"`
	LastName    string `json:"last name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Role        string `json:"role"`
	Phone       string `json:"phone"`
}

// check errors and if error accures, exit the program
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		// change it into other func
		//os.Exit(1)
	}
}

func parsingCSVData(records [][]string) []Employee {
	var emp Employee
	var employees []Employee
	// parsing the data
	for _, rec := range records {
		emp.ID = rec[0]
		emp.FirstName = rec[1]
		emp.LastName = rec[2]
		emp.Email = rec[3]
		emp.Description = rec[4]
		emp.Role = rec[5]
		emp.Phone = rec[6]
		employees = append(employees, emp)
	}

	return employees
}

func parsingJSONData(records []byte) []Employee {
	var employees []Employee
	// parsing the data
	for _, emp := range employees {
		var row []string
		row = append(row, emp.ID)
		row = append(row, emp.FirstName)
		row = append(row, emp.LastName)
		row = append(row, emp.Email)
		row = append(row, emp.Description)
		row = append(row, emp.Role)
		row = append(row, emp.Phone)
	}
	return employees
}

// check both file and path is correct
func checkValidFile(filename string) (bool, error) {

	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", filename)
	}

	return true, nil
}

// open csv file
func openFile(path *string) *os.File {
	var pathStr string
	pathStr = *path
	_, err := checkValidFile(pathStr)
	checkError(err)
	csvFile, err := os.Open(pathStr)
	checkError(err)
	return csvFile
}

// read csv file
func readCSV(csvFile *os.File) [][]string {

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	checkError(err)

	return records
}

// convert csv to json file
// verb first
// rewrite the json part and separate out the function of save file function
func convertCSVToJSON(records [][]string) []byte {

	var employees = parsingCSVData(records)

	// create json data
	jsonData, err := json.Marshal(employees)
	checkError(err)

	fmt.Println(string(jsonData))
	return jsonData
}

func saveJSON(jsonData []byte) {
	// create json file
	jsonFile, err := os.Create("./result.json")
	checkError(err)

	defer jsonFile.Close()

	// write json data into json file
	jsonFile.Write(jsonData)
	jsonFile.Close()
}

// read json file
func readJSON(jsonFile *os.File) []byte {

	records, err := ioutil.ReadAll(jsonFile)
	checkError(err)

	return records
}

func convertJSONToCSV(records []byte) {

	var employees []Employee

	var err = json.Unmarshal(records, &employees)
	checkError(err)

	csvFile, err := os.Create("./result.csv")
	checkError(err)

	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, emp := range employees {
		var row []string
		row = append(row, emp.ID)
		row = append(row, emp.FirstName)
		row = append(row, emp.LastName)
		row = append(row, emp.Email)
		row = append(row, emp.Description)
		row = append(row, emp.Role)
		row = append(row, emp.Phone)
		writer.Write(row)
	}
	writer.Flush()
}

func main() {

	// input path

	path := flag.String("csv path", "./sample.csv", "Path of the file")
	flag.Parse()
	// google pacakge yyy.openCSV
	csvFile := openFile(path)
	recordsCSV := readCSV(csvFile)
	jsonData := convertCSVToJSON(recordsCSV)
	saveJSON(jsonData)

	pathJSON := flag.String("json path", "./sample.json", "Path of the file")
	flag.Parse()
	jsonFile := openFile(pathJSON)
	recordsJSON := readJSON(jsonFile)
	convertJSONToCSV(recordsJSON)
	//saveCSV(jsonData)

}
