package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

const debug = false

func main() {
	start := time.Now()

	inputSample := `{
		"name": "John Doe",
		"age": 15,
		"hobbies": "[{\"name\":\"climbing\",\"isFavorite\":true},{\"name\":\"cycling\",\"isFavorite\":true},{\"name\":\"running\",\"isFavorite\":true}]",
		"esc": "{\"a\":3,\"b\":\"haha\",\"c\":\"{\\\"na\\\":3,\\\"nb\\\":\\\"haha\\\",\\\"nc\\\":\\\"{\\\\\\\"nna\\\\\\\":3,\\\\\\\"nnb\\\\\\\":\\\\\\\"haha\\\\\\\",\\\\\\\"nnc\\\\\\\":\\\\\\\"huhu\\\\\\\"}\\\"}\"}"
	}`

	format := pflag.BoolP("format", "f", false, "Format json with indentation")
	pflag.Parse()
	args := pflag.Args()

	if !debug {
		if len(args) == 0 {
			fmt.Println("Must have an argument")
			return
		} else if len(args) > 1 {
			fmt.Println("Only accept only single argument")
			return
		}
	}

	var input string
	if debug {
		input = inputSample
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
		var t = strings.Repeat("    ", len(nest))
		if ok {
			if debug {
				fmt.Printf("%s%s:\n", t, k)
			}
			v = recusiveFormat(m, append(nest, k)...)
		} else {
			if debug {
				fmt.Printf("%s%s: %v\n", t, k, v)
			}
		}
		target[k] = v
	}
	return target
}
