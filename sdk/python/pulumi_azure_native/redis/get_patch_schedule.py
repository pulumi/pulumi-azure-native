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
    'GetPatchScheduleResult',
    'AwaitableGetPatchScheduleResult',
    'get_patch_schedule',
    'get_patch_schedule_output',
]

@pulumi.output_type
class GetPatchScheduleResult:
    """
    Response to put/get patch schedules for Redis cache.
    """
    def __init__(__self__, azure_api_version=None, id=None, location=None, name=None, schedule_entries=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if schedule_entries and not isinstance(schedule_entries, list):
            raise TypeError("Expected argument 'schedule_entries' to be a list")
        pulumi.set(__self__, "schedule_entries", schedule_entries)
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
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
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
    @pulumi.getter(name="scheduleEntries")
    def schedule_entries(self) -> Sequence['outputs.ScheduleEntryResponse']:
        """
        List of patch schedules for a Redis cache.
        """
        return pulumi.get(self, "schedule_entries")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetPatchScheduleResult(GetPatchScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPatchScheduleResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            location=self.location,
            name=self.name,
            schedule_entries=self.schedule_entries,
            type=self.type)


def get_patch_schedule(default: Optional[builtins.str] = None,
                       name: Optional[builtins.str] = None,
                       resource_group_name: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPatchScheduleResult:
    """
    Gets the patching schedule of a redis cache.

    Uses Azure REST API version 2024-11-01.

    Other available API versions: 2016-04-01, 2017-02-01, 2017-10-01, 2018-03-01, 2019-07-01, 2020-06-01, 2020-12-01, 2021-06-01, 2022-05-01, 2022-06-01, 2023-04-01, 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str default: Default string modeled as parameter for auto generation to work correctly.
    :param builtins.str name: The name of the redis cache.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['default'] = default
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:redis:getPatchSchedule', __args__, opts=opts, typ=GetPatchScheduleResult).value

    return AwaitableGetPatchScheduleResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        schedule_entries=pulumi.get(__ret__, 'schedule_entries'),
        type=pulumi.get(__ret__, 'type'))
def get_patch_schedule_output(default: Optional[pulumi.Input[builtins.str]] = None,
                              name: Optional[pulumi.Input[builtins.str]] = None,
                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPatchScheduleResult]:
    """
    Gets the patching schedule of a redis cache.

    Uses Azure REST API version 2024-11-01.

    Other available API versions: 2016-04-01, 2017-02-01, 2017-10-01, 2018-03-01, 2019-07-01, 2020-06-01, 2020-12-01, 2021-06-01, 2022-05-01, 2022-06-01, 2023-04-01, 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str default: Default string modeled as parameter for auto generation to work correctly.
    :param builtins.str name: The name of the redis cache.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['default'] = default
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:redis:getPatchSchedule', __args__, opts=opts, typ=GetPatchScheduleResult)
    return __ret__.apply(lambda __response__: GetPatchScheduleResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        schedule_entries=pulumi.get(__response__, 'schedule_entries'),
        type=pulumi.get(__response__, 'type')))
