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

func main() {
	data, err := ioutil.ReadFile("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	correct := 0
	total := 0
	csvr := csv.NewReader(strings.NewReader(string(data)))
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		q := record[0] + "?"
		color.Yellow(q)
		var response int
		fmt.Scan(&response)
		answer, _ := strconv.Atoi(record[1])
		total++
		if answer != response {
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
