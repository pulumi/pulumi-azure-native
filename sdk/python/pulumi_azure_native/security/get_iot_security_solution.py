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
    'GetIotSecuritySolutionResult',
    'AwaitableGetIotSecuritySolutionResult',
    'get_iot_security_solution',
    'get_iot_security_solution_output',
]

@pulumi.output_type
class GetIotSecuritySolutionResult:
    """
    IoT Security solution configuration and resource information.
    """
    def __init__(__self__, additional_workspaces=None, auto_discovered_resources=None, azure_api_version=None, disabled_data_sources=None, display_name=None, export=None, id=None, iot_hubs=None, location=None, name=None, recommendations_configuration=None, status=None, system_data=None, tags=None, type=None, unmasked_ip_logging_status=None, user_defined_resources=None, workspace=None):
        if additional_workspaces and not isinstance(additional_workspaces, list):
            raise TypeError("Expected argument 'additional_workspaces' to be a list")
        pulumi.set(__self__, "additional_workspaces", additional_workspaces)
        if auto_discovered_resources and not isinstance(auto_discovered_resources, list):
            raise TypeError("Expected argument 'auto_discovered_resources' to be a list")
        pulumi.set(__self__, "auto_discovered_resources", auto_discovered_resources)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if disabled_data_sources and not isinstance(disabled_data_sources, list):
            raise TypeError("Expected argument 'disabled_data_sources' to be a list")
        pulumi.set(__self__, "disabled_data_sources", disabled_data_sources)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if export and not isinstance(export, list):
            raise TypeError("Expected argument 'export' to be a list")
        pulumi.set(__self__, "export", export)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iot_hubs and not isinstance(iot_hubs, list):
            raise TypeError("Expected argument 'iot_hubs' to be a list")
        pulumi.set(__self__, "iot_hubs", iot_hubs)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if recommendations_configuration and not isinstance(recommendations_configuration, list):
            raise TypeError("Expected argument 'recommendations_configuration' to be a list")
        pulumi.set(__self__, "recommendations_configuration", recommendations_configuration)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if unmasked_ip_logging_status and not isinstance(unmasked_ip_logging_status, str):
            raise TypeError("Expected argument 'unmasked_ip_logging_status' to be a str")
        pulumi.set(__self__, "unmasked_ip_logging_status", unmasked_ip_logging_status)
        if user_defined_resources and not isinstance(user_defined_resources, dict):
            raise TypeError("Expected argument 'user_defined_resources' to be a dict")
        pulumi.set(__self__, "user_defined_resources", user_defined_resources)
        if workspace and not isinstance(workspace, str):
            raise TypeError("Expected argument 'workspace' to be a str")
        pulumi.set(__self__, "workspace", workspace)

    @property
    @pulumi.getter(name="additionalWorkspaces")
    def additional_workspaces(self) -> Optional[Sequence['outputs.AdditionalWorkspacesPropertiesResponse']]:
        """
        List of additional workspaces
        """
        return pulumi.get(self, "additional_workspaces")

    @property
    @pulumi.getter(name="autoDiscoveredResources")
    def auto_discovered_resources(self) -> Sequence[builtins.str]:
        """
        List of resources that were automatically discovered as relevant to the security solution.
        """
        return pulumi.get(self, "auto_discovered_resources")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="disabledDataSources")
    def disabled_data_sources(self) -> Optional[Sequence[builtins.str]]:
        """
        Disabled data sources. Disabling these data sources compromises the system.
        """
        return pulumi.get(self, "disabled_data_sources")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> builtins.str:
        """
        Resource display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def export(self) -> Optional[Sequence[builtins.str]]:
        """
        List of additional options for exporting to workspace data.
        """
        return pulumi.get(self, "export")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="iotHubs")
    def iot_hubs(self) -> Sequence[builtins.str]:
        """
        IoT Hub resource IDs
        """
        return pulumi.get(self, "iot_hubs")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="recommendationsConfiguration")
    def recommendations_configuration(self) -> Optional[Sequence['outputs.RecommendationConfigurationPropertiesResponse']]:
        """
        List of the configuration status for each recommendation type.
        """
        return pulumi.get(self, "recommendations_configuration")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        """
        Status of the IoT Security solution.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="unmaskedIpLoggingStatus")
    def unmasked_ip_logging_status(self) -> Optional[builtins.str]:
        """
        Unmasked IP address logging status
        """
        return pulumi.get(self, "unmasked_ip_logging_status")

    @property
    @pulumi.getter(name="userDefinedResources")
    def user_defined_resources(self) -> Optional['outputs.UserDefinedResourcesPropertiesResponse']:
        """
        Properties of the IoT Security solution's user defined resources.
        """
        return pulumi.get(self, "user_defined_resources")

    @property
    @pulumi.getter
    def workspace(self) -> Optional[builtins.str]:
        """
        Workspace resource ID
        """
        return pulumi.get(self, "workspace")


class AwaitableGetIotSecuritySolutionResult(GetIotSecuritySolutionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIotSecuritySolutionResult(
            additional_workspaces=self.additional_workspaces,
            auto_discovered_resources=self.auto_discovered_resources,
            azure_api_version=self.azure_api_version,
            disabled_data_sources=self.disabled_data_sources,
            display_name=self.display_name,
            export=self.export,
            id=self.id,
            iot_hubs=self.iot_hubs,
            location=self.location,
            name=self.name,
            recommendations_configuration=self.recommendations_configuration,
            status=self.status,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            unmasked_ip_logging_status=self.unmasked_ip_logging_status,
            user_defined_resources=self.user_defined_resources,
            workspace=self.workspace)


def get_iot_security_solution(resource_group_name: Optional[builtins.str] = None,
                              solution_name: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIotSecuritySolutionResult:
    """
    User this method to get details of a specific IoT Security solution based on solution name

    Uses Azure REST API version 2019-08-01.

    Other available API versions: 2017-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native security [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    :param builtins.str solution_name: The name of the IoT Security solution.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['solutionName'] = solution_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:security:getIotSecuritySolution', __args__, opts=opts, typ=GetIotSecuritySolutionResult).value

    return AwaitableGetIotSecuritySolutionResult(
        additional_workspaces=pulumi.get(__ret__, 'additional_workspaces'),
        auto_discovered_resources=pulumi.get(__ret__, 'auto_discovered_resources'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        disabled_data_sources=pulumi.get(__ret__, 'disabled_data_sources'),
        display_name=pulumi.get(__ret__, 'display_name'),
        export=pulumi.get(__ret__, 'export'),
        id=pulumi.get(__ret__, 'id'),
        iot_hubs=pulumi.get(__ret__, 'iot_hubs'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        recommendations_configuration=pulumi.get(__ret__, 'recommendations_configuration'),
        status=pulumi.get(__ret__, 'status'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        unmasked_ip_logging_status=pulumi.get(__ret__, 'unmasked_ip_logging_status'),
        user_defined_resources=pulumi.get(__ret__, 'user_defined_resources'),
        workspace=pulumi.get(__ret__, 'workspace'))
def get_iot_security_solution_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                     solution_name: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIotSecuritySolutionResult]:
    """
    User this method to get details of a specific IoT Security solution based on solution name

    Uses Azure REST API version 2019-08-01.

    Other available API versions: 2017-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native security [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    :param builtins.str solution_name: The name of the IoT Security solution.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['solutionName'] = solution_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:security:getIotSecuritySolution', __args__, opts=opts, typ=GetIotSecuritySolutionResult)
    return __ret__.apply(lambda __response__: GetIotSecuritySolutionResult(
        additional_workspaces=pulumi.get(__response__, 'additional_workspaces'),
        auto_discovered_resources=pulumi.get(__response__, 'auto_discovered_resources'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        disabled_data_sources=pulumi.get(__response__, 'disabled_data_sources'),
        display_name=pulumi.get(__response__, 'display_name'),
        export=pulumi.get(__response__, 'export'),
        id=pulumi.get(__response__, 'id'),
        iot_hubs=pulumi.get(__response__, 'iot_hubs'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        recommendations_configuration=pulumi.get(__response__, 'recommendations_configuration'),
        status=pulumi.get(__response__, 'status'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        unmasked_ip_logging_status=pulumi.get(__response__, 'unmasked_ip_logging_status'),
        user_defined_resources=pulumi.get(__response__, 'user_defined_resources'),
        workspace=pulumi.get(__response__, 'workspace')))
