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
from ._inputs import *

__all__ = ['IoTHubEventSourceArgs', 'IoTHubEventSource']

@pulumi.input_type
class IoTHubEventSourceArgs:
    def __init__(__self__, *,
                 consumer_group_name: pulumi.Input[builtins.str],
                 environment_name: pulumi.Input[builtins.str],
                 event_source_resource_id: pulumi.Input[builtins.str],
                 iot_hub_name: pulumi.Input[builtins.str],
                 key_name: pulumi.Input[builtins.str],
                 kind: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 shared_access_key: pulumi.Input[builtins.str],
                 event_source_name: Optional[pulumi.Input[builtins.str]] = None,
                 local_timestamp: Optional[pulumi.Input['LocalTimestampArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 time: Optional[pulumi.Input[builtins.str]] = None,
                 timestamp_property_name: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[Union[builtins.str, 'IngressStartAtType']]] = None):
        """
        The set of arguments for constructing a IoTHubEventSource resource.
        :param pulumi.Input[builtins.str] consumer_group_name: The name of the iot hub's consumer group that holds the partitions from which events will be read.
        :param pulumi.Input[builtins.str] environment_name: The name of the Time Series Insights environment associated with the specified resource group.
        :param pulumi.Input[builtins.str] event_source_resource_id: The resource id of the event source in Azure Resource Manager.
        :param pulumi.Input[builtins.str] iot_hub_name: The name of the iot hub.
        :param pulumi.Input[builtins.str] key_name: The name of the Shared Access Policy key that grants the Time Series Insights service access to the iot hub. This shared access policy key must grant 'service connect' permissions to the iot hub.
        :param pulumi.Input[builtins.str] kind: The kind of the event source.
               Expected value is 'Microsoft.IoTHub'.
        :param pulumi.Input[builtins.str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[builtins.str] shared_access_key: The value of the Shared Access Policy key that grants the Time Series Insights service read access to the iot hub. This property is not shown in event source responses.
        :param pulumi.Input[builtins.str] event_source_name: Name of the event source.
        :param pulumi.Input['LocalTimestampArgs'] local_timestamp: An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
        :param pulumi.Input[builtins.str] location: The location of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Key-value pairs of additional properties for the resource.
        :param pulumi.Input[builtins.str] time: ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
        :param pulumi.Input[builtins.str] timestamp_property_name: The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
        :param pulumi.Input[Union[builtins.str, 'IngressStartAtType']] type: The type of the ingressStartAt, It can be "EarliestAvailable", "EventSourceCreationTime", "CustomEnqueuedTime".
        """
        pulumi.set(__self__, "consumer_group_name", consumer_group_name)
        pulumi.set(__self__, "environment_name", environment_name)
        pulumi.set(__self__, "event_source_resource_id", event_source_resource_id)
        pulumi.set(__self__, "iot_hub_name", iot_hub_name)
        pulumi.set(__self__, "key_name", key_name)
        pulumi.set(__self__, "kind", 'Microsoft.IoTHub')
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "shared_access_key", shared_access_key)
        if event_source_name is not None:
            pulumi.set(__self__, "event_source_name", event_source_name)
        if local_timestamp is not None:
            pulumi.set(__self__, "local_timestamp", local_timestamp)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if time is not None:
            pulumi.set(__self__, "time", time)
        if timestamp_property_name is not None:
            pulumi.set(__self__, "timestamp_property_name", timestamp_property_name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="consumerGroupName")
    def consumer_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the iot hub's consumer group that holds the partitions from which events will be read.
        """
        return pulumi.get(self, "consumer_group_name")

    @consumer_group_name.setter
    def consumer_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "consumer_group_name", value)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Time Series Insights environment associated with the specified resource group.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "environment_name", value)

    @property
    @pulumi.getter(name="eventSourceResourceId")
    def event_source_resource_id(self) -> pulumi.Input[builtins.str]:
        """
        The resource id of the event source in Azure Resource Manager.
        """
        return pulumi.get(self, "event_source_resource_id")

    @event_source_resource_id.setter
    def event_source_resource_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "event_source_resource_id", value)

    @property
    @pulumi.getter(name="iotHubName")
    def iot_hub_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the iot hub.
        """
        return pulumi.get(self, "iot_hub_name")

    @iot_hub_name.setter
    def iot_hub_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "iot_hub_name", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Shared Access Policy key that grants the Time Series Insights service access to the iot hub. This shared access policy key must grant 'service connect' permissions to the iot hub.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[builtins.str]:
        """
        The kind of the event source.
        Expected value is 'Microsoft.IoTHub'.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of an Azure Resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="sharedAccessKey")
    def shared_access_key(self) -> pulumi.Input[builtins.str]:
        """
        The value of the Shared Access Policy key that grants the Time Series Insights service read access to the iot hub. This property is not shown in event source responses.
        """
        return pulumi.get(self, "shared_access_key")

    @shared_access_key.setter
    def shared_access_key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "shared_access_key", value)

    @property
    @pulumi.getter(name="eventSourceName")
    def event_source_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the event source.
        """
        return pulumi.get(self, "event_source_name")

    @event_source_name.setter
    def event_source_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "event_source_name", value)

    @property
    @pulumi.getter(name="localTimestamp")
    def local_timestamp(self) -> Optional[pulumi.Input['LocalTimestampArgs']]:
        """
        An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
        """
        return pulumi.get(self, "local_timestamp")

    @local_timestamp.setter
    def local_timestamp(self, value: Optional[pulumi.Input['LocalTimestampArgs']]):
        pulumi.set(self, "local_timestamp", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Key-value pairs of additional properties for the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def time(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
        """
        return pulumi.get(self, "time")

    @time.setter
    def time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "time", value)

    @property
    @pulumi.getter(name="timestampPropertyName")
    def timestamp_property_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
        """
        return pulumi.get(self, "timestamp_property_name")

    @timestamp_property_name.setter
    def timestamp_property_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "timestamp_property_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[builtins.str, 'IngressStartAtType']]]:
        """
        The type of the ingressStartAt, It can be "EarliestAvailable", "EventSourceCreationTime", "CustomEnqueuedTime".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[builtins.str, 'IngressStartAtType']]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("azure-native:timeseriesinsights:IoTHubEventSource")
class IoTHubEventSource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 event_source_name: Optional[pulumi.Input[builtins.str]] = None,
                 event_source_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 iot_hub_name: Optional[pulumi.Input[builtins.str]] = None,
                 key_name: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 local_timestamp: Optional[pulumi.Input[Union['LocalTimestampArgs', 'LocalTimestampArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 shared_access_key: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 time: Optional[pulumi.Input[builtins.str]] = None,
                 timestamp_property_name: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[Union[builtins.str, 'IngressStartAtType']]] = None,
                 __props__=None):
        """
        An event source that receives its data from an Azure IoTHub.

        Uses Azure REST API version 2020-05-15. In version 2.x of the Azure Native provider, it used API version 2020-05-15.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] consumer_group_name: The name of the iot hub's consumer group that holds the partitions from which events will be read.
        :param pulumi.Input[builtins.str] environment_name: The name of the Time Series Insights environment associated with the specified resource group.
        :param pulumi.Input[builtins.str] event_source_name: Name of the event source.
        :param pulumi.Input[builtins.str] event_source_resource_id: The resource id of the event source in Azure Resource Manager.
        :param pulumi.Input[builtins.str] iot_hub_name: The name of the iot hub.
        :param pulumi.Input[builtins.str] key_name: The name of the Shared Access Policy key that grants the Time Series Insights service access to the iot hub. This shared access policy key must grant 'service connect' permissions to the iot hub.
        :param pulumi.Input[builtins.str] kind: The kind of the event source.
               Expected value is 'Microsoft.IoTHub'.
        :param pulumi.Input[Union['LocalTimestampArgs', 'LocalTimestampArgsDict']] local_timestamp: An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
        :param pulumi.Input[builtins.str] location: The location of the resource.
        :param pulumi.Input[builtins.str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[builtins.str] shared_access_key: The value of the Shared Access Policy key that grants the Time Series Insights service read access to the iot hub. This property is not shown in event source responses.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Key-value pairs of additional properties for the resource.
        :param pulumi.Input[builtins.str] time: ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
        :param pulumi.Input[builtins.str] timestamp_property_name: The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
        :param pulumi.Input[Union[builtins.str, 'IngressStartAtType']] type: The type of the ingressStartAt, It can be "EarliestAvailable", "EventSourceCreationTime", "CustomEnqueuedTime".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IoTHubEventSourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An event source that receives its data from an Azure IoTHub.

        Uses Azure REST API version 2020-05-15. In version 2.x of the Azure Native provider, it used API version 2020-05-15.

        :param str resource_name: The name of the resource.
        :param IoTHubEventSourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IoTHubEventSourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 event_source_name: Optional[pulumi.Input[builtins.str]] = None,
                 event_source_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 iot_hub_name: Optional[pulumi.Input[builtins.str]] = None,
                 key_name: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 local_timestamp: Optional[pulumi.Input[Union['LocalTimestampArgs', 'LocalTimestampArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 shared_access_key: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 time: Optional[pulumi.Input[builtins.str]] = None,
                 timestamp_property_name: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[Union[builtins.str, 'IngressStartAtType']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IoTHubEventSourceArgs.__new__(IoTHubEventSourceArgs)

            if consumer_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'consumer_group_name'")
            __props__.__dict__["consumer_group_name"] = consumer_group_name
            if environment_name is None and not opts.urn:
                raise TypeError("Missing required property 'environment_name'")
            __props__.__dict__["environment_name"] = environment_name
            __props__.__dict__["event_source_name"] = event_source_name
            if event_source_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'event_source_resource_id'")
            __props__.__dict__["event_source_resource_id"] = event_source_resource_id
            if iot_hub_name is None and not opts.urn:
                raise TypeError("Missing required property 'iot_hub_name'")
            __props__.__dict__["iot_hub_name"] = iot_hub_name
            if key_name is None and not opts.urn:
                raise TypeError("Missing required property 'key_name'")
            __props__.__dict__["key_name"] = key_name
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = 'Microsoft.IoTHub'
            __props__.__dict__["local_timestamp"] = local_timestamp
            __props__.__dict__["location"] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if shared_access_key is None and not opts.urn:
                raise TypeError("Missing required property 'shared_access_key'")
            __props__.__dict__["shared_access_key"] = shared_access_key
            __props__.__dict__["tags"] = tags
            __props__.__dict__["time"] = time
            __props__.__dict__["timestamp_property_name"] = timestamp_property_name
            __props__.__dict__["type"] = type
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:timeseriesinsights/v20170228preview:IoTHubEventSource"), pulumi.Alias(type_="azure-native:timeseriesinsights/v20171115:IoTHubEventSource"), pulumi.Alias(type_="azure-native:timeseriesinsights/v20180815preview:IoTHubEventSource"), pulumi.Alias(type_="azure-native:timeseriesinsights/v20200515:IoTHubEventSource"), pulumi.Alias(type_="azure-native:timeseriesinsights/v20210331preview:IoTHubEventSource"), pulumi.Alias(type_="azure-native:timeseriesinsights/v20210630preview:EventHubEventSource"), pulumi.Alias(type_="azure-native:timeseriesinsights/v20210630preview:IoTHubEventSource"), pulumi.Alias(type_="azure-native:timeseriesinsights:EventHubEventSource")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(IoTHubEventSource, __self__).__init__(
            'azure-native:timeseriesinsights:IoTHubEventSource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IoTHubEventSource':
        """
        Get an existing IoTHubEventSource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IoTHubEventSourceArgs.__new__(IoTHubEventSourceArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["consumer_group_name"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["event_source_resource_id"] = None
        __props__.__dict__["iot_hub_name"] = None
        __props__.__dict__["key_name"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["local_timestamp"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["time"] = None
        __props__.__dict__["timestamp_property_name"] = None
        __props__.__dict__["type"] = None
        return IoTHubEventSource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="consumerGroupName")
    def consumer_group_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the iot hub's consumer group that holds the partitions from which events will be read.
        """
        return pulumi.get(self, "consumer_group_name")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[builtins.str]:
        """
        The time the resource was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="eventSourceResourceId")
    def event_source_resource_id(self) -> pulumi.Output[builtins.str]:
        """
        The resource id of the event source in Azure Resource Manager.
        """
        return pulumi.get(self, "event_source_resource_id")

    @property
    @pulumi.getter(name="iotHubName")
    def iot_hub_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the iot hub.
        """
        return pulumi.get(self, "iot_hub_name")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the Shared Access Policy key that grants the Time Series Insights service access to the iot hub. This shared access policy key must grant 'service connect' permissions to the iot hub.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        The kind of the event source.
        Expected value is 'Microsoft.IoTHub'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="localTimestamp")
    def local_timestamp(self) -> pulumi.Output[Optional['outputs.LocalTimestampResponse']]:
        """
        An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
        """
        return pulumi.get(self, "local_timestamp")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def time(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
        """
        return pulumi.get(self, "time")

    @property
    @pulumi.getter(name="timestampPropertyName")
    def timestamp_property_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
        """
        return pulumi.get(self, "timestamp_property_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

