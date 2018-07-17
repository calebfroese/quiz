package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type challenge struct {
	question string
	answer   string
}

type score struct {
	total     int
	correct   int
	incorrect int
}

var s = score{0, 0, 0}
var countdown = flag.Int("countdown", 30, "amount of seconds until game over")
var qPath = flag.String("questions", "./problems.csv", "path to the csv questions")

func main() {
	flag.Parse()

	onboard()
	qs := parseCSV(*qPath)
	go timer(*countdown)
	for _, c := range qs {
		color.Yellow(c.question + "?")
		var response string
		fmt.Scan(&response)
		if c.answer != response {
			color.Red("Incorrect")
		} else {
			color.Green("Correct")
			s.correct++
		}
	}
	gameEnd()
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
	s.total = len(challenges)
	return
}

func onboard() {
	color.Green("Welcome to the gophercises quiz test.")
	color.White("You will be tested with a series of numeric questions.")
	color.White("Press enter when you are ready to begin")
	fmt.Scanln()
}

func timer(seconds int) {
	defer os.Exit(0)
	time.Sleep(time.Duration(seconds) * time.Second)
	color.Red("Out of time!")
	gameEnd()
}

func gameEnd() {
	color.White("")
	color.Yellow("Final Score")
	finalScore := strconv.Itoa(s.correct) + "/" + strconv.Itoa(s.total)
	color.Yellow(finalScore)
}
