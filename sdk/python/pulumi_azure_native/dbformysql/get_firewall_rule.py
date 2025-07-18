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
    'GetFirewallRuleResult',
    'AwaitableGetFirewallRuleResult',
    'get_firewall_rule',
    'get_firewall_rule_output',
]

@pulumi.output_type
class GetFirewallRuleResult:
    """
    Represents a server firewall rule.
    """
    def __init__(__self__, azure_api_version=None, end_ip_address=None, id=None, name=None, start_ip_address=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if end_ip_address and not isinstance(end_ip_address, str):
            raise TypeError("Expected argument 'end_ip_address' to be a str")
        pulumi.set(__self__, "end_ip_address", end_ip_address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start_ip_address and not isinstance(start_ip_address, str):
            raise TypeError("Expected argument 'start_ip_address' to be a str")
        pulumi.set(__self__, "start_ip_address", start_ip_address)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
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
    @pulumi.getter(name="endIpAddress")
    def end_ip_address(self) -> builtins.str:
        """
        The end IP address of the server firewall rule. Must be IPv4 format.
        """
        return pulumi.get(self, "end_ip_address")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startIpAddress")
    def start_ip_address(self) -> builtins.str:
        """
        The start IP address of the server firewall rule. Must be IPv4 format.
        """
        return pulumi.get(self, "start_ip_address")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetFirewallRuleResult(GetFirewallRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallRuleResult(
            azure_api_version=self.azure_api_version,
            end_ip_address=self.end_ip_address,
            id=self.id,
            name=self.name,
            start_ip_address=self.start_ip_address,
            system_data=self.system_data,
            type=self.type)


def get_firewall_rule(firewall_rule_name: Optional[builtins.str] = None,
                      resource_group_name: Optional[builtins.str] = None,
                      server_name: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallRuleResult:
    """
    Gets information about a server firewall rule.

    Uses Azure REST API version 2023-12-30.

    Other available API versions: 2022-01-01, 2023-06-01-preview, 2023-06-30, 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbformysql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str firewall_rule_name: The name of the server firewall rule.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['firewallRuleName'] = firewall_rule_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:dbformysql:getFirewallRule', __args__, opts=opts, typ=GetFirewallRuleResult).value

    return AwaitableGetFirewallRuleResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        end_ip_address=pulumi.get(__ret__, 'end_ip_address'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        start_ip_address=pulumi.get(__ret__, 'start_ip_address'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_firewall_rule_output(firewall_rule_name: Optional[pulumi.Input[builtins.str]] = None,
                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                             server_name: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFirewallRuleResult]:
    """
    Gets information about a server firewall rule.

    Uses Azure REST API version 2023-12-30.

    Other available API versions: 2022-01-01, 2023-06-01-preview, 2023-06-30, 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbformysql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str firewall_rule_name: The name of the server firewall rule.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['firewallRuleName'] = firewall_rule_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:dbformysql:getFirewallRule', __args__, opts=opts, typ=GetFirewallRuleResult)
    return __ret__.apply(lambda __response__: GetFirewallRuleResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        end_ip_address=pulumi.get(__response__, 'end_ip_address'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        start_ip_address=pulumi.get(__response__, 'start_ip_address'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
