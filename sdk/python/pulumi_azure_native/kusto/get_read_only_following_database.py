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
    'GetReadOnlyFollowingDatabaseResult',
    'AwaitableGetReadOnlyFollowingDatabaseResult',
    'get_read_only_following_database',
    'get_read_only_following_database_output',
]

@pulumi.output_type
class GetReadOnlyFollowingDatabaseResult:
    """
    Class representing a read only following database.
    """
    def __init__(__self__, attached_database_configuration_name=None, azure_api_version=None, database_share_origin=None, hot_cache_period=None, id=None, kind=None, leader_cluster_resource_id=None, location=None, name=None, original_database_name=None, principals_modification_kind=None, provisioning_state=None, soft_delete_period=None, statistics=None, suspension_details=None, table_level_sharing_properties=None, type=None):
        if attached_database_configuration_name and not isinstance(attached_database_configuration_name, str):
            raise TypeError("Expected argument 'attached_database_configuration_name' to be a str")
        pulumi.set(__self__, "attached_database_configuration_name", attached_database_configuration_name)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if database_share_origin and not isinstance(database_share_origin, str):
            raise TypeError("Expected argument 'database_share_origin' to be a str")
        pulumi.set(__self__, "database_share_origin", database_share_origin)
        if hot_cache_period and not isinstance(hot_cache_period, str):
            raise TypeError("Expected argument 'hot_cache_period' to be a str")
        pulumi.set(__self__, "hot_cache_period", hot_cache_period)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if leader_cluster_resource_id and not isinstance(leader_cluster_resource_id, str):
            raise TypeError("Expected argument 'leader_cluster_resource_id' to be a str")
        pulumi.set(__self__, "leader_cluster_resource_id", leader_cluster_resource_id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if original_database_name and not isinstance(original_database_name, str):
            raise TypeError("Expected argument 'original_database_name' to be a str")
        pulumi.set(__self__, "original_database_name", original_database_name)
        if principals_modification_kind and not isinstance(principals_modification_kind, str):
            raise TypeError("Expected argument 'principals_modification_kind' to be a str")
        pulumi.set(__self__, "principals_modification_kind", principals_modification_kind)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if soft_delete_period and not isinstance(soft_delete_period, str):
            raise TypeError("Expected argument 'soft_delete_period' to be a str")
        pulumi.set(__self__, "soft_delete_period", soft_delete_period)
        if statistics and not isinstance(statistics, dict):
            raise TypeError("Expected argument 'statistics' to be a dict")
        pulumi.set(__self__, "statistics", statistics)
        if suspension_details and not isinstance(suspension_details, dict):
            raise TypeError("Expected argument 'suspension_details' to be a dict")
        pulumi.set(__self__, "suspension_details", suspension_details)
        if table_level_sharing_properties and not isinstance(table_level_sharing_properties, dict):
            raise TypeError("Expected argument 'table_level_sharing_properties' to be a dict")
        pulumi.set(__self__, "table_level_sharing_properties", table_level_sharing_properties)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="attachedDatabaseConfigurationName")
    def attached_database_configuration_name(self) -> builtins.str:
        """
        The name of the attached database configuration cluster
        """
        return pulumi.get(self, "attached_database_configuration_name")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="databaseShareOrigin")
    def database_share_origin(self) -> builtins.str:
        """
        The origin of the following setup.
        """
        return pulumi.get(self, "database_share_origin")

    @property
    @pulumi.getter(name="hotCachePeriod")
    def hot_cache_period(self) -> Optional[builtins.str]:
        """
        The time the data should be kept in cache for fast queries in TimeSpan.
        """
        return pulumi.get(self, "hot_cache_period")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> builtins.str:
        """
        Kind of the database
        Expected value is 'ReadOnlyFollowing'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="leaderClusterResourceId")
    def leader_cluster_resource_id(self) -> builtins.str:
        """
        The name of the leader cluster
        """
        return pulumi.get(self, "leader_cluster_resource_id")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        Resource location.
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
    @pulumi.getter(name="originalDatabaseName")
    def original_database_name(self) -> builtins.str:
        """
        The original database name, before databaseNameOverride or databaseNamePrefix where applied.
        """
        return pulumi.get(self, "original_database_name")

    @property
    @pulumi.getter(name="principalsModificationKind")
    def principals_modification_kind(self) -> builtins.str:
        """
        The principals modification kind of the database
        """
        return pulumi.get(self, "principals_modification_kind")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioned state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="softDeletePeriod")
    def soft_delete_period(self) -> builtins.str:
        """
        The time the data should be kept before it stops being accessible to queries in TimeSpan.
        """
        return pulumi.get(self, "soft_delete_period")

    @property
    @pulumi.getter
    def statistics(self) -> 'outputs.DatabaseStatisticsResponse':
        """
        The statistics of the database.
        """
        return pulumi.get(self, "statistics")

    @property
    @pulumi.getter(name="suspensionDetails")
    def suspension_details(self) -> 'outputs.SuspensionDetailsResponse':
        """
        The database suspension details. If the database is suspended, this object contains information related to the database's suspension state.
        """
        return pulumi.get(self, "suspension_details")

    @property
    @pulumi.getter(name="tableLevelSharingProperties")
    def table_level_sharing_properties(self) -> 'outputs.TableLevelSharingPropertiesResponse':
        """
        Table level sharing specifications
        """
        return pulumi.get(self, "table_level_sharing_properties")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetReadOnlyFollowingDatabaseResult(GetReadOnlyFollowingDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReadOnlyFollowingDatabaseResult(
            attached_database_configuration_name=self.attached_database_configuration_name,
            azure_api_version=self.azure_api_version,
            database_share_origin=self.database_share_origin,
            hot_cache_period=self.hot_cache_period,
            id=self.id,
            kind=self.kind,
            leader_cluster_resource_id=self.leader_cluster_resource_id,
            location=self.location,
            name=self.name,
            original_database_name=self.original_database_name,
            principals_modification_kind=self.principals_modification_kind,
            provisioning_state=self.provisioning_state,
            soft_delete_period=self.soft_delete_period,
            statistics=self.statistics,
            suspension_details=self.suspension_details,
            table_level_sharing_properties=self.table_level_sharing_properties,
            type=self.type)


def get_read_only_following_database(cluster_name: Optional[builtins.str] = None,
                                     database_name: Optional[builtins.str] = None,
                                     resource_group_name: Optional[builtins.str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReadOnlyFollowingDatabaseResult:
    """
    Returns a database.

    Uses Azure REST API version 2024-04-13.


    :param builtins.str cluster_name: The name of the Kusto cluster.
    :param builtins.str database_name: The name of the database in the Kusto cluster.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:kusto:getReadOnlyFollowingDatabase', __args__, opts=opts, typ=GetReadOnlyFollowingDatabaseResult).value

    return AwaitableGetReadOnlyFollowingDatabaseResult(
        attached_database_configuration_name=pulumi.get(__ret__, 'attached_database_configuration_name'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        database_share_origin=pulumi.get(__ret__, 'database_share_origin'),
        hot_cache_period=pulumi.get(__ret__, 'hot_cache_period'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        leader_cluster_resource_id=pulumi.get(__ret__, 'leader_cluster_resource_id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        original_database_name=pulumi.get(__ret__, 'original_database_name'),
        principals_modification_kind=pulumi.get(__ret__, 'principals_modification_kind'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        soft_delete_period=pulumi.get(__ret__, 'soft_delete_period'),
        statistics=pulumi.get(__ret__, 'statistics'),
        suspension_details=pulumi.get(__ret__, 'suspension_details'),
        table_level_sharing_properties=pulumi.get(__ret__, 'table_level_sharing_properties'),
        type=pulumi.get(__ret__, 'type'))
def get_read_only_following_database_output(cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                                            database_name: Optional[pulumi.Input[builtins.str]] = None,
                                            resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetReadOnlyFollowingDatabaseResult]:
    """
    Returns a database.

    Uses Azure REST API version 2024-04-13.


    :param builtins.str cluster_name: The name of the Kusto cluster.
    :param builtins.str database_name: The name of the database in the Kusto cluster.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:kusto:getReadOnlyFollowingDatabase', __args__, opts=opts, typ=GetReadOnlyFollowingDatabaseResult)
    return __ret__.apply(lambda __response__: GetReadOnlyFollowingDatabaseResult(
        attached_database_configuration_name=pulumi.get(__response__, 'attached_database_configuration_name'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        database_share_origin=pulumi.get(__response__, 'database_share_origin'),
        hot_cache_period=pulumi.get(__response__, 'hot_cache_period'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        leader_cluster_resource_id=pulumi.get(__response__, 'leader_cluster_resource_id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        original_database_name=pulumi.get(__response__, 'original_database_name'),
        principals_modification_kind=pulumi.get(__response__, 'principals_modification_kind'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        soft_delete_period=pulumi.get(__response__, 'soft_delete_period'),
        statistics=pulumi.get(__response__, 'statistics'),
        suspension_details=pulumi.get(__response__, 'suspension_details'),
        table_level_sharing_properties=pulumi.get(__response__, 'table_level_sharing_properties'),
        type=pulumi.get(__response__, 'type')))
