package main

import (
	"encoding/json"
	"fmt"
	"github.com/droundy/goopt"
	"github.com/xaionaro/gograndstream/grandstream"
	"os"
	"strings"
)

type format int

const (
	FORMAT_UNDEFINED          format = 0
	FORMAT_GRANDSTREAM_PCODES format = 1
	FORMAT_JSON               format = 2
)

var formatMap = map[string]format{"GrandstreamPCodes": FORMAT_GRANDSTREAM_PCODES, "JSON": FORMAT_JSON}

func parseFromTo(out *format, arg string) (err error) {
	*out = formatMap[arg]
	if *out == 0 {
		return fmt.Errorf("Invalid format name: \"%v\"", arg)
	}

	return nil
}

func main() {
	var from format
	var to format

	var possibleToFromValuesSlice []string
	for formatName, _ := range formatMap {
		possibleToFromValuesSlice = append(possibleToFromValuesSlice, formatName)
	}
	possibleToFromValuesString := strings.Join(possibleToFromValuesSlice, ", ")

	goopt.ReqArg([]string{"-f", "--from"}, "formatName", "format of data in stdin   (possible values: "+possibleToFromValuesString+")", func(arg string) error { return parseFromTo(&from, arg) })
	goopt.ReqArg([]string{"-t", "--to"}, "formatName", "format of data for stdout (possible values: "+possibleToFromValuesString+")", func(arg string) error { return parseFromTo(&to, arg) })

	goopt.Description = func() string {
		return "Converter of Grandstream configuration P-codes to more suitable formats"
	}
	goopt.Version = "0.0"
	goopt.Summary = "grandstream P-codes parser"
	goopt.Parse(nil)

	if from == FORMAT_UNDEFINED || to == FORMAT_UNDEFINED {
		panic(fmt.Errorf("--from or --to is not set"))
		os.Exit(-1)
	}

	var configuration map[string]map[string]map[string]string
	switch from {
	case FORMAT_GRANDSTREAM_PCODES:
		var err error
		configuration, err = grandstream.ParseFile(os.Stdin)
		if err != nil {
			panic(err)
		}
		break
	case FORMAT_JSON:
		decoder := json.NewDecoder(os.Stdin)
		decoder.Decode(&configuration)
		break
	}

	switch to {
	case FORMAT_GRANDSTREAM_PCODES:
		err := grandstream.WriteToFile(os.Stdout, configuration)
		if err != nil {
			panic(err)
		}
		break
	case FORMAT_JSON:
		b, err := json.Marshal(configuration)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b)
		break
	}
}
