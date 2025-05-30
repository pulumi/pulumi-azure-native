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
    'GetRegistrationAssignmentResult',
    'AwaitableGetRegistrationAssignmentResult',
    'get_registration_assignment',
    'get_registration_assignment_output',
]

@pulumi.output_type
class GetRegistrationAssignmentResult:
    """
    The registration assignment.
    """
    def __init__(__self__, azure_api_version=None, id=None, name=None, properties=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
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
        The fully qualified path of the registration assignment.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the registration assignment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.RegistrationAssignmentPropertiesResponse':
        """
        The properties of a registration assignment.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The metadata for the registration assignment resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the Azure resource (Microsoft.ManagedServices/registrationAssignments).
        """
        return pulumi.get(self, "type")


class AwaitableGetRegistrationAssignmentResult(GetRegistrationAssignmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistrationAssignmentResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            name=self.name,
            properties=self.properties,
            system_data=self.system_data,
            type=self.type)


def get_registration_assignment(expand_registration_definition: Optional[builtins.bool] = None,
                                registration_assignment_id: Optional[builtins.str] = None,
                                scope: Optional[builtins.str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistrationAssignmentResult:
    """
    Gets the details of the specified registration assignment.

    Uses Azure REST API version 2022-10-01.


    :param builtins.bool expand_registration_definition: The flag indicating whether to return the registration definition details along with the registration assignment details.
    :param builtins.str registration_assignment_id: The GUID of the registration assignment.
    :param builtins.str scope: The scope of the resource.
    """
    __args__ = dict()
    __args__['expandRegistrationDefinition'] = expand_registration_definition
    __args__['registrationAssignmentId'] = registration_assignment_id
    __args__['scope'] = scope
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:managedservices:getRegistrationAssignment', __args__, opts=opts, typ=GetRegistrationAssignmentResult).value

    return AwaitableGetRegistrationAssignmentResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_registration_assignment_output(expand_registration_definition: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                       registration_assignment_id: Optional[pulumi.Input[builtins.str]] = None,
                                       scope: Optional[pulumi.Input[builtins.str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRegistrationAssignmentResult]:
    """
    Gets the details of the specified registration assignment.

    Uses Azure REST API version 2022-10-01.


    :param builtins.bool expand_registration_definition: The flag indicating whether to return the registration definition details along with the registration assignment details.
    :param builtins.str registration_assignment_id: The GUID of the registration assignment.
    :param builtins.str scope: The scope of the resource.
    """
    __args__ = dict()
    __args__['expandRegistrationDefinition'] = expand_registration_definition
    __args__['registrationAssignmentId'] = registration_assignment_id
    __args__['scope'] = scope
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:managedservices:getRegistrationAssignment', __args__, opts=opts, typ=GetRegistrationAssignmentResult)
    return __ret__.apply(lambda __response__: GetRegistrationAssignmentResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
