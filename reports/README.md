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

## `flattenedPropertyConflicts.json`

The Azure spec has an annotation [x-ms-client-flatten](https://github.com/Azure/autorest/blob/main/docs/extensions/readme.md#x-ms-client-flatten) which instructs the schema generator to flatten the property into the parent object. However, there are cases where this leads to overwriting a property, creating incorrect schema and SDKs, when inner and outer property have the same name.

## `allEndpoints.json`

This file contains a list of all endpoints defined in the Azure spec, grouped by resource provider and then by URL path. It can be useful for exploration and analytics purposes, since it contains endpoints that are not included in our schema, but in a compact format.

This, most likely very unelegant, `jq` query filters the endpoints to only those with inner POST endpoints, and may serve as a template for other queries:

```jq
jq '[ keys[] as $rp | .[$rp] | keys[] as $path | .[$path][] ] | map(select(.HttpVerbs != null) | select(.HttpVerbs | contains(["POST"])))' allEndpoints.json
```

## `oldApiVersions.json`

This file contains a list of all API versions that are older than the oldest previous default version for each resource provider. These API versions are good candidates for removal. Versions that are already explicitly removed via `vN-removed.json` are not included.
