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
    'GetVirtualEndpointResult',
    'AwaitableGetVirtualEndpointResult',
    'get_virtual_endpoint',
    'get_virtual_endpoint_output',
]

@pulumi.output_type
class GetVirtualEndpointResult:
    """
    Pair of virtual endpoints for a flexible server.
    """
    def __init__(__self__, azure_api_version=None, endpoint_type=None, id=None, members=None, name=None, system_data=None, type=None, virtual_endpoints=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if endpoint_type and not isinstance(endpoint_type, str):
            raise TypeError("Expected argument 'endpoint_type' to be a str")
        pulumi.set(__self__, "endpoint_type", endpoint_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_endpoints and not isinstance(virtual_endpoints, list):
            raise TypeError("Expected argument 'virtual_endpoints' to be a list")
        pulumi.set(__self__, "virtual_endpoints", virtual_endpoints)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[builtins.str]:
        """
        Type of endpoint for the virtual endpoints.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def members(self) -> Optional[Sequence[builtins.str]]:
        """
        List of flexible servers that one of the virtual endpoints can refer to.
        """
        return pulumi.get(self, "members")

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

    @property
    @pulumi.getter(name="virtualEndpoints")
    def virtual_endpoints(self) -> Sequence[builtins.str]:
        """
        List of virtual endpoints for a flexible server.
        """
        return pulumi.get(self, "virtual_endpoints")


class AwaitableGetVirtualEndpointResult(GetVirtualEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualEndpointResult(
            azure_api_version=self.azure_api_version,
            endpoint_type=self.endpoint_type,
            id=self.id,
            members=self.members,
            name=self.name,
            system_data=self.system_data,
            type=self.type,
            virtual_endpoints=self.virtual_endpoints)


def get_virtual_endpoint(resource_group_name: Optional[builtins.str] = None,
                         server_name: Optional[builtins.str] = None,
                         virtual_endpoint_name: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualEndpointResult:
    """
    Gets information about a pair of virtual endpoints.

    Uses Azure REST API version 2024-08-01.

    Other available API versions: 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_name: The name of the server.
    :param builtins.str virtual_endpoint_name: Base name of the virtual endpoints.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    __args__['virtualEndpointName'] = virtual_endpoint_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:dbforpostgresql:getVirtualEndpoint', __args__, opts=opts, typ=GetVirtualEndpointResult).value

    return AwaitableGetVirtualEndpointResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        endpoint_type=pulumi.get(__ret__, 'endpoint_type'),
        id=pulumi.get(__ret__, 'id'),
        members=pulumi.get(__ret__, 'members'),
        name=pulumi.get(__ret__, 'name'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'),
        virtual_endpoints=pulumi.get(__ret__, 'virtual_endpoints'))
def get_virtual_endpoint_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                server_name: Optional[pulumi.Input[builtins.str]] = None,
                                virtual_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVirtualEndpointResult]:
    """
    Gets information about a pair of virtual endpoints.

    Uses Azure REST API version 2024-08-01.

    Other available API versions: 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_name: The name of the server.
    :param builtins.str virtual_endpoint_name: Base name of the virtual endpoints.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    __args__['virtualEndpointName'] = virtual_endpoint_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:dbforpostgresql:getVirtualEndpoint', __args__, opts=opts, typ=GetVirtualEndpointResult)
    return __ret__.apply(lambda __response__: GetVirtualEndpointResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        endpoint_type=pulumi.get(__response__, 'endpoint_type'),
        id=pulumi.get(__response__, 'id'),
        members=pulumi.get(__response__, 'members'),
        name=pulumi.get(__response__, 'name'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type'),
        virtual_endpoints=pulumi.get(__response__, 'virtual_endpoints')))
