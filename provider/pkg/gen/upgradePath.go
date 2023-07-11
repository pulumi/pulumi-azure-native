package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateUpgradeWarnings() (map[string]string, error) {
	// Note - though this function will not return an empty map as of v2,
	// none of the affected resources will be in the v2 schema, so any
	// generated deprectation warnings will be ignored.
	outputMap := make(map[string]string)
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// TODO: This version is hardcoded - we should find a way to inspect at least what major
	// version we're building and read the proper file here.
	upgradePathFile := filepath.Join(wd, "versions", "v3-removed-resources.yaml")
	data, err := os.ReadFile(upgradePathFile)
	if err != nil {
		return nil, err
	}

	text := string(data)
	textLines := strings.Split(text, "\n")
	for _, line := range textLines {
		res := strings.Split(line, ": ")
		// This is a best-effort reading of a mostly-structured file.
		// We'll skip things like preambles etc
		if len(res) != 2 {
			continue
		}
		outputMap[res[0]] = fmt.Sprintf(
			"%s is being removed in the next major version of this provider. "+
				"Upgrade to at least %s to guarantee forwards compatibility.", res[0], res[1],
		)
	}
	return outputMap, nil
}
