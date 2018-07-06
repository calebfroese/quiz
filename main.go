package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvr := csv.NewReader(strings.NewReader(string(data)))
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Question", record[0])
		var response int
		fmt.Scan(&response)
		answer, _ := strconv.Atoi(record[1])
		if answer != response {
			fmt.Println("Incorrect")
		} else {
			fmt.Println("Correct")
		}
	}
}
