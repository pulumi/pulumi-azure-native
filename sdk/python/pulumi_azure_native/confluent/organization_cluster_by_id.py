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

__all__ = ['OrganizationClusterByIdArgs', 'OrganizationClusterById']

@pulumi.input_type
class OrganizationClusterByIdArgs:
    def __init__(__self__, *,
                 environment_id: pulumi.Input[builtins.str],
                 organization_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 metadata: Optional[pulumi.Input['SCMetadataEntityArgs']] = None,
                 spec: Optional[pulumi.Input['SCClusterSpecEntityArgs']] = None,
                 status: Optional[pulumi.Input['ClusterStatusEntityArgs']] = None):
        """
        The set of arguments for constructing a OrganizationClusterById resource.
        :param pulumi.Input[builtins.str] environment_id: Confluent environment id
        :param pulumi.Input[builtins.str] organization_name: Organization resource name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] cluster_id: Confluent kafka or schema registry cluster id
        :param pulumi.Input[builtins.str] kind: Type of cluster
        :param pulumi.Input['SCMetadataEntityArgs'] metadata: Metadata of the record
        :param pulumi.Input['SCClusterSpecEntityArgs'] spec: Specification of the cluster
        :param pulumi.Input['ClusterStatusEntityArgs'] status: Specification of the cluster status
        """
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "organization_name", organization_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[builtins.str]:
        """
        Confluent environment id
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="organizationName")
    def organization_name(self) -> pulumi.Input[builtins.str]:
        """
        Organization resource name
        """
        return pulumi.get(self, "organization_name")

    @organization_name.setter
    def organization_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "organization_name", value)

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
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Confluent kafka or schema registry cluster id
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Type of cluster
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['SCMetadataEntityArgs']]:
        """
        Metadata of the record
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['SCMetadataEntityArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['SCClusterSpecEntityArgs']]:
        """
        Specification of the cluster
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['SCClusterSpecEntityArgs']]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['ClusterStatusEntityArgs']]:
        """
        Specification of the cluster status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['ClusterStatusEntityArgs']]):
        pulumi.set(self, "status", value)


@pulumi.type_token("azure-native:confluent:OrganizationClusterById")
class OrganizationClusterById(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 metadata: Optional[pulumi.Input[Union['SCMetadataEntityArgs', 'SCMetadataEntityArgsDict']]] = None,
                 organization_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 spec: Optional[pulumi.Input[Union['SCClusterSpecEntityArgs', 'SCClusterSpecEntityArgsDict']]] = None,
                 status: Optional[pulumi.Input[Union['ClusterStatusEntityArgs', 'ClusterStatusEntityArgsDict']]] = None,
                 __props__=None):
        """
        Details of cluster record

        Uses Azure REST API version 2024-07-01. In version 2.x of the Azure Native provider, it used API version 2024-07-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_id: Confluent kafka or schema registry cluster id
        :param pulumi.Input[builtins.str] environment_id: Confluent environment id
        :param pulumi.Input[builtins.str] kind: Type of cluster
        :param pulumi.Input[Union['SCMetadataEntityArgs', 'SCMetadataEntityArgsDict']] metadata: Metadata of the record
        :param pulumi.Input[builtins.str] organization_name: Organization resource name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['SCClusterSpecEntityArgs', 'SCClusterSpecEntityArgsDict']] spec: Specification of the cluster
        :param pulumi.Input[Union['ClusterStatusEntityArgs', 'ClusterStatusEntityArgsDict']] status: Specification of the cluster status
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationClusterByIdArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Details of cluster record

        Uses Azure REST API version 2024-07-01. In version 2.x of the Azure Native provider, it used API version 2024-07-01.

        :param str resource_name: The name of the resource.
        :param OrganizationClusterByIdArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationClusterByIdArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 metadata: Optional[pulumi.Input[Union['SCMetadataEntityArgs', 'SCMetadataEntityArgsDict']]] = None,
                 organization_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 spec: Optional[pulumi.Input[Union['SCClusterSpecEntityArgs', 'SCClusterSpecEntityArgsDict']]] = None,
                 status: Optional[pulumi.Input[Union['ClusterStatusEntityArgs', 'ClusterStatusEntityArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationClusterByIdArgs.__new__(OrganizationClusterByIdArgs)

            __props__.__dict__["cluster_id"] = cluster_id
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["kind"] = kind
            __props__.__dict__["metadata"] = metadata
            if organization_name is None and not opts.urn:
                raise TypeError("Missing required property 'organization_name'")
            __props__.__dict__["organization_name"] = organization_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["spec"] = spec
            __props__.__dict__["status"] = status
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:confluent/v20240701:OrganizationClusterById")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(OrganizationClusterById, __self__).__init__(
            'azure-native:confluent:OrganizationClusterById',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationClusterById':
        """
        Get an existing OrganizationClusterById resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationClusterByIdArgs.__new__(OrganizationClusterByIdArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["spec"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return OrganizationClusterById(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Type of cluster
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional['outputs.SCMetadataEntityResponse']]:
        """
        Metadata of the record
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[Optional['outputs.SCClusterSpecEntityResponse']]:
        """
        Specification of the cluster
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['outputs.ClusterStatusEntityResponse']]:
        """
        Specification of the cluster status
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
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

