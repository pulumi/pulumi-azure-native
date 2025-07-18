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
    'GetAssetEndpointProfileResult',
    'AwaitableGetAssetEndpointProfileResult',
    'get_asset_endpoint_profile',
    'get_asset_endpoint_profile_output',
]

@pulumi.output_type
class GetAssetEndpointProfileResult:
    """
    Asset Endpoint Profile definition.
    """
    def __init__(__self__, additional_configuration=None, authentication=None, azure_api_version=None, discovered_asset_endpoint_profile_ref=None, endpoint_profile_type=None, extended_location=None, id=None, location=None, name=None, provisioning_state=None, status=None, system_data=None, tags=None, target_address=None, type=None, uuid=None):
        if additional_configuration and not isinstance(additional_configuration, str):
            raise TypeError("Expected argument 'additional_configuration' to be a str")
        pulumi.set(__self__, "additional_configuration", additional_configuration)
        if authentication and not isinstance(authentication, dict):
            raise TypeError("Expected argument 'authentication' to be a dict")
        pulumi.set(__self__, "authentication", authentication)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if discovered_asset_endpoint_profile_ref and not isinstance(discovered_asset_endpoint_profile_ref, str):
            raise TypeError("Expected argument 'discovered_asset_endpoint_profile_ref' to be a str")
        pulumi.set(__self__, "discovered_asset_endpoint_profile_ref", discovered_asset_endpoint_profile_ref)
        if endpoint_profile_type and not isinstance(endpoint_profile_type, str):
            raise TypeError("Expected argument 'endpoint_profile_type' to be a str")
        pulumi.set(__self__, "endpoint_profile_type", endpoint_profile_type)
        if extended_location and not isinstance(extended_location, dict):
            raise TypeError("Expected argument 'extended_location' to be a dict")
        pulumi.set(__self__, "extended_location", extended_location)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target_address and not isinstance(target_address, str):
            raise TypeError("Expected argument 'target_address' to be a str")
        pulumi.set(__self__, "target_address", target_address)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="additionalConfiguration")
    def additional_configuration(self) -> Optional[builtins.str]:
        """
        Stringified JSON that contains connectivity type specific further configuration (e.g. OPC UA, Modbus, ONVIF).
        """
        return pulumi.get(self, "additional_configuration")

    @property
    @pulumi.getter
    def authentication(self) -> Optional['outputs.AuthenticationResponse']:
        """
        Defines the client authentication mechanism to the server.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="discoveredAssetEndpointProfileRef")
    def discovered_asset_endpoint_profile_ref(self) -> Optional[builtins.str]:
        """
        Reference to a discovered asset endpoint profile. Populated only if the asset endpoint profile has been created from discovery flow. Discovered asset endpoint profile name must be provided.
        """
        return pulumi.get(self, "discovered_asset_endpoint_profile_ref")

    @property
    @pulumi.getter(name="endpointProfileType")
    def endpoint_profile_type(self) -> builtins.str:
        """
        Defines the configuration for the connector type that is being used with the endpoint profile.
        """
        return pulumi.get(self, "endpoint_profile_type")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> 'outputs.ExtendedLocationResponse':
        """
        The extended location.
        """
        return pulumi.get(self, "extended_location")

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
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.AssetEndpointProfileStatusResponse':
        """
        Read only object to reflect changes that have occurred on the Edge. Similar to Kubernetes status property for custom resources.
        """
        return pulumi.get(self, "status")

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
    @pulumi.getter(name="targetAddress")
    def target_address(self) -> builtins.str:
        """
        The local valid URI specifying the network address/DNS name of a southbound device. The scheme part of the targetAddress URI specifies the type of the device. The additionalConfiguration field holds further connector type specific configuration.
        """
        return pulumi.get(self, "target_address")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> builtins.str:
        """
        Globally unique, immutable, non-reusable id.
        """
        return pulumi.get(self, "uuid")


class AwaitableGetAssetEndpointProfileResult(GetAssetEndpointProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssetEndpointProfileResult(
            additional_configuration=self.additional_configuration,
            authentication=self.authentication,
            azure_api_version=self.azure_api_version,
            discovered_asset_endpoint_profile_ref=self.discovered_asset_endpoint_profile_ref,
            endpoint_profile_type=self.endpoint_profile_type,
            extended_location=self.extended_location,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            status=self.status,
            system_data=self.system_data,
            tags=self.tags,
            target_address=self.target_address,
            type=self.type,
            uuid=self.uuid)


def get_asset_endpoint_profile(asset_endpoint_profile_name: Optional[builtins.str] = None,
                               resource_group_name: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAssetEndpointProfileResult:
    """
    Get a AssetEndpointProfile

    Uses Azure REST API version 2024-11-01.

    Other available API versions: 2023-11-01-preview, 2024-09-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str asset_endpoint_profile_name: Asset Endpoint Profile name parameter.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['assetEndpointProfileName'] = asset_endpoint_profile_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:deviceregistry:getAssetEndpointProfile', __args__, opts=opts, typ=GetAssetEndpointProfileResult).value

    return AwaitableGetAssetEndpointProfileResult(
        additional_configuration=pulumi.get(__ret__, 'additional_configuration'),
        authentication=pulumi.get(__ret__, 'authentication'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        discovered_asset_endpoint_profile_ref=pulumi.get(__ret__, 'discovered_asset_endpoint_profile_ref'),
        endpoint_profile_type=pulumi.get(__ret__, 'endpoint_profile_type'),
        extended_location=pulumi.get(__ret__, 'extended_location'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        status=pulumi.get(__ret__, 'status'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        target_address=pulumi.get(__ret__, 'target_address'),
        type=pulumi.get(__ret__, 'type'),
        uuid=pulumi.get(__ret__, 'uuid'))
def get_asset_endpoint_profile_output(asset_endpoint_profile_name: Optional[pulumi.Input[builtins.str]] = None,
                                      resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAssetEndpointProfileResult]:
    """
    Get a AssetEndpointProfile

    Uses Azure REST API version 2024-11-01.

    Other available API versions: 2023-11-01-preview, 2024-09-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str asset_endpoint_profile_name: Asset Endpoint Profile name parameter.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['assetEndpointProfileName'] = asset_endpoint_profile_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:deviceregistry:getAssetEndpointProfile', __args__, opts=opts, typ=GetAssetEndpointProfileResult)
    return __ret__.apply(lambda __response__: GetAssetEndpointProfileResult(
        additional_configuration=pulumi.get(__response__, 'additional_configuration'),
        authentication=pulumi.get(__response__, 'authentication'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        discovered_asset_endpoint_profile_ref=pulumi.get(__response__, 'discovered_asset_endpoint_profile_ref'),
        endpoint_profile_type=pulumi.get(__response__, 'endpoint_profile_type'),
        extended_location=pulumi.get(__response__, 'extended_location'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        status=pulumi.get(__response__, 'status'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        target_address=pulumi.get(__response__, 'target_address'),
        type=pulumi.get(__response__, 'type'),
        uuid=pulumi.get(__response__, 'uuid')))
