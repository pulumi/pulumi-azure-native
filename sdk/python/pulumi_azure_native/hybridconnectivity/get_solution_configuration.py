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
    'GetSolutionConfigurationResult',
    'AwaitableGetSolutionConfigurationResult',
    'get_solution_configuration',
    'get_solution_configuration_output',
]

@pulumi.output_type
class GetSolutionConfigurationResult:
    """
    Solution Configuration
    """
    def __init__(__self__, azure_api_version=None, id=None, last_sync_time=None, name=None, provisioning_state=None, solution_settings=None, solution_type=None, status=None, status_details=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_sync_time and not isinstance(last_sync_time, str):
            raise TypeError("Expected argument 'last_sync_time' to be a str")
        pulumi.set(__self__, "last_sync_time", last_sync_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if solution_settings and not isinstance(solution_settings, dict):
            raise TypeError("Expected argument 'solution_settings' to be a dict")
        pulumi.set(__self__, "solution_settings", solution_settings)
        if solution_type and not isinstance(solution_type, str):
            raise TypeError("Expected argument 'solution_type' to be a str")
        pulumi.set(__self__, "solution_type", solution_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_details and not isinstance(status_details, str):
            raise TypeError("Expected argument 'status_details' to be a str")
        pulumi.set(__self__, "status_details", status_details)
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
    @pulumi.getter(name="lastSyncTime")
    def last_sync_time(self) -> builtins.str:
        """
        The last time resources were inventoried
        """
        return pulumi.get(self, "last_sync_time")

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
        The resource provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="solutionSettings")
    def solution_settings(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Solution settings
        """
        return pulumi.get(self, "solution_settings")

    @property
    @pulumi.getter(name="solutionType")
    def solution_type(self) -> builtins.str:
        """
        The type of the solution
        """
        return pulumi.get(self, "solution_type")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        The status of solution configurations
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusDetails")
    def status_details(self) -> builtins.str:
        """
        The detailed message of status details
        """
        return pulumi.get(self, "status_details")

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


class AwaitableGetSolutionConfigurationResult(GetSolutionConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSolutionConfigurationResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            last_sync_time=self.last_sync_time,
            name=self.name,
            provisioning_state=self.provisioning_state,
            solution_settings=self.solution_settings,
            solution_type=self.solution_type,
            status=self.status,
            status_details=self.status_details,
            system_data=self.system_data,
            type=self.type)


def get_solution_configuration(resource_uri: Optional[builtins.str] = None,
                               solution_configuration: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSolutionConfigurationResult:
    """
    Get a SolutionConfiguration

    Uses Azure REST API version 2024-12-01.


    :param builtins.str resource_uri: The fully qualified Azure Resource manager identifier of the resource.
    :param builtins.str solution_configuration: Represent Solution Configuration Resource.
    """
    __args__ = dict()
    __args__['resourceUri'] = resource_uri
    __args__['solutionConfiguration'] = solution_configuration
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:hybridconnectivity:getSolutionConfiguration', __args__, opts=opts, typ=GetSolutionConfigurationResult).value

    return AwaitableGetSolutionConfigurationResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        last_sync_time=pulumi.get(__ret__, 'last_sync_time'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        solution_settings=pulumi.get(__ret__, 'solution_settings'),
        solution_type=pulumi.get(__ret__, 'solution_type'),
        status=pulumi.get(__ret__, 'status'),
        status_details=pulumi.get(__ret__, 'status_details'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_solution_configuration_output(resource_uri: Optional[pulumi.Input[builtins.str]] = None,
                                      solution_configuration: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSolutionConfigurationResult]:
    """
    Get a SolutionConfiguration

    Uses Azure REST API version 2024-12-01.


    :param builtins.str resource_uri: The fully qualified Azure Resource manager identifier of the resource.
    :param builtins.str solution_configuration: Represent Solution Configuration Resource.
    """
    __args__ = dict()
    __args__['resourceUri'] = resource_uri
    __args__['solutionConfiguration'] = solution_configuration
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:hybridconnectivity:getSolutionConfiguration', __args__, opts=opts, typ=GetSolutionConfigurationResult)
    return __ret__.apply(lambda __response__: GetSolutionConfigurationResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        last_sync_time=pulumi.get(__response__, 'last_sync_time'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        solution_settings=pulumi.get(__response__, 'solution_settings'),
        solution_type=pulumi.get(__response__, 'solution_type'),
        status=pulumi.get(__response__, 'status'),
        status_details=pulumi.get(__response__, 'status_details'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
