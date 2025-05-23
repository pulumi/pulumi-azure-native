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
from ._enums import *

__all__ = [
    'DatadogOrganizationPropertiesArgs',
    'DatadogOrganizationPropertiesArgsDict',
    'FilteringTagArgs',
    'FilteringTagArgsDict',
    'IdentityPropertiesArgs',
    'IdentityPropertiesArgsDict',
    'LogRulesArgs',
    'LogRulesArgsDict',
    'MetricRulesArgs',
    'MetricRulesArgsDict',
    'MonitorPropertiesArgs',
    'MonitorPropertiesArgsDict',
    'MonitoredSubscriptionArgs',
    'MonitoredSubscriptionArgsDict',
    'MonitoringTagRulesPropertiesArgs',
    'MonitoringTagRulesPropertiesArgsDict',
    'ResourceSkuArgs',
    'ResourceSkuArgsDict',
    'SubscriptionListArgs',
    'SubscriptionListArgsDict',
    'UserInfoArgs',
    'UserInfoArgsDict',
]

MYPY = False

if not MYPY:
    class DatadogOrganizationPropertiesArgsDict(TypedDict):
        """
        Specify the Datadog organization name. In the case of linking to existing organizations, Id, ApiKey, and Applicationkey is required as well.
        """
        api_key: NotRequired[pulumi.Input[builtins.str]]
        """
        Api key associated to the Datadog organization.
        """
        application_key: NotRequired[pulumi.Input[builtins.str]]
        """
        Application key associated to the Datadog organization.
        """
        cspm: NotRequired[pulumi.Input[builtins.bool]]
        """
        The configuration which describes the state of cloud security posture management. This collects configuration information for all resources in a subscription and track conformance to industry benchmarks.
        """
        enterprise_app_id: NotRequired[pulumi.Input[builtins.str]]
        """
        The Id of the Enterprise App used for Single sign on.
        """
        id: NotRequired[pulumi.Input[builtins.str]]
        """
        Id of the Datadog organization.
        """
        linking_auth_code: NotRequired[pulumi.Input[builtins.str]]
        """
        The auth code used to linking to an existing Datadog organization.
        """
        linking_client_id: NotRequired[pulumi.Input[builtins.str]]
        """
        The client_id from an existing in exchange for an auth token to link organization.
        """
        name: NotRequired[pulumi.Input[builtins.str]]
        """
        Name of the Datadog organization.
        """
        redirect_uri: NotRequired[pulumi.Input[builtins.str]]
        """
        The redirect URI for linking.
        """
elif False:
    DatadogOrganizationPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DatadogOrganizationPropertiesArgs:
    def __init__(__self__, *,
                 api_key: Optional[pulumi.Input[builtins.str]] = None,
                 application_key: Optional[pulumi.Input[builtins.str]] = None,
                 cspm: Optional[pulumi.Input[builtins.bool]] = None,
                 enterprise_app_id: Optional[pulumi.Input[builtins.str]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 linking_auth_code: Optional[pulumi.Input[builtins.str]] = None,
                 linking_client_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 redirect_uri: Optional[pulumi.Input[builtins.str]] = None):
        """
        Specify the Datadog organization name. In the case of linking to existing organizations, Id, ApiKey, and Applicationkey is required as well.
        :param pulumi.Input[builtins.str] api_key: Api key associated to the Datadog organization.
        :param pulumi.Input[builtins.str] application_key: Application key associated to the Datadog organization.
        :param pulumi.Input[builtins.bool] cspm: The configuration which describes the state of cloud security posture management. This collects configuration information for all resources in a subscription and track conformance to industry benchmarks.
        :param pulumi.Input[builtins.str] enterprise_app_id: The Id of the Enterprise App used for Single sign on.
        :param pulumi.Input[builtins.str] id: Id of the Datadog organization.
        :param pulumi.Input[builtins.str] linking_auth_code: The auth code used to linking to an existing Datadog organization.
        :param pulumi.Input[builtins.str] linking_client_id: The client_id from an existing in exchange for an auth token to link organization.
        :param pulumi.Input[builtins.str] name: Name of the Datadog organization.
        :param pulumi.Input[builtins.str] redirect_uri: The redirect URI for linking.
        """
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if application_key is not None:
            pulumi.set(__self__, "application_key", application_key)
        if cspm is not None:
            pulumi.set(__self__, "cspm", cspm)
        if enterprise_app_id is not None:
            pulumi.set(__self__, "enterprise_app_id", enterprise_app_id)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if linking_auth_code is not None:
            pulumi.set(__self__, "linking_auth_code", linking_auth_code)
        if linking_client_id is not None:
            pulumi.set(__self__, "linking_client_id", linking_client_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if redirect_uri is not None:
            pulumi.set(__self__, "redirect_uri", redirect_uri)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Api key associated to the Datadog organization.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter(name="applicationKey")
    def application_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Application key associated to the Datadog organization.
        """
        return pulumi.get(self, "application_key")

    @application_key.setter
    def application_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "application_key", value)

    @property
    @pulumi.getter
    def cspm(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        The configuration which describes the state of cloud security posture management. This collects configuration information for all resources in a subscription and track conformance to industry benchmarks.
        """
        return pulumi.get(self, "cspm")

    @cspm.setter
    def cspm(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "cspm", value)

    @property
    @pulumi.getter(name="enterpriseAppId")
    def enterprise_app_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Id of the Enterprise App used for Single sign on.
        """
        return pulumi.get(self, "enterprise_app_id")

    @enterprise_app_id.setter
    def enterprise_app_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "enterprise_app_id", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Id of the Datadog organization.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="linkingAuthCode")
    def linking_auth_code(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The auth code used to linking to an existing Datadog organization.
        """
        return pulumi.get(self, "linking_auth_code")

    @linking_auth_code.setter
    def linking_auth_code(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "linking_auth_code", value)

    @property
    @pulumi.getter(name="linkingClientId")
    def linking_client_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The client_id from an existing in exchange for an auth token to link organization.
        """
        return pulumi.get(self, "linking_client_id")

    @linking_client_id.setter
    def linking_client_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "linking_client_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Datadog organization.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="redirectUri")
    def redirect_uri(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The redirect URI for linking.
        """
        return pulumi.get(self, "redirect_uri")

    @redirect_uri.setter
    def redirect_uri(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "redirect_uri", value)


if not MYPY:
    class FilteringTagArgsDict(TypedDict):
        """
        The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them from being monitored.
        """
        action: NotRequired[pulumi.Input[Union[builtins.str, 'TagAction']]]
        """
        Valid actions for a filtering tag. Exclusion takes priority over inclusion.
        """
        name: NotRequired[pulumi.Input[builtins.str]]
        """
        The name (also known as the key) of the tag.
        """
        value: NotRequired[pulumi.Input[builtins.str]]
        """
        The value of the tag.
        """
elif False:
    FilteringTagArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FilteringTagArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[Union[builtins.str, 'TagAction']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None):
        """
        The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them from being monitored.
        :param pulumi.Input[Union[builtins.str, 'TagAction']] action: Valid actions for a filtering tag. Exclusion takes priority over inclusion.
        :param pulumi.Input[builtins.str] name: The name (also known as the key) of the tag.
        :param pulumi.Input[builtins.str] value: The value of the tag.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[Union[builtins.str, 'TagAction']]]:
        """
        Valid actions for a filtering tag. Exclusion takes priority over inclusion.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[Union[builtins.str, 'TagAction']]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name (also known as the key) of the tag.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The value of the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "value", value)


if not MYPY:
    class IdentityPropertiesArgsDict(TypedDict):
        type: NotRequired[pulumi.Input[Union[builtins.str, 'ManagedIdentityTypes']]]
        """
        Specifies the identity type of the Datadog Monitor. At this time the only allowed value is 'SystemAssigned'.
        """
elif False:
    IdentityPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IdentityPropertiesArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[Union[builtins.str, 'ManagedIdentityTypes']]] = None):
        """
        :param pulumi.Input[Union[builtins.str, 'ManagedIdentityTypes']] type: Specifies the identity type of the Datadog Monitor. At this time the only allowed value is 'SystemAssigned'.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[builtins.str, 'ManagedIdentityTypes']]]:
        """
        Specifies the identity type of the Datadog Monitor. At this time the only allowed value is 'SystemAssigned'.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[builtins.str, 'ManagedIdentityTypes']]]):
        pulumi.set(self, "type", value)


if not MYPY:
    class LogRulesArgsDict(TypedDict):
        """
        Set of rules for sending logs for the Monitor resource.
        """
        filtering_tags: NotRequired[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgsDict']]]]
        """
        List of filtering tags to be used for capturing logs. This only takes effect if SendResourceLogs flag is enabled. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        """
        send_aad_logs: NotRequired[pulumi.Input[builtins.bool]]
        """
        Flag specifying if AAD logs should be sent for the Monitor resource.
        """
        send_resource_logs: NotRequired[pulumi.Input[builtins.bool]]
        """
        Flag specifying if Azure resource logs should be sent for the Monitor resource.
        """
        send_subscription_logs: NotRequired[pulumi.Input[builtins.bool]]
        """
        Flag specifying if Azure subscription logs should be sent for the Monitor resource.
        """
elif False:
    LogRulesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class LogRulesArgs:
    def __init__(__self__, *,
                 filtering_tags: Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]] = None,
                 send_aad_logs: Optional[pulumi.Input[builtins.bool]] = None,
                 send_resource_logs: Optional[pulumi.Input[builtins.bool]] = None,
                 send_subscription_logs: Optional[pulumi.Input[builtins.bool]] = None):
        """
        Set of rules for sending logs for the Monitor resource.
        :param pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]] filtering_tags: List of filtering tags to be used for capturing logs. This only takes effect if SendResourceLogs flag is enabled. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        :param pulumi.Input[builtins.bool] send_aad_logs: Flag specifying if AAD logs should be sent for the Monitor resource.
        :param pulumi.Input[builtins.bool] send_resource_logs: Flag specifying if Azure resource logs should be sent for the Monitor resource.
        :param pulumi.Input[builtins.bool] send_subscription_logs: Flag specifying if Azure subscription logs should be sent for the Monitor resource.
        """
        if filtering_tags is not None:
            pulumi.set(__self__, "filtering_tags", filtering_tags)
        if send_aad_logs is not None:
            pulumi.set(__self__, "send_aad_logs", send_aad_logs)
        if send_resource_logs is not None:
            pulumi.set(__self__, "send_resource_logs", send_resource_logs)
        if send_subscription_logs is not None:
            pulumi.set(__self__, "send_subscription_logs", send_subscription_logs)

    @property
    @pulumi.getter(name="filteringTags")
    def filtering_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]]:
        """
        List of filtering tags to be used for capturing logs. This only takes effect if SendResourceLogs flag is enabled. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        """
        return pulumi.get(self, "filtering_tags")

    @filtering_tags.setter
    def filtering_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]]):
        pulumi.set(self, "filtering_tags", value)

    @property
    @pulumi.getter(name="sendAadLogs")
    def send_aad_logs(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Flag specifying if AAD logs should be sent for the Monitor resource.
        """
        return pulumi.get(self, "send_aad_logs")

    @send_aad_logs.setter
    def send_aad_logs(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "send_aad_logs", value)

    @property
    @pulumi.getter(name="sendResourceLogs")
    def send_resource_logs(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Flag specifying if Azure resource logs should be sent for the Monitor resource.
        """
        return pulumi.get(self, "send_resource_logs")

    @send_resource_logs.setter
    def send_resource_logs(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "send_resource_logs", value)

    @property
    @pulumi.getter(name="sendSubscriptionLogs")
    def send_subscription_logs(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Flag specifying if Azure subscription logs should be sent for the Monitor resource.
        """
        return pulumi.get(self, "send_subscription_logs")

    @send_subscription_logs.setter
    def send_subscription_logs(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "send_subscription_logs", value)


if not MYPY:
    class MetricRulesArgsDict(TypedDict):
        """
        Set of rules for sending metrics for the Monitor resource.
        """
        filtering_tags: NotRequired[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgsDict']]]]
        """
        List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        """
elif False:
    MetricRulesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MetricRulesArgs:
    def __init__(__self__, *,
                 filtering_tags: Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]] = None):
        """
        Set of rules for sending metrics for the Monitor resource.
        :param pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]] filtering_tags: List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        """
        if filtering_tags is not None:
            pulumi.set(__self__, "filtering_tags", filtering_tags)

    @property
    @pulumi.getter(name="filteringTags")
    def filtering_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]]:
        """
        List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        """
        return pulumi.get(self, "filtering_tags")

    @filtering_tags.setter
    def filtering_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]]):
        pulumi.set(self, "filtering_tags", value)


if not MYPY:
    class MonitorPropertiesArgsDict(TypedDict):
        """
        Properties specific to the monitor resource.
        """
        datadog_organization_properties: NotRequired[pulumi.Input['DatadogOrganizationPropertiesArgsDict']]
        """
        Specify the Datadog organization name. In the case of linking to existing organizations, Id, ApiKey, and Applicationkey is required as well.
        """
        monitoring_status: NotRequired[pulumi.Input[Union[builtins.str, 'MonitoringStatus']]]
        """
        Flag specifying if the resource monitoring is enabled or disabled.
        """
        user_info: NotRequired[pulumi.Input['UserInfoArgsDict']]
        """
        Includes name, email and optionally, phone number. User Information can't be null.
        """
elif False:
    MonitorPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MonitorPropertiesArgs:
    def __init__(__self__, *,
                 datadog_organization_properties: Optional[pulumi.Input['DatadogOrganizationPropertiesArgs']] = None,
                 monitoring_status: Optional[pulumi.Input[Union[builtins.str, 'MonitoringStatus']]] = None,
                 user_info: Optional[pulumi.Input['UserInfoArgs']] = None):
        """
        Properties specific to the monitor resource.
        :param pulumi.Input['DatadogOrganizationPropertiesArgs'] datadog_organization_properties: Specify the Datadog organization name. In the case of linking to existing organizations, Id, ApiKey, and Applicationkey is required as well.
        :param pulumi.Input[Union[builtins.str, 'MonitoringStatus']] monitoring_status: Flag specifying if the resource monitoring is enabled or disabled.
        :param pulumi.Input['UserInfoArgs'] user_info: Includes name, email and optionally, phone number. User Information can't be null.
        """
        if datadog_organization_properties is not None:
            pulumi.set(__self__, "datadog_organization_properties", datadog_organization_properties)
        if monitoring_status is not None:
            pulumi.set(__self__, "monitoring_status", monitoring_status)
        if user_info is not None:
            pulumi.set(__self__, "user_info", user_info)

    @property
    @pulumi.getter(name="datadogOrganizationProperties")
    def datadog_organization_properties(self) -> Optional[pulumi.Input['DatadogOrganizationPropertiesArgs']]:
        """
        Specify the Datadog organization name. In the case of linking to existing organizations, Id, ApiKey, and Applicationkey is required as well.
        """
        return pulumi.get(self, "datadog_organization_properties")

    @datadog_organization_properties.setter
    def datadog_organization_properties(self, value: Optional[pulumi.Input['DatadogOrganizationPropertiesArgs']]):
        pulumi.set(self, "datadog_organization_properties", value)

    @property
    @pulumi.getter(name="monitoringStatus")
    def monitoring_status(self) -> Optional[pulumi.Input[Union[builtins.str, 'MonitoringStatus']]]:
        """
        Flag specifying if the resource monitoring is enabled or disabled.
        """
        return pulumi.get(self, "monitoring_status")

    @monitoring_status.setter
    def monitoring_status(self, value: Optional[pulumi.Input[Union[builtins.str, 'MonitoringStatus']]]):
        pulumi.set(self, "monitoring_status", value)

    @property
    @pulumi.getter(name="userInfo")
    def user_info(self) -> Optional[pulumi.Input['UserInfoArgs']]:
        """
        Includes name, email and optionally, phone number. User Information can't be null.
        """
        return pulumi.get(self, "user_info")

    @user_info.setter
    def user_info(self, value: Optional[pulumi.Input['UserInfoArgs']]):
        pulumi.set(self, "user_info", value)


if not MYPY:
    class MonitoredSubscriptionArgsDict(TypedDict):
        """
        The list of subscriptions and it's monitoring status by current Datadog monitor.
        """
        error: NotRequired[pulumi.Input[builtins.str]]
        """
        The reason of not monitoring the subscription.
        """
        status: NotRequired[pulumi.Input[Union[builtins.str, 'Status']]]
        """
        The state of monitoring.
        """
        subscription_id: NotRequired[pulumi.Input[builtins.str]]
        """
        The subscriptionId to be monitored.
        """
        tag_rules: NotRequired[pulumi.Input['MonitoringTagRulesPropertiesArgsDict']]
        """
        Definition of the properties for a TagRules resource.
        """
elif False:
    MonitoredSubscriptionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MonitoredSubscriptionArgs:
    def __init__(__self__, *,
                 error: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[Union[builtins.str, 'Status']]] = None,
                 subscription_id: Optional[pulumi.Input[builtins.str]] = None,
                 tag_rules: Optional[pulumi.Input['MonitoringTagRulesPropertiesArgs']] = None):
        """
        The list of subscriptions and it's monitoring status by current Datadog monitor.
        :param pulumi.Input[builtins.str] error: The reason of not monitoring the subscription.
        :param pulumi.Input[Union[builtins.str, 'Status']] status: The state of monitoring.
        :param pulumi.Input[builtins.str] subscription_id: The subscriptionId to be monitored.
        :param pulumi.Input['MonitoringTagRulesPropertiesArgs'] tag_rules: Definition of the properties for a TagRules resource.
        """
        if error is not None:
            pulumi.set(__self__, "error", error)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subscription_id is not None:
            pulumi.set(__self__, "subscription_id", subscription_id)
        if tag_rules is not None:
            pulumi.set(__self__, "tag_rules", tag_rules)

    @property
    @pulumi.getter
    def error(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The reason of not monitoring the subscription.
        """
        return pulumi.get(self, "error")

    @error.setter
    def error(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "error", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[Union[builtins.str, 'Status']]]:
        """
        The state of monitoring.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[Union[builtins.str, 'Status']]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The subscriptionId to be monitored.
        """
        return pulumi.get(self, "subscription_id")

    @subscription_id.setter
    def subscription_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subscription_id", value)

    @property
    @pulumi.getter(name="tagRules")
    def tag_rules(self) -> Optional[pulumi.Input['MonitoringTagRulesPropertiesArgs']]:
        """
        Definition of the properties for a TagRules resource.
        """
        return pulumi.get(self, "tag_rules")

    @tag_rules.setter
    def tag_rules(self, value: Optional[pulumi.Input['MonitoringTagRulesPropertiesArgs']]):
        pulumi.set(self, "tag_rules", value)


if not MYPY:
    class MonitoringTagRulesPropertiesArgsDict(TypedDict):
        """
        Definition of the properties for a TagRules resource.
        """
        automuting: NotRequired[pulumi.Input[builtins.bool]]
        """
        Configuration to enable/disable auto-muting flag
        """
        custom_metrics: NotRequired[pulumi.Input[builtins.bool]]
        """
        Configuration to enable/disable custom metrics. If enabled, custom metrics from app insights will be sent.
        """
        log_rules: NotRequired[pulumi.Input['LogRulesArgsDict']]
        """
        Set of rules for sending logs for the Monitor resource.
        """
        metric_rules: NotRequired[pulumi.Input['MetricRulesArgsDict']]
        """
        Set of rules for sending metrics for the Monitor resource.
        """
elif False:
    MonitoringTagRulesPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MonitoringTagRulesPropertiesArgs:
    def __init__(__self__, *,
                 automuting: Optional[pulumi.Input[builtins.bool]] = None,
                 custom_metrics: Optional[pulumi.Input[builtins.bool]] = None,
                 log_rules: Optional[pulumi.Input['LogRulesArgs']] = None,
                 metric_rules: Optional[pulumi.Input['MetricRulesArgs']] = None):
        """
        Definition of the properties for a TagRules resource.
        :param pulumi.Input[builtins.bool] automuting: Configuration to enable/disable auto-muting flag
        :param pulumi.Input[builtins.bool] custom_metrics: Configuration to enable/disable custom metrics. If enabled, custom metrics from app insights will be sent.
        :param pulumi.Input['LogRulesArgs'] log_rules: Set of rules for sending logs for the Monitor resource.
        :param pulumi.Input['MetricRulesArgs'] metric_rules: Set of rules for sending metrics for the Monitor resource.
        """
        if automuting is not None:
            pulumi.set(__self__, "automuting", automuting)
        if custom_metrics is not None:
            pulumi.set(__self__, "custom_metrics", custom_metrics)
        if log_rules is not None:
            pulumi.set(__self__, "log_rules", log_rules)
        if metric_rules is not None:
            pulumi.set(__self__, "metric_rules", metric_rules)

    @property
    @pulumi.getter
    def automuting(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Configuration to enable/disable auto-muting flag
        """
        return pulumi.get(self, "automuting")

    @automuting.setter
    def automuting(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "automuting", value)

    @property
    @pulumi.getter(name="customMetrics")
    def custom_metrics(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Configuration to enable/disable custom metrics. If enabled, custom metrics from app insights will be sent.
        """
        return pulumi.get(self, "custom_metrics")

    @custom_metrics.setter
    def custom_metrics(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "custom_metrics", value)

    @property
    @pulumi.getter(name="logRules")
    def log_rules(self) -> Optional[pulumi.Input['LogRulesArgs']]:
        """
        Set of rules for sending logs for the Monitor resource.
        """
        return pulumi.get(self, "log_rules")

    @log_rules.setter
    def log_rules(self, value: Optional[pulumi.Input['LogRulesArgs']]):
        pulumi.set(self, "log_rules", value)

    @property
    @pulumi.getter(name="metricRules")
    def metric_rules(self) -> Optional[pulumi.Input['MetricRulesArgs']]:
        """
        Set of rules for sending metrics for the Monitor resource.
        """
        return pulumi.get(self, "metric_rules")

    @metric_rules.setter
    def metric_rules(self, value: Optional[pulumi.Input['MetricRulesArgs']]):
        pulumi.set(self, "metric_rules", value)


if not MYPY:
    class ResourceSkuArgsDict(TypedDict):
        name: pulumi.Input[builtins.str]
        """
        Name of the SKU in {PlanId} format. For Terraform, the only allowed value is 'Linked'.
        """
elif False:
    ResourceSkuArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ResourceSkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] name: Name of the SKU in {PlanId} format. For Terraform, the only allowed value is 'Linked'.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the SKU in {PlanId} format. For Terraform, the only allowed value is 'Linked'.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)


if not MYPY:
    class SubscriptionListArgsDict(TypedDict):
        """
        The request to update subscriptions needed to be monitored by the Datadog monitor resource.
        """
        monitored_subscription_list: NotRequired[pulumi.Input[Sequence[pulumi.Input['MonitoredSubscriptionArgsDict']]]]
        """
        List of subscriptions and the state of the monitoring.
        """
        operation: NotRequired[pulumi.Input[Union[builtins.str, 'Operation']]]
        """
        The operation for the patch on the resource.
        """
elif False:
    SubscriptionListArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubscriptionListArgs:
    def __init__(__self__, *,
                 monitored_subscription_list: Optional[pulumi.Input[Sequence[pulumi.Input['MonitoredSubscriptionArgs']]]] = None,
                 operation: Optional[pulumi.Input[Union[builtins.str, 'Operation']]] = None):
        """
        The request to update subscriptions needed to be monitored by the Datadog monitor resource.
        :param pulumi.Input[Sequence[pulumi.Input['MonitoredSubscriptionArgs']]] monitored_subscription_list: List of subscriptions and the state of the monitoring.
        :param pulumi.Input[Union[builtins.str, 'Operation']] operation: The operation for the patch on the resource.
        """
        if monitored_subscription_list is not None:
            pulumi.set(__self__, "monitored_subscription_list", monitored_subscription_list)
        if operation is not None:
            pulumi.set(__self__, "operation", operation)

    @property
    @pulumi.getter(name="monitoredSubscriptionList")
    def monitored_subscription_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MonitoredSubscriptionArgs']]]]:
        """
        List of subscriptions and the state of the monitoring.
        """
        return pulumi.get(self, "monitored_subscription_list")

    @monitored_subscription_list.setter
    def monitored_subscription_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MonitoredSubscriptionArgs']]]]):
        pulumi.set(self, "monitored_subscription_list", value)

    @property
    @pulumi.getter
    def operation(self) -> Optional[pulumi.Input[Union[builtins.str, 'Operation']]]:
        """
        The operation for the patch on the resource.
        """
        return pulumi.get(self, "operation")

    @operation.setter
    def operation(self, value: Optional[pulumi.Input[Union[builtins.str, 'Operation']]]):
        pulumi.set(self, "operation", value)


if not MYPY:
    class UserInfoArgsDict(TypedDict):
        """
        Includes name, email and optionally, phone number. User Information can't be null.
        """
        email_address: NotRequired[pulumi.Input[builtins.str]]
        """
        Email of the user used by Datadog for contacting them if needed
        """
        name: NotRequired[pulumi.Input[builtins.str]]
        """
        Name of the user
        """
        phone_number: NotRequired[pulumi.Input[builtins.str]]
        """
        Phone number of the user used by Datadog for contacting them if needed
        """
elif False:
    UserInfoArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class UserInfoArgs:
    def __init__(__self__, *,
                 email_address: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 phone_number: Optional[pulumi.Input[builtins.str]] = None):
        """
        Includes name, email and optionally, phone number. User Information can't be null.
        :param pulumi.Input[builtins.str] email_address: Email of the user used by Datadog for contacting them if needed
        :param pulumi.Input[builtins.str] name: Name of the user
        :param pulumi.Input[builtins.str] phone_number: Phone number of the user used by Datadog for contacting them if needed
        """
        if email_address is not None:
            pulumi.set(__self__, "email_address", email_address)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if phone_number is not None:
            pulumi.set(__self__, "phone_number", phone_number)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Email of the user used by Datadog for contacting them if needed
        """
        return pulumi.get(self, "email_address")

    @email_address.setter
    def email_address(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "email_address", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the user
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Phone number of the user used by Datadog for contacting them if needed
        """
        return pulumi.get(self, "phone_number")

    @phone_number.setter
    def phone_number(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "phone_number", value)


