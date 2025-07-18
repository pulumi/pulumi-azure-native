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

__all__ = ['DevCenterArgs', 'DevCenter']

@pulumi.input_type
class DevCenterArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 dev_center_name: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 encryption: Optional[pulumi.Input['EncryptionArgs']] = None,
                 identity: Optional[pulumi.Input['ManagedServiceIdentityArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 project_catalog_settings: Optional[pulumi.Input['DevCenterProjectCatalogSettingsArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a DevCenter resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] dev_center_name: The name of the devcenter.
        :param pulumi.Input[builtins.str] display_name: The display name of the devcenter.
        :param pulumi.Input['EncryptionArgs'] encryption: Encryption settings to be used for server-side encryption for proprietary content (such as catalogs, logs, customizations).
        :param pulumi.Input['ManagedServiceIdentityArgs'] identity: Managed identity properties
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input['DevCenterProjectCatalogSettingsArgs'] project_catalog_settings: Dev Center settings to be used when associating a project with a catalog.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if dev_center_name is not None:
            pulumi.set(__self__, "dev_center_name", dev_center_name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project_catalog_settings is not None:
            pulumi.set(__self__, "project_catalog_settings", project_catalog_settings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter(name="devCenterName")
    def dev_center_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the devcenter.
        """
        return pulumi.get(self, "dev_center_name")

    @dev_center_name.setter
    def dev_center_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "dev_center_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The display name of the devcenter.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input['EncryptionArgs']]:
        """
        Encryption settings to be used for server-side encryption for proprietary content (such as catalogs, logs, customizations).
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input['EncryptionArgs']]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ManagedServiceIdentityArgs']]:
        """
        Managed identity properties
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ManagedServiceIdentityArgs']]):
        pulumi.set(self, "identity", value)

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
    @pulumi.getter(name="projectCatalogSettings")
    def project_catalog_settings(self) -> Optional[pulumi.Input['DevCenterProjectCatalogSettingsArgs']]:
        """
        Dev Center settings to be used when associating a project with a catalog.
        """
        return pulumi.get(self, "project_catalog_settings")

    @project_catalog_settings.setter
    def project_catalog_settings(self, value: Optional[pulumi.Input['DevCenterProjectCatalogSettingsArgs']]):
        pulumi.set(self, "project_catalog_settings", value)

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


@pulumi.type_token("azure-native:devcenter:DevCenter")
class DevCenter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dev_center_name: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 encryption: Optional[pulumi.Input[Union['EncryptionArgs', 'EncryptionArgsDict']]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 project_catalog_settings: Optional[pulumi.Input[Union['DevCenterProjectCatalogSettingsArgs', 'DevCenterProjectCatalogSettingsArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Represents a devcenter resource.

        Uses Azure REST API version 2024-02-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.

        Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] dev_center_name: The name of the devcenter.
        :param pulumi.Input[builtins.str] display_name: The display name of the devcenter.
        :param pulumi.Input[Union['EncryptionArgs', 'EncryptionArgsDict']] encryption: Encryption settings to be used for server-side encryption for proprietary content (such as catalogs, logs, customizations).
        :param pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']] identity: Managed identity properties
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Union['DevCenterProjectCatalogSettingsArgs', 'DevCenterProjectCatalogSettingsArgsDict']] project_catalog_settings: Dev Center settings to be used when associating a project with a catalog.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DevCenterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a devcenter resource.

        Uses Azure REST API version 2024-02-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.

        Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DevCenterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DevCenterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dev_center_name: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 encryption: Optional[pulumi.Input[Union['EncryptionArgs', 'EncryptionArgsDict']]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 project_catalog_settings: Optional[pulumi.Input[Union['DevCenterProjectCatalogSettingsArgs', 'DevCenterProjectCatalogSettingsArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DevCenterArgs.__new__(DevCenterArgs)

            __props__.__dict__["dev_center_name"] = dev_center_name
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["encryption"] = encryption
            __props__.__dict__["identity"] = identity
            __props__.__dict__["location"] = location
            __props__.__dict__["project_catalog_settings"] = project_catalog_settings
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["dev_center_uri"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:devcenter/v20220801preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20220901preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20221012preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20221111preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20230101preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20230401:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20230801preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20231001preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20240201:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20240501preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20240601preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20240701preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20240801preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20241001preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20250201:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20250401preview:DevCenter"), pulumi.Alias(type_="azure-native:devcenter/v20250701preview:DevCenter")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DevCenter, __self__).__init__(
            'azure-native:devcenter:DevCenter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DevCenter':
        """
        Get an existing DevCenter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DevCenterArgs.__new__(DevCenterArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["dev_center_uri"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["encryption"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project_catalog_settings"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return DevCenter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="devCenterUri")
    def dev_center_uri(self) -> pulumi.Output[builtins.str]:
        """
        The URI of the Dev Center.
        """
        return pulumi.get(self, "dev_center_uri")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The display name of the devcenter.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[Optional['outputs.EncryptionResponse']]:
        """
        Encryption settings to be used for server-side encryption for proprietary content (such as catalogs, logs, customizations).
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ManagedServiceIdentityResponse']]:
        """
        Managed identity properties
        """
        return pulumi.get(self, "identity")

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
    @pulumi.getter(name="projectCatalogSettings")
    def project_catalog_settings(self) -> pulumi.Output[Optional['outputs.DevCenterProjectCatalogSettingsResponse']]:
        """
        Dev Center settings to be used when associating a project with a catalog.
        """
        return pulumi.get(self, "project_catalog_settings")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

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
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

