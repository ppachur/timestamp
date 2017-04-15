package main

import (
	"flag"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetBaseTime(t *testing.T) {
	Convey("aquires baseTime", t, func() {

		Convey("when no timestamp or time are provided", func() {
			result := getBaseTime()
			So(result, ShouldHaveSameTypeAs, int64(time.Now().Unix()))
		})

		Convey("when time is provided", func() {
			timeFlag = "2016-01-01T05:00"
			result := getBaseTime()
			So(result, ShouldEqual, 1451624400)
		})

		Convey("when timestamp is provided", func() {
			timestampFlag = 1492267919
			timeFlag = ""
			result := getBaseTime()
			So(result, ShouldEqual, 1492267919)
		})

		Convey("panic when both timestamp and time are provided", func() {
			timestampFlag = 1492267919
			timeFlag = "2016-01-01T05:00"
			So(func() { getBaseTime() }, ShouldPanic)
		})
		timeFlag = ""
	})
}

func TestGetResultTime(t *testing.T) {
	Convey("correctly add / substract flag --diff", t, func() {

		Convey("when --diff is 0", func() {
			diffFlag = 0
			baseTime := getBaseTime()
			result := getResultTime(baseTime)
			So(result, ShouldEqual, int64(time.Unix(baseTime, 0).Add(time.Duration(diffFlag)*time.Hour).Unix()))
		})

		Convey("when --diff is negative", func() {
			diffFlag = -5
			baseTime := getBaseTime()
			result := getResultTime(baseTime)
			So(result, ShouldEqual, int64(time.Unix(baseTime, 0).Add(time.Duration(diffFlag)*time.Hour).Unix()))
		})
	})
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	flag.Parse()
	os.Exit(m.Run())
}
