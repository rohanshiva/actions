package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

type Match struct {
	Path         string
	MatchContent string
}

type Detectors struct {
	Matches []Match
	Strict  bool // strict requires all the matches to pass
}

type Framework struct {
	Name      string
	Detectors Detectors
}

var Frameworks = [...]Framework{
	{
		Name: "create-react-app",
		Detectors: Detectors{
			Matches: []Match{
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"react-scripts":\s*".+?"[^}]*}`},
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"react-dev-utils":\s*".+?"[^}]*}`},
			},
			Strict: false,
		},
	},
	{
		Name: "svelte",
		Detectors: Detectors{
			Matches: []Match{
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"svelte":\s*".+?"[^}]*}`},
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"sirv-cli":\s*".+?"[^}]*}`},
			},
			Strict: true,
		},
	},
	{
		Name: "angular",
		Detectors: Detectors{
			Matches: []Match{
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"@angular\/cli":\s*".+?"[^}]*}`},
			},
			Strict: true,
		},
	},
	{
		Name: "svelte-kit",
		Detectors: Detectors{
			Matches: []Match{
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"@sveltejs\/kit":\s*".+?"[^}]*}`},
			},
			Strict: true,
		},
	},
	{
		Name: "next",
		Detectors: Detectors{
			Matches: []Match{
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"next":\s*".+?"[^}]*}`},
			},
			Strict: true,
		},
	},
	{
		Name: "nuxt",
		Detectors: Detectors{
			Matches: []Match{
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"nuxt3?(-edge)?":\s*".+?"[^}]*}`},
			},
			Strict: true,
		},
	},
	{
		Name: "vue",
		Detectors: Detectors{
			Matches: []Match{
				{Path: "package.json", MatchContent: `"(dev)?(d|D)ependencies":\s*{[^}]*"@vue\/cli-service":\s*".+?"[^}]*}`},
			},
			Strict: true,
		},
	},
}

func check(dir string, framework *Framework) (bool, error) {
	passed := false
	for _, match := range (*framework).Detectors.Matches {

		// check to see if the file exists before check for pattern match
		if !fileExists(dir, match.Path) {
			return false, nil
		}

		path := filepath.Join(dir, match.Path)

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return false, err
		}

		pass, _ := regexp.MatchString(match.MatchContent, string(b))

		if !pass && (*framework).Detectors.Strict {
			return false, nil
		}
		if pass {
			passed = true
		}
	}
	return passed, nil
}

func detectFramework(dir string) (string, error) {
	for _, framework := range Frameworks {
		check, err := check(dir, &framework)
		if err != nil {
			return "", err
		}
		if check {
			return framework.Name, nil
		}
	}
	return "node", nil
}
