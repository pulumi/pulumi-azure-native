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
    'GetRegistryEnvironmentContainerResult',
    'AwaitableGetRegistryEnvironmentContainerResult',
    'get_registry_environment_container',
    'get_registry_environment_container_output',
]

@pulumi.output_type
class GetRegistryEnvironmentContainerResult:
    """
    Azure Resource Manager resource envelope.
    """
    def __init__(__self__, azure_api_version=None, environment_container_properties=None, id=None, name=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if environment_container_properties and not isinstance(environment_container_properties, dict):
            raise TypeError("Expected argument 'environment_container_properties' to be a dict")
        pulumi.set(__self__, "environment_container_properties", environment_container_properties)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="environmentContainerProperties")
    def environment_container_properties(self) -> 'outputs.EnvironmentContainerResponse':
        """
        [Required] Additional attributes of the entity.
        """
        return pulumi.get(self, "environment_container_properties")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
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


class AwaitableGetRegistryEnvironmentContainerResult(GetRegistryEnvironmentContainerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryEnvironmentContainerResult(
            azure_api_version=self.azure_api_version,
            environment_container_properties=self.environment_container_properties,
            id=self.id,
            name=self.name,
            system_data=self.system_data,
            type=self.type)


def get_registry_environment_container(environment_name: Optional[builtins.str] = None,
                                       registry_name: Optional[builtins.str] = None,
                                       resource_group_name: Optional[builtins.str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistryEnvironmentContainerResult:
    """
    Azure Resource Manager resource envelope.

    Uses Azure REST API version 2024-10-01.

    Other available API versions: 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str environment_name: Container name. This is case-sensitive.
    :param builtins.str registry_name: Name of Azure Machine Learning registry. This is case-insensitive
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['environmentName'] = environment_name
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:machinelearningservices:getRegistryEnvironmentContainer', __args__, opts=opts, typ=GetRegistryEnvironmentContainerResult).value

    return AwaitableGetRegistryEnvironmentContainerResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        environment_container_properties=pulumi.get(__ret__, 'environment_container_properties'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_registry_environment_container_output(environment_name: Optional[pulumi.Input[builtins.str]] = None,
                                              registry_name: Optional[pulumi.Input[builtins.str]] = None,
                                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRegistryEnvironmentContainerResult]:
    """
    Azure Resource Manager resource envelope.

    Uses Azure REST API version 2024-10-01.

    Other available API versions: 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str environment_name: Container name. This is case-sensitive.
    :param builtins.str registry_name: Name of Azure Machine Learning registry. This is case-insensitive
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['environmentName'] = environment_name
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:machinelearningservices:getRegistryEnvironmentContainer', __args__, opts=opts, typ=GetRegistryEnvironmentContainerResult)
    return __ret__.apply(lambda __response__: GetRegistryEnvironmentContainerResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        environment_container_properties=pulumi.get(__response__, 'environment_container_properties'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
