# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetMonitoredSubscriptionResult',
    'AwaitableGetMonitoredSubscriptionResult',
    'get_monitored_subscription',
    'get_monitored_subscription_output',
]

@pulumi.output_type
class GetMonitoredSubscriptionResult:
    """
    The request to update subscriptions needed to be monitored by the Elastic monitor resource.
    """
    def __init__(__self__, id=None, name=None, properties=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the monitored subscription resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the monitored subscription resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.SubscriptionListResponse':
        """
        The request to update subscriptions needed to be monitored by the Elastic monitor resource.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the monitored subscription resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetMonitoredSubscriptionResult(GetMonitoredSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMonitoredSubscriptionResult(
            id=self.id,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_monitored_subscription(configuration_name: Optional[str] = None,
                               monitor_name: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMonitoredSubscriptionResult:
    """
    The request to update subscriptions needed to be monitored by the Elastic monitor resource.

    Uses Azure REST API version 2024-05-01-preview.

    Other available API versions: 2024-06-15-preview, 2024-10-01-preview, 2025-01-15-preview.


    :param str configuration_name: The configuration name. Only 'default' value is supported.
    :param str monitor_name: Monitor resource name
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['configurationName'] = configuration_name
    __args__['monitorName'] = monitor_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:elastic:getMonitoredSubscription', __args__, opts=opts, typ=GetMonitoredSubscriptionResult).value

    return AwaitableGetMonitoredSubscriptionResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        type=pulumi.get(__ret__, 'type'))
def get_monitored_subscription_output(configuration_name: Optional[pulumi.Input[str]] = None,
                                      monitor_name: Optional[pulumi.Input[str]] = None,
                                      resource_group_name: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMonitoredSubscriptionResult]:
    """
    The request to update subscriptions needed to be monitored by the Elastic monitor resource.

    Uses Azure REST API version 2024-05-01-preview.

    Other available API versions: 2024-06-15-preview, 2024-10-01-preview, 2025-01-15-preview.


    :param str configuration_name: The configuration name. Only 'default' value is supported.
    :param str monitor_name: Monitor resource name
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['configurationName'] = configuration_name
    __args__['monitorName'] = monitor_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:elastic:getMonitoredSubscription', __args__, opts=opts, typ=GetMonitoredSubscriptionResult)
    return __ret__.apply(lambda __response__: GetMonitoredSubscriptionResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        type=pulumi.get(__response__, 'type')))
