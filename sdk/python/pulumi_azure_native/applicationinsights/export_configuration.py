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

__all__ = ['ExportConfigurationArgs', 'ExportConfiguration']

@pulumi.input_type
class ExportConfigurationArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 resource_name: pulumi.Input[builtins.str],
                 destination_account_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_address: Optional[pulumi.Input[builtins.str]] = None,
                 destination_storage_location_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_storage_subscription_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_type: Optional[pulumi.Input[builtins.str]] = None,
                 export_id: Optional[pulumi.Input[builtins.str]] = None,
                 is_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 notification_queue_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 notification_queue_uri: Optional[pulumi.Input[builtins.str]] = None,
                 record_types: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ExportConfiguration resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_name: The name of the Application Insights component resource.
        :param pulumi.Input[builtins.str] destination_account_id: The name of destination storage account.
        :param pulumi.Input[builtins.str] destination_address: The SAS URL for the destination storage container. It must grant write permission.
        :param pulumi.Input[builtins.str] destination_storage_location_id: The location ID of the destination storage container.
        :param pulumi.Input[builtins.str] destination_storage_subscription_id: The subscription ID of the destination storage container.
        :param pulumi.Input[builtins.str] destination_type: The Continuous Export destination type. This has to be 'Blob'.
        :param pulumi.Input[builtins.str] export_id: The Continuous Export configuration ID. This is unique within a Application Insights component.
        :param pulumi.Input[builtins.str] is_enabled: Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'.
        :param pulumi.Input[builtins.str] notification_queue_enabled: Deprecated
        :param pulumi.Input[builtins.str] notification_queue_uri: Deprecated
        :param pulumi.Input[builtins.str] record_types: The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "resource_name", resource_name)
        if destination_account_id is not None:
            pulumi.set(__self__, "destination_account_id", destination_account_id)
        if destination_address is not None:
            pulumi.set(__self__, "destination_address", destination_address)
        if destination_storage_location_id is not None:
            pulumi.set(__self__, "destination_storage_location_id", destination_storage_location_id)
        if destination_storage_subscription_id is not None:
            pulumi.set(__self__, "destination_storage_subscription_id", destination_storage_subscription_id)
        if destination_type is not None:
            pulumi.set(__self__, "destination_type", destination_type)
        if export_id is not None:
            pulumi.set(__self__, "export_id", export_id)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if notification_queue_enabled is not None:
            pulumi.set(__self__, "notification_queue_enabled", notification_queue_enabled)
        if notification_queue_uri is not None:
            pulumi.set(__self__, "notification_queue_uri", notification_queue_uri)
        if record_types is not None:
            pulumi.set(__self__, "record_types", record_types)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Application Insights component resource.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter(name="destinationAccountId")
    def destination_account_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of destination storage account.
        """
        return pulumi.get(self, "destination_account_id")

    @destination_account_id.setter
    def destination_account_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_account_id", value)

    @property
    @pulumi.getter(name="destinationAddress")
    def destination_address(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The SAS URL for the destination storage container. It must grant write permission.
        """
        return pulumi.get(self, "destination_address")

    @destination_address.setter
    def destination_address(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_address", value)

    @property
    @pulumi.getter(name="destinationStorageLocationId")
    def destination_storage_location_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The location ID of the destination storage container.
        """
        return pulumi.get(self, "destination_storage_location_id")

    @destination_storage_location_id.setter
    def destination_storage_location_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_storage_location_id", value)

    @property
    @pulumi.getter(name="destinationStorageSubscriptionId")
    def destination_storage_subscription_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The subscription ID of the destination storage container.
        """
        return pulumi.get(self, "destination_storage_subscription_id")

    @destination_storage_subscription_id.setter
    def destination_storage_subscription_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_storage_subscription_id", value)

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Continuous Export destination type. This has to be 'Blob'.
        """
        return pulumi.get(self, "destination_type")

    @destination_type.setter
    def destination_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "destination_type", value)

    @property
    @pulumi.getter(name="exportId")
    def export_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Continuous Export configuration ID. This is unique within a Application Insights component.
        """
        return pulumi.get(self, "export_id")

    @export_id.setter
    def export_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "export_id", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="notificationQueueEnabled")
    def notification_queue_enabled(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Deprecated
        """
        return pulumi.get(self, "notification_queue_enabled")

    @notification_queue_enabled.setter
    def notification_queue_enabled(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "notification_queue_enabled", value)

    @property
    @pulumi.getter(name="notificationQueueUri")
    def notification_queue_uri(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Deprecated
        """
        return pulumi.get(self, "notification_queue_uri")

    @notification_queue_uri.setter
    def notification_queue_uri(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "notification_queue_uri", value)

    @property
    @pulumi.getter(name="recordTypes")
    def record_types(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
        """
        return pulumi.get(self, "record_types")

    @record_types.setter
    def record_types(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "record_types", value)


@pulumi.type_token("azure-native:applicationinsights:ExportConfiguration")
class ExportConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_account_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_address: Optional[pulumi.Input[builtins.str]] = None,
                 destination_storage_location_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_storage_subscription_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_type: Optional[pulumi.Input[builtins.str]] = None,
                 export_id: Optional[pulumi.Input[builtins.str]] = None,
                 is_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 notification_queue_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 notification_queue_uri: Optional[pulumi.Input[builtins.str]] = None,
                 record_types: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Properties that define a Continuous Export configuration.

        Uses Azure REST API version 2015-05-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] destination_account_id: The name of destination storage account.
        :param pulumi.Input[builtins.str] destination_address: The SAS URL for the destination storage container. It must grant write permission.
        :param pulumi.Input[builtins.str] destination_storage_location_id: The location ID of the destination storage container.
        :param pulumi.Input[builtins.str] destination_storage_subscription_id: The subscription ID of the destination storage container.
        :param pulumi.Input[builtins.str] destination_type: The Continuous Export destination type. This has to be 'Blob'.
        :param pulumi.Input[builtins.str] export_id: The Continuous Export configuration ID. This is unique within a Application Insights component.
        :param pulumi.Input[builtins.str] is_enabled: Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'.
        :param pulumi.Input[builtins.str] notification_queue_enabled: Deprecated
        :param pulumi.Input[builtins.str] notification_queue_uri: Deprecated
        :param pulumi.Input[builtins.str] record_types: The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_name_: The name of the Application Insights component resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExportConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Properties that define a Continuous Export configuration.

        Uses Azure REST API version 2015-05-01.

        :param str resource_name: The name of the resource.
        :param ExportConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExportConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_account_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_address: Optional[pulumi.Input[builtins.str]] = None,
                 destination_storage_location_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_storage_subscription_id: Optional[pulumi.Input[builtins.str]] = None,
                 destination_type: Optional[pulumi.Input[builtins.str]] = None,
                 export_id: Optional[pulumi.Input[builtins.str]] = None,
                 is_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 notification_queue_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 notification_queue_uri: Optional[pulumi.Input[builtins.str]] = None,
                 record_types: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExportConfigurationArgs.__new__(ExportConfigurationArgs)

            __props__.__dict__["destination_account_id"] = destination_account_id
            __props__.__dict__["destination_address"] = destination_address
            __props__.__dict__["destination_storage_location_id"] = destination_storage_location_id
            __props__.__dict__["destination_storage_subscription_id"] = destination_storage_subscription_id
            __props__.__dict__["destination_type"] = destination_type
            __props__.__dict__["export_id"] = export_id
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["notification_queue_enabled"] = notification_queue_enabled
            __props__.__dict__["notification_queue_uri"] = notification_queue_uri
            __props__.__dict__["record_types"] = record_types
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__.__dict__["resource_name"] = resource_name_
            __props__.__dict__["application_name"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["container_name"] = None
            __props__.__dict__["export_status"] = None
            __props__.__dict__["instrumentation_key"] = None
            __props__.__dict__["is_user_enabled"] = None
            __props__.__dict__["last_gap_time"] = None
            __props__.__dict__["last_success_time"] = None
            __props__.__dict__["last_user_update"] = None
            __props__.__dict__["permanent_error_reason"] = None
            __props__.__dict__["resource_group"] = None
            __props__.__dict__["storage_name"] = None
            __props__.__dict__["subscription_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:applicationinsights/v20150501:ExportConfiguration"), pulumi.Alias(type_="azure-native:insights/v20150501:ExportConfiguration"), pulumi.Alias(type_="azure-native:insights:ExportConfiguration")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ExportConfiguration, __self__).__init__(
            'azure-native:applicationinsights:ExportConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExportConfiguration':
        """
        Get an existing ExportConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ExportConfigurationArgs.__new__(ExportConfigurationArgs)

        __props__.__dict__["application_name"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["container_name"] = None
        __props__.__dict__["destination_account_id"] = None
        __props__.__dict__["destination_storage_location_id"] = None
        __props__.__dict__["destination_storage_subscription_id"] = None
        __props__.__dict__["destination_type"] = None
        __props__.__dict__["export_id"] = None
        __props__.__dict__["export_status"] = None
        __props__.__dict__["instrumentation_key"] = None
        __props__.__dict__["is_user_enabled"] = None
        __props__.__dict__["last_gap_time"] = None
        __props__.__dict__["last_success_time"] = None
        __props__.__dict__["last_user_update"] = None
        __props__.__dict__["notification_queue_enabled"] = None
        __props__.__dict__["permanent_error_reason"] = None
        __props__.__dict__["record_types"] = None
        __props__.__dict__["resource_group"] = None
        __props__.__dict__["storage_name"] = None
        __props__.__dict__["subscription_id"] = None
        return ExportConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the Application Insights component.
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the destination storage container.
        """
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter(name="destinationAccountId")
    def destination_account_id(self) -> pulumi.Output[builtins.str]:
        """
        The name of destination account.
        """
        return pulumi.get(self, "destination_account_id")

    @property
    @pulumi.getter(name="destinationStorageLocationId")
    def destination_storage_location_id(self) -> pulumi.Output[builtins.str]:
        """
        The destination account location ID.
        """
        return pulumi.get(self, "destination_storage_location_id")

    @property
    @pulumi.getter(name="destinationStorageSubscriptionId")
    def destination_storage_subscription_id(self) -> pulumi.Output[builtins.str]:
        """
        The destination storage account subscription ID.
        """
        return pulumi.get(self, "destination_storage_subscription_id")

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> pulumi.Output[builtins.str]:
        """
        The destination type.
        """
        return pulumi.get(self, "destination_type")

    @property
    @pulumi.getter(name="exportId")
    def export_id(self) -> pulumi.Output[builtins.str]:
        """
        The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.
        """
        return pulumi.get(self, "export_id")

    @property
    @pulumi.getter(name="exportStatus")
    def export_status(self) -> pulumi.Output[builtins.str]:
        """
        This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
        """
        return pulumi.get(self, "export_status")

    @property
    @pulumi.getter(name="instrumentationKey")
    def instrumentation_key(self) -> pulumi.Output[builtins.str]:
        """
        The instrumentation key of the Application Insights component.
        """
        return pulumi.get(self, "instrumentation_key")

    @property
    @pulumi.getter(name="isUserEnabled")
    def is_user_enabled(self) -> pulumi.Output[builtins.str]:
        """
        This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
        """
        return pulumi.get(self, "is_user_enabled")

    @property
    @pulumi.getter(name="lastGapTime")
    def last_gap_time(self) -> pulumi.Output[builtins.str]:
        """
        The last time the Continuous Export configuration started failing.
        """
        return pulumi.get(self, "last_gap_time")

    @property
    @pulumi.getter(name="lastSuccessTime")
    def last_success_time(self) -> pulumi.Output[builtins.str]:
        """
        The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
        """
        return pulumi.get(self, "last_success_time")

    @property
    @pulumi.getter(name="lastUserUpdate")
    def last_user_update(self) -> pulumi.Output[builtins.str]:
        """
        Last time the Continuous Export configuration was updated.
        """
        return pulumi.get(self, "last_user_update")

    @property
    @pulumi.getter(name="notificationQueueEnabled")
    def notification_queue_enabled(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Deprecated
        """
        return pulumi.get(self, "notification_queue_enabled")

    @property
    @pulumi.getter(name="permanentErrorReason")
    def permanent_error_reason(self) -> pulumi.Output[builtins.str]:
        """
        This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
        """
        return pulumi.get(self, "permanent_error_reason")

    @property
    @pulumi.getter(name="recordTypes")
    def record_types(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
        """
        return pulumi.get(self, "record_types")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> pulumi.Output[builtins.str]:
        """
        The resource group of the Application Insights component.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="storageName")
    def storage_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the destination storage account.
        """
        return pulumi.get(self, "storage_name")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Output[builtins.str]:
        """
        The subscription of the Application Insights component.
        """
        return pulumi.get(self, "subscription_id")

