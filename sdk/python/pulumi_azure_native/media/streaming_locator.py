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
from ._inputs import *

__all__ = ['StreamingLocatorArgs', 'StreamingLocator']

@pulumi.input_type
class StreamingLocatorArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[builtins.str],
                 asset_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 streaming_policy_name: pulumi.Input[builtins.str],
                 alternative_media_id: Optional[pulumi.Input[builtins.str]] = None,
                 content_keys: Optional[pulumi.Input[Sequence[pulumi.Input['StreamingLocatorContentKeyArgs']]]] = None,
                 default_content_key_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 end_time: Optional[pulumi.Input[builtins.str]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 start_time: Optional[pulumi.Input[builtins.str]] = None,
                 streaming_locator_id: Optional[pulumi.Input[builtins.str]] = None,
                 streaming_locator_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a StreamingLocator resource.
        :param pulumi.Input[builtins.str] account_name: The Media Services account name.
        :param pulumi.Input[builtins.str] asset_name: Asset Name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[builtins.str] streaming_policy_name: Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
        :param pulumi.Input[builtins.str] alternative_media_id: Alternative Media ID of this Streaming Locator
        :param pulumi.Input[Sequence[pulumi.Input['StreamingLocatorContentKeyArgs']]] content_keys: The ContentKeys used by this Streaming Locator.
        :param pulumi.Input[builtins.str] default_content_key_policy_name: Name of the default ContentKeyPolicy used by this Streaming Locator.
        :param pulumi.Input[builtins.str] end_time: The end time of the Streaming Locator.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] filters: A list of asset or account filters which apply to this streaming locator
        :param pulumi.Input[builtins.str] start_time: The start time of the Streaming Locator.
        :param pulumi.Input[builtins.str] streaming_locator_id: The StreamingLocatorId of the Streaming Locator.
        :param pulumi.Input[builtins.str] streaming_locator_name: The Streaming Locator name.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "asset_name", asset_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "streaming_policy_name", streaming_policy_name)
        if alternative_media_id is not None:
            pulumi.set(__self__, "alternative_media_id", alternative_media_id)
        if content_keys is not None:
            pulumi.set(__self__, "content_keys", content_keys)
        if default_content_key_policy_name is not None:
            pulumi.set(__self__, "default_content_key_policy_name", default_content_key_policy_name)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if streaming_locator_id is not None:
            pulumi.set(__self__, "streaming_locator_id", streaming_locator_id)
        if streaming_locator_name is not None:
            pulumi.set(__self__, "streaming_locator_name", streaming_locator_name)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[builtins.str]:
        """
        The Media Services account name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="assetName")
    def asset_name(self) -> pulumi.Input[builtins.str]:
        """
        Asset Name
        """
        return pulumi.get(self, "asset_name")

    @asset_name.setter
    def asset_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "asset_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group within the Azure subscription.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="streamingPolicyName")
    def streaming_policy_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
        """
        return pulumi.get(self, "streaming_policy_name")

    @streaming_policy_name.setter
    def streaming_policy_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "streaming_policy_name", value)

    @property
    @pulumi.getter(name="alternativeMediaId")
    def alternative_media_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Alternative Media ID of this Streaming Locator
        """
        return pulumi.get(self, "alternative_media_id")

    @alternative_media_id.setter
    def alternative_media_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "alternative_media_id", value)

    @property
    @pulumi.getter(name="contentKeys")
    def content_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StreamingLocatorContentKeyArgs']]]]:
        """
        The ContentKeys used by this Streaming Locator.
        """
        return pulumi.get(self, "content_keys")

    @content_keys.setter
    def content_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StreamingLocatorContentKeyArgs']]]]):
        pulumi.set(self, "content_keys", value)

    @property
    @pulumi.getter(name="defaultContentKeyPolicyName")
    def default_content_key_policy_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the default ContentKeyPolicy used by this Streaming Locator.
        """
        return pulumi.get(self, "default_content_key_policy_name")

    @default_content_key_policy_name.setter
    def default_content_key_policy_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_content_key_policy_name", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The end time of the Streaming Locator.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of asset or account filters which apply to this streaming locator
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The start time of the Streaming Locator.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter(name="streamingLocatorId")
    def streaming_locator_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The StreamingLocatorId of the Streaming Locator.
        """
        return pulumi.get(self, "streaming_locator_id")

    @streaming_locator_id.setter
    def streaming_locator_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "streaming_locator_id", value)

    @property
    @pulumi.getter(name="streamingLocatorName")
    def streaming_locator_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Streaming Locator name.
        """
        return pulumi.get(self, "streaming_locator_name")

    @streaming_locator_name.setter
    def streaming_locator_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "streaming_locator_name", value)


@pulumi.type_token("azure-native:media:StreamingLocator")
class StreamingLocator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 alternative_media_id: Optional[pulumi.Input[builtins.str]] = None,
                 asset_name: Optional[pulumi.Input[builtins.str]] = None,
                 content_keys: Optional[pulumi.Input[Sequence[pulumi.Input[Union['StreamingLocatorContentKeyArgs', 'StreamingLocatorContentKeyArgsDict']]]]] = None,
                 default_content_key_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 end_time: Optional[pulumi.Input[builtins.str]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 start_time: Optional[pulumi.Input[builtins.str]] = None,
                 streaming_locator_id: Optional[pulumi.Input[builtins.str]] = None,
                 streaming_locator_name: Optional[pulumi.Input[builtins.str]] = None,
                 streaming_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A Streaming Locator resource

        Uses Azure REST API version 2023-01-01. In version 2.x of the Azure Native provider, it used API version 2023-01-01.

        Other available API versions: 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-06-01, 2021-11-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] account_name: The Media Services account name.
        :param pulumi.Input[builtins.str] alternative_media_id: Alternative Media ID of this Streaming Locator
        :param pulumi.Input[builtins.str] asset_name: Asset Name
        :param pulumi.Input[Sequence[pulumi.Input[Union['StreamingLocatorContentKeyArgs', 'StreamingLocatorContentKeyArgsDict']]]] content_keys: The ContentKeys used by this Streaming Locator.
        :param pulumi.Input[builtins.str] default_content_key_policy_name: Name of the default ContentKeyPolicy used by this Streaming Locator.
        :param pulumi.Input[builtins.str] end_time: The end time of the Streaming Locator.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] filters: A list of asset or account filters which apply to this streaming locator
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[builtins.str] start_time: The start time of the Streaming Locator.
        :param pulumi.Input[builtins.str] streaming_locator_id: The StreamingLocatorId of the Streaming Locator.
        :param pulumi.Input[builtins.str] streaming_locator_name: The Streaming Locator name.
        :param pulumi.Input[builtins.str] streaming_policy_name: Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamingLocatorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Streaming Locator resource

        Uses Azure REST API version 2023-01-01. In version 2.x of the Azure Native provider, it used API version 2023-01-01.

        Other available API versions: 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-06-01, 2021-11-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param StreamingLocatorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamingLocatorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 alternative_media_id: Optional[pulumi.Input[builtins.str]] = None,
                 asset_name: Optional[pulumi.Input[builtins.str]] = None,
                 content_keys: Optional[pulumi.Input[Sequence[pulumi.Input[Union['StreamingLocatorContentKeyArgs', 'StreamingLocatorContentKeyArgsDict']]]]] = None,
                 default_content_key_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 end_time: Optional[pulumi.Input[builtins.str]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 start_time: Optional[pulumi.Input[builtins.str]] = None,
                 streaming_locator_id: Optional[pulumi.Input[builtins.str]] = None,
                 streaming_locator_name: Optional[pulumi.Input[builtins.str]] = None,
                 streaming_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StreamingLocatorArgs.__new__(StreamingLocatorArgs)

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["alternative_media_id"] = alternative_media_id
            if asset_name is None and not opts.urn:
                raise TypeError("Missing required property 'asset_name'")
            __props__.__dict__["asset_name"] = asset_name
            __props__.__dict__["content_keys"] = content_keys
            __props__.__dict__["default_content_key_policy_name"] = default_content_key_policy_name
            __props__.__dict__["end_time"] = end_time
            __props__.__dict__["filters"] = filters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["start_time"] = start_time
            __props__.__dict__["streaming_locator_id"] = streaming_locator_id
            __props__.__dict__["streaming_locator_name"] = streaming_locator_name
            if streaming_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'streaming_policy_name'")
            __props__.__dict__["streaming_policy_name"] = streaming_policy_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["created"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:media/v20180330preview:StreamingLocator"), pulumi.Alias(type_="azure-native:media/v20180601preview:StreamingLocator"), pulumi.Alias(type_="azure-native:media/v20180701:StreamingLocator"), pulumi.Alias(type_="azure-native:media/v20200501:StreamingLocator"), pulumi.Alias(type_="azure-native:media/v20210601:StreamingLocator"), pulumi.Alias(type_="azure-native:media/v20211101:StreamingLocator"), pulumi.Alias(type_="azure-native:media/v20220801:StreamingLocator"), pulumi.Alias(type_="azure-native:media/v20230101:StreamingLocator")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(StreamingLocator, __self__).__init__(
            'azure-native:media:StreamingLocator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'StreamingLocator':
        """
        Get an existing StreamingLocator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StreamingLocatorArgs.__new__(StreamingLocatorArgs)

        __props__.__dict__["alternative_media_id"] = None
        __props__.__dict__["asset_name"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["content_keys"] = None
        __props__.__dict__["created"] = None
        __props__.__dict__["default_content_key_policy_name"] = None
        __props__.__dict__["end_time"] = None
        __props__.__dict__["filters"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["start_time"] = None
        __props__.__dict__["streaming_locator_id"] = None
        __props__.__dict__["streaming_policy_name"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return StreamingLocator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alternativeMediaId")
    def alternative_media_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Alternative Media ID of this Streaming Locator
        """
        return pulumi.get(self, "alternative_media_id")

    @property
    @pulumi.getter(name="assetName")
    def asset_name(self) -> pulumi.Output[builtins.str]:
        """
        Asset Name
        """
        return pulumi.get(self, "asset_name")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="contentKeys")
    def content_keys(self) -> pulumi.Output[Optional[Sequence['outputs.StreamingLocatorContentKeyResponse']]]:
        """
        The ContentKeys used by this Streaming Locator.
        """
        return pulumi.get(self, "content_keys")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[builtins.str]:
        """
        The creation time of the Streaming Locator.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="defaultContentKeyPolicyName")
    def default_content_key_policy_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Name of the default ContentKeyPolicy used by this Streaming Locator.
        """
        return pulumi.get(self, "default_content_key_policy_name")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The end time of the Streaming Locator.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        A list of asset or account filters which apply to this streaming locator
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The start time of the Streaming Locator.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="streamingLocatorId")
    def streaming_locator_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The StreamingLocatorId of the Streaming Locator.
        """
        return pulumi.get(self, "streaming_locator_id")

    @property
    @pulumi.getter(name="streamingPolicyName")
    def streaming_policy_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
        """
        return pulumi.get(self, "streaming_policy_name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system metadata relating to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

