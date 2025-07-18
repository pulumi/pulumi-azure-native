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
    'GetTrustedAccessRoleBindingResult',
    'AwaitableGetTrustedAccessRoleBindingResult',
    'get_trusted_access_role_binding',
    'get_trusted_access_role_binding_output',
]

@pulumi.output_type
class GetTrustedAccessRoleBindingResult:
    """
    Defines binding between a resource and role
    """
    def __init__(__self__, azure_api_version=None, id=None, name=None, provisioning_state=None, roles=None, source_resource_id=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if source_resource_id and not isinstance(source_resource_id, str):
            raise TypeError("Expected argument 'source_resource_id' to be a str")
        pulumi.set(__self__, "source_resource_id", source_resource_id)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The current provisioning state of trusted access role binding.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def roles(self) -> Sequence[builtins.str]:
        """
        A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> builtins.str:
        """
        The ARM resource ID of source resource that trusted access is configured for.
        """
        return pulumi.get(self, "source_resource_id")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetTrustedAccessRoleBindingResult(GetTrustedAccessRoleBindingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrustedAccessRoleBindingResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            roles=self.roles,
            source_resource_id=self.source_resource_id,
            system_data=self.system_data,
            type=self.type)


def get_trusted_access_role_binding(resource_group_name: Optional[builtins.str] = None,
                                    resource_name: Optional[builtins.str] = None,
                                    trusted_access_role_binding_name: Optional[builtins.str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrustedAccessRoleBindingResult:
    """
    Defines binding between a resource and role

    Uses Azure REST API version 2024-10-01.

    Other available API versions: 2022-04-02-preview, 2022-05-02-preview, 2022-06-02-preview, 2022-07-02-preview, 2022-08-02-preview, 2022-08-03-preview, 2022-09-02-preview, 2022-10-02-preview, 2022-11-02-preview, 2023-01-02-preview, 2023-02-02-preview, 2023-03-02-preview, 2023-04-02-preview, 2023-05-02-preview, 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview, 2023-09-01, 2023-09-02-preview, 2023-10-01, 2023-10-02-preview, 2023-11-01, 2023-11-02-preview, 2024-01-01, 2024-01-02-preview, 2024-02-01, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-01, 2024-05-02-preview, 2024-06-02-preview, 2024-07-01, 2024-07-02-preview, 2024-08-01, 2024-09-01, 2024-09-02-preview, 2024-10-02-preview, 2025-01-01, 2025-01-02-preview, 2025-02-01, 2025-02-02-preview, 2025-03-01, 2025-03-02-preview, 2025-04-01, 2025-04-02-preview, 2025-05-01, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str resource_name: The name of the managed cluster resource.
    :param builtins.str trusted_access_role_binding_name: The name of trusted access role binding.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    __args__['trustedAccessRoleBindingName'] = trusted_access_role_binding_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:containerservice:getTrustedAccessRoleBinding', __args__, opts=opts, typ=GetTrustedAccessRoleBindingResult).value

    return AwaitableGetTrustedAccessRoleBindingResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        roles=pulumi.get(__ret__, 'roles'),
        source_resource_id=pulumi.get(__ret__, 'source_resource_id'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_trusted_access_role_binding_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                           resource_name: Optional[pulumi.Input[builtins.str]] = None,
                                           trusted_access_role_binding_name: Optional[pulumi.Input[builtins.str]] = None,
                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTrustedAccessRoleBindingResult]:
    """
    Defines binding between a resource and role

    Uses Azure REST API version 2024-10-01.

    Other available API versions: 2022-04-02-preview, 2022-05-02-preview, 2022-06-02-preview, 2022-07-02-preview, 2022-08-02-preview, 2022-08-03-preview, 2022-09-02-preview, 2022-10-02-preview, 2022-11-02-preview, 2023-01-02-preview, 2023-02-02-preview, 2023-03-02-preview, 2023-04-02-preview, 2023-05-02-preview, 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview, 2023-09-01, 2023-09-02-preview, 2023-10-01, 2023-10-02-preview, 2023-11-01, 2023-11-02-preview, 2024-01-01, 2024-01-02-preview, 2024-02-01, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-01, 2024-05-02-preview, 2024-06-02-preview, 2024-07-01, 2024-07-02-preview, 2024-08-01, 2024-09-01, 2024-09-02-preview, 2024-10-02-preview, 2025-01-01, 2025-01-02-preview, 2025-02-01, 2025-02-02-preview, 2025-03-01, 2025-03-02-preview, 2025-04-01, 2025-04-02-preview, 2025-05-01, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str resource_name: The name of the managed cluster resource.
    :param builtins.str trusted_access_role_binding_name: The name of trusted access role binding.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    __args__['trustedAccessRoleBindingName'] = trusted_access_role_binding_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:containerservice:getTrustedAccessRoleBinding', __args__, opts=opts, typ=GetTrustedAccessRoleBindingResult)
    return __ret__.apply(lambda __response__: GetTrustedAccessRoleBindingResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        roles=pulumi.get(__response__, 'roles'),
        source_resource_id=pulumi.get(__response__, 'source_resource_id'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
