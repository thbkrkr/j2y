package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

func main() {
	mode := flag.String("mode", "j2y", "mode: j2y or y2j")
	flag.Parse()

	stat, err := os.Stdin.Stat()
	if err != nil {
		exitWith(err)
	}
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		exitWith(errors.New("Nothing on stdin"))
	}

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		exitWith(err)
	}

	if len(in) == 0 {
		fmt.Printf("Nothing on STDIN", err)
	}

	input := string(in)
	if *mode == "y2j" && (strings.HasPrefix(input, "{") || strings.HasPrefix(input, "[")) {
		exitWith(errors.New("Invalid YAML in input"))
	}

	var bytes []byte
	var obj interface{}
	switch {
	case *mode == "j2y":

		err = json.Unmarshal(in, &obj)
		if err != nil {
			exitWith(err)
		}

		bytes, err = yaml.Marshal(obj)
		if err != nil {
			exitWith(err)
		}

	case *mode == "y2j":

		err = yaml.Unmarshal(in, &obj)
		if err != nil {
			exitWith(err)
		}

		bytes, err = json.Marshal(obj)
		if err != nil {
			exitWith(err)
		}
	}
	if err != nil {
		exitWith(err)
	}

	result := string(bytes)

	if result == "" {
		exitWith(errors.New("Flag -mode invalid: j2y or y2j expected"))
	}

	if *mode == "j2y" && (strings.HasPrefix(result, "{") || strings.HasPrefix(result, "[")) {
		exitWith(errors.New("Invalid JSON in input"))
	}

	fmt.Println(string(result))
	return
}

func exitWith(err error) {
	fmt.Printf("err: %v\n", err)
	os.Exit(1)
}
