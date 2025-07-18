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
    'ListSkusByProjectResult',
    'AwaitableListSkusByProjectResult',
    'list_skus_by_project',
    'list_skus_by_project_output',
]

@pulumi.output_type
class ListSkusByProjectResult:
    """
    Results of the Microsoft.DevCenter SKU list operation.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> builtins.str:
        """
        URL to get the next set of results if there are any.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Sequence['outputs.DevCenterSkuResponse']:
        """
        Current page of results.
        """
        return pulumi.get(self, "value")


class AwaitableListSkusByProjectResult(ListSkusByProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListSkusByProjectResult(
            next_link=self.next_link,
            value=self.value)


def list_skus_by_project(project_name: Optional[builtins.str] = None,
                         resource_group_name: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListSkusByProjectResult:
    """
    Lists SKUs available to the project

    Uses Azure REST API version 2024-10-01-preview.

    Other available API versions: 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str project_name: The name of the project.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['projectName'] = project_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:devcenter:listSkusByProject', __args__, opts=opts, typ=ListSkusByProjectResult).value

    return AwaitableListSkusByProjectResult(
        next_link=pulumi.get(__ret__, 'next_link'),
        value=pulumi.get(__ret__, 'value'))
def list_skus_by_project_output(project_name: Optional[pulumi.Input[builtins.str]] = None,
                                resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListSkusByProjectResult]:
    """
    Lists SKUs available to the project

    Uses Azure REST API version 2024-10-01-preview.

    Other available API versions: 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str project_name: The name of the project.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['projectName'] = project_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:devcenter:listSkusByProject', __args__, opts=opts, typ=ListSkusByProjectResult)
    return __ret__.apply(lambda __response__: ListSkusByProjectResult(
        next_link=pulumi.get(__response__, 'next_link'),
        value=pulumi.get(__response__, 'value')))
