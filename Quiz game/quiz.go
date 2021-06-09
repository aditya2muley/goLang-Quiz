package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var fileName *string
var result ResultSet

func init() {
	fileName = flag.String("fname", "", "CSV File containing quiz questions")
	flag.Parse()

	result = ResultSet{
		correct_answer_count: 0,
		question_count:       0,
	}
}

type FileNotPresent struct {
	When time.Time
	What string
}

type Questions struct {
	expression string
	answer     string
}

type ResultSet struct {
	correct_answer_count int
	question_count       int
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
	reader := bufio.NewReader(os.Stdin)

	if *fileName == "" {
		err := &FileNotPresent{
			time.Now(),
			"Please provide file name using flag -fname i.e( -fname 'filename')",
		}
		fmt.Println(err)
	}
	questions := readFile(*fileName)
	result.question_count = len(questions)

	for _, question := range questions {
		fmt.Println(question.expression, "=")
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(input) == question.answer {
			result.correct_answer_count = result.correct_answer_count + 1
		}

	}
	fmt.Println(result)
}
