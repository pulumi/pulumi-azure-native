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
    'ListLocalRulestackCountriesResult',
    'AwaitableListLocalRulestackCountriesResult',
    'list_local_rulestack_countries',
    'list_local_rulestack_countries_output',
]

@pulumi.output_type
class ListLocalRulestackCountriesResult:
    """
    Countries Response Object
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
    def next_link(self) -> Optional[builtins.str]:
        """
        next link
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Sequence['outputs.CountryResponse']:
        """
        List of countries
        """
        return pulumi.get(self, "value")


class AwaitableListLocalRulestackCountriesResult(ListLocalRulestackCountriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListLocalRulestackCountriesResult(
            next_link=self.next_link,
            value=self.value)


def list_local_rulestack_countries(local_rulestack_name: Optional[builtins.str] = None,
                                   resource_group_name: Optional[builtins.str] = None,
                                   skip: Optional[builtins.str] = None,
                                   top: Optional[builtins.int] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListLocalRulestackCountriesResult:
    """
    List of countries for Rulestack

    Uses Azure REST API version 2025-02-06-preview.

    Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str local_rulestack_name: LocalRulestack resource name
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['localRulestackName'] = local_rulestack_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['skip'] = skip
    __args__['top'] = top
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:cloudngfw:listLocalRulestackCountries', __args__, opts=opts, typ=ListLocalRulestackCountriesResult).value

    return AwaitableListLocalRulestackCountriesResult(
        next_link=pulumi.get(__ret__, 'next_link'),
        value=pulumi.get(__ret__, 'value'))
def list_local_rulestack_countries_output(local_rulestack_name: Optional[pulumi.Input[builtins.str]] = None,
                                          resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                          skip: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          top: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListLocalRulestackCountriesResult]:
    """
    List of countries for Rulestack

    Uses Azure REST API version 2025-02-06-preview.

    Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str local_rulestack_name: LocalRulestack resource name
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['localRulestackName'] = local_rulestack_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['skip'] = skip
    __args__['top'] = top
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:cloudngfw:listLocalRulestackCountries', __args__, opts=opts, typ=ListLocalRulestackCountriesResult)
    return __ret__.apply(lambda __response__: ListLocalRulestackCountriesResult(
        next_link=pulumi.get(__response__, 'next_link'),
        value=pulumi.get(__response__, 'value')))
