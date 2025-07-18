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
    'GetPrivateZoneResult',
    'AwaitableGetPrivateZoneResult',
    'get_private_zone',
    'get_private_zone_output',
]

@pulumi.output_type
class GetPrivateZoneResult:
    """
    Describes a Private DNS zone.
    """
    def __init__(__self__, azure_api_version=None, etag=None, id=None, internal_id=None, location=None, max_number_of_record_sets=None, max_number_of_virtual_network_links=None, max_number_of_virtual_network_links_with_registration=None, name=None, number_of_record_sets=None, number_of_virtual_network_links=None, number_of_virtual_network_links_with_registration=None, provisioning_state=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal_id and not isinstance(internal_id, str):
            raise TypeError("Expected argument 'internal_id' to be a str")
        pulumi.set(__self__, "internal_id", internal_id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if max_number_of_record_sets and not isinstance(max_number_of_record_sets, float):
            raise TypeError("Expected argument 'max_number_of_record_sets' to be a float")
        pulumi.set(__self__, "max_number_of_record_sets", max_number_of_record_sets)
        if max_number_of_virtual_network_links and not isinstance(max_number_of_virtual_network_links, float):
            raise TypeError("Expected argument 'max_number_of_virtual_network_links' to be a float")
        pulumi.set(__self__, "max_number_of_virtual_network_links", max_number_of_virtual_network_links)
        if max_number_of_virtual_network_links_with_registration and not isinstance(max_number_of_virtual_network_links_with_registration, float):
            raise TypeError("Expected argument 'max_number_of_virtual_network_links_with_registration' to be a float")
        pulumi.set(__self__, "max_number_of_virtual_network_links_with_registration", max_number_of_virtual_network_links_with_registration)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number_of_record_sets and not isinstance(number_of_record_sets, float):
            raise TypeError("Expected argument 'number_of_record_sets' to be a float")
        pulumi.set(__self__, "number_of_record_sets", number_of_record_sets)
        if number_of_virtual_network_links and not isinstance(number_of_virtual_network_links, float):
            raise TypeError("Expected argument 'number_of_virtual_network_links' to be a float")
        pulumi.set(__self__, "number_of_virtual_network_links", number_of_virtual_network_links)
        if number_of_virtual_network_links_with_registration and not isinstance(number_of_virtual_network_links_with_registration, float):
            raise TypeError("Expected argument 'number_of_virtual_network_links_with_registration' to be a float")
        pulumi.set(__self__, "number_of_virtual_network_links_with_registration", number_of_virtual_network_links_with_registration)
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
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def etag(self) -> Optional[builtins.str]:
        """
        The ETag of the zone.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="internalId")
    def internal_id(self) -> builtins.str:
        """
        Private zone internal Id
        """
        return pulumi.get(self, "internal_id")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        The Azure Region where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxNumberOfRecordSets")
    def max_number_of_record_sets(self) -> builtins.float:
        """
        The maximum number of record sets that can be created in this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        """
        return pulumi.get(self, "max_number_of_record_sets")

    @property
    @pulumi.getter(name="maxNumberOfVirtualNetworkLinks")
    def max_number_of_virtual_network_links(self) -> builtins.float:
        """
        The maximum number of virtual networks that can be linked to this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        """
        return pulumi.get(self, "max_number_of_virtual_network_links")

    @property
    @pulumi.getter(name="maxNumberOfVirtualNetworkLinksWithRegistration")
    def max_number_of_virtual_network_links_with_registration(self) -> builtins.float:
        """
        The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
        """
        return pulumi.get(self, "max_number_of_virtual_network_links_with_registration")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfRecordSets")
    def number_of_record_sets(self) -> builtins.float:
        """
        The current number of record sets in this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        """
        return pulumi.get(self, "number_of_record_sets")

    @property
    @pulumi.getter(name="numberOfVirtualNetworkLinks")
    def number_of_virtual_network_links(self) -> builtins.float:
        """
        The current number of virtual networks that are linked to this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        """
        return pulumi.get(self, "number_of_virtual_network_links")

    @property
    @pulumi.getter(name="numberOfVirtualNetworkLinksWithRegistration")
    def number_of_virtual_network_links_with_registration(self) -> builtins.float:
        """
        The current number of virtual networks that are linked to this Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
        """
        return pulumi.get(self, "number_of_virtual_network_links_with_registration")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the resource. This is a read-only property and any attempt to set this value will be ignored.
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


class AwaitableGetPrivateZoneResult(GetPrivateZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateZoneResult(
            azure_api_version=self.azure_api_version,
            etag=self.etag,
            id=self.id,
            internal_id=self.internal_id,
            location=self.location,
            max_number_of_record_sets=self.max_number_of_record_sets,
            max_number_of_virtual_network_links=self.max_number_of_virtual_network_links,
            max_number_of_virtual_network_links_with_registration=self.max_number_of_virtual_network_links_with_registration,
            name=self.name,
            number_of_record_sets=self.number_of_record_sets,
            number_of_virtual_network_links=self.number_of_virtual_network_links,
            number_of_virtual_network_links_with_registration=self.number_of_virtual_network_links_with_registration,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_private_zone(private_zone_name: Optional[builtins.str] = None,
                     resource_group_name: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateZoneResult:
    """
    Gets a Private DNS zone. Retrieves the zone properties, but not the virtual networks links or the record sets within the zone.

    Uses Azure REST API version 2024-06-01.

    Other available API versions: 2018-09-01, 2020-01-01, 2020-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native privatedns [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str private_zone_name: The name of the Private DNS zone (without a terminating dot).
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['privateZoneName'] = private_zone_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:privatedns:getPrivateZone', __args__, opts=opts, typ=GetPrivateZoneResult).value

    return AwaitableGetPrivateZoneResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        internal_id=pulumi.get(__ret__, 'internal_id'),
        location=pulumi.get(__ret__, 'location'),
        max_number_of_record_sets=pulumi.get(__ret__, 'max_number_of_record_sets'),
        max_number_of_virtual_network_links=pulumi.get(__ret__, 'max_number_of_virtual_network_links'),
        max_number_of_virtual_network_links_with_registration=pulumi.get(__ret__, 'max_number_of_virtual_network_links_with_registration'),
        name=pulumi.get(__ret__, 'name'),
        number_of_record_sets=pulumi.get(__ret__, 'number_of_record_sets'),
        number_of_virtual_network_links=pulumi.get(__ret__, 'number_of_virtual_network_links'),
        number_of_virtual_network_links_with_registration=pulumi.get(__ret__, 'number_of_virtual_network_links_with_registration'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_private_zone_output(private_zone_name: Optional[pulumi.Input[builtins.str]] = None,
                            resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPrivateZoneResult]:
    """
    Gets a Private DNS zone. Retrieves the zone properties, but not the virtual networks links or the record sets within the zone.

    Uses Azure REST API version 2024-06-01.

    Other available API versions: 2018-09-01, 2020-01-01, 2020-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native privatedns [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str private_zone_name: The name of the Private DNS zone (without a terminating dot).
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['privateZoneName'] = private_zone_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:privatedns:getPrivateZone', __args__, opts=opts, typ=GetPrivateZoneResult)
    return __ret__.apply(lambda __response__: GetPrivateZoneResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        internal_id=pulumi.get(__response__, 'internal_id'),
        location=pulumi.get(__response__, 'location'),
        max_number_of_record_sets=pulumi.get(__response__, 'max_number_of_record_sets'),
        max_number_of_virtual_network_links=pulumi.get(__response__, 'max_number_of_virtual_network_links'),
        max_number_of_virtual_network_links_with_registration=pulumi.get(__response__, 'max_number_of_virtual_network_links_with_registration'),
        name=pulumi.get(__response__, 'name'),
        number_of_record_sets=pulumi.get(__response__, 'number_of_record_sets'),
        number_of_virtual_network_links=pulumi.get(__response__, 'number_of_virtual_network_links'),
        number_of_virtual_network_links_with_registration=pulumi.get(__response__, 'number_of_virtual_network_links_with_registration'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
