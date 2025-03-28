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
from .. import _utilities
from . import outputs

__all__ = [
    'GetDevOpsPolicyAssignmentResult',
    'AwaitableGetDevOpsPolicyAssignmentResult',
    'get_dev_ops_policy_assignment',
    'get_dev_ops_policy_assignment_output',
]

@pulumi.output_type
class GetDevOpsPolicyAssignmentResult:
    """
    DevOps Policy assignment resource.
    """
    def __init__(__self__, id=None, name=None, properties=None, system_data=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.DevOpsPolicyAssignmentPropertiesResponse':
        """
        Properties of the DevOps policy assignment resource.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetDevOpsPolicyAssignmentResult(GetDevOpsPolicyAssignmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDevOpsPolicyAssignmentResult(
            id=self.id,
            name=self.name,
            properties=self.properties,
            system_data=self.system_data,
            type=self.type)


def get_dev_ops_policy_assignment(policy_assignment_id: Optional[str] = None,
                                  resource_group_name: Optional[str] = None,
                                  security_connector_name: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDevOpsPolicyAssignmentResult:
    """
    DevOps Policy assignment resource.

    Uses Azure REST API version 2024-05-15-preview.


    :param str policy_assignment_id: The policy assignment Id.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str security_connector_name: The security connector name.
    """
    __args__ = dict()
    __args__['policyAssignmentId'] = policy_assignment_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['securityConnectorName'] = security_connector_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:security:getDevOpsPolicyAssignment', __args__, opts=opts, typ=GetDevOpsPolicyAssignmentResult).value

    return AwaitableGetDevOpsPolicyAssignmentResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_dev_ops_policy_assignment_output(policy_assignment_id: Optional[pulumi.Input[str]] = None,
                                         resource_group_name: Optional[pulumi.Input[str]] = None,
                                         security_connector_name: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDevOpsPolicyAssignmentResult]:
    """
    DevOps Policy assignment resource.

    Uses Azure REST API version 2024-05-15-preview.


    :param str policy_assignment_id: The policy assignment Id.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str security_connector_name: The security connector name.
    """
    __args__ = dict()
    __args__['policyAssignmentId'] = policy_assignment_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['securityConnectorName'] = security_connector_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:security:getDevOpsPolicyAssignment', __args__, opts=opts, typ=GetDevOpsPolicyAssignmentResult)
    return __ret__.apply(lambda __response__: GetDevOpsPolicyAssignmentResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
