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

__all__ = ['MongoClusterArgs', 'MongoCluster']

@pulumi.input_type
class MongoClusterArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 mongo_cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input['MongoClusterPropertiesArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a MongoCluster resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] mongo_cluster_name: The name of the mongo cluster.
        :param pulumi.Input['MongoClusterPropertiesArgs'] properties: The resource-specific properties for this resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if mongo_cluster_name is not None:
            pulumi.set(__self__, "mongo_cluster_name", mongo_cluster_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
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
    @pulumi.getter(name="mongoClusterName")
    def mongo_cluster_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the mongo cluster.
        """
        return pulumi.get(self, "mongo_cluster_name")

    @mongo_cluster_name.setter
    def mongo_cluster_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "mongo_cluster_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['MongoClusterPropertiesArgs']]:
        """
        The resource-specific properties for this resource.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['MongoClusterPropertiesArgs']]):
        pulumi.set(self, "properties", value)

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


@pulumi.type_token("azure-native:mongocluster:MongoCluster")
class MongoCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 mongo_cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['MongoClusterPropertiesArgs', 'MongoClusterPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Represents a mongo cluster resource.

        Uses Azure REST API version 2024-07-01.

        Other available API versions: 2024-03-01-preview, 2024-06-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mongocluster [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] mongo_cluster_name: The name of the mongo cluster.
        :param pulumi.Input[Union['MongoClusterPropertiesArgs', 'MongoClusterPropertiesArgsDict']] properties: The resource-specific properties for this resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MongoClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a mongo cluster resource.

        Uses Azure REST API version 2024-07-01.

        Other available API versions: 2024-03-01-preview, 2024-06-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mongocluster [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param MongoClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MongoClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 mongo_cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['MongoClusterPropertiesArgs', 'MongoClusterPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MongoClusterArgs.__new__(MongoClusterArgs)

            __props__.__dict__["location"] = location
            __props__.__dict__["mongo_cluster_name"] = mongo_cluster_name
            __props__.__dict__["properties"] = properties
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:documentdb/v20230315preview:MongoCluster"), pulumi.Alias(type_="azure-native:documentdb/v20230915preview:MongoCluster"), pulumi.Alias(type_="azure-native:documentdb/v20231115preview:MongoCluster"), pulumi.Alias(type_="azure-native:documentdb/v20240215preview:MongoCluster"), pulumi.Alias(type_="azure-native:documentdb/v20240301preview:MongoCluster"), pulumi.Alias(type_="azure-native:documentdb/v20240601preview:MongoCluster"), pulumi.Alias(type_="azure-native:documentdb/v20240701:MongoCluster"), pulumi.Alias(type_="azure-native:documentdb/v20241001preview:MongoCluster"), pulumi.Alias(type_="azure-native:documentdb:MongoCluster"), pulumi.Alias(type_="azure-native:mongocluster/v20240301preview:MongoCluster"), pulumi.Alias(type_="azure-native:mongocluster/v20240601preview:MongoCluster"), pulumi.Alias(type_="azure-native:mongocluster/v20240701:MongoCluster"), pulumi.Alias(type_="azure-native:mongocluster/v20241001preview:MongoCluster"), pulumi.Alias(type_="azure-native:mongocluster/v20250401preview:MongoCluster")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(MongoCluster, __self__).__init__(
            'azure-native:mongocluster:MongoCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MongoCluster':
        """
        Get an existing MongoCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MongoClusterArgs.__new__(MongoClusterArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return MongoCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

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
    def properties(self) -> pulumi.Output['outputs.MongoClusterPropertiesResponse']:
        """
        The resource-specific properties for this resource.
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

