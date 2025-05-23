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
    'GetNetworkDeviceResult',
    'AwaitableGetNetworkDeviceResult',
    'get_network_device',
    'get_network_device_output',
]

@pulumi.output_type
class GetNetworkDeviceResult:
    """
    The Network Device resource definition.
    """
    def __init__(__self__, administrative_state=None, annotation=None, azure_api_version=None, configuration_state=None, host_name=None, id=None, location=None, management_ipv4_address=None, management_ipv6_address=None, name=None, network_device_role=None, network_device_sku=None, network_rack_id=None, provisioning_state=None, serial_number=None, system_data=None, tags=None, type=None, version=None):
        if administrative_state and not isinstance(administrative_state, str):
            raise TypeError("Expected argument 'administrative_state' to be a str")
        pulumi.set(__self__, "administrative_state", administrative_state)
        if annotation and not isinstance(annotation, str):
            raise TypeError("Expected argument 'annotation' to be a str")
        pulumi.set(__self__, "annotation", annotation)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if configuration_state and not isinstance(configuration_state, str):
            raise TypeError("Expected argument 'configuration_state' to be a str")
        pulumi.set(__self__, "configuration_state", configuration_state)
        if host_name and not isinstance(host_name, str):
            raise TypeError("Expected argument 'host_name' to be a str")
        pulumi.set(__self__, "host_name", host_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if management_ipv4_address and not isinstance(management_ipv4_address, str):
            raise TypeError("Expected argument 'management_ipv4_address' to be a str")
        pulumi.set(__self__, "management_ipv4_address", management_ipv4_address)
        if management_ipv6_address and not isinstance(management_ipv6_address, str):
            raise TypeError("Expected argument 'management_ipv6_address' to be a str")
        pulumi.set(__self__, "management_ipv6_address", management_ipv6_address)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_device_role and not isinstance(network_device_role, str):
            raise TypeError("Expected argument 'network_device_role' to be a str")
        pulumi.set(__self__, "network_device_role", network_device_role)
        if network_device_sku and not isinstance(network_device_sku, str):
            raise TypeError("Expected argument 'network_device_sku' to be a str")
        pulumi.set(__self__, "network_device_sku", network_device_sku)
        if network_rack_id and not isinstance(network_rack_id, str):
            raise TypeError("Expected argument 'network_rack_id' to be a str")
        pulumi.set(__self__, "network_rack_id", network_rack_id)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if serial_number and not isinstance(serial_number, str):
            raise TypeError("Expected argument 'serial_number' to be a str")
        pulumi.set(__self__, "serial_number", serial_number)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="administrativeState")
    def administrative_state(self) -> builtins.str:
        """
        Administrative state of the resource.
        """
        return pulumi.get(self, "administrative_state")

    @property
    @pulumi.getter
    def annotation(self) -> Optional[builtins.str]:
        """
        Switch configuration description.
        """
        return pulumi.get(self, "annotation")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="configurationState")
    def configuration_state(self) -> builtins.str:
        """
        Configuration state of the resource.
        """
        return pulumi.get(self, "configuration_state")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> Optional[builtins.str]:
        """
        The host name of the device.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementIpv4Address")
    def management_ipv4_address(self) -> builtins.str:
        """
        Management IPv4 Address.
        """
        return pulumi.get(self, "management_ipv4_address")

    @property
    @pulumi.getter(name="managementIpv6Address")
    def management_ipv6_address(self) -> builtins.str:
        """
        Management IPv6 Address.
        """
        return pulumi.get(self, "management_ipv6_address")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkDeviceRole")
    def network_device_role(self) -> builtins.str:
        """
        NetworkDeviceRole is the device role: Example: CE | ToR.
        """
        return pulumi.get(self, "network_device_role")

    @property
    @pulumi.getter(name="networkDeviceSku")
    def network_device_sku(self) -> Optional[builtins.str]:
        """
        Network Device SKU name.
        """
        return pulumi.get(self, "network_device_sku")

    @property
    @pulumi.getter(name="networkRackId")
    def network_rack_id(self) -> builtins.str:
        """
        Reference to network rack resource id.
        """
        return pulumi.get(self, "network_rack_id")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> builtins.str:
        """
        Serial number of the device. Format of serial Number - Make;Model;HardwareRevisionId;SerialNumber.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> builtins.str:
        """
        Current version of the device as defined in SKU.
        """
        return pulumi.get(self, "version")


class AwaitableGetNetworkDeviceResult(GetNetworkDeviceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkDeviceResult(
            administrative_state=self.administrative_state,
            annotation=self.annotation,
            azure_api_version=self.azure_api_version,
            configuration_state=self.configuration_state,
            host_name=self.host_name,
            id=self.id,
            location=self.location,
            management_ipv4_address=self.management_ipv4_address,
            management_ipv6_address=self.management_ipv6_address,
            name=self.name,
            network_device_role=self.network_device_role,
            network_device_sku=self.network_device_sku,
            network_rack_id=self.network_rack_id,
            provisioning_state=self.provisioning_state,
            serial_number=self.serial_number,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            version=self.version)


def get_network_device(network_device_name: Optional[builtins.str] = None,
                       resource_group_name: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkDeviceResult:
    """
    Gets the Network Device resource details.

    Uses Azure REST API version 2023-06-15.

    Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str network_device_name: Name of the Network Device.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['networkDeviceName'] = network_device_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:managednetworkfabric:getNetworkDevice', __args__, opts=opts, typ=GetNetworkDeviceResult).value

    return AwaitableGetNetworkDeviceResult(
        administrative_state=pulumi.get(__ret__, 'administrative_state'),
        annotation=pulumi.get(__ret__, 'annotation'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        configuration_state=pulumi.get(__ret__, 'configuration_state'),
        host_name=pulumi.get(__ret__, 'host_name'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        management_ipv4_address=pulumi.get(__ret__, 'management_ipv4_address'),
        management_ipv6_address=pulumi.get(__ret__, 'management_ipv6_address'),
        name=pulumi.get(__ret__, 'name'),
        network_device_role=pulumi.get(__ret__, 'network_device_role'),
        network_device_sku=pulumi.get(__ret__, 'network_device_sku'),
        network_rack_id=pulumi.get(__ret__, 'network_rack_id'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        serial_number=pulumi.get(__ret__, 'serial_number'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        version=pulumi.get(__ret__, 'version'))
def get_network_device_output(network_device_name: Optional[pulumi.Input[builtins.str]] = None,
                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNetworkDeviceResult]:
    """
    Gets the Network Device resource details.

    Uses Azure REST API version 2023-06-15.

    Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str network_device_name: Name of the Network Device.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['networkDeviceName'] = network_device_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:managednetworkfabric:getNetworkDevice', __args__, opts=opts, typ=GetNetworkDeviceResult)
    return __ret__.apply(lambda __response__: GetNetworkDeviceResult(
        administrative_state=pulumi.get(__response__, 'administrative_state'),
        annotation=pulumi.get(__response__, 'annotation'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        configuration_state=pulumi.get(__response__, 'configuration_state'),
        host_name=pulumi.get(__response__, 'host_name'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        management_ipv4_address=pulumi.get(__response__, 'management_ipv4_address'),
        management_ipv6_address=pulumi.get(__response__, 'management_ipv6_address'),
        name=pulumi.get(__response__, 'name'),
        network_device_role=pulumi.get(__response__, 'network_device_role'),
        network_device_sku=pulumi.get(__response__, 'network_device_sku'),
        network_rack_id=pulumi.get(__response__, 'network_rack_id'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        serial_number=pulumi.get(__response__, 'serial_number'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        version=pulumi.get(__response__, 'version')))
