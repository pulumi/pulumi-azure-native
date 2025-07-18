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

__all__ = ['VirtualEndpointArgs', 'VirtualEndpoint']

@pulumi.input_type
class VirtualEndpointArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 server_name: pulumi.Input[builtins.str],
                 endpoint_type: Optional[pulumi.Input[Union[builtins.str, 'VirtualEndpointType']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 virtual_endpoint_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a VirtualEndpoint resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] server_name: The name of the server.
        :param pulumi.Input[Union[builtins.str, 'VirtualEndpointType']] endpoint_type: Type of endpoint for the virtual endpoints.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] members: List of flexible servers that one of the virtual endpoints can refer to.
        :param pulumi.Input[builtins.str] virtual_endpoint_name: Base name of the virtual endpoints.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "server_name", server_name)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if virtual_endpoint_name is not None:
            pulumi.set(__self__, "virtual_endpoint_name", virtual_endpoint_name)

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
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the server.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[pulumi.Input[Union[builtins.str, 'VirtualEndpointType']]]:
        """
        Type of endpoint for the virtual endpoints.
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: Optional[pulumi.Input[Union[builtins.str, 'VirtualEndpointType']]]):
        pulumi.set(self, "endpoint_type", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of flexible servers that one of the virtual endpoints can refer to.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="virtualEndpointName")
    def virtual_endpoint_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Base name of the virtual endpoints.
        """
        return pulumi.get(self, "virtual_endpoint_name")

    @virtual_endpoint_name.setter
    def virtual_endpoint_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "virtual_endpoint_name", value)


@pulumi.type_token("azure-native:dbforpostgresql:VirtualEndpoint")
class VirtualEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_type: Optional[pulumi.Input[Union[builtins.str, 'VirtualEndpointType']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                 virtual_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Pair of virtual endpoints for a flexible server.

        Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.

        Other available API versions: 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[builtins.str, 'VirtualEndpointType']] endpoint_type: Type of endpoint for the virtual endpoints.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] members: List of flexible servers that one of the virtual endpoints can refer to.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] server_name: The name of the server.
        :param pulumi.Input[builtins.str] virtual_endpoint_name: Base name of the virtual endpoints.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Pair of virtual endpoints for a flexible server.

        Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.

        Other available API versions: 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param VirtualEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_type: Optional[pulumi.Input[Union[builtins.str, 'VirtualEndpointType']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                 virtual_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualEndpointArgs.__new__(VirtualEndpointArgs)

            __props__.__dict__["endpoint_type"] = endpoint_type
            __props__.__dict__["members"] = members
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__.__dict__["server_name"] = server_name
            __props__.__dict__["virtual_endpoint_name"] = virtual_endpoint_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["virtual_endpoints"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:dbforpostgresql/v20230601preview:VirtualEndpoint"), pulumi.Alias(type_="azure-native:dbforpostgresql/v20231201preview:VirtualEndpoint"), pulumi.Alias(type_="azure-native:dbforpostgresql/v20240301preview:VirtualEndpoint"), pulumi.Alias(type_="azure-native:dbforpostgresql/v20240801:VirtualEndpoint"), pulumi.Alias(type_="azure-native:dbforpostgresql/v20241101preview:VirtualEndpoint"), pulumi.Alias(type_="azure-native:dbforpostgresql/v20250101preview:VirtualEndpoint")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualEndpoint, __self__).__init__(
            'azure-native:dbforpostgresql:VirtualEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualEndpoint':
        """
        Get an existing VirtualEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VirtualEndpointArgs.__new__(VirtualEndpointArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["endpoint_type"] = None
        __props__.__dict__["members"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["virtual_endpoints"] = None
        return VirtualEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Type of endpoint for the virtual endpoints.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        List of flexible servers that one of the virtual endpoints can refer to.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualEndpoints")
    def virtual_endpoints(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        List of virtual endpoints for a flexible server.
        """
        return pulumi.get(self, "virtual_endpoints")

