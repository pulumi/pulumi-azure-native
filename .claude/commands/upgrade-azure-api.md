---
description: Upgrade Azure API versions for a service
---

Guide the user through upgrading Azure API versions for a specific service/namespace. This command should be run from the pulumi-azure-native repository root.

## Prerequisites Check
1. Verify we're in the pulumi-azure-native repository
2. Check that the current branch is clean or prompt user to stash changes
3. Ensure azure-rest-api-specs submodule is initialized

## Workflow

### 1. Gather Information
Ask the user:
- Which Azure service/namespace to upgrade (e.g., "cosmosdb", "compute", "storage")
- What is the target API version? (Optional - if not provided, we'll identify latest available)
- Is there a related GitHub issue number?

### 2. Check Current State
- Read `versions/v3-spec.yaml` to find the current `tracking` version for the service
- Display current version to user
- If user didn't specify target version, proceed to step 3; otherwise skip to step 4

### 3. Identify Available Versions
- Run `make update_submodules` to update azure-rest-api-specs and az-provider-list.json
- Check `azure-rest-api-specs/specification/{service-name}/resource-manager/*/stable/` for available stable versions
- Check `azure-rest-api-specs/specification/{service-name}/resource-manager/*/preview/` for preview versions
- Present options to user and ask which version to use
- **Note**: Prefer stable versions over preview unless specifically needed

### 4. Update Version Configuration
- Create a new branch: `feature/{service}-api-{version}-{issue-number}`
- Edit `versions/v3-spec.yaml`:
  - Update the `tracking` field for the service to the new API version
  - Keep any existing `additions` unless they conflict with the new version
- Commit changes: `git commit -m "Update {Service} to API version {version}"`

### 5. Regenerate Schema (Fast Iteration)
- Run `DEBUG_CODEGEN_NAMESPACES="{service}" make schema` to regenerate schema for just this service
- **Important**: This is faster for testing but generates full schema at the end
- Check for errors in the output

### 6. Verify Schema Changes
- Run `jq '.types | keys[] | select(contains("{Service}"))' bin/schema-default-versions.json | head -20` to preview new types
- Look for new properties that were requested (if upgrading for specific features)
- Ask user if the changes look correct

### 7. Full Schema Regeneration
- Run `make schema` to regenerate the complete schema (this updates both schema files and version metadata)
- Review the changes to `versions/v3.yaml` (auto-generated)

### 8. Generate SDKs
- Run `make` to build all SDKs and the provider binary
- **Warning**: This takes 5-10 minutes and generates ~150 files for affected services
- Monitor for build errors

### 9. Commit Generated Code
- Stage all changes: `git add -A` (includes both modified and new files)
- **Important**: Use `-A` not `-u` to catch new type files
- Review staged files to ensure no unintended changes
- Commit: `git commit -m "Regenerate SDKs with {Service} API version {version}"`

### 10. Create Pull Request
- Push branch: `git push -u origin {branch-name}`
- Create PR with title: "Add support for {Service} {feature-description} (API version {version})"
- Use PR body template:
  ```markdown
  ## Summary
  Updates the {Service} module to API version **{version}**, enabling {feature-description}.

  ## Changes
  - Updated {Service} tracking version from `{old-version}` to `{new-version}` in `versions/v3-spec.yaml`
  - Regenerated version metadata in `versions/v3.yaml`

  ## New Features
  {List new properties/resources available}

  ## Testing
  - ✅ Full schema regeneration completed successfully
  - ✅ All SDKs built without errors (Node.js, Python, .NET, Go, Java)
  - ✅ Provider binary built successfully
  - ✅ Verified new properties exist in generated schema

  ## Related Issues
  Fixes #{issue-number}
  ```

## Common Issues & Troubleshooting

**API Merge Conflicts**
- See `playbooks/Resolving-api-merge-conflicts.md` for type reconciliation issues
- May need to add type overrides in `versions/v3-config.yaml`

**Breaking Changes**
- Review the OpenAPI spec diff in `azure-rest-api-specs` for breaking changes
- May need custom resource handling in `provider/pkg/resources/customresources/`

**Missing Resources**
- Check if resources were intentionally removed in `versions/v3-removed-resources.json`
- See `playbooks/customizing-resources.md` for override procedures

**Build Failures**
- Use `DEBUG_CODEGEN_NAMESPACES="{service}" make schema` to isolate issues
- Check `provider/pkg/gen/` for schema generation errors
- Review `reports/` directory for generation warnings

## Reference
- Playbook: `playbooks/customizing-resources.md`
- Playbook: `playbooks/Resolving-api-merge-conflicts.md`
- Version files: `versions/v3-spec.yaml`, `versions/v3-config.yaml`
- CLAUDE.md: Project-specific conventions and commands
