package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type MicroInfo struct {
	Name      string
	Directory string
	Engine    string
}

func Scan(sourceDir string) ([]*MicroInfo, error) {
	files, err := os.ReadDir(sourceDir)

	if err != nil {
		return nil, nil
	}
	var micros []*MicroInfo

	// scan root source dir for a micro
	m, err := scanDir(sourceDir)
	if err != nil {
		return nil, err
	}
	if m != nil {
		micros = append(micros, m)
	}

	// scan subfolders for micros
	for _, file := range files {
		if file.IsDir() {
			m, err = scanDir(filepath.Join(sourceDir, file.Name()))
			if err != nil {
				return nil, err
			}
			if m != nil {
				micros = append(micros, m)
			}
		}
	}

	return micros, nil
}

func scanDir(dir string) (*MicroInfo, error) {
	runtimeDetectors := []runtimeDetector{
		python,
		node,
	}

	for _, scanner := range runtimeDetectors {
		m, err := scanner(dir)
		if err != nil {
			return nil, err
		}
		if m != nil {
			return m, nil
		}
	}
	return nil, nil
}

func main() {
	micros, err := Scan("./")
	if err != nil {
		log.Fatal(err)
	}

	cfg := NewAppConfig("./", "app", micros)
	err = cfg.SaveConfig("./")
	if err != nil {
		log.Fatal(err)
	}

	jsonMicros, err := json.Marshal(micros)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("python2.7 -m pip --version")
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
	fmt.Println(string(jsonMicros))
}
