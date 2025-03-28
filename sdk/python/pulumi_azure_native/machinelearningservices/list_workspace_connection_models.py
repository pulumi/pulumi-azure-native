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
    'ListWorkspaceConnectionModelsResult',
    'AwaitableListWorkspaceConnectionModelsResult',
    'list_workspace_connection_models',
    'list_workspace_connection_models_output',
]

@pulumi.output_type
class ListWorkspaceConnectionModelsResult:
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> Optional[str]:
        """
        The link to the next page constructed using the continuationToken.  If null, there are no additional pages.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Optional[Sequence['outputs.AccountModelResponse']]:
        """
        List of models.
        """
        return pulumi.get(self, "value")


class AwaitableListWorkspaceConnectionModelsResult(ListWorkspaceConnectionModelsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListWorkspaceConnectionModelsResult(
            next_link=self.next_link,
            value=self.value)


def list_workspace_connection_models(resource_group_name: Optional[str] = None,
                                     workspace_name: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListWorkspaceConnectionModelsResult:
    """
    List available models from all connections.

    Uses Azure REST API version 2024-04-01-preview.


    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str workspace_name: Azure Machine Learning Workspace Name
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:machinelearningservices:listWorkspaceConnectionModels', __args__, opts=opts, typ=ListWorkspaceConnectionModelsResult).value

    return AwaitableListWorkspaceConnectionModelsResult(
        next_link=pulumi.get(__ret__, 'next_link'),
        value=pulumi.get(__ret__, 'value'))
def list_workspace_connection_models_output(resource_group_name: Optional[pulumi.Input[str]] = None,
                                            workspace_name: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListWorkspaceConnectionModelsResult]:
    """
    List available models from all connections.

    Uses Azure REST API version 2024-04-01-preview.


    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str workspace_name: Azure Machine Learning Workspace Name
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:machinelearningservices:listWorkspaceConnectionModels', __args__, opts=opts, typ=ListWorkspaceConnectionModelsResult)
    return __ret__.apply(lambda __response__: ListWorkspaceConnectionModelsResult(
        next_link=pulumi.get(__response__, 'next_link'),
        value=pulumi.get(__response__, 'value')))
