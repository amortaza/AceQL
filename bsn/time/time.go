package bsntime

/*
todo
1. import data
2. export data
3. BR nodes, to pull data from bsn
   like a BR shoudl be able to query a record o more or update records
4. BR scripts to come from BSN
5. work on TODOs
got it?
*/
import (
	"fmt"
	"github.com/amortaza/aceql/logger"
	"strconv"
	"strings"
	"time"
)

func Format(in time.Time) string {
	return in.Format("2006-01-02 15:04:05")
}

func Now() string {
	return Format(time.Now())
}

func Normalize(s string) string {
	if s == "" {
		return s
	}

	s = strings.Trim(s, " ")

	return strings.ReplaceAll(s, "/", "-")
}

func IsAfter(d1, d2 string) bool {
	return d1 >= d2
}

func AddSeconds(timeAsStr, secAsStr string) (string, error) {
	timeAsStr = strings.Trim(timeAsStr, " ")
	secAsStr = strings.Trim(secAsStr, " ")

	seconds, err := strconv.Atoi(secAsStr)
	if err != nil {
		return "", logger.Error(fmt.Sprintf("Cannot parse seconds, see \"%s\"", secAsStr), "bsntime.AddSeconds")
	}

	t, err := time.Parse("2006-01-02 15:04:05", timeAsStr)
	if err != nil {
		return "", logger.Error(fmt.Sprintf("Cannot parse time, see \"%s\"", timeAsStr), "bsntime.AddSeconds")
	}

	next := t.Add(time.Duration(seconds) * time.Second)

	return Format(next), nil
}
