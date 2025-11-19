# Fix for Issue #4400: API Version Upgrade with pulumi up --refresh

## Problem Summary

When upgrading a resource (e.g., ManagedCluster) from one preview API version to another using `pulumi up --refresh`, the provider incorrectly triggered resource replacement even though no actual changes were made. The workaround `pulumi refresh --run-program && pulumi up` worked correctly.

## Root Cause

The bug occurred in the `Read` method when handling API version changes during refresh:

1. During `pulumi up --refresh`, the resource URN type changes to the new API version (via alias)
2. The `Read` method looks up resource metadata using the NEW URN type
3. When normalizing old state to inputs, it incorrectly uses the NEW API version's schema instead of the OLD schema
4. This causes properties that exist only in the new API version to appear as "additions"
5. The diff calculation shows spurious changes, triggering replacement

## The Fix

### Changes Made

**File: `provider/pkg/provider/provider.go`**

1. **API Version Detection** (lines 1315-1331):
   - Extract `azureApiVersion` from old state
   - Detect when API version has changed
   - Look up old resource metadata if available

2. **New Helper Methods** (lines 178-226):
   - `lookupResourceWithAPIVersion`: Looks up resource metadata for a specific API version
   - `apiVersionToVersionPart`: Converts API version format (e.g., "2024-01-02-preview" → "v20240102preview")

3. **Schema-Aware Normalization** (lines 1398-1438):
   - Use old schema when converting old state to inputs
   - Use new schema when converting new state to inputs
   - Preserve old inputs if old metadata cannot be found (fallback safety)

**File: `provider/pkg/provider/resource_lookup_test.go`** (new file)
- Unit tests for API version format conversion

### Key Implementation Details

#### API Version Tracking
```go
var oldApiVersion string
var oldRes *resources.AzureAPIResource
if azureApiVersion, ok := oldState["azureApiVersion"]; ok && azureApiVersion.IsString() {
    oldApiVersion = azureApiVersion.StringValue()
    if oldApiVersion != res.APIVersion {
        // API version changed - look up old metadata
        oldRes, err = k.lookupResourceWithAPIVersion(urn, oldApiVersion)
    }
}
```

#### Schema-Aware Conversion
```go
// Use old schema for old state, new schema for new state
if oldRes != nil {
    oldInputProjection = k.converter.SdkOutputsToSdkInputs(oldRes.PutParameters, plainOldState)
} else {
    oldInputProjection = k.converter.SdkOutputsToSdkInputs(res.PutParameters, plainOldState)
}
newInputProjection := k.converter.SdkOutputsToSdkInputs(res.PutParameters, outputsWithoutIgnores)
```

#### Fallback Safety
```go
// If old metadata not found, preserve old inputs to avoid spurious diffs
if oldApiVersion != "" && oldApiVersion != res.APIVersion && oldRes == nil {
    logging.V(5).Infof("%s: API version changed but old metadata not found, preserving old inputs", label)
    // Don't apply diff - keep old inputs as-is
} else {
    inputs = applyDiff(inputs, diff)
}
```

## Why This Fixes the Issue

The fix ensures that when converting Azure outputs back to inputs during refresh, we use the schema that matches the data:

- **Old state** (from previous deployment) → converted using **OLD API version schema**
- **New state** (from Azure) → converted using **NEW API version schema**
- **Diff calculation** → now compares apples-to-apples instead of mixing schemas

This prevents properties that exist in the new API version but not the old from appearing as "additions".

## Testing

1. **Unit Tests**: Added tests for API version format conversion
2. **Provider Tests**: All existing unit tests pass (no regressions)
3. **Build**: Provider builds successfully

## Workaround (Until Fixed)

Users affected by this issue can use either:
1. Two-step process: `pulumi refresh --run-program && pulumi up`
2. Combined flag: `pulumi up --refresh --run-program` (as suggested by maintainer)

## Related Issues

- #4400 - Original report
- #4015 - PR that added `azureApiVersion` tracking to state
- #4408 - Similar issue with WebAppBackupConfiguration
- #4231 - WebAppSwiftVirtualNetworkConnection triggers unexpected replace

## Notes

This fix leverages the `azureApiVersion` output property that was added in PR #4015 to enable proper state migration during API version upgrades. The provider's "input-input diffing" architecture is sound - this was a bug in how schemas were selected during the normalization process, not a fundamental design flaw.
