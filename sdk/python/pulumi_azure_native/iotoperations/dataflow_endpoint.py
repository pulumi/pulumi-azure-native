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

__all__ = ['DataflowEndpointArgs', 'DataflowEndpoint']

@pulumi.input_type
class DataflowEndpointArgs:
    def __init__(__self__, *,
                 extended_location: pulumi.Input['ExtendedLocationArgs'],
                 instance_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 dataflow_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input['DataflowEndpointPropertiesArgs']] = None):
        """
        The set of arguments for constructing a DataflowEndpoint resource.
        :param pulumi.Input['ExtendedLocationArgs'] extended_location: Edge location of the resource.
        :param pulumi.Input[builtins.str] instance_name: Name of instance.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] dataflow_endpoint_name: Name of Instance dataflowEndpoint resource
        :param pulumi.Input['DataflowEndpointPropertiesArgs'] properties: The resource-specific properties for this resource.
        """
        pulumi.set(__self__, "extended_location", extended_location)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if dataflow_endpoint_name is not None:
            pulumi.set(__self__, "dataflow_endpoint_name", dataflow_endpoint_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> pulumi.Input['ExtendedLocationArgs']:
        """
        Edge location of the resource.
        """
        return pulumi.get(self, "extended_location")

    @extended_location.setter
    def extended_location(self, value: pulumi.Input['ExtendedLocationArgs']):
        pulumi.set(self, "extended_location", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_name", value)

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
    @pulumi.getter(name="dataflowEndpointName")
    def dataflow_endpoint_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of Instance dataflowEndpoint resource
        """
        return pulumi.get(self, "dataflow_endpoint_name")

    @dataflow_endpoint_name.setter
    def dataflow_endpoint_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "dataflow_endpoint_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['DataflowEndpointPropertiesArgs']]:
        """
        The resource-specific properties for this resource.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['DataflowEndpointPropertiesArgs']]):
        pulumi.set(self, "properties", value)


@pulumi.type_token("azure-native:iotoperations:DataflowEndpoint")
class DataflowEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataflow_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 extended_location: Optional[pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']]] = None,
                 instance_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['DataflowEndpointPropertiesArgs', 'DataflowEndpointPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Instance dataflowEndpoint resource

        Uses Azure REST API version 2024-11-01.

        Other available API versions: 2024-08-15-preview, 2024-09-15-preview, 2025-04-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotoperations [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] dataflow_endpoint_name: Name of Instance dataflowEndpoint resource
        :param pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']] extended_location: Edge location of the resource.
        :param pulumi.Input[builtins.str] instance_name: Name of instance.
        :param pulumi.Input[Union['DataflowEndpointPropertiesArgs', 'DataflowEndpointPropertiesArgsDict']] properties: The resource-specific properties for this resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataflowEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Instance dataflowEndpoint resource

        Uses Azure REST API version 2024-11-01.

        Other available API versions: 2024-08-15-preview, 2024-09-15-preview, 2025-04-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotoperations [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DataflowEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataflowEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataflow_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 extended_location: Optional[pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']]] = None,
                 instance_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['DataflowEndpointPropertiesArgs', 'DataflowEndpointPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataflowEndpointArgs.__new__(DataflowEndpointArgs)

            __props__.__dict__["dataflow_endpoint_name"] = dataflow_endpoint_name
            if extended_location is None and not opts.urn:
                raise TypeError("Missing required property 'extended_location'")
            __props__.__dict__["extended_location"] = extended_location
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__.__dict__["instance_name"] = instance_name
            __props__.__dict__["properties"] = properties
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:iotoperations/v20240701preview:DataFlowEndpoint"), pulumi.Alias(type_="azure-native:iotoperations/v20240701preview:DataflowEndpoint"), pulumi.Alias(type_="azure-native:iotoperations/v20240815preview:DataflowEndpoint"), pulumi.Alias(type_="azure-native:iotoperations/v20240915preview:DataflowEndpoint"), pulumi.Alias(type_="azure-native:iotoperations/v20241101:DataflowEndpoint"), pulumi.Alias(type_="azure-native:iotoperations/v20250401:DataflowEndpoint"), pulumi.Alias(type_="azure-native:iotoperations/v20250701preview:DataflowEndpoint"), pulumi.Alias(type_="azure-native:iotoperations:DataFlowEndpoint")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DataflowEndpoint, __self__).__init__(
            'azure-native:iotoperations:DataflowEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataflowEndpoint':
        """
        Get an existing DataflowEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataflowEndpointArgs.__new__(DataflowEndpointArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["extended_location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return DataflowEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> pulumi.Output['outputs.ExtendedLocationResponse']:
        """
        Edge location of the resource.
        """
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.DataflowEndpointPropertiesResponse']:
        """
        The resource-specific properties for this resource.
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

