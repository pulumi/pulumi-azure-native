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
    'GetAutoscaleSettingResult',
    'AwaitableGetAutoscaleSettingResult',
    'get_autoscale_setting',
    'get_autoscale_setting_output',
]

@pulumi.output_type
class GetAutoscaleSettingResult:
    """
    The autoscale setting resource.
    """
    def __init__(__self__, enabled=None, id=None, location=None, name=None, notifications=None, predictive_autoscale_policy=None, profiles=None, system_data=None, tags=None, target_resource_location=None, target_resource_uri=None, type=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notifications and not isinstance(notifications, list):
            raise TypeError("Expected argument 'notifications' to be a list")
        pulumi.set(__self__, "notifications", notifications)
        if predictive_autoscale_policy and not isinstance(predictive_autoscale_policy, dict):
            raise TypeError("Expected argument 'predictive_autoscale_policy' to be a dict")
        pulumi.set(__self__, "predictive_autoscale_policy", predictive_autoscale_policy)
        if profiles and not isinstance(profiles, list):
            raise TypeError("Expected argument 'profiles' to be a list")
        pulumi.set(__self__, "profiles", profiles)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target_resource_location and not isinstance(target_resource_location, str):
            raise TypeError("Expected argument 'target_resource_location' to be a str")
        pulumi.set(__self__, "target_resource_location", target_resource_location)
        if target_resource_uri and not isinstance(target_resource_uri, str):
            raise TypeError("Expected argument 'target_resource_uri' to be a str")
        pulumi.set(__self__, "target_resource_uri", target_resource_uri)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'false'.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Azure resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> Optional[Sequence['outputs.AutoscaleNotificationResponse']]:
        """
        the collection of notifications.
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter(name="predictiveAutoscalePolicy")
    def predictive_autoscale_policy(self) -> Optional['outputs.PredictiveAutoscalePolicyResponse']:
        """
        the predictive autoscale policy mode.
        """
        return pulumi.get(self, "predictive_autoscale_policy")

    @property
    @pulumi.getter
    def profiles(self) -> Sequence['outputs.AutoscaleProfileResponse']:
        """
        the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
        """
        return pulumi.get(self, "profiles")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system metadata related to the response.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater in length than 128 characters and a value no greater in length than 256 characters.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetResourceLocation")
    def target_resource_location(self) -> Optional[str]:
        """
        the location of the resource that the autoscale setting should be added to.
        """
        return pulumi.get(self, "target_resource_location")

    @property
    @pulumi.getter(name="targetResourceUri")
    def target_resource_uri(self) -> Optional[str]:
        """
        the resource identifier of the resource that the autoscale setting should be added to.
        """
        return pulumi.get(self, "target_resource_uri")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetAutoscaleSettingResult(GetAutoscaleSettingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAutoscaleSettingResult(
            enabled=self.enabled,
            id=self.id,
            location=self.location,
            name=self.name,
            notifications=self.notifications,
            predictive_autoscale_policy=self.predictive_autoscale_policy,
            profiles=self.profiles,
            system_data=self.system_data,
            tags=self.tags,
            target_resource_location=self.target_resource_location,
            target_resource_uri=self.target_resource_uri,
            type=self.type)


def get_autoscale_setting(autoscale_setting_name: Optional[str] = None,
                          resource_group_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAutoscaleSettingResult:
    """
    Gets an autoscale setting

    Uses Azure REST API version 2022-10-01.


    :param str autoscale_setting_name: The autoscale setting name.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['autoscaleSettingName'] = autoscale_setting_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:insights:getAutoscaleSetting', __args__, opts=opts, typ=GetAutoscaleSettingResult).value

    return AwaitableGetAutoscaleSettingResult(
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        notifications=pulumi.get(__ret__, 'notifications'),
        predictive_autoscale_policy=pulumi.get(__ret__, 'predictive_autoscale_policy'),
        profiles=pulumi.get(__ret__, 'profiles'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        target_resource_location=pulumi.get(__ret__, 'target_resource_location'),
        target_resource_uri=pulumi.get(__ret__, 'target_resource_uri'),
        type=pulumi.get(__ret__, 'type'))
def get_autoscale_setting_output(autoscale_setting_name: Optional[pulumi.Input[str]] = None,
                                 resource_group_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAutoscaleSettingResult]:
    """
    Gets an autoscale setting

    Uses Azure REST API version 2022-10-01.


    :param str autoscale_setting_name: The autoscale setting name.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['autoscaleSettingName'] = autoscale_setting_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:insights:getAutoscaleSetting', __args__, opts=opts, typ=GetAutoscaleSettingResult)
    return __ret__.apply(lambda __response__: GetAutoscaleSettingResult(
        enabled=pulumi.get(__response__, 'enabled'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        notifications=pulumi.get(__response__, 'notifications'),
        predictive_autoscale_policy=pulumi.get(__response__, 'predictive_autoscale_policy'),
        profiles=pulumi.get(__response__, 'profiles'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        target_resource_location=pulumi.get(__response__, 'target_resource_location'),
        target_resource_uri=pulumi.get(__response__, 'target_resource_uri'),
        type=pulumi.get(__response__, 'type')))
