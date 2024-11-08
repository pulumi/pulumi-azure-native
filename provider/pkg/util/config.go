package util

import "os"

// EnableAzcoreBackend is a feature toggle that returns true if the newer backend using azcore and
// azidentity for REST and authentication should be used. Otherwise, the previous autorest backend
// is used.
//
// Tracked in epic #3576, the new backend was added to upgrade from unmaintained libraries that
// don't receive security and other updates. It uses the latest official Azure packages.
//
// The new backend is enabled by default but can be disabled via this feature toggle.
func EnableAzcoreBackend() bool {
	if enable, ok := os.LookupEnv("PULUMI_ENABLE_AZCORE_BACKEND"); ok {
		return enable == "true"
	}
	return true
}
