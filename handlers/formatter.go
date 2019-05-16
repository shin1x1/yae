package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type YamlFormatter struct {
	filepath string
}

func NewYamlFormatter(filepath string) *YamlFormatter {
	return &YamlFormatter{
		filepath: filepath,
	}
}

func (y *YamlFormatter) Run(w io.Writer) {
	var yml interface{}

	in, err := ioutil.ReadFile(y.filepath)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	if err := yaml.Unmarshal(in, &yml); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	out, err := yaml.Marshal(yml)
	if err != nil {
		log.Fatalf("marshall error: %v", err)
	}

	_, _ = fmt.Fprint(w, string(out))
}
