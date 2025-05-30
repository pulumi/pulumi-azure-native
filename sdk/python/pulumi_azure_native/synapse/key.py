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

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 workspace_name: pulumi.Input[builtins.str],
                 is_active_cmk: Optional[pulumi.Input[builtins.bool]] = None,
                 key_name: Optional[pulumi.Input[builtins.str]] = None,
                 key_vault_url: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        :param pulumi.Input[builtins.bool] is_active_cmk: Used to activate the workspace after a customer managed key is provided.
        :param pulumi.Input[builtins.str] key_name: The name of the workspace key
        :param pulumi.Input[builtins.str] key_vault_url: The Key Vault Url of the workspace key.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if is_active_cmk is not None:
            pulumi.set(__self__, "is_active_cmk", is_active_cmk)
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_vault_url is not None:
            pulumi.set(__self__, "key_vault_url", key_vault_url)

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
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "workspace_name", value)

    @property
    @pulumi.getter(name="isActiveCMK")
    def is_active_cmk(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Used to activate the workspace after a customer managed key is provided.
        """
        return pulumi.get(self, "is_active_cmk")

    @is_active_cmk.setter
    def is_active_cmk(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_active_cmk", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the workspace key
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyVaultUrl")
    def key_vault_url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Key Vault Url of the workspace key.
        """
        return pulumi.get(self, "key_vault_url")

    @key_vault_url.setter
    def key_vault_url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_vault_url", value)


@pulumi.type_token("azure-native:synapse:Key")
class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_active_cmk: Optional[pulumi.Input[builtins.bool]] = None,
                 key_name: Optional[pulumi.Input[builtins.str]] = None,
                 key_vault_url: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A workspace key

        Uses Azure REST API version 2021-06-01. In version 2.x of the Azure Native provider, it used API version 2021-06-01.

        Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] is_active_cmk: Used to activate the workspace after a customer managed key is provided.
        :param pulumi.Input[builtins.str] key_name: The name of the workspace key
        :param pulumi.Input[builtins.str] key_vault_url: The Key Vault Url of the workspace key.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A workspace key

        Uses Azure REST API version 2021-06-01. In version 2.x of the Azure Native provider, it used API version 2021-06-01.

        Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_active_cmk: Optional[pulumi.Input[builtins.bool]] = None,
                 key_name: Optional[pulumi.Input[builtins.str]] = None,
                 key_vault_url: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeyArgs.__new__(KeyArgs)

            __props__.__dict__["is_active_cmk"] = is_active_cmk
            __props__.__dict__["key_name"] = key_name
            __props__.__dict__["key_vault_url"] = key_vault_url
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:synapse/v20190601preview:Key"), pulumi.Alias(type_="azure-native:synapse/v20201201:Key"), pulumi.Alias(type_="azure-native:synapse/v20210301:Key"), pulumi.Alias(type_="azure-native:synapse/v20210401preview:Key"), pulumi.Alias(type_="azure-native:synapse/v20210501:Key"), pulumi.Alias(type_="azure-native:synapse/v20210601:Key"), pulumi.Alias(type_="azure-native:synapse/v20210601preview:Key")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Key, __self__).__init__(
            'azure-native:synapse:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KeyArgs.__new__(KeyArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["is_active_cmk"] = None
        __props__.__dict__["key_vault_url"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["type"] = None
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="isActiveCMK")
    def is_active_cmk(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Used to activate the workspace after a customer managed key is provided.
        """
        return pulumi.get(self, "is_active_cmk")

    @property
    @pulumi.getter(name="keyVaultUrl")
    def key_vault_url(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Key Vault Url of the workspace key.
        """
        return pulumi.get(self, "key_vault_url")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

