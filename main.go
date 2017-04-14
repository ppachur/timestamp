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

func main() {
	tNow := time.Now()
	tProvided := tNow
	baseIsCurrentTime := true
	differencePresent := false
	t := tNow

	if timestampFlag != 0 {
		tProvided = time.Unix(timestampFlag, 0)
		t = tProvided
		baseIsCurrentTime = false
	} else if timeFlag != "" {
		const longForm = "2006-01-02T15:04"
		tTemp, err := time.Parse(longForm, timeFlag)
		if err != nil {
			fmt.Println("Unable to parse provided time. Please use exactly this format: 'YYYY-MM-DDTHH:mm'")
			os.Exit(1)
		}
		tProvided = tTemp
		t = tTemp
		baseIsCurrentTime = false
	}

	if diffFlag != 0 {
		t = t.Add(time.Duration(diffFlag) * time.Hour)
		differencePresent = true
	}

	if cleanFlag && !differencePresent {
		if dateFlag {
			fmt.Println(tNow.UTC().Format(time.UnixDate))
		} else {
			fmt.Println(int64(t.Unix()))
		}

		os.Exit(0)
	} else if !cleanFlag {
		fmt.Printf("Current time is %s / %s :: Unix Time: %s\n\n", tNow.Format(time.UnixDate), tNow.UTC().Format(time.UnixDate), strconv.FormatInt(int64(tNow.Unix()), 10))
	}
	if !baseIsCurrentTime {
		if cleanFlag && !differencePresent {
			if dateFlag {
				fmt.Println(tProvided.UTC().Format(time.UnixDate))
			} else {
				fmt.Println(int64(t.Unix()), 10)
			}

			os.Exit(0)
		} else if !cleanFlag {
			fmt.Printf("Provided time is %s / %s :: Unix Time: %s\nDifference to provided time: %f hours or %f days from now\n\n", tProvided.Format(time.UnixDate), tProvided.UTC().Format(time.UnixDate), strconv.FormatInt(int64(tProvided.Unix()), 10), tProvided.Sub(tNow).Hours(), (tProvided.Sub(tNow).Hours() / 24))
		}
	}
	if differencePresent {
		if cleanFlag {
			if dateFlag {
				fmt.Println(t.UTC().Format(time.UnixDate))
			} else {
				fmt.Println(int64(t.Unix()))
			}

			os.Exit(0)
		} else {
			fmt.Printf("Calculated time is %s / %s :: Unix Time: %s\nDifference to calculated time: %f hours or %f days from now\n\n", t.Format(time.UnixDate), t.UTC().Format(time.UnixDate), strconv.FormatInt(int64(t.Unix()), 10), t.Sub(tNow).Hours(), (t.Sub(tNow).Hours() / 24))
		}
	}
}
