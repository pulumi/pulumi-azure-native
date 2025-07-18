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
    'ListWorkspaceStorageAccountKeysResult',
    'AwaitableListWorkspaceStorageAccountKeysResult',
    'list_workspace_storage_account_keys',
    'list_workspace_storage_account_keys_output',
]

@pulumi.output_type
class ListWorkspaceStorageAccountKeysResult:
    def __init__(__self__, user_storage_key=None):
        if user_storage_key and not isinstance(user_storage_key, str):
            raise TypeError("Expected argument 'user_storage_key' to be a str")
        pulumi.set(__self__, "user_storage_key", user_storage_key)

    @property
    @pulumi.getter(name="userStorageKey")
    def user_storage_key(self) -> builtins.str:
        return pulumi.get(self, "user_storage_key")


class AwaitableListWorkspaceStorageAccountKeysResult(ListWorkspaceStorageAccountKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListWorkspaceStorageAccountKeysResult(
            user_storage_key=self.user_storage_key)


def list_workspace_storage_account_keys(resource_group_name: Optional[builtins.str] = None,
                                        workspace_name: Optional[builtins.str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListWorkspaceStorageAccountKeysResult:
    """
    List storage account keys of a workspace.

    Uses Azure REST API version 2024-10-01.

    Other available API versions: 2021-03-01-preview, 2021-07-01, 2022-01-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: Name of Azure Machine Learning workspace.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:machinelearningservices:listWorkspaceStorageAccountKeys', __args__, opts=opts, typ=ListWorkspaceStorageAccountKeysResult).value

    return AwaitableListWorkspaceStorageAccountKeysResult(
        user_storage_key=pulumi.get(__ret__, 'user_storage_key'))
def list_workspace_storage_account_keys_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                               workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListWorkspaceStorageAccountKeysResult]:
    """
    List storage account keys of a workspace.

    Uses Azure REST API version 2024-10-01.

    Other available API versions: 2021-03-01-preview, 2021-07-01, 2022-01-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: Name of Azure Machine Learning workspace.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:machinelearningservices:listWorkspaceStorageAccountKeys', __args__, opts=opts, typ=ListWorkspaceStorageAccountKeysResult)
    return __ret__.apply(lambda __response__: ListWorkspaceStorageAccountKeysResult(
        user_storage_key=pulumi.get(__response__, 'user_storage_key')))
