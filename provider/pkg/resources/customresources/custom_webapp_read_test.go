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

func TestReconcileVnetRouteAllEnabled(t *testing.T) {
	for name, testCase := range map[string]struct {
		oldProjection                    map[string]interface{}
		newProjection                    map[string]interface{}
		shouldContainVnetRouteAllEnabled bool
	}{
		"equivalent true, dropped from old projection": {
			oldProjection: map[string]interface{}{
				"vnetRouteAllEnabled": true,
			},
			newProjection: map[string]interface{}{
				"outboundVnetRouting": map[string]interface{}{"applicationTraffic": true},
			},
			shouldContainVnetRouteAllEnabled: false,
		},
		"equivalent false (new property entirely absent), dropped from old projection": {
			oldProjection: map[string]interface{}{
				"vnetRouteAllEnabled": false,
			},
			newProjection:                    map[string]interface{}{},
			shouldContainVnetRouteAllEnabled: false,
		},
		"genuinely different value, kept so the diff still surfaces": {
			oldProjection: map[string]interface{}{
				"vnetRouteAllEnabled": true,
			},
			newProjection: map[string]interface{}{
				"outboundVnetRouting": map[string]interface{}{"applicationTraffic": false},
			},
			shouldContainVnetRouteAllEnabled: true,
		},
		"new schema still has the old property, nothing to reconcile": {
			oldProjection: map[string]interface{}{
				"vnetRouteAllEnabled": true,
			},
			newProjection: map[string]interface{}{
				"vnetRouteAllEnabled": true,
			},
			shouldContainVnetRouteAllEnabled: true,
		},
		"old property absent, no-op": {
			oldProjection:                    map[string]interface{}{},
			newProjection:                    map[string]interface{}{},
			shouldContainVnetRouteAllEnabled: false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			reconcileVnetRouteAllEnabled(testCase.oldProjection, testCase.newProjection)
			_, ok := testCase.oldProjection["vnetRouteAllEnabled"]
			require.Equal(t, testCase.shouldContainVnetRouteAllEnabled, ok)
		})
	}
}
