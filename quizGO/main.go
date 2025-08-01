package main

import (
	"flag"
	"fmt"
	"time"
	"os"
	"encoding/csv"
)

type problem struct {
	q string
	a string
}

func problemPuller(fileName string) ([]problem, error) {
	//read all the problems from quiz.csv

	//1. open the file
	if fObj, err := os.Open(fileName); err== nil {
		//2 create a new instance of reader
		csvR := csv.NewReader(fObj)
		//3. it will need to read the file
		if clines, err := csvR.ReadAll(); err == nil {
			//4. call the parseProblem function on the lines to get the problems
			return parseProblem(clines), nil
		}else {
			return nil, fmt.Errorf("error reading csv file: %s", err.Error())
		}
	}else {
		return nil, fmt.Errorf("error opening file: %s", err.Error())
	}
}

func parseProblem(lines [][]string) []problem {
	//go over the lines and parse them, with the problem struct
	r := make([]problem, len(lines)) //make a slice of problems of size of length of lines
	for i :=0 ; i<len(lines); i++ {
		r[i] = problem{
			q : lines[i][0],
			a : lines[i][1],
		}
	}
	return r
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	//1. input the name of the file
	fName := flag.String("f", "quiz.csv", "path of csv file")
	//2. set the duration of the timer
	// This defines a command-line flag -t that:

	// Name: "t" → This is the flag name (you pass it as -t when running the program)

	// Default value: 30 → If the user doesn't pass -t, it will default to 30

	// Description: "timer for the quiz" → This shows up when the user runs -h or --help
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()
	//3. pull the problems from the file (calling our problem puller func)
	problems, err := problemPuller(*fName)
	//4. handle the error
	if err!=nil {
		exit(fmt.Sprintf("something went wrong: %v", err))
	}
	//5. create a variable to count our correct answers
	correctAns := 0
	//6. using the duration of the timer, we want to initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	//7. loop thru the problems, print the questions, we'll accept the answers
problemLoop:
    for i, p := range problems {
		
		fmt.Printf("Problem %d: %s\n", i+1, p.q)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansC <- answer
		}()
		select{
		case <- tObj.C: // if the timer runs out
			fmt.Println("Time's up!")
			break problemLoop
		case iAns:= <-ansC:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}
	//8. we'll calculate and print out the result
	fmt.Printf("Your result is: %d/%d\n", correctAns, len(problems))
	
}