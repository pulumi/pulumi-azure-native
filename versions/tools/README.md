# About

This folder contains tools for working with Azure modules and resources and their versions, both API versions and provider major versions.

## compare_major_versions.go

This script compares two major provider versions and outputs either a JSON structure of the resource changes, or a Markdown table of all resources with their REST API versions in the two major versions. The table is the one seen in the upgrade guides like [here](https://www.pulumi.com/registry/packages/azure-native/from-v2-to-v3/resource-versions/).

```
go run compare_major_versions.go -olderMajorVersion 2 -out changes
```

```
go run compare_major_versions.go -olderMajorVersion 2 -out table
```

