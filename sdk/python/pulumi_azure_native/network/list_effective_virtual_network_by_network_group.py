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
    'ListEffectiveVirtualNetworkByNetworkGroupResult',
    'AwaitableListEffectiveVirtualNetworkByNetworkGroupResult',
    'list_effective_virtual_network_by_network_group',
    'list_effective_virtual_network_by_network_group_output',
]

@pulumi.output_type
class ListEffectiveVirtualNetworkByNetworkGroupResult:
    """
    Result of the request to list Effective Virtual Network. It contains a list of groups and a URL link to get the next set of results.
    """
    def __init__(__self__, skip_token=None, value=None):
        if skip_token and not isinstance(skip_token, str):
            raise TypeError("Expected argument 'skip_token' to be a str")
        pulumi.set(__self__, "skip_token", skip_token)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="skipToken")
    def skip_token(self) -> Optional[builtins.str]:
        """
        When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
        """
        return pulumi.get(self, "skip_token")

    @property
    @pulumi.getter
    def value(self) -> Optional[Sequence['outputs.EffectiveVirtualNetworkResponse']]:
        """
        Gets a page of EffectiveVirtualNetwork
        """
        return pulumi.get(self, "value")


class AwaitableListEffectiveVirtualNetworkByNetworkGroupResult(ListEffectiveVirtualNetworkByNetworkGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListEffectiveVirtualNetworkByNetworkGroupResult(
            skip_token=self.skip_token,
            value=self.value)


def list_effective_virtual_network_by_network_group(network_group_name: Optional[builtins.str] = None,
                                                    network_manager_name: Optional[builtins.str] = None,
                                                    resource_group_name: Optional[builtins.str] = None,
                                                    skip_token: Optional[builtins.str] = None,
                                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListEffectiveVirtualNetworkByNetworkGroupResult:
    """
    Lists all effective virtual networks by specified network group.

    Uses Azure REST API version 2021-02-01-preview.


    :param builtins.str network_group_name: The name of the network group to get.
    :param builtins.str network_manager_name: The name of the network manager.
    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str skip_token: When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
    """
    __args__ = dict()
    __args__['networkGroupName'] = network_group_name
    __args__['networkManagerName'] = network_manager_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['skipToken'] = skip_token
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:network:listEffectiveVirtualNetworkByNetworkGroup', __args__, opts=opts, typ=ListEffectiveVirtualNetworkByNetworkGroupResult).value

    return AwaitableListEffectiveVirtualNetworkByNetworkGroupResult(
        skip_token=pulumi.get(__ret__, 'skip_token'),
        value=pulumi.get(__ret__, 'value'))
def list_effective_virtual_network_by_network_group_output(network_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                           network_manager_name: Optional[pulumi.Input[builtins.str]] = None,
                                                           resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                           skip_token: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListEffectiveVirtualNetworkByNetworkGroupResult]:
    """
    Lists all effective virtual networks by specified network group.

    Uses Azure REST API version 2021-02-01-preview.


    :param builtins.str network_group_name: The name of the network group to get.
    :param builtins.str network_manager_name: The name of the network manager.
    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str skip_token: When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
    """
    __args__ = dict()
    __args__['networkGroupName'] = network_group_name
    __args__['networkManagerName'] = network_manager_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['skipToken'] = skip_token
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:network:listEffectiveVirtualNetworkByNetworkGroup', __args__, opts=opts, typ=ListEffectiveVirtualNetworkByNetworkGroupResult)
    return __ret__.apply(lambda __response__: ListEffectiveVirtualNetworkByNetworkGroupResult(
        skip_token=pulumi.get(__response__, 'skip_token'),
        value=pulumi.get(__response__, 'value')))
