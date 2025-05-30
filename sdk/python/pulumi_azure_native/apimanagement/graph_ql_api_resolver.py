# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['GraphQLApiResolverArgs', 'GraphQLApiResolver']

@pulumi.input_type
class GraphQLApiResolverArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 service_name: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 path: Optional[pulumi.Input[builtins.str]] = None,
                 resolver_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a GraphQLApiResolver resource.
        :param pulumi.Input[builtins.str] api_id: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] service_name: The name of the API Management service.
        :param pulumi.Input[builtins.str] description: Description of the resolver. May include HTML formatting tags.
        :param pulumi.Input[builtins.str] display_name: Resolver Name.
        :param pulumi.Input[builtins.str] path: Path is type/field being resolved.
        :param pulumi.Input[builtins.str] resolver_id: Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "service_name", service_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if resolver_id is not None:
            pulumi.set(__self__, "resolver_id", resolver_id)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[builtins.str]:
        """
        API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the API Management service.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the resolver. May include HTML formatting tags.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resolver Name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Path is type/field being resolved.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="resolverId")
    def resolver_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
        """
        return pulumi.get(self, "resolver_id")

    @resolver_id.setter
    def resolver_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resolver_id", value)


@pulumi.type_token("azure-native:apimanagement:GraphQLApiResolver")
class GraphQLApiResolver(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 path: Optional[pulumi.Input[builtins.str]] = None,
                 resolver_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        GraphQL API Resolver details.

        Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.

        Other available API versions: 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] api_id: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        :param pulumi.Input[builtins.str] description: Description of the resolver. May include HTML formatting tags.
        :param pulumi.Input[builtins.str] display_name: Resolver Name.
        :param pulumi.Input[builtins.str] path: Path is type/field being resolved.
        :param pulumi.Input[builtins.str] resolver_id: Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] service_name: The name of the API Management service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GraphQLApiResolverArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        GraphQL API Resolver details.

        Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.

        Other available API versions: 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param GraphQLApiResolverArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GraphQLApiResolverArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 path: Optional[pulumi.Input[builtins.str]] = None,
                 resolver_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GraphQLApiResolverArgs.__new__(GraphQLApiResolverArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["path"] = path
            __props__.__dict__["resolver_id"] = resolver_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:apimanagement/v20220801:GraphQLApiResolver"), pulumi.Alias(type_="azure-native:apimanagement/v20220901preview:GraphQLApiResolver"), pulumi.Alias(type_="azure-native:apimanagement/v20230301preview:GraphQLApiResolver"), pulumi.Alias(type_="azure-native:apimanagement/v20230501preview:GraphQLApiResolver"), pulumi.Alias(type_="azure-native:apimanagement/v20230901preview:GraphQLApiResolver"), pulumi.Alias(type_="azure-native:apimanagement/v20240501:GraphQLApiResolver"), pulumi.Alias(type_="azure-native:apimanagement/v20240601preview:GraphQLApiResolver")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GraphQLApiResolver, __self__).__init__(
            'azure-native:apimanagement:GraphQLApiResolver',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GraphQLApiResolver':
        """
        Get an existing GraphQLApiResolver resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GraphQLApiResolverArgs.__new__(GraphQLApiResolverArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["path"] = None
        __props__.__dict__["type"] = None
        return GraphQLApiResolver(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the resolver. May include HTML formatting tags.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Resolver Name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Path is type/field being resolved.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

