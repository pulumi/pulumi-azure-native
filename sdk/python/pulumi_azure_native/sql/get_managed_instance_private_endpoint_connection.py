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
    'GetManagedInstancePrivateEndpointConnectionResult',
    'AwaitableGetManagedInstancePrivateEndpointConnectionResult',
    'get_managed_instance_private_endpoint_connection',
    'get_managed_instance_private_endpoint_connection_output',
]

@pulumi.output_type
class GetManagedInstancePrivateEndpointConnectionResult:
    """
    A private endpoint connection
    """
    def __init__(__self__, azure_api_version=None, id=None, name=None, private_endpoint=None, private_link_service_connection_state=None, provisioning_state=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_endpoint and not isinstance(private_endpoint, dict):
            raise TypeError("Expected argument 'private_endpoint' to be a dict")
        pulumi.set(__self__, "private_endpoint", private_endpoint)
        if private_link_service_connection_state and not isinstance(private_link_service_connection_state, dict):
            raise TypeError("Expected argument 'private_link_service_connection_state' to be a dict")
        pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
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
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> Optional['outputs.ManagedInstancePrivateEndpointPropertyResponse']:
        """
        Private endpoint which the connection belongs to.
        """
        return pulumi.get(self, "private_endpoint")

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> Optional['outputs.ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse']:
        """
        Connection State of the Private Endpoint Connection.
        """
        return pulumi.get(self, "private_link_service_connection_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        State of the Private Endpoint Connection.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetManagedInstancePrivateEndpointConnectionResult(GetManagedInstancePrivateEndpointConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedInstancePrivateEndpointConnectionResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            name=self.name,
            private_endpoint=self.private_endpoint,
            private_link_service_connection_state=self.private_link_service_connection_state,
            provisioning_state=self.provisioning_state,
            type=self.type)


def get_managed_instance_private_endpoint_connection(managed_instance_name: Optional[builtins.str] = None,
                                                     private_endpoint_connection_name: Optional[builtins.str] = None,
                                                     resource_group_name: Optional[builtins.str] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagedInstancePrivateEndpointConnectionResult:
    """
    Gets a private endpoint connection.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str managed_instance_name: The name of the managed instance.
    :param builtins.str private_endpoint_connection_name: The name of the private endpoint connection.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    """
    __args__ = dict()
    __args__['managedInstanceName'] = managed_instance_name
    __args__['privateEndpointConnectionName'] = private_endpoint_connection_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:sql:getManagedInstancePrivateEndpointConnection', __args__, opts=opts, typ=GetManagedInstancePrivateEndpointConnectionResult).value

    return AwaitableGetManagedInstancePrivateEndpointConnectionResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        private_endpoint=pulumi.get(__ret__, 'private_endpoint'),
        private_link_service_connection_state=pulumi.get(__ret__, 'private_link_service_connection_state'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        type=pulumi.get(__ret__, 'type'))
def get_managed_instance_private_endpoint_connection_output(managed_instance_name: Optional[pulumi.Input[builtins.str]] = None,
                                                            private_endpoint_connection_name: Optional[pulumi.Input[builtins.str]] = None,
                                                            resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetManagedInstancePrivateEndpointConnectionResult]:
    """
    Gets a private endpoint connection.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str managed_instance_name: The name of the managed instance.
    :param builtins.str private_endpoint_connection_name: The name of the private endpoint connection.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    """
    __args__ = dict()
    __args__['managedInstanceName'] = managed_instance_name
    __args__['privateEndpointConnectionName'] = private_endpoint_connection_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:sql:getManagedInstancePrivateEndpointConnection', __args__, opts=opts, typ=GetManagedInstancePrivateEndpointConnectionResult)
    return __ret__.apply(lambda __response__: GetManagedInstancePrivateEndpointConnectionResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        private_endpoint=pulumi.get(__response__, 'private_endpoint'),
        private_link_service_connection_state=pulumi.get(__response__, 'private_link_service_connection_state'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        type=pulumi.get(__response__, 'type')))
