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
    'GetStorageTargetResult',
    'AwaitableGetStorageTargetResult',
    'get_storage_target',
    'get_storage_target_output',
]

@pulumi.output_type
class GetStorageTargetResult:
    """
    Type of the Storage Target.
    """
    def __init__(__self__, allocation_percentage=None, azure_api_version=None, blob_nfs=None, clfs=None, id=None, junctions=None, location=None, name=None, nfs3=None, provisioning_state=None, state=None, system_data=None, target_type=None, type=None, unknown=None):
        if allocation_percentage and not isinstance(allocation_percentage, int):
            raise TypeError("Expected argument 'allocation_percentage' to be a int")
        pulumi.set(__self__, "allocation_percentage", allocation_percentage)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if blob_nfs and not isinstance(blob_nfs, dict):
            raise TypeError("Expected argument 'blob_nfs' to be a dict")
        pulumi.set(__self__, "blob_nfs", blob_nfs)
        if clfs and not isinstance(clfs, dict):
            raise TypeError("Expected argument 'clfs' to be a dict")
        pulumi.set(__self__, "clfs", clfs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if junctions and not isinstance(junctions, list):
            raise TypeError("Expected argument 'junctions' to be a list")
        pulumi.set(__self__, "junctions", junctions)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nfs3 and not isinstance(nfs3, dict):
            raise TypeError("Expected argument 'nfs3' to be a dict")
        pulumi.set(__self__, "nfs3", nfs3)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if target_type and not isinstance(target_type, str):
            raise TypeError("Expected argument 'target_type' to be a str")
        pulumi.set(__self__, "target_type", target_type)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if unknown and not isinstance(unknown, dict):
            raise TypeError("Expected argument 'unknown' to be a dict")
        pulumi.set(__self__, "unknown", unknown)

    @property
    @pulumi.getter(name="allocationPercentage")
    def allocation_percentage(self) -> builtins.int:
        """
        The percentage of cache space allocated for this storage target
        """
        return pulumi.get(self, "allocation_percentage")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="blobNfs")
    def blob_nfs(self) -> Optional['outputs.BlobNfsTargetResponse']:
        """
        Properties when targetType is blobNfs.
        """
        return pulumi.get(self, "blob_nfs")

    @property
    @pulumi.getter
    def clfs(self) -> Optional['outputs.ClfsTargetResponse']:
        """
        Properties when targetType is clfs.
        """
        return pulumi.get(self, "clfs")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID of the Storage Target.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def junctions(self) -> Optional[Sequence['outputs.NamespaceJunctionResponse']]:
        """
        List of cache namespace junctions to target for namespace associations.
        """
        return pulumi.get(self, "junctions")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Region name string.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the Storage Target.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nfs3(self) -> Optional['outputs.Nfs3TargetResponse']:
        """
        Properties when targetType is nfs3.
        """
        return pulumi.get(self, "nfs3")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def state(self) -> Optional[builtins.str]:
        """
        Storage target operational state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system meta data relating to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> builtins.str:
        """
        Type of the Storage Target.
        """
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def unknown(self) -> Optional['outputs.UnknownTargetResponse']:
        """
        Properties when targetType is unknown.
        """
        return pulumi.get(self, "unknown")


class AwaitableGetStorageTargetResult(GetStorageTargetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageTargetResult(
            allocation_percentage=self.allocation_percentage,
            azure_api_version=self.azure_api_version,
            blob_nfs=self.blob_nfs,
            clfs=self.clfs,
            id=self.id,
            junctions=self.junctions,
            location=self.location,
            name=self.name,
            nfs3=self.nfs3,
            provisioning_state=self.provisioning_state,
            state=self.state,
            system_data=self.system_data,
            target_type=self.target_type,
            type=self.type,
            unknown=self.unknown)


def get_storage_target(cache_name: Optional[builtins.str] = None,
                       resource_group_name: Optional[builtins.str] = None,
                       storage_target_name: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageTargetResult:
    """
    Returns a Storage Target from a cache.

    Uses Azure REST API version 2024-03-01.

    Other available API versions: 2023-05-01, 2023-11-01-preview, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagecache [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str cache_name: Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str storage_target_name: Name of Storage Target.
    """
    __args__ = dict()
    __args__['cacheName'] = cache_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['storageTargetName'] = storage_target_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:storagecache:getStorageTarget', __args__, opts=opts, typ=GetStorageTargetResult).value

    return AwaitableGetStorageTargetResult(
        allocation_percentage=pulumi.get(__ret__, 'allocation_percentage'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        blob_nfs=pulumi.get(__ret__, 'blob_nfs'),
        clfs=pulumi.get(__ret__, 'clfs'),
        id=pulumi.get(__ret__, 'id'),
        junctions=pulumi.get(__ret__, 'junctions'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        nfs3=pulumi.get(__ret__, 'nfs3'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        state=pulumi.get(__ret__, 'state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        target_type=pulumi.get(__ret__, 'target_type'),
        type=pulumi.get(__ret__, 'type'),
        unknown=pulumi.get(__ret__, 'unknown'))
def get_storage_target_output(cache_name: Optional[pulumi.Input[builtins.str]] = None,
                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                              storage_target_name: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStorageTargetResult]:
    """
    Returns a Storage Target from a cache.

    Uses Azure REST API version 2024-03-01.

    Other available API versions: 2023-05-01, 2023-11-01-preview, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagecache [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str cache_name: Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str storage_target_name: Name of Storage Target.
    """
    __args__ = dict()
    __args__['cacheName'] = cache_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['storageTargetName'] = storage_target_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:storagecache:getStorageTarget', __args__, opts=opts, typ=GetStorageTargetResult)
    return __ret__.apply(lambda __response__: GetStorageTargetResult(
        allocation_percentage=pulumi.get(__response__, 'allocation_percentage'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        blob_nfs=pulumi.get(__response__, 'blob_nfs'),
        clfs=pulumi.get(__response__, 'clfs'),
        id=pulumi.get(__response__, 'id'),
        junctions=pulumi.get(__response__, 'junctions'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        nfs3=pulumi.get(__response__, 'nfs3'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        state=pulumi.get(__response__, 'state'),
        system_data=pulumi.get(__response__, 'system_data'),
        target_type=pulumi.get(__response__, 'target_type'),
        type=pulumi.get(__response__, 'type'),
        unknown=pulumi.get(__response__, 'unknown')))
