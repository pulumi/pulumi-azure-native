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
    'GetMediaServiceResult',
    'AwaitableGetMediaServiceResult',
    'get_media_service',
    'get_media_service_output',
]

@pulumi.output_type
class GetMediaServiceResult:
    """
    A Media Services account.
    """
    def __init__(__self__, azure_api_version=None, encryption=None, id=None, identity=None, key_delivery=None, location=None, media_service_id=None, minimum_tls_version=None, name=None, private_endpoint_connections=None, provisioning_state=None, public_network_access=None, storage_accounts=None, storage_authentication=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if encryption and not isinstance(encryption, dict):
            raise TypeError("Expected argument 'encryption' to be a dict")
        pulumi.set(__self__, "encryption", encryption)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if key_delivery and not isinstance(key_delivery, dict):
            raise TypeError("Expected argument 'key_delivery' to be a dict")
        pulumi.set(__self__, "key_delivery", key_delivery)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if media_service_id and not isinstance(media_service_id, str):
            raise TypeError("Expected argument 'media_service_id' to be a str")
        pulumi.set(__self__, "media_service_id", media_service_id)
        if minimum_tls_version and not isinstance(minimum_tls_version, str):
            raise TypeError("Expected argument 'minimum_tls_version' to be a str")
        pulumi.set(__self__, "minimum_tls_version", minimum_tls_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_endpoint_connections and not isinstance(private_endpoint_connections, list):
            raise TypeError("Expected argument 'private_endpoint_connections' to be a list")
        pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if public_network_access and not isinstance(public_network_access, str):
            raise TypeError("Expected argument 'public_network_access' to be a str")
        pulumi.set(__self__, "public_network_access", public_network_access)
        if storage_accounts and not isinstance(storage_accounts, list):
            raise TypeError("Expected argument 'storage_accounts' to be a list")
        pulumi.set(__self__, "storage_accounts", storage_accounts)
        if storage_authentication and not isinstance(storage_authentication, str):
            raise TypeError("Expected argument 'storage_authentication' to be a str")
        pulumi.set(__self__, "storage_authentication", storage_authentication)
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
    def encryption(self) -> Optional['outputs.AccountEncryptionResponse']:
        """
        The account encryption properties.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.MediaServiceIdentityResponse']:
        """
        The Managed Identity for the Media Services account.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="keyDelivery")
    def key_delivery(self) -> Optional['outputs.KeyDeliveryResponse']:
        """
        The Key Delivery properties for Media Services account.
        """
        return pulumi.get(self, "key_delivery")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mediaServiceId")
    def media_service_id(self) -> builtins.str:
        """
        The Media Services account ID.
        """
        return pulumi.get(self, "media_service_id")

    @property
    @pulumi.getter(name="minimumTlsVersion")
    def minimum_tls_version(self) -> Optional[builtins.str]:
        """
        The minimum TLS version allowed for this account's requests. This is an optional property. If unspecified, a secure default value will be used.
        """
        return pulumi.get(self, "minimum_tls_version")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Sequence['outputs.PrivateEndpointConnectionResponse']:
        """
        The Private Endpoint Connections created for the Media Service account.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning state of the Media Services account.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[builtins.str]:
        """
        Whether or not public network access is allowed for resources under the Media Services account.
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter(name="storageAccounts")
    def storage_accounts(self) -> Optional[Sequence['outputs.StorageAccountResponse']]:
        """
        The storage accounts for this resource.
        """
        return pulumi.get(self, "storage_accounts")

    @property
    @pulumi.getter(name="storageAuthentication")
    def storage_authentication(self) -> Optional[builtins.str]:
        return pulumi.get(self, "storage_authentication")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system metadata relating to this resource.
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


class AwaitableGetMediaServiceResult(GetMediaServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMediaServiceResult(
            azure_api_version=self.azure_api_version,
            encryption=self.encryption,
            id=self.id,
            identity=self.identity,
            key_delivery=self.key_delivery,
            location=self.location,
            media_service_id=self.media_service_id,
            minimum_tls_version=self.minimum_tls_version,
            name=self.name,
            private_endpoint_connections=self.private_endpoint_connections,
            provisioning_state=self.provisioning_state,
            public_network_access=self.public_network_access,
            storage_accounts=self.storage_accounts,
            storage_authentication=self.storage_authentication,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_media_service(account_name: Optional[builtins.str] = None,
                      resource_group_name: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMediaServiceResult:
    """
    Get the details of a Media Services account

    Uses Azure REST API version 2023-01-01.

    Other available API versions: 2015-10-01, 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-05-01, 2021-06-01, 2021-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: The Media Services account name.
    :param builtins.str resource_group_name: The name of the resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:media:getMediaService', __args__, opts=opts, typ=GetMediaServiceResult).value

    return AwaitableGetMediaServiceResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        encryption=pulumi.get(__ret__, 'encryption'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        key_delivery=pulumi.get(__ret__, 'key_delivery'),
        location=pulumi.get(__ret__, 'location'),
        media_service_id=pulumi.get(__ret__, 'media_service_id'),
        minimum_tls_version=pulumi.get(__ret__, 'minimum_tls_version'),
        name=pulumi.get(__ret__, 'name'),
        private_endpoint_connections=pulumi.get(__ret__, 'private_endpoint_connections'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        public_network_access=pulumi.get(__ret__, 'public_network_access'),
        storage_accounts=pulumi.get(__ret__, 'storage_accounts'),
        storage_authentication=pulumi.get(__ret__, 'storage_authentication'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_media_service_output(account_name: Optional[pulumi.Input[builtins.str]] = None,
                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMediaServiceResult]:
    """
    Get the details of a Media Services account

    Uses Azure REST API version 2023-01-01.

    Other available API versions: 2015-10-01, 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-05-01, 2021-06-01, 2021-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: The Media Services account name.
    :param builtins.str resource_group_name: The name of the resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:media:getMediaService', __args__, opts=opts, typ=GetMediaServiceResult)
    return __ret__.apply(lambda __response__: GetMediaServiceResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        encryption=pulumi.get(__response__, 'encryption'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        key_delivery=pulumi.get(__response__, 'key_delivery'),
        location=pulumi.get(__response__, 'location'),
        media_service_id=pulumi.get(__response__, 'media_service_id'),
        minimum_tls_version=pulumi.get(__response__, 'minimum_tls_version'),
        name=pulumi.get(__response__, 'name'),
        private_endpoint_connections=pulumi.get(__response__, 'private_endpoint_connections'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        public_network_access=pulumi.get(__response__, 'public_network_access'),
        storage_accounts=pulumi.get(__response__, 'storage_accounts'),
        storage_authentication=pulumi.get(__response__, 'storage_authentication'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
