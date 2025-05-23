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
    'ListAccountKeysResult',
    'AwaitableListAccountKeysResult',
    'list_account_keys',
    'list_account_keys_output',
]

@pulumi.output_type
class ListAccountKeysResult:
    """
    The set of keys which can be used to access the Maps REST APIs. Two keys are provided for key rotation without interruption.
    """
    def __init__(__self__, primary_key=None, primary_key_last_updated=None, secondary_key=None, secondary_key_last_updated=None):
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        pulumi.set(__self__, "primary_key", primary_key)
        if primary_key_last_updated and not isinstance(primary_key_last_updated, str):
            raise TypeError("Expected argument 'primary_key_last_updated' to be a str")
        pulumi.set(__self__, "primary_key_last_updated", primary_key_last_updated)
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        pulumi.set(__self__, "secondary_key", secondary_key)
        if secondary_key_last_updated and not isinstance(secondary_key_last_updated, str):
            raise TypeError("Expected argument 'secondary_key_last_updated' to be a str")
        pulumi.set(__self__, "secondary_key_last_updated", secondary_key_last_updated)

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> builtins.str:
        """
        The primary key for accessing the Maps REST APIs.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="primaryKeyLastUpdated")
    def primary_key_last_updated(self) -> builtins.str:
        """
        The last updated date and time of the primary key.
        """
        return pulumi.get(self, "primary_key_last_updated")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> builtins.str:
        """
        The secondary key for accessing the Maps REST APIs.
        """
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter(name="secondaryKeyLastUpdated")
    def secondary_key_last_updated(self) -> builtins.str:
        """
        The last updated date and time of the secondary key.
        """
        return pulumi.get(self, "secondary_key_last_updated")


class AwaitableListAccountKeysResult(ListAccountKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListAccountKeysResult(
            primary_key=self.primary_key,
            primary_key_last_updated=self.primary_key_last_updated,
            secondary_key=self.secondary_key,
            secondary_key_last_updated=self.secondary_key_last_updated)


def list_account_keys(account_name: Optional[builtins.str] = None,
                      resource_group_name: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListAccountKeysResult:
    """
    Get the keys to use with the Maps APIs. A key is used to authenticate and authorize access to the Maps REST APIs. Only one key is needed at a time; two are given to provide seamless key regeneration.

    Uses Azure REST API version 2024-07-01-preview.

    Other available API versions: 2020-02-01-preview, 2021-02-01, 2021-07-01-preview, 2021-12-01-preview, 2023-06-01, 2023-08-01-preview, 2023-12-01-preview, 2024-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maps [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: The name of the Maps Account.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:maps:listAccountKeys', __args__, opts=opts, typ=ListAccountKeysResult).value

    return AwaitableListAccountKeysResult(
        primary_key=pulumi.get(__ret__, 'primary_key'),
        primary_key_last_updated=pulumi.get(__ret__, 'primary_key_last_updated'),
        secondary_key=pulumi.get(__ret__, 'secondary_key'),
        secondary_key_last_updated=pulumi.get(__ret__, 'secondary_key_last_updated'))
def list_account_keys_output(account_name: Optional[pulumi.Input[builtins.str]] = None,
                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListAccountKeysResult]:
    """
    Get the keys to use with the Maps APIs. A key is used to authenticate and authorize access to the Maps REST APIs. Only one key is needed at a time; two are given to provide seamless key regeneration.

    Uses Azure REST API version 2024-07-01-preview.

    Other available API versions: 2020-02-01-preview, 2021-02-01, 2021-07-01-preview, 2021-12-01-preview, 2023-06-01, 2023-08-01-preview, 2023-12-01-preview, 2024-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maps [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: The name of the Maps Account.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:maps:listAccountKeys', __args__, opts=opts, typ=ListAccountKeysResult)
    return __ret__.apply(lambda __response__: ListAccountKeysResult(
        primary_key=pulumi.get(__response__, 'primary_key'),
        primary_key_last_updated=pulumi.get(__response__, 'primary_key_last_updated'),
        secondary_key=pulumi.get(__response__, 'secondary_key'),
        secondary_key_last_updated=pulumi.get(__response__, 'secondary_key_last_updated')))
