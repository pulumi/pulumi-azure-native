package customresources

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestFilterRedactedPublishingUsername(t *testing.T) {
	for name, testCase := range map[string]struct {
		response                        map[string]any
		shouldContainPublishingUsername bool
	}{
		"realName": {
			response: map[string]any{
				"siteConfig": map[string]any{"publishingUsername": "$1234"},
			},
			shouldContainPublishingUsername: true,
		},
		"redactedName": {
			response: map[string]any{
				"siteConfig": map[string]any{"publishingUsername": "REDACTED"},
			},
			shouldContainPublishingUsername: false,
		},
		"noName": { // don't panic
			response: map[string]any{
				"siteConfig": map[string]any{},
			},
			shouldContainPublishingUsername: false,
		},
		"noSiteConfig": { // don't panic
			response:                        map[string]any{},
			shouldContainPublishingUsername: false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			filterResponse(testCase.response)
			siteConfig, ok := util.GetInnerMap(testCase.response, "siteConfig")
			require.Equal(t, name != "noSiteConfig", ok)
			_, ok = siteConfig["publishingUsername"]
			require.Equal(t, testCase.shouldContainPublishingUsername, ok)
		})
	}
}
