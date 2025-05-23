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

__all__ = ['DeviceArgs', 'Device']

@pulumi.input_type
class DeviceArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 data_residency: Optional[pulumi.Input['DataResidencyArgs']] = None,
                 device_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input['ResourceIdentityArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input['SkuArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Device resource.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input['DataResidencyArgs'] data_residency: The details of data-residency related properties for this resource
        :param pulumi.Input[builtins.str] device_name: The device name.
        :param pulumi.Input['ResourceIdentityArgs'] identity: Msi identity of the resource
        :param pulumi.Input[builtins.str] location: The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
        :param pulumi.Input['SkuArgs'] sku: The SKU type.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if data_residency is not None:
            pulumi.set(__self__, "data_residency", data_residency)
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The resource group name.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="dataResidency")
    def data_residency(self) -> Optional[pulumi.Input['DataResidencyArgs']]:
        """
        The details of data-residency related properties for this resource
        """
        return pulumi.get(self, "data_residency")

    @data_residency.setter
    def data_residency(self, value: Optional[pulumi.Input['DataResidencyArgs']]):
        pulumi.set(self, "data_residency", value)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The device name.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ResourceIdentityArgs']]:
        """
        Msi identity of the resource
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ResourceIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def sku(self) -> Optional[pulumi.Input['SkuArgs']]:
        """
        The SKU type.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input['SkuArgs']]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("azure-native:databoxedge:Device")
class Device(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_residency: Optional[pulumi.Input[Union['DataResidencyArgs', 'DataResidencyArgsDict']]] = None,
                 device_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input[Union['ResourceIdentityArgs', 'ResourceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        The Data Box Edge/Gateway device.

        Uses Azure REST API version 2023-07-01. In version 2.x of the Azure Native provider, it used API version 2022-03-01.

        Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['DataResidencyArgs', 'DataResidencyArgsDict']] data_residency: The details of data-residency related properties for this resource
        :param pulumi.Input[builtins.str] device_name: The device name.
        :param pulumi.Input[Union['ResourceIdentityArgs', 'ResourceIdentityArgsDict']] identity: Msi identity of the resource
        :param pulumi.Input[builtins.str] location: The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input[Union['SkuArgs', 'SkuArgsDict']] sku: The SKU type.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The Data Box Edge/Gateway device.

        Uses Azure REST API version 2023-07-01. In version 2.x of the Azure Native provider, it used API version 2022-03-01.

        Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_residency: Optional[pulumi.Input[Union['DataResidencyArgs', 'DataResidencyArgsDict']]] = None,
                 device_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input[Union['ResourceIdentityArgs', 'ResourceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceArgs.__new__(DeviceArgs)

            __props__.__dict__["data_residency"] = data_residency
            __props__.__dict__["device_name"] = device_name
            __props__.__dict__["identity"] = identity
            __props__.__dict__["location"] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["configured_role_types"] = None
            __props__.__dict__["culture"] = None
            __props__.__dict__["data_box_edge_device_status"] = None
            __props__.__dict__["description"] = None
            __props__.__dict__["device_hcs_version"] = None
            __props__.__dict__["device_local_capacity"] = None
            __props__.__dict__["device_model"] = None
            __props__.__dict__["device_software_version"] = None
            __props__.__dict__["device_type"] = None
            __props__.__dict__["edge_profile"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["friendly_name"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["kubernetes_workload_profile"] = None
            __props__.__dict__["model_description"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["node_count"] = None
            __props__.__dict__["resource_move_details"] = None
            __props__.__dict__["serial_number"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["time_zone"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:databoxedge/v20190301:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20190701:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20190801:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20200501preview:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20200901:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20200901preview:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20201201:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20210201:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20210201preview:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20210601:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20210601preview:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20220301:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20220401preview:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20221201preview:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20230101preview:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20230701:Device"), pulumi.Alias(type_="azure-native:databoxedge/v20231201:Device")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Device, __self__).__init__(
            'azure-native:databoxedge:Device',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Device':
        """
        Get an existing Device resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeviceArgs.__new__(DeviceArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["configured_role_types"] = None
        __props__.__dict__["culture"] = None
        __props__.__dict__["data_box_edge_device_status"] = None
        __props__.__dict__["data_residency"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["device_hcs_version"] = None
        __props__.__dict__["device_local_capacity"] = None
        __props__.__dict__["device_model"] = None
        __props__.__dict__["device_software_version"] = None
        __props__.__dict__["device_type"] = None
        __props__.__dict__["edge_profile"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["friendly_name"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["kubernetes_workload_profile"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["model_description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["node_count"] = None
        __props__.__dict__["resource_move_details"] = None
        __props__.__dict__["serial_number"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["time_zone"] = None
        __props__.__dict__["type"] = None
        return Device(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="configuredRoleTypes")
    def configured_role_types(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Type of compute roles configured.
        """
        return pulumi.get(self, "configured_role_types")

    @property
    @pulumi.getter
    def culture(self) -> pulumi.Output[builtins.str]:
        """
        The Data Box Edge/Gateway device culture.
        """
        return pulumi.get(self, "culture")

    @property
    @pulumi.getter(name="dataBoxEdgeDeviceStatus")
    def data_box_edge_device_status(self) -> pulumi.Output[builtins.str]:
        """
        The status of the Data Box Edge/Gateway device.
        """
        return pulumi.get(self, "data_box_edge_device_status")

    @property
    @pulumi.getter(name="dataResidency")
    def data_residency(self) -> pulumi.Output[Optional['outputs.DataResidencyResponse']]:
        """
        The details of data-residency related properties for this resource
        """
        return pulumi.get(self, "data_residency")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[builtins.str]:
        """
        The Description of the Data Box Edge/Gateway device.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceHcsVersion")
    def device_hcs_version(self) -> pulumi.Output[builtins.str]:
        """
        The device software version number of the device (eg: 1.2.18105.6).
        """
        return pulumi.get(self, "device_hcs_version")

    @property
    @pulumi.getter(name="deviceLocalCapacity")
    def device_local_capacity(self) -> pulumi.Output[builtins.float]:
        """
        The Data Box Edge/Gateway device local capacity in MB.
        """
        return pulumi.get(self, "device_local_capacity")

    @property
    @pulumi.getter(name="deviceModel")
    def device_model(self) -> pulumi.Output[builtins.str]:
        """
        The Data Box Edge/Gateway device model.
        """
        return pulumi.get(self, "device_model")

    @property
    @pulumi.getter(name="deviceSoftwareVersion")
    def device_software_version(self) -> pulumi.Output[builtins.str]:
        """
        The Data Box Edge/Gateway device software version.
        """
        return pulumi.get(self, "device_software_version")

    @property
    @pulumi.getter(name="deviceType")
    def device_type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the Data Box Edge/Gateway device.
        """
        return pulumi.get(self, "device_type")

    @property
    @pulumi.getter(name="edgeProfile")
    def edge_profile(self) -> pulumi.Output['outputs.EdgeProfileResponse']:
        """
        The details of Edge Profile for this resource
        """
        return pulumi.get(self, "edge_profile")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The etag for the devices.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[builtins.str]:
        """
        The Data Box Edge/Gateway device name.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ResourceIdentityResponse']]:
        """
        Msi identity of the resource
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        The kind of the device.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="kubernetesWorkloadProfile")
    def kubernetes_workload_profile(self) -> pulumi.Output[builtins.str]:
        """
        Kubernetes Workload Profile
        """
        return pulumi.get(self, "kubernetes_workload_profile")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="modelDescription")
    def model_description(self) -> pulumi.Output[builtins.str]:
        """
        The description of the Data Box Edge/Gateway device model.
        """
        return pulumi.get(self, "model_description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The object name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[builtins.int]:
        """
        The number of nodes in the cluster.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter(name="resourceMoveDetails")
    def resource_move_details(self) -> pulumi.Output['outputs.ResourceMoveDetailsResponse']:
        """
        The details of the move operation on this resource.
        """
        return pulumi.get(self, "resource_move_details")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output[builtins.str]:
        """
        The Serial Number of Data Box Edge/Gateway device.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        The SKU type.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        DataBoxEdge Resource
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> pulumi.Output[builtins.str]:
        """
        The Data Box Edge/Gateway device timezone.
        """
        return pulumi.get(self, "time_zone")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")

