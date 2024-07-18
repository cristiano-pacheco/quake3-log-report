package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cristiano-pacheco/quake3-log-report/internal/logparser"
)

func main() {
	// Define a command-line flag for the file path
	filePath := flag.String("file", "quake3.log", "path to the quake3 log file")
	outputType := flag.String("outputType", "ranking", "ranking|report")
	flag.Parse()

	if *outputType != "ranking" && *outputType != "report" {
		fmt.Println(fmt.Errorf("the mode %s is invalid", *outputType))
		os.Exit(1)
	}

	matches, err := logparser.ParseQuakeLog(*filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *outputType == "ranking" {
		output, err := logparser.MatchesToGameRankingJSON(matches)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(output)
		return
	}

	if *outputType == "report" {
		output, err := logparser.MatchesToGameDeathCausesJSON(matches)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(output)
		return
	}
}
