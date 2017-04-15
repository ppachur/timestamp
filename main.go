package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var diffFlag int64
var timeFlag string
var timestampFlag int64
var cleanFlag bool
var dateFlag bool

func init() {
	flag.Int64Var(&diffFlag, "diff", 0, "difference relative to current or provided (if present) time in hours. Eg.: -1")
	flag.StringVar(&timeFlag, "time", "", "a date provided similar to '2006-01-22T15:04'")
	flag.Int64Var(&timestampFlag, "timestamp", 0, "a date provided as a unix timestamp")
	flag.BoolVar(&cleanFlag, "clean", false, "produce 'clean' output for scripting, defaults to false")
	flag.BoolVar(&dateFlag, "date", false, "print output as date when using --clean option, defaults to false")
	flag.Parse()
}

func getBaseTime() int64 {
	var baseTime int64
	if timeFlag != "" && timestampFlag != 0 {
		panic("cannot use both timestamp and time")
	} else if timeFlag != "" {
		const longForm = "2006-01-02T15:04"
		tTemp, err := time.Parse(longForm, timeFlag)
		if err != nil {
			fmt.Println("Unable to parse provided time. Please use exactly this format: 'YYYY-MM-DDTHH:mm'")
			os.Exit(1)
		}
		baseTime = int64(tTemp.Unix())

	} else if timestampFlag != 0 {
		baseTime = timestampFlag
	} else {
		baseTime = int64(time.Now().Unix())
	}
	return baseTime
}

func getResultTime(value int64) int64 {
	var resultTime int64
	if diffFlag != 0 {
		resultTime = int64(time.Unix(value, 0).Add(time.Duration(diffFlag) * time.Hour).Unix())
	} else {
		resultTime = value
	}
	return resultTime
}

func printOutput(value int64, resultType string) {
	if cleanFlag && !dateFlag {
		fmt.Println(value)
	} else if cleanFlag && dateFlag {
		fmt.Println(time.Unix(value, 0).Format(time.UnixDate))
	} else {
		tNow := time.Now()
		resultTime := time.Unix(value, 0)
		if resultType == "current time" {
			fmt.Printf("%s: %s / %s :: Unix Time: %s\n\n", resultType, resultTime.Format(time.UnixDate), resultTime.UTC().Format(time.UnixDate), strconv.FormatInt(int64(resultTime.Unix()), 10))
		} else {
			fmt.Printf("%s: %s / %s :: Unix Time: %s\ndiff to %s: %f hours or %f days from now\n\n", resultType, resultTime.Format(time.UnixDate), resultTime.UTC().Format(time.UnixDate), strconv.FormatInt(int64(resultTime.Unix()), 10), resultType, resultTime.Sub(tNow).Hours(), (resultTime.Sub(tNow).Hours() / 24))
		}
	}
}

func main() {
	base := getBaseTime()
	if diffFlag != 0 && (timeFlag == "" || timestampFlag == 0) && !cleanFlag {
		printOutput(base, "current time")
	}
	if (timeFlag != "" || timestampFlag != 0) && !cleanFlag {
		printOutput(base, "provided")
	}

	result := getResultTime(base)
	printOutput(result, "calculated")
	os.Exit(0)

}
