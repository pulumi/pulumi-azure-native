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

__all__ = ['DedicatedHostGroupArgs', 'DedicatedHostGroup']

@pulumi.input_type
class DedicatedHostGroupArgs:
    def __init__(__self__, *,
                 platform_fault_domain_count: pulumi.Input[builtins.int],
                 resource_group_name: pulumi.Input[builtins.str],
                 additional_capabilities: Optional[pulumi.Input['DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs']] = None,
                 host_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 support_automatic_placement: Optional[pulumi.Input[builtins.bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a DedicatedHostGroup resource.
        :param pulumi.Input[builtins.int] platform_fault_domain_count: Number of fault domains that the host group can span.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input['DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs'] additional_capabilities: Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
        :param pulumi.Input[builtins.str] host_group_name: The name of the dedicated host group.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.bool] support_automatic_placement: Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] zones: The availability zones.
        """
        pulumi.set(__self__, "platform_fault_domain_count", platform_fault_domain_count)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if additional_capabilities is not None:
            pulumi.set(__self__, "additional_capabilities", additional_capabilities)
        if host_group_name is not None:
            pulumi.set(__self__, "host_group_name", host_group_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if support_automatic_placement is not None:
            pulumi.set(__self__, "support_automatic_placement", support_automatic_placement)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zones is not None:
            pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> pulumi.Input[builtins.int]:
        """
        Number of fault domains that the host group can span.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @platform_fault_domain_count.setter
    def platform_fault_domain_count(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "platform_fault_domain_count", value)

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
    @pulumi.getter(name="additionalCapabilities")
    def additional_capabilities(self) -> Optional[pulumi.Input['DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs']]:
        """
        Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
        """
        return pulumi.get(self, "additional_capabilities")

    @additional_capabilities.setter
    def additional_capabilities(self, value: Optional[pulumi.Input['DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs']]):
        pulumi.set(self, "additional_capabilities", value)

    @property
    @pulumi.getter(name="hostGroupName")
    def host_group_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the dedicated host group.
        """
        return pulumi.get(self, "host_group_name")

    @host_group_name.setter
    def host_group_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "host_group_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="supportAutomaticPlacement")
    def support_automatic_placement(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
        """
        return pulumi.get(self, "support_automatic_placement")

    @support_automatic_placement.setter
    def support_automatic_placement(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "support_automatic_placement", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The availability zones.
        """
        return pulumi.get(self, "zones")

    @zones.setter
    def zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "zones", value)


@pulumi.type_token("azure-native:compute:DedicatedHostGroup")
class DedicatedHostGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_capabilities: Optional[pulumi.Input[Union['DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs', 'DedicatedHostGroupPropertiesAdditionalCapabilitiesArgsDict']]] = None,
                 host_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[builtins.int]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 support_automatic_placement: Optional[pulumi.Input[builtins.bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Specifies information about the dedicated host group that the dedicated hosts should be assigned to. Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.

        Uses Azure REST API version 2024-11-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.

        Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs', 'DedicatedHostGroupPropertiesAdditionalCapabilitiesArgsDict']] additional_capabilities: Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
        :param pulumi.Input[builtins.str] host_group_name: The name of the dedicated host group.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.int] platform_fault_domain_count: Number of fault domains that the host group can span.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.bool] support_automatic_placement: Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] zones: The availability zones.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DedicatedHostGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Specifies information about the dedicated host group that the dedicated hosts should be assigned to. Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.

        Uses Azure REST API version 2024-11-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.

        Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DedicatedHostGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedHostGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_capabilities: Optional[pulumi.Input[Union['DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs', 'DedicatedHostGroupPropertiesAdditionalCapabilitiesArgsDict']]] = None,
                 host_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[builtins.int]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 support_automatic_placement: Optional[pulumi.Input[builtins.bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DedicatedHostGroupArgs.__new__(DedicatedHostGroupArgs)

            __props__.__dict__["additional_capabilities"] = additional_capabilities
            __props__.__dict__["host_group_name"] = host_group_name
            __props__.__dict__["location"] = location
            if platform_fault_domain_count is None and not opts.urn:
                raise TypeError("Missing required property 'platform_fault_domain_count'")
            __props__.__dict__["platform_fault_domain_count"] = platform_fault_domain_count
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["support_automatic_placement"] = support_automatic_placement
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zones"] = zones
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["hosts"] = None
            __props__.__dict__["instance_view"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:compute/v20190301:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20190701:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20191201:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20200601:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20201201:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20210301:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20210401:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20210701:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20211101:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20220301:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20220801:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20221101:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20230301:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20230701:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20230901:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20240301:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20240701:DedicatedHostGroup"), pulumi.Alias(type_="azure-native:compute/v20241101:DedicatedHostGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DedicatedHostGroup, __self__).__init__(
            'azure-native:compute:DedicatedHostGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DedicatedHostGroup':
        """
        Get an existing DedicatedHostGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DedicatedHostGroupArgs.__new__(DedicatedHostGroupArgs)

        __props__.__dict__["additional_capabilities"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["hosts"] = None
        __props__.__dict__["instance_view"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["platform_fault_domain_count"] = None
        __props__.__dict__["support_automatic_placement"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["zones"] = None
        return DedicatedHostGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalCapabilities")
    def additional_capabilities(self) -> pulumi.Output[Optional['outputs.DedicatedHostGroupPropertiesAdditionalCapabilitiesResponse']]:
        """
        Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
        """
        return pulumi.get(self, "additional_capabilities")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def hosts(self) -> pulumi.Output[Sequence['outputs.SubResourceReadOnlyResponse']]:
        """
        A list of references to all dedicated hosts in the dedicated host group.
        """
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter(name="instanceView")
    def instance_view(self) -> pulumi.Output['outputs.DedicatedHostGroupInstanceViewResponse']:
        """
        The dedicated host group instance view, which has the list of instance view of the dedicated hosts under the dedicated host group.
        """
        return pulumi.get(self, "instance_view")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> pulumi.Output[builtins.int]:
        """
        Number of fault domains that the host group can span.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @property
    @pulumi.getter(name="supportAutomaticPlacement")
    def support_automatic_placement(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
        """
        return pulumi.get(self, "support_automatic_placement")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The availability zones.
        """
        return pulumi.get(self, "zones")

