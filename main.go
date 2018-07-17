package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type challenge struct {
	question string
	answer   string
}

func main() {
	qs := parseCSV("./problems.csv")
	correct := 0
	total := 0

	for _, c := range qs {
		color.Yellow(c.question + "?")
		var response string
		fmt.Scan(&response)
		total++
		if c.answer != response {
			color.Red("Incorrect")
		} else {
			color.Green("Correct")
			correct++
		}
	}

	fmt.Println("")
	color.Yellow("Final Score")
	finalScore := strconv.Itoa(correct) + "/" + strconv.Itoa(total)
	color.Yellow(finalScore)
}

func parseCSV(path string) (challenges []challenge) {
	challenges = make([]challenge, 0)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	csvr := csv.NewReader(strings.NewReader(string(data)))
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			// end of input
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		c := challenge{record[0], record[1]}
		challenges = append(challenges, c)
	}
	return
}
