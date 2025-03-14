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
from ._enums import *

__all__ = [
    'FilteringTagArgs',
    'FilteringTagArgsDict',
    'IdentityPropertiesArgs',
    'IdentityPropertiesArgsDict',
    'LogRulesArgs',
    'LogRulesArgsDict',
    'LogzOrganizationPropertiesArgs',
    'LogzOrganizationPropertiesArgsDict',
    'MetricRulesArgs',
    'MetricRulesArgsDict',
    'MetricsTagRulesPropertiesArgs',
    'MetricsTagRulesPropertiesArgsDict',
    'MonitorPropertiesArgs',
    'MonitorPropertiesArgsDict',
    'MonitoringTagRulesPropertiesArgs',
    'MonitoringTagRulesPropertiesArgsDict',
    'PlanDataArgs',
    'PlanDataArgsDict',
    'UserInfoArgs',
    'UserInfoArgsDict',
]

MYPY = False

if not MYPY:
    class FilteringTagArgsDict(TypedDict):
        """
        The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them from being monitored.
        """
        action: NotRequired[pulumi.Input[Union[str, 'TagAction']]]
        """
        Valid actions for a filtering tag. Exclusion takes priority over inclusion.
        """
        name: NotRequired[pulumi.Input[str]]
        """
        The name (also known as the key) of the tag.
        """
        value: NotRequired[pulumi.Input[str]]
        """
        The value of the tag.
        """
elif False:
    FilteringTagArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FilteringTagArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[Union[str, 'TagAction']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them from being monitored.
        :param pulumi.Input[Union[str, 'TagAction']] action: Valid actions for a filtering tag. Exclusion takes priority over inclusion.
        :param pulumi.Input[str] name: The name (also known as the key) of the tag.
        :param pulumi.Input[str] value: The value of the tag.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[Union[str, 'TagAction']]]:
        """
        Valid actions for a filtering tag. Exclusion takes priority over inclusion.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[Union[str, 'TagAction']]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name (also known as the key) of the tag.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


if not MYPY:
    class IdentityPropertiesArgsDict(TypedDict):
        type: NotRequired[pulumi.Input[Union[str, 'ManagedIdentityTypes']]]
elif False:
    IdentityPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IdentityPropertiesArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[Union[str, 'ManagedIdentityTypes']]] = None):
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[str, 'ManagedIdentityTypes']]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[str, 'ManagedIdentityTypes']]]):
        pulumi.set(self, "type", value)


if not MYPY:
    class LogRulesArgsDict(TypedDict):
        """
        Set of rules for sending logs for the Monitor resource.
        """
        filtering_tags: NotRequired[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgsDict']]]]
        """
        List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        """
        send_aad_logs: NotRequired[pulumi.Input[bool]]
        """
        Flag specifying if AAD logs should be sent for the Monitor resource.
        """
        send_activity_logs: NotRequired[pulumi.Input[bool]]
        """
        Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
        """
        send_subscription_logs: NotRequired[pulumi.Input[bool]]
        """
        Flag specifying if subscription logs should be sent for the Monitor resource.
        """
elif False:
    LogRulesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class LogRulesArgs:
    def __init__(__self__, *,
                 filtering_tags: Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]] = None,
                 send_aad_logs: Optional[pulumi.Input[bool]] = None,
                 send_activity_logs: Optional[pulumi.Input[bool]] = None,
                 send_subscription_logs: Optional[pulumi.Input[bool]] = None):
        """
        Set of rules for sending logs for the Monitor resource.
        :param pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]] filtering_tags: List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        :param pulumi.Input[bool] send_aad_logs: Flag specifying if AAD logs should be sent for the Monitor resource.
        :param pulumi.Input[bool] send_activity_logs: Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
        :param pulumi.Input[bool] send_subscription_logs: Flag specifying if subscription logs should be sent for the Monitor resource.
        """
        if filtering_tags is not None:
            pulumi.set(__self__, "filtering_tags", filtering_tags)
        if send_aad_logs is not None:
            pulumi.set(__self__, "send_aad_logs", send_aad_logs)
        if send_activity_logs is not None:
            pulumi.set(__self__, "send_activity_logs", send_activity_logs)
        if send_subscription_logs is not None:
            pulumi.set(__self__, "send_subscription_logs", send_subscription_logs)

    @property
    @pulumi.getter(name="filteringTags")
    def filtering_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]]:
        """
        List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        """
        return pulumi.get(self, "filtering_tags")

    @filtering_tags.setter
    def filtering_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]]):
        pulumi.set(self, "filtering_tags", value)

    @property
    @pulumi.getter(name="sendAadLogs")
    def send_aad_logs(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag specifying if AAD logs should be sent for the Monitor resource.
        """
        return pulumi.get(self, "send_aad_logs")

    @send_aad_logs.setter
    def send_aad_logs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_aad_logs", value)

    @property
    @pulumi.getter(name="sendActivityLogs")
    def send_activity_logs(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
        """
        return pulumi.get(self, "send_activity_logs")

    @send_activity_logs.setter
    def send_activity_logs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_activity_logs", value)

    @property
    @pulumi.getter(name="sendSubscriptionLogs")
    def send_subscription_logs(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag specifying if subscription logs should be sent for the Monitor resource.
        """
        return pulumi.get(self, "send_subscription_logs")

    @send_subscription_logs.setter
    def send_subscription_logs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_subscription_logs", value)


if not MYPY:
    class LogzOrganizationPropertiesArgsDict(TypedDict):
        company_name: NotRequired[pulumi.Input[str]]
        """
        Name of the Logz organization.
        """
        enterprise_app_id: NotRequired[pulumi.Input[str]]
        """
        The Id of the Enterprise App used for Single sign on.
        """
        single_sign_on_url: NotRequired[pulumi.Input[str]]
        """
        The login URL specific to this Logz Organization.
        """
elif False:
    LogzOrganizationPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class LogzOrganizationPropertiesArgs:
    def __init__(__self__, *,
                 company_name: Optional[pulumi.Input[str]] = None,
                 enterprise_app_id: Optional[pulumi.Input[str]] = None,
                 single_sign_on_url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] company_name: Name of the Logz organization.
        :param pulumi.Input[str] enterprise_app_id: The Id of the Enterprise App used for Single sign on.
        :param pulumi.Input[str] single_sign_on_url: The login URL specific to this Logz Organization.
        """
        if company_name is not None:
            pulumi.set(__self__, "company_name", company_name)
        if enterprise_app_id is not None:
            pulumi.set(__self__, "enterprise_app_id", enterprise_app_id)
        if single_sign_on_url is not None:
            pulumi.set(__self__, "single_sign_on_url", single_sign_on_url)

    @property
    @pulumi.getter(name="companyName")
    def company_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Logz organization.
        """
        return pulumi.get(self, "company_name")

    @company_name.setter
    def company_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "company_name", value)

    @property
    @pulumi.getter(name="enterpriseAppId")
    def enterprise_app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of the Enterprise App used for Single sign on.
        """
        return pulumi.get(self, "enterprise_app_id")

    @enterprise_app_id.setter
    def enterprise_app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enterprise_app_id", value)

    @property
    @pulumi.getter(name="singleSignOnUrl")
    def single_sign_on_url(self) -> Optional[pulumi.Input[str]]:
        """
        The login URL specific to this Logz Organization.
        """
        return pulumi.get(self, "single_sign_on_url")

    @single_sign_on_url.setter
    def single_sign_on_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "single_sign_on_url", value)


if not MYPY:
    class MetricRulesArgsDict(TypedDict):
        """
        Set of rules for sending metrics for the Monitor resource.
        """
        filtering_tags: NotRequired[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgsDict']]]]
        """
        List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        """
        subscription_id: NotRequired[pulumi.Input[str]]
        """
        Subscription Id for which filtering tags are applicable
        """
elif False:
    MetricRulesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MetricRulesArgs:
    def __init__(__self__, *,
                 filtering_tags: Optional[pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None):
        """
        Set of rules for sending metrics for the Monitor resource.
        :param pulumi.Input[Sequence[pulumi.Input['FilteringTagArgs']]] filtering_tags: List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        :param pulumi.Input[str] subscription_id: Subscription Id for which filtering tags are applicable
        """
        if filtering_tags is not None:
            pulumi.set(__self__, "filtering_tags", filtering_tags)
        if subscription_id is not None:
            pulumi.set(__self__, "subscription_id", subscription_id)

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

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subscription Id for which filtering tags are applicable
        """
        return pulumi.get(self, "subscription_id")

    @subscription_id.setter
    def subscription_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription_id", value)


if not MYPY:
    class MetricsTagRulesPropertiesArgsDict(TypedDict):
        """
        Definition of the properties for a TagRules resource.
        """
        metric_rules: NotRequired[pulumi.Input[Sequence[pulumi.Input['MetricRulesArgsDict']]]]
        send_metrics: NotRequired[pulumi.Input[bool]]
        """
        Flag specifying if metrics from Azure resources should be sent for the Monitor resource.
        """
elif False:
    MetricsTagRulesPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MetricsTagRulesPropertiesArgs:
    def __init__(__self__, *,
                 metric_rules: Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesArgs']]]] = None,
                 send_metrics: Optional[pulumi.Input[bool]] = None):
        """
        Definition of the properties for a TagRules resource.
        :param pulumi.Input[bool] send_metrics: Flag specifying if metrics from Azure resources should be sent for the Monitor resource.
        """
        if metric_rules is not None:
            pulumi.set(__self__, "metric_rules", metric_rules)
        if send_metrics is not None:
            pulumi.set(__self__, "send_metrics", send_metrics)

    @property
    @pulumi.getter(name="metricRules")
    def metric_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesArgs']]]]:
        return pulumi.get(self, "metric_rules")

    @metric_rules.setter
    def metric_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesArgs']]]]):
        pulumi.set(self, "metric_rules", value)

    @property
    @pulumi.getter(name="sendMetrics")
    def send_metrics(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag specifying if metrics from Azure resources should be sent for the Monitor resource.
        """
        return pulumi.get(self, "send_metrics")

    @send_metrics.setter
    def send_metrics(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_metrics", value)


if not MYPY:
    class MonitorPropertiesArgsDict(TypedDict):
        """
        Properties specific to the monitor resource.
        """
        logz_organization_properties: NotRequired[pulumi.Input['LogzOrganizationPropertiesArgsDict']]
        marketplace_subscription_status: NotRequired[pulumi.Input[Union[str, 'MarketplaceSubscriptionStatus']]]
        """
        Flag specifying the Marketplace Subscription Status of the resource. If payment is not made in time, the resource will go in Suspended state.
        """
        monitoring_status: NotRequired[pulumi.Input[Union[str, 'MonitoringStatus']]]
        """
        Flag specifying if the resource monitoring is enabled or disabled.
        """
        plan_data: NotRequired[pulumi.Input['PlanDataArgsDict']]
        user_info: NotRequired[pulumi.Input['UserInfoArgsDict']]
elif False:
    MonitorPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MonitorPropertiesArgs:
    def __init__(__self__, *,
                 logz_organization_properties: Optional[pulumi.Input['LogzOrganizationPropertiesArgs']] = None,
                 marketplace_subscription_status: Optional[pulumi.Input[Union[str, 'MarketplaceSubscriptionStatus']]] = None,
                 monitoring_status: Optional[pulumi.Input[Union[str, 'MonitoringStatus']]] = None,
                 plan_data: Optional[pulumi.Input['PlanDataArgs']] = None,
                 user_info: Optional[pulumi.Input['UserInfoArgs']] = None):
        """
        Properties specific to the monitor resource.
        :param pulumi.Input[Union[str, 'MarketplaceSubscriptionStatus']] marketplace_subscription_status: Flag specifying the Marketplace Subscription Status of the resource. If payment is not made in time, the resource will go in Suspended state.
        :param pulumi.Input[Union[str, 'MonitoringStatus']] monitoring_status: Flag specifying if the resource monitoring is enabled or disabled.
        """
        if logz_organization_properties is not None:
            pulumi.set(__self__, "logz_organization_properties", logz_organization_properties)
        if marketplace_subscription_status is not None:
            pulumi.set(__self__, "marketplace_subscription_status", marketplace_subscription_status)
        if monitoring_status is not None:
            pulumi.set(__self__, "monitoring_status", monitoring_status)
        if plan_data is not None:
            pulumi.set(__self__, "plan_data", plan_data)
        if user_info is not None:
            pulumi.set(__self__, "user_info", user_info)

    @property
    @pulumi.getter(name="logzOrganizationProperties")
    def logz_organization_properties(self) -> Optional[pulumi.Input['LogzOrganizationPropertiesArgs']]:
        return pulumi.get(self, "logz_organization_properties")

    @logz_organization_properties.setter
    def logz_organization_properties(self, value: Optional[pulumi.Input['LogzOrganizationPropertiesArgs']]):
        pulumi.set(self, "logz_organization_properties", value)

    @property
    @pulumi.getter(name="marketplaceSubscriptionStatus")
    def marketplace_subscription_status(self) -> Optional[pulumi.Input[Union[str, 'MarketplaceSubscriptionStatus']]]:
        """
        Flag specifying the Marketplace Subscription Status of the resource. If payment is not made in time, the resource will go in Suspended state.
        """
        return pulumi.get(self, "marketplace_subscription_status")

    @marketplace_subscription_status.setter
    def marketplace_subscription_status(self, value: Optional[pulumi.Input[Union[str, 'MarketplaceSubscriptionStatus']]]):
        pulumi.set(self, "marketplace_subscription_status", value)

    @property
    @pulumi.getter(name="monitoringStatus")
    def monitoring_status(self) -> Optional[pulumi.Input[Union[str, 'MonitoringStatus']]]:
        """
        Flag specifying if the resource monitoring is enabled or disabled.
        """
        return pulumi.get(self, "monitoring_status")

    @monitoring_status.setter
    def monitoring_status(self, value: Optional[pulumi.Input[Union[str, 'MonitoringStatus']]]):
        pulumi.set(self, "monitoring_status", value)

    @property
    @pulumi.getter(name="planData")
    def plan_data(self) -> Optional[pulumi.Input['PlanDataArgs']]:
        return pulumi.get(self, "plan_data")

    @plan_data.setter
    def plan_data(self, value: Optional[pulumi.Input['PlanDataArgs']]):
        pulumi.set(self, "plan_data", value)

    @property
    @pulumi.getter(name="userInfo")
    def user_info(self) -> Optional[pulumi.Input['UserInfoArgs']]:
        return pulumi.get(self, "user_info")

    @user_info.setter
    def user_info(self, value: Optional[pulumi.Input['UserInfoArgs']]):
        pulumi.set(self, "user_info", value)


if not MYPY:
    class MonitoringTagRulesPropertiesArgsDict(TypedDict):
        """
        Definition of the properties for a TagRules resource.
        """
        log_rules: NotRequired[pulumi.Input['LogRulesArgsDict']]
        """
        Set of rules for sending logs for the Monitor resource.
        """
elif False:
    MonitoringTagRulesPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MonitoringTagRulesPropertiesArgs:
    def __init__(__self__, *,
                 log_rules: Optional[pulumi.Input['LogRulesArgs']] = None):
        """
        Definition of the properties for a TagRules resource.
        :param pulumi.Input['LogRulesArgs'] log_rules: Set of rules for sending logs for the Monitor resource.
        """
        if log_rules is not None:
            pulumi.set(__self__, "log_rules", log_rules)

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


if not MYPY:
    class PlanDataArgsDict(TypedDict):
        billing_cycle: NotRequired[pulumi.Input[str]]
        """
        different billing cycles like MONTHLY/WEEKLY. this could be enum
        """
        effective_date: NotRequired[pulumi.Input[str]]
        """
        date when plan was applied
        """
        plan_details: NotRequired[pulumi.Input[str]]
        """
        plan id as published by Logz
        """
        usage_type: NotRequired[pulumi.Input[str]]
        """
        different usage type like PAYG/COMMITTED. this could be enum
        """
elif False:
    PlanDataArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class PlanDataArgs:
    def __init__(__self__, *,
                 billing_cycle: Optional[pulumi.Input[str]] = None,
                 effective_date: Optional[pulumi.Input[str]] = None,
                 plan_details: Optional[pulumi.Input[str]] = None,
                 usage_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] billing_cycle: different billing cycles like MONTHLY/WEEKLY. this could be enum
        :param pulumi.Input[str] effective_date: date when plan was applied
        :param pulumi.Input[str] plan_details: plan id as published by Logz
        :param pulumi.Input[str] usage_type: different usage type like PAYG/COMMITTED. this could be enum
        """
        if billing_cycle is not None:
            pulumi.set(__self__, "billing_cycle", billing_cycle)
        if effective_date is not None:
            pulumi.set(__self__, "effective_date", effective_date)
        if plan_details is not None:
            pulumi.set(__self__, "plan_details", plan_details)
        if usage_type is not None:
            pulumi.set(__self__, "usage_type", usage_type)

    @property
    @pulumi.getter(name="billingCycle")
    def billing_cycle(self) -> Optional[pulumi.Input[str]]:
        """
        different billing cycles like MONTHLY/WEEKLY. this could be enum
        """
        return pulumi.get(self, "billing_cycle")

    @billing_cycle.setter
    def billing_cycle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_cycle", value)

    @property
    @pulumi.getter(name="effectiveDate")
    def effective_date(self) -> Optional[pulumi.Input[str]]:
        """
        date when plan was applied
        """
        return pulumi.get(self, "effective_date")

    @effective_date.setter
    def effective_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "effective_date", value)

    @property
    @pulumi.getter(name="planDetails")
    def plan_details(self) -> Optional[pulumi.Input[str]]:
        """
        plan id as published by Logz
        """
        return pulumi.get(self, "plan_details")

    @plan_details.setter
    def plan_details(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_details", value)

    @property
    @pulumi.getter(name="usageType")
    def usage_type(self) -> Optional[pulumi.Input[str]]:
        """
        different usage type like PAYG/COMMITTED. this could be enum
        """
        return pulumi.get(self, "usage_type")

    @usage_type.setter
    def usage_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage_type", value)


if not MYPY:
    class UserInfoArgsDict(TypedDict):
        email_address: NotRequired[pulumi.Input[str]]
        """
        Email of the user used by Logz for contacting them if needed
        """
        first_name: NotRequired[pulumi.Input[str]]
        """
        First Name of the user
        """
        last_name: NotRequired[pulumi.Input[str]]
        """
        Last Name of the user
        """
        phone_number: NotRequired[pulumi.Input[str]]
        """
        Phone number of the user used by Logz for contacting them if needed
        """
elif False:
    UserInfoArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class UserInfoArgs:
    def __init__(__self__, *,
                 email_address: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 phone_number: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] email_address: Email of the user used by Logz for contacting them if needed
        :param pulumi.Input[str] first_name: First Name of the user
        :param pulumi.Input[str] last_name: Last Name of the user
        :param pulumi.Input[str] phone_number: Phone number of the user used by Logz for contacting them if needed
        """
        if email_address is not None:
            pulumi.set(__self__, "email_address", email_address)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if phone_number is not None:
            pulumi.set(__self__, "phone_number", phone_number)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the user used by Logz for contacting them if needed
        """
        return pulumi.get(self, "email_address")

    @email_address.setter
    def email_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_address", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        First Name of the user
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        Last Name of the user
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> Optional[pulumi.Input[str]]:
        """
        Phone number of the user used by Logz for contacting them if needed
        """
        return pulumi.get(self, "phone_number")

    @phone_number.setter
    def phone_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_number", value)


