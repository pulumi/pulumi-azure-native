# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetNetworkRackResult',
    'AwaitableGetNetworkRackResult',
    'get_network_rack',
    'get_network_rack_output',
]

@pulumi.output_type
class GetNetworkRackResult:
    """
    The NetworkRack resource definition.
    """
    def __init__(__self__, annotation=None, id=None, location=None, name=None, network_devices=None, network_fabric_id=None, network_rack_sku=None, provisioning_state=None, system_data=None, tags=None, type=None):
        if annotation and not isinstance(annotation, str):
            raise TypeError("Expected argument 'annotation' to be a str")
        pulumi.set(__self__, "annotation", annotation)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_devices and not isinstance(network_devices, list):
            raise TypeError("Expected argument 'network_devices' to be a list")
        pulumi.set(__self__, "network_devices", network_devices)
        if network_fabric_id and not isinstance(network_fabric_id, str):
            raise TypeError("Expected argument 'network_fabric_id' to be a str")
        pulumi.set(__self__, "network_fabric_id", network_fabric_id)
        if network_rack_sku and not isinstance(network_rack_sku, str):
            raise TypeError("Expected argument 'network_rack_sku' to be a str")
        pulumi.set(__self__, "network_rack_sku", network_rack_sku)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def annotation(self) -> Optional[str]:
        """
        Switch configuration description.
        """
        return pulumi.get(self, "annotation")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkDevices")
    def network_devices(self) -> Sequence[str]:
        """
        List of network device ARM resource ids.
        """
        return pulumi.get(self, "network_devices")

    @property
    @pulumi.getter(name="networkFabricId")
    def network_fabric_id(self) -> str:
        """
        Network Fabric ARM resource id.
        """
        return pulumi.get(self, "network_fabric_id")

    @property
    @pulumi.getter(name="networkRackSku")
    def network_rack_sku(self) -> str:
        """
        Network Rack SKU name.
        """
        return pulumi.get(self, "network_rack_sku")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Gets the provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetNetworkRackResult(GetNetworkRackResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkRackResult(
            annotation=self.annotation,
            id=self.id,
            location=self.location,
            name=self.name,
            network_devices=self.network_devices,
            network_fabric_id=self.network_fabric_id,
            network_rack_sku=self.network_rack_sku,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_network_rack(network_rack_name: Optional[str] = None,
                     resource_group_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkRackResult:
    """
    Get Network Rack resource details.

    Uses Azure REST API version 2023-02-01-preview.

    Other available API versions: 2023-06-15.


    :param str network_rack_name: Name of the Network Rack
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['networkRackName'] = network_rack_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:managednetworkfabric:getNetworkRack', __args__, opts=opts, typ=GetNetworkRackResult).value

    return AwaitableGetNetworkRackResult(
        annotation=pulumi.get(__ret__, 'annotation'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        network_devices=pulumi.get(__ret__, 'network_devices'),
        network_fabric_id=pulumi.get(__ret__, 'network_fabric_id'),
        network_rack_sku=pulumi.get(__ret__, 'network_rack_sku'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_network_rack_output(network_rack_name: Optional[pulumi.Input[str]] = None,
                            resource_group_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNetworkRackResult]:
    """
    Get Network Rack resource details.

    Uses Azure REST API version 2023-02-01-preview.

    Other available API versions: 2023-06-15.


    :param str network_rack_name: Name of the Network Rack
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['networkRackName'] = network_rack_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:managednetworkfabric:getNetworkRack', __args__, opts=opts, typ=GetNetworkRackResult)
    return __ret__.apply(lambda __response__: GetNetworkRackResult(
        annotation=pulumi.get(__response__, 'annotation'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        network_devices=pulumi.get(__response__, 'network_devices'),
        network_fabric_id=pulumi.get(__response__, 'network_fabric_id'),
        network_rack_sku=pulumi.get(__response__, 'network_rack_sku'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
