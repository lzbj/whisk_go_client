package whisk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func Propfile(base string) string {
	if base != "" {
		var filename = fmt.Sprintf("%s/whisk.properties", base)
		if _, err := os.Stat(base); err == nil {
			return filename
		} else {
			var parent = path.Base(base)
			if parent != base {
				return Propfile(parent)
			}
			return ""
		}
	}
	return ""
}

func ImportPropsIfAvailable(file string) map[string]string {
	prop := make(map[string]string)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %v\n", err)
		return nil
	}
	for _, line := range strings.Split(string(data), "\n") {
		var key, value string
		var parts = strings.Split(string(line), "=")
		if len(parts) >= 1 {
			key = parts[0]
		}
		if len(parts) >= 2 {
			value = parts[1]
		}

		if key != "" && value != "" {
			prop[strings.Replace(strings.ToUpper(key), ".", "-", -1)] = value
		}
		if key != "" && value == "" {
			prop[strings.Replace(strings.ToUpper(key), ".", "-", -1)] = ""
		}
	}
	return prop
}

func importDefaultProps() {

}

func importProps(f *os.File) {

}

func updateProps() {

}

func writeProps() {

}

func checkRequiredProperties() {

}

func getPropertyvalue(key string, prop map[string]string) string {
	if evalue := os.Getenv(key); evalue != "" {
		return evalue
	} else if prop[key] != "" {
		return prop[key]
	}
	return ""
}

func merge() {

}
