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
    'GetJobPrivateEndpointResult',
    'AwaitableGetJobPrivateEndpointResult',
    'get_job_private_endpoint',
    'get_job_private_endpoint_output',
]

@pulumi.output_type
class GetJobPrivateEndpointResult:
    """
    A job agent private endpoint.
    """
    def __init__(__self__, azure_api_version=None, id=None, name=None, private_endpoint_id=None, target_server_azure_resource_id=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_endpoint_id and not isinstance(private_endpoint_id, str):
            raise TypeError("Expected argument 'private_endpoint_id' to be a str")
        pulumi.set(__self__, "private_endpoint_id", private_endpoint_id)
        if target_server_azure_resource_id and not isinstance(target_server_azure_resource_id, str):
            raise TypeError("Expected argument 'target_server_azure_resource_id' to be a str")
        pulumi.set(__self__, "target_server_azure_resource_id", target_server_azure_resource_id)
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
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointId")
    def private_endpoint_id(self) -> builtins.str:
        """
        Private endpoint id of the private endpoint.
        """
        return pulumi.get(self, "private_endpoint_id")

    @property
    @pulumi.getter(name="targetServerAzureResourceId")
    def target_server_azure_resource_id(self) -> builtins.str:
        """
        ARM resource id of the server the private endpoint will target.
        """
        return pulumi.get(self, "target_server_azure_resource_id")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetJobPrivateEndpointResult(GetJobPrivateEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobPrivateEndpointResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            name=self.name,
            private_endpoint_id=self.private_endpoint_id,
            target_server_azure_resource_id=self.target_server_azure_resource_id,
            type=self.type)


def get_job_private_endpoint(job_agent_name: Optional[builtins.str] = None,
                             private_endpoint_name: Optional[builtins.str] = None,
                             resource_group_name: Optional[builtins.str] = None,
                             server_name: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobPrivateEndpointResult:
    """
    Gets a private endpoint.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str job_agent_name: The name of the job agent.
    :param builtins.str private_endpoint_name: The name of the private endpoint to get.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['jobAgentName'] = job_agent_name
    __args__['privateEndpointName'] = private_endpoint_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:sql:getJobPrivateEndpoint', __args__, opts=opts, typ=GetJobPrivateEndpointResult).value

    return AwaitableGetJobPrivateEndpointResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        private_endpoint_id=pulumi.get(__ret__, 'private_endpoint_id'),
        target_server_azure_resource_id=pulumi.get(__ret__, 'target_server_azure_resource_id'),
        type=pulumi.get(__ret__, 'type'))
def get_job_private_endpoint_output(job_agent_name: Optional[pulumi.Input[builtins.str]] = None,
                                    private_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                                    resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                    server_name: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetJobPrivateEndpointResult]:
    """
    Gets a private endpoint.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str job_agent_name: The name of the job agent.
    :param builtins.str private_endpoint_name: The name of the private endpoint to get.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['jobAgentName'] = job_agent_name
    __args__['privateEndpointName'] = private_endpoint_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:sql:getJobPrivateEndpoint', __args__, opts=opts, typ=GetJobPrivateEndpointResult)
    return __ret__.apply(lambda __response__: GetJobPrivateEndpointResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        private_endpoint_id=pulumi.get(__response__, 'private_endpoint_id'),
        target_server_azure_resource_id=pulumi.get(__response__, 'target_server_azure_resource_id'),
        type=pulumi.get(__response__, 'type')))
