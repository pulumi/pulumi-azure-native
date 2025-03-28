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
    'ListIntegrationAccountCallbackUrlResult',
    'AwaitableListIntegrationAccountCallbackUrlResult',
    'list_integration_account_callback_url',
    'list_integration_account_callback_url_output',
]

@pulumi.output_type
class ListIntegrationAccountCallbackUrlResult:
    def __init__(__self__, value=None):
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The URL value.
        """
        return pulumi.get(self, "value")


class AwaitableListIntegrationAccountCallbackUrlResult(ListIntegrationAccountCallbackUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListIntegrationAccountCallbackUrlResult(
            value=self.value)


def list_integration_account_callback_url(integration_account_name: Optional[str] = None,
                                          not_after: Optional[str] = None,
                                          resource_group_name: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIntegrationAccountCallbackUrlResult:
    """
    Lists the integration account callback URL.


    :param str integration_account_name: The integration account name.
    :param str not_after: The expiry time.
    :param str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['integrationAccountName'] = integration_account_name
    __args__['notAfter'] = not_after
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:logic/v20150801preview:listIntegrationAccountCallbackUrl', __args__, opts=opts, typ=ListIntegrationAccountCallbackUrlResult).value

    return AwaitableListIntegrationAccountCallbackUrlResult(
        value=pulumi.get(__ret__, 'value'))
def list_integration_account_callback_url_output(integration_account_name: Optional[pulumi.Input[str]] = None,
                                                 not_after: Optional[pulumi.Input[Optional[str]]] = None,
                                                 resource_group_name: Optional[pulumi.Input[str]] = None,
                                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListIntegrationAccountCallbackUrlResult]:
    """
    Lists the integration account callback URL.


    :param str integration_account_name: The integration account name.
    :param str not_after: The expiry time.
    :param str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['integrationAccountName'] = integration_account_name
    __args__['notAfter'] = not_after
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:logic/v20150801preview:listIntegrationAccountCallbackUrl', __args__, opts=opts, typ=ListIntegrationAccountCallbackUrlResult)
    return __ret__.apply(lambda __response__: ListIntegrationAccountCallbackUrlResult(
        value=pulumi.get(__response__, 'value')))
