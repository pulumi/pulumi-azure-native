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

__all__ = ['ConfigurationAssignmentArgs', 'ConfigurationAssignment']

@pulumi.input_type
class ConfigurationAssignmentArgs:
    def __init__(__self__, *,
                 provider_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 resource_name: pulumi.Input[builtins.str],
                 resource_type: pulumi.Input[builtins.str],
                 configuration_assignment_name: Optional[pulumi.Input[builtins.str]] = None,
                 filter: Optional[pulumi.Input['ConfigurationAssignmentFilterPropertiesArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 maintenance_configuration_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ConfigurationAssignment resource.
        :param pulumi.Input[builtins.str] provider_name: Resource provider name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_name: Resource parent name
        :param pulumi.Input[builtins.str] resource_type: Resource parent type
        :param pulumi.Input[builtins.str] configuration_assignment_name: The name of the ConfigurationAssignment
        :param pulumi.Input['ConfigurationAssignmentFilterPropertiesArgs'] filter: Properties of the configuration assignment
        :param pulumi.Input[builtins.str] location: Location of the resource
        :param pulumi.Input[builtins.str] maintenance_configuration_id: The maintenance configuration Id
        :param pulumi.Input[builtins.str] resource_id: The unique resourceId
        """
        pulumi.set(__self__, "provider_name", provider_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "resource_name", resource_name)
        pulumi.set(__self__, "resource_type", resource_type)
        if configuration_assignment_name is not None:
            pulumi.set(__self__, "configuration_assignment_name", configuration_assignment_name)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if maintenance_configuration_id is not None:
            pulumi.set(__self__, "maintenance_configuration_id", maintenance_configuration_id)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Input[builtins.str]:
        """
        Resource provider name
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "provider_name", value)

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
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Input[builtins.str]:
        """
        Resource parent name
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[builtins.str]:
        """
        Resource parent type
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="configurationAssignmentName")
    def configuration_assignment_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the ConfigurationAssignment
        """
        return pulumi.get(self, "configuration_assignment_name")

    @configuration_assignment_name.setter
    def configuration_assignment_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "configuration_assignment_name", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input['ConfigurationAssignmentFilterPropertiesArgs']]:
        """
        Properties of the configuration assignment
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input['ConfigurationAssignmentFilterPropertiesArgs']]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Location of the resource
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="maintenanceConfigurationId")
    def maintenance_configuration_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The maintenance configuration Id
        """
        return pulumi.get(self, "maintenance_configuration_id")

    @maintenance_configuration_id.setter
    def maintenance_configuration_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "maintenance_configuration_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The unique resourceId
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_id", value)


@pulumi.type_token("azure-native:maintenance:ConfigurationAssignment")
class ConfigurationAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_assignment_name: Optional[pulumi.Input[builtins.str]] = None,
                 filter: Optional[pulumi.Input[Union['ConfigurationAssignmentFilterPropertiesArgs', 'ConfigurationAssignmentFilterPropertiesArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 maintenance_configuration_id: Optional[pulumi.Input[builtins.str]] = None,
                 provider_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Configuration Assignment

        Uses Azure REST API version 2023-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-11-01-preview.

        Other available API versions: 2022-11-01-preview, 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] configuration_assignment_name: The name of the ConfigurationAssignment
        :param pulumi.Input[Union['ConfigurationAssignmentFilterPropertiesArgs', 'ConfigurationAssignmentFilterPropertiesArgsDict']] filter: Properties of the configuration assignment
        :param pulumi.Input[builtins.str] location: Location of the resource
        :param pulumi.Input[builtins.str] maintenance_configuration_id: The maintenance configuration Id
        :param pulumi.Input[builtins.str] provider_name: Resource provider name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_id: The unique resourceId
        :param pulumi.Input[builtins.str] resource_name_: Resource parent name
        :param pulumi.Input[builtins.str] resource_type: Resource parent type
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigurationAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configuration Assignment

        Uses Azure REST API version 2023-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-11-01-preview.

        Other available API versions: 2022-11-01-preview, 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ConfigurationAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_assignment_name: Optional[pulumi.Input[builtins.str]] = None,
                 filter: Optional[pulumi.Input[Union['ConfigurationAssignmentFilterPropertiesArgs', 'ConfigurationAssignmentFilterPropertiesArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 maintenance_configuration_id: Optional[pulumi.Input[builtins.str]] = None,
                 provider_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigurationAssignmentArgs.__new__(ConfigurationAssignmentArgs)

            __props__.__dict__["configuration_assignment_name"] = configuration_assignment_name
            __props__.__dict__["filter"] = filter
            __props__.__dict__["location"] = location
            __props__.__dict__["maintenance_configuration_id"] = maintenance_configuration_id
            if provider_name is None and not opts.urn:
                raise TypeError("Missing required property 'provider_name'")
            __props__.__dict__["provider_name"] = provider_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["resource_id"] = resource_id
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__.__dict__["resource_name"] = resource_name_
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:maintenance/v20210401preview:ConfigurationAssignment"), pulumi.Alias(type_="azure-native:maintenance/v20210901preview:ConfigurationAssignment"), pulumi.Alias(type_="azure-native:maintenance/v20220701preview:ConfigurationAssignment"), pulumi.Alias(type_="azure-native:maintenance/v20221101preview:ConfigurationAssignment"), pulumi.Alias(type_="azure-native:maintenance/v20230401:ConfigurationAssignment"), pulumi.Alias(type_="azure-native:maintenance/v20230901preview:ConfigurationAssignment"), pulumi.Alias(type_="azure-native:maintenance/v20231001preview:ConfigurationAssignment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ConfigurationAssignment, __self__).__init__(
            'azure-native:maintenance:ConfigurationAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConfigurationAssignment':
        """
        Get an existing ConfigurationAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConfigurationAssignmentArgs.__new__(ConfigurationAssignmentArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["filter"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["maintenance_configuration_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["resource_id"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return ConfigurationAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional['outputs.ConfigurationAssignmentFilterPropertiesResponse']]:
        """
        Properties of the configuration assignment
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Location of the resource
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maintenanceConfigurationId")
    def maintenance_configuration_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The maintenance configuration Id
        """
        return pulumi.get(self, "maintenance_configuration_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The unique resourceId
        """
        return pulumi.get(self, "resource_id")

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
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

