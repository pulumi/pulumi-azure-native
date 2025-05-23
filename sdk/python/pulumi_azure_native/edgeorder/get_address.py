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
    'GetAddressResult',
    'AwaitableGetAddressResult',
    'get_address',
    'get_address_output',
]

@pulumi.output_type
class GetAddressResult:
    """
    Address Resource.
    """
    def __init__(__self__, address_classification=None, address_validation_status=None, azure_api_version=None, contact_details=None, id=None, location=None, name=None, provisioning_state=None, shipping_address=None, system_data=None, tags=None, type=None):
        if address_classification and not isinstance(address_classification, str):
            raise TypeError("Expected argument 'address_classification' to be a str")
        pulumi.set(__self__, "address_classification", address_classification)
        if address_validation_status and not isinstance(address_validation_status, str):
            raise TypeError("Expected argument 'address_validation_status' to be a str")
        pulumi.set(__self__, "address_validation_status", address_validation_status)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if contact_details and not isinstance(contact_details, dict):
            raise TypeError("Expected argument 'contact_details' to be a dict")
        pulumi.set(__self__, "contact_details", contact_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if shipping_address and not isinstance(shipping_address, dict):
            raise TypeError("Expected argument 'shipping_address' to be a dict")
        pulumi.set(__self__, "shipping_address", shipping_address)
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
    @pulumi.getter(name="addressClassification")
    def address_classification(self) -> Optional[builtins.str]:
        """
        Type of address based on its usage context.
        """
        return pulumi.get(self, "address_classification")

    @property
    @pulumi.getter(name="addressValidationStatus")
    def address_validation_status(self) -> builtins.str:
        """
        Status of address validation.
        """
        return pulumi.get(self, "address_validation_status")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="contactDetails")
    def contact_details(self) -> Optional['outputs.ContactDetailsResponse']:
        """
        Contact details for the address.
        """
        return pulumi.get(self, "contact_details")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
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
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="shippingAddress")
    def shipping_address(self) -> Optional['outputs.ShippingAddressResponse']:
        """
        Shipping details for the address.
        """
        return pulumi.get(self, "shipping_address")

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


class AwaitableGetAddressResult(GetAddressResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddressResult(
            address_classification=self.address_classification,
            address_validation_status=self.address_validation_status,
            azure_api_version=self.azure_api_version,
            contact_details=self.contact_details,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            shipping_address=self.shipping_address,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_address(address_name: Optional[builtins.str] = None,
                resource_group_name: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddressResult:
    """
    Get information about the specified address.

    Uses Azure REST API version 2024-02-01.

    Other available API versions: 2022-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native edgeorder [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str address_name: The name of the address Resource within the specified resource group. address names must be between 3 and 24 characters in length and use any alphanumeric and underscore only.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['addressName'] = address_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:edgeorder:getAddress', __args__, opts=opts, typ=GetAddressResult).value

    return AwaitableGetAddressResult(
        address_classification=pulumi.get(__ret__, 'address_classification'),
        address_validation_status=pulumi.get(__ret__, 'address_validation_status'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        contact_details=pulumi.get(__ret__, 'contact_details'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        shipping_address=pulumi.get(__ret__, 'shipping_address'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_address_output(address_name: Optional[pulumi.Input[builtins.str]] = None,
                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAddressResult]:
    """
    Get information about the specified address.

    Uses Azure REST API version 2024-02-01.

    Other available API versions: 2022-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native edgeorder [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str address_name: The name of the address Resource within the specified resource group. address names must be between 3 and 24 characters in length and use any alphanumeric and underscore only.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['addressName'] = address_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:edgeorder:getAddress', __args__, opts=opts, typ=GetAddressResult)
    return __ret__.apply(lambda __response__: GetAddressResult(
        address_classification=pulumi.get(__response__, 'address_classification'),
        address_validation_status=pulumi.get(__response__, 'address_validation_status'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        contact_details=pulumi.get(__response__, 'contact_details'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        shipping_address=pulumi.get(__response__, 'shipping_address'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
