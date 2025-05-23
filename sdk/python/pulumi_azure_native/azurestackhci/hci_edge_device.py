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
from ._inputs import *

__all__ = ['HciEdgeDeviceArgs', 'HciEdgeDevice']

@pulumi.input_type
class HciEdgeDeviceArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[builtins.str],
                 resource_uri: pulumi.Input[builtins.str],
                 edge_device_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input['HciEdgeDevicePropertiesArgs']] = None):
        """
        The set of arguments for constructing a HciEdgeDevice resource.
        :param pulumi.Input[builtins.str] kind: Edge device kind.
               Expected value is 'HCI'.
        :param pulumi.Input[builtins.str] resource_uri: The fully qualified Azure Resource manager identifier of the resource.
        :param pulumi.Input[builtins.str] edge_device_name: Name of Device
        :param pulumi.Input['HciEdgeDevicePropertiesArgs'] properties: properties for Arc-enabled edge device with HCI OS.
        """
        pulumi.set(__self__, "kind", 'HCI')
        pulumi.set(__self__, "resource_uri", resource_uri)
        if edge_device_name is not None:
            pulumi.set(__self__, "edge_device_name", edge_device_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[builtins.str]:
        """
        Edge device kind.
        Expected value is 'HCI'.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="resourceUri")
    def resource_uri(self) -> pulumi.Input[builtins.str]:
        """
        The fully qualified Azure Resource manager identifier of the resource.
        """
        return pulumi.get(self, "resource_uri")

    @resource_uri.setter
    def resource_uri(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_uri", value)

    @property
    @pulumi.getter(name="edgeDeviceName")
    def edge_device_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of Device
        """
        return pulumi.get(self, "edge_device_name")

    @edge_device_name.setter
    def edge_device_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "edge_device_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['HciEdgeDevicePropertiesArgs']]:
        """
        properties for Arc-enabled edge device with HCI OS.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['HciEdgeDevicePropertiesArgs']]):
        pulumi.set(self, "properties", value)


@pulumi.type_token("azure-native:azurestackhci:HciEdgeDevice")
class HciEdgeDevice(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 edge_device_name: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['HciEdgeDevicePropertiesArgs', 'HciEdgeDevicePropertiesArgsDict']]] = None,
                 resource_uri: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Arc-enabled edge device with HCI OS.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-08-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] edge_device_name: Name of Device
        :param pulumi.Input[builtins.str] kind: Edge device kind.
               Expected value is 'HCI'.
        :param pulumi.Input[Union['HciEdgeDevicePropertiesArgs', 'HciEdgeDevicePropertiesArgsDict']] properties: properties for Arc-enabled edge device with HCI OS.
        :param pulumi.Input[builtins.str] resource_uri: The fully qualified Azure Resource manager identifier of the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HciEdgeDeviceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Arc-enabled edge device with HCI OS.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-08-01-preview.

        :param str resource_name: The name of the resource.
        :param HciEdgeDeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HciEdgeDeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 edge_device_name: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['HciEdgeDevicePropertiesArgs', 'HciEdgeDevicePropertiesArgsDict']]] = None,
                 resource_uri: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HciEdgeDeviceArgs.__new__(HciEdgeDeviceArgs)

            __props__.__dict__["edge_device_name"] = edge_device_name
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = 'HCI'
            __props__.__dict__["properties"] = properties
            if resource_uri is None and not opts.urn:
                raise TypeError("Missing required property 'resource_uri'")
            __props__.__dict__["resource_uri"] = resource_uri
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:azurestackhci/v20230801preview:EdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20230801preview:HciEdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20231101preview:EdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20231101preview:HciEdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20240101:EdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20240101:HciEdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20240215preview:HciEdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20240401:HciEdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20240901preview:HciEdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20241201preview:HciEdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci/v20250201preview:HciEdgeDevice"), pulumi.Alias(type_="azure-native:azurestackhci:EdgeDevice")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(HciEdgeDevice, __self__).__init__(
            'azure-native:azurestackhci:HciEdgeDevice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HciEdgeDevice':
        """
        Get an existing HciEdgeDevice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = HciEdgeDeviceArgs.__new__(HciEdgeDeviceArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return HciEdgeDevice(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        Edge device kind.
        Expected value is 'HCI'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.HciEdgeDevicePropertiesResponse']:
        """
        properties for Arc-enabled edge device with HCI OS.
        """
        return pulumi.get(self, "properties")

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

