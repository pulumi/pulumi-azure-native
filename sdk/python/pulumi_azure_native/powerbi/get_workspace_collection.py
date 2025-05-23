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
    'GetWorkspaceCollectionResult',
    'AwaitableGetWorkspaceCollectionResult',
    'get_workspace_collection',
    'get_workspace_collection_output',
]

@pulumi.output_type
class GetWorkspaceCollectionResult:
    def __init__(__self__, azure_api_version=None, id=None, location=None, name=None, properties=None, sku=None, tags=None, type=None):
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
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
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
        Resource id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        Azure location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        Workspace collection name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> Any:
        """
        Properties
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.AzureSkuResponse']:
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetWorkspaceCollectionResult(GetWorkspaceCollectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkspaceCollectionResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            location=self.location,
            name=self.name,
            properties=self.properties,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_workspace_collection(resource_group_name: Optional[builtins.str] = None,
                             workspace_collection_name: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkspaceCollectionResult:
    """
    Retrieves an existing Power BI Workspace Collection.

    Uses Azure REST API version 2016-01-29.


    :param builtins.str resource_group_name: Azure resource group
    :param builtins.str workspace_collection_name: Power BI Embedded Workspace Collection name
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceCollectionName'] = workspace_collection_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:powerbi:getWorkspaceCollection', __args__, opts=opts, typ=GetWorkspaceCollectionResult).value

    return AwaitableGetWorkspaceCollectionResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        sku=pulumi.get(__ret__, 'sku'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_workspace_collection_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                    workspace_collection_name: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWorkspaceCollectionResult]:
    """
    Retrieves an existing Power BI Workspace Collection.

    Uses Azure REST API version 2016-01-29.


    :param builtins.str resource_group_name: Azure resource group
    :param builtins.str workspace_collection_name: Power BI Embedded Workspace Collection name
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceCollectionName'] = workspace_collection_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:powerbi:getWorkspaceCollection', __args__, opts=opts, typ=GetWorkspaceCollectionResult)
    return __ret__.apply(lambda __response__: GetWorkspaceCollectionResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        sku=pulumi.get(__response__, 'sku'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
