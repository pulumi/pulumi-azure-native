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
    'GetChannelResult',
    'AwaitableGetChannelResult',
    'get_channel',
    'get_channel_output',
]

@pulumi.output_type
class GetChannelResult:
    """
    The EngagementFabric channel
    """
    def __init__(__self__, azure_api_version=None, channel_functions=None, channel_type=None, credentials=None, id=None, name=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if channel_functions and not isinstance(channel_functions, list):
            raise TypeError("Expected argument 'channel_functions' to be a list")
        pulumi.set(__self__, "channel_functions", channel_functions)
        if channel_type and not isinstance(channel_type, str):
            raise TypeError("Expected argument 'channel_type' to be a str")
        pulumi.set(__self__, "channel_type", channel_type)
        if credentials and not isinstance(credentials, dict):
            raise TypeError("Expected argument 'credentials' to be a dict")
        pulumi.set(__self__, "credentials", credentials)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="channelFunctions")
    def channel_functions(self) -> Optional[Sequence[builtins.str]]:
        """
        The functions to be enabled for the channel
        """
        return pulumi.get(self, "channel_functions")

    @property
    @pulumi.getter(name="channelType")
    def channel_type(self) -> builtins.str:
        """
        The channel type
        """
        return pulumi.get(self, "channel_type")

    @property
    @pulumi.getter
    def credentials(self) -> Optional[Mapping[str, builtins.str]]:
        """
        The channel credentials
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The ID of the resource
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
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The fully qualified type of the resource
        """
        return pulumi.get(self, "type")


class AwaitableGetChannelResult(GetChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetChannelResult(
            azure_api_version=self.azure_api_version,
            channel_functions=self.channel_functions,
            channel_type=self.channel_type,
            credentials=self.credentials,
            id=self.id,
            name=self.name,
            type=self.type)


def get_channel(account_name: Optional[builtins.str] = None,
                channel_name: Optional[builtins.str] = None,
                resource_group_name: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetChannelResult:
    """
    The EngagementFabric channel

    Uses Azure REST API version 2018-09-01-preview.


    :param builtins.str account_name: Account Name
    :param builtins.str channel_name: Channel Name
    :param builtins.str resource_group_name: Resource Group Name
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['channelName'] = channel_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:engagementfabric:getChannel', __args__, opts=opts, typ=GetChannelResult).value

    return AwaitableGetChannelResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        channel_functions=pulumi.get(__ret__, 'channel_functions'),
        channel_type=pulumi.get(__ret__, 'channel_type'),
        credentials=pulumi.get(__ret__, 'credentials'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        type=pulumi.get(__ret__, 'type'))
def get_channel_output(account_name: Optional[pulumi.Input[builtins.str]] = None,
                       channel_name: Optional[pulumi.Input[builtins.str]] = None,
                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetChannelResult]:
    """
    The EngagementFabric channel

    Uses Azure REST API version 2018-09-01-preview.


    :param builtins.str account_name: Account Name
    :param builtins.str channel_name: Channel Name
    :param builtins.str resource_group_name: Resource Group Name
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['channelName'] = channel_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:engagementfabric:getChannel', __args__, opts=opts, typ=GetChannelResult)
    return __ret__.apply(lambda __response__: GetChannelResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        channel_functions=pulumi.get(__response__, 'channel_functions'),
        channel_type=pulumi.get(__response__, 'channel_type'),
        credentials=pulumi.get(__response__, 'credentials'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        type=pulumi.get(__response__, 'type')))
