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
    'ListNamespaceTopicSharedAccessKeysResult',
    'AwaitableListNamespaceTopicSharedAccessKeysResult',
    'list_namespace_topic_shared_access_keys',
    'list_namespace_topic_shared_access_keys_output',
]

@pulumi.output_type
class ListNamespaceTopicSharedAccessKeysResult:
    """
    Shared access keys of the Topic
    """
    def __init__(__self__, key1=None, key2=None):
        if key1 and not isinstance(key1, str):
            raise TypeError("Expected argument 'key1' to be a str")
        pulumi.set(__self__, "key1", key1)
        if key2 and not isinstance(key2, str):
            raise TypeError("Expected argument 'key2' to be a str")
        pulumi.set(__self__, "key2", key2)

    @property
    @pulumi.getter
    def key1(self) -> Optional[builtins.str]:
        """
        Shared access key1 for the topic.
        """
        return pulumi.get(self, "key1")

    @property
    @pulumi.getter
    def key2(self) -> Optional[builtins.str]:
        """
        Shared access key2 for the topic.
        """
        return pulumi.get(self, "key2")


class AwaitableListNamespaceTopicSharedAccessKeysResult(ListNamespaceTopicSharedAccessKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListNamespaceTopicSharedAccessKeysResult(
            key1=self.key1,
            key2=self.key2)


def list_namespace_topic_shared_access_keys(namespace_name: Optional[builtins.str] = None,
                                            resource_group_name: Optional[builtins.str] = None,
                                            topic_name: Optional[builtins.str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListNamespaceTopicSharedAccessKeysResult:
    """
    List the two keys used to publish to a namespace topic.

    Uses Azure REST API version 2025-02-15.

    Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str namespace_name: Name of the namespace.
    :param builtins.str resource_group_name: The name of the resource group within the user's subscription.
    :param builtins.str topic_name: Name of the topic.
    """
    __args__ = dict()
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['topicName'] = topic_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:eventgrid:listNamespaceTopicSharedAccessKeys', __args__, opts=opts, typ=ListNamespaceTopicSharedAccessKeysResult).value

    return AwaitableListNamespaceTopicSharedAccessKeysResult(
        key1=pulumi.get(__ret__, 'key1'),
        key2=pulumi.get(__ret__, 'key2'))
def list_namespace_topic_shared_access_keys_output(namespace_name: Optional[pulumi.Input[builtins.str]] = None,
                                                   resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                   topic_name: Optional[pulumi.Input[builtins.str]] = None,
                                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListNamespaceTopicSharedAccessKeysResult]:
    """
    List the two keys used to publish to a namespace topic.

    Uses Azure REST API version 2025-02-15.

    Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str namespace_name: Name of the namespace.
    :param builtins.str resource_group_name: The name of the resource group within the user's subscription.
    :param builtins.str topic_name: Name of the topic.
    """
    __args__ = dict()
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['topicName'] = topic_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:eventgrid:listNamespaceTopicSharedAccessKeys', __args__, opts=opts, typ=ListNamespaceTopicSharedAccessKeysResult)
    return __ret__.apply(lambda __response__: ListNamespaceTopicSharedAccessKeysResult(
        key1=pulumi.get(__response__, 'key1'),
        key2=pulumi.get(__response__, 'key2')))
