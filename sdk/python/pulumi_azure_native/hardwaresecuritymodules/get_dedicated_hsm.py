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
    'GetDedicatedHsmResult',
    'AwaitableGetDedicatedHsmResult',
    'get_dedicated_hsm',
    'get_dedicated_hsm_output',
]

@pulumi.output_type
class GetDedicatedHsmResult:
    """
    Resource information with extended details.
    """
    def __init__(__self__, azure_api_version=None, id=None, location=None, management_network_profile=None, name=None, network_profile=None, provisioning_state=None, sku=None, stamp_id=None, status_message=None, system_data=None, tags=None, type=None, zones=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if management_network_profile and not isinstance(management_network_profile, dict):
            raise TypeError("Expected argument 'management_network_profile' to be a dict")
        pulumi.set(__self__, "management_network_profile", management_network_profile)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_profile and not isinstance(network_profile, dict):
            raise TypeError("Expected argument 'network_profile' to be a dict")
        pulumi.set(__self__, "network_profile", network_profile)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if stamp_id and not isinstance(stamp_id, str):
            raise TypeError("Expected argument 'stamp_id' to be a str")
        pulumi.set(__self__, "stamp_id", stamp_id)
        if status_message and not isinstance(status_message, str):
            raise TypeError("Expected argument 'status_message' to be a str")
        pulumi.set(__self__, "status_message", status_message)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

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
    @pulumi.getter(name="managementNetworkProfile")
    def management_network_profile(self) -> Optional['outputs.NetworkProfileResponse']:
        """
        Specifies the management network interfaces of the dedicated hsm.
        """
        return pulumi.get(self, "management_network_profile")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkProfile")
    def network_profile(self) -> Optional['outputs.NetworkProfileResponse']:
        """
        Specifies the network interfaces of the dedicated hsm.
        """
        return pulumi.get(self, "network_profile")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.SkuResponse':
        """
        SKU details
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="stampId")
    def stamp_id(self) -> Optional[builtins.str]:
        """
        This field will be used when RP does not support Availability zones.
        """
        return pulumi.get(self, "stamp_id")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> builtins.str:
        """
        Resource Status Message.
        """
        return pulumi.get(self, "status_message")

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
    def zones(self) -> Optional[Sequence[builtins.str]]:
        """
        The Dedicated Hsm zones.
        """
        return pulumi.get(self, "zones")


class AwaitableGetDedicatedHsmResult(GetDedicatedHsmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDedicatedHsmResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            location=self.location,
            management_network_profile=self.management_network_profile,
            name=self.name,
            network_profile=self.network_profile,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            stamp_id=self.stamp_id,
            status_message=self.status_message,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            zones=self.zones)


def get_dedicated_hsm(name: Optional[builtins.str] = None,
                      resource_group_name: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDedicatedHsmResult:
    """
    Gets the specified Azure dedicated HSM.

    Uses Azure REST API version 2024-06-30-preview.

    Other available API versions: 2021-11-30, 2025-03-31. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hardwaresecuritymodules [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str name: Name of the dedicated Hsm
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:hardwaresecuritymodules:getDedicatedHsm', __args__, opts=opts, typ=GetDedicatedHsmResult).value

    return AwaitableGetDedicatedHsmResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        management_network_profile=pulumi.get(__ret__, 'management_network_profile'),
        name=pulumi.get(__ret__, 'name'),
        network_profile=pulumi.get(__ret__, 'network_profile'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        sku=pulumi.get(__ret__, 'sku'),
        stamp_id=pulumi.get(__ret__, 'stamp_id'),
        status_message=pulumi.get(__ret__, 'status_message'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        zones=pulumi.get(__ret__, 'zones'))
def get_dedicated_hsm_output(name: Optional[pulumi.Input[builtins.str]] = None,
                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDedicatedHsmResult]:
    """
    Gets the specified Azure dedicated HSM.

    Uses Azure REST API version 2024-06-30-preview.

    Other available API versions: 2021-11-30, 2025-03-31. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hardwaresecuritymodules [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str name: Name of the dedicated Hsm
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:hardwaresecuritymodules:getDedicatedHsm', __args__, opts=opts, typ=GetDedicatedHsmResult)
    return __ret__.apply(lambda __response__: GetDedicatedHsmResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        management_network_profile=pulumi.get(__response__, 'management_network_profile'),
        name=pulumi.get(__response__, 'name'),
        network_profile=pulumi.get(__response__, 'network_profile'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        sku=pulumi.get(__response__, 'sku'),
        stamp_id=pulumi.get(__response__, 'stamp_id'),
        status_message=pulumi.get(__response__, 'status_message'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        zones=pulumi.get(__response__, 'zones')))
