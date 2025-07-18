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

__all__ = ['DedicatedHostArgs', 'DedicatedHost']

@pulumi.input_type
class DedicatedHostArgs:
    def __init__(__self__, *,
                 host_group_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 sku: pulumi.Input['SkuArgs'],
                 auto_replace_on_failure: Optional[pulumi.Input[builtins.bool]] = None,
                 host_name: Optional[pulumi.Input[builtins.str]] = None,
                 license_type: Optional[pulumi.Input['DedicatedHostLicenseTypes']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 platform_fault_domain: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a DedicatedHost resource.
        :param pulumi.Input[builtins.str] host_group_name: The name of the dedicated host group.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input['SkuArgs'] sku: SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
        :param pulumi.Input[builtins.bool] auto_replace_on_failure: Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
        :param pulumi.Input[builtins.str] host_name: The name of the dedicated host.
        :param pulumi.Input['DedicatedHostLicenseTypes'] license_type: Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.int] platform_fault_domain: Fault domain of the dedicated host within a dedicated host group.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "host_group_name", host_group_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku", sku)
        if auto_replace_on_failure is not None:
            pulumi.set(__self__, "auto_replace_on_failure", auto_replace_on_failure)
        if host_name is not None:
            pulumi.set(__self__, "host_name", host_name)
        if license_type is not None:
            pulumi.set(__self__, "license_type", license_type)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if platform_fault_domain is not None:
            pulumi.set(__self__, "platform_fault_domain", platform_fault_domain)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="hostGroupName")
    def host_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the dedicated host group.
        """
        return pulumi.get(self, "host_group_name")

    @host_group_name.setter
    def host_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "host_group_name", value)

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
    @pulumi.getter
    def sku(self) -> pulumi.Input['SkuArgs']:
        """
        SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input['SkuArgs']):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter(name="autoReplaceOnFailure")
    def auto_replace_on_failure(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
        """
        return pulumi.get(self, "auto_replace_on_failure")

    @auto_replace_on_failure.setter
    def auto_replace_on_failure(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "auto_replace_on_failure", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the dedicated host.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> Optional[pulumi.Input['DedicatedHostLicenseTypes']]:
        """
        Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
        """
        return pulumi.get(self, "license_type")

    @license_type.setter
    def license_type(self, value: Optional[pulumi.Input['DedicatedHostLicenseTypes']]):
        pulumi.set(self, "license_type", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="platformFaultDomain")
    def platform_fault_domain(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Fault domain of the dedicated host within a dedicated host group.
        """
        return pulumi.get(self, "platform_fault_domain")

    @platform_fault_domain.setter
    def platform_fault_domain(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "platform_fault_domain", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("azure-native:compute:DedicatedHost")
class DedicatedHost(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_replace_on_failure: Optional[pulumi.Input[builtins.bool]] = None,
                 host_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 host_name: Optional[pulumi.Input[builtins.str]] = None,
                 license_type: Optional[pulumi.Input['DedicatedHostLicenseTypes']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 platform_fault_domain: Optional[pulumi.Input[builtins.int]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Specifies information about the Dedicated host.

        Uses Azure REST API version 2024-11-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.

        Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] auto_replace_on_failure: Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
        :param pulumi.Input[builtins.str] host_group_name: The name of the dedicated host group.
        :param pulumi.Input[builtins.str] host_name: The name of the dedicated host.
        :param pulumi.Input['DedicatedHostLicenseTypes'] license_type: Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.int] platform_fault_domain: Fault domain of the dedicated host within a dedicated host group.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['SkuArgs', 'SkuArgsDict']] sku: SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DedicatedHostArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Specifies information about the Dedicated host.

        Uses Azure REST API version 2024-11-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.

        Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DedicatedHostArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedHostArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_replace_on_failure: Optional[pulumi.Input[builtins.bool]] = None,
                 host_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 host_name: Optional[pulumi.Input[builtins.str]] = None,
                 license_type: Optional[pulumi.Input['DedicatedHostLicenseTypes']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 platform_fault_domain: Optional[pulumi.Input[builtins.int]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DedicatedHostArgs.__new__(DedicatedHostArgs)

            __props__.__dict__["auto_replace_on_failure"] = auto_replace_on_failure
            if host_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'host_group_name'")
            __props__.__dict__["host_group_name"] = host_group_name
            __props__.__dict__["host_name"] = host_name
            __props__.__dict__["license_type"] = license_type
            __props__.__dict__["location"] = location
            __props__.__dict__["platform_fault_domain"] = platform_fault_domain
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["host_id"] = None
            __props__.__dict__["instance_view"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["provisioning_time"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["time_created"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["virtual_machines"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:compute/v20190301:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20190701:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20191201:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20200601:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20201201:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20210301:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20210401:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20210701:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20211101:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20220301:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20220801:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20221101:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20230301:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20230701:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20230901:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20240301:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20240701:DedicatedHost"), pulumi.Alias(type_="azure-native:compute/v20241101:DedicatedHost")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DedicatedHost, __self__).__init__(
            'azure-native:compute:DedicatedHost',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DedicatedHost':
        """
        Get an existing DedicatedHost resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DedicatedHostArgs.__new__(DedicatedHostArgs)

        __props__.__dict__["auto_replace_on_failure"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["host_id"] = None
        __props__.__dict__["instance_view"] = None
        __props__.__dict__["license_type"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["platform_fault_domain"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["provisioning_time"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["time_created"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["virtual_machines"] = None
        return DedicatedHost(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoReplaceOnFailure")
    def auto_replace_on_failure(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
        """
        return pulumi.get(self, "auto_replace_on_failure")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> pulumi.Output[builtins.str]:
        """
        A unique id generated and assigned to the dedicated host by the platform. Does not change throughout the lifetime of the host.
        """
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter(name="instanceView")
    def instance_view(self) -> pulumi.Output['outputs.DedicatedHostInstanceViewResponse']:
        """
        The dedicated host instance view.
        """
        return pulumi.get(self, "instance_view")

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
        """
        return pulumi.get(self, "license_type")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformFaultDomain")
    def platform_fault_domain(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Fault domain of the dedicated host within a dedicated host group.
        """
        return pulumi.get(self, "platform_fault_domain")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="provisioningTime")
    def provisioning_time(self) -> pulumi.Output[builtins.str]:
        """
        The date when the host was first provisioned.
        """
        return pulumi.get(self, "provisioning_time")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.SkuResponse']:
        """
        SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the time at which the Dedicated Host resource was created. Minimum api-version: 2021-11-01.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualMachines")
    def virtual_machines(self) -> pulumi.Output[Sequence['outputs.SubResourceReadOnlyResponse']]:
        """
        A list of references to all virtual machines in the Dedicated Host.
        """
        return pulumi.get(self, "virtual_machines")

