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
    'GetVirtualNetworkResult',
    'AwaitableGetVirtualNetworkResult',
    'get_virtual_network',
    'get_virtual_network_output',
]

@pulumi.output_type
class GetVirtualNetworkResult:
    """
    A virtual network.
    """
    def __init__(__self__, allowed_subnets=None, azure_api_version=None, created_date=None, description=None, external_provider_resource_id=None, external_subnets=None, id=None, location=None, name=None, provisioning_state=None, subnet_overrides=None, tags=None, type=None, unique_identifier=None):
        if allowed_subnets and not isinstance(allowed_subnets, list):
            raise TypeError("Expected argument 'allowed_subnets' to be a list")
        pulumi.set(__self__, "allowed_subnets", allowed_subnets)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if external_provider_resource_id and not isinstance(external_provider_resource_id, str):
            raise TypeError("Expected argument 'external_provider_resource_id' to be a str")
        pulumi.set(__self__, "external_provider_resource_id", external_provider_resource_id)
        if external_subnets and not isinstance(external_subnets, list):
            raise TypeError("Expected argument 'external_subnets' to be a list")
        pulumi.set(__self__, "external_subnets", external_subnets)
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
        if subnet_overrides and not isinstance(subnet_overrides, list):
            raise TypeError("Expected argument 'subnet_overrides' to be a list")
        pulumi.set(__self__, "subnet_overrides", subnet_overrides)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if unique_identifier and not isinstance(unique_identifier, str):
            raise TypeError("Expected argument 'unique_identifier' to be a str")
        pulumi.set(__self__, "unique_identifier", unique_identifier)

    @property
    @pulumi.getter(name="allowedSubnets")
    def allowed_subnets(self) -> Optional[Sequence['outputs.SubnetResponse']]:
        """
        The allowed subnets of the virtual network.
        """
        return pulumi.get(self, "allowed_subnets")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> builtins.str:
        """
        The creation date of the virtual network.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The description of the virtual network.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalProviderResourceId")
    def external_provider_resource_id(self) -> Optional[builtins.str]:
        """
        The Microsoft.Network resource identifier of the virtual network.
        """
        return pulumi.get(self, "external_provider_resource_id")

    @property
    @pulumi.getter(name="externalSubnets")
    def external_subnets(self) -> Sequence['outputs.ExternalSubnetResponse']:
        """
        The external subnet properties.
        """
        return pulumi.get(self, "external_subnets")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The identifier of the resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        The location of the resource.
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
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="subnetOverrides")
    def subnet_overrides(self) -> Optional[Sequence['outputs.SubnetOverrideResponse']]:
        """
        The subnet overrides of the virtual network.
        """
        return pulumi.get(self, "subnet_overrides")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> builtins.str:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")


class AwaitableGetVirtualNetworkResult(GetVirtualNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualNetworkResult(
            allowed_subnets=self.allowed_subnets,
            azure_api_version=self.azure_api_version,
            created_date=self.created_date,
            description=self.description,
            external_provider_resource_id=self.external_provider_resource_id,
            external_subnets=self.external_subnets,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            subnet_overrides=self.subnet_overrides,
            tags=self.tags,
            type=self.type,
            unique_identifier=self.unique_identifier)


def get_virtual_network(expand: Optional[builtins.str] = None,
                        lab_name: Optional[builtins.str] = None,
                        name: Optional[builtins.str] = None,
                        resource_group_name: Optional[builtins.str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualNetworkResult:
    """
    Get virtual network.

    Uses Azure REST API version 2018-09-15.


    :param builtins.str expand: Specify the $expand query. Example: 'properties($expand=externalSubnets)'
    :param builtins.str lab_name: The name of the lab.
    :param builtins.str name: The name of the VirtualNetwork
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['labName'] = lab_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:devtestlab:getVirtualNetwork', __args__, opts=opts, typ=GetVirtualNetworkResult).value

    return AwaitableGetVirtualNetworkResult(
        allowed_subnets=pulumi.get(__ret__, 'allowed_subnets'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        created_date=pulumi.get(__ret__, 'created_date'),
        description=pulumi.get(__ret__, 'description'),
        external_provider_resource_id=pulumi.get(__ret__, 'external_provider_resource_id'),
        external_subnets=pulumi.get(__ret__, 'external_subnets'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        subnet_overrides=pulumi.get(__ret__, 'subnet_overrides'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        unique_identifier=pulumi.get(__ret__, 'unique_identifier'))
def get_virtual_network_output(expand: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                               lab_name: Optional[pulumi.Input[builtins.str]] = None,
                               name: Optional[pulumi.Input[builtins.str]] = None,
                               resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVirtualNetworkResult]:
    """
    Get virtual network.

    Uses Azure REST API version 2018-09-15.


    :param builtins.str expand: Specify the $expand query. Example: 'properties($expand=externalSubnets)'
    :param builtins.str lab_name: The name of the lab.
    :param builtins.str name: The name of the VirtualNetwork
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['labName'] = lab_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:devtestlab:getVirtualNetwork', __args__, opts=opts, typ=GetVirtualNetworkResult)
    return __ret__.apply(lambda __response__: GetVirtualNetworkResult(
        allowed_subnets=pulumi.get(__response__, 'allowed_subnets'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        created_date=pulumi.get(__response__, 'created_date'),
        description=pulumi.get(__response__, 'description'),
        external_provider_resource_id=pulumi.get(__response__, 'external_provider_resource_id'),
        external_subnets=pulumi.get(__response__, 'external_subnets'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        subnet_overrides=pulumi.get(__response__, 'subnet_overrides'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        unique_identifier=pulumi.get(__response__, 'unique_identifier')))
