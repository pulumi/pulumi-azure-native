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
    'GetDotNetComponentResult',
    'AwaitableGetDotNetComponentResult',
    'get_dot_net_component',
    'get_dot_net_component_output',
]

@pulumi.output_type
class GetDotNetComponentResult:
    """
    .NET Component.
    """
    def __init__(__self__, azure_api_version=None, component_type=None, configurations=None, id=None, name=None, provisioning_state=None, service_binds=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if component_type and not isinstance(component_type, str):
            raise TypeError("Expected argument 'component_type' to be a str")
        pulumi.set(__self__, "component_type", component_type)
        if configurations and not isinstance(configurations, list):
            raise TypeError("Expected argument 'configurations' to be a list")
        pulumi.set(__self__, "configurations", configurations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if service_binds and not isinstance(service_binds, list):
            raise TypeError("Expected argument 'service_binds' to be a list")
        pulumi.set(__self__, "service_binds", service_binds)
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
    @pulumi.getter(name="componentType")
    def component_type(self) -> Optional[builtins.str]:
        """
        Type of the .NET Component.
        """
        return pulumi.get(self, "component_type")

    @property
    @pulumi.getter
    def configurations(self) -> Optional[Sequence['outputs.DotNetComponentConfigurationPropertyResponse']]:
        """
        List of .NET Components configuration properties
        """
        return pulumi.get(self, "configurations")

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
        Provisioning state of the .NET Component.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceBinds")
    def service_binds(self) -> Optional[Sequence['outputs.DotNetComponentServiceBindResponse']]:
        """
        List of .NET Components that are bound to the .NET component
        """
        return pulumi.get(self, "service_binds")

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


class AwaitableGetDotNetComponentResult(GetDotNetComponentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDotNetComponentResult(
            azure_api_version=self.azure_api_version,
            component_type=self.component_type,
            configurations=self.configurations,
            id=self.id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            service_binds=self.service_binds,
            system_data=self.system_data,
            type=self.type)


def get_dot_net_component(environment_name: Optional[builtins.str] = None,
                          name: Optional[builtins.str] = None,
                          resource_group_name: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDotNetComponentResult:
    """
    .NET Component.

    Uses Azure REST API version 2024-10-02-preview.

    Other available API versions: 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str environment_name: Name of the Managed Environment.
    :param builtins.str name: Name of the .NET Component.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['environmentName'] = environment_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:app:getDotNetComponent', __args__, opts=opts, typ=GetDotNetComponentResult).value

    return AwaitableGetDotNetComponentResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        component_type=pulumi.get(__ret__, 'component_type'),
        configurations=pulumi.get(__ret__, 'configurations'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        service_binds=pulumi.get(__ret__, 'service_binds'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_dot_net_component_output(environment_name: Optional[pulumi.Input[builtins.str]] = None,
                                 name: Optional[pulumi.Input[builtins.str]] = None,
                                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDotNetComponentResult]:
    """
    .NET Component.

    Uses Azure REST API version 2024-10-02-preview.

    Other available API versions: 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str environment_name: Name of the Managed Environment.
    :param builtins.str name: Name of the .NET Component.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['environmentName'] = environment_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:app:getDotNetComponent', __args__, opts=opts, typ=GetDotNetComponentResult)
    return __ret__.apply(lambda __response__: GetDotNetComponentResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        component_type=pulumi.get(__response__, 'component_type'),
        configurations=pulumi.get(__response__, 'configurations'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        service_binds=pulumi.get(__response__, 'service_binds'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
