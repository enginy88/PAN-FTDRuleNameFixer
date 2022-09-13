package app

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var LogErr *log.Logger
var LogWarn *log.Logger
var LogInfo *log.Logger
var LogAlways *log.Logger

func init() {

	LogErr = log.New(os.Stderr, "(PAN-FTDRuleNameFixer) ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	LogWarn = log.New(os.Stdout, "(PAN-FTDRuleNameFixer) WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	LogInfo = log.New(os.Stdout, "(PAN-FTDRuleNameFixer) INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	LogAlways = log.New(os.Stdout, "(PAN-FTDRuleNameFixer) ALWAYS: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

}

func Typeof(v interface{}) string {

	return fmt.Sprintf("%T", v)

}

func SearchString(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func MatchSubstrings(str string, subs ...string) (bool, int) {

	matchCount := 0
	isCompleteMatch := true

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matchCount += 1
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matchCount
}

func FindBetween(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func FindBefore(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func FindAfter(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}
