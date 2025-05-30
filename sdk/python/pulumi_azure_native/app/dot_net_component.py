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

__all__ = ['DotNetComponentArgs', 'DotNetComponent']

@pulumi.input_type
class DotNetComponentArgs:
    def __init__(__self__, *,
                 environment_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 component_type: Optional[pulumi.Input[Union[builtins.str, 'DotNetComponentType']]] = None,
                 configurations: Optional[pulumi.Input[Sequence[pulumi.Input['DotNetComponentConfigurationPropertyArgs']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 service_binds: Optional[pulumi.Input[Sequence[pulumi.Input['DotNetComponentServiceBindArgs']]]] = None):
        """
        The set of arguments for constructing a DotNetComponent resource.
        :param pulumi.Input[builtins.str] environment_name: Name of the Managed Environment.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union[builtins.str, 'DotNetComponentType']] component_type: Type of the .NET Component.
        :param pulumi.Input[Sequence[pulumi.Input['DotNetComponentConfigurationPropertyArgs']]] configurations: List of .NET Components configuration properties
        :param pulumi.Input[builtins.str] name: Name of the .NET Component.
        :param pulumi.Input[Sequence[pulumi.Input['DotNetComponentServiceBindArgs']]] service_binds: List of .NET Components that are bound to the .NET component
        """
        pulumi.set(__self__, "environment_name", environment_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if component_type is not None:
            pulumi.set(__self__, "component_type", component_type)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_binds is not None:
            pulumi.set(__self__, "service_binds", service_binds)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the Managed Environment.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "environment_name", value)

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
    @pulumi.getter(name="componentType")
    def component_type(self) -> Optional[pulumi.Input[Union[builtins.str, 'DotNetComponentType']]]:
        """
        Type of the .NET Component.
        """
        return pulumi.get(self, "component_type")

    @component_type.setter
    def component_type(self, value: Optional[pulumi.Input[Union[builtins.str, 'DotNetComponentType']]]):
        pulumi.set(self, "component_type", value)

    @property
    @pulumi.getter
    def configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DotNetComponentConfigurationPropertyArgs']]]]:
        """
        List of .NET Components configuration properties
        """
        return pulumi.get(self, "configurations")

    @configurations.setter
    def configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DotNetComponentConfigurationPropertyArgs']]]]):
        pulumi.set(self, "configurations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the .NET Component.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceBinds")
    def service_binds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DotNetComponentServiceBindArgs']]]]:
        """
        List of .NET Components that are bound to the .NET component
        """
        return pulumi.get(self, "service_binds")

    @service_binds.setter
    def service_binds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DotNetComponentServiceBindArgs']]]]):
        pulumi.set(self, "service_binds", value)


@pulumi.type_token("azure-native:app:DotNetComponent")
class DotNetComponent(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 component_type: Optional[pulumi.Input[Union[builtins.str, 'DotNetComponentType']]] = None,
                 configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DotNetComponentConfigurationPropertyArgs', 'DotNetComponentConfigurationPropertyArgsDict']]]]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_binds: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DotNetComponentServiceBindArgs', 'DotNetComponentServiceBindArgsDict']]]]] = None,
                 __props__=None):
        """
        .NET Component.

        Uses Azure REST API version 2024-10-02-preview. In version 2.x of the Azure Native provider, it used API version 2023-11-02-preview.

        Other available API versions: 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[builtins.str, 'DotNetComponentType']] component_type: Type of the .NET Component.
        :param pulumi.Input[Sequence[pulumi.Input[Union['DotNetComponentConfigurationPropertyArgs', 'DotNetComponentConfigurationPropertyArgsDict']]]] configurations: List of .NET Components configuration properties
        :param pulumi.Input[builtins.str] environment_name: Name of the Managed Environment.
        :param pulumi.Input[builtins.str] name: Name of the .NET Component.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Sequence[pulumi.Input[Union['DotNetComponentServiceBindArgs', 'DotNetComponentServiceBindArgsDict']]]] service_binds: List of .NET Components that are bound to the .NET component
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DotNetComponentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        .NET Component.

        Uses Azure REST API version 2024-10-02-preview. In version 2.x of the Azure Native provider, it used API version 2023-11-02-preview.

        Other available API versions: 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DotNetComponentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DotNetComponentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 component_type: Optional[pulumi.Input[Union[builtins.str, 'DotNetComponentType']]] = None,
                 configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DotNetComponentConfigurationPropertyArgs', 'DotNetComponentConfigurationPropertyArgsDict']]]]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_binds: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DotNetComponentServiceBindArgs', 'DotNetComponentServiceBindArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DotNetComponentArgs.__new__(DotNetComponentArgs)

            __props__.__dict__["component_type"] = component_type
            __props__.__dict__["configurations"] = configurations
            if environment_name is None and not opts.urn:
                raise TypeError("Missing required property 'environment_name'")
            __props__.__dict__["environment_name"] = environment_name
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["service_binds"] = service_binds
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:app/v20231102preview:DotNetComponent"), pulumi.Alias(type_="azure-native:app/v20240202preview:DotNetComponent"), pulumi.Alias(type_="azure-native:app/v20240802preview:DotNetComponent"), pulumi.Alias(type_="azure-native:app/v20241002preview:DotNetComponent"), pulumi.Alias(type_="azure-native:app/v20250202preview:DotNetComponent")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DotNetComponent, __self__).__init__(
            'azure-native:app:DotNetComponent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DotNetComponent':
        """
        Get an existing DotNetComponent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DotNetComponentArgs.__new__(DotNetComponentArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["component_type"] = None
        __props__.__dict__["configurations"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["service_binds"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return DotNetComponent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="componentType")
    def component_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Type of the .NET Component.
        """
        return pulumi.get(self, "component_type")

    @property
    @pulumi.getter
    def configurations(self) -> pulumi.Output[Optional[Sequence['outputs.DotNetComponentConfigurationPropertyResponse']]]:
        """
        List of .NET Components configuration properties
        """
        return pulumi.get(self, "configurations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of the .NET Component.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceBinds")
    def service_binds(self) -> pulumi.Output[Optional[Sequence['outputs.DotNetComponentServiceBindResponse']]]:
        """
        List of .NET Components that are bound to the .NET component
        """
        return pulumi.get(self, "service_binds")

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

