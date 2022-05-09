package main

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func loadValues(contents []byte) (CogValues, error) {
	var values CogValues
	err := yaml.Unmarshal(contents, &values)
	return values, err
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		filecontents, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
		values, err := loadValues(filecontents)
		if err != nil {
			log.Fatal(err)
		}
		v := kubecogValidator{}
		print(v.ValidateToSingleString(values))
	} else {
		log.Fatal("Please pipe in a cogvalues.yaml")
	}
}
