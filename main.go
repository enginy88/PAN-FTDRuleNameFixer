package main

import (
	"PAN-FTDRuleNameFixer/app"
	"PAN-FTDRuleNameFixer/convert"

	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	start := time.Now()
	app.LogAlways.Println("HELLO MSG: Welcome to PAN-FTDRuleNameFixer v1.1 by EY!")

	if len(os.Args) < 2 {
		app.LogErr.Fatalln("MISSING INPUT: Program argument needed, exiting! (Run with '-h' to see details.)")
	}
	inputArg := flag.String("input", "", "Input file location. (Mandatory!)")
	flag.Parse()
	if *inputArg == "" {
		app.LogErr.Fatalln("MISSING INPUT: Program argument needed, exiting! (Run with '-h' to see details.)")
	}

	convert.RunConvertJobs(*inputArg)

	duration := fmt.Sprintf("%.1f", time.Since(start).Seconds())
	app.LogAlways.Println("BYE MSG: All done in " + duration + "s, bye!")

}
