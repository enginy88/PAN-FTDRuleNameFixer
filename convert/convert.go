package convert

import (
	"PAN-FTDRuleNameFixer/app"

	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func RunConvertJobs(filePath string) {

	readFile, err := os.Open(filePath)
	if err != nil {
		app.LogErr.Fatalln(err)
	}

	rulenamesMap := map[string]string{}

	re := regexp.MustCompile("(RULE:\\W)(\\w.*)(\"$)")

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var line = fileScanner.Text()

		isCompleteMatch, _ := app.MatchSubstrings(line, "set rulebase security rules", "description")
		if isCompleteMatch == false {
			continue
		}

		currentRulename := app.FindBetween(line, "set rulebase security rules ", " description")

		matches := re.FindStringSubmatch(line)
		if matches == nil || matches[2] == "" {
			continue
		}

		r := regexp.MustCompile("\\W")
		escapedFutureRulename := r.ReplaceAllLiteralString(matches[2], "_")

		uniqEscapedFutureRulename := escapedFutureRulename
		for i := 0; i < 10000; i++ {
			if i != 0 {
				uniqEscapedFutureRulename = escapedFutureRulename + "-" + strconv.Itoa(i)
			}
			_, ok := rulenamesMap[uniqEscapedFutureRulename]
			if ok == true {
				continue
			}
			break
		}

		rulenamesMap[uniqEscapedFutureRulename] = currentRulename
	}

	err = readFile.Close()
	if err != nil {
		app.LogWarn.Println(err)
	}

	rulenamesMapJSON, err := json.Marshal(rulenamesMap)
	if err != nil {
		app.LogErr.Fatalln(err)
	}
	app.LogInfo.Println("CONVERSION LIST: " + string(rulenamesMapJSON))

	writeFile, err := os.Create(filepath.Dir(filePath) + "/cli-commands.txt")
	if err != nil {
		app.LogErr.Fatalln(err)
	}

	writer := bufio.NewWriter(writeFile)

	for newRulename, currentRulename := range rulenamesMap {
		_, err := writer.WriteString("rename rulebase security rules " + currentRulename + " to " + newRulename + "\n")
		if err != nil {
			app.LogWarn.Println(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		app.LogWarn.Println(err)
	}

	err = writeFile.Close()
	if err != nil {
		app.LogWarn.Println(err)
	}

	app.LogInfo.Println("CREATED FILE: " + "PAN-OS CLI commands needed to be run written on file: " + filepath.Dir(filePath) + "/cli-commands.txt")
}
