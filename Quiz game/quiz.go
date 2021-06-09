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
var timeout *int

func init() {
	fileName = flag.String("fname", "question.csv", "CSV File containing quiz questions")
	timeout = flag.Int("timeout-interval", 30, "Time out in seconds")
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

func (res *ResultSet) ask_questions(questions []Questions, reader bufio.Reader) {
	res.correct_answer_count = 0
	for _, question := range questions {
		fmt.Print(question.expression, "=")
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(input) == question.answer {
			res.correct_answer_count++
		}
	}
}

func display_result(result ResultSet) {
	fmt.Println("-----------------------------******************-----------------------------------------")
	fmt.Printf("Correct answers:-%v \nTotal number of questions:-%v", result.correct_answer_count, result.question_count)
	fmt.Println("\n")
}

func main() {
	result := ResultSet{
		correct_answer_count: 0,
		question_count:       0,
	}
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

	timeout_trigger := time.After(time.Duration(*timeout) * time.Second)
	result_channel := make(chan ResultSet)
	go func() {
		result.ask_questions(questions, *reader)
		result_channel <- result
	}()

	select {
	case <-result_channel:
	case <-timeout_trigger:
		fmt.Println("\nTimeout!!!!\n")
	}
	display_result(result)
}
