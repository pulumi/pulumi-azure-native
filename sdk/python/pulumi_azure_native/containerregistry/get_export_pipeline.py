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
    'GetExportPipelineResult',
    'AwaitableGetExportPipelineResult',
    'get_export_pipeline',
    'get_export_pipeline_output',
]

@pulumi.output_type
class GetExportPipelineResult:
    """
    An object that represents an export pipeline for a container registry.
    """
    def __init__(__self__, azure_api_version=None, id=None, identity=None, location=None, name=None, options=None, provisioning_state=None, system_data=None, target=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if options and not isinstance(options, list):
            raise TypeError("Expected argument 'options' to be a list")
        pulumi.set(__self__, "options", options)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if target and not isinstance(target, dict):
            raise TypeError("Expected argument 'target' to be a dict")
        pulumi.set(__self__, "target", target)
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
        The resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.IdentityPropertiesResponse']:
        """
        The identity of the export pipeline.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        The location of the export pipeline.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> Optional[Sequence[builtins.str]]:
        """
        The list of all options configured for the pipeline.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the pipeline at the time the operation was called.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def target(self) -> 'outputs.ExportPipelineTargetPropertiesResponse':
        """
        The target properties of the export pipeline.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetExportPipelineResult(GetExportPipelineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExportPipelineResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            identity=self.identity,
            location=self.location,
            name=self.name,
            options=self.options,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            target=self.target,
            type=self.type)


def get_export_pipeline(export_pipeline_name: Optional[builtins.str] = None,
                        registry_name: Optional[builtins.str] = None,
                        resource_group_name: Optional[builtins.str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExportPipelineResult:
    """
    Gets the properties of the export pipeline.

    Uses Azure REST API version 2023-01-01-preview.

    Other available API versions: 2019-12-01-preview, 2020-11-01-preview, 2021-06-01-preview, 2021-08-01-preview, 2021-12-01-preview, 2022-02-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-11-01-preview, 2024-11-01-preview, 2025-03-01-preview, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str export_pipeline_name: The name of the export pipeline.
    :param builtins.str registry_name: The name of the container registry.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['exportPipelineName'] = export_pipeline_name
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:containerregistry:getExportPipeline', __args__, opts=opts, typ=GetExportPipelineResult).value

    return AwaitableGetExportPipelineResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        options=pulumi.get(__ret__, 'options'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        target=pulumi.get(__ret__, 'target'),
        type=pulumi.get(__ret__, 'type'))
def get_export_pipeline_output(export_pipeline_name: Optional[pulumi.Input[builtins.str]] = None,
                               registry_name: Optional[pulumi.Input[builtins.str]] = None,
                               resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetExportPipelineResult]:
    """
    Gets the properties of the export pipeline.

    Uses Azure REST API version 2023-01-01-preview.

    Other available API versions: 2019-12-01-preview, 2020-11-01-preview, 2021-06-01-preview, 2021-08-01-preview, 2021-12-01-preview, 2022-02-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-11-01-preview, 2024-11-01-preview, 2025-03-01-preview, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str export_pipeline_name: The name of the export pipeline.
    :param builtins.str registry_name: The name of the container registry.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['exportPipelineName'] = export_pipeline_name
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:containerregistry:getExportPipeline', __args__, opts=opts, typ=GetExportPipelineResult)
    return __ret__.apply(lambda __response__: GetExportPipelineResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        options=pulumi.get(__response__, 'options'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        target=pulumi.get(__response__, 'target'),
        type=pulumi.get(__response__, 'type')))
