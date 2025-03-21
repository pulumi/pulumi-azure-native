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

__all__ = [
    'GetLinkedServiceResult',
    'AwaitableGetLinkedServiceResult',
    'get_linked_service',
    'get_linked_service_output',
]

@pulumi.output_type
class GetLinkedServiceResult:
    """
    The top level Linked service resource container.
    """
    def __init__(__self__, id=None, name=None, resource_id=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The resource id of the resource that will be linked to the workspace.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetLinkedServiceResult(GetLinkedServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLinkedServiceResult(
            id=self.id,
            name=self.name,
            resource_id=self.resource_id,
            type=self.type)


def get_linked_service(linked_service_name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       workspace_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLinkedServiceResult:
    """
    Gets a linked service instance.


    :param str linked_service_name: Name of the linked service.
    :param str resource_group_name: The name of the resource group to get. The name is case insensitive.
    :param str workspace_name: Name of the Log Analytics Workspace that contains the linkedServices resource
    """
    __args__ = dict()
    __args__['linkedServiceName'] = linked_service_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:operationalinsights/v20151101preview:getLinkedService', __args__, opts=opts, typ=GetLinkedServiceResult).value

    return AwaitableGetLinkedServiceResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        resource_id=pulumi.get(__ret__, 'resource_id'),
        type=pulumi.get(__ret__, 'type'))
def get_linked_service_output(linked_service_name: Optional[pulumi.Input[str]] = None,
                              resource_group_name: Optional[pulumi.Input[str]] = None,
                              workspace_name: Optional[pulumi.Input[str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLinkedServiceResult]:
    """
    Gets a linked service instance.


    :param str linked_service_name: Name of the linked service.
    :param str resource_group_name: The name of the resource group to get. The name is case insensitive.
    :param str workspace_name: Name of the Log Analytics Workspace that contains the linkedServices resource
    """
    __args__ = dict()
    __args__['linkedServiceName'] = linked_service_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:operationalinsights/v20151101preview:getLinkedService', __args__, opts=opts, typ=GetLinkedServiceResult)
    return __ret__.apply(lambda __response__: GetLinkedServiceResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        resource_id=pulumi.get(__response__, 'resource_id'),
        type=pulumi.get(__response__, 'type')))
