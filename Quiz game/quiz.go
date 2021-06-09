package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var fileName *string

func init() {
	fileName = flag.String("fname", "", "CSV File containing quiz questions")
	flag.Parse()
}

type FileNotPresent struct {
	When time.Time
	What string
}

type Questions struct {
	expression string
	answer     string
}

func (e *FileNotPresent) Error() string {
	return fmt.Sprintf("Error:- at %v, %s", e.When.Format("01-JAN-2006 15:04:00"), e.What)
}

func readFile(questions_file string) []Questions {
	var questions_list []Questions
	csv_file, err := os.Open(questions_file)

	if err != nil {
		fmt.Println(err)
	}

	defer csv_file.Close()

	quiz_data, err := csv.NewReader(csv_file).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, row := range quiz_data {
		individual_question := Questions{
			expression: row[0],
			answer:     row[1],
		}
		questions_list = append(questions_list, individual_question)
	}

	return questions_list
}

func main() {
	if *fileName == "" {
		err := &FileNotPresent{
			time.Now(),
			"Please provide file name using flag -fname i.e( -fname 'filename')",
		}
		fmt.Println(err)
	}
	// readFile(*fileName)
	fmt.Println(readFile(*fileName))
}
