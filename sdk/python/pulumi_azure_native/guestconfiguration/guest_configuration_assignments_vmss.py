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
from ._enums import *
from ._inputs import *

__all__ = ['GuestConfigurationAssignmentsVMSSArgs', 'GuestConfigurationAssignmentsVMSS']

@pulumi.input_type
class GuestConfigurationAssignmentsVMSSArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 vmss_name: pulumi.Input[builtins.str],
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input['GuestConfigurationAssignmentPropertiesArgs']] = None):
        """
        The set of arguments for constructing a GuestConfigurationAssignmentsVMSS resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] vmss_name: The name of the virtual machine scale set.
        :param pulumi.Input[builtins.str] location: Region where the VM is located.
        :param pulumi.Input[builtins.str] name: The guest configuration assignment name.
        :param pulumi.Input['GuestConfigurationAssignmentPropertiesArgs'] properties: Properties of the Guest configuration assignment.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "vmss_name", vmss_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="vmssName")
    def vmss_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the virtual machine scale set.
        """
        return pulumi.get(self, "vmss_name")

    @vmss_name.setter
    def vmss_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "vmss_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Region where the VM is located.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The guest configuration assignment name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['GuestConfigurationAssignmentPropertiesArgs']]:
        """
        Properties of the Guest configuration assignment.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['GuestConfigurationAssignmentPropertiesArgs']]):
        pulumi.set(self, "properties", value)


@pulumi.type_token("azure-native:guestconfiguration:GuestConfigurationAssignmentsVMSS")
class GuestConfigurationAssignmentsVMSS(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['GuestConfigurationAssignmentPropertiesArgs', 'GuestConfigurationAssignmentPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 vmss_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Guest configuration assignment is an association between a machine and guest configuration.

        Uses Azure REST API version 2024-04-05. In version 2.x of the Azure Native provider, it used API version 2022-01-25.

        Other available API versions: 2022-01-25. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native guestconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] location: Region where the VM is located.
        :param pulumi.Input[builtins.str] name: The guest configuration assignment name.
        :param pulumi.Input[Union['GuestConfigurationAssignmentPropertiesArgs', 'GuestConfigurationAssignmentPropertiesArgsDict']] properties: Properties of the Guest configuration assignment.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] vmss_name: The name of the virtual machine scale set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GuestConfigurationAssignmentsVMSSArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Guest configuration assignment is an association between a machine and guest configuration.

        Uses Azure REST API version 2024-04-05. In version 2.x of the Azure Native provider, it used API version 2022-01-25.

        Other available API versions: 2022-01-25. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native guestconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param GuestConfigurationAssignmentsVMSSArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GuestConfigurationAssignmentsVMSSArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['GuestConfigurationAssignmentPropertiesArgs', 'GuestConfigurationAssignmentPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 vmss_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GuestConfigurationAssignmentsVMSSArgs.__new__(GuestConfigurationAssignmentsVMSSArgs)

            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["properties"] = properties
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if vmss_name is None and not opts.urn:
                raise TypeError("Missing required property 'vmss_name'")
            __props__.__dict__["vmss_name"] = vmss_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:guestconfiguration/v20220125:GuestConfigurationAssignmentsVMSS"), pulumi.Alias(type_="azure-native:guestconfiguration/v20240405:GuestConfigurationAssignmentsVMSS")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GuestConfigurationAssignmentsVMSS, __self__).__init__(
            'azure-native:guestconfiguration:GuestConfigurationAssignmentsVMSS',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GuestConfigurationAssignmentsVMSS':
        """
        Get an existing GuestConfigurationAssignmentsVMSS resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GuestConfigurationAssignmentsVMSSArgs.__new__(GuestConfigurationAssignmentsVMSSArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return GuestConfigurationAssignmentsVMSS(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Region where the VM is located.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The guest configuration assignment name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.GuestConfigurationAssignmentPropertiesResponse']:
        """
        Properties of the Guest configuration assignment.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

