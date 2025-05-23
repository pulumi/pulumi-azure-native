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
    'GetDeviceExtendedInformationResult',
    'AwaitableGetDeviceExtendedInformationResult',
    'get_device_extended_information',
    'get_device_extended_information_output',
]

@pulumi.output_type
class GetDeviceExtendedInformationResult:
    """
    The extended Info of the Data Box Edge/Gateway device.
    """
    def __init__(__self__, channel_integrity_key_name=None, channel_integrity_key_version=None, client_secret_store_id=None, client_secret_store_url=None, cloud_witness_container_name=None, cloud_witness_storage_account_name=None, cloud_witness_storage_endpoint=None, cluster_witness_type=None, device_secrets=None, encryption_key=None, encryption_key_thumbprint=None, file_share_witness_location=None, file_share_witness_username=None, id=None, key_vault_sync_status=None, name=None, resource_key=None, system_data=None, type=None):
        if channel_integrity_key_name and not isinstance(channel_integrity_key_name, str):
            raise TypeError("Expected argument 'channel_integrity_key_name' to be a str")
        pulumi.set(__self__, "channel_integrity_key_name", channel_integrity_key_name)
        if channel_integrity_key_version and not isinstance(channel_integrity_key_version, str):
            raise TypeError("Expected argument 'channel_integrity_key_version' to be a str")
        pulumi.set(__self__, "channel_integrity_key_version", channel_integrity_key_version)
        if client_secret_store_id and not isinstance(client_secret_store_id, str):
            raise TypeError("Expected argument 'client_secret_store_id' to be a str")
        pulumi.set(__self__, "client_secret_store_id", client_secret_store_id)
        if client_secret_store_url and not isinstance(client_secret_store_url, str):
            raise TypeError("Expected argument 'client_secret_store_url' to be a str")
        pulumi.set(__self__, "client_secret_store_url", client_secret_store_url)
        if cloud_witness_container_name and not isinstance(cloud_witness_container_name, str):
            raise TypeError("Expected argument 'cloud_witness_container_name' to be a str")
        pulumi.set(__self__, "cloud_witness_container_name", cloud_witness_container_name)
        if cloud_witness_storage_account_name and not isinstance(cloud_witness_storage_account_name, str):
            raise TypeError("Expected argument 'cloud_witness_storage_account_name' to be a str")
        pulumi.set(__self__, "cloud_witness_storage_account_name", cloud_witness_storage_account_name)
        if cloud_witness_storage_endpoint and not isinstance(cloud_witness_storage_endpoint, str):
            raise TypeError("Expected argument 'cloud_witness_storage_endpoint' to be a str")
        pulumi.set(__self__, "cloud_witness_storage_endpoint", cloud_witness_storage_endpoint)
        if cluster_witness_type and not isinstance(cluster_witness_type, str):
            raise TypeError("Expected argument 'cluster_witness_type' to be a str")
        pulumi.set(__self__, "cluster_witness_type", cluster_witness_type)
        if device_secrets and not isinstance(device_secrets, dict):
            raise TypeError("Expected argument 'device_secrets' to be a dict")
        pulumi.set(__self__, "device_secrets", device_secrets)
        if encryption_key and not isinstance(encryption_key, str):
            raise TypeError("Expected argument 'encryption_key' to be a str")
        pulumi.set(__self__, "encryption_key", encryption_key)
        if encryption_key_thumbprint and not isinstance(encryption_key_thumbprint, str):
            raise TypeError("Expected argument 'encryption_key_thumbprint' to be a str")
        pulumi.set(__self__, "encryption_key_thumbprint", encryption_key_thumbprint)
        if file_share_witness_location and not isinstance(file_share_witness_location, str):
            raise TypeError("Expected argument 'file_share_witness_location' to be a str")
        pulumi.set(__self__, "file_share_witness_location", file_share_witness_location)
        if file_share_witness_username and not isinstance(file_share_witness_username, str):
            raise TypeError("Expected argument 'file_share_witness_username' to be a str")
        pulumi.set(__self__, "file_share_witness_username", file_share_witness_username)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_vault_sync_status and not isinstance(key_vault_sync_status, str):
            raise TypeError("Expected argument 'key_vault_sync_status' to be a str")
        pulumi.set(__self__, "key_vault_sync_status", key_vault_sync_status)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_key and not isinstance(resource_key, str):
            raise TypeError("Expected argument 'resource_key' to be a str")
        pulumi.set(__self__, "resource_key", resource_key)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="channelIntegrityKeyName")
    def channel_integrity_key_name(self) -> Optional[builtins.str]:
        """
        The name of Channel Integrity Key stored in the Client Key Vault
        """
        return pulumi.get(self, "channel_integrity_key_name")

    @property
    @pulumi.getter(name="channelIntegrityKeyVersion")
    def channel_integrity_key_version(self) -> Optional[builtins.str]:
        """
        The version of Channel Integrity Key stored in the Client Key Vault
        """
        return pulumi.get(self, "channel_integrity_key_version")

    @property
    @pulumi.getter(name="clientSecretStoreId")
    def client_secret_store_id(self) -> Optional[builtins.str]:
        """
        The Key Vault ARM Id for client secrets
        """
        return pulumi.get(self, "client_secret_store_id")

    @property
    @pulumi.getter(name="clientSecretStoreUrl")
    def client_secret_store_url(self) -> Optional[builtins.str]:
        """
        The url to access the Client Key Vault
        """
        return pulumi.get(self, "client_secret_store_url")

    @property
    @pulumi.getter(name="cloudWitnessContainerName")
    def cloud_witness_container_name(self) -> builtins.str:
        """
        The Container for cloud witness in the storage account.
        """
        return pulumi.get(self, "cloud_witness_container_name")

    @property
    @pulumi.getter(name="cloudWitnessStorageAccountName")
    def cloud_witness_storage_account_name(self) -> builtins.str:
        """
        The Cloud Witness Storage account name.
        """
        return pulumi.get(self, "cloud_witness_storage_account_name")

    @property
    @pulumi.getter(name="cloudWitnessStorageEndpoint")
    def cloud_witness_storage_endpoint(self) -> builtins.str:
        """
        The Azure service endpoint of the cloud witness storage account.
        """
        return pulumi.get(self, "cloud_witness_storage_endpoint")

    @property
    @pulumi.getter(name="clusterWitnessType")
    def cluster_witness_type(self) -> builtins.str:
        """
        Cluster Witness Type
        """
        return pulumi.get(self, "cluster_witness_type")

    @property
    @pulumi.getter(name="deviceSecrets")
    def device_secrets(self) -> Mapping[str, 'outputs.SecretResponse']:
        """
        Device secrets, will be returned only with ODataFilter $expand=deviceSecrets
        """
        return pulumi.get(self, "device_secrets")

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional[builtins.str]:
        """
        The public part of the encryption certificate. Client uses this to encrypt any secret.
        """
        return pulumi.get(self, "encryption_key")

    @property
    @pulumi.getter(name="encryptionKeyThumbprint")
    def encryption_key_thumbprint(self) -> Optional[builtins.str]:
        """
        The digital signature of encrypted certificate.
        """
        return pulumi.get(self, "encryption_key_thumbprint")

    @property
    @pulumi.getter(name="fileShareWitnessLocation")
    def file_share_witness_location(self) -> builtins.str:
        """
        The witness location of file share.
        """
        return pulumi.get(self, "file_share_witness_location")

    @property
    @pulumi.getter(name="fileShareWitnessUsername")
    def file_share_witness_username(self) -> builtins.str:
        """
        The username of file share.
        """
        return pulumi.get(self, "file_share_witness_username")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The path ID that uniquely identifies the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyVaultSyncStatus")
    def key_vault_sync_status(self) -> Optional[builtins.str]:
        """
        Key vault sync status
        """
        return pulumi.get(self, "key_vault_sync_status")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The object name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceKey")
    def resource_key(self) -> builtins.str:
        """
        The Resource ID of the Resource.
        """
        return pulumi.get(self, "resource_key")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Metadata pertaining to creation and last modification of DataBoxEdgeDevice
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")


class AwaitableGetDeviceExtendedInformationResult(GetDeviceExtendedInformationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceExtendedInformationResult(
            channel_integrity_key_name=self.channel_integrity_key_name,
            channel_integrity_key_version=self.channel_integrity_key_version,
            client_secret_store_id=self.client_secret_store_id,
            client_secret_store_url=self.client_secret_store_url,
            cloud_witness_container_name=self.cloud_witness_container_name,
            cloud_witness_storage_account_name=self.cloud_witness_storage_account_name,
            cloud_witness_storage_endpoint=self.cloud_witness_storage_endpoint,
            cluster_witness_type=self.cluster_witness_type,
            device_secrets=self.device_secrets,
            encryption_key=self.encryption_key,
            encryption_key_thumbprint=self.encryption_key_thumbprint,
            file_share_witness_location=self.file_share_witness_location,
            file_share_witness_username=self.file_share_witness_username,
            id=self.id,
            key_vault_sync_status=self.key_vault_sync_status,
            name=self.name,
            resource_key=self.resource_key,
            system_data=self.system_data,
            type=self.type)


def get_device_extended_information(device_name: Optional[builtins.str] = None,
                                    resource_group_name: Optional[builtins.str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceExtendedInformationResult:
    """
    Gets additional information for the specified Azure Stack Edge/Data Box Gateway device.

    Uses Azure REST API version 2023-07-01.

    Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str device_name: The device name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:databoxedge:getDeviceExtendedInformation', __args__, opts=opts, typ=GetDeviceExtendedInformationResult).value

    return AwaitableGetDeviceExtendedInformationResult(
        channel_integrity_key_name=pulumi.get(__ret__, 'channel_integrity_key_name'),
        channel_integrity_key_version=pulumi.get(__ret__, 'channel_integrity_key_version'),
        client_secret_store_id=pulumi.get(__ret__, 'client_secret_store_id'),
        client_secret_store_url=pulumi.get(__ret__, 'client_secret_store_url'),
        cloud_witness_container_name=pulumi.get(__ret__, 'cloud_witness_container_name'),
        cloud_witness_storage_account_name=pulumi.get(__ret__, 'cloud_witness_storage_account_name'),
        cloud_witness_storage_endpoint=pulumi.get(__ret__, 'cloud_witness_storage_endpoint'),
        cluster_witness_type=pulumi.get(__ret__, 'cluster_witness_type'),
        device_secrets=pulumi.get(__ret__, 'device_secrets'),
        encryption_key=pulumi.get(__ret__, 'encryption_key'),
        encryption_key_thumbprint=pulumi.get(__ret__, 'encryption_key_thumbprint'),
        file_share_witness_location=pulumi.get(__ret__, 'file_share_witness_location'),
        file_share_witness_username=pulumi.get(__ret__, 'file_share_witness_username'),
        id=pulumi.get(__ret__, 'id'),
        key_vault_sync_status=pulumi.get(__ret__, 'key_vault_sync_status'),
        name=pulumi.get(__ret__, 'name'),
        resource_key=pulumi.get(__ret__, 'resource_key'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_device_extended_information_output(device_name: Optional[pulumi.Input[builtins.str]] = None,
                                           resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDeviceExtendedInformationResult]:
    """
    Gets additional information for the specified Azure Stack Edge/Data Box Gateway device.

    Uses Azure REST API version 2023-07-01.

    Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str device_name: The device name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:databoxedge:getDeviceExtendedInformation', __args__, opts=opts, typ=GetDeviceExtendedInformationResult)
    return __ret__.apply(lambda __response__: GetDeviceExtendedInformationResult(
        channel_integrity_key_name=pulumi.get(__response__, 'channel_integrity_key_name'),
        channel_integrity_key_version=pulumi.get(__response__, 'channel_integrity_key_version'),
        client_secret_store_id=pulumi.get(__response__, 'client_secret_store_id'),
        client_secret_store_url=pulumi.get(__response__, 'client_secret_store_url'),
        cloud_witness_container_name=pulumi.get(__response__, 'cloud_witness_container_name'),
        cloud_witness_storage_account_name=pulumi.get(__response__, 'cloud_witness_storage_account_name'),
        cloud_witness_storage_endpoint=pulumi.get(__response__, 'cloud_witness_storage_endpoint'),
        cluster_witness_type=pulumi.get(__response__, 'cluster_witness_type'),
        device_secrets=pulumi.get(__response__, 'device_secrets'),
        encryption_key=pulumi.get(__response__, 'encryption_key'),
        encryption_key_thumbprint=pulumi.get(__response__, 'encryption_key_thumbprint'),
        file_share_witness_location=pulumi.get(__response__, 'file_share_witness_location'),
        file_share_witness_username=pulumi.get(__response__, 'file_share_witness_username'),
        id=pulumi.get(__response__, 'id'),
        key_vault_sync_status=pulumi.get(__response__, 'key_vault_sync_status'),
        name=pulumi.get(__response__, 'name'),
        resource_key=pulumi.get(__response__, 'resource_key'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
