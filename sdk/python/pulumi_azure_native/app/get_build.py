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
    'GetBuildResult',
    'AwaitableGetBuildResult',
    'get_build',
    'get_build_output',
]

@pulumi.output_type
class GetBuildResult:
    """
    Information pertaining to an individual build.
    """
    def __init__(__self__, azure_api_version=None, build_status=None, configuration=None, destination_container_registry=None, id=None, log_stream_endpoint=None, name=None, provisioning_state=None, system_data=None, token_endpoint=None, type=None, upload_endpoint=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if build_status and not isinstance(build_status, str):
            raise TypeError("Expected argument 'build_status' to be a str")
        pulumi.set(__self__, "build_status", build_status)
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        pulumi.set(__self__, "configuration", configuration)
        if destination_container_registry and not isinstance(destination_container_registry, dict):
            raise TypeError("Expected argument 'destination_container_registry' to be a dict")
        pulumi.set(__self__, "destination_container_registry", destination_container_registry)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if log_stream_endpoint and not isinstance(log_stream_endpoint, str):
            raise TypeError("Expected argument 'log_stream_endpoint' to be a str")
        pulumi.set(__self__, "log_stream_endpoint", log_stream_endpoint)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if token_endpoint and not isinstance(token_endpoint, str):
            raise TypeError("Expected argument 'token_endpoint' to be a str")
        pulumi.set(__self__, "token_endpoint", token_endpoint)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if upload_endpoint and not isinstance(upload_endpoint, str):
            raise TypeError("Expected argument 'upload_endpoint' to be a str")
        pulumi.set(__self__, "upload_endpoint", upload_endpoint)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="buildStatus")
    def build_status(self) -> builtins.str:
        """
        Status of the build once it has been provisioned.
        """
        return pulumi.get(self, "build_status")

    @property
    @pulumi.getter
    def configuration(self) -> Optional['outputs.BuildConfigurationResponse']:
        """
        Configuration of the build.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="destinationContainerRegistry")
    def destination_container_registry(self) -> Optional['outputs.ContainerRegistryWithCustomImageResponse']:
        """
        Container registry that the final image will be uploaded to.
        """
        return pulumi.get(self, "destination_container_registry")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logStreamEndpoint")
    def log_stream_endpoint(self) -> builtins.str:
        """
        Endpoint from which the build logs can be streamed.
        """
        return pulumi.get(self, "log_stream_endpoint")

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
        Build provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter(name="tokenEndpoint")
    def token_endpoint(self) -> builtins.str:
        """
        Endpoint to use to retrieve an authentication token for log streaming and uploading source code.
        """
        return pulumi.get(self, "token_endpoint")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uploadEndpoint")
    def upload_endpoint(self) -> builtins.str:
        """
        Endpoint to which the source code should be uploaded.
        """
        return pulumi.get(self, "upload_endpoint")


class AwaitableGetBuildResult(GetBuildResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBuildResult(
            azure_api_version=self.azure_api_version,
            build_status=self.build_status,
            configuration=self.configuration,
            destination_container_registry=self.destination_container_registry,
            id=self.id,
            log_stream_endpoint=self.log_stream_endpoint,
            name=self.name,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            token_endpoint=self.token_endpoint,
            type=self.type,
            upload_endpoint=self.upload_endpoint)


def get_build(build_name: Optional[builtins.str] = None,
              builder_name: Optional[builtins.str] = None,
              resource_group_name: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBuildResult:
    """
    Get a BuildResource

    Uses Azure REST API version 2024-10-02-preview.

    Other available API versions: 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str build_name: The name of a build.
    :param builtins.str builder_name: The name of the builder.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['buildName'] = build_name
    __args__['builderName'] = builder_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:app:getBuild', __args__, opts=opts, typ=GetBuildResult).value

    return AwaitableGetBuildResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        build_status=pulumi.get(__ret__, 'build_status'),
        configuration=pulumi.get(__ret__, 'configuration'),
        destination_container_registry=pulumi.get(__ret__, 'destination_container_registry'),
        id=pulumi.get(__ret__, 'id'),
        log_stream_endpoint=pulumi.get(__ret__, 'log_stream_endpoint'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        token_endpoint=pulumi.get(__ret__, 'token_endpoint'),
        type=pulumi.get(__ret__, 'type'),
        upload_endpoint=pulumi.get(__ret__, 'upload_endpoint'))
def get_build_output(build_name: Optional[pulumi.Input[builtins.str]] = None,
                     builder_name: Optional[pulumi.Input[builtins.str]] = None,
                     resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetBuildResult]:
    """
    Get a BuildResource

    Uses Azure REST API version 2024-10-02-preview.

    Other available API versions: 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str build_name: The name of a build.
    :param builtins.str builder_name: The name of the builder.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['buildName'] = build_name
    __args__['builderName'] = builder_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:app:getBuild', __args__, opts=opts, typ=GetBuildResult)
    return __ret__.apply(lambda __response__: GetBuildResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        build_status=pulumi.get(__response__, 'build_status'),
        configuration=pulumi.get(__response__, 'configuration'),
        destination_container_registry=pulumi.get(__response__, 'destination_container_registry'),
        id=pulumi.get(__response__, 'id'),
        log_stream_endpoint=pulumi.get(__response__, 'log_stream_endpoint'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        token_endpoint=pulumi.get(__response__, 'token_endpoint'),
        type=pulumi.get(__response__, 'type'),
        upload_endpoint=pulumi.get(__response__, 'upload_endpoint')))
