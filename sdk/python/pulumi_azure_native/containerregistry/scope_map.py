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

__all__ = ['ScopeMapArgs', 'ScopeMap']

@pulumi.input_type
class ScopeMapArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 registry_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 scope_map_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ScopeMap resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] actions: The list of scoped permissions for registry artifacts.
               E.g. repositories/repository-name/content/read,
               repositories/repository-name/metadata/write
        :param pulumi.Input[builtins.str] registry_name: The name of the container registry.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] description: The user friendly description of the scope map.
        :param pulumi.Input[builtins.str] scope_map_name: The name of the scope map.
        """
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "registry_name", registry_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if scope_map_name is not None:
            pulumi.set(__self__, "scope_map_name", scope_map_name)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        The list of scoped permissions for registry artifacts.
        E.g. repositories/repository-name/content/read,
        repositories/repository-name/metadata/write
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter(name="registryName")
    def registry_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the container registry.
        """
        return pulumi.get(self, "registry_name")

    @registry_name.setter
    def registry_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "registry_name", value)

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
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The user friendly description of the scope map.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="scopeMapName")
    def scope_map_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the scope map.
        """
        return pulumi.get(self, "scope_map_name")

    @scope_map_name.setter
    def scope_map_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "scope_map_name", value)


@pulumi.type_token("azure-native:containerregistry:ScopeMap")
class ScopeMap(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 registry_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope_map_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        An object that represents a scope map for a container registry.

        Uses Azure REST API version 2024-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-12-01.

        Other available API versions: 2020-11-01-preview, 2021-06-01-preview, 2021-08-01-preview, 2021-12-01-preview, 2022-02-01-preview, 2022-12-01, 2023-01-01-preview, 2023-06-01-preview, 2023-07-01, 2023-08-01-preview, 2023-11-01-preview, 2025-03-01-preview, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] actions: The list of scoped permissions for registry artifacts.
               E.g. repositories/repository-name/content/read,
               repositories/repository-name/metadata/write
        :param pulumi.Input[builtins.str] description: The user friendly description of the scope map.
        :param pulumi.Input[builtins.str] registry_name: The name of the container registry.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] scope_map_name: The name of the scope map.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScopeMapArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An object that represents a scope map for a container registry.

        Uses Azure REST API version 2024-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-12-01.

        Other available API versions: 2020-11-01-preview, 2021-06-01-preview, 2021-08-01-preview, 2021-12-01-preview, 2022-02-01-preview, 2022-12-01, 2023-01-01-preview, 2023-06-01-preview, 2023-07-01, 2023-08-01-preview, 2023-11-01-preview, 2025-03-01-preview, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ScopeMapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScopeMapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 registry_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope_map_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScopeMapArgs.__new__(ScopeMapArgs)

            if actions is None and not opts.urn:
                raise TypeError("Missing required property 'actions'")
            __props__.__dict__["actions"] = actions
            __props__.__dict__["description"] = description
            if registry_name is None and not opts.urn:
                raise TypeError("Missing required property 'registry_name'")
            __props__.__dict__["registry_name"] = registry_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["scope_map_name"] = scope_map_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["creation_date"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:containerregistry/v20190501preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20201101preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20210601preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20210801preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20211201preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20220201preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20221201:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20230101preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20230601preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20230701:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20230801preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20231101preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20241101preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20250301preview:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20250401:ScopeMap"), pulumi.Alias(type_="azure-native:containerregistry/v20250501preview:ScopeMap")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ScopeMap, __self__).__init__(
            'azure-native:containerregistry:ScopeMap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ScopeMap':
        """
        Get an existing ScopeMap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ScopeMapArgs.__new__(ScopeMapArgs)

        __props__.__dict__["actions"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["creation_date"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return ScopeMap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        The list of scoped permissions for registry artifacts.
        E.g. repositories/repository-name/content/read,
        repositories/repository-name/metadata/write
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[builtins.str]:
        """
        The creation date of scope map.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The user friendly description of the scope map.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

