package main 

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const defaultProplemFileName = "proplems.csv"

func main(){
	// os open return pointer to file not file.
	f, err := os.Open(defaultProplemFileName)
	
	if err != nil {
		fmt.Printf("failed to open the file: %v\n", err)
		return
	}
	defer f.Close()
	
	r := csv.NewReader(f)
	questions, err := r.ReadAll()

	if err != nil {
		fmt.Printf("failed to read csv file questions: %v", err)
		return
	}
	correctAnswers := 0
	for i, record := range questions {
		question, correctAnswer := record[0], record[1] 
		fmt.Printf("%d. %s ?\n", i+1, question)
		var answer string
		_, err := fmt.Scan(&answer)
		if err != nil {
			fmt.Printf("Failed to scan input %v", err)
			return
		}
		if answer == correctAnswer {
			correctAnswers++
		}
	}
	fmt.Printf("Result: %d/%d", correctAnswers, len(questions))
}

func startTimer() {
	time.Sleep(30 * time.Second)
}