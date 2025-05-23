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
from ._enums import *

__all__ = [
    'ListIntegrationAccountCallbackUrlResult',
    'AwaitableListIntegrationAccountCallbackUrlResult',
    'list_integration_account_callback_url',
    'list_integration_account_callback_url_output',
]

@pulumi.output_type
class ListIntegrationAccountCallbackUrlResult:
    """
    The callback url.
    """
    def __init__(__self__, value=None):
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[builtins.str]:
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


def list_integration_account_callback_url(integration_account_name: Optional[builtins.str] = None,
                                          key_type: Optional[Union[builtins.str, 'KeyType']] = None,
                                          not_after: Optional[builtins.str] = None,
                                          resource_group_name: Optional[builtins.str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIntegrationAccountCallbackUrlResult:
    """
    Gets the integration account callback URL.

    Uses Azure REST API version 2019-05-01.

    Other available API versions: 2015-08-01-preview, 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str integration_account_name: The integration account name.
    :param Union[builtins.str, 'KeyType'] key_type: The key type.
    :param builtins.str not_after: The expiry time.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['integrationAccountName'] = integration_account_name
    __args__['keyType'] = key_type
    __args__['notAfter'] = not_after
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:logic:listIntegrationAccountCallbackUrl', __args__, opts=opts, typ=ListIntegrationAccountCallbackUrlResult).value

    return AwaitableListIntegrationAccountCallbackUrlResult(
        value=pulumi.get(__ret__, 'value'))
def list_integration_account_callback_url_output(integration_account_name: Optional[pulumi.Input[builtins.str]] = None,
                                                 key_type: Optional[pulumi.Input[Optional[Union[builtins.str, 'KeyType']]]] = None,
                                                 not_after: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListIntegrationAccountCallbackUrlResult]:
    """
    Gets the integration account callback URL.

    Uses Azure REST API version 2019-05-01.

    Other available API versions: 2015-08-01-preview, 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str integration_account_name: The integration account name.
    :param Union[builtins.str, 'KeyType'] key_type: The key type.
    :param builtins.str not_after: The expiry time.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['integrationAccountName'] = integration_account_name
    __args__['keyType'] = key_type
    __args__['notAfter'] = not_after
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:logic:listIntegrationAccountCallbackUrl', __args__, opts=opts, typ=ListIntegrationAccountCallbackUrlResult)
    return __ret__.apply(lambda __response__: ListIntegrationAccountCallbackUrlResult(
        value=pulumi.get(__response__, 'value')))
