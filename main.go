package main

import (
	"bufio"
	"github.com/crumbhole/kubecog-helper/src/schema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func loadValues(contents []byte) (schema.CogValues, error) {
	var values schema.CogValues
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
		v := schema.KubecogValidator{}
		print(v.ValidateToSingleString(values))
	} else {
		log.Fatal("Please pipe in a cogvalues.yaml")
	}
}
