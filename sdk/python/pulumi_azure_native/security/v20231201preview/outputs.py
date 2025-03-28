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
from ._enums import *

__all__ = [
    'AutomationActionEventHubResponse',
    'AutomationActionLogicAppResponse',
    'AutomationActionWorkspaceResponse',
    'AutomationRuleSetResponse',
    'AutomationScopeResponse',
    'AutomationSourceResponse',
    'AutomationTriggeringRuleResponse',
    'NotificationsSourceAlertResponse',
    'NotificationsSourceAttackPathResponse',
    'SecurityContactPropertiesResponseNotificationsByRole',
]

@pulumi.output_type
class AutomationActionEventHubResponse(dict):
    """
    The target Event Hub to which event data will be exported. To learn more about Microsoft Defender for Cloud continuous export capabilities, visit https://aka.ms/ASCExportLearnMore
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionType":
            suggest = "action_type"
        elif key == "sasPolicyName":
            suggest = "sas_policy_name"
        elif key == "connectionString":
            suggest = "connection_string"
        elif key == "eventHubResourceId":
            suggest = "event_hub_resource_id"
        elif key == "isTrustedServiceEnabled":
            suggest = "is_trusted_service_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AutomationActionEventHubResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AutomationActionEventHubResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AutomationActionEventHubResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action_type: str,
                 sas_policy_name: str,
                 connection_string: Optional[str] = None,
                 event_hub_resource_id: Optional[str] = None,
                 is_trusted_service_enabled: Optional[bool] = None):
        """
        The target Event Hub to which event data will be exported. To learn more about Microsoft Defender for Cloud continuous export capabilities, visit https://aka.ms/ASCExportLearnMore
        :param str action_type: The type of the action that will be triggered by the Automation
               Expected value is 'EventHub'.
        :param str sas_policy_name: The target Event Hub SAS policy name.
        :param str connection_string: The target Event Hub connection string (it will not be included in any response).
        :param str event_hub_resource_id: The target Event Hub Azure Resource ID.
        :param bool is_trusted_service_enabled: Indicates whether the trusted service is enabled or not.
        """
        pulumi.set(__self__, "action_type", 'EventHub')
        pulumi.set(__self__, "sas_policy_name", sas_policy_name)
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if event_hub_resource_id is not None:
            pulumi.set(__self__, "event_hub_resource_id", event_hub_resource_id)
        if is_trusted_service_enabled is not None:
            pulumi.set(__self__, "is_trusted_service_enabled", is_trusted_service_enabled)

    @property
    @pulumi.getter(name="actionType")
    def action_type(self) -> str:
        """
        The type of the action that will be triggered by the Automation
        Expected value is 'EventHub'.
        """
        return pulumi.get(self, "action_type")

    @property
    @pulumi.getter(name="sasPolicyName")
    def sas_policy_name(self) -> str:
        """
        The target Event Hub SAS policy name.
        """
        return pulumi.get(self, "sas_policy_name")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[str]:
        """
        The target Event Hub connection string (it will not be included in any response).
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="eventHubResourceId")
    def event_hub_resource_id(self) -> Optional[str]:
        """
        The target Event Hub Azure Resource ID.
        """
        return pulumi.get(self, "event_hub_resource_id")

    @property
    @pulumi.getter(name="isTrustedServiceEnabled")
    def is_trusted_service_enabled(self) -> Optional[bool]:
        """
        Indicates whether the trusted service is enabled or not.
        """
        return pulumi.get(self, "is_trusted_service_enabled")


@pulumi.output_type
class AutomationActionLogicAppResponse(dict):
    """
    The logic app action that should be triggered. To learn more about Microsoft Defender for Cloud's Workflow Automation capabilities, visit https://aka.ms/ASCWorkflowAutomationLearnMore
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionType":
            suggest = "action_type"
        elif key == "logicAppResourceId":
            suggest = "logic_app_resource_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AutomationActionLogicAppResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AutomationActionLogicAppResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AutomationActionLogicAppResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action_type: str,
                 logic_app_resource_id: Optional[str] = None,
                 uri: Optional[str] = None):
        """
        The logic app action that should be triggered. To learn more about Microsoft Defender for Cloud's Workflow Automation capabilities, visit https://aka.ms/ASCWorkflowAutomationLearnMore
        :param str action_type: The type of the action that will be triggered by the Automation
               Expected value is 'LogicApp'.
        :param str logic_app_resource_id: The triggered Logic App Azure Resource ID. This can also reside on other subscriptions, given that you have permissions to trigger the Logic App
        :param str uri: The Logic App trigger URI endpoint (it will not be included in any response).
        """
        pulumi.set(__self__, "action_type", 'LogicApp')
        if logic_app_resource_id is not None:
            pulumi.set(__self__, "logic_app_resource_id", logic_app_resource_id)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="actionType")
    def action_type(self) -> str:
        """
        The type of the action that will be triggered by the Automation
        Expected value is 'LogicApp'.
        """
        return pulumi.get(self, "action_type")

    @property
    @pulumi.getter(name="logicAppResourceId")
    def logic_app_resource_id(self) -> Optional[str]:
        """
        The triggered Logic App Azure Resource ID. This can also reside on other subscriptions, given that you have permissions to trigger the Logic App
        """
        return pulumi.get(self, "logic_app_resource_id")

    @property
    @pulumi.getter
    def uri(self) -> Optional[str]:
        """
        The Logic App trigger URI endpoint (it will not be included in any response).
        """
        return pulumi.get(self, "uri")


@pulumi.output_type
class AutomationActionWorkspaceResponse(dict):
    """
    The Log Analytics Workspace to which event data will be exported. Security alerts data will reside in the 'SecurityAlert' table and the assessments data will reside in the 'SecurityRecommendation' table (under the 'Security'/'SecurityCenterFree' solutions). Note that in order to view the data in the workspace, the Security Center Log Analytics free/standard solution needs to be enabled on that workspace. To learn more about Microsoft Defender for Cloud continuous export capabilities, visit https://aka.ms/ASCExportLearnMore
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionType":
            suggest = "action_type"
        elif key == "workspaceResourceId":
            suggest = "workspace_resource_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AutomationActionWorkspaceResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AutomationActionWorkspaceResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AutomationActionWorkspaceResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action_type: str,
                 workspace_resource_id: Optional[str] = None):
        """
        The Log Analytics Workspace to which event data will be exported. Security alerts data will reside in the 'SecurityAlert' table and the assessments data will reside in the 'SecurityRecommendation' table (under the 'Security'/'SecurityCenterFree' solutions). Note that in order to view the data in the workspace, the Security Center Log Analytics free/standard solution needs to be enabled on that workspace. To learn more about Microsoft Defender for Cloud continuous export capabilities, visit https://aka.ms/ASCExportLearnMore
        :param str action_type: The type of the action that will be triggered by the Automation
               Expected value is 'Workspace'.
        :param str workspace_resource_id: The fully qualified Log Analytics Workspace Azure Resource ID.
        """
        pulumi.set(__self__, "action_type", 'Workspace')
        if workspace_resource_id is not None:
            pulumi.set(__self__, "workspace_resource_id", workspace_resource_id)

    @property
    @pulumi.getter(name="actionType")
    def action_type(self) -> str:
        """
        The type of the action that will be triggered by the Automation
        Expected value is 'Workspace'.
        """
        return pulumi.get(self, "action_type")

    @property
    @pulumi.getter(name="workspaceResourceId")
    def workspace_resource_id(self) -> Optional[str]:
        """
        The fully qualified Log Analytics Workspace Azure Resource ID.
        """
        return pulumi.get(self, "workspace_resource_id")


@pulumi.output_type
class AutomationRuleSetResponse(dict):
    """
    A rule set which evaluates all its rules upon an event interception. Only when all the included rules in the rule set will be evaluated as 'true', will the event trigger the defined actions. 
    """
    def __init__(__self__, *,
                 rules: Optional[Sequence['outputs.AutomationTriggeringRuleResponse']] = None):
        """
        A rule set which evaluates all its rules upon an event interception. Only when all the included rules in the rule set will be evaluated as 'true', will the event trigger the defined actions. 
        """
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> Optional[Sequence['outputs.AutomationTriggeringRuleResponse']]:
        return pulumi.get(self, "rules")


@pulumi.output_type
class AutomationScopeResponse(dict):
    """
    A single automation scope.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "scopePath":
            suggest = "scope_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AutomationScopeResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AutomationScopeResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AutomationScopeResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[str] = None,
                 scope_path: Optional[str] = None):
        """
        A single automation scope.
        :param str description: The resources scope description.
        :param str scope_path: The resources scope path. Can be the subscription on which the automation is defined on or a resource group under that subscription (fully qualified Azure resource IDs).
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if scope_path is not None:
            pulumi.set(__self__, "scope_path", scope_path)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The resources scope description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="scopePath")
    def scope_path(self) -> Optional[str]:
        """
        The resources scope path. Can be the subscription on which the automation is defined on or a resource group under that subscription (fully qualified Azure resource IDs).
        """
        return pulumi.get(self, "scope_path")


@pulumi.output_type
class AutomationSourceResponse(dict):
    """
    The source event types which evaluate the security automation set of rules. For example - security alerts and security assessments. To learn more about the supported security events data models schemas - please visit https://aka.ms/ASCAutomationSchemas.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "eventSource":
            suggest = "event_source"
        elif key == "ruleSets":
            suggest = "rule_sets"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AutomationSourceResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AutomationSourceResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AutomationSourceResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 event_source: Optional[str] = None,
                 rule_sets: Optional[Sequence['outputs.AutomationRuleSetResponse']] = None):
        """
        The source event types which evaluate the security automation set of rules. For example - security alerts and security assessments. To learn more about the supported security events data models schemas - please visit https://aka.ms/ASCAutomationSchemas.
        :param str event_source: A valid event source type.
        :param Sequence['AutomationRuleSetResponse'] rule_sets: A set of rules which evaluate upon event interception. A logical disjunction is applied between defined rule sets (logical 'or').
        """
        if event_source is not None:
            pulumi.set(__self__, "event_source", event_source)
        if rule_sets is not None:
            pulumi.set(__self__, "rule_sets", rule_sets)

    @property
    @pulumi.getter(name="eventSource")
    def event_source(self) -> Optional[str]:
        """
        A valid event source type.
        """
        return pulumi.get(self, "event_source")

    @property
    @pulumi.getter(name="ruleSets")
    def rule_sets(self) -> Optional[Sequence['outputs.AutomationRuleSetResponse']]:
        """
        A set of rules which evaluate upon event interception. A logical disjunction is applied between defined rule sets (logical 'or').
        """
        return pulumi.get(self, "rule_sets")


@pulumi.output_type
class AutomationTriggeringRuleResponse(dict):
    """
    A rule which is evaluated upon event interception. The rule is configured by comparing a specific value from the event model to an expected value. This comparison is done by using one of the supported operators set.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expectedValue":
            suggest = "expected_value"
        elif key == "propertyJPath":
            suggest = "property_j_path"
        elif key == "propertyType":
            suggest = "property_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AutomationTriggeringRuleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AutomationTriggeringRuleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AutomationTriggeringRuleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expected_value: Optional[str] = None,
                 operator: Optional[str] = None,
                 property_j_path: Optional[str] = None,
                 property_type: Optional[str] = None):
        """
        A rule which is evaluated upon event interception. The rule is configured by comparing a specific value from the event model to an expected value. This comparison is done by using one of the supported operators set.
        :param str expected_value: The expected value.
        :param str operator: A valid comparer operator to use. A case-insensitive comparison will be applied for String PropertyType.
        :param str property_j_path: The JPath of the entity model property that should be checked.
        :param str property_type: The data type of the compared operands (string, integer, floating point number or a boolean [true/false]]
        """
        if expected_value is not None:
            pulumi.set(__self__, "expected_value", expected_value)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if property_j_path is not None:
            pulumi.set(__self__, "property_j_path", property_j_path)
        if property_type is not None:
            pulumi.set(__self__, "property_type", property_type)

    @property
    @pulumi.getter(name="expectedValue")
    def expected_value(self) -> Optional[str]:
        """
        The expected value.
        """
        return pulumi.get(self, "expected_value")

    @property
    @pulumi.getter
    def operator(self) -> Optional[str]:
        """
        A valid comparer operator to use. A case-insensitive comparison will be applied for String PropertyType.
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter(name="propertyJPath")
    def property_j_path(self) -> Optional[str]:
        """
        The JPath of the entity model property that should be checked.
        """
        return pulumi.get(self, "property_j_path")

    @property
    @pulumi.getter(name="propertyType")
    def property_type(self) -> Optional[str]:
        """
        The data type of the compared operands (string, integer, floating point number or a boolean [true/false]]
        """
        return pulumi.get(self, "property_type")


@pulumi.output_type
class NotificationsSourceAlertResponse(dict):
    """
    Alert notification source
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceType":
            suggest = "source_type"
        elif key == "minimalSeverity":
            suggest = "minimal_severity"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotificationsSourceAlertResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotificationsSourceAlertResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotificationsSourceAlertResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 source_type: str,
                 minimal_severity: Optional[str] = None):
        """
        Alert notification source
        :param str source_type: The source type that will trigger the notification
               Expected value is 'Alert'.
        :param str minimal_severity: Defines the minimal alert severity which will be sent as email notifications
        """
        pulumi.set(__self__, "source_type", 'Alert')
        if minimal_severity is not None:
            pulumi.set(__self__, "minimal_severity", minimal_severity)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        """
        The source type that will trigger the notification
        Expected value is 'Alert'.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="minimalSeverity")
    def minimal_severity(self) -> Optional[str]:
        """
        Defines the minimal alert severity which will be sent as email notifications
        """
        return pulumi.get(self, "minimal_severity")


@pulumi.output_type
class NotificationsSourceAttackPathResponse(dict):
    """
    Attack path notification source
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceType":
            suggest = "source_type"
        elif key == "minimalRiskLevel":
            suggest = "minimal_risk_level"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotificationsSourceAttackPathResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotificationsSourceAttackPathResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotificationsSourceAttackPathResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 source_type: str,
                 minimal_risk_level: Optional[str] = None):
        """
        Attack path notification source
        :param str source_type: The source type that will trigger the notification
               Expected value is 'AttackPath'.
        :param str minimal_risk_level: Defines the minimal attach path risk level which will be sent as email notifications
        """
        pulumi.set(__self__, "source_type", 'AttackPath')
        if minimal_risk_level is not None:
            pulumi.set(__self__, "minimal_risk_level", minimal_risk_level)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        """
        The source type that will trigger the notification
        Expected value is 'AttackPath'.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="minimalRiskLevel")
    def minimal_risk_level(self) -> Optional[str]:
        """
        Defines the minimal attach path risk level which will be sent as email notifications
        """
        return pulumi.get(self, "minimal_risk_level")


@pulumi.output_type
class SecurityContactPropertiesResponseNotificationsByRole(dict):
    """
    Defines whether to send email notifications from Microsoft Defender for Cloud to persons with specific RBAC roles on the subscription.
    """
    def __init__(__self__, *,
                 roles: Optional[Sequence[str]] = None,
                 state: Optional[str] = None):
        """
        Defines whether to send email notifications from Microsoft Defender for Cloud to persons with specific RBAC roles on the subscription.
        :param Sequence[str] roles: Defines which RBAC roles will get email notifications from Microsoft Defender for Cloud. List of allowed RBAC roles: 
        :param str state: Defines whether to send email notifications from AMicrosoft Defender for Cloud to persons with specific RBAC roles on the subscription.
        """
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def roles(self) -> Optional[Sequence[str]]:
        """
        Defines which RBAC roles will get email notifications from Microsoft Defender for Cloud. List of allowed RBAC roles: 
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        Defines whether to send email notifications from AMicrosoft Defender for Cloud to persons with specific RBAC roles on the subscription.
        """
        return pulumi.get(self, "state")


