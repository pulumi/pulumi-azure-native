# Help, Azure Native isn’t doing the right thing! How can I fix it?

First, try to improve the code to do the right thing automatically. Special cases are expensive to maintain. If that doesn't work out, see the below list of existing override/customization locations.

## Availability of resources

### The wrong API version is used
Override: change the default API version in versions/v2-spec.yaml and versions/v2-lock.yaml.

See: [Specify custom API version on resources #2467](https://github.com/pulumi/pulumi-azure-native/issues/2467), [Change the default API version of AlertProcessingRuleByName to 2021-08-08 #2921](https://github.com/pulumi/pulumi-azure-native/pull/2921) (example), [Change the default version of insights/ActivityLogAlert to supported 2020-10-01 #2941](https://github.com/pulumi/pulumi-azure-native/pull/2941) (example)

### A GET method (invoke) is missing
Override: openapi/additionalInvokes.go.

Note: don’t do this for POST endpoints, invokes need to be side-effect free.

See: Add selected GET invokes #2922

### An auto-generated resource name looks wrong
Override: resources/resources.go `ResourceName()` has a list of exceptions.

### A resource should not be included in Azure Native
Override: gen/exclude.go.

### Two identical resource names clash
A whole API version with all resources and types should not be included
We usually do this when there are naming conflicts (which is an upstream issue) and we only need one of the two versions.

Override: `excludeRegexes` in openapi/discover.go.

### A new Azure API version changed the path and the operation ID of a resource
In that case, it looks like a new resource because we don’t know it’s the same than the existing one.
Override: openapi/paths/normalisePath.go legacyPathMappings


## Behavior of resources

### Removal of a property doesn’t have any effect
A possible cause is that for some resources, Azure treats the omission as "no change" instead of removing the value.

Override: gen/propertyDefaults.go.

See: #2507, [Allow removing network rule sets from storage accounts #2915](https://github.com/pulumi/pulumi-azure-native/pull/2915)

### Azure complains about a missing “properties” but it’s empty
Override: gen/requiredContainers.go (but check first why it’s not auto-detected).

See: [Ensure empty required object properties are in requests #2856](https://github.com/pulumi/pulumi-azure-native/pull/2856)

### When diffing, a property should be compared in a custom way
Override: either provider/diff.go applyAzureSpecificDiff, or add a special case to valueDiff.

### A resource is a singleton and cannot be deleted
Override: add its default state to openapi/defaults/defaultResourcesState.go so it can be reset upon delete.

### A resource should not be created in the configured default location
Override: resources/resources.go AutoLocationDisabled.
