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
    'GetFirewallPolicyRuleGroupResult',
    'AwaitableGetFirewallPolicyRuleGroupResult',
    'get_firewall_policy_rule_group',
    'get_firewall_policy_rule_group_output',
]

@pulumi.output_type
class GetFirewallPolicyRuleGroupResult:
    """
    Rule Group resource.
    """
    def __init__(__self__, azure_api_version=None, etag=None, id=None, name=None, priority=None, provisioning_state=None, rules=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
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
    def etag(self) -> builtins.str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> Optional[builtins.int]:
        """
        Priority of the Firewall Policy Rule Group resource.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the firewall policy rule group resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def rules(self) -> Optional[Sequence[Any]]:
        """
        Group of Firewall Policy rules.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Rule Group type.
        """
        return pulumi.get(self, "type")


class AwaitableGetFirewallPolicyRuleGroupResult(GetFirewallPolicyRuleGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallPolicyRuleGroupResult(
            azure_api_version=self.azure_api_version,
            etag=self.etag,
            id=self.id,
            name=self.name,
            priority=self.priority,
            provisioning_state=self.provisioning_state,
            rules=self.rules,
            type=self.type)


def get_firewall_policy_rule_group(firewall_policy_name: Optional[builtins.str] = None,
                                   resource_group_name: Optional[builtins.str] = None,
                                   rule_group_name: Optional[builtins.str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallPolicyRuleGroupResult:
    """
    Gets the specified FirewallPolicyRuleGroup.

    Uses Azure REST API version 2020-04-01.

    Other available API versions: 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str firewall_policy_name: The name of the Firewall Policy.
    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str rule_group_name: The name of the FirewallPolicyRuleGroup.
    """
    __args__ = dict()
    __args__['firewallPolicyName'] = firewall_policy_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleGroupName'] = rule_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:network:getFirewallPolicyRuleGroup', __args__, opts=opts, typ=GetFirewallPolicyRuleGroupResult).value

    return AwaitableGetFirewallPolicyRuleGroupResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        priority=pulumi.get(__ret__, 'priority'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        rules=pulumi.get(__ret__, 'rules'),
        type=pulumi.get(__ret__, 'type'))
def get_firewall_policy_rule_group_output(firewall_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                                          resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                          rule_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFirewallPolicyRuleGroupResult]:
    """
    Gets the specified FirewallPolicyRuleGroup.

    Uses Azure REST API version 2020-04-01.

    Other available API versions: 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str firewall_policy_name: The name of the Firewall Policy.
    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str rule_group_name: The name of the FirewallPolicyRuleGroup.
    """
    __args__ = dict()
    __args__['firewallPolicyName'] = firewall_policy_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleGroupName'] = rule_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:network:getFirewallPolicyRuleGroup', __args__, opts=opts, typ=GetFirewallPolicyRuleGroupResult)
    return __ret__.apply(lambda __response__: GetFirewallPolicyRuleGroupResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        priority=pulumi.get(__response__, 'priority'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        rules=pulumi.get(__response__, 'rules'),
        type=pulumi.get(__response__, 'type')))
