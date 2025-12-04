---
description: Upgrade Azure API versions for a service
---

Guide the user through upgrading Azure API versions for a specific service/namespace. This command should be run from the pulumi-azure-native repository root.

## Prerequisites Check
1. Verify we're in the pulumi-azure-native repository
2. Check that the current branch is clean or prompt user to stash changes
3. Ensure azure-rest-api-specs submodule is initialized

## Submodule Management

**Important**: The `azure-rest-api-specs` submodule contains OpenAPI specs for ALL Azure services. Updating it pulls in changes across potentially hundreds of services, which can:
- Introduce unrelated breaking changes
- Cause merge conflicts in version files
- Make PRs harder to review (mixing your changes with unrelated updates)

**Default recommendation**: Do NOT update the submodule unless necessary. Most API versions are already available in the current submodule.

**How to check if a version is available without updating**:
```bash
# Check for stable versions
ls azure-rest-api-specs/specification/{service}/resource-manager/Microsoft.{Service}/stable/

# Check for preview versions
ls azure-rest-api-specs/specification/{service}/resource-manager/Microsoft.{Service}/preview/
```

**When to update the submodule**:
- The target API version doesn't exist in the current submodule
- You're doing a bulk update of multiple services
- You explicitly want to pull in the latest specs

**If you must update**:
```bash
make update_submodules
```
Be prepared for a larger PR that may require additional testing and review.

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
- First, check if the target version already exists in the current submodule (see "Submodule Management" above)
- Check `azure-rest-api-specs/specification/{service-name}/resource-manager/*/stable/` for available stable versions
- Check `azure-rest-api-specs/specification/{service-name}/resource-manager/*/preview/` for preview versions
- Present options to user and ask which version to use
- **Only if needed**: Run `make update_submodules` if the target version isn't available
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
- Run `make generate` to regenerate all SDKs (Node.js, Python, .NET, Go, Java)
- **Warning**: This takes 5-10 minutes and generates ~150 files for affected services
- Monitor for build errors

### 9. Generate Docs Schema
- Run `make generate_docs` to regenerate the docs schema with examples
- This updates `provider/cmd/pulumi-resource-azure-native/schema.json` (committed to repo)
- **Critical**: This step is often forgotten but required for complete PRs

### 10. Build Provider (Optional Verification)
- Run `make provider` to build the provider binary
- Verifies that the generated code compiles correctly

### 11. Commit Generated Code
- Stage specific paths to avoid committing untracked files:
  ```bash
  git add versions/ sdk/ provider/cmd/pulumi-resource-azure-native/schema.json
  ```
- Review staged files: `git diff --cached --stat`
- Commit: `git commit -m "Regenerate SDKs with {Service} API version {version}"`

### 12. Create Pull Request
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

**Stale PR / Rebasing**
- If master has moved forward with other API updates while your PR was in review:
  ```bash
  git fetch origin
  git rebase origin/master
  ```
- After rebasing, re-run ALL generation steps:
  ```bash
  make schema
  make generate
  make generate_docs
  ```
- This ensures your PR doesn't accidentally revert other changes

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
