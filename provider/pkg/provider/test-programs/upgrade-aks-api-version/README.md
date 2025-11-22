# AKS API Version Upgrade Test (v2.90.0 → v3)

## Purpose

Regression test for [issue #4400](https://github.com/pulumi/pulumi-azure-native/issues/4400):
Spurious AKS cluster replacements during refresh after v2 → v3 upgrade.

## The Bug

When upgrading from v2.90.0 (versioned resources) to v3 (unversioned/default resources):

1. Initial deployment with v2.90.0 using versioned API: `containerservice/v20240102preview:ManagedCluster`
2. State stores `azureApiVersion: "2024-01-02-preview"` in outputs
3. Code updated to v3 with unversioned/default API: `containerservice:ManagedCluster` (with alias)
4. v3 provider metadata has different default API version (e.g., `2024-10-02-preview`)
5. During refresh, provider incorrectly used NEW API version from metadata instead of OLD version from state
6. Schema differences between API versions caused false positive property diffs
7. Provider incorrectly reported spurious replacements

### The "Mixed State" Condition

After provider upgrade, state contains:
- Type token: `containerservice:ManagedCluster` (unversioned)
- Output: `azureApiVersion: "2024-01-02-preview"` (OLD API version from v2)
- Provider metadata: Default API version is `2024-10-02-preview` (NEW)

**The bug**: Provider used NEW API version from metadata to read the resource, instead of OLD API version from state outputs.

### v2 vs v3 Resource Types

**v2.90.0** (explicit API versions):
- Uses versioned types: `azure-native:containerservice/v20240102preview:ManagedCluster`
- User explicitly selects API version

**v3** (default API versions):
- Uses unversioned types: `azure-native:containerservice:ManagedCluster`
- Provider selects appropriate default API version
- **No explicit alias needed** - Pulumi automatically aliases compatible type changes within the same provider

## The Fix

Modified `lookupResourceWithAPIVersion()` in `provider.go` to:
- Check for `azureApiVersion` in state outputs during refresh
- Use OLD API version from state (not NEW default from metadata) to read resource
- Use `openapi.ApiToSdkVersion()` for canonical API version conversion
- Validate ARM path stability across API versions
- Avoid spurious replacements by comparing with correct schema version

**Key insight**: The `azureApiVersion` output in state is the source of truth for which API version was used to create the resource.

## Test Strategy

This test uses the `upgradeTest` infrastructure with **two separate programs**:

### Program Structure

```
upgrade-aks-api-version/
├── Pulumi.yaml           # v2 program (versioned type)
├── README.md
└── v3/
    └── Pulumi.yaml       # v3 program (unversioned type + alias)
```

### Test Flow

1. **Recording Phase** (v2.90.0 + root `Pulumi.yaml`):
   - Uses versioned type: `containerservice/v20240102preview:ManagedCluster`
   - Deploys AKS cluster with API version `2024-01-02-preview`
   - State stores `azureApiVersion: "2024-01-02-preview"` in outputs
   - Records GRPC interactions to `grpc.json`

2. **Upgrade Preview Phase** (v3/master + `v3/Pulumi.yaml`):
   - Uses unversioned type: `containerservice:ManagedCluster`
   - Alias connects to old versioned type
   - v3 provider's default API version is different (e.g., `2024-10-02-preview`)
   - During refresh, provider uses `azureApiVersion` from state (OLD version)
   - Validates no spurious replacements occur

3. **Replay Mode** (CI):
   - Uses recorded GRPC interactions (no Azure connection needed)
   - Completes in <1 second

## Running the Test

### First Time (Recording)

To record GRPC interactions with v2.90.0:

```bash
cd provider/pkg/provider

# Ensure you have Azure credentials configured
export ARM_SUBSCRIPTION_ID="..."
export ARM_CLIENT_ID="..."
export ARM_CLIENT_SECRET="..."
export ARM_TENANT_ID="..."

# Run test to record interactions
# This will deploy real Azure resources with v2.90.0 and record GRPC calls
go test -v -run TestUpgradeAksApiVersion_2_90_0 -timeout 30m

# The recording will be saved to:
# testdata/recorded/TestProviderUpgrade/upgrade-aks-api-version/2.90.0/grpc.json
```

**NOTE**: The first run will:
- Deploy an AKS cluster (takes ~3-5 minutes)
- Create a resource group
- Record all provider interactions to grpc.json
- Clean up resources after test completes

### Subsequent Runs (Replay)

Once recorded, the test runs fast using the recording:

```bash
# Run test in replay mode (no Azure connection needed)
go test -v -run TestUpgradeAksApiVersion_2_90_0

# Should complete in < 1 second using recorded GRPC interactions
```

### CI Execution

In CI, the test automatically uses the committed grpc.json recording, requiring no live Azure credentials.

## Test Assertions

The test verifies:
- `assertpreview.HasNoReplacements()` - No spurious replacements during upgrade preview
- `assertpreview.HasNoDeletes()` - No resources deleted during upgrade

## Files

- `Pulumi.yaml` - v2 program using versioned type `containerservice/v20240102preview:ManagedCluster` (for recording with v2.90.0)
- `v3/Pulumi.yaml` - v3 program using unversioned type `containerservice:ManagedCluster` with alias to versioned type (for upgrade preview)
- `README.md` - This file
- `testdata/recorded/TestProviderUpgrade/upgrade-aks-api-version/2.90.0/grpc.json` - Recorded GRPC interactions with v2.90.0 (created on first run)
- `testdata/recorded/TestProviderUpgrade/upgrade-aks-api-version/2.90.0/stack.json` - Stack state snapshot from v2.90.0 recording

## Related

- Issue: https://github.com/pulumi/pulumi-azure-native/issues/4400
- PR: https://github.com/pulumi/pulumi-azure-native/pull/4429
- Manual test logs: `/tmp/preview-refresh-FIXED.log`, `/tmp/up-refresh-FIXED.log`
