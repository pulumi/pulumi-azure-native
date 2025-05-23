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
    'GetBackupScheduleResult',
    'AwaitableGetBackupScheduleResult',
    'get_backup_schedule',
    'get_backup_schedule_output',
]

@pulumi.output_type
class GetBackupScheduleResult:
    """
    The backup schedule.
    """
    def __init__(__self__, azure_api_version=None, backup_type=None, id=None, kind=None, last_successful_run=None, name=None, retention_count=None, schedule_recurrence=None, schedule_status=None, start_time=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if backup_type and not isinstance(backup_type, str):
            raise TypeError("Expected argument 'backup_type' to be a str")
        pulumi.set(__self__, "backup_type", backup_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if last_successful_run and not isinstance(last_successful_run, str):
            raise TypeError("Expected argument 'last_successful_run' to be a str")
        pulumi.set(__self__, "last_successful_run", last_successful_run)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if retention_count and not isinstance(retention_count, float):
            raise TypeError("Expected argument 'retention_count' to be a float")
        pulumi.set(__self__, "retention_count", retention_count)
        if schedule_recurrence and not isinstance(schedule_recurrence, dict):
            raise TypeError("Expected argument 'schedule_recurrence' to be a dict")
        pulumi.set(__self__, "schedule_recurrence", schedule_recurrence)
        if schedule_status and not isinstance(schedule_status, str):
            raise TypeError("Expected argument 'schedule_status' to be a str")
        pulumi.set(__self__, "schedule_status", schedule_status)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
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
    @pulumi.getter(name="backupType")
    def backup_type(self) -> builtins.str:
        """
        The type of backup which needs to be taken.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The path ID that uniquely identifies the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[builtins.str]:
        """
        The Kind of the object. Currently only Series8000 is supported
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastSuccessfulRun")
    def last_successful_run(self) -> builtins.str:
        """
        The last successful backup run which was triggered for the schedule.
        """
        return pulumi.get(self, "last_successful_run")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the object.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionCount")
    def retention_count(self) -> builtins.float:
        """
        The number of backups to be retained.
        """
        return pulumi.get(self, "retention_count")

    @property
    @pulumi.getter(name="scheduleRecurrence")
    def schedule_recurrence(self) -> 'outputs.ScheduleRecurrenceResponse':
        """
        The schedule recurrence.
        """
        return pulumi.get(self, "schedule_recurrence")

    @property
    @pulumi.getter(name="scheduleStatus")
    def schedule_status(self) -> builtins.str:
        """
        The schedule status.
        """
        return pulumi.get(self, "schedule_status")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> builtins.str:
        """
        The start time of the schedule.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")


class AwaitableGetBackupScheduleResult(GetBackupScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupScheduleResult(
            azure_api_version=self.azure_api_version,
            backup_type=self.backup_type,
            id=self.id,
            kind=self.kind,
            last_successful_run=self.last_successful_run,
            name=self.name,
            retention_count=self.retention_count,
            schedule_recurrence=self.schedule_recurrence,
            schedule_status=self.schedule_status,
            start_time=self.start_time,
            type=self.type)


def get_backup_schedule(backup_policy_name: Optional[builtins.str] = None,
                        backup_schedule_name: Optional[builtins.str] = None,
                        device_name: Optional[builtins.str] = None,
                        manager_name: Optional[builtins.str] = None,
                        resource_group_name: Optional[builtins.str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupScheduleResult:
    """
    Gets the properties of the specified backup schedule name.

    Uses Azure REST API version 2017-06-01.


    :param builtins.str backup_policy_name: The backup policy name.
    :param builtins.str backup_schedule_name: The name of the backup schedule to be fetched
    :param builtins.str device_name: The device name
    :param builtins.str manager_name: The manager name
    :param builtins.str resource_group_name: The resource group name
    """
    __args__ = dict()
    __args__['backupPolicyName'] = backup_policy_name
    __args__['backupScheduleName'] = backup_schedule_name
    __args__['deviceName'] = device_name
    __args__['managerName'] = manager_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:storsimple:getBackupSchedule', __args__, opts=opts, typ=GetBackupScheduleResult).value

    return AwaitableGetBackupScheduleResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        backup_type=pulumi.get(__ret__, 'backup_type'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        last_successful_run=pulumi.get(__ret__, 'last_successful_run'),
        name=pulumi.get(__ret__, 'name'),
        retention_count=pulumi.get(__ret__, 'retention_count'),
        schedule_recurrence=pulumi.get(__ret__, 'schedule_recurrence'),
        schedule_status=pulumi.get(__ret__, 'schedule_status'),
        start_time=pulumi.get(__ret__, 'start_time'),
        type=pulumi.get(__ret__, 'type'))
def get_backup_schedule_output(backup_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                               backup_schedule_name: Optional[pulumi.Input[builtins.str]] = None,
                               device_name: Optional[pulumi.Input[builtins.str]] = None,
                               manager_name: Optional[pulumi.Input[builtins.str]] = None,
                               resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetBackupScheduleResult]:
    """
    Gets the properties of the specified backup schedule name.

    Uses Azure REST API version 2017-06-01.


    :param builtins.str backup_policy_name: The backup policy name.
    :param builtins.str backup_schedule_name: The name of the backup schedule to be fetched
    :param builtins.str device_name: The device name
    :param builtins.str manager_name: The manager name
    :param builtins.str resource_group_name: The resource group name
    """
    __args__ = dict()
    __args__['backupPolicyName'] = backup_policy_name
    __args__['backupScheduleName'] = backup_schedule_name
    __args__['deviceName'] = device_name
    __args__['managerName'] = manager_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:storsimple:getBackupSchedule', __args__, opts=opts, typ=GetBackupScheduleResult)
    return __ret__.apply(lambda __response__: GetBackupScheduleResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        backup_type=pulumi.get(__response__, 'backup_type'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        last_successful_run=pulumi.get(__response__, 'last_successful_run'),
        name=pulumi.get(__response__, 'name'),
        retention_count=pulumi.get(__response__, 'retention_count'),
        schedule_recurrence=pulumi.get(__response__, 'schedule_recurrence'),
        schedule_status=pulumi.get(__response__, 'schedule_status'),
        start_time=pulumi.get(__response__, 'start_time'),
        type=pulumi.get(__response__, 'type')))
