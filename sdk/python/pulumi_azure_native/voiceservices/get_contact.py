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
    'GetContactResult',
    'AwaitableGetContactResult',
    'get_contact',
    'get_contact_output',
]

@pulumi.output_type
class GetContactResult:
    """
    A Contact resource
    """
    def __init__(__self__, azure_api_version=None, contact_name=None, email=None, id=None, location=None, name=None, phone_number=None, provisioning_state=None, role=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if contact_name and not isinstance(contact_name, str):
            raise TypeError("Expected argument 'contact_name' to be a str")
        pulumi.set(__self__, "contact_name", contact_name)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if phone_number and not isinstance(phone_number, str):
            raise TypeError("Expected argument 'phone_number' to be a str")
        pulumi.set(__self__, "phone_number", phone_number)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
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
    @pulumi.getter(name="contactName")
    def contact_name(self) -> builtins.str:
        """
        Full name of contact
        """
        return pulumi.get(self, "contact_name")

    @property
    @pulumi.getter
    def email(self) -> builtins.str:
        """
        Email address of contact
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> builtins.str:
        """
        Telephone number of contact
        """
        return pulumi.get(self, "phone_number")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Resource provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def role(self) -> builtins.str:
        """
        Job title of contact
        """
        return pulumi.get(self, "role")

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
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetContactResult(GetContactResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContactResult(
            azure_api_version=self.azure_api_version,
            contact_name=self.contact_name,
            email=self.email,
            id=self.id,
            location=self.location,
            name=self.name,
            phone_number=self.phone_number,
            provisioning_state=self.provisioning_state,
            role=self.role,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_contact(communications_gateway_name: Optional[builtins.str] = None,
                contact_name: Optional[builtins.str] = None,
                resource_group_name: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContactResult:
    """
    Get a Contact

    Uses Azure REST API version 2022-12-01-preview.


    :param builtins.str communications_gateway_name: Unique identifier for this deployment
    :param builtins.str contact_name: Unique identifier for this contact
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['communicationsGatewayName'] = communications_gateway_name
    __args__['contactName'] = contact_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:voiceservices:getContact', __args__, opts=opts, typ=GetContactResult).value

    return AwaitableGetContactResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        contact_name=pulumi.get(__ret__, 'contact_name'),
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        phone_number=pulumi.get(__ret__, 'phone_number'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        role=pulumi.get(__ret__, 'role'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_contact_output(communications_gateway_name: Optional[pulumi.Input[builtins.str]] = None,
                       contact_name: Optional[pulumi.Input[builtins.str]] = None,
                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetContactResult]:
    """
    Get a Contact

    Uses Azure REST API version 2022-12-01-preview.


    :param builtins.str communications_gateway_name: Unique identifier for this deployment
    :param builtins.str contact_name: Unique identifier for this contact
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['communicationsGatewayName'] = communications_gateway_name
    __args__['contactName'] = contact_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:voiceservices:getContact', __args__, opts=opts, typ=GetContactResult)
    return __ret__.apply(lambda __response__: GetContactResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        contact_name=pulumi.get(__response__, 'contact_name'),
        email=pulumi.get(__response__, 'email'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        phone_number=pulumi.get(__response__, 'phone_number'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        role=pulumi.get(__response__, 'role'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
