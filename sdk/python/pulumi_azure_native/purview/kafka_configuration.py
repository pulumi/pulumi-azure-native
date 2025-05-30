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

__all__ = ['KafkaConfigurationArgs', 'KafkaConfiguration']

@pulumi.input_type
class KafkaConfigurationArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 consumer_group: Optional[pulumi.Input[builtins.str]] = None,
                 credentials: Optional[pulumi.Input['CredentialsArgs']] = None,
                 event_hub_partition_id: Optional[pulumi.Input[builtins.str]] = None,
                 event_hub_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 event_hub_type: Optional[pulumi.Input[Union[builtins.str, 'EventHubType']]] = None,
                 event_streaming_state: Optional[pulumi.Input[Union[builtins.str, 'EventStreamingState']]] = None,
                 event_streaming_type: Optional[pulumi.Input[Union[builtins.str, 'EventStreamingType']]] = None,
                 kafka_configuration_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a KafkaConfiguration resource.
        :param pulumi.Input[builtins.str] account_name: The name of the account.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input[builtins.str] consumer_group: Consumer group for hook event hub.
        :param pulumi.Input['CredentialsArgs'] credentials: Credentials to access the event streaming service attached to the purview account.
        :param pulumi.Input[builtins.str] event_hub_partition_id: Optional partition Id for notification event hub. If not set, all partitions will be leveraged.
        :param pulumi.Input[Union[builtins.str, 'EventHubType']] event_hub_type: The event hub type.
        :param pulumi.Input[Union[builtins.str, 'EventStreamingState']] event_streaming_state: The state of the event streaming service
        :param pulumi.Input[Union[builtins.str, 'EventStreamingType']] event_streaming_type: The event streaming service type
        :param pulumi.Input[builtins.str] kafka_configuration_name: The kafka configuration name.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if consumer_group is not None:
            pulumi.set(__self__, "consumer_group", consumer_group)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if event_hub_partition_id is not None:
            pulumi.set(__self__, "event_hub_partition_id", event_hub_partition_id)
        if event_hub_resource_id is not None:
            pulumi.set(__self__, "event_hub_resource_id", event_hub_resource_id)
        if event_hub_type is not None:
            pulumi.set(__self__, "event_hub_type", event_hub_type)
        if event_streaming_state is None:
            event_streaming_state = 'Enabled'
        if event_streaming_state is not None:
            pulumi.set(__self__, "event_streaming_state", event_streaming_state)
        if event_streaming_type is None:
            event_streaming_type = 'None'
        if event_streaming_type is not None:
            pulumi.set(__self__, "event_streaming_type", event_streaming_type)
        if kafka_configuration_name is not None:
            pulumi.set(__self__, "kafka_configuration_name", kafka_configuration_name)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the account.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The resource group name.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Consumer group for hook event hub.
        """
        return pulumi.get(self, "consumer_group")

    @consumer_group.setter
    def consumer_group(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "consumer_group", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input['CredentialsArgs']]:
        """
        Credentials to access the event streaming service attached to the purview account.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input['CredentialsArgs']]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="eventHubPartitionId")
    def event_hub_partition_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Optional partition Id for notification event hub. If not set, all partitions will be leveraged.
        """
        return pulumi.get(self, "event_hub_partition_id")

    @event_hub_partition_id.setter
    def event_hub_partition_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "event_hub_partition_id", value)

    @property
    @pulumi.getter(name="eventHubResourceId")
    def event_hub_resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "event_hub_resource_id")

    @event_hub_resource_id.setter
    def event_hub_resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "event_hub_resource_id", value)

    @property
    @pulumi.getter(name="eventHubType")
    def event_hub_type(self) -> Optional[pulumi.Input[Union[builtins.str, 'EventHubType']]]:
        """
        The event hub type.
        """
        return pulumi.get(self, "event_hub_type")

    @event_hub_type.setter
    def event_hub_type(self, value: Optional[pulumi.Input[Union[builtins.str, 'EventHubType']]]):
        pulumi.set(self, "event_hub_type", value)

    @property
    @pulumi.getter(name="eventStreamingState")
    def event_streaming_state(self) -> Optional[pulumi.Input[Union[builtins.str, 'EventStreamingState']]]:
        """
        The state of the event streaming service
        """
        return pulumi.get(self, "event_streaming_state")

    @event_streaming_state.setter
    def event_streaming_state(self, value: Optional[pulumi.Input[Union[builtins.str, 'EventStreamingState']]]):
        pulumi.set(self, "event_streaming_state", value)

    @property
    @pulumi.getter(name="eventStreamingType")
    def event_streaming_type(self) -> Optional[pulumi.Input[Union[builtins.str, 'EventStreamingType']]]:
        """
        The event streaming service type
        """
        return pulumi.get(self, "event_streaming_type")

    @event_streaming_type.setter
    def event_streaming_type(self, value: Optional[pulumi.Input[Union[builtins.str, 'EventStreamingType']]]):
        pulumi.set(self, "event_streaming_type", value)

    @property
    @pulumi.getter(name="kafkaConfigurationName")
    def kafka_configuration_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The kafka configuration name.
        """
        return pulumi.get(self, "kafka_configuration_name")

    @kafka_configuration_name.setter
    def kafka_configuration_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kafka_configuration_name", value)


@pulumi.type_token("azure-native:purview:KafkaConfiguration")
class KafkaConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 consumer_group: Optional[pulumi.Input[builtins.str]] = None,
                 credentials: Optional[pulumi.Input[Union['CredentialsArgs', 'CredentialsArgsDict']]] = None,
                 event_hub_partition_id: Optional[pulumi.Input[builtins.str]] = None,
                 event_hub_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 event_hub_type: Optional[pulumi.Input[Union[builtins.str, 'EventHubType']]] = None,
                 event_streaming_state: Optional[pulumi.Input[Union[builtins.str, 'EventStreamingState']]] = None,
                 event_streaming_type: Optional[pulumi.Input[Union[builtins.str, 'EventStreamingType']]] = None,
                 kafka_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The configuration of the event streaming service resource attached to the Purview account for kafka notifications.

        Uses Azure REST API version 2024-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-12-01.

        Other available API versions: 2021-12-01, 2023-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native purview [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] account_name: The name of the account.
        :param pulumi.Input[builtins.str] consumer_group: Consumer group for hook event hub.
        :param pulumi.Input[Union['CredentialsArgs', 'CredentialsArgsDict']] credentials: Credentials to access the event streaming service attached to the purview account.
        :param pulumi.Input[builtins.str] event_hub_partition_id: Optional partition Id for notification event hub. If not set, all partitions will be leveraged.
        :param pulumi.Input[Union[builtins.str, 'EventHubType']] event_hub_type: The event hub type.
        :param pulumi.Input[Union[builtins.str, 'EventStreamingState']] event_streaming_state: The state of the event streaming service
        :param pulumi.Input[Union[builtins.str, 'EventStreamingType']] event_streaming_type: The event streaming service type
        :param pulumi.Input[builtins.str] kafka_configuration_name: The kafka configuration name.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KafkaConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The configuration of the event streaming service resource attached to the Purview account for kafka notifications.

        Uses Azure REST API version 2024-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-12-01.

        Other available API versions: 2021-12-01, 2023-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native purview [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param KafkaConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KafkaConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 consumer_group: Optional[pulumi.Input[builtins.str]] = None,
                 credentials: Optional[pulumi.Input[Union['CredentialsArgs', 'CredentialsArgsDict']]] = None,
                 event_hub_partition_id: Optional[pulumi.Input[builtins.str]] = None,
                 event_hub_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 event_hub_type: Optional[pulumi.Input[Union[builtins.str, 'EventHubType']]] = None,
                 event_streaming_state: Optional[pulumi.Input[Union[builtins.str, 'EventStreamingState']]] = None,
                 event_streaming_type: Optional[pulumi.Input[Union[builtins.str, 'EventStreamingType']]] = None,
                 kafka_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KafkaConfigurationArgs.__new__(KafkaConfigurationArgs)

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["consumer_group"] = consumer_group
            __props__.__dict__["credentials"] = credentials
            __props__.__dict__["event_hub_partition_id"] = event_hub_partition_id
            __props__.__dict__["event_hub_resource_id"] = event_hub_resource_id
            __props__.__dict__["event_hub_type"] = event_hub_type
            if event_streaming_state is None:
                event_streaming_state = 'Enabled'
            __props__.__dict__["event_streaming_state"] = event_streaming_state
            if event_streaming_type is None:
                event_streaming_type = 'None'
            __props__.__dict__["event_streaming_type"] = event_streaming_type
            __props__.__dict__["kafka_configuration_name"] = kafka_configuration_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:purview/v20211201:KafkaConfiguration"), pulumi.Alias(type_="azure-native:purview/v20230501preview:KafkaConfiguration"), pulumi.Alias(type_="azure-native:purview/v20240401preview:KafkaConfiguration")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(KafkaConfiguration, __self__).__init__(
            'azure-native:purview:KafkaConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'KafkaConfiguration':
        """
        Get an existing KafkaConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KafkaConfigurationArgs.__new__(KafkaConfigurationArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["consumer_group"] = None
        __props__.__dict__["credentials"] = None
        __props__.__dict__["event_hub_partition_id"] = None
        __props__.__dict__["event_hub_resource_id"] = None
        __props__.__dict__["event_hub_type"] = None
        __props__.__dict__["event_streaming_state"] = None
        __props__.__dict__["event_streaming_type"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return KafkaConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Consumer group for hook event hub.
        """
        return pulumi.get(self, "consumer_group")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional['outputs.CredentialsResponse']]:
        """
        Credentials to access the event streaming service attached to the purview account.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="eventHubPartitionId")
    def event_hub_partition_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Optional partition Id for notification event hub. If not set, all partitions will be leveraged.
        """
        return pulumi.get(self, "event_hub_partition_id")

    @property
    @pulumi.getter(name="eventHubResourceId")
    def event_hub_resource_id(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "event_hub_resource_id")

    @property
    @pulumi.getter(name="eventHubType")
    def event_hub_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The event hub type.
        """
        return pulumi.get(self, "event_hub_type")

    @property
    @pulumi.getter(name="eventStreamingState")
    def event_streaming_state(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The state of the event streaming service
        """
        return pulumi.get(self, "event_streaming_state")

    @property
    @pulumi.getter(name="eventStreamingType")
    def event_streaming_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The event streaming service type
        """
        return pulumi.get(self, "event_streaming_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.ProxyResourceResponseSystemData']:
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the type.
        """
        return pulumi.get(self, "type")

