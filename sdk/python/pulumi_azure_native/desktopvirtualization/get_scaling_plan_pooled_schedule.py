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
    'GetScalingPlanPooledScheduleResult',
    'AwaitableGetScalingPlanPooledScheduleResult',
    'get_scaling_plan_pooled_schedule',
    'get_scaling_plan_pooled_schedule_output',
]

@pulumi.output_type
class GetScalingPlanPooledScheduleResult:
    """
    Represents a ScalingPlanPooledSchedule definition.
    """
    def __init__(__self__, azure_api_version=None, days_of_week=None, id=None, name=None, off_peak_load_balancing_algorithm=None, off_peak_start_time=None, peak_load_balancing_algorithm=None, peak_start_time=None, ramp_down_capacity_threshold_pct=None, ramp_down_force_logoff_users=None, ramp_down_load_balancing_algorithm=None, ramp_down_minimum_hosts_pct=None, ramp_down_notification_message=None, ramp_down_start_time=None, ramp_down_stop_hosts_when=None, ramp_down_wait_time_minutes=None, ramp_up_capacity_threshold_pct=None, ramp_up_load_balancing_algorithm=None, ramp_up_minimum_hosts_pct=None, ramp_up_start_time=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if days_of_week and not isinstance(days_of_week, list):
            raise TypeError("Expected argument 'days_of_week' to be a list")
        pulumi.set(__self__, "days_of_week", days_of_week)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if off_peak_load_balancing_algorithm and not isinstance(off_peak_load_balancing_algorithm, str):
            raise TypeError("Expected argument 'off_peak_load_balancing_algorithm' to be a str")
        pulumi.set(__self__, "off_peak_load_balancing_algorithm", off_peak_load_balancing_algorithm)
        if off_peak_start_time and not isinstance(off_peak_start_time, dict):
            raise TypeError("Expected argument 'off_peak_start_time' to be a dict")
        pulumi.set(__self__, "off_peak_start_time", off_peak_start_time)
        if peak_load_balancing_algorithm and not isinstance(peak_load_balancing_algorithm, str):
            raise TypeError("Expected argument 'peak_load_balancing_algorithm' to be a str")
        pulumi.set(__self__, "peak_load_balancing_algorithm", peak_load_balancing_algorithm)
        if peak_start_time and not isinstance(peak_start_time, dict):
            raise TypeError("Expected argument 'peak_start_time' to be a dict")
        pulumi.set(__self__, "peak_start_time", peak_start_time)
        if ramp_down_capacity_threshold_pct and not isinstance(ramp_down_capacity_threshold_pct, int):
            raise TypeError("Expected argument 'ramp_down_capacity_threshold_pct' to be a int")
        pulumi.set(__self__, "ramp_down_capacity_threshold_pct", ramp_down_capacity_threshold_pct)
        if ramp_down_force_logoff_users and not isinstance(ramp_down_force_logoff_users, bool):
            raise TypeError("Expected argument 'ramp_down_force_logoff_users' to be a bool")
        pulumi.set(__self__, "ramp_down_force_logoff_users", ramp_down_force_logoff_users)
        if ramp_down_load_balancing_algorithm and not isinstance(ramp_down_load_balancing_algorithm, str):
            raise TypeError("Expected argument 'ramp_down_load_balancing_algorithm' to be a str")
        pulumi.set(__self__, "ramp_down_load_balancing_algorithm", ramp_down_load_balancing_algorithm)
        if ramp_down_minimum_hosts_pct and not isinstance(ramp_down_minimum_hosts_pct, int):
            raise TypeError("Expected argument 'ramp_down_minimum_hosts_pct' to be a int")
        pulumi.set(__self__, "ramp_down_minimum_hosts_pct", ramp_down_minimum_hosts_pct)
        if ramp_down_notification_message and not isinstance(ramp_down_notification_message, str):
            raise TypeError("Expected argument 'ramp_down_notification_message' to be a str")
        pulumi.set(__self__, "ramp_down_notification_message", ramp_down_notification_message)
        if ramp_down_start_time and not isinstance(ramp_down_start_time, dict):
            raise TypeError("Expected argument 'ramp_down_start_time' to be a dict")
        pulumi.set(__self__, "ramp_down_start_time", ramp_down_start_time)
        if ramp_down_stop_hosts_when and not isinstance(ramp_down_stop_hosts_when, str):
            raise TypeError("Expected argument 'ramp_down_stop_hosts_when' to be a str")
        pulumi.set(__self__, "ramp_down_stop_hosts_when", ramp_down_stop_hosts_when)
        if ramp_down_wait_time_minutes and not isinstance(ramp_down_wait_time_minutes, int):
            raise TypeError("Expected argument 'ramp_down_wait_time_minutes' to be a int")
        pulumi.set(__self__, "ramp_down_wait_time_minutes", ramp_down_wait_time_minutes)
        if ramp_up_capacity_threshold_pct and not isinstance(ramp_up_capacity_threshold_pct, int):
            raise TypeError("Expected argument 'ramp_up_capacity_threshold_pct' to be a int")
        pulumi.set(__self__, "ramp_up_capacity_threshold_pct", ramp_up_capacity_threshold_pct)
        if ramp_up_load_balancing_algorithm and not isinstance(ramp_up_load_balancing_algorithm, str):
            raise TypeError("Expected argument 'ramp_up_load_balancing_algorithm' to be a str")
        pulumi.set(__self__, "ramp_up_load_balancing_algorithm", ramp_up_load_balancing_algorithm)
        if ramp_up_minimum_hosts_pct and not isinstance(ramp_up_minimum_hosts_pct, int):
            raise TypeError("Expected argument 'ramp_up_minimum_hosts_pct' to be a int")
        pulumi.set(__self__, "ramp_up_minimum_hosts_pct", ramp_up_minimum_hosts_pct)
        if ramp_up_start_time and not isinstance(ramp_up_start_time, dict):
            raise TypeError("Expected argument 'ramp_up_start_time' to be a dict")
        pulumi.set(__self__, "ramp_up_start_time", ramp_up_start_time)
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
    @pulumi.getter(name="daysOfWeek")
    def days_of_week(self) -> Optional[Sequence[builtins.str]]:
        """
        Set of days of the week on which this schedule is active.
        """
        return pulumi.get(self, "days_of_week")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="offPeakLoadBalancingAlgorithm")
    def off_peak_load_balancing_algorithm(self) -> Optional[builtins.str]:
        """
        Load balancing algorithm for off-peak period.
        """
        return pulumi.get(self, "off_peak_load_balancing_algorithm")

    @property
    @pulumi.getter(name="offPeakStartTime")
    def off_peak_start_time(self) -> Optional['outputs.TimeResponse']:
        """
        Starting time for off-peak period.
        """
        return pulumi.get(self, "off_peak_start_time")

    @property
    @pulumi.getter(name="peakLoadBalancingAlgorithm")
    def peak_load_balancing_algorithm(self) -> Optional[builtins.str]:
        """
        Load balancing algorithm for peak period.
        """
        return pulumi.get(self, "peak_load_balancing_algorithm")

    @property
    @pulumi.getter(name="peakStartTime")
    def peak_start_time(self) -> Optional['outputs.TimeResponse']:
        """
        Starting time for peak period.
        """
        return pulumi.get(self, "peak_start_time")

    @property
    @pulumi.getter(name="rampDownCapacityThresholdPct")
    def ramp_down_capacity_threshold_pct(self) -> Optional[builtins.int]:
        """
        Capacity threshold for ramp down period.
        """
        return pulumi.get(self, "ramp_down_capacity_threshold_pct")

    @property
    @pulumi.getter(name="rampDownForceLogoffUsers")
    def ramp_down_force_logoff_users(self) -> Optional[builtins.bool]:
        """
        Should users be logged off forcefully from hosts.
        """
        return pulumi.get(self, "ramp_down_force_logoff_users")

    @property
    @pulumi.getter(name="rampDownLoadBalancingAlgorithm")
    def ramp_down_load_balancing_algorithm(self) -> Optional[builtins.str]:
        """
        Load balancing algorithm for ramp down period.
        """
        return pulumi.get(self, "ramp_down_load_balancing_algorithm")

    @property
    @pulumi.getter(name="rampDownMinimumHostsPct")
    def ramp_down_minimum_hosts_pct(self) -> Optional[builtins.int]:
        """
        Minimum host percentage for ramp down period.
        """
        return pulumi.get(self, "ramp_down_minimum_hosts_pct")

    @property
    @pulumi.getter(name="rampDownNotificationMessage")
    def ramp_down_notification_message(self) -> Optional[builtins.str]:
        """
        Notification message for users during ramp down period.
        """
        return pulumi.get(self, "ramp_down_notification_message")

    @property
    @pulumi.getter(name="rampDownStartTime")
    def ramp_down_start_time(self) -> Optional['outputs.TimeResponse']:
        """
        Starting time for ramp down period.
        """
        return pulumi.get(self, "ramp_down_start_time")

    @property
    @pulumi.getter(name="rampDownStopHostsWhen")
    def ramp_down_stop_hosts_when(self) -> Optional[builtins.str]:
        """
        Specifies when to stop hosts during ramp down period.
        """
        return pulumi.get(self, "ramp_down_stop_hosts_when")

    @property
    @pulumi.getter(name="rampDownWaitTimeMinutes")
    def ramp_down_wait_time_minutes(self) -> Optional[builtins.int]:
        """
        Number of minutes to wait to stop hosts during ramp down period.
        """
        return pulumi.get(self, "ramp_down_wait_time_minutes")

    @property
    @pulumi.getter(name="rampUpCapacityThresholdPct")
    def ramp_up_capacity_threshold_pct(self) -> Optional[builtins.int]:
        """
        Capacity threshold for ramp up period.
        """
        return pulumi.get(self, "ramp_up_capacity_threshold_pct")

    @property
    @pulumi.getter(name="rampUpLoadBalancingAlgorithm")
    def ramp_up_load_balancing_algorithm(self) -> Optional[builtins.str]:
        """
        Load balancing algorithm for ramp up period.
        """
        return pulumi.get(self, "ramp_up_load_balancing_algorithm")

    @property
    @pulumi.getter(name="rampUpMinimumHostsPct")
    def ramp_up_minimum_hosts_pct(self) -> Optional[builtins.int]:
        """
        Minimum host percentage for ramp up period.
        """
        return pulumi.get(self, "ramp_up_minimum_hosts_pct")

    @property
    @pulumi.getter(name="rampUpStartTime")
    def ramp_up_start_time(self) -> Optional['outputs.TimeResponse']:
        """
        Starting time for ramp up period.
        """
        return pulumi.get(self, "ramp_up_start_time")

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


class AwaitableGetScalingPlanPooledScheduleResult(GetScalingPlanPooledScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScalingPlanPooledScheduleResult(
            azure_api_version=self.azure_api_version,
            days_of_week=self.days_of_week,
            id=self.id,
            name=self.name,
            off_peak_load_balancing_algorithm=self.off_peak_load_balancing_algorithm,
            off_peak_start_time=self.off_peak_start_time,
            peak_load_balancing_algorithm=self.peak_load_balancing_algorithm,
            peak_start_time=self.peak_start_time,
            ramp_down_capacity_threshold_pct=self.ramp_down_capacity_threshold_pct,
            ramp_down_force_logoff_users=self.ramp_down_force_logoff_users,
            ramp_down_load_balancing_algorithm=self.ramp_down_load_balancing_algorithm,
            ramp_down_minimum_hosts_pct=self.ramp_down_minimum_hosts_pct,
            ramp_down_notification_message=self.ramp_down_notification_message,
            ramp_down_start_time=self.ramp_down_start_time,
            ramp_down_stop_hosts_when=self.ramp_down_stop_hosts_when,
            ramp_down_wait_time_minutes=self.ramp_down_wait_time_minutes,
            ramp_up_capacity_threshold_pct=self.ramp_up_capacity_threshold_pct,
            ramp_up_load_balancing_algorithm=self.ramp_up_load_balancing_algorithm,
            ramp_up_minimum_hosts_pct=self.ramp_up_minimum_hosts_pct,
            ramp_up_start_time=self.ramp_up_start_time,
            system_data=self.system_data,
            type=self.type)


def get_scaling_plan_pooled_schedule(resource_group_name: Optional[builtins.str] = None,
                                     scaling_plan_name: Optional[builtins.str] = None,
                                     scaling_plan_schedule_name: Optional[builtins.str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScalingPlanPooledScheduleResult:
    """
    Get a ScalingPlanPooledSchedule.

    Uses Azure REST API version 2024-04-03.

    Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str scaling_plan_name: The name of the scaling plan.
    :param builtins.str scaling_plan_schedule_name: The name of the ScalingPlanSchedule
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['scalingPlanName'] = scaling_plan_name
    __args__['scalingPlanScheduleName'] = scaling_plan_schedule_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:desktopvirtualization:getScalingPlanPooledSchedule', __args__, opts=opts, typ=GetScalingPlanPooledScheduleResult).value

    return AwaitableGetScalingPlanPooledScheduleResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        days_of_week=pulumi.get(__ret__, 'days_of_week'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        off_peak_load_balancing_algorithm=pulumi.get(__ret__, 'off_peak_load_balancing_algorithm'),
        off_peak_start_time=pulumi.get(__ret__, 'off_peak_start_time'),
        peak_load_balancing_algorithm=pulumi.get(__ret__, 'peak_load_balancing_algorithm'),
        peak_start_time=pulumi.get(__ret__, 'peak_start_time'),
        ramp_down_capacity_threshold_pct=pulumi.get(__ret__, 'ramp_down_capacity_threshold_pct'),
        ramp_down_force_logoff_users=pulumi.get(__ret__, 'ramp_down_force_logoff_users'),
        ramp_down_load_balancing_algorithm=pulumi.get(__ret__, 'ramp_down_load_balancing_algorithm'),
        ramp_down_minimum_hosts_pct=pulumi.get(__ret__, 'ramp_down_minimum_hosts_pct'),
        ramp_down_notification_message=pulumi.get(__ret__, 'ramp_down_notification_message'),
        ramp_down_start_time=pulumi.get(__ret__, 'ramp_down_start_time'),
        ramp_down_stop_hosts_when=pulumi.get(__ret__, 'ramp_down_stop_hosts_when'),
        ramp_down_wait_time_minutes=pulumi.get(__ret__, 'ramp_down_wait_time_minutes'),
        ramp_up_capacity_threshold_pct=pulumi.get(__ret__, 'ramp_up_capacity_threshold_pct'),
        ramp_up_load_balancing_algorithm=pulumi.get(__ret__, 'ramp_up_load_balancing_algorithm'),
        ramp_up_minimum_hosts_pct=pulumi.get(__ret__, 'ramp_up_minimum_hosts_pct'),
        ramp_up_start_time=pulumi.get(__ret__, 'ramp_up_start_time'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_scaling_plan_pooled_schedule_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                            scaling_plan_name: Optional[pulumi.Input[builtins.str]] = None,
                                            scaling_plan_schedule_name: Optional[pulumi.Input[builtins.str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetScalingPlanPooledScheduleResult]:
    """
    Get a ScalingPlanPooledSchedule.

    Uses Azure REST API version 2024-04-03.

    Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str scaling_plan_name: The name of the scaling plan.
    :param builtins.str scaling_plan_schedule_name: The name of the ScalingPlanSchedule
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['scalingPlanName'] = scaling_plan_name
    __args__['scalingPlanScheduleName'] = scaling_plan_schedule_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:desktopvirtualization:getScalingPlanPooledSchedule', __args__, opts=opts, typ=GetScalingPlanPooledScheduleResult)
    return __ret__.apply(lambda __response__: GetScalingPlanPooledScheduleResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        days_of_week=pulumi.get(__response__, 'days_of_week'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        off_peak_load_balancing_algorithm=pulumi.get(__response__, 'off_peak_load_balancing_algorithm'),
        off_peak_start_time=pulumi.get(__response__, 'off_peak_start_time'),
        peak_load_balancing_algorithm=pulumi.get(__response__, 'peak_load_balancing_algorithm'),
        peak_start_time=pulumi.get(__response__, 'peak_start_time'),
        ramp_down_capacity_threshold_pct=pulumi.get(__response__, 'ramp_down_capacity_threshold_pct'),
        ramp_down_force_logoff_users=pulumi.get(__response__, 'ramp_down_force_logoff_users'),
        ramp_down_load_balancing_algorithm=pulumi.get(__response__, 'ramp_down_load_balancing_algorithm'),
        ramp_down_minimum_hosts_pct=pulumi.get(__response__, 'ramp_down_minimum_hosts_pct'),
        ramp_down_notification_message=pulumi.get(__response__, 'ramp_down_notification_message'),
        ramp_down_start_time=pulumi.get(__response__, 'ramp_down_start_time'),
        ramp_down_stop_hosts_when=pulumi.get(__response__, 'ramp_down_stop_hosts_when'),
        ramp_down_wait_time_minutes=pulumi.get(__response__, 'ramp_down_wait_time_minutes'),
        ramp_up_capacity_threshold_pct=pulumi.get(__response__, 'ramp_up_capacity_threshold_pct'),
        ramp_up_load_balancing_algorithm=pulumi.get(__response__, 'ramp_up_load_balancing_algorithm'),
        ramp_up_minimum_hosts_pct=pulumi.get(__response__, 'ramp_up_minimum_hosts_pct'),
        ramp_up_start_time=pulumi.get(__response__, 'ramp_up_start_time'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
