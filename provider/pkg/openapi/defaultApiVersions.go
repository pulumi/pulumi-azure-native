package openapi

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// # How to determine the default API version of a service?
//
// It's in the `readme.md` file of the service in azure-rest-api-specs, in a yaml code block with a `tag` field.
//
// Example:
//
// ```yaml
// openapi-type: arm
// tag: package-2025-03-01
// ```
//
// There will be other fields as well.
//
// The version can be just a year and month, in which case the API version is the only one released in that month.
//
// The version can also be a composite version like `tag: package-composite-v5`. I don't fully understand those yet.
//
// ## Spec version
// Since we're reading the readme.md files from the azure-rest-api-specs repo, we need to know which commit of the
// file to read. The `main` branch of the spec can be updated with new API versions that SDKs are not using yet. In
// case of stable versions, this is _probably_ ok, in case of preview versions it's more questionable. But we cannot
// simply use the latest stable version since some services barely release stable versions.

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
