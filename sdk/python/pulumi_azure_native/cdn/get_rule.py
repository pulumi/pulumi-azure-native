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
    'GetRuleResult',
    'AwaitableGetRuleResult',
    'get_rule',
    'get_rule_output',
]

@pulumi.output_type
class GetRuleResult:
    """
    Friendly Rules name mapping to the any Rules or secret related information.
    """
    def __init__(__self__, actions=None, azure_api_version=None, conditions=None, deployment_status=None, id=None, match_processing_behavior=None, name=None, order=None, provisioning_state=None, rule_set_name=None, system_data=None, type=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if conditions and not isinstance(conditions, list):
            raise TypeError("Expected argument 'conditions' to be a list")
        pulumi.set(__self__, "conditions", conditions)
        if deployment_status and not isinstance(deployment_status, str):
            raise TypeError("Expected argument 'deployment_status' to be a str")
        pulumi.set(__self__, "deployment_status", deployment_status)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if match_processing_behavior and not isinstance(match_processing_behavior, str):
            raise TypeError("Expected argument 'match_processing_behavior' to be a str")
        pulumi.set(__self__, "match_processing_behavior", match_processing_behavior)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if order and not isinstance(order, int):
            raise TypeError("Expected argument 'order' to be a int")
        pulumi.set(__self__, "order", order)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if rule_set_name and not isinstance(rule_set_name, str):
            raise TypeError("Expected argument 'rule_set_name' to be a str")
        pulumi.set(__self__, "rule_set_name", rule_set_name)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def actions(self) -> Sequence[Any]:
        """
        A list of actions that are executed when all the conditions of a rule are satisfied.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def conditions(self) -> Optional[Sequence[Any]]:
        """
        A list of conditions that must be matched for the actions to be executed
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="deploymentStatus")
    def deployment_status(self) -> builtins.str:
        return pulumi.get(self, "deployment_status")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="matchProcessingBehavior")
    def match_processing_behavior(self) -> Optional[builtins.str]:
        """
        If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
        """
        return pulumi.get(self, "match_processing_behavior")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> builtins.int:
        """
        The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning status
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="ruleSetName")
    def rule_set_name(self) -> builtins.str:
        """
        The name of the rule set containing the rule.
        """
        return pulumi.get(self, "rule_set_name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Read only system data
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetRuleResult(GetRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRuleResult(
            actions=self.actions,
            azure_api_version=self.azure_api_version,
            conditions=self.conditions,
            deployment_status=self.deployment_status,
            id=self.id,
            match_processing_behavior=self.match_processing_behavior,
            name=self.name,
            order=self.order,
            provisioning_state=self.provisioning_state,
            rule_set_name=self.rule_set_name,
            system_data=self.system_data,
            type=self.type)


def get_rule(profile_name: Optional[builtins.str] = None,
             resource_group_name: Optional[builtins.str] = None,
             rule_name: Optional[builtins.str] = None,
             rule_set_name: Optional[builtins.str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRuleResult:
    """
    Gets an existing delivery rule within a rule set.

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str profile_name: Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
    :param builtins.str resource_group_name: Name of the Resource group within the Azure subscription.
    :param builtins.str rule_name: Name of the delivery rule which is unique within the endpoint.
    :param builtins.str rule_set_name: Name of the rule set under the profile.
    """
    __args__ = dict()
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleName'] = rule_name
    __args__['ruleSetName'] = rule_set_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:cdn:getRule', __args__, opts=opts, typ=GetRuleResult).value

    return AwaitableGetRuleResult(
        actions=pulumi.get(__ret__, 'actions'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        conditions=pulumi.get(__ret__, 'conditions'),
        deployment_status=pulumi.get(__ret__, 'deployment_status'),
        id=pulumi.get(__ret__, 'id'),
        match_processing_behavior=pulumi.get(__ret__, 'match_processing_behavior'),
        name=pulumi.get(__ret__, 'name'),
        order=pulumi.get(__ret__, 'order'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        rule_set_name=pulumi.get(__ret__, 'rule_set_name'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_rule_output(profile_name: Optional[pulumi.Input[builtins.str]] = None,
                    resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                    rule_name: Optional[pulumi.Input[builtins.str]] = None,
                    rule_set_name: Optional[pulumi.Input[builtins.str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRuleResult]:
    """
    Gets an existing delivery rule within a rule set.

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str profile_name: Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
    :param builtins.str resource_group_name: Name of the Resource group within the Azure subscription.
    :param builtins.str rule_name: Name of the delivery rule which is unique within the endpoint.
    :param builtins.str rule_set_name: Name of the rule set under the profile.
    """
    __args__ = dict()
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleName'] = rule_name
    __args__['ruleSetName'] = rule_set_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:cdn:getRule', __args__, opts=opts, typ=GetRuleResult)
    return __ret__.apply(lambda __response__: GetRuleResult(
        actions=pulumi.get(__response__, 'actions'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        conditions=pulumi.get(__response__, 'conditions'),
        deployment_status=pulumi.get(__response__, 'deployment_status'),
        id=pulumi.get(__response__, 'id'),
        match_processing_behavior=pulumi.get(__response__, 'match_processing_behavior'),
        name=pulumi.get(__response__, 'name'),
        order=pulumi.get(__response__, 'order'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        rule_set_name=pulumi.get(__response__, 'rule_set_name'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
