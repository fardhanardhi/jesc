package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	start := time.Now()

	format := pflag.BoolP("format", "f", false, "Format json with indentation")
	filePath := pflag.String("file", "", "Input filepath")
	pflag.Parse()
	args := pflag.Args()

	if *filePath != "" && len(args) > 0 {
		fmt.Println("Either pass json directly as inline argument or specify the file path using --file")
		fmt.Println("")
		os.Exit(1)
	}

	var input string

	if *filePath != "" {
		// Open our jsonFile
		jsonFile, err := os.Open(*filePath)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nSuccessfully Opened %s", *filePath)
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		// read our opened xmlFile as a byte array.
		byteValue, _ := io.ReadAll(jsonFile)
		input = string(byteValue)
	} else {
		input = args[0]
	}

	var target map[string]any
	err := json.Unmarshal([]byte(input), &target)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	var data = recusiveFormat(target)
	var jsonBytes []byte
	if *format {
		jsonBytes, _ = json.MarshalIndent(data, "", "    ")
	} else {
		jsonBytes, _ = json.Marshal(data)
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("[INFO] Result:")
	fmt.Println("")
	fmt.Println(string(jsonBytes))
	fmt.Println("")
	fmt.Println("[INFO] JSON successfuly escaped ✨")
	fmt.Printf("[INFO] Execution took %s\n", time.Since(start))
	os.Exit(0)
}

func recusiveFormat(target map[string]any, nest ...string) map[string]any {
	for k, v := range target {
		var va = v
		str, ok := v.(string)
		if ok {
			var noSpaceStr = strings.ReplaceAll(str, " ", "")
			var charBracket = ""
			if len(noSpaceStr) > 0 {
				charBracket = string(noSpaceStr[0])
			}
			if charBracket == "{" {
				var mapTarget map[string]any
				err := json.Unmarshal([]byte(str), &mapTarget)
				if err == nil {
					va = mapTarget
				}
			} else if charBracket == "[" {
				var sliceTarget []map[string]any
				err := json.Unmarshal([]byte(str), &sliceTarget)
				if err == nil {
					va = sliceTarget
				}
			}
		}
		v = va
		m, ok := v.(map[string]any)
		if ok {
			v = recusiveFormat(m, append(nest, k)...)
		}
		target[k] = v
	}
	return target
}
