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
    def __init__(__self__, attached_database_configuration_name=None, azure_api_version=None, hot_cache_period=None, id=None, kind=None, leader_cluster_resource_id=None, location=None, name=None, principals_modification_kind=None, provisioning_state=None, soft_delete_period=None, statistics=None, system_data=None, type=None):
        if attached_database_configuration_name and not isinstance(attached_database_configuration_name, str):
            raise TypeError("Expected argument 'attached_database_configuration_name' to be a str")
        pulumi.set(__self__, "attached_database_configuration_name", attached_database_configuration_name)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
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
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
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


class AwaitableGetReadOnlyFollowingDatabaseResult(GetReadOnlyFollowingDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReadOnlyFollowingDatabaseResult(
            attached_database_configuration_name=self.attached_database_configuration_name,
            azure_api_version=self.azure_api_version,
            hot_cache_period=self.hot_cache_period,
            id=self.id,
            kind=self.kind,
            leader_cluster_resource_id=self.leader_cluster_resource_id,
            location=self.location,
            name=self.name,
            principals_modification_kind=self.principals_modification_kind,
            provisioning_state=self.provisioning_state,
            soft_delete_period=self.soft_delete_period,
            statistics=self.statistics,
            system_data=self.system_data,
            type=self.type)


def get_read_only_following_database(database_name: Optional[builtins.str] = None,
                                     kusto_pool_name: Optional[builtins.str] = None,
                                     resource_group_name: Optional[builtins.str] = None,
                                     workspace_name: Optional[builtins.str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReadOnlyFollowingDatabaseResult:
    """
    Returns a database.

    Uses Azure REST API version 2021-06-01-preview.


    :param builtins.str database_name: The name of the database in the Kusto pool.
    :param builtins.str kusto_pool_name: The name of the Kusto pool.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['kustoPoolName'] = kusto_pool_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:synapse:getReadOnlyFollowingDatabase', __args__, opts=opts, typ=GetReadOnlyFollowingDatabaseResult).value

    return AwaitableGetReadOnlyFollowingDatabaseResult(
        attached_database_configuration_name=pulumi.get(__ret__, 'attached_database_configuration_name'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        hot_cache_period=pulumi.get(__ret__, 'hot_cache_period'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        leader_cluster_resource_id=pulumi.get(__ret__, 'leader_cluster_resource_id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        principals_modification_kind=pulumi.get(__ret__, 'principals_modification_kind'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        soft_delete_period=pulumi.get(__ret__, 'soft_delete_period'),
        statistics=pulumi.get(__ret__, 'statistics'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_read_only_following_database_output(database_name: Optional[pulumi.Input[builtins.str]] = None,
                                            kusto_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                                            resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                            workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetReadOnlyFollowingDatabaseResult]:
    """
    Returns a database.

    Uses Azure REST API version 2021-06-01-preview.


    :param builtins.str database_name: The name of the database in the Kusto pool.
    :param builtins.str kusto_pool_name: The name of the Kusto pool.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['kustoPoolName'] = kusto_pool_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:synapse:getReadOnlyFollowingDatabase', __args__, opts=opts, typ=GetReadOnlyFollowingDatabaseResult)
    return __ret__.apply(lambda __response__: GetReadOnlyFollowingDatabaseResult(
        attached_database_configuration_name=pulumi.get(__response__, 'attached_database_configuration_name'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        hot_cache_period=pulumi.get(__response__, 'hot_cache_period'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        leader_cluster_resource_id=pulumi.get(__response__, 'leader_cluster_resource_id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        principals_modification_kind=pulumi.get(__response__, 'principals_modification_kind'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        soft_delete_period=pulumi.get(__response__, 'soft_delete_period'),
        statistics=pulumi.get(__response__, 'statistics'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
