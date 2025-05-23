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
    'GetResourceSyncRuleResult',
    'AwaitableGetResourceSyncRuleResult',
    'get_resource_sync_rule',
    'get_resource_sync_rule_output',
]

@pulumi.output_type
class GetResourceSyncRuleResult:
    """
    Resource Sync Rules definition.
    """
    def __init__(__self__, azure_api_version=None, id=None, location=None, name=None, priority=None, provisioning_state=None, selector=None, system_data=None, tags=None, target_resource_group=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if selector and not isinstance(selector, dict):
            raise TypeError("Expected argument 'selector' to be a dict")
        pulumi.set(__self__, "selector", selector)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target_resource_group and not isinstance(target_resource_group, str):
            raise TypeError("Expected argument 'target_resource_group' to be a str")
        pulumi.set(__self__, "target_resource_group", target_resource_group)
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
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> Optional[builtins.int]:
        """
        Priority represents a priority of the Resource Sync Rule
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning State for the Resource Sync Rule.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def selector(self) -> Optional['outputs.ResourceSyncRulePropertiesResponseSelector']:
        """
        A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
        """
        return pulumi.get(self, "selector")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Metadata pertaining to creation and last modification of the resource
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetResourceGroup")
    def target_resource_group(self) -> Optional[builtins.str]:
        """
        For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule is one of the matching rules with highest priority, then the unmapped custom resource will be projected to the target resource group associated with this resource sync rule. The user creating this resource sync rule should have write permissions on the target resource group and this write permission will be validated when creating the resource sync rule.
        """
        return pulumi.get(self, "target_resource_group")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetResourceSyncRuleResult(GetResourceSyncRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceSyncRuleResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            location=self.location,
            name=self.name,
            priority=self.priority,
            provisioning_state=self.provisioning_state,
            selector=self.selector,
            system_data=self.system_data,
            tags=self.tags,
            target_resource_group=self.target_resource_group,
            type=self.type)


def get_resource_sync_rule(child_resource_name: Optional[builtins.str] = None,
                           resource_group_name: Optional[builtins.str] = None,
                           resource_name: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceSyncRuleResult:
    """
    Gets the details of the resourceSyncRule with a specified resource group, subscription id Custom Location resource name and Resource Sync Rule name.

    Uses Azure REST API version 2021-08-31-preview.


    :param builtins.str child_resource_name: Resource Sync Rule name.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str resource_name: Custom Locations name.
    """
    __args__ = dict()
    __args__['childResourceName'] = child_resource_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:extendedlocation:getResourceSyncRule', __args__, opts=opts, typ=GetResourceSyncRuleResult).value

    return AwaitableGetResourceSyncRuleResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        priority=pulumi.get(__ret__, 'priority'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        selector=pulumi.get(__ret__, 'selector'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        target_resource_group=pulumi.get(__ret__, 'target_resource_group'),
        type=pulumi.get(__ret__, 'type'))
def get_resource_sync_rule_output(child_resource_name: Optional[pulumi.Input[builtins.str]] = None,
                                  resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                  resource_name: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetResourceSyncRuleResult]:
    """
    Gets the details of the resourceSyncRule with a specified resource group, subscription id Custom Location resource name and Resource Sync Rule name.

    Uses Azure REST API version 2021-08-31-preview.


    :param builtins.str child_resource_name: Resource Sync Rule name.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str resource_name: Custom Locations name.
    """
    __args__ = dict()
    __args__['childResourceName'] = child_resource_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:extendedlocation:getResourceSyncRule', __args__, opts=opts, typ=GetResourceSyncRuleResult)
    return __ret__.apply(lambda __response__: GetResourceSyncRuleResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        priority=pulumi.get(__response__, 'priority'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        selector=pulumi.get(__response__, 'selector'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        target_resource_group=pulumi.get(__response__, 'target_resource_group'),
        type=pulumi.get(__response__, 'type')))
