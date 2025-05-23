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
    'GetRoleDefinitionResult',
    'AwaitableGetRoleDefinitionResult',
    'get_role_definition',
    'get_role_definition_output',
]

@pulumi.output_type
class GetRoleDefinitionResult:
    """
    Role definition.
    """
    def __init__(__self__, assignable_scopes=None, azure_api_version=None, created_by=None, created_on=None, description=None, id=None, name=None, permissions=None, role_name=None, role_type=None, type=None, updated_by=None, updated_on=None):
        if assignable_scopes and not isinstance(assignable_scopes, list):
            raise TypeError("Expected argument 'assignable_scopes' to be a list")
        pulumi.set(__self__, "assignable_scopes", assignable_scopes)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if created_on and not isinstance(created_on, str):
            raise TypeError("Expected argument 'created_on' to be a str")
        pulumi.set(__self__, "created_on", created_on)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)
        if role_name and not isinstance(role_name, str):
            raise TypeError("Expected argument 'role_name' to be a str")
        pulumi.set(__self__, "role_name", role_name)
        if role_type and not isinstance(role_type, str):
            raise TypeError("Expected argument 'role_type' to be a str")
        pulumi.set(__self__, "role_type", role_type)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_by and not isinstance(updated_by, str):
            raise TypeError("Expected argument 'updated_by' to be a str")
        pulumi.set(__self__, "updated_by", updated_by)
        if updated_on and not isinstance(updated_on, str):
            raise TypeError("Expected argument 'updated_on' to be a str")
        pulumi.set(__self__, "updated_on", updated_on)

    @property
    @pulumi.getter(name="assignableScopes")
    def assignable_scopes(self) -> Optional[Sequence[builtins.str]]:
        """
        Role definition assignable scopes.
        """
        return pulumi.get(self, "assignable_scopes")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> builtins.str:
        """
        Id of the user who created the assignment
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> builtins.str:
        """
        Time it was created
        """
        return pulumi.get(self, "created_on")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The role definition description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The role definition ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The role definition name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> Optional[Sequence['outputs.PermissionResponse']]:
        """
        Role definition permissions.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[builtins.str]:
        """
        The role name.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter(name="roleType")
    def role_type(self) -> Optional[builtins.str]:
        """
        The role type.
        """
        return pulumi.get(self, "role_type")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The role definition type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> builtins.str:
        """
        Id of the user who updated the assignment
        """
        return pulumi.get(self, "updated_by")

    @property
    @pulumi.getter(name="updatedOn")
    def updated_on(self) -> builtins.str:
        """
        Time it was updated
        """
        return pulumi.get(self, "updated_on")


class AwaitableGetRoleDefinitionResult(GetRoleDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleDefinitionResult(
            assignable_scopes=self.assignable_scopes,
            azure_api_version=self.azure_api_version,
            created_by=self.created_by,
            created_on=self.created_on,
            description=self.description,
            id=self.id,
            name=self.name,
            permissions=self.permissions,
            role_name=self.role_name,
            role_type=self.role_type,
            type=self.type,
            updated_by=self.updated_by,
            updated_on=self.updated_on)


def get_role_definition(role_definition_id: Optional[builtins.str] = None,
                        scope: Optional[builtins.str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleDefinitionResult:
    """
    Get role definition by ID (GUID).

    Uses Azure REST API version 2022-05-01-preview.

    Other available API versions: 2022-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str role_definition_id: The ID of the role definition.
    :param builtins.str scope: The scope of the operation or resource. Valid scopes are: subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
    """
    __args__ = dict()
    __args__['roleDefinitionId'] = role_definition_id
    __args__['scope'] = scope
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:authorization:getRoleDefinition', __args__, opts=opts, typ=GetRoleDefinitionResult).value

    return AwaitableGetRoleDefinitionResult(
        assignable_scopes=pulumi.get(__ret__, 'assignable_scopes'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        created_by=pulumi.get(__ret__, 'created_by'),
        created_on=pulumi.get(__ret__, 'created_on'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        permissions=pulumi.get(__ret__, 'permissions'),
        role_name=pulumi.get(__ret__, 'role_name'),
        role_type=pulumi.get(__ret__, 'role_type'),
        type=pulumi.get(__ret__, 'type'),
        updated_by=pulumi.get(__ret__, 'updated_by'),
        updated_on=pulumi.get(__ret__, 'updated_on'))
def get_role_definition_output(role_definition_id: Optional[pulumi.Input[builtins.str]] = None,
                               scope: Optional[pulumi.Input[builtins.str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRoleDefinitionResult]:
    """
    Get role definition by ID (GUID).

    Uses Azure REST API version 2022-05-01-preview.

    Other available API versions: 2022-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str role_definition_id: The ID of the role definition.
    :param builtins.str scope: The scope of the operation or resource. Valid scopes are: subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
    """
    __args__ = dict()
    __args__['roleDefinitionId'] = role_definition_id
    __args__['scope'] = scope
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:authorization:getRoleDefinition', __args__, opts=opts, typ=GetRoleDefinitionResult)
    return __ret__.apply(lambda __response__: GetRoleDefinitionResult(
        assignable_scopes=pulumi.get(__response__, 'assignable_scopes'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        created_by=pulumi.get(__response__, 'created_by'),
        created_on=pulumi.get(__response__, 'created_on'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        permissions=pulumi.get(__response__, 'permissions'),
        role_name=pulumi.get(__response__, 'role_name'),
        role_type=pulumi.get(__response__, 'role_type'),
        type=pulumi.get(__response__, 'type'),
        updated_by=pulumi.get(__response__, 'updated_by'),
        updated_on=pulumi.get(__response__, 'updated_on')))
