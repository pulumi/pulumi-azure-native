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
    'GetSystemTopicResult',
    'AwaitableGetSystemTopicResult',
    'get_system_topic',
    'get_system_topic_output',
]

@pulumi.output_type
class GetSystemTopicResult:
    """
    EventGrid System Topic.
    """
    def __init__(__self__, azure_api_version=None, id=None, identity=None, location=None, metric_resource_id=None, name=None, provisioning_state=None, source=None, system_data=None, tags=None, topic_type=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if metric_resource_id and not isinstance(metric_resource_id, str):
            raise TypeError("Expected argument 'metric_resource_id' to be a str")
        pulumi.set(__self__, "metric_resource_id", metric_resource_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if topic_type and not isinstance(topic_type, str):
            raise TypeError("Expected argument 'topic_type' to be a str")
        pulumi.set(__self__, "topic_type", topic_type)
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
        Fully qualified identifier of the resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.IdentityInfoResponse']:
        """
        Identity information for the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="metricResourceId")
    def metric_resource_id(self) -> builtins.str:
        """
        Metric resource id for the system topic.
        """
        return pulumi.get(self, "metric_resource_id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning state of the system topic.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def source(self) -> Optional[builtins.str]:
        """
        Source for the system topic.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system metadata relating to the Event Grid resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="topicType")
    def topic_type(self) -> Optional[builtins.str]:
        """
        TopicType for the system topic.
        """
        return pulumi.get(self, "topic_type")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetSystemTopicResult(GetSystemTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemTopicResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            identity=self.identity,
            location=self.location,
            metric_resource_id=self.metric_resource_id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            source=self.source,
            system_data=self.system_data,
            tags=self.tags,
            topic_type=self.topic_type,
            type=self.type)


def get_system_topic(resource_group_name: Optional[builtins.str] = None,
                     system_topic_name: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemTopicResult:
    """
    Get properties of a system topic.

    Uses Azure REST API version 2025-02-15.

    Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group within the user's subscription.
    :param builtins.str system_topic_name: Name of the system topic.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['systemTopicName'] = system_topic_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:eventgrid:getSystemTopic', __args__, opts=opts, typ=GetSystemTopicResult).value

    return AwaitableGetSystemTopicResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        location=pulumi.get(__ret__, 'location'),
        metric_resource_id=pulumi.get(__ret__, 'metric_resource_id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        source=pulumi.get(__ret__, 'source'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        topic_type=pulumi.get(__ret__, 'topic_type'),
        type=pulumi.get(__ret__, 'type'))
def get_system_topic_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                            system_topic_name: Optional[pulumi.Input[builtins.str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSystemTopicResult]:
    """
    Get properties of a system topic.

    Uses Azure REST API version 2025-02-15.

    Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group within the user's subscription.
    :param builtins.str system_topic_name: Name of the system topic.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['systemTopicName'] = system_topic_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:eventgrid:getSystemTopic', __args__, opts=opts, typ=GetSystemTopicResult)
    return __ret__.apply(lambda __response__: GetSystemTopicResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        location=pulumi.get(__response__, 'location'),
        metric_resource_id=pulumi.get(__response__, 'metric_resource_id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        source=pulumi.get(__response__, 'source'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        topic_type=pulumi.get(__response__, 'topic_type'),
        type=pulumi.get(__response__, 'type')))
