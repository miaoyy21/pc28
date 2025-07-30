package ifs

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
	"text/template"
)

func Exec(name string, inData interface{}, outData interface{}) error {
	tpl := template.Must(template.ParseFiles(name))

	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, inData); err != nil {
		return err
	}

	log.Println("\n", buf.String())
	out, err := exec.Command("/bin/bash", "-c", buf.String()).Output()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(out, outData); err != nil {
		return err
	}

	return nil
}
