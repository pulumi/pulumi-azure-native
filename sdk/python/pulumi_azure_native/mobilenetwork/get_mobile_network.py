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
    'GetMobileNetworkResult',
    'AwaitableGetMobileNetworkResult',
    'get_mobile_network',
    'get_mobile_network_output',
]

@pulumi.output_type
class GetMobileNetworkResult:
    """
    Mobile network resource.
    """
    def __init__(__self__, azure_api_version=None, id=None, identity=None, location=None, name=None, provisioning_state=None, public_land_mobile_network_identifier=None, public_land_mobile_networks=None, service_key=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if public_land_mobile_network_identifier and not isinstance(public_land_mobile_network_identifier, dict):
            raise TypeError("Expected argument 'public_land_mobile_network_identifier' to be a dict")
        pulumi.set(__self__, "public_land_mobile_network_identifier", public_land_mobile_network_identifier)
        if public_land_mobile_networks and not isinstance(public_land_mobile_networks, list):
            raise TypeError("Expected argument 'public_land_mobile_networks' to be a list")
        pulumi.set(__self__, "public_land_mobile_networks", public_land_mobile_networks)
        if service_key and not isinstance(service_key, str):
            raise TypeError("Expected argument 'service_key' to be a str")
        pulumi.set(__self__, "service_key", service_key)
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
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.ManagedServiceIdentityResponse']:
        """
        The identity used to retrieve any private keys used for SUPI concealment from Azure key vault.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the mobile network resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicLandMobileNetworkIdentifier")
    def public_land_mobile_network_identifier(self) -> 'outputs.PlmnIdResponse':
        """
        The unique public land mobile network identifier for the network. If both 'publicLandMobileNetworks' and 'publicLandMobileNetworkIdentifier' are specified, then the 'publicLandMobileNetworks' will take precedence.
        """
        return pulumi.get(self, "public_land_mobile_network_identifier")

    @property
    @pulumi.getter(name="publicLandMobileNetworks")
    def public_land_mobile_networks(self) -> Optional[Sequence['outputs.PublicLandMobileNetworkResponse']]:
        """
        A list of public land mobile networks including their identifiers. If both 'publicLandMobileNetworks' and 'publicLandMobileNetworkIdentifier' are specified, then the 'publicLandMobileNetworks' will take precedence.
        """
        return pulumi.get(self, "public_land_mobile_networks")

    @property
    @pulumi.getter(name="serviceKey")
    def service_key(self) -> builtins.str:
        """
        The mobile network resource identifier
        """
        return pulumi.get(self, "service_key")

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


class AwaitableGetMobileNetworkResult(GetMobileNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMobileNetworkResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            identity=self.identity,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            public_land_mobile_network_identifier=self.public_land_mobile_network_identifier,
            public_land_mobile_networks=self.public_land_mobile_networks,
            service_key=self.service_key,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_mobile_network(mobile_network_name: Optional[builtins.str] = None,
                       resource_group_name: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMobileNetworkResult:
    """
    Gets information about the specified mobile network.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str mobile_network_name: The name of the mobile network.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['mobileNetworkName'] = mobile_network_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:mobilenetwork:getMobileNetwork', __args__, opts=opts, typ=GetMobileNetworkResult).value

    return AwaitableGetMobileNetworkResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        public_land_mobile_network_identifier=pulumi.get(__ret__, 'public_land_mobile_network_identifier'),
        public_land_mobile_networks=pulumi.get(__ret__, 'public_land_mobile_networks'),
        service_key=pulumi.get(__ret__, 'service_key'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_mobile_network_output(mobile_network_name: Optional[pulumi.Input[builtins.str]] = None,
                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMobileNetworkResult]:
    """
    Gets information about the specified mobile network.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str mobile_network_name: The name of the mobile network.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['mobileNetworkName'] = mobile_network_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:mobilenetwork:getMobileNetwork', __args__, opts=opts, typ=GetMobileNetworkResult)
    return __ret__.apply(lambda __response__: GetMobileNetworkResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        public_land_mobile_network_identifier=pulumi.get(__response__, 'public_land_mobile_network_identifier'),
        public_land_mobile_networks=pulumi.get(__response__, 'public_land_mobile_networks'),
        service_key=pulumi.get(__response__, 'service_key'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
