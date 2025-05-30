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
    'GetMECRoleResult',
    'AwaitableGetMECRoleResult',
    'get_mec_role',
    'get_mec_role_output',
]

@pulumi.output_type
class GetMECRoleResult:
    """
    MEC role.
    """
    def __init__(__self__, azure_api_version=None, connection_string=None, controller_endpoint=None, id=None, kind=None, name=None, resource_unique_id=None, role_status=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if connection_string and not isinstance(connection_string, dict):
            raise TypeError("Expected argument 'connection_string' to be a dict")
        pulumi.set(__self__, "connection_string", connection_string)
        if controller_endpoint and not isinstance(controller_endpoint, str):
            raise TypeError("Expected argument 'controller_endpoint' to be a str")
        pulumi.set(__self__, "controller_endpoint", controller_endpoint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_unique_id and not isinstance(resource_unique_id, str):
            raise TypeError("Expected argument 'resource_unique_id' to be a str")
        pulumi.set(__self__, "resource_unique_id", resource_unique_id)
        if role_status and not isinstance(role_status, str):
            raise TypeError("Expected argument 'role_status' to be a str")
        pulumi.set(__self__, "role_status", role_status)
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
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional['outputs.AsymmetricEncryptedSecretResponse']:
        """
        Activation key of the MEC.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="controllerEndpoint")
    def controller_endpoint(self) -> Optional[builtins.str]:
        """
        Controller Endpoint.
        """
        return pulumi.get(self, "controller_endpoint")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The path ID that uniquely identifies the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> builtins.str:
        """
        Role type.
        Expected value is 'MEC'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The object name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceUniqueId")
    def resource_unique_id(self) -> Optional[builtins.str]:
        """
        Unique Id of the Resource.
        """
        return pulumi.get(self, "resource_unique_id")

    @property
    @pulumi.getter(name="roleStatus")
    def role_status(self) -> builtins.str:
        """
        Role status.
        """
        return pulumi.get(self, "role_status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Metadata pertaining to creation and last modification of Role
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")


class AwaitableGetMECRoleResult(GetMECRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMECRoleResult(
            azure_api_version=self.azure_api_version,
            connection_string=self.connection_string,
            controller_endpoint=self.controller_endpoint,
            id=self.id,
            kind=self.kind,
            name=self.name,
            resource_unique_id=self.resource_unique_id,
            role_status=self.role_status,
            system_data=self.system_data,
            type=self.type)


def get_mec_role(device_name: Optional[builtins.str] = None,
                 name: Optional[builtins.str] = None,
                 resource_group_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMECRoleResult:
    """
    Gets a specific role by name.

    Uses Azure REST API version 2023-07-01.


    :param builtins.str device_name: The device name.
    :param builtins.str name: The role name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:databoxedge:getMECRole', __args__, opts=opts, typ=GetMECRoleResult).value

    return AwaitableGetMECRoleResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        connection_string=pulumi.get(__ret__, 'connection_string'),
        controller_endpoint=pulumi.get(__ret__, 'controller_endpoint'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        name=pulumi.get(__ret__, 'name'),
        resource_unique_id=pulumi.get(__ret__, 'resource_unique_id'),
        role_status=pulumi.get(__ret__, 'role_status'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_mec_role_output(device_name: Optional[pulumi.Input[builtins.str]] = None,
                        name: Optional[pulumi.Input[builtins.str]] = None,
                        resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMECRoleResult]:
    """
    Gets a specific role by name.

    Uses Azure REST API version 2023-07-01.


    :param builtins.str device_name: The device name.
    :param builtins.str name: The role name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:databoxedge:getMECRole', __args__, opts=opts, typ=GetMECRoleResult)
    return __ret__.apply(lambda __response__: GetMECRoleResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        connection_string=pulumi.get(__response__, 'connection_string'),
        controller_endpoint=pulumi.get(__response__, 'controller_endpoint'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        name=pulumi.get(__response__, 'name'),
        resource_unique_id=pulumi.get(__response__, 'resource_unique_id'),
        role_status=pulumi.get(__response__, 'role_status'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
