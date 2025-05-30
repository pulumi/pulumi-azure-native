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
    'ListConnectedEnvironmentsDaprComponentSecretsResult',
    'AwaitableListConnectedEnvironmentsDaprComponentSecretsResult',
    'list_connected_environments_dapr_component_secrets',
    'list_connected_environments_dapr_component_secrets_output',
]

@pulumi.output_type
class ListConnectedEnvironmentsDaprComponentSecretsResult:
    """
    Dapr component Secrets Collection for ListSecrets Action.
    """
    def __init__(__self__, value=None):
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Sequence['outputs.DaprSecretResponse']:
        """
        Collection of secrets used by a Dapr component
        """
        return pulumi.get(self, "value")


class AwaitableListConnectedEnvironmentsDaprComponentSecretsResult(ListConnectedEnvironmentsDaprComponentSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListConnectedEnvironmentsDaprComponentSecretsResult(
            value=self.value)


def list_connected_environments_dapr_component_secrets(component_name: Optional[builtins.str] = None,
                                                       connected_environment_name: Optional[builtins.str] = None,
                                                       resource_group_name: Optional[builtins.str] = None,
                                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListConnectedEnvironmentsDaprComponentSecretsResult:
    """
    Dapr component Secrets Collection for ListSecrets Action.

    Uses Azure REST API version 2024-03-01.

    Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str component_name: Name of the Dapr Component.
    :param builtins.str connected_environment_name: Name of the connected environment.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['componentName'] = component_name
    __args__['connectedEnvironmentName'] = connected_environment_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:app:listConnectedEnvironmentsDaprComponentSecrets', __args__, opts=opts, typ=ListConnectedEnvironmentsDaprComponentSecretsResult).value

    return AwaitableListConnectedEnvironmentsDaprComponentSecretsResult(
        value=pulumi.get(__ret__, 'value'))
def list_connected_environments_dapr_component_secrets_output(component_name: Optional[pulumi.Input[builtins.str]] = None,
                                                              connected_environment_name: Optional[pulumi.Input[builtins.str]] = None,
                                                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListConnectedEnvironmentsDaprComponentSecretsResult]:
    """
    Dapr component Secrets Collection for ListSecrets Action.

    Uses Azure REST API version 2024-03-01.

    Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str component_name: Name of the Dapr Component.
    :param builtins.str connected_environment_name: Name of the connected environment.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['componentName'] = component_name
    __args__['connectedEnvironmentName'] = connected_environment_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:app:listConnectedEnvironmentsDaprComponentSecrets', __args__, opts=opts, typ=ListConnectedEnvironmentsDaprComponentSecretsResult)
    return __ret__.apply(lambda __response__: ListConnectedEnvironmentsDaprComponentSecretsResult(
        value=pulumi.get(__response__, 'value')))
