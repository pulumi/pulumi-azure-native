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
    'GetFileShareResult',
    'AwaitableGetFileShareResult',
    'get_file_share',
    'get_file_share_output',
]

@pulumi.output_type
class GetFileShareResult:
    """
    Properties of the file share, including Id, resource name, resource type, Etag.
    """
    def __init__(__self__, access_tier=None, access_tier_change_time=None, access_tier_status=None, azure_api_version=None, deleted=None, deleted_time=None, enabled_protocols=None, etag=None, file_share_paid_bursting=None, id=None, included_burst_iops=None, last_modified_time=None, lease_duration=None, lease_state=None, lease_status=None, max_burst_credits_for_iops=None, metadata=None, name=None, next_allowed_provisioned_bandwidth_downgrade_time=None, next_allowed_provisioned_iops_downgrade_time=None, next_allowed_quota_downgrade_time=None, provisioned_bandwidth_mibps=None, provisioned_iops=None, remaining_retention_days=None, root_squash=None, share_quota=None, share_usage_bytes=None, signed_identifiers=None, snapshot_time=None, type=None, version=None):
        if access_tier and not isinstance(access_tier, str):
            raise TypeError("Expected argument 'access_tier' to be a str")
        pulumi.set(__self__, "access_tier", access_tier)
        if access_tier_change_time and not isinstance(access_tier_change_time, str):
            raise TypeError("Expected argument 'access_tier_change_time' to be a str")
        pulumi.set(__self__, "access_tier_change_time", access_tier_change_time)
        if access_tier_status and not isinstance(access_tier_status, str):
            raise TypeError("Expected argument 'access_tier_status' to be a str")
        pulumi.set(__self__, "access_tier_status", access_tier_status)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if deleted and not isinstance(deleted, bool):
            raise TypeError("Expected argument 'deleted' to be a bool")
        pulumi.set(__self__, "deleted", deleted)
        if deleted_time and not isinstance(deleted_time, str):
            raise TypeError("Expected argument 'deleted_time' to be a str")
        pulumi.set(__self__, "deleted_time", deleted_time)
        if enabled_protocols and not isinstance(enabled_protocols, str):
            raise TypeError("Expected argument 'enabled_protocols' to be a str")
        pulumi.set(__self__, "enabled_protocols", enabled_protocols)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if file_share_paid_bursting and not isinstance(file_share_paid_bursting, dict):
            raise TypeError("Expected argument 'file_share_paid_bursting' to be a dict")
        pulumi.set(__self__, "file_share_paid_bursting", file_share_paid_bursting)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if included_burst_iops and not isinstance(included_burst_iops, int):
            raise TypeError("Expected argument 'included_burst_iops' to be a int")
        pulumi.set(__self__, "included_burst_iops", included_burst_iops)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if lease_duration and not isinstance(lease_duration, str):
            raise TypeError("Expected argument 'lease_duration' to be a str")
        pulumi.set(__self__, "lease_duration", lease_duration)
        if lease_state and not isinstance(lease_state, str):
            raise TypeError("Expected argument 'lease_state' to be a str")
        pulumi.set(__self__, "lease_state", lease_state)
        if lease_status and not isinstance(lease_status, str):
            raise TypeError("Expected argument 'lease_status' to be a str")
        pulumi.set(__self__, "lease_status", lease_status)
        if max_burst_credits_for_iops and not isinstance(max_burst_credits_for_iops, float):
            raise TypeError("Expected argument 'max_burst_credits_for_iops' to be a float")
        pulumi.set(__self__, "max_burst_credits_for_iops", max_burst_credits_for_iops)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if next_allowed_provisioned_bandwidth_downgrade_time and not isinstance(next_allowed_provisioned_bandwidth_downgrade_time, str):
            raise TypeError("Expected argument 'next_allowed_provisioned_bandwidth_downgrade_time' to be a str")
        pulumi.set(__self__, "next_allowed_provisioned_bandwidth_downgrade_time", next_allowed_provisioned_bandwidth_downgrade_time)
        if next_allowed_provisioned_iops_downgrade_time and not isinstance(next_allowed_provisioned_iops_downgrade_time, str):
            raise TypeError("Expected argument 'next_allowed_provisioned_iops_downgrade_time' to be a str")
        pulumi.set(__self__, "next_allowed_provisioned_iops_downgrade_time", next_allowed_provisioned_iops_downgrade_time)
        if next_allowed_quota_downgrade_time and not isinstance(next_allowed_quota_downgrade_time, str):
            raise TypeError("Expected argument 'next_allowed_quota_downgrade_time' to be a str")
        pulumi.set(__self__, "next_allowed_quota_downgrade_time", next_allowed_quota_downgrade_time)
        if provisioned_bandwidth_mibps and not isinstance(provisioned_bandwidth_mibps, int):
            raise TypeError("Expected argument 'provisioned_bandwidth_mibps' to be a int")
        pulumi.set(__self__, "provisioned_bandwidth_mibps", provisioned_bandwidth_mibps)
        if provisioned_iops and not isinstance(provisioned_iops, int):
            raise TypeError("Expected argument 'provisioned_iops' to be a int")
        pulumi.set(__self__, "provisioned_iops", provisioned_iops)
        if remaining_retention_days and not isinstance(remaining_retention_days, int):
            raise TypeError("Expected argument 'remaining_retention_days' to be a int")
        pulumi.set(__self__, "remaining_retention_days", remaining_retention_days)
        if root_squash and not isinstance(root_squash, str):
            raise TypeError("Expected argument 'root_squash' to be a str")
        pulumi.set(__self__, "root_squash", root_squash)
        if share_quota and not isinstance(share_quota, int):
            raise TypeError("Expected argument 'share_quota' to be a int")
        pulumi.set(__self__, "share_quota", share_quota)
        if share_usage_bytes and not isinstance(share_usage_bytes, float):
            raise TypeError("Expected argument 'share_usage_bytes' to be a float")
        pulumi.set(__self__, "share_usage_bytes", share_usage_bytes)
        if signed_identifiers and not isinstance(signed_identifiers, list):
            raise TypeError("Expected argument 'signed_identifiers' to be a list")
        pulumi.set(__self__, "signed_identifiers", signed_identifiers)
        if snapshot_time and not isinstance(snapshot_time, str):
            raise TypeError("Expected argument 'snapshot_time' to be a str")
        pulumi.set(__self__, "snapshot_time", snapshot_time)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="accessTier")
    def access_tier(self) -> Optional[builtins.str]:
        """
        Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
        """
        return pulumi.get(self, "access_tier")

    @property
    @pulumi.getter(name="accessTierChangeTime")
    def access_tier_change_time(self) -> builtins.str:
        """
        Indicates the last modification time for share access tier.
        """
        return pulumi.get(self, "access_tier_change_time")

    @property
    @pulumi.getter(name="accessTierStatus")
    def access_tier_status(self) -> builtins.str:
        """
        Indicates if there is a pending transition for access tier.
        """
        return pulumi.get(self, "access_tier_status")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def deleted(self) -> builtins.bool:
        """
        Indicates whether the share was deleted.
        """
        return pulumi.get(self, "deleted")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> builtins.str:
        """
        The deleted time if the share was deleted.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter(name="enabledProtocols")
    def enabled_protocols(self) -> Optional[builtins.str]:
        """
        The authentication protocol that is used for the file share. Can only be specified when creating a share.
        """
        return pulumi.get(self, "enabled_protocols")

    @property
    @pulumi.getter
    def etag(self) -> builtins.str:
        """
        Resource Etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="fileSharePaidBursting")
    def file_share_paid_bursting(self) -> Optional['outputs.FileSharePropertiesResponseFileSharePaidBursting']:
        """
        File Share Paid Bursting properties.
        """
        return pulumi.get(self, "file_share_paid_bursting")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includedBurstIops")
    def included_burst_iops(self) -> builtins.int:
        """
        The calculated burst IOPS of the share. This property is only for file shares created under Files Provisioned v2 account type.
        """
        return pulumi.get(self, "included_burst_iops")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> builtins.str:
        """
        Returns the date and time the share was last modified.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> builtins.str:
        """
        Specifies whether the lease on a share is of infinite or fixed duration, only when the share is leased.
        """
        return pulumi.get(self, "lease_duration")

    @property
    @pulumi.getter(name="leaseState")
    def lease_state(self) -> builtins.str:
        """
        Lease state of the share.
        """
        return pulumi.get(self, "lease_state")

    @property
    @pulumi.getter(name="leaseStatus")
    def lease_status(self) -> builtins.str:
        """
        The lease status of the share.
        """
        return pulumi.get(self, "lease_status")

    @property
    @pulumi.getter(name="maxBurstCreditsForIops")
    def max_burst_credits_for_iops(self) -> builtins.float:
        """
        The calculated maximum burst credits for the share. This property is only for file shares created under Files Provisioned v2 account type.
        """
        return pulumi.get(self, "max_burst_credits_for_iops")

    @property
    @pulumi.getter
    def metadata(self) -> Optional[Mapping[str, builtins.str]]:
        """
        A name-value pair to associate with the share as metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextAllowedProvisionedBandwidthDowngradeTime")
    def next_allowed_provisioned_bandwidth_downgrade_time(self) -> builtins.str:
        """
        Returns the next allowed provisioned bandwidth downgrade time for the share. This property is only for file shares created under Files Provisioned v2 account type.
        """
        return pulumi.get(self, "next_allowed_provisioned_bandwidth_downgrade_time")

    @property
    @pulumi.getter(name="nextAllowedProvisionedIopsDowngradeTime")
    def next_allowed_provisioned_iops_downgrade_time(self) -> builtins.str:
        """
        Returns the next allowed provisioned IOPS downgrade time for the share. This property is only for file shares created under Files Provisioned v2 account type.
        """
        return pulumi.get(self, "next_allowed_provisioned_iops_downgrade_time")

    @property
    @pulumi.getter(name="nextAllowedQuotaDowngradeTime")
    def next_allowed_quota_downgrade_time(self) -> builtins.str:
        """
        Returns the next allowed provisioned storage size downgrade time for the share. This property is only for file shares created under Files Provisioned v1 SSD and Files Provisioned v2 account type
        """
        return pulumi.get(self, "next_allowed_quota_downgrade_time")

    @property
    @pulumi.getter(name="provisionedBandwidthMibps")
    def provisioned_bandwidth_mibps(self) -> Optional[builtins.int]:
        """
        The provisioned bandwidth of the share, in mebibytes per second. This property is only for file shares created under Files Provisioned v2 account type. Please refer to the GetFileServiceUsage API response for the minimum and maximum allowed value for provisioned bandwidth.
        """
        return pulumi.get(self, "provisioned_bandwidth_mibps")

    @property
    @pulumi.getter(name="provisionedIops")
    def provisioned_iops(self) -> Optional[builtins.int]:
        """
        The provisioned IOPS of the share. This property is only for file shares created under Files Provisioned v2 account type. Please refer to the GetFileServiceUsage API response for the minimum and maximum allowed value for provisioned IOPS.
        """
        return pulumi.get(self, "provisioned_iops")

    @property
    @pulumi.getter(name="remainingRetentionDays")
    def remaining_retention_days(self) -> builtins.int:
        """
        Remaining retention days for share that was soft deleted.
        """
        return pulumi.get(self, "remaining_retention_days")

    @property
    @pulumi.getter(name="rootSquash")
    def root_squash(self) -> Optional[builtins.str]:
        """
        The property is for NFS share only. The default is NoRootSquash.
        """
        return pulumi.get(self, "root_squash")

    @property
    @pulumi.getter(name="shareQuota")
    def share_quota(self) -> Optional[builtins.int]:
        """
        The provisioned size of the share, in gibibytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400. For file shares created under Files Provisioned v2 account type, please refer to the GetFileServiceUsage API response for the minimum and maximum allowed provisioned storage size.
        """
        return pulumi.get(self, "share_quota")

    @property
    @pulumi.getter(name="shareUsageBytes")
    def share_usage_bytes(self) -> builtins.float:
        """
        The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.
        """
        return pulumi.get(self, "share_usage_bytes")

    @property
    @pulumi.getter(name="signedIdentifiers")
    def signed_identifiers(self) -> Optional[Sequence['outputs.SignedIdentifierResponse']]:
        """
        List of stored access policies specified on the share.
        """
        return pulumi.get(self, "signed_identifiers")

    @property
    @pulumi.getter(name="snapshotTime")
    def snapshot_time(self) -> builtins.str:
        """
        Creation time of share snapshot returned in the response of list shares with expand param "snapshots".
        """
        return pulumi.get(self, "snapshot_time")

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
        The version of the share.
        """
        return pulumi.get(self, "version")


class AwaitableGetFileShareResult(GetFileShareResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileShareResult(
            access_tier=self.access_tier,
            access_tier_change_time=self.access_tier_change_time,
            access_tier_status=self.access_tier_status,
            azure_api_version=self.azure_api_version,
            deleted=self.deleted,
            deleted_time=self.deleted_time,
            enabled_protocols=self.enabled_protocols,
            etag=self.etag,
            file_share_paid_bursting=self.file_share_paid_bursting,
            id=self.id,
            included_burst_iops=self.included_burst_iops,
            last_modified_time=self.last_modified_time,
            lease_duration=self.lease_duration,
            lease_state=self.lease_state,
            lease_status=self.lease_status,
            max_burst_credits_for_iops=self.max_burst_credits_for_iops,
            metadata=self.metadata,
            name=self.name,
            next_allowed_provisioned_bandwidth_downgrade_time=self.next_allowed_provisioned_bandwidth_downgrade_time,
            next_allowed_provisioned_iops_downgrade_time=self.next_allowed_provisioned_iops_downgrade_time,
            next_allowed_quota_downgrade_time=self.next_allowed_quota_downgrade_time,
            provisioned_bandwidth_mibps=self.provisioned_bandwidth_mibps,
            provisioned_iops=self.provisioned_iops,
            remaining_retention_days=self.remaining_retention_days,
            root_squash=self.root_squash,
            share_quota=self.share_quota,
            share_usage_bytes=self.share_usage_bytes,
            signed_identifiers=self.signed_identifiers,
            snapshot_time=self.snapshot_time,
            type=self.type,
            version=self.version)


def get_file_share(account_name: Optional[builtins.str] = None,
                   expand: Optional[builtins.str] = None,
                   resource_group_name: Optional[builtins.str] = None,
                   share_name: Optional[builtins.str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileShareResult:
    """
    Gets properties of a specified share.

    Uses Azure REST API version 2024-01-01.

    Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
    :param builtins.str expand: Optional, used to expand the properties within share's properties. Valid values are: stats. Should be passed as a string with delimiter ','.
    :param builtins.str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    :param builtins.str share_name: The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['expand'] = expand
    __args__['resourceGroupName'] = resource_group_name
    __args__['shareName'] = share_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:storage:getFileShare', __args__, opts=opts, typ=GetFileShareResult).value

    return AwaitableGetFileShareResult(
        access_tier=pulumi.get(__ret__, 'access_tier'),
        access_tier_change_time=pulumi.get(__ret__, 'access_tier_change_time'),
        access_tier_status=pulumi.get(__ret__, 'access_tier_status'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        deleted=pulumi.get(__ret__, 'deleted'),
        deleted_time=pulumi.get(__ret__, 'deleted_time'),
        enabled_protocols=pulumi.get(__ret__, 'enabled_protocols'),
        etag=pulumi.get(__ret__, 'etag'),
        file_share_paid_bursting=pulumi.get(__ret__, 'file_share_paid_bursting'),
        id=pulumi.get(__ret__, 'id'),
        included_burst_iops=pulumi.get(__ret__, 'included_burst_iops'),
        last_modified_time=pulumi.get(__ret__, 'last_modified_time'),
        lease_duration=pulumi.get(__ret__, 'lease_duration'),
        lease_state=pulumi.get(__ret__, 'lease_state'),
        lease_status=pulumi.get(__ret__, 'lease_status'),
        max_burst_credits_for_iops=pulumi.get(__ret__, 'max_burst_credits_for_iops'),
        metadata=pulumi.get(__ret__, 'metadata'),
        name=pulumi.get(__ret__, 'name'),
        next_allowed_provisioned_bandwidth_downgrade_time=pulumi.get(__ret__, 'next_allowed_provisioned_bandwidth_downgrade_time'),
        next_allowed_provisioned_iops_downgrade_time=pulumi.get(__ret__, 'next_allowed_provisioned_iops_downgrade_time'),
        next_allowed_quota_downgrade_time=pulumi.get(__ret__, 'next_allowed_quota_downgrade_time'),
        provisioned_bandwidth_mibps=pulumi.get(__ret__, 'provisioned_bandwidth_mibps'),
        provisioned_iops=pulumi.get(__ret__, 'provisioned_iops'),
        remaining_retention_days=pulumi.get(__ret__, 'remaining_retention_days'),
        root_squash=pulumi.get(__ret__, 'root_squash'),
        share_quota=pulumi.get(__ret__, 'share_quota'),
        share_usage_bytes=pulumi.get(__ret__, 'share_usage_bytes'),
        signed_identifiers=pulumi.get(__ret__, 'signed_identifiers'),
        snapshot_time=pulumi.get(__ret__, 'snapshot_time'),
        type=pulumi.get(__ret__, 'type'),
        version=pulumi.get(__ret__, 'version'))
def get_file_share_output(account_name: Optional[pulumi.Input[builtins.str]] = None,
                          expand: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                          resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                          share_name: Optional[pulumi.Input[builtins.str]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFileShareResult]:
    """
    Gets properties of a specified share.

    Uses Azure REST API version 2024-01-01.

    Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
    :param builtins.str expand: Optional, used to expand the properties within share's properties. Valid values are: stats. Should be passed as a string with delimiter ','.
    :param builtins.str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    :param builtins.str share_name: The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['expand'] = expand
    __args__['resourceGroupName'] = resource_group_name
    __args__['shareName'] = share_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:storage:getFileShare', __args__, opts=opts, typ=GetFileShareResult)
    return __ret__.apply(lambda __response__: GetFileShareResult(
        access_tier=pulumi.get(__response__, 'access_tier'),
        access_tier_change_time=pulumi.get(__response__, 'access_tier_change_time'),
        access_tier_status=pulumi.get(__response__, 'access_tier_status'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        deleted=pulumi.get(__response__, 'deleted'),
        deleted_time=pulumi.get(__response__, 'deleted_time'),
        enabled_protocols=pulumi.get(__response__, 'enabled_protocols'),
        etag=pulumi.get(__response__, 'etag'),
        file_share_paid_bursting=pulumi.get(__response__, 'file_share_paid_bursting'),
        id=pulumi.get(__response__, 'id'),
        included_burst_iops=pulumi.get(__response__, 'included_burst_iops'),
        last_modified_time=pulumi.get(__response__, 'last_modified_time'),
        lease_duration=pulumi.get(__response__, 'lease_duration'),
        lease_state=pulumi.get(__response__, 'lease_state'),
        lease_status=pulumi.get(__response__, 'lease_status'),
        max_burst_credits_for_iops=pulumi.get(__response__, 'max_burst_credits_for_iops'),
        metadata=pulumi.get(__response__, 'metadata'),
        name=pulumi.get(__response__, 'name'),
        next_allowed_provisioned_bandwidth_downgrade_time=pulumi.get(__response__, 'next_allowed_provisioned_bandwidth_downgrade_time'),
        next_allowed_provisioned_iops_downgrade_time=pulumi.get(__response__, 'next_allowed_provisioned_iops_downgrade_time'),
        next_allowed_quota_downgrade_time=pulumi.get(__response__, 'next_allowed_quota_downgrade_time'),
        provisioned_bandwidth_mibps=pulumi.get(__response__, 'provisioned_bandwidth_mibps'),
        provisioned_iops=pulumi.get(__response__, 'provisioned_iops'),
        remaining_retention_days=pulumi.get(__response__, 'remaining_retention_days'),
        root_squash=pulumi.get(__response__, 'root_squash'),
        share_quota=pulumi.get(__response__, 'share_quota'),
        share_usage_bytes=pulumi.get(__response__, 'share_usage_bytes'),
        signed_identifiers=pulumi.get(__response__, 'signed_identifiers'),
        snapshot_time=pulumi.get(__response__, 'snapshot_time'),
        type=pulumi.get(__response__, 'type'),
        version=pulumi.get(__response__, 'version')))
