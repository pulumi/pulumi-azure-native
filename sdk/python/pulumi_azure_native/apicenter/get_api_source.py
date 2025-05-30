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
    'GetApiSourceResult',
    'AwaitableGetApiSourceResult',
    'get_api_source',
    'get_api_source_output',
]

@pulumi.output_type
class GetApiSourceResult:
    """
    API source entity.
    """
    def __init__(__self__, azure_api_management_source=None, azure_api_version=None, id=None, import_specification=None, link_state=None, name=None, system_data=None, target_environment_id=None, target_lifecycle_stage=None, type=None):
        if azure_api_management_source and not isinstance(azure_api_management_source, dict):
            raise TypeError("Expected argument 'azure_api_management_source' to be a dict")
        pulumi.set(__self__, "azure_api_management_source", azure_api_management_source)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if import_specification and not isinstance(import_specification, str):
            raise TypeError("Expected argument 'import_specification' to be a str")
        pulumi.set(__self__, "import_specification", import_specification)
        if link_state and not isinstance(link_state, dict):
            raise TypeError("Expected argument 'link_state' to be a dict")
        pulumi.set(__self__, "link_state", link_state)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if target_environment_id and not isinstance(target_environment_id, str):
            raise TypeError("Expected argument 'target_environment_id' to be a str")
        pulumi.set(__self__, "target_environment_id", target_environment_id)
        if target_lifecycle_stage and not isinstance(target_lifecycle_stage, str):
            raise TypeError("Expected argument 'target_lifecycle_stage' to be a str")
        pulumi.set(__self__, "target_lifecycle_stage", target_lifecycle_stage)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="azureApiManagementSource")
    def azure_api_management_source(self) -> Optional['outputs.AzureApiManagementSourceResponse']:
        """
        API source configuration for Azure API Management.
        """
        return pulumi.get(self, "azure_api_management_source")

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
    @pulumi.getter(name="importSpecification")
    def import_specification(self) -> Optional[builtins.str]:
        """
        Indicates if the specification should be imported along with metadata.
        """
        return pulumi.get(self, "import_specification")

    @property
    @pulumi.getter(name="linkState")
    def link_state(self) -> 'outputs.LinkStateResponse':
        """
        The state of the API source link
        """
        return pulumi.get(self, "link_state")

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
    @pulumi.getter(name="targetEnvironmentId")
    def target_environment_id(self) -> Optional[builtins.str]:
        """
        The target environment resource ID.
        """
        return pulumi.get(self, "target_environment_id")

    @property
    @pulumi.getter(name="targetLifecycleStage")
    def target_lifecycle_stage(self) -> Optional[builtins.str]:
        """
        The target lifecycle stage.
        """
        return pulumi.get(self, "target_lifecycle_stage")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetApiSourceResult(GetApiSourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiSourceResult(
            azure_api_management_source=self.azure_api_management_source,
            azure_api_version=self.azure_api_version,
            id=self.id,
            import_specification=self.import_specification,
            link_state=self.link_state,
            name=self.name,
            system_data=self.system_data,
            target_environment_id=self.target_environment_id,
            target_lifecycle_stage=self.target_lifecycle_stage,
            type=self.type)


def get_api_source(api_source_name: Optional[builtins.str] = None,
                   resource_group_name: Optional[builtins.str] = None,
                   service_name: Optional[builtins.str] = None,
                   workspace_name: Optional[builtins.str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiSourceResult:
    """
    Returns details of the API source.

    Uses Azure REST API version 2024-06-01-preview.


    :param builtins.str api_source_name: The name of the API.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str service_name: The name of Azure API Center service.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['apiSourceName'] = api_source_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:apicenter:getApiSource', __args__, opts=opts, typ=GetApiSourceResult).value

    return AwaitableGetApiSourceResult(
        azure_api_management_source=pulumi.get(__ret__, 'azure_api_management_source'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        import_specification=pulumi.get(__ret__, 'import_specification'),
        link_state=pulumi.get(__ret__, 'link_state'),
        name=pulumi.get(__ret__, 'name'),
        system_data=pulumi.get(__ret__, 'system_data'),
        target_environment_id=pulumi.get(__ret__, 'target_environment_id'),
        target_lifecycle_stage=pulumi.get(__ret__, 'target_lifecycle_stage'),
        type=pulumi.get(__ret__, 'type'))
def get_api_source_output(api_source_name: Optional[pulumi.Input[builtins.str]] = None,
                          resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                          service_name: Optional[pulumi.Input[builtins.str]] = None,
                          workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApiSourceResult]:
    """
    Returns details of the API source.

    Uses Azure REST API version 2024-06-01-preview.


    :param builtins.str api_source_name: The name of the API.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str service_name: The name of Azure API Center service.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['apiSourceName'] = api_source_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:apicenter:getApiSource', __args__, opts=opts, typ=GetApiSourceResult)
    return __ret__.apply(lambda __response__: GetApiSourceResult(
        azure_api_management_source=pulumi.get(__response__, 'azure_api_management_source'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        import_specification=pulumi.get(__response__, 'import_specification'),
        link_state=pulumi.get(__response__, 'link_state'),
        name=pulumi.get(__response__, 'name'),
        system_data=pulumi.get(__response__, 'system_data'),
        target_environment_id=pulumi.get(__response__, 'target_environment_id'),
        target_lifecycle_stage=pulumi.get(__response__, 'target_lifecycle_stage'),
        type=pulumi.get(__response__, 'type')))
