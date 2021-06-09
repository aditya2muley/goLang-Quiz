package main

import (
	"flag"
	"fmt"
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

func (e *FileNotPresent) Error() string {
	return fmt.Sprintf("Error:- at %v, %s", e.When.Format("01-JAN-2006 15:04:00"), e.What)
}

func main() {
	if *fileName == "" {
		err := &FileNotPresent{
			time.Now(),
			"Please provide file name using flag -fname i.e( -fname 'filename')",
		}
		fmt.Println(err)
	}
	fmt.Println(*fileName == "")
}
