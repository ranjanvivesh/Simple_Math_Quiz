package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

)

type problems struct {
	q string
	a string
}

func main() {
	fmt.Println("MATH QUIZ GAME")

	csvfilename := flag.String("csv","problems.csv","a csv file for the problems in the format of question,answer")
	flag.Parse()
	
	file,err := os.Open(*csvfilename)

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Failed to open the csv file: %s \n ",*csvfilename)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	
	
	lines,err := r.ReadAll()

	if err!= nil {
		log.Fatal(err)
		fmt.Printf("Failed to Parse the scv file: \n", err)
	}

	problem := parseLine(lines)

	correct := 0
	for i,p := range problem{
		fmt.Printf("The Problem  # %d is %s = \n",i+1,p.q)
		var answer string
		fmt.Scanf("%s\n",&answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %v out of %v \n",correct,len(problem))
}

func parseLine(lines [][]string) []problems{
	ret :=  make([]problems,len(lines))

	for i, line := range lines {
		ret[i] = problems{
			q:line[0],
			a:line[1], 
		}
	}
	return ret
}

