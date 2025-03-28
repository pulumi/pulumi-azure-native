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
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 cluster_size: pulumi.Input[int],
                 private_cloud_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 sku: pulumi.Input['SkuArgs'],
                 cluster_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[int] cluster_size: The cluster size
        :param pulumi.Input[str] private_cloud_name: The name of the private cloud.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input['SkuArgs'] sku: The cluster SKU
        :param pulumi.Input[str] cluster_name: Name of the cluster in the private cloud
        """
        pulumi.set(__self__, "cluster_size", cluster_size)
        pulumi.set(__self__, "private_cloud_name", private_cloud_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku", sku)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)

    @property
    @pulumi.getter(name="clusterSize")
    def cluster_size(self) -> pulumi.Input[int]:
        """
        The cluster size
        """
        return pulumi.get(self, "cluster_size")

    @cluster_size.setter
    def cluster_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "cluster_size", value)

    @property
    @pulumi.getter(name="privateCloudName")
    def private_cloud_name(self) -> pulumi.Input[str]:
        """
        The name of the private cloud.
        """
        return pulumi.get(self, "private_cloud_name")

    @private_cloud_name.setter
    def private_cloud_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_cloud_name", value)

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
    @pulumi.getter
    def sku(self) -> pulumi.Input['SkuArgs']:
        """
        The cluster SKU
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input['SkuArgs']):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the cluster in the private cloud
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 cluster_size: Optional[pulumi.Input[int]] = None,
                 private_cloud_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 __props__=None):
        """
        A cluster resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Name of the cluster in the private cloud
        :param pulumi.Input[int] cluster_size: The cluster size
        :param pulumi.Input[str] private_cloud_name: The name of the private cloud.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['SkuArgs', 'SkuArgsDict']] sku: The cluster SKU
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A cluster resource

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 cluster_size: Optional[pulumi.Input[int]] = None,
                 private_cloud_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            __props__.__dict__["cluster_name"] = cluster_name
            if cluster_size is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_size'")
            __props__.__dict__["cluster_size"] = cluster_size
            if private_cloud_name is None and not opts.urn:
                raise TypeError("Missing required property 'private_cloud_name'")
            __props__.__dict__["private_cloud_name"] = private_cloud_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__.__dict__["sku"] = sku
            __props__.__dict__["cluster_id"] = None
            __props__.__dict__["hosts"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:avs/v20200717preview:Cluster"), pulumi.Alias(type_="azure-native:avs/v20210101preview:Cluster"), pulumi.Alias(type_="azure-native:avs/v20210601:Cluster"), pulumi.Alias(type_="azure-native:avs/v20211201:Cluster"), pulumi.Alias(type_="azure-native:avs/v20220501:Cluster"), pulumi.Alias(type_="azure-native:avs/v20230301:Cluster"), pulumi.Alias(type_="azure-native:avs/v20230901:Cluster"), pulumi.Alias(type_="azure-native:avs:Cluster")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Cluster, __self__).__init__(
            'azure-native:avs/v20200320:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ClusterArgs.__new__(ClusterArgs)

        __props__.__dict__["cluster_id"] = None
        __props__.__dict__["cluster_size"] = None
        __props__.__dict__["hosts"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["type"] = None
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[int]:
        """
        The identity
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterSize")
    def cluster_size(self) -> pulumi.Output[int]:
        """
        The cluster size
        """
        return pulumi.get(self, "cluster_size")

    @property
    @pulumi.getter
    def hosts(self) -> pulumi.Output[Sequence[str]]:
        """
        The hosts
        """
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The state of the cluster provisioning
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.SkuResponse']:
        """
        The cluster SKU
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

