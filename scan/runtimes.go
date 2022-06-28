package main

import (
	"path/filepath"
)

type runtimeDetector func(dir string) (*MicroInfo, error)

func python(dir string) (*MicroInfo, error) {
	// if any of the following files exist detect as python app
	if !checkFiles("requirements.txt", "Pipfile", "setup.py")(dir) {
		return nil, nil
	}

	m := &MicroInfo{
		Name:      filepath.Base(dir),
		Directory: dir,
		Engine:    "py3.10",
	}

	return m, nil
}

func node(dir string) (*MicroInfo, error) {
	// if any of the following files exist detect as a node app
	if !checkFiles("package.json")(dir) {
		return nil, nil
	}

	m := &MicroInfo{
		Name:      filepath.Base(dir),
		Directory: dir,
		Engine:    "node",
	}

	framework, err := detectFramework(dir)

	if err != nil {
		return nil, err
	}

	m.Engine = framework

	return m, nil
}