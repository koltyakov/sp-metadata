package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func genDoc(templatePath string, filePath string, data interface{}) error {
	tData, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}

	t, err := template.New(templatePath).Parse(fmt.Sprintf("%s", tData))
	if err != nil {
		return err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, data); err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(&b)
	if err != nil {
		return err
	}

	folderPath := filepath.Dir(filePath)
	_ = os.MkdirAll(folderPath, os.ModePerm)

	file, _ := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	_ = file.Close()

	trimmed := fmt.Sprintf("%s", bytes)
	trimmed = strings.TrimSpace(trimmed)
	trimmed = strings.Replace(trimmed, "\n\n\n", "\n\n", -1)

	if err := ioutil.WriteFile(filePath, []byte(trimmed), 0644); err != nil {
		return err
	}

	return nil
}
