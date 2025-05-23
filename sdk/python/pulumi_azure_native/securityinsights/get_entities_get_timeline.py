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
from ._enums import *

__all__ = [
    'GetEntitiesGetTimelineResult',
    'AwaitableGetEntitiesGetTimelineResult',
    'get_entities_get_timeline',
    'get_entities_get_timeline_output',
]

@pulumi.output_type
class GetEntitiesGetTimelineResult:
    """
    The entity timeline result operation response.
    """
    def __init__(__self__, meta_data=None, value=None):
        if meta_data and not isinstance(meta_data, dict):
            raise TypeError("Expected argument 'meta_data' to be a dict")
        pulumi.set(__self__, "meta_data", meta_data)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="metaData")
    def meta_data(self) -> Optional['outputs.TimelineResultsMetadataResponse']:
        """
        The metadata from the timeline operation results.
        """
        return pulumi.get(self, "meta_data")

    @property
    @pulumi.getter
    def value(self) -> Optional[Sequence[Any]]:
        """
        The timeline result values.
        """
        return pulumi.get(self, "value")


class AwaitableGetEntitiesGetTimelineResult(GetEntitiesGetTimelineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEntitiesGetTimelineResult(
            meta_data=self.meta_data,
            value=self.value)


def get_entities_get_timeline(end_time: Optional[builtins.str] = None,
                              entity_id: Optional[builtins.str] = None,
                              kinds: Optional[Sequence[Union[builtins.str, 'EntityTimelineKind']]] = None,
                              number_of_bucket: Optional[builtins.int] = None,
                              resource_group_name: Optional[builtins.str] = None,
                              start_time: Optional[builtins.str] = None,
                              workspace_name: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEntitiesGetTimelineResult:
    """
    Timeline for an entity.

    Uses Azure REST API version 2025-01-01-preview.

    Other available API versions: 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str end_time: The end timeline date, so the results returned are before this date.
    :param builtins.str entity_id: entity ID
    :param Sequence[Union[builtins.str, 'EntityTimelineKind']] kinds: Array of timeline Item kinds.
    :param builtins.int number_of_bucket: The number of bucket for timeline queries aggregation.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str start_time: The start timeline date, so the results returned are after this date.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['endTime'] = end_time
    __args__['entityId'] = entity_id
    __args__['kinds'] = kinds
    __args__['numberOfBucket'] = number_of_bucket
    __args__['resourceGroupName'] = resource_group_name
    __args__['startTime'] = start_time
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:securityinsights:getEntitiesGetTimeline', __args__, opts=opts, typ=GetEntitiesGetTimelineResult).value

    return AwaitableGetEntitiesGetTimelineResult(
        meta_data=pulumi.get(__ret__, 'meta_data'),
        value=pulumi.get(__ret__, 'value'))
def get_entities_get_timeline_output(end_time: Optional[pulumi.Input[builtins.str]] = None,
                                     entity_id: Optional[pulumi.Input[builtins.str]] = None,
                                     kinds: Optional[pulumi.Input[Optional[Sequence[Union[builtins.str, 'EntityTimelineKind']]]]] = None,
                                     number_of_bucket: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                     resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                     start_time: Optional[pulumi.Input[builtins.str]] = None,
                                     workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEntitiesGetTimelineResult]:
    """
    Timeline for an entity.

    Uses Azure REST API version 2025-01-01-preview.

    Other available API versions: 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str end_time: The end timeline date, so the results returned are before this date.
    :param builtins.str entity_id: entity ID
    :param Sequence[Union[builtins.str, 'EntityTimelineKind']] kinds: Array of timeline Item kinds.
    :param builtins.int number_of_bucket: The number of bucket for timeline queries aggregation.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str start_time: The start timeline date, so the results returned are after this date.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['endTime'] = end_time
    __args__['entityId'] = entity_id
    __args__['kinds'] = kinds
    __args__['numberOfBucket'] = number_of_bucket
    __args__['resourceGroupName'] = resource_group_name
    __args__['startTime'] = start_time
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:securityinsights:getEntitiesGetTimeline', __args__, opts=opts, typ=GetEntitiesGetTimelineResult)
    return __ret__.apply(lambda __response__: GetEntitiesGetTimelineResult(
        meta_data=pulumi.get(__response__, 'meta_data'),
        value=pulumi.get(__response__, 'value')))
