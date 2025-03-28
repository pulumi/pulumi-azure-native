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
from ... import _utilities
from . import outputs

__all__ = [
    'GetStreamingLocatorResult',
    'AwaitableGetStreamingLocatorResult',
    'get_streaming_locator',
    'get_streaming_locator_output',
]

@pulumi.output_type
class GetStreamingLocatorResult:
    """
    A Streaming Locator resource
    """
    def __init__(__self__, asset_name=None, content_keys=None, created=None, default_content_key_policy_name=None, end_time=None, id=None, name=None, start_time=None, streaming_locator_id=None, streaming_policy_name=None, type=None):
        if asset_name and not isinstance(asset_name, str):
            raise TypeError("Expected argument 'asset_name' to be a str")
        pulumi.set(__self__, "asset_name", asset_name)
        if content_keys and not isinstance(content_keys, list):
            raise TypeError("Expected argument 'content_keys' to be a list")
        pulumi.set(__self__, "content_keys", content_keys)
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if default_content_key_policy_name and not isinstance(default_content_key_policy_name, str):
            raise TypeError("Expected argument 'default_content_key_policy_name' to be a str")
        pulumi.set(__self__, "default_content_key_policy_name", default_content_key_policy_name)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if streaming_locator_id and not isinstance(streaming_locator_id, str):
            raise TypeError("Expected argument 'streaming_locator_id' to be a str")
        pulumi.set(__self__, "streaming_locator_id", streaming_locator_id)
        if streaming_policy_name and not isinstance(streaming_policy_name, str):
            raise TypeError("Expected argument 'streaming_policy_name' to be a str")
        pulumi.set(__self__, "streaming_policy_name", streaming_policy_name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="assetName")
    def asset_name(self) -> str:
        """
        Asset Name
        """
        return pulumi.get(self, "asset_name")

    @property
    @pulumi.getter(name="contentKeys")
    def content_keys(self) -> Optional[Sequence['outputs.StreamingLocatorUserDefinedContentKeyResponse']]:
        """
        ContentKeys used by this Streaming Locator
        """
        return pulumi.get(self, "content_keys")

    @property
    @pulumi.getter
    def created(self) -> str:
        """
        Creation time of Streaming Locator
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="defaultContentKeyPolicyName")
    def default_content_key_policy_name(self) -> Optional[str]:
        """
        Default ContentKeyPolicy used by this Streaming Locator
        """
        return pulumi.get(self, "default_content_key_policy_name")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
        """
        EndTime of Streaming Locator
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[str]:
        """
        StartTime of Streaming Locator
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="streamingLocatorId")
    def streaming_locator_id(self) -> Optional[str]:
        """
        StreamingLocatorId of Streaming Locator
        """
        return pulumi.get(self, "streaming_locator_id")

    @property
    @pulumi.getter(name="streamingPolicyName")
    def streaming_policy_name(self) -> str:
        """
        Streaming policy name used by this streaming locator. Either specify the name of streaming policy you created or use one of the predefined streaming polices. The predefined streaming policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_SecureStreaming' and 'Predefined_SecureStreamingWithFairPlay'
        """
        return pulumi.get(self, "streaming_policy_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetStreamingLocatorResult(GetStreamingLocatorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStreamingLocatorResult(
            asset_name=self.asset_name,
            content_keys=self.content_keys,
            created=self.created,
            default_content_key_policy_name=self.default_content_key_policy_name,
            end_time=self.end_time,
            id=self.id,
            name=self.name,
            start_time=self.start_time,
            streaming_locator_id=self.streaming_locator_id,
            streaming_policy_name=self.streaming_policy_name,
            type=self.type)


def get_streaming_locator(account_name: Optional[str] = None,
                          resource_group_name: Optional[str] = None,
                          streaming_locator_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStreamingLocatorResult:
    """
    Get the details of a Streaming Locator in the Media Services account


    :param str account_name: The Media Services account name.
    :param str resource_group_name: The name of the resource group within the Azure subscription.
    :param str streaming_locator_name: The Streaming Locator name.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['streamingLocatorName'] = streaming_locator_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:media/v20180330preview:getStreamingLocator', __args__, opts=opts, typ=GetStreamingLocatorResult).value

    return AwaitableGetStreamingLocatorResult(
        asset_name=pulumi.get(__ret__, 'asset_name'),
        content_keys=pulumi.get(__ret__, 'content_keys'),
        created=pulumi.get(__ret__, 'created'),
        default_content_key_policy_name=pulumi.get(__ret__, 'default_content_key_policy_name'),
        end_time=pulumi.get(__ret__, 'end_time'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        start_time=pulumi.get(__ret__, 'start_time'),
        streaming_locator_id=pulumi.get(__ret__, 'streaming_locator_id'),
        streaming_policy_name=pulumi.get(__ret__, 'streaming_policy_name'),
        type=pulumi.get(__ret__, 'type'))
def get_streaming_locator_output(account_name: Optional[pulumi.Input[str]] = None,
                                 resource_group_name: Optional[pulumi.Input[str]] = None,
                                 streaming_locator_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStreamingLocatorResult]:
    """
    Get the details of a Streaming Locator in the Media Services account


    :param str account_name: The Media Services account name.
    :param str resource_group_name: The name of the resource group within the Azure subscription.
    :param str streaming_locator_name: The Streaming Locator name.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['streamingLocatorName'] = streaming_locator_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:media/v20180330preview:getStreamingLocator', __args__, opts=opts, typ=GetStreamingLocatorResult)
    return __ret__.apply(lambda __response__: GetStreamingLocatorResult(
        asset_name=pulumi.get(__response__, 'asset_name'),
        content_keys=pulumi.get(__response__, 'content_keys'),
        created=pulumi.get(__response__, 'created'),
        default_content_key_policy_name=pulumi.get(__response__, 'default_content_key_policy_name'),
        end_time=pulumi.get(__response__, 'end_time'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        start_time=pulumi.get(__response__, 'start_time'),
        streaming_locator_id=pulumi.get(__response__, 'streaming_locator_id'),
        streaming_policy_name=pulumi.get(__response__, 'streaming_policy_name'),
        type=pulumi.get(__response__, 'type')))
