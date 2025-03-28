# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['SAPVirtualInstanceArgs', 'SAPVirtualInstance']

@pulumi.input_type
class SAPVirtualInstanceArgs:
    def __init__(__self__, *,
                 configuration: pulumi.Input[Union['DeploymentConfigurationArgs', 'DeploymentWithOSConfigurationArgs', 'DiscoveryConfigurationArgs']],
                 environment: pulumi.Input[Union[str, 'SAPEnvironmentType']],
                 resource_group_name: pulumi.Input[str],
                 sap_product: pulumi.Input[Union[str, 'SAPProductType']],
                 identity: Optional[pulumi.Input['UserAssignedServiceIdentityArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_resource_group_configuration: Optional[pulumi.Input['ManagedRGConfigurationArgs']] = None,
                 sap_virtual_instance_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SAPVirtualInstance resource.
        :param pulumi.Input[Union['DeploymentConfigurationArgs', 'DeploymentWithOSConfigurationArgs', 'DiscoveryConfigurationArgs']] configuration: Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
        :param pulumi.Input[Union[str, 'SAPEnvironmentType']] environment: Defines the environment type - Production/Non Production.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union[str, 'SAPProductType']] sap_product: Defines the SAP Product type.
        :param pulumi.Input['UserAssignedServiceIdentityArgs'] identity: A pre-created user assigned identity with appropriate roles assigned. To learn more on identity and roles required, visit the ACSS how-to-guide.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input['ManagedRGConfigurationArgs'] managed_resource_group_configuration: Managed resource group configuration
        :param pulumi.Input[str] sap_virtual_instance_name: The name of the Virtual Instances for SAP solutions resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "configuration", configuration)
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sap_product", sap_product)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed_resource_group_configuration is not None:
            pulumi.set(__self__, "managed_resource_group_configuration", managed_resource_group_configuration)
        if sap_virtual_instance_name is not None:
            pulumi.set(__self__, "sap_virtual_instance_name", sap_virtual_instance_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Input[Union['DeploymentConfigurationArgs', 'DeploymentWithOSConfigurationArgs', 'DiscoveryConfigurationArgs']]:
        """
        Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: pulumi.Input[Union['DeploymentConfigurationArgs', 'DeploymentWithOSConfigurationArgs', 'DiscoveryConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[Union[str, 'SAPEnvironmentType']]:
        """
        Defines the environment type - Production/Non Production.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[Union[str, 'SAPEnvironmentType']]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="sapProduct")
    def sap_product(self) -> pulumi.Input[Union[str, 'SAPProductType']]:
        """
        Defines the SAP Product type.
        """
        return pulumi.get(self, "sap_product")

    @sap_product.setter
    def sap_product(self, value: pulumi.Input[Union[str, 'SAPProductType']]):
        pulumi.set(self, "sap_product", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['UserAssignedServiceIdentityArgs']]:
        """
        A pre-created user assigned identity with appropriate roles assigned. To learn more on identity and roles required, visit the ACSS how-to-guide.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['UserAssignedServiceIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="managedResourceGroupConfiguration")
    def managed_resource_group_configuration(self) -> Optional[pulumi.Input['ManagedRGConfigurationArgs']]:
        """
        Managed resource group configuration
        """
        return pulumi.get(self, "managed_resource_group_configuration")

    @managed_resource_group_configuration.setter
    def managed_resource_group_configuration(self, value: Optional[pulumi.Input['ManagedRGConfigurationArgs']]):
        pulumi.set(self, "managed_resource_group_configuration", value)

    @property
    @pulumi.getter(name="sapVirtualInstanceName")
    def sap_virtual_instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Virtual Instances for SAP solutions resource
        """
        return pulumi.get(self, "sap_virtual_instance_name")

    @sap_virtual_instance_name.setter
    def sap_virtual_instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sap_virtual_instance_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class SAPVirtualInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[Union[Union['DeploymentConfigurationArgs', 'DeploymentConfigurationArgsDict'], Union['DeploymentWithOSConfigurationArgs', 'DeploymentWithOSConfigurationArgsDict'], Union['DiscoveryConfigurationArgs', 'DiscoveryConfigurationArgsDict']]]] = None,
                 environment: Optional[pulumi.Input[Union[str, 'SAPEnvironmentType']]] = None,
                 identity: Optional[pulumi.Input[Union['UserAssignedServiceIdentityArgs', 'UserAssignedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_resource_group_configuration: Optional[pulumi.Input[Union['ManagedRGConfigurationArgs', 'ManagedRGConfigurationArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sap_product: Optional[pulumi.Input[Union[str, 'SAPProductType']]] = None,
                 sap_virtual_instance_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Define the Virtual Instance for SAP solutions resource.

        Uses Azure REST API version 2023-04-01. In version 1.x of the Azure Native provider, it used API version 2021-12-01-preview.

        Other available API versions: 2023-10-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[Union['DeploymentConfigurationArgs', 'DeploymentConfigurationArgsDict'], Union['DeploymentWithOSConfigurationArgs', 'DeploymentWithOSConfigurationArgsDict'], Union['DiscoveryConfigurationArgs', 'DiscoveryConfigurationArgsDict']]] configuration: Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
        :param pulumi.Input[Union[str, 'SAPEnvironmentType']] environment: Defines the environment type - Production/Non Production.
        :param pulumi.Input[Union['UserAssignedServiceIdentityArgs', 'UserAssignedServiceIdentityArgsDict']] identity: A pre-created user assigned identity with appropriate roles assigned. To learn more on identity and roles required, visit the ACSS how-to-guide.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[Union['ManagedRGConfigurationArgs', 'ManagedRGConfigurationArgsDict']] managed_resource_group_configuration: Managed resource group configuration
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union[str, 'SAPProductType']] sap_product: Defines the SAP Product type.
        :param pulumi.Input[str] sap_virtual_instance_name: The name of the Virtual Instances for SAP solutions resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SAPVirtualInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Define the Virtual Instance for SAP solutions resource.

        Uses Azure REST API version 2023-04-01. In version 1.x of the Azure Native provider, it used API version 2021-12-01-preview.

        Other available API versions: 2023-10-01-preview.

        :param str resource_name: The name of the resource.
        :param SAPVirtualInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SAPVirtualInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[Union[Union['DeploymentConfigurationArgs', 'DeploymentConfigurationArgsDict'], Union['DeploymentWithOSConfigurationArgs', 'DeploymentWithOSConfigurationArgsDict'], Union['DiscoveryConfigurationArgs', 'DiscoveryConfigurationArgsDict']]]] = None,
                 environment: Optional[pulumi.Input[Union[str, 'SAPEnvironmentType']]] = None,
                 identity: Optional[pulumi.Input[Union['UserAssignedServiceIdentityArgs', 'UserAssignedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_resource_group_configuration: Optional[pulumi.Input[Union['ManagedRGConfigurationArgs', 'ManagedRGConfigurationArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sap_product: Optional[pulumi.Input[Union[str, 'SAPProductType']]] = None,
                 sap_virtual_instance_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SAPVirtualInstanceArgs.__new__(SAPVirtualInstanceArgs)

            if configuration is None and not opts.urn:
                raise TypeError("Missing required property 'configuration'")
            __props__.__dict__["configuration"] = configuration
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            __props__.__dict__["identity"] = identity
            __props__.__dict__["location"] = location
            __props__.__dict__["managed_resource_group_configuration"] = managed_resource_group_configuration
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if sap_product is None and not opts.urn:
                raise TypeError("Missing required property 'sap_product'")
            __props__.__dict__["sap_product"] = sap_product
            __props__.__dict__["sap_virtual_instance_name"] = sap_virtual_instance_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["errors"] = None
            __props__.__dict__["health"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:workloads/v20211201preview:SAPVirtualInstance"), pulumi.Alias(type_="azure-native:workloads/v20221101preview:SAPVirtualInstance"), pulumi.Alias(type_="azure-native:workloads/v20230401:SAPVirtualInstance"), pulumi.Alias(type_="azure-native:workloads/v20231001preview:SAPVirtualInstance"), pulumi.Alias(type_="azure-native:workloads/v20240901:SAPVirtualInstance")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SAPVirtualInstance, __self__).__init__(
            'azure-native:workloads:SAPVirtualInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SAPVirtualInstance':
        """
        Get an existing SAPVirtualInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SAPVirtualInstanceArgs.__new__(SAPVirtualInstanceArgs)

        __props__.__dict__["configuration"] = None
        __props__.__dict__["environment"] = None
        __props__.__dict__["errors"] = None
        __props__.__dict__["health"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["managed_resource_group_configuration"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["sap_product"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return SAPVirtualInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[Any]:
        """
        Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        Defines the environment type - Production/Non Production.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def errors(self) -> pulumi.Output['outputs.SAPVirtualInstanceErrorResponse']:
        """
        Indicates any errors on the Virtual Instance for SAP solutions resource.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter
    def health(self) -> pulumi.Output[str]:
        """
        Defines the health of SAP Instances.
        """
        return pulumi.get(self, "health")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.UserAssignedServiceIdentityResponse']]:
        """
        A pre-created user assigned identity with appropriate roles assigned. To learn more on identity and roles required, visit the ACSS how-to-guide.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedResourceGroupConfiguration")
    def managed_resource_group_configuration(self) -> pulumi.Output[Optional['outputs.ManagedRGConfigurationResponse']]:
        """
        Managed resource group configuration
        """
        return pulumi.get(self, "managed_resource_group_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Defines the provisioning states.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="sapProduct")
    def sap_product(self) -> pulumi.Output[str]:
        """
        Defines the SAP Product type.
        """
        return pulumi.get(self, "sap_product")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Defines the Virtual Instance for SAP state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Defines the SAP Instance status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

