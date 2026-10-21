### Default Resources

Azure has the concept of _default_ resource. These are resources, usually child resources of other parent resources that simply exist by default when their parent resource is created. 

When working with these resources, it is important to not attempt to create them via `Create` operations, as they are automatically provisioned by Azure when their parent resource is created. 

In AzureNative terms, we identify these resources as _singletons_ and their lifecycle operations are managed separately from regular resources. Typically a `Delete` is a no-op. 

The provider annotates certain resources as default with some initial state. See `provider/pkg/openapi/defaults/defaultResourcesState.go` 

Sometimes, a resource requires a custom lifecycle implementation because of non-standard API endpoint by Azure. In these cases we handle these as custom resources. For example: https://github.com/pulumi/pulumi-azure-native/pull/4799