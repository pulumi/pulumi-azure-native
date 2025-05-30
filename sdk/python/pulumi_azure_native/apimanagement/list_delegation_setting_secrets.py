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
    'ListDelegationSettingSecretsResult',
    'AwaitableListDelegationSettingSecretsResult',
    'list_delegation_setting_secrets',
    'list_delegation_setting_secrets_output',
]

@pulumi.output_type
class ListDelegationSettingSecretsResult:
    """
    Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
    """
    def __init__(__self__, validation_key=None):
        if validation_key and not isinstance(validation_key, str):
            raise TypeError("Expected argument 'validation_key' to be a str")
        pulumi.set(__self__, "validation_key", validation_key)

    @property
    @pulumi.getter(name="validationKey")
    def validation_key(self) -> Optional[builtins.str]:
        """
        This is secret value of the validation key in portal settings.
        """
        return pulumi.get(self, "validation_key")


class AwaitableListDelegationSettingSecretsResult(ListDelegationSettingSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDelegationSettingSecretsResult(
            validation_key=self.validation_key)


def list_delegation_setting_secrets(resource_group_name: Optional[builtins.str] = None,
                                    service_name: Optional[builtins.str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDelegationSettingSecretsResult:
    """
    Gets the secret validation key of the DelegationSettings.

    Uses Azure REST API version 2021-08-01.

    Other available API versions: 2021-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:apimanagement:listDelegationSettingSecrets', __args__, opts=opts, typ=ListDelegationSettingSecretsResult).value

    return AwaitableListDelegationSettingSecretsResult(
        validation_key=pulumi.get(__ret__, 'validation_key'))
def list_delegation_setting_secrets_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                           service_name: Optional[pulumi.Input[builtins.str]] = None,
                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListDelegationSettingSecretsResult]:
    """
    Gets the secret validation key of the DelegationSettings.

    Uses Azure REST API version 2021-08-01.

    Other available API versions: 2021-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:apimanagement:listDelegationSettingSecrets', __args__, opts=opts, typ=ListDelegationSettingSecretsResult)
    return __ret__.apply(lambda __response__: ListDelegationSettingSecretsResult(
        validation_key=pulumi.get(__response__, 'validation_key')))
