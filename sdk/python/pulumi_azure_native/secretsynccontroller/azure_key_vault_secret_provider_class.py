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

__all__ = ['AzureKeyVaultSecretProviderClassArgs', 'AzureKeyVaultSecretProviderClass']

@pulumi.input_type
class AzureKeyVaultSecretProviderClassArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[builtins.str],
                 keyvault_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 tenant_id: pulumi.Input[builtins.str],
                 azure_key_vault_secret_provider_class_name: Optional[pulumi.Input[builtins.str]] = None,
                 extended_location: Optional[pulumi.Input['AzureResourceManagerCommonTypesExtendedLocationArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 objects: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a AzureKeyVaultSecretProviderClass resource.
        :param pulumi.Input[builtins.str] client_id: The user assigned managed identity client ID that should be used to access the Azure Key Vault.
        :param pulumi.Input[builtins.str] keyvault_name: The name of the Azure Key Vault to sync secrets from.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] tenant_id: The Azure Active Directory tenant ID that should be used for authenticating requests to the Azure Key Vault.
        :param pulumi.Input[builtins.str] azure_key_vault_secret_provider_class_name: The name of the AzureKeyVaultSecretProviderClass
        :param pulumi.Input['AzureResourceManagerCommonTypesExtendedLocationArgs'] extended_location: The complex type of the extended location.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] objects: Objects defines the desired state of synced K8s secret objects
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "keyvault_name", keyvault_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if azure_key_vault_secret_provider_class_name is not None:
            pulumi.set(__self__, "azure_key_vault_secret_provider_class_name", azure_key_vault_secret_provider_class_name)
        if extended_location is not None:
            pulumi.set(__self__, "extended_location", extended_location)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if objects is not None:
            pulumi.set(__self__, "objects", objects)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[builtins.str]:
        """
        The user assigned managed identity client ID that should be used to access the Azure Key Vault.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="keyvaultName")
    def keyvault_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Azure Key Vault to sync secrets from.
        """
        return pulumi.get(self, "keyvault_name")

    @keyvault_name.setter
    def keyvault_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "keyvault_name", value)

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
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[builtins.str]:
        """
        The Azure Active Directory tenant ID that should be used for authenticating requests to the Azure Key Vault.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="azureKeyVaultSecretProviderClassName")
    def azure_key_vault_secret_provider_class_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the AzureKeyVaultSecretProviderClass
        """
        return pulumi.get(self, "azure_key_vault_secret_provider_class_name")

    @azure_key_vault_secret_provider_class_name.setter
    def azure_key_vault_secret_provider_class_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "azure_key_vault_secret_provider_class_name", value)

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> Optional[pulumi.Input['AzureResourceManagerCommonTypesExtendedLocationArgs']]:
        """
        The complex type of the extended location.
        """
        return pulumi.get(self, "extended_location")

    @extended_location.setter
    def extended_location(self, value: Optional[pulumi.Input['AzureResourceManagerCommonTypesExtendedLocationArgs']]):
        pulumi.set(self, "extended_location", value)

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
    @pulumi.getter
    def objects(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Objects defines the desired state of synced K8s secret objects
        """
        return pulumi.get(self, "objects")

    @objects.setter
    def objects(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "objects", value)

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


@pulumi.type_token("azure-native:secretsynccontroller:AzureKeyVaultSecretProviderClass")
class AzureKeyVaultSecretProviderClass(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_key_vault_secret_provider_class_name: Optional[pulumi.Input[builtins.str]] = None,
                 client_id: Optional[pulumi.Input[builtins.str]] = None,
                 extended_location: Optional[pulumi.Input[Union['AzureResourceManagerCommonTypesExtendedLocationArgs', 'AzureResourceManagerCommonTypesExtendedLocationArgsDict']]] = None,
                 keyvault_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 objects: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The AzureKeyVaultSecretProviderClass resource.

        Uses Azure REST API version 2024-08-21-preview. In version 2.x of the Azure Native provider, it used API version 2024-08-21-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] azure_key_vault_secret_provider_class_name: The name of the AzureKeyVaultSecretProviderClass
        :param pulumi.Input[builtins.str] client_id: The user assigned managed identity client ID that should be used to access the Azure Key Vault.
        :param pulumi.Input[Union['AzureResourceManagerCommonTypesExtendedLocationArgs', 'AzureResourceManagerCommonTypesExtendedLocationArgsDict']] extended_location: The complex type of the extended location.
        :param pulumi.Input[builtins.str] keyvault_name: The name of the Azure Key Vault to sync secrets from.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] objects: Objects defines the desired state of synced K8s secret objects
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] tenant_id: The Azure Active Directory tenant ID that should be used for authenticating requests to the Azure Key Vault.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AzureKeyVaultSecretProviderClassArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AzureKeyVaultSecretProviderClass resource.

        Uses Azure REST API version 2024-08-21-preview. In version 2.x of the Azure Native provider, it used API version 2024-08-21-preview.

        :param str resource_name: The name of the resource.
        :param AzureKeyVaultSecretProviderClassArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AzureKeyVaultSecretProviderClassArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_key_vault_secret_provider_class_name: Optional[pulumi.Input[builtins.str]] = None,
                 client_id: Optional[pulumi.Input[builtins.str]] = None,
                 extended_location: Optional[pulumi.Input[Union['AzureResourceManagerCommonTypesExtendedLocationArgs', 'AzureResourceManagerCommonTypesExtendedLocationArgsDict']]] = None,
                 keyvault_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 objects: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AzureKeyVaultSecretProviderClassArgs.__new__(AzureKeyVaultSecretProviderClassArgs)

            __props__.__dict__["azure_key_vault_secret_provider_class_name"] = azure_key_vault_secret_provider_class_name
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["extended_location"] = extended_location
            if keyvault_name is None and not opts.urn:
                raise TypeError("Missing required property 'keyvault_name'")
            __props__.__dict__["keyvault_name"] = keyvault_name
            __props__.__dict__["location"] = location
            __props__.__dict__["objects"] = objects
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:secretsynccontroller/v20240821preview:AzureKeyVaultSecretProviderClass")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AzureKeyVaultSecretProviderClass, __self__).__init__(
            'azure-native:secretsynccontroller:AzureKeyVaultSecretProviderClass',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AzureKeyVaultSecretProviderClass':
        """
        Get an existing AzureKeyVaultSecretProviderClass resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AzureKeyVaultSecretProviderClassArgs.__new__(AzureKeyVaultSecretProviderClassArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["client_id"] = None
        __props__.__dict__["extended_location"] = None
        __props__.__dict__["keyvault_name"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["objects"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tenant_id"] = None
        __props__.__dict__["type"] = None
        return AzureKeyVaultSecretProviderClass(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[builtins.str]:
        """
        The user assigned managed identity client ID that should be used to access the Azure Key Vault.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> pulumi.Output[Optional['outputs.AzureResourceManagerCommonTypesExtendedLocationResponse']]:
        """
        The complex type of the extended location.
        """
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter(name="keyvaultName")
    def keyvault_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the Azure Key Vault to sync secrets from.
        """
        return pulumi.get(self, "keyvault_name")

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
    @pulumi.getter
    def objects(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Objects defines the desired state of synced K8s secret objects
        """
        return pulumi.get(self, "objects")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of the AzureKeyVaultSecretProviderClass instance.
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
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[builtins.str]:
        """
        The Azure Active Directory tenant ID that should be used for authenticating requests to the Azure Key Vault.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

