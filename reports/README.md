# Reports

## `active.json`

Summary of API versions actually deployed via `az provider list`.

## `allResourcesByVersion.json`

Index of all namespaces, versions and resources.

## `allResourceVersionsByResource.json`

Similar to above but listing the resources within each namespace with a list of spec versions available listed for each resource.

## `curationViolations.json`

Warnings about when the expectations in the v2-config.yaml don't match with reality (`expectAdditions` and `expectTracking`)

## `genWarnings.json`

Pretty much just tracks where we manipulate fields to help with type merging between different versions.

## `pathChanges.json`

Identify all resources with different paths between the current and previous default major version's default resources.

Additional highlights:

- Where we didn't have a resource in the previous major version's default resources.
- Resources listed in the default version that don't exist.

Calculated by `findPathChanges`. Discussed in <https://github.com/pulumi/pulumi-azure-native/pull/2422>

## `pending.json`

A list of versions of each namespace which might be available to upgrade in the default version.
