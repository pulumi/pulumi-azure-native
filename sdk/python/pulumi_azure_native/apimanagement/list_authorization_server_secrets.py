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

__all__ = [
    'ListAuthorizationServerSecretsResult',
    'AwaitableListAuthorizationServerSecretsResult',
    'list_authorization_server_secrets',
    'list_authorization_server_secrets_output',
]

@pulumi.output_type
class ListAuthorizationServerSecretsResult:
    """
    OAuth Server Secrets Contract.
    """
    def __init__(__self__, client_secret=None, resource_owner_password=None, resource_owner_username=None):
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if resource_owner_password and not isinstance(resource_owner_password, str):
            raise TypeError("Expected argument 'resource_owner_password' to be a str")
        pulumi.set(__self__, "resource_owner_password", resource_owner_password)
        if resource_owner_username and not isinstance(resource_owner_username, str):
            raise TypeError("Expected argument 'resource_owner_username' to be a str")
        pulumi.set(__self__, "resource_owner_username", resource_owner_username)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[builtins.str]:
        """
        oAuth Authorization Server Secrets.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="resourceOwnerPassword")
    def resource_owner_password(self) -> Optional[builtins.str]:
        """
        Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
        """
        return pulumi.get(self, "resource_owner_password")

    @property
    @pulumi.getter(name="resourceOwnerUsername")
    def resource_owner_username(self) -> Optional[builtins.str]:
        """
        Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
        """
        return pulumi.get(self, "resource_owner_username")


class AwaitableListAuthorizationServerSecretsResult(ListAuthorizationServerSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListAuthorizationServerSecretsResult(
            client_secret=self.client_secret,
            resource_owner_password=self.resource_owner_password,
            resource_owner_username=self.resource_owner_username)


def list_authorization_server_secrets(authsid: Optional[builtins.str] = None,
                                      resource_group_name: Optional[builtins.str] = None,
                                      service_name: Optional[builtins.str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListAuthorizationServerSecretsResult:
    """
    Gets the client secret details of the authorization server.

    Uses Azure REST API version 2022-09-01-preview.

    Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str authsid: Identifier of the authorization server.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['authsid'] = authsid
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:apimanagement:listAuthorizationServerSecrets', __args__, opts=opts, typ=ListAuthorizationServerSecretsResult).value

    return AwaitableListAuthorizationServerSecretsResult(
        client_secret=pulumi.get(__ret__, 'client_secret'),
        resource_owner_password=pulumi.get(__ret__, 'resource_owner_password'),
        resource_owner_username=pulumi.get(__ret__, 'resource_owner_username'))
def list_authorization_server_secrets_output(authsid: Optional[pulumi.Input[builtins.str]] = None,
                                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                             service_name: Optional[pulumi.Input[builtins.str]] = None,
                                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListAuthorizationServerSecretsResult]:
    """
    Gets the client secret details of the authorization server.

    Uses Azure REST API version 2022-09-01-preview.

    Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str authsid: Identifier of the authorization server.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['authsid'] = authsid
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:apimanagement:listAuthorizationServerSecrets', __args__, opts=opts, typ=ListAuthorizationServerSecretsResult)
    return __ret__.apply(lambda __response__: ListAuthorizationServerSecretsResult(
        client_secret=pulumi.get(__response__, 'client_secret'),
        resource_owner_password=pulumi.get(__response__, 'resource_owner_password'),
        resource_owner_username=pulumi.get(__response__, 'resource_owner_username')))
