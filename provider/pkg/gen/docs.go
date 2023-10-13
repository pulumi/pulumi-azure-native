package gen

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func getAdditionalDocs(rootDir, resourceTok string) *string {
	withoutPrefix := strings.TrimPrefix(resourceTok, "azure-native:")
	filename := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(withoutPrefix, "-")
	file, err := os.ReadFile(filepath.Join(rootDir, "docs", "resources", filename+".md"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	content := string(file)
	return &content
}
