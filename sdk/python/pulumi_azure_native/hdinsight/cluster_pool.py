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

__all__ = ['ClusterPoolArgs', 'ClusterPool']

@pulumi.input_type
class ClusterPoolArgs:
    def __init__(__self__, *,
                 compute_profile: pulumi.Input['ClusterPoolResourcePropertiesComputeProfileArgs'],
                 resource_group_name: pulumi.Input[builtins.str],
                 cluster_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                 cluster_pool_profile: Optional[pulumi.Input['ClusterPoolResourcePropertiesClusterPoolProfileArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_analytics_profile: Optional[pulumi.Input['ClusterPoolResourcePropertiesLogAnalyticsProfileArgs']] = None,
                 managed_resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 network_profile: Optional[pulumi.Input['ClusterPoolResourcePropertiesNetworkProfileArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a ClusterPool resource.
        :param pulumi.Input['ClusterPoolResourcePropertiesComputeProfileArgs'] compute_profile: CLuster pool compute profile.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] cluster_pool_name: The name of the cluster pool.
        :param pulumi.Input['ClusterPoolResourcePropertiesClusterPoolProfileArgs'] cluster_pool_profile: CLuster pool profile.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input['ClusterPoolResourcePropertiesLogAnalyticsProfileArgs'] log_analytics_profile: Cluster pool log analytics profile to enable OMS agent for AKS cluster.
        :param pulumi.Input[builtins.str] managed_resource_group_name: A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
        :param pulumi.Input['ClusterPoolResourcePropertiesNetworkProfileArgs'] network_profile: Cluster pool network profile.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "compute_profile", compute_profile)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if cluster_pool_name is not None:
            pulumi.set(__self__, "cluster_pool_name", cluster_pool_name)
        if cluster_pool_profile is not None:
            pulumi.set(__self__, "cluster_pool_profile", cluster_pool_profile)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log_analytics_profile is not None:
            pulumi.set(__self__, "log_analytics_profile", log_analytics_profile)
        if managed_resource_group_name is not None:
            pulumi.set(__self__, "managed_resource_group_name", managed_resource_group_name)
        if network_profile is not None:
            pulumi.set(__self__, "network_profile", network_profile)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="computeProfile")
    def compute_profile(self) -> pulumi.Input['ClusterPoolResourcePropertiesComputeProfileArgs']:
        """
        CLuster pool compute profile.
        """
        return pulumi.get(self, "compute_profile")

    @compute_profile.setter
    def compute_profile(self, value: pulumi.Input['ClusterPoolResourcePropertiesComputeProfileArgs']):
        pulumi.set(self, "compute_profile", value)

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
    @pulumi.getter(name="clusterPoolName")
    def cluster_pool_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the cluster pool.
        """
        return pulumi.get(self, "cluster_pool_name")

    @cluster_pool_name.setter
    def cluster_pool_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cluster_pool_name", value)

    @property
    @pulumi.getter(name="clusterPoolProfile")
    def cluster_pool_profile(self) -> Optional[pulumi.Input['ClusterPoolResourcePropertiesClusterPoolProfileArgs']]:
        """
        CLuster pool profile.
        """
        return pulumi.get(self, "cluster_pool_profile")

    @cluster_pool_profile.setter
    def cluster_pool_profile(self, value: Optional[pulumi.Input['ClusterPoolResourcePropertiesClusterPoolProfileArgs']]):
        pulumi.set(self, "cluster_pool_profile", value)

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
    @pulumi.getter(name="logAnalyticsProfile")
    def log_analytics_profile(self) -> Optional[pulumi.Input['ClusterPoolResourcePropertiesLogAnalyticsProfileArgs']]:
        """
        Cluster pool log analytics profile to enable OMS agent for AKS cluster.
        """
        return pulumi.get(self, "log_analytics_profile")

    @log_analytics_profile.setter
    def log_analytics_profile(self, value: Optional[pulumi.Input['ClusterPoolResourcePropertiesLogAnalyticsProfileArgs']]):
        pulumi.set(self, "log_analytics_profile", value)

    @property
    @pulumi.getter(name="managedResourceGroupName")
    def managed_resource_group_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
        """
        return pulumi.get(self, "managed_resource_group_name")

    @managed_resource_group_name.setter
    def managed_resource_group_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "managed_resource_group_name", value)

    @property
    @pulumi.getter(name="networkProfile")
    def network_profile(self) -> Optional[pulumi.Input['ClusterPoolResourcePropertiesNetworkProfileArgs']]:
        """
        Cluster pool network profile.
        """
        return pulumi.get(self, "network_profile")

    @network_profile.setter
    def network_profile(self, value: Optional[pulumi.Input['ClusterPoolResourcePropertiesNetworkProfileArgs']]):
        pulumi.set(self, "network_profile", value)

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


@pulumi.type_token("azure-native:hdinsight:ClusterPool")
class ClusterPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                 cluster_pool_profile: Optional[pulumi.Input[Union['ClusterPoolResourcePropertiesClusterPoolProfileArgs', 'ClusterPoolResourcePropertiesClusterPoolProfileArgsDict']]] = None,
                 compute_profile: Optional[pulumi.Input[Union['ClusterPoolResourcePropertiesComputeProfileArgs', 'ClusterPoolResourcePropertiesComputeProfileArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_analytics_profile: Optional[pulumi.Input[Union['ClusterPoolResourcePropertiesLogAnalyticsProfileArgs', 'ClusterPoolResourcePropertiesLogAnalyticsProfileArgsDict']]] = None,
                 managed_resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 network_profile: Optional[pulumi.Input[Union['ClusterPoolResourcePropertiesNetworkProfileArgs', 'ClusterPoolResourcePropertiesNetworkProfileArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Cluster pool.

        Uses Azure REST API version 2024-05-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.

        Other available API versions: 2023-06-01-preview, 2023-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hdinsight [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_pool_name: The name of the cluster pool.
        :param pulumi.Input[Union['ClusterPoolResourcePropertiesClusterPoolProfileArgs', 'ClusterPoolResourcePropertiesClusterPoolProfileArgsDict']] cluster_pool_profile: CLuster pool profile.
        :param pulumi.Input[Union['ClusterPoolResourcePropertiesComputeProfileArgs', 'ClusterPoolResourcePropertiesComputeProfileArgsDict']] compute_profile: CLuster pool compute profile.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Union['ClusterPoolResourcePropertiesLogAnalyticsProfileArgs', 'ClusterPoolResourcePropertiesLogAnalyticsProfileArgsDict']] log_analytics_profile: Cluster pool log analytics profile to enable OMS agent for AKS cluster.
        :param pulumi.Input[builtins.str] managed_resource_group_name: A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
        :param pulumi.Input[Union['ClusterPoolResourcePropertiesNetworkProfileArgs', 'ClusterPoolResourcePropertiesNetworkProfileArgsDict']] network_profile: Cluster pool network profile.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Cluster pool.

        Uses Azure REST API version 2024-05-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.

        Other available API versions: 2023-06-01-preview, 2023-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hdinsight [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ClusterPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                 cluster_pool_profile: Optional[pulumi.Input[Union['ClusterPoolResourcePropertiesClusterPoolProfileArgs', 'ClusterPoolResourcePropertiesClusterPoolProfileArgsDict']]] = None,
                 compute_profile: Optional[pulumi.Input[Union['ClusterPoolResourcePropertiesComputeProfileArgs', 'ClusterPoolResourcePropertiesComputeProfileArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_analytics_profile: Optional[pulumi.Input[Union['ClusterPoolResourcePropertiesLogAnalyticsProfileArgs', 'ClusterPoolResourcePropertiesLogAnalyticsProfileArgsDict']]] = None,
                 managed_resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 network_profile: Optional[pulumi.Input[Union['ClusterPoolResourcePropertiesNetworkProfileArgs', 'ClusterPoolResourcePropertiesNetworkProfileArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterPoolArgs.__new__(ClusterPoolArgs)

            __props__.__dict__["cluster_pool_name"] = cluster_pool_name
            __props__.__dict__["cluster_pool_profile"] = cluster_pool_profile
            if compute_profile is None and not opts.urn:
                raise TypeError("Missing required property 'compute_profile'")
            __props__.__dict__["compute_profile"] = compute_profile
            __props__.__dict__["location"] = location
            __props__.__dict__["log_analytics_profile"] = log_analytics_profile
            __props__.__dict__["managed_resource_group_name"] = managed_resource_group_name
            __props__.__dict__["network_profile"] = network_profile
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["aks_cluster_profile"] = None
            __props__.__dict__["aks_managed_resource_group_name"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["deployment_id"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:hdinsight/v20230601preview:ClusterPool"), pulumi.Alias(type_="azure-native:hdinsight/v20231101preview:ClusterPool"), pulumi.Alias(type_="azure-native:hdinsight/v20240501preview:ClusterPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ClusterPool, __self__).__init__(
            'azure-native:hdinsight:ClusterPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ClusterPool':
        """
        Get an existing ClusterPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ClusterPoolArgs.__new__(ClusterPoolArgs)

        __props__.__dict__["aks_cluster_profile"] = None
        __props__.__dict__["aks_managed_resource_group_name"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["cluster_pool_profile"] = None
        __props__.__dict__["compute_profile"] = None
        __props__.__dict__["deployment_id"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["log_analytics_profile"] = None
        __props__.__dict__["managed_resource_group_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_profile"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return ClusterPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aksClusterProfile")
    def aks_cluster_profile(self) -> pulumi.Output['outputs.ClusterPoolResourcePropertiesResponseAksClusterProfile']:
        """
        Properties of underlying AKS cluster.
        """
        return pulumi.get(self, "aks_cluster_profile")

    @property
    @pulumi.getter(name="aksManagedResourceGroupName")
    def aks_managed_resource_group_name(self) -> pulumi.Output[builtins.str]:
        """
        A resource group created by AKS, to hold the infrastructure resources created by AKS on-behalf of customers. It is generated by cluster pool name and managed resource group name by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}
        """
        return pulumi.get(self, "aks_managed_resource_group_name")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="clusterPoolProfile")
    def cluster_pool_profile(self) -> pulumi.Output[Optional['outputs.ClusterPoolResourcePropertiesResponseClusterPoolProfile']]:
        """
        CLuster pool profile.
        """
        return pulumi.get(self, "cluster_pool_profile")

    @property
    @pulumi.getter(name="computeProfile")
    def compute_profile(self) -> pulumi.Output['outputs.ClusterPoolResourcePropertiesResponseComputeProfile']:
        """
        CLuster pool compute profile.
        """
        return pulumi.get(self, "compute_profile")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Output[builtins.str]:
        """
        A unique id generated by the RP to identify the resource.
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logAnalyticsProfile")
    def log_analytics_profile(self) -> pulumi.Output[Optional['outputs.ClusterPoolResourcePropertiesResponseLogAnalyticsProfile']]:
        """
        Cluster pool log analytics profile to enable OMS agent for AKS cluster.
        """
        return pulumi.get(self, "log_analytics_profile")

    @property
    @pulumi.getter(name="managedResourceGroupName")
    def managed_resource_group_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
        """
        return pulumi.get(self, "managed_resource_group_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkProfile")
    def network_profile(self) -> pulumi.Output[Optional['outputs.ClusterPoolResourcePropertiesResponseNetworkProfile']]:
        """
        Cluster pool network profile.
        """
        return pulumi.get(self, "network_profile")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Business status of the resource.
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

