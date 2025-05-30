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
    'GetUserSharedAccessTokenResult',
    'AwaitableGetUserSharedAccessTokenResult',
    'get_user_shared_access_token',
    'get_user_shared_access_token_output',
]

@pulumi.output_type
class GetUserSharedAccessTokenResult:
    """
    Get User Token response details.
    """
    def __init__(__self__, value=None):
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[builtins.str]:
        """
        Shared Access Authorization token for the User.
        """
        return pulumi.get(self, "value")


class AwaitableGetUserSharedAccessTokenResult(GetUserSharedAccessTokenResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserSharedAccessTokenResult(
            value=self.value)


def get_user_shared_access_token(expiry: Optional[builtins.str] = None,
                                 key_type: Optional['KeyType'] = None,
                                 resource_group_name: Optional[builtins.str] = None,
                                 service_name: Optional[builtins.str] = None,
                                 user_id: Optional[builtins.str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserSharedAccessTokenResult:
    """
    Gets the Shared Access Authorization Token for the User.

    Uses Azure REST API version 2022-09-01-preview.

    Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str expiry: The Expiry time of the Token. Maximum token expiry time is set to 30 days. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
    :param 'KeyType' key_type: The Key to be used to generate token for user.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str service_name: The name of the API Management service.
    :param builtins.str user_id: User identifier. Must be unique in the current API Management service instance.
    """
    __args__ = dict()
    __args__['expiry'] = expiry
    __args__['keyType'] = key_type
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:apimanagement:getUserSharedAccessToken', __args__, opts=opts, typ=GetUserSharedAccessTokenResult).value

    return AwaitableGetUserSharedAccessTokenResult(
        value=pulumi.get(__ret__, 'value'))
def get_user_shared_access_token_output(expiry: Optional[pulumi.Input[builtins.str]] = None,
                                        key_type: Optional[pulumi.Input['KeyType']] = None,
                                        resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                        service_name: Optional[pulumi.Input[builtins.str]] = None,
                                        user_id: Optional[pulumi.Input[builtins.str]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserSharedAccessTokenResult]:
    """
    Gets the Shared Access Authorization Token for the User.

    Uses Azure REST API version 2022-09-01-preview.

    Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str expiry: The Expiry time of the Token. Maximum token expiry time is set to 30 days. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
    :param 'KeyType' key_type: The Key to be used to generate token for user.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str service_name: The name of the API Management service.
    :param builtins.str user_id: User identifier. Must be unique in the current API Management service instance.
    """
    __args__ = dict()
    __args__['expiry'] = expiry
    __args__['keyType'] = key_type
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:apimanagement:getUserSharedAccessToken', __args__, opts=opts, typ=GetUserSharedAccessTokenResult)
    return __ret__.apply(lambda __response__: GetUserSharedAccessTokenResult(
        value=pulumi.get(__response__, 'value')))
