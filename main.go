package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	configFilename = "config.json"
	logFilename    = "sample.log"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("cant find working directory")
	}

	var (
		configPath = path + configFilename
		logPath    = path + logFilename
	)

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln("cant open config")
	}

	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		log.Fatalln("cant open log")
	}

	var cfg config
	err = json.Unmarshal(configBytes, &cfg)
	if err != nil {
		log.Fatalln("cant unmarshal config")
	}

	result, _ := json.MarshalIndent(
		parse(logBytes, cfg),
		"", "\t",
	)
	log.Println(string(result))
}

func parse(data []byte, cfg config) (r []parsedOutput) {
	entries := strings.Split(string(data), cfg.Delimiter)

	for i, e := range entries {
		log.Println("parsing: ", e)
		po := parsedOutput{Index: i}

		for _, t := range cfg.Templates {
			if fields := parseEntry(e, t); fields != nil {
				po.Template = t.Name
				po.Fields = fields
				break
			}
		}

		r = append(r, po)
	}

	return
}

func parseEntry(entry string, t template) map[string]string {
	var (
		groups = strings.Split(entry, t.Delimiter)
		r      = make(map[string]string)
	)

	if len(groups) != len(t.Groups) {
		return nil
	}

	for i := 0; i < len(groups); i++ {
		var (
			bytes       = []byte(groups[i])
			groupFields = t.Groups[i].Fields
		)

		for _, f := range groupFields {
			re, err := regexp.Compile(f.Pattern)
			if err != nil {
				log.Fatalln("cant compile pattern: ", f.Pattern)
			}

			var result = string(re.Find(bytes))
			if result == "" {
				return nil
			}

			r[f.Name] = result
		}
	}

	return r
}
