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
from ._inputs import *

__all__ = [
    'ListIntegrationAccountKeyVaultKeysResult',
    'AwaitableListIntegrationAccountKeyVaultKeysResult',
    'list_integration_account_key_vault_keys',
    'list_integration_account_key_vault_keys_output',
]

@pulumi.output_type
class ListIntegrationAccountKeyVaultKeysResult:
    """
    Collection of key vault keys.
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
        The skip token.
        """
        return pulumi.get(self, "skip_token")

    @property
    @pulumi.getter
    def value(self) -> Optional[Sequence['outputs.KeyVaultKeyResponse']]:
        """
        The key vault keys.
        """
        return pulumi.get(self, "value")


class AwaitableListIntegrationAccountKeyVaultKeysResult(ListIntegrationAccountKeyVaultKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListIntegrationAccountKeyVaultKeysResult(
            skip_token=self.skip_token,
            value=self.value)


def list_integration_account_key_vault_keys(integration_account_name: Optional[builtins.str] = None,
                                            key_vault: Optional[Union['KeyVaultReference', 'KeyVaultReferenceDict']] = None,
                                            resource_group_name: Optional[builtins.str] = None,
                                            skip_token: Optional[builtins.str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIntegrationAccountKeyVaultKeysResult:
    """
    Gets the integration account's Key Vault keys.

    Uses Azure REST API version 2019-05-01.

    Other available API versions: 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str integration_account_name: The integration account name.
    :param Union['KeyVaultReference', 'KeyVaultReferenceDict'] key_vault: The key vault reference.
    :param builtins.str resource_group_name: The resource group name.
    :param builtins.str skip_token: The skip token.
    """
    __args__ = dict()
    __args__['integrationAccountName'] = integration_account_name
    __args__['keyVault'] = key_vault
    __args__['resourceGroupName'] = resource_group_name
    __args__['skipToken'] = skip_token
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:logic:listIntegrationAccountKeyVaultKeys', __args__, opts=opts, typ=ListIntegrationAccountKeyVaultKeysResult).value

    return AwaitableListIntegrationAccountKeyVaultKeysResult(
        skip_token=pulumi.get(__ret__, 'skip_token'),
        value=pulumi.get(__ret__, 'value'))
def list_integration_account_key_vault_keys_output(integration_account_name: Optional[pulumi.Input[builtins.str]] = None,
                                                   key_vault: Optional[pulumi.Input[Union['KeyVaultReference', 'KeyVaultReferenceDict']]] = None,
                                                   resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                   skip_token: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListIntegrationAccountKeyVaultKeysResult]:
    """
    Gets the integration account's Key Vault keys.

    Uses Azure REST API version 2019-05-01.

    Other available API versions: 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str integration_account_name: The integration account name.
    :param Union['KeyVaultReference', 'KeyVaultReferenceDict'] key_vault: The key vault reference.
    :param builtins.str resource_group_name: The resource group name.
    :param builtins.str skip_token: The skip token.
    """
    __args__ = dict()
    __args__['integrationAccountName'] = integration_account_name
    __args__['keyVault'] = key_vault
    __args__['resourceGroupName'] = resource_group_name
    __args__['skipToken'] = skip_token
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:logic:listIntegrationAccountKeyVaultKeys', __args__, opts=opts, typ=ListIntegrationAccountKeyVaultKeysResult)
    return __ret__.apply(lambda __response__: ListIntegrationAccountKeyVaultKeysResult(
        skip_token=pulumi.get(__response__, 'skip_token'),
        value=pulumi.get(__response__, 'value')))
