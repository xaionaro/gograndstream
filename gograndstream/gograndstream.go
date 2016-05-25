package gograndstream

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var entryNameMap = map[int]string{
	301:	"SpeedDialKey1Account",
	302:	"SpeedDialKey1Name",
	303:	"SpeedDialKey1UserID",
}

func parseSpeedDialParam(pCode int, arg string) (entryName string, entryValue string, err error) {
	entryValue = arg
	entryName  = entryNameMap[pCode]

	if (entryName == "") {
		return "", "", fmt.Errorf("Unknown P-Code: %v", pCode)
	}

	return entryName, entryValue, nil
}

func ParseFile(in *os.File) (map[string]string, error) {
	result := make(map[string]string)

	reader := bufio.NewReader(in)

	for {
		lineBytes, isPrefix, err := reader.ReadLine()
		if (err != nil) {
			break
		}
		if (isPrefix) {
			return nil, fmt.Errorf("line is too long");
		}
		line	  := string(lineBytes)
		words	  := strings.Split(line, "=")
		pCodeStr  := words[0]
		arg	  := words[1]

		if (pCodeStr[0:1] != "P") {
			return nil, fmt.Errorf("Invalid P-Code: \"%v\"", pCodeStr)
		}

		pCode,err := strconv.Atoi(pCodeStr[1:])
		if (err != nil) {
			return nil, err
		}

		switch {
			case pCode >= 6000 && pCode < 7000, pCode >= 300 && pCode < 400:
				entryName, entryValue, err := parseSpeedDialParam(pCode, arg)
				if (err != nil) {
					return nil, err
				}
				result[entryName] = entryValue
				break
		}
	}

	return result, nil
}
func WriteToFile(in *os.File, configuration map[string]string) error {
	return nil
}
