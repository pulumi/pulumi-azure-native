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
    'GetPeriodicTimerEventTriggerResult',
    'AwaitableGetPeriodicTimerEventTriggerResult',
    'get_periodic_timer_event_trigger',
    'get_periodic_timer_event_trigger_output',
]

@pulumi.output_type
class GetPeriodicTimerEventTriggerResult:
    """
    Trigger details.
    """
    def __init__(__self__, azure_api_version=None, custom_context_tag=None, id=None, kind=None, name=None, sink_info=None, source_info=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if custom_context_tag and not isinstance(custom_context_tag, str):
            raise TypeError("Expected argument 'custom_context_tag' to be a str")
        pulumi.set(__self__, "custom_context_tag", custom_context_tag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sink_info and not isinstance(sink_info, dict):
            raise TypeError("Expected argument 'sink_info' to be a dict")
        pulumi.set(__self__, "sink_info", sink_info)
        if source_info and not isinstance(source_info, dict):
            raise TypeError("Expected argument 'source_info' to be a dict")
        pulumi.set(__self__, "source_info", source_info)
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
    @pulumi.getter(name="customContextTag")
    def custom_context_tag(self) -> Optional[builtins.str]:
        """
        A custom context tag typically used to correlate the trigger against its usage. For example, if a periodic timer trigger is intended for certain specific IoT modules in the device, the tag can be the name or the image URL of the module.
        """
        return pulumi.get(self, "custom_context_tag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The path ID that uniquely identifies the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> builtins.str:
        """
        Trigger Kind.
        Expected value is 'PeriodicTimerEvent'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The object name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sinkInfo")
    def sink_info(self) -> 'outputs.RoleSinkInfoResponse':
        """
        Role Sink information.
        """
        return pulumi.get(self, "sink_info")

    @property
    @pulumi.getter(name="sourceInfo")
    def source_info(self) -> 'outputs.PeriodicTimerSourceInfoResponse':
        """
        Periodic timer details.
        """
        return pulumi.get(self, "source_info")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Metadata pertaining to creation and last modification of Trigger
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")


class AwaitableGetPeriodicTimerEventTriggerResult(GetPeriodicTimerEventTriggerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPeriodicTimerEventTriggerResult(
            azure_api_version=self.azure_api_version,
            custom_context_tag=self.custom_context_tag,
            id=self.id,
            kind=self.kind,
            name=self.name,
            sink_info=self.sink_info,
            source_info=self.source_info,
            system_data=self.system_data,
            type=self.type)


def get_periodic_timer_event_trigger(device_name: Optional[builtins.str] = None,
                                     name: Optional[builtins.str] = None,
                                     resource_group_name: Optional[builtins.str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPeriodicTimerEventTriggerResult:
    """
    Get a specific trigger by name.

    Uses Azure REST API version 2023-07-01.


    :param builtins.str device_name: The device name.
    :param builtins.str name: The trigger name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:databoxedge:getPeriodicTimerEventTrigger', __args__, opts=opts, typ=GetPeriodicTimerEventTriggerResult).value

    return AwaitableGetPeriodicTimerEventTriggerResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        custom_context_tag=pulumi.get(__ret__, 'custom_context_tag'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        name=pulumi.get(__ret__, 'name'),
        sink_info=pulumi.get(__ret__, 'sink_info'),
        source_info=pulumi.get(__ret__, 'source_info'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_periodic_timer_event_trigger_output(device_name: Optional[pulumi.Input[builtins.str]] = None,
                                            name: Optional[pulumi.Input[builtins.str]] = None,
                                            resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPeriodicTimerEventTriggerResult]:
    """
    Get a specific trigger by name.

    Uses Azure REST API version 2023-07-01.


    :param builtins.str device_name: The device name.
    :param builtins.str name: The trigger name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:databoxedge:getPeriodicTimerEventTrigger', __args__, opts=opts, typ=GetPeriodicTimerEventTriggerResult)
    return __ret__.apply(lambda __response__: GetPeriodicTimerEventTriggerResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        custom_context_tag=pulumi.get(__response__, 'custom_context_tag'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        name=pulumi.get(__response__, 'name'),
        sink_info=pulumi.get(__response__, 'sink_info'),
        source_info=pulumi.get(__response__, 'source_info'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
