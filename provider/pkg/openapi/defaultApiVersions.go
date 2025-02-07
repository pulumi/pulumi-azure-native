package openapi

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func readDefaultVersionFromReadme(readme io.Reader) (string, error) {
	var version string
	var inYamlBlock bool

	scanner := bufio.NewScanner(readme)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "```yaml" || line == "``` yaml" {
			inYamlBlock = true
		}
		if inYamlBlock && line == "```" {
			inYamlBlock = false
		}

		if inYamlBlock && strings.HasPrefix(line, "tag: package-") {
			version = strings.TrimSpace(strings.TrimPrefix(line, "tag: package-"))
			break
		}
	}

	// if version == "" {
	// 	return "", fmt.Errorf("no version found in readme")
	// }

	// Those occur in sql, security, and synapse
	// if strings.HasPrefix(version, "composite-") {
	// 	return "", fmt.Errorf("composite versions are not supported yet")
	// }
	return version, nil
}

func ReadAllDefaultVersionsFromReadmes(specBaseDir string) (map[string]string, error) {
	versions := make(map[string]string)

	pattern := filepath.Join(specBaseDir, "specification", "*", "resource-manager", "readme.md")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		pathParts := strings.Split(file, string(filepath.Separator))
		moduleFolder := pathParts[len(pathParts)-3]

		version, err := readDefaultVersionFromReadme(f)
		if err != nil {
			return nil, err
		}
		versions[moduleFolder] = version
	}

	return versions, nil
}
