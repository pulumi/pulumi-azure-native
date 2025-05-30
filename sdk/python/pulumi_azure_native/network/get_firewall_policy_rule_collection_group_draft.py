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
    'GetFirewallPolicyRuleCollectionGroupDraftResult',
    'AwaitableGetFirewallPolicyRuleCollectionGroupDraftResult',
    'get_firewall_policy_rule_collection_group_draft',
    'get_firewall_policy_rule_collection_group_draft_output',
]

@pulumi.output_type
class GetFirewallPolicyRuleCollectionGroupDraftResult:
    """
    Rule Collection Group resource.
    """
    def __init__(__self__, azure_api_version=None, id=None, name=None, priority=None, rule_collections=None, size=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if rule_collections and not isinstance(rule_collections, list):
            raise TypeError("Expected argument 'rule_collections' to be a list")
        pulumi.set(__self__, "rule_collections", rule_collections)
        if size and not isinstance(size, str):
            raise TypeError("Expected argument 'size' to be a str")
        pulumi.set(__self__, "size", size)
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
        Priority of the Firewall Policy Rule Collection Group resource.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="ruleCollections")
    def rule_collections(self) -> Optional[Sequence[Any]]:
        """
        Group of Firewall Policy rule collections.
        """
        return pulumi.get(self, "rule_collections")

    @property
    @pulumi.getter
    def size(self) -> builtins.str:
        """
        A read-only string that represents the size of the FirewallPolicyRuleCollectionGroupProperties in MB. (ex 1.2MB)
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Rule Group type.
        """
        return pulumi.get(self, "type")


class AwaitableGetFirewallPolicyRuleCollectionGroupDraftResult(GetFirewallPolicyRuleCollectionGroupDraftResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallPolicyRuleCollectionGroupDraftResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            name=self.name,
            priority=self.priority,
            rule_collections=self.rule_collections,
            size=self.size,
            type=self.type)


def get_firewall_policy_rule_collection_group_draft(firewall_policy_name: Optional[builtins.str] = None,
                                                    resource_group_name: Optional[builtins.str] = None,
                                                    rule_collection_group_name: Optional[builtins.str] = None,
                                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallPolicyRuleCollectionGroupDraftResult:
    """
    Get Rule Collection Group Draft.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str firewall_policy_name: The name of the Firewall Policy.
    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str rule_collection_group_name: The name of the FirewallPolicyRuleCollectionGroup.
    """
    __args__ = dict()
    __args__['firewallPolicyName'] = firewall_policy_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleCollectionGroupName'] = rule_collection_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:network:getFirewallPolicyRuleCollectionGroupDraft', __args__, opts=opts, typ=GetFirewallPolicyRuleCollectionGroupDraftResult).value

    return AwaitableGetFirewallPolicyRuleCollectionGroupDraftResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        priority=pulumi.get(__ret__, 'priority'),
        rule_collections=pulumi.get(__ret__, 'rule_collections'),
        size=pulumi.get(__ret__, 'size'),
        type=pulumi.get(__ret__, 'type'))
def get_firewall_policy_rule_collection_group_draft_output(firewall_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                                                           resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                           rule_collection_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFirewallPolicyRuleCollectionGroupDraftResult]:
    """
    Get Rule Collection Group Draft.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str firewall_policy_name: The name of the Firewall Policy.
    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str rule_collection_group_name: The name of the FirewallPolicyRuleCollectionGroup.
    """
    __args__ = dict()
    __args__['firewallPolicyName'] = firewall_policy_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleCollectionGroupName'] = rule_collection_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:network:getFirewallPolicyRuleCollectionGroupDraft', __args__, opts=opts, typ=GetFirewallPolicyRuleCollectionGroupDraftResult)
    return __ret__.apply(lambda __response__: GetFirewallPolicyRuleCollectionGroupDraftResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        priority=pulumi.get(__response__, 'priority'),
        rule_collections=pulumi.get(__response__, 'rule_collections'),
        size=pulumi.get(__response__, 'size'),
        type=pulumi.get(__response__, 'type')))
