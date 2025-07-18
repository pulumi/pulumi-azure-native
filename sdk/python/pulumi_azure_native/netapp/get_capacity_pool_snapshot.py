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
    'GetCapacityPoolSnapshotResult',
    'AwaitableGetCapacityPoolSnapshotResult',
    'get_capacity_pool_snapshot',
    'get_capacity_pool_snapshot_output',
]

@pulumi.output_type
class GetCapacityPoolSnapshotResult:
    """
    Snapshot of a Volume
    """
    def __init__(__self__, azure_api_version=None, created=None, id=None, location=None, name=None, provisioning_state=None, snapshot_id=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
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
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
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
    def created(self) -> builtins.str:
        """
        The creation date of the snapshot
        """
        return pulumi.get(self, "created")

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
        Resource location
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
        Azure lifecycle management
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> builtins.str:
        """
        UUID v4 used to identify the Snapshot
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetCapacityPoolSnapshotResult(GetCapacityPoolSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCapacityPoolSnapshotResult(
            azure_api_version=self.azure_api_version,
            created=self.created,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            snapshot_id=self.snapshot_id,
            system_data=self.system_data,
            type=self.type)


def get_capacity_pool_snapshot(account_name: Optional[builtins.str] = None,
                               pool_name: Optional[builtins.str] = None,
                               resource_group_name: Optional[builtins.str] = None,
                               snapshot_name: Optional[builtins.str] = None,
                               volume_name: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCapacityPoolSnapshotResult:
    """
    Get details of the specified snapshot

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2022-11-01, 2022-11-01-preview, 2023-05-01, 2023-05-01-preview, 2023-07-01, 2023-07-01-preview, 2023-11-01, 2023-11-01-preview, 2024-01-01, 2024-03-01, 2024-03-01-preview, 2024-05-01, 2024-05-01-preview, 2024-07-01, 2024-07-01-preview, 2024-09-01-preview, 2025-01-01, 2025-01-01-preview, 2025-03-01, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native netapp [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: The name of the NetApp account
    :param builtins.str pool_name: The name of the capacity pool
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str snapshot_name: The name of the snapshot
    :param builtins.str volume_name: The name of the volume
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['poolName'] = pool_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['snapshotName'] = snapshot_name
    __args__['volumeName'] = volume_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:netapp:getCapacityPoolSnapshot', __args__, opts=opts, typ=GetCapacityPoolSnapshotResult).value

    return AwaitableGetCapacityPoolSnapshotResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        created=pulumi.get(__ret__, 'created'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_capacity_pool_snapshot_output(account_name: Optional[pulumi.Input[builtins.str]] = None,
                                      pool_name: Optional[pulumi.Input[builtins.str]] = None,
                                      resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                      snapshot_name: Optional[pulumi.Input[builtins.str]] = None,
                                      volume_name: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCapacityPoolSnapshotResult]:
    """
    Get details of the specified snapshot

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2022-11-01, 2022-11-01-preview, 2023-05-01, 2023-05-01-preview, 2023-07-01, 2023-07-01-preview, 2023-11-01, 2023-11-01-preview, 2024-01-01, 2024-03-01, 2024-03-01-preview, 2024-05-01, 2024-05-01-preview, 2024-07-01, 2024-07-01-preview, 2024-09-01-preview, 2025-01-01, 2025-01-01-preview, 2025-03-01, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native netapp [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: The name of the NetApp account
    :param builtins.str pool_name: The name of the capacity pool
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str snapshot_name: The name of the snapshot
    :param builtins.str volume_name: The name of the volume
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['poolName'] = pool_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['snapshotName'] = snapshot_name
    __args__['volumeName'] = volume_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:netapp:getCapacityPoolSnapshot', __args__, opts=opts, typ=GetCapacityPoolSnapshotResult)
    return __ret__.apply(lambda __response__: GetCapacityPoolSnapshotResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        created=pulumi.get(__response__, 'created'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        snapshot_id=pulumi.get(__response__, 'snapshot_id'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
