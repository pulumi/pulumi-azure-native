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
    'GetSolutionResult',
    'AwaitableGetSolutionResult',
    'get_solution',
    'get_solution_output',
]

@pulumi.output_type
class GetSolutionResult:
    """
    Solution REST Resource.
    """
    def __init__(__self__, azure_api_version=None, etag=None, id=None, name=None, properties=None, type=None):
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
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
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
    def etag(self) -> Optional[builtins.str]:
        """
        Gets or sets the ETAG for optimistic concurrency control.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Gets the relative URL to get to this REST resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Gets the name of this REST resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.SolutionPropertiesResponse':
        """
        Gets or sets the properties of the solution.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Gets the type of this REST resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetSolutionResult(GetSolutionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSolutionResult(
            azure_api_version=self.azure_api_version,
            etag=self.etag,
            id=self.id,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_solution(migrate_project_name: Optional[builtins.str] = None,
                 resource_group_name: Optional[builtins.str] = None,
                 solution_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSolutionResult:
    """
    Solution REST Resource.

    Uses Azure REST API version 2018-09-01-preview.


    :param builtins.str migrate_project_name: Name of the Azure Migrate project.
    :param builtins.str resource_group_name: Name of the Azure Resource Group that migrate project is part of.
    :param builtins.str solution_name: Unique name of a migration solution within a migrate project.
    """
    __args__ = dict()
    __args__['migrateProjectName'] = migrate_project_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['solutionName'] = solution_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:migrate:getSolution', __args__, opts=opts, typ=GetSolutionResult).value

    return AwaitableGetSolutionResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        type=pulumi.get(__ret__, 'type'))
def get_solution_output(migrate_project_name: Optional[pulumi.Input[builtins.str]] = None,
                        resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                        solution_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSolutionResult]:
    """
    Solution REST Resource.

    Uses Azure REST API version 2018-09-01-preview.


    :param builtins.str migrate_project_name: Name of the Azure Migrate project.
    :param builtins.str resource_group_name: Name of the Azure Resource Group that migrate project is part of.
    :param builtins.str solution_name: Unique name of a migration solution within a migrate project.
    """
    __args__ = dict()
    __args__['migrateProjectName'] = migrate_project_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['solutionName'] = solution_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:migrate:getSolution', __args__, opts=opts, typ=GetSolutionResult)
    return __ret__.apply(lambda __response__: GetSolutionResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        type=pulumi.get(__response__, 'type')))
