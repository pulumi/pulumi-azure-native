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
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['PrivateEndpointConnectionArgs', 'PrivateEndpointConnection']

@pulumi.input_type
class PrivateEndpointConnectionArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 search_service_name: pulumi.Input[builtins.str],
                 private_endpoint_connection_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input['PrivateEndpointConnectionPropertiesArgs']] = None):
        """
        The set of arguments for constructing a PrivateEndpointConnection resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] search_service_name: The name of the search service associated with the specified resource group.
        :param pulumi.Input[builtins.str] private_endpoint_connection_name: The name of the private endpoint connection to the search service with the specified resource group.
        :param pulumi.Input['PrivateEndpointConnectionPropertiesArgs'] properties: Describes the properties of an existing private endpoint connection to the search service.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "search_service_name", search_service_name)
        if private_endpoint_connection_name is not None:
            pulumi.set(__self__, "private_endpoint_connection_name", private_endpoint_connection_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="searchServiceName")
    def search_service_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the search service associated with the specified resource group.
        """
        return pulumi.get(self, "search_service_name")

    @search_service_name.setter
    def search_service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "search_service_name", value)

    @property
    @pulumi.getter(name="privateEndpointConnectionName")
    def private_endpoint_connection_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the private endpoint connection to the search service with the specified resource group.
        """
        return pulumi.get(self, "private_endpoint_connection_name")

    @private_endpoint_connection_name.setter
    def private_endpoint_connection_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "private_endpoint_connection_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['PrivateEndpointConnectionPropertiesArgs']]:
        """
        Describes the properties of an existing private endpoint connection to the search service.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['PrivateEndpointConnectionPropertiesArgs']]):
        pulumi.set(self, "properties", value)


@pulumi.type_token("azure-native:search:PrivateEndpointConnection")
class PrivateEndpointConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 private_endpoint_connection_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['PrivateEndpointConnectionPropertiesArgs', 'PrivateEndpointConnectionPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 search_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Describes an existing private endpoint connection to the search service.

        Uses Azure REST API version 2023-11-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2022-09-01, 2024-03-01-preview, 2024-06-01-preview, 2025-02-01-preview, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native search [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] private_endpoint_connection_name: The name of the private endpoint connection to the search service with the specified resource group.
        :param pulumi.Input[Union['PrivateEndpointConnectionPropertiesArgs', 'PrivateEndpointConnectionPropertiesArgsDict']] properties: Describes the properties of an existing private endpoint connection to the search service.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] search_service_name: The name of the search service associated with the specified resource group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateEndpointConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Describes an existing private endpoint connection to the search service.

        Uses Azure REST API version 2023-11-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2022-09-01, 2024-03-01-preview, 2024-06-01-preview, 2025-02-01-preview, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native search [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param PrivateEndpointConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateEndpointConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 private_endpoint_connection_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['PrivateEndpointConnectionPropertiesArgs', 'PrivateEndpointConnectionPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 search_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateEndpointConnectionArgs.__new__(PrivateEndpointConnectionArgs)

            __props__.__dict__["private_endpoint_connection_name"] = private_endpoint_connection_name
            __props__.__dict__["properties"] = properties
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if search_service_name is None and not opts.urn:
                raise TypeError("Missing required property 'search_service_name'")
            __props__.__dict__["search_service_name"] = search_service_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:search/v20191001preview:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20200313:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20200801:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20200801preview:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20210401preview:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20220901:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20231101:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20240301preview:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20240601preview:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20250201preview:PrivateEndpointConnection"), pulumi.Alias(type_="azure-native:search/v20250501:PrivateEndpointConnection")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PrivateEndpointConnection, __self__).__init__(
            'azure-native:search:PrivateEndpointConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateEndpointConnection':
        """
        Get an existing PrivateEndpointConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PrivateEndpointConnectionArgs.__new__(PrivateEndpointConnectionArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["type"] = None
        return PrivateEndpointConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.PrivateEndpointConnectionPropertiesResponse']:
        """
        Describes the properties of an existing private endpoint connection to the search service.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

