package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 2 {
		t := time.Now()
		tUTC := t.UTC()
		if strings.Index(os.Args[1], "+") >= 0 || strings.Index(os.Args[1], "-") >= 0 { // difference provided: get timestamp
			fmt.Printf("Calculating current date %s hour(s).\n", os.Args[1])
			multiplyer, err := strconv.ParseInt(os.Args[1], 10, 64)
			if err != nil {
				fmt.Printf("Error parsing input: %s\n", err)
			}
			d := time.Duration(multiplyer)
			tThen := t.Add(d * time.Hour)
			tThenUTC := tThen.UTC()
			fmt.Printf("Current time is %s / %s :: Unix Time: %s\n", t.Format(time.UnixDate), tUTC.Format(time.UnixDate), strconv.FormatInt(int64(t.Unix()), 10))
			fmt.Printf("Calculated time is %s / %s :: Unix Time: %s\n\n", tThen.Format(time.UnixDate), tThenUTC.Format(time.UnixDate), strconv.FormatInt(int64(tThen.Unix()), 10))
		} else { // timestamp provided, get difference
			fmt.Printf("Calculating time difference between provided timestamp %s an current time.\n", os.Args[1])
			timestamp, err := strconv.ParseInt(os.Args[1], 10, 64)
			if err != nil {
				fmt.Printf("Error parsing input: %s\n", err)
			}
			tThen := time.Unix(timestamp, 0)
			tDiff := tThen.Sub(t)
			tThenUTC := tThen.UTC()
			fmt.Printf("Current time is %s / %s :: Unix Time: %s\n", t.Format(time.UnixDate), tUTC.Format(time.UnixDate), strconv.FormatInt(int64(t.Unix()), 10))
			fmt.Printf("Provided time is %s / %s :: Unix Time: %s\nDifference: %f hours or %f days\n\n", tThen.Format(time.UnixDate), tThenUTC.Format(time.UnixDate), strconv.FormatInt(int64(tThen.Unix()), 10), tDiff.Hours(), (tDiff.Hours() / 24))
		}
	} else { // wrong number of arguments, get help
		fmt.Printf("usage: %[1]s [+|-][hours : int] gives date and timestamp. E.g. '%[1]s +1'\n", os.Args[0])
		fmt.Printf("              [timestamp : int] gives the time difference between a provided unix-timestamp and current time. E.g. '%s 1234567890'\n", os.Args[0])
		os.Exit(1)
	}
}
