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

__all__ = ['VirtualMachineTemplateArgs', 'VirtualMachineTemplate']

@pulumi.input_type
class VirtualMachineTemplateArgs:
    def __init__(__self__, *,
                 extended_location: pulumi.Input['ExtendedLocationArgs'],
                 resource_group_name: pulumi.Input[builtins.str],
                 inventory_item_id: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 uuid: Optional[pulumi.Input[builtins.str]] = None,
                 virtual_machine_template_name: Optional[pulumi.Input[builtins.str]] = None,
                 vmm_server_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a VirtualMachineTemplate resource.
        :param pulumi.Input['ExtendedLocationArgs'] extended_location: The extended location.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        :param pulumi.Input[builtins.str] inventory_item_id: Gets or sets the inventory Item ID for the resource.
        :param pulumi.Input[builtins.str] location: Gets or sets the location.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags
        :param pulumi.Input[builtins.str] uuid: Unique ID of the virtual machine template.
        :param pulumi.Input[builtins.str] virtual_machine_template_name: Name of the VirtualMachineTemplate.
        :param pulumi.Input[builtins.str] vmm_server_id: ARM Id of the vmmServer resource in which this resource resides.
        """
        pulumi.set(__self__, "extended_location", extended_location)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if inventory_item_id is not None:
            pulumi.set(__self__, "inventory_item_id", inventory_item_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if virtual_machine_template_name is not None:
            pulumi.set(__self__, "virtual_machine_template_name", virtual_machine_template_name)
        if vmm_server_id is not None:
            pulumi.set(__self__, "vmm_server_id", vmm_server_id)

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> pulumi.Input['ExtendedLocationArgs']:
        """
        The extended location.
        """
        return pulumi.get(self, "extended_location")

    @extended_location.setter
    def extended_location(self, value: pulumi.Input['ExtendedLocationArgs']):
        pulumi.set(self, "extended_location", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="inventoryItemId")
    def inventory_item_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Gets or sets the inventory Item ID for the resource.
        """
        return pulumi.get(self, "inventory_item_id")

    @inventory_item_id.setter
    def inventory_item_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "inventory_item_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Gets or sets the location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Unique ID of the virtual machine template.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "uuid", value)

    @property
    @pulumi.getter(name="virtualMachineTemplateName")
    def virtual_machine_template_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the VirtualMachineTemplate.
        """
        return pulumi.get(self, "virtual_machine_template_name")

    @virtual_machine_template_name.setter
    def virtual_machine_template_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "virtual_machine_template_name", value)

    @property
    @pulumi.getter(name="vmmServerId")
    def vmm_server_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ARM Id of the vmmServer resource in which this resource resides.
        """
        return pulumi.get(self, "vmm_server_id")

    @vmm_server_id.setter
    def vmm_server_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vmm_server_id", value)


@pulumi.type_token("azure-native:scvmm:VirtualMachineTemplate")
class VirtualMachineTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extended_location: Optional[pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']]] = None,
                 inventory_item_id: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 uuid: Optional[pulumi.Input[builtins.str]] = None,
                 virtual_machine_template_name: Optional[pulumi.Input[builtins.str]] = None,
                 vmm_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The VirtualMachineTemplates resource definition.

        Uses Azure REST API version 2023-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-05-21-preview.

        Other available API versions: 2022-05-21-preview, 2023-10-07, 2024-06-01, 2025-03-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']] extended_location: The extended location.
        :param pulumi.Input[builtins.str] inventory_item_id: Gets or sets the inventory Item ID for the resource.
        :param pulumi.Input[builtins.str] location: Gets or sets the location.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags
        :param pulumi.Input[builtins.str] uuid: Unique ID of the virtual machine template.
        :param pulumi.Input[builtins.str] virtual_machine_template_name: Name of the VirtualMachineTemplate.
        :param pulumi.Input[builtins.str] vmm_server_id: ARM Id of the vmmServer resource in which this resource resides.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualMachineTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The VirtualMachineTemplates resource definition.

        Uses Azure REST API version 2023-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-05-21-preview.

        Other available API versions: 2022-05-21-preview, 2023-10-07, 2024-06-01, 2025-03-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param VirtualMachineTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualMachineTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extended_location: Optional[pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']]] = None,
                 inventory_item_id: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 uuid: Optional[pulumi.Input[builtins.str]] = None,
                 virtual_machine_template_name: Optional[pulumi.Input[builtins.str]] = None,
                 vmm_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualMachineTemplateArgs.__new__(VirtualMachineTemplateArgs)

            if extended_location is None and not opts.urn:
                raise TypeError("Missing required property 'extended_location'")
            __props__.__dict__["extended_location"] = extended_location
            __props__.__dict__["inventory_item_id"] = inventory_item_id
            __props__.__dict__["location"] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["virtual_machine_template_name"] = virtual_machine_template_name
            __props__.__dict__["vmm_server_id"] = vmm_server_id
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["computer_name"] = None
            __props__.__dict__["cpu_count"] = None
            __props__.__dict__["disks"] = None
            __props__.__dict__["dynamic_memory_enabled"] = None
            __props__.__dict__["dynamic_memory_max_mb"] = None
            __props__.__dict__["dynamic_memory_min_mb"] = None
            __props__.__dict__["generation"] = None
            __props__.__dict__["is_customizable"] = None
            __props__.__dict__["is_highly_available"] = None
            __props__.__dict__["limit_cpu_for_migration"] = None
            __props__.__dict__["memory_mb"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["network_interfaces"] = None
            __props__.__dict__["os_name"] = None
            __props__.__dict__["os_type"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:scvmm/v20200605preview:VirtualMachineTemplate"), pulumi.Alias(type_="azure-native:scvmm/v20220521preview:VirtualMachineTemplate"), pulumi.Alias(type_="azure-native:scvmm/v20230401preview:VirtualMachineTemplate"), pulumi.Alias(type_="azure-native:scvmm/v20231007:VirtualMachineTemplate"), pulumi.Alias(type_="azure-native:scvmm/v20240601:VirtualMachineTemplate"), pulumi.Alias(type_="azure-native:scvmm/v20250313:VirtualMachineTemplate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualMachineTemplate, __self__).__init__(
            'azure-native:scvmm:VirtualMachineTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualMachineTemplate':
        """
        Get an existing VirtualMachineTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VirtualMachineTemplateArgs.__new__(VirtualMachineTemplateArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["computer_name"] = None
        __props__.__dict__["cpu_count"] = None
        __props__.__dict__["disks"] = None
        __props__.__dict__["dynamic_memory_enabled"] = None
        __props__.__dict__["dynamic_memory_max_mb"] = None
        __props__.__dict__["dynamic_memory_min_mb"] = None
        __props__.__dict__["extended_location"] = None
        __props__.__dict__["generation"] = None
        __props__.__dict__["inventory_item_id"] = None
        __props__.__dict__["is_customizable"] = None
        __props__.__dict__["is_highly_available"] = None
        __props__.__dict__["limit_cpu_for_migration"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["memory_mb"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_interfaces"] = None
        __props__.__dict__["os_name"] = None
        __props__.__dict__["os_type"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["uuid"] = None
        __props__.__dict__["vmm_server_id"] = None
        return VirtualMachineTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="computerName")
    def computer_name(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets computer name.
        """
        return pulumi.get(self, "computer_name")

    @property
    @pulumi.getter(name="cpuCount")
    def cpu_count(self) -> pulumi.Output[builtins.int]:
        """
        Gets or sets the desired number of vCPUs for the vm.
        """
        return pulumi.get(self, "cpu_count")

    @property
    @pulumi.getter
    def disks(self) -> pulumi.Output[Sequence['outputs.VirtualDiskResponse']]:
        """
        Gets or sets the disks of the template.
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter(name="dynamicMemoryEnabled")
    def dynamic_memory_enabled(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets a value indicating whether to enable dynamic memory or not.
        """
        return pulumi.get(self, "dynamic_memory_enabled")

    @property
    @pulumi.getter(name="dynamicMemoryMaxMB")
    def dynamic_memory_max_mb(self) -> pulumi.Output[builtins.int]:
        """
        Gets or sets the max dynamic memory for the vm.
        """
        return pulumi.get(self, "dynamic_memory_max_mb")

    @property
    @pulumi.getter(name="dynamicMemoryMinMB")
    def dynamic_memory_min_mb(self) -> pulumi.Output[builtins.int]:
        """
        Gets or sets the min dynamic memory for the vm.
        """
        return pulumi.get(self, "dynamic_memory_min_mb")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> pulumi.Output['outputs.ExtendedLocationResponse']:
        """
        The extended location.
        """
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter
    def generation(self) -> pulumi.Output[builtins.int]:
        """
        Gets or sets the generation for the vm.
        """
        return pulumi.get(self, "generation")

    @property
    @pulumi.getter(name="inventoryItemId")
    def inventory_item_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets the inventory Item ID for the resource.
        """
        return pulumi.get(self, "inventory_item_id")

    @property
    @pulumi.getter(name="isCustomizable")
    def is_customizable(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets a value indicating whether the vm template is customizable or not.
        """
        return pulumi.get(self, "is_customizable")

    @property
    @pulumi.getter(name="isHighlyAvailable")
    def is_highly_available(self) -> pulumi.Output[builtins.str]:
        """
        Gets highly available property.
        """
        return pulumi.get(self, "is_highly_available")

    @property
    @pulumi.getter(name="limitCpuForMigration")
    def limit_cpu_for_migration(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets a value indicating whether to enable processor compatibility mode for live migration of VMs.
        """
        return pulumi.get(self, "limit_cpu_for_migration")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="memoryMB")
    def memory_mb(self) -> pulumi.Output[builtins.int]:
        """
        MemoryMB is the desired size of a virtual machine's memory, in MB.
        """
        return pulumi.get(self, "memory_mb")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> pulumi.Output[Sequence['outputs.NetworkInterfacesResponse']]:
        """
        Gets or sets the network interfaces of the template.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="osName")
    def os_name(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets os name.
        """
        return pulumi.get(self, "os_name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the type of the os.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system data.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource Type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Unique ID of the virtual machine template.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter(name="vmmServerId")
    def vmm_server_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        ARM Id of the vmmServer resource in which this resource resides.
        """
        return pulumi.get(self, "vmm_server_id")

