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

__all__ = [
    'GetBackupShortTermRetentionPolicyResult',
    'AwaitableGetBackupShortTermRetentionPolicyResult',
    'get_backup_short_term_retention_policy',
    'get_backup_short_term_retention_policy_output',
]

@pulumi.output_type
class GetBackupShortTermRetentionPolicyResult:
    """
    A short term retention policy.
    """
    def __init__(__self__, azure_api_version=None, diff_backup_interval_in_hours=None, id=None, name=None, retention_days=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if diff_backup_interval_in_hours and not isinstance(diff_backup_interval_in_hours, int):
            raise TypeError("Expected argument 'diff_backup_interval_in_hours' to be a int")
        pulumi.set(__self__, "diff_backup_interval_in_hours", diff_backup_interval_in_hours)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if retention_days and not isinstance(retention_days, int):
            raise TypeError("Expected argument 'retention_days' to be a int")
        pulumi.set(__self__, "retention_days", retention_days)
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
    @pulumi.getter(name="diffBackupIntervalInHours")
    def diff_backup_interval_in_hours(self) -> Optional[builtins.int]:
        """
        The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
        """
        return pulumi.get(self, "diff_backup_interval_in_hours")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[builtins.int]:
        """
        The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetBackupShortTermRetentionPolicyResult(GetBackupShortTermRetentionPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupShortTermRetentionPolicyResult(
            azure_api_version=self.azure_api_version,
            diff_backup_interval_in_hours=self.diff_backup_interval_in_hours,
            id=self.id,
            name=self.name,
            retention_days=self.retention_days,
            type=self.type)


def get_backup_short_term_retention_policy(database_name: Optional[builtins.str] = None,
                                           policy_name: Optional[builtins.str] = None,
                                           resource_group_name: Optional[builtins.str] = None,
                                           server_name: Optional[builtins.str] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupShortTermRetentionPolicyResult:
    """
    Gets a database's short term retention policy.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2017-10-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str database_name: The name of the database.
    :param builtins.str policy_name: The policy name. Should always be "default".
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['policyName'] = policy_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:sql:getBackupShortTermRetentionPolicy', __args__, opts=opts, typ=GetBackupShortTermRetentionPolicyResult).value

    return AwaitableGetBackupShortTermRetentionPolicyResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        diff_backup_interval_in_hours=pulumi.get(__ret__, 'diff_backup_interval_in_hours'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        retention_days=pulumi.get(__ret__, 'retention_days'),
        type=pulumi.get(__ret__, 'type'))
def get_backup_short_term_retention_policy_output(database_name: Optional[pulumi.Input[builtins.str]] = None,
                                                  policy_name: Optional[pulumi.Input[builtins.str]] = None,
                                                  resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                  server_name: Optional[pulumi.Input[builtins.str]] = None,
                                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetBackupShortTermRetentionPolicyResult]:
    """
    Gets a database's short term retention policy.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2017-10-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str database_name: The name of the database.
    :param builtins.str policy_name: The policy name. Should always be "default".
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['policyName'] = policy_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:sql:getBackupShortTermRetentionPolicy', __args__, opts=opts, typ=GetBackupShortTermRetentionPolicyResult)
    return __ret__.apply(lambda __response__: GetBackupShortTermRetentionPolicyResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        diff_backup_interval_in_hours=pulumi.get(__response__, 'diff_backup_interval_in_hours'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        retention_days=pulumi.get(__response__, 'retention_days'),
        type=pulumi.get(__response__, 'type')))
