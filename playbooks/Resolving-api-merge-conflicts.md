New Azure API updates sometimes produce merge conflicts as we attempt to reconcile the types
within and between versions. If you encounter one of these errors there are a few different 
approaches to resolve them depending on the source of the conflict:

 - Excluding a version from the default version selection through exclusions in the config.
 - Modifying the merge logic to handle a new case.
 - Add an exception before calling merge to pick one type or the other.
 - Exclude an API version entirely from discovery.


## Example PR

[Fix nightly upgrade due to incompatible refs · pulumi-azure-native/2961](https://github.com/pulumi/pulumi-azure-native/pull/2961) is an example PR for one of these fixes
