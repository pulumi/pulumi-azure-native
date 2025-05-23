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
    'GetVariableResult',
    'AwaitableGetVariableResult',
    'get_variable',
    'get_variable_output',
]

@pulumi.output_type
class GetVariableResult:
    """
    Definition of the variable.
    """
    def __init__(__self__, azure_api_version=None, creation_time=None, description=None, id=None, is_encrypted=None, last_modified_time=None, name=None, type=None, value=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_encrypted and not isinstance(is_encrypted, bool):
            raise TypeError("Expected argument 'is_encrypted' to be a bool")
        pulumi.set(__self__, "is_encrypted", is_encrypted)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[builtins.str]:
        """
        Gets or sets the creation time.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        Gets or sets the description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource Id for the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isEncrypted")
    def is_encrypted(self) -> Optional[builtins.bool]:
        """
        Gets or sets the encrypted flag of the variable.
        """
        return pulumi.get(self, "is_encrypted")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[builtins.str]:
        """
        Gets or sets the last modified time.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional[builtins.str]:
        """
        Gets or sets the value of the variable.
        """
        return pulumi.get(self, "value")


class AwaitableGetVariableResult(GetVariableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVariableResult(
            azure_api_version=self.azure_api_version,
            creation_time=self.creation_time,
            description=self.description,
            id=self.id,
            is_encrypted=self.is_encrypted,
            last_modified_time=self.last_modified_time,
            name=self.name,
            type=self.type,
            value=self.value)


def get_variable(automation_account_name: Optional[builtins.str] = None,
                 resource_group_name: Optional[builtins.str] = None,
                 variable_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVariableResult:
    """
    Retrieve the variable identified by variable name.

    Uses Azure REST API version 2023-11-01.

    Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2022-08-08, 2023-05-15-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str automation_account_name: The name of the automation account.
    :param builtins.str resource_group_name: Name of an Azure Resource group.
    :param builtins.str variable_name: The name of variable.
    """
    __args__ = dict()
    __args__['automationAccountName'] = automation_account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['variableName'] = variable_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:automation:getVariable', __args__, opts=opts, typ=GetVariableResult).value

    return AwaitableGetVariableResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        is_encrypted=pulumi.get(__ret__, 'is_encrypted'),
        last_modified_time=pulumi.get(__ret__, 'last_modified_time'),
        name=pulumi.get(__ret__, 'name'),
        type=pulumi.get(__ret__, 'type'),
        value=pulumi.get(__ret__, 'value'))
def get_variable_output(automation_account_name: Optional[pulumi.Input[builtins.str]] = None,
                        resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                        variable_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVariableResult]:
    """
    Retrieve the variable identified by variable name.

    Uses Azure REST API version 2023-11-01.

    Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2022-08-08, 2023-05-15-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str automation_account_name: The name of the automation account.
    :param builtins.str resource_group_name: Name of an Azure Resource group.
    :param builtins.str variable_name: The name of variable.
    """
    __args__ = dict()
    __args__['automationAccountName'] = automation_account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['variableName'] = variable_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:automation:getVariable', __args__, opts=opts, typ=GetVariableResult)
    return __ret__.apply(lambda __response__: GetVariableResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        creation_time=pulumi.get(__response__, 'creation_time'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        is_encrypted=pulumi.get(__response__, 'is_encrypted'),
        last_modified_time=pulumi.get(__response__, 'last_modified_time'),
        name=pulumi.get(__response__, 'name'),
        type=pulumi.get(__response__, 'type'),
        value=pulumi.get(__response__, 'value')))
