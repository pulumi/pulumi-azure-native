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
    'GetConsoleWithLocationResult',
    'AwaitableGetConsoleWithLocationResult',
    'get_console_with_location',
    'get_console_with_location_output',
]

@pulumi.output_type
class GetConsoleWithLocationResult:
    """
    Cloud shell console
    """
    def __init__(__self__, azure_api_version=None, properties=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.ConsolePropertiesResponse':
        """
        Cloud shell console properties.
        """
        return pulumi.get(self, "properties")


class AwaitableGetConsoleWithLocationResult(GetConsoleWithLocationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConsoleWithLocationResult(
            azure_api_version=self.azure_api_version,
            properties=self.properties)


def get_console_with_location(console_name: Optional[builtins.str] = None,
                              location: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConsoleWithLocationResult:
    """
    Gets the console for the user.

    Uses Azure REST API version 2018-10-01.


    :param builtins.str console_name: The name of the console
    :param builtins.str location: The provider location
    """
    __args__ = dict()
    __args__['consoleName'] = console_name
    __args__['location'] = location
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:portal:getConsoleWithLocation', __args__, opts=opts, typ=GetConsoleWithLocationResult).value

    return AwaitableGetConsoleWithLocationResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        properties=pulumi.get(__ret__, 'properties'))
def get_console_with_location_output(console_name: Optional[pulumi.Input[builtins.str]] = None,
                                     location: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetConsoleWithLocationResult]:
    """
    Gets the console for the user.

    Uses Azure REST API version 2018-10-01.


    :param builtins.str console_name: The name of the console
    :param builtins.str location: The provider location
    """
    __args__ = dict()
    __args__['consoleName'] = console_name
    __args__['location'] = location
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:portal:getConsoleWithLocation', __args__, opts=opts, typ=GetConsoleWithLocationResult)
    return __ret__.apply(lambda __response__: GetConsoleWithLocationResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        properties=pulumi.get(__response__, 'properties')))
