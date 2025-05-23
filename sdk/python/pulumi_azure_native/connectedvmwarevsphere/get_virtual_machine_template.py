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

__all__ = [
    'GetVirtualMachineTemplateResult',
    'AwaitableGetVirtualMachineTemplateResult',
    'get_virtual_machine_template',
    'get_virtual_machine_template_output',
]

@pulumi.output_type
class GetVirtualMachineTemplateResult:
    """
    Define the virtualMachineTemplate.
    """
    def __init__(__self__, azure_api_version=None, custom_resource_name=None, disks=None, extended_location=None, firmware_type=None, folder_path=None, id=None, inventory_item_id=None, kind=None, location=None, memory_size_mb=None, mo_name=None, mo_ref_id=None, name=None, network_interfaces=None, num_cpus=None, num_cores_per_socket=None, os_name=None, os_type=None, provisioning_state=None, statuses=None, system_data=None, tags=None, tools_version=None, tools_version_status=None, type=None, uuid=None, v_center_id=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if custom_resource_name and not isinstance(custom_resource_name, str):
            raise TypeError("Expected argument 'custom_resource_name' to be a str")
        pulumi.set(__self__, "custom_resource_name", custom_resource_name)
        if disks and not isinstance(disks, list):
            raise TypeError("Expected argument 'disks' to be a list")
        pulumi.set(__self__, "disks", disks)
        if extended_location and not isinstance(extended_location, dict):
            raise TypeError("Expected argument 'extended_location' to be a dict")
        pulumi.set(__self__, "extended_location", extended_location)
        if firmware_type and not isinstance(firmware_type, str):
            raise TypeError("Expected argument 'firmware_type' to be a str")
        pulumi.set(__self__, "firmware_type", firmware_type)
        if folder_path and not isinstance(folder_path, str):
            raise TypeError("Expected argument 'folder_path' to be a str")
        pulumi.set(__self__, "folder_path", folder_path)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inventory_item_id and not isinstance(inventory_item_id, str):
            raise TypeError("Expected argument 'inventory_item_id' to be a str")
        pulumi.set(__self__, "inventory_item_id", inventory_item_id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if memory_size_mb and not isinstance(memory_size_mb, int):
            raise TypeError("Expected argument 'memory_size_mb' to be a int")
        pulumi.set(__self__, "memory_size_mb", memory_size_mb)
        if mo_name and not isinstance(mo_name, str):
            raise TypeError("Expected argument 'mo_name' to be a str")
        pulumi.set(__self__, "mo_name", mo_name)
        if mo_ref_id and not isinstance(mo_ref_id, str):
            raise TypeError("Expected argument 'mo_ref_id' to be a str")
        pulumi.set(__self__, "mo_ref_id", mo_ref_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_interfaces and not isinstance(network_interfaces, list):
            raise TypeError("Expected argument 'network_interfaces' to be a list")
        pulumi.set(__self__, "network_interfaces", network_interfaces)
        if num_cpus and not isinstance(num_cpus, int):
            raise TypeError("Expected argument 'num_cpus' to be a int")
        pulumi.set(__self__, "num_cpus", num_cpus)
        if num_cores_per_socket and not isinstance(num_cores_per_socket, int):
            raise TypeError("Expected argument 'num_cores_per_socket' to be a int")
        pulumi.set(__self__, "num_cores_per_socket", num_cores_per_socket)
        if os_name and not isinstance(os_name, str):
            raise TypeError("Expected argument 'os_name' to be a str")
        pulumi.set(__self__, "os_name", os_name)
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if tools_version and not isinstance(tools_version, str):
            raise TypeError("Expected argument 'tools_version' to be a str")
        pulumi.set(__self__, "tools_version", tools_version)
        if tools_version_status and not isinstance(tools_version_status, str):
            raise TypeError("Expected argument 'tools_version_status' to be a str")
        pulumi.set(__self__, "tools_version_status", tools_version_status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if v_center_id and not isinstance(v_center_id, str):
            raise TypeError("Expected argument 'v_center_id' to be a str")
        pulumi.set(__self__, "v_center_id", v_center_id)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="customResourceName")
    def custom_resource_name(self) -> builtins.str:
        """
        Gets the name of the corresponding resource in Kubernetes.
        """
        return pulumi.get(self, "custom_resource_name")

    @property
    @pulumi.getter
    def disks(self) -> Sequence['outputs.VirtualDiskResponse']:
        """
        Gets or sets the disks the template.
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> Optional['outputs.ExtendedLocationResponse']:
        """
        Gets or sets the extended location.
        """
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter(name="firmwareType")
    def firmware_type(self) -> builtins.str:
        """
        Firmware type
        """
        return pulumi.get(self, "firmware_type")

    @property
    @pulumi.getter(name="folderPath")
    def folder_path(self) -> builtins.str:
        """
        Gets or sets the folder path of the template.
        """
        return pulumi.get(self, "folder_path")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Gets or sets the Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inventoryItemId")
    def inventory_item_id(self) -> Optional[builtins.str]:
        """
        Gets or sets the inventory Item ID for the virtual machine template.
        """
        return pulumi.get(self, "inventory_item_id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[builtins.str]:
        """
        Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Gets or sets the location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="memorySizeMB")
    def memory_size_mb(self) -> builtins.int:
        """
        Gets or sets memory size in MBs for the template.
        """
        return pulumi.get(self, "memory_size_mb")

    @property
    @pulumi.getter(name="moName")
    def mo_name(self) -> builtins.str:
        """
        Gets or sets the vCenter Managed Object name for the virtual machine template.
        """
        return pulumi.get(self, "mo_name")

    @property
    @pulumi.getter(name="moRefId")
    def mo_ref_id(self) -> Optional[builtins.str]:
        """
        Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual machine
        template.
        """
        return pulumi.get(self, "mo_ref_id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Gets or sets the name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> Sequence['outputs.NetworkInterfaceResponse']:
        """
        Gets or sets the network interfaces of the template.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="numCPUs")
    def num_cpus(self) -> builtins.int:
        """
        Gets or sets the number of vCPUs for the template.
        """
        return pulumi.get(self, "num_cpus")

    @property
    @pulumi.getter(name="numCoresPerSocket")
    def num_cores_per_socket(self) -> builtins.int:
        """
        Gets or sets the number of cores per socket for the template.
        Defaults to 1 if unspecified.
        """
        return pulumi.get(self, "num_cores_per_socket")

    @property
    @pulumi.getter(name="osName")
    def os_name(self) -> builtins.str:
        """
        Gets or sets os name.
        """
        return pulumi.get(self, "os_name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> builtins.str:
        """
        Gets or sets the type of the os.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Gets the provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def statuses(self) -> Sequence['outputs.ResourceStatusResponse']:
        """
        The resource status information.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system data.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Gets or sets the Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="toolsVersion")
    def tools_version(self) -> builtins.str:
        """
        Gets or sets the current version of VMware Tools.
        """
        return pulumi.get(self, "tools_version")

    @property
    @pulumi.getter(name="toolsVersionStatus")
    def tools_version_status(self) -> builtins.str:
        """
        Gets or sets the current version status of VMware Tools installed in the guest operating system.
        """
        return pulumi.get(self, "tools_version_status")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Gets or sets the type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> builtins.str:
        """
        Gets or sets a unique identifier for this resource.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter(name="vCenterId")
    def v_center_id(self) -> Optional[builtins.str]:
        """
        Gets or sets the ARM Id of the vCenter resource in which this template resides.
        """
        return pulumi.get(self, "v_center_id")


class AwaitableGetVirtualMachineTemplateResult(GetVirtualMachineTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualMachineTemplateResult(
            azure_api_version=self.azure_api_version,
            custom_resource_name=self.custom_resource_name,
            disks=self.disks,
            extended_location=self.extended_location,
            firmware_type=self.firmware_type,
            folder_path=self.folder_path,
            id=self.id,
            inventory_item_id=self.inventory_item_id,
            kind=self.kind,
            location=self.location,
            memory_size_mb=self.memory_size_mb,
            mo_name=self.mo_name,
            mo_ref_id=self.mo_ref_id,
            name=self.name,
            network_interfaces=self.network_interfaces,
            num_cpus=self.num_cpus,
            num_cores_per_socket=self.num_cores_per_socket,
            os_name=self.os_name,
            os_type=self.os_type,
            provisioning_state=self.provisioning_state,
            statuses=self.statuses,
            system_data=self.system_data,
            tags=self.tags,
            tools_version=self.tools_version,
            tools_version_status=self.tools_version_status,
            type=self.type,
            uuid=self.uuid,
            v_center_id=self.v_center_id)


def get_virtual_machine_template(resource_group_name: Optional[builtins.str] = None,
                                 virtual_machine_template_name: Optional[builtins.str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualMachineTemplateResult:
    """
    Implements virtual machine template GET method.

    Uses Azure REST API version 2023-12-01.

    Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The Resource Group Name.
    :param builtins.str virtual_machine_template_name: Name of the virtual machine template resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['virtualMachineTemplateName'] = virtual_machine_template_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:connectedvmwarevsphere:getVirtualMachineTemplate', __args__, opts=opts, typ=GetVirtualMachineTemplateResult).value

    return AwaitableGetVirtualMachineTemplateResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        custom_resource_name=pulumi.get(__ret__, 'custom_resource_name'),
        disks=pulumi.get(__ret__, 'disks'),
        extended_location=pulumi.get(__ret__, 'extended_location'),
        firmware_type=pulumi.get(__ret__, 'firmware_type'),
        folder_path=pulumi.get(__ret__, 'folder_path'),
        id=pulumi.get(__ret__, 'id'),
        inventory_item_id=pulumi.get(__ret__, 'inventory_item_id'),
        kind=pulumi.get(__ret__, 'kind'),
        location=pulumi.get(__ret__, 'location'),
        memory_size_mb=pulumi.get(__ret__, 'memory_size_mb'),
        mo_name=pulumi.get(__ret__, 'mo_name'),
        mo_ref_id=pulumi.get(__ret__, 'mo_ref_id'),
        name=pulumi.get(__ret__, 'name'),
        network_interfaces=pulumi.get(__ret__, 'network_interfaces'),
        num_cpus=pulumi.get(__ret__, 'num_cpus'),
        num_cores_per_socket=pulumi.get(__ret__, 'num_cores_per_socket'),
        os_name=pulumi.get(__ret__, 'os_name'),
        os_type=pulumi.get(__ret__, 'os_type'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        statuses=pulumi.get(__ret__, 'statuses'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        tools_version=pulumi.get(__ret__, 'tools_version'),
        tools_version_status=pulumi.get(__ret__, 'tools_version_status'),
        type=pulumi.get(__ret__, 'type'),
        uuid=pulumi.get(__ret__, 'uuid'),
        v_center_id=pulumi.get(__ret__, 'v_center_id'))
def get_virtual_machine_template_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                        virtual_machine_template_name: Optional[pulumi.Input[builtins.str]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVirtualMachineTemplateResult]:
    """
    Implements virtual machine template GET method.

    Uses Azure REST API version 2023-12-01.

    Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The Resource Group Name.
    :param builtins.str virtual_machine_template_name: Name of the virtual machine template resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['virtualMachineTemplateName'] = virtual_machine_template_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:connectedvmwarevsphere:getVirtualMachineTemplate', __args__, opts=opts, typ=GetVirtualMachineTemplateResult)
    return __ret__.apply(lambda __response__: GetVirtualMachineTemplateResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        custom_resource_name=pulumi.get(__response__, 'custom_resource_name'),
        disks=pulumi.get(__response__, 'disks'),
        extended_location=pulumi.get(__response__, 'extended_location'),
        firmware_type=pulumi.get(__response__, 'firmware_type'),
        folder_path=pulumi.get(__response__, 'folder_path'),
        id=pulumi.get(__response__, 'id'),
        inventory_item_id=pulumi.get(__response__, 'inventory_item_id'),
        kind=pulumi.get(__response__, 'kind'),
        location=pulumi.get(__response__, 'location'),
        memory_size_mb=pulumi.get(__response__, 'memory_size_mb'),
        mo_name=pulumi.get(__response__, 'mo_name'),
        mo_ref_id=pulumi.get(__response__, 'mo_ref_id'),
        name=pulumi.get(__response__, 'name'),
        network_interfaces=pulumi.get(__response__, 'network_interfaces'),
        num_cpus=pulumi.get(__response__, 'num_cpus'),
        num_cores_per_socket=pulumi.get(__response__, 'num_cores_per_socket'),
        os_name=pulumi.get(__response__, 'os_name'),
        os_type=pulumi.get(__response__, 'os_type'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        statuses=pulumi.get(__response__, 'statuses'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        tools_version=pulumi.get(__response__, 'tools_version'),
        tools_version_status=pulumi.get(__response__, 'tools_version_status'),
        type=pulumi.get(__response__, 'type'),
        uuid=pulumi.get(__response__, 'uuid'),
        v_center_id=pulumi.get(__response__, 'v_center_id')))
