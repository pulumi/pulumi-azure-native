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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    Represents a user who has access to one or more shares on the Data Box Edge/Gateway device.
    """
    def __init__(__self__, azure_api_version=None, encrypted_password=None, id=None, name=None, share_access_rights=None, system_data=None, type=None, user_type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if encrypted_password and not isinstance(encrypted_password, dict):
            raise TypeError("Expected argument 'encrypted_password' to be a dict")
        pulumi.set(__self__, "encrypted_password", encrypted_password)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if share_access_rights and not isinstance(share_access_rights, list):
            raise TypeError("Expected argument 'share_access_rights' to be a list")
        pulumi.set(__self__, "share_access_rights", share_access_rights)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user_type and not isinstance(user_type, str):
            raise TypeError("Expected argument 'user_type' to be a str")
        pulumi.set(__self__, "user_type", user_type)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="encryptedPassword")
    def encrypted_password(self) -> Optional['outputs.AsymmetricEncryptedSecretResponse']:
        """
        The password details.
        """
        return pulumi.get(self, "encrypted_password")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The path ID that uniquely identifies the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The object name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="shareAccessRights")
    def share_access_rights(self) -> Sequence['outputs.ShareAccessRightResponse']:
        """
        List of shares that the user has rights on. This field should not be specified during user creation.
        """
        return pulumi.get(self, "share_access_rights")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Metadata pertaining to creation and last modification of User
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userType")
    def user_type(self) -> builtins.str:
        """
        Type of the user.
        """
        return pulumi.get(self, "user_type")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            azure_api_version=self.azure_api_version,
            encrypted_password=self.encrypted_password,
            id=self.id,
            name=self.name,
            share_access_rights=self.share_access_rights,
            system_data=self.system_data,
            type=self.type,
            user_type=self.user_type)


def get_user(device_name: Optional[builtins.str] = None,
             name: Optional[builtins.str] = None,
             resource_group_name: Optional[builtins.str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Gets the properties of the specified user.

    Uses Azure REST API version 2023-07-01.

    Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str device_name: The device name.
    :param builtins.str name: The user name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:databoxedge:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        encrypted_password=pulumi.get(__ret__, 'encrypted_password'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        share_access_rights=pulumi.get(__ret__, 'share_access_rights'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'),
        user_type=pulumi.get(__ret__, 'user_type'))
def get_user_output(device_name: Optional[pulumi.Input[builtins.str]] = None,
                    name: Optional[pulumi.Input[builtins.str]] = None,
                    resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserResult]:
    """
    Gets the properties of the specified user.

    Uses Azure REST API version 2023-07-01.

    Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str device_name: The device name.
    :param builtins.str name: The user name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:databoxedge:getUser', __args__, opts=opts, typ=GetUserResult)
    return __ret__.apply(lambda __response__: GetUserResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        encrypted_password=pulumi.get(__response__, 'encrypted_password'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        share_access_rights=pulumi.get(__response__, 'share_access_rights'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type'),
        user_type=pulumi.get(__response__, 'user_type')))
