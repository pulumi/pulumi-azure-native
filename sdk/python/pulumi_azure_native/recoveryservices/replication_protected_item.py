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

__all__ = ['ReplicationProtectedItemArgs', 'ReplicationProtectedItem']

@pulumi.input_type
class ReplicationProtectedItemArgs:
    def __init__(__self__, *,
                 fabric_name: pulumi.Input[builtins.str],
                 protection_container_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 resource_name: pulumi.Input[builtins.str],
                 properties: Optional[pulumi.Input['EnableProtectionInputPropertiesArgs']] = None,
                 replicated_protected_item_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ReplicationProtectedItem resource.
        :param pulumi.Input[builtins.str] fabric_name: Name of the fabric.
        :param pulumi.Input[builtins.str] protection_container_name: Protection container name.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[builtins.str] resource_name: The name of the recovery services vault.
        :param pulumi.Input['EnableProtectionInputPropertiesArgs'] properties: Enable protection input properties.
        :param pulumi.Input[builtins.str] replicated_protected_item_name: A name for the replication protected item.
        """
        pulumi.set(__self__, "fabric_name", fabric_name)
        pulumi.set(__self__, "protection_container_name", protection_container_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "resource_name", resource_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if replicated_protected_item_name is not None:
            pulumi.set(__self__, "replicated_protected_item_name", replicated_protected_item_name)

    @property
    @pulumi.getter(name="fabricName")
    def fabric_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the fabric.
        """
        return pulumi.get(self, "fabric_name")

    @fabric_name.setter
    def fabric_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "fabric_name", value)

    @property
    @pulumi.getter(name="protectionContainerName")
    def protection_container_name(self) -> pulumi.Input[builtins.str]:
        """
        Protection container name.
        """
        return pulumi.get(self, "protection_container_name")

    @protection_container_name.setter
    def protection_container_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "protection_container_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group where the recovery services vault is present.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the recovery services vault.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['EnableProtectionInputPropertiesArgs']]:
        """
        Enable protection input properties.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['EnableProtectionInputPropertiesArgs']]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter(name="replicatedProtectedItemName")
    def replicated_protected_item_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A name for the replication protected item.
        """
        return pulumi.get(self, "replicated_protected_item_name")

    @replicated_protected_item_name.setter
    def replicated_protected_item_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "replicated_protected_item_name", value)


@pulumi.type_token("azure-native:recoveryservices:ReplicationProtectedItem")
class ReplicationProtectedItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fabric_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['EnableProtectionInputPropertiesArgs', 'EnableProtectionInputPropertiesArgsDict']]] = None,
                 protection_container_name: Optional[pulumi.Input[builtins.str]] = None,
                 replicated_protected_item_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Replication protected item.

        Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.

        Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] fabric_name: Name of the fabric.
        :param pulumi.Input[Union['EnableProtectionInputPropertiesArgs', 'EnableProtectionInputPropertiesArgsDict']] properties: Enable protection input properties.
        :param pulumi.Input[builtins.str] protection_container_name: Protection container name.
        :param pulumi.Input[builtins.str] replicated_protected_item_name: A name for the replication protected item.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[builtins.str] resource_name_: The name of the recovery services vault.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicationProtectedItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Replication protected item.

        Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.

        Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ReplicationProtectedItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicationProtectedItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fabric_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['EnableProtectionInputPropertiesArgs', 'EnableProtectionInputPropertiesArgsDict']]] = None,
                 protection_container_name: Optional[pulumi.Input[builtins.str]] = None,
                 replicated_protected_item_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReplicationProtectedItemArgs.__new__(ReplicationProtectedItemArgs)

            if fabric_name is None and not opts.urn:
                raise TypeError("Missing required property 'fabric_name'")
            __props__.__dict__["fabric_name"] = fabric_name
            __props__.__dict__["properties"] = properties
            if protection_container_name is None and not opts.urn:
                raise TypeError("Missing required property 'protection_container_name'")
            __props__.__dict__["protection_container_name"] = protection_container_name
            __props__.__dict__["replicated_protected_item_name"] = replicated_protected_item_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__.__dict__["resource_name"] = resource_name_
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["location"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:recoveryservices/v20160810:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20180110:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20180710:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20210210:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20210301:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20210401:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20210601:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20210701:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20210801:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20211001:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20211101:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20211201:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20220101:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20220201:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20220301:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20220401:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20220501:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20220801:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20220910:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20221001:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20230101:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20230201:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20230401:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20230601:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20230801:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20240101:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20240201:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20240401:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20241001:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20250101:ReplicationProtectedItem"), pulumi.Alias(type_="azure-native:recoveryservices/v20250201:ReplicationProtectedItem")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ReplicationProtectedItem, __self__).__init__(
            'azure-native:recoveryservices:ReplicationProtectedItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ReplicationProtectedItem':
        """
        Get an existing ReplicationProtectedItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReplicationProtectedItemArgs.__new__(ReplicationProtectedItemArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["type"] = None
        return ReplicationProtectedItem(resource_name, opts=opts, __props__=__props__)

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
        Resource Location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.ReplicationProtectedItemPropertiesResponse']:
        """
        The custom data.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource Type
        """
        return pulumi.get(self, "type")

