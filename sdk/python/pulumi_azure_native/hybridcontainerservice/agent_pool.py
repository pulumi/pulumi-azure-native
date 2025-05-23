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

__all__ = ['AgentPoolArgs', 'AgentPool']

@pulumi.input_type
class AgentPoolArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 resource_name: pulumi.Input[builtins.str],
                 agent_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 cloud_provider_profile: Optional[pulumi.Input['CloudProviderProfileArgs']] = None,
                 count: Optional[pulumi.Input[builtins.int]] = None,
                 extended_location: Optional[pulumi.Input['AgentPoolExtendedLocationArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 max_count: Optional[pulumi.Input[builtins.int]] = None,
                 max_pods: Optional[pulumi.Input[builtins.int]] = None,
                 min_count: Optional[pulumi.Input[builtins.int]] = None,
                 mode: Optional[pulumi.Input[Union[builtins.str, 'Mode']]] = None,
                 node_image_version: Optional[pulumi.Input[builtins.str]] = None,
                 node_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 node_taints: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 os_type: Optional[pulumi.Input[Union[builtins.str, 'OsType']]] = None,
                 status: Optional[pulumi.Input['AgentPoolProvisioningStatusStatusArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 vm_size: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a AgentPool resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_name: Parameter for the name of the provisioned cluster
        :param pulumi.Input[builtins.str] agent_pool_name: Parameter for the name of the agent pool in the provisioned cluster
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] availability_zones: AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
        :param pulumi.Input['CloudProviderProfileArgs'] cloud_provider_profile: The underlying cloud infra provider properties.
        :param pulumi.Input[builtins.int] count: Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        :param pulumi.Input[builtins.str] location: The resource location
        :param pulumi.Input[builtins.int] max_count: The maximum number of nodes for auto-scaling
        :param pulumi.Input[builtins.int] max_pods: The maximum number of pods that can run on a node.
        :param pulumi.Input[builtins.int] min_count: The minimum number of nodes for auto-scaling
        :param pulumi.Input[Union[builtins.str, 'Mode']] mode: Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
        :param pulumi.Input[builtins.str] node_image_version: The version of node image
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] node_labels: NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] node_taints: NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        :param pulumi.Input[Union[builtins.str, 'OsType']] os_type: OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
        :param pulumi.Input['AgentPoolProvisioningStatusStatusArgs'] status: HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags
        :param pulumi.Input[builtins.str] vm_size: VmSize - The size of the agent pool VMs.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "resource_name", resource_name)
        if agent_pool_name is not None:
            pulumi.set(__self__, "agent_pool_name", agent_pool_name)
        if availability_zones is not None:
            pulumi.set(__self__, "availability_zones", availability_zones)
        if cloud_provider_profile is not None:
            pulumi.set(__self__, "cloud_provider_profile", cloud_provider_profile)
        if count is None:
            count = 1
        if count is not None:
            pulumi.set(__self__, "count", count)
        if extended_location is not None:
            pulumi.set(__self__, "extended_location", extended_location)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if max_count is not None:
            pulumi.set(__self__, "max_count", max_count)
        if max_pods is not None:
            pulumi.set(__self__, "max_pods", max_pods)
        if min_count is not None:
            pulumi.set(__self__, "min_count", min_count)
        if mode is None:
            mode = 'User'
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if node_image_version is not None:
            pulumi.set(__self__, "node_image_version", node_image_version)
        if node_labels is not None:
            pulumi.set(__self__, "node_labels", node_labels)
        if node_taints is not None:
            pulumi.set(__self__, "node_taints", node_taints)
        if os_type is not None:
            pulumi.set(__self__, "os_type", os_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vm_size is not None:
            pulumi.set(__self__, "vm_size", vm_size)

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
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Input[builtins.str]:
        """
        Parameter for the name of the provisioned cluster
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter(name="agentPoolName")
    def agent_pool_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Parameter for the name of the agent pool in the provisioned cluster
        """
        return pulumi.get(self, "agent_pool_name")

    @agent_pool_name.setter
    def agent_pool_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "agent_pool_name", value)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
        """
        return pulumi.get(self, "availability_zones")

    @availability_zones.setter
    def availability_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "availability_zones", value)

    @property
    @pulumi.getter(name="cloudProviderProfile")
    def cloud_provider_profile(self) -> Optional[pulumi.Input['CloudProviderProfileArgs']]:
        """
        The underlying cloud infra provider properties.
        """
        return pulumi.get(self, "cloud_provider_profile")

    @cloud_provider_profile.setter
    def cloud_provider_profile(self, value: Optional[pulumi.Input['CloudProviderProfileArgs']]):
        pulumi.set(self, "cloud_provider_profile", value)

    @property
    @pulumi.getter
    def count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> Optional[pulumi.Input['AgentPoolExtendedLocationArgs']]:
        return pulumi.get(self, "extended_location")

    @extended_location.setter
    def extended_location(self, value: Optional[pulumi.Input['AgentPoolExtendedLocationArgs']]):
        pulumi.set(self, "extended_location", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource location
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of nodes for auto-scaling
        """
        return pulumi.get(self, "max_count")

    @max_count.setter
    def max_count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_count", value)

    @property
    @pulumi.getter(name="maxPods")
    def max_pods(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum number of pods that can run on a node.
        """
        return pulumi.get(self, "max_pods")

    @max_pods.setter
    def max_pods(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_pods", value)

    @property
    @pulumi.getter(name="minCount")
    def min_count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The minimum number of nodes for auto-scaling
        """
        return pulumi.get(self, "min_count")

    @min_count.setter
    def min_count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "min_count", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[Union[builtins.str, 'Mode']]]:
        """
        Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[Union[builtins.str, 'Mode']]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="nodeImageVersion")
    def node_image_version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The version of node image
        """
        return pulumi.get(self, "node_image_version")

    @node_image_version.setter
    def node_image_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "node_image_version", value)

    @property
    @pulumi.getter(name="nodeLabels")
    def node_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
        """
        return pulumi.get(self, "node_labels")

    @node_labels.setter
    def node_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "node_labels", value)

    @property
    @pulumi.getter(name="nodeTaints")
    def node_taints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        """
        return pulumi.get(self, "node_taints")

    @node_taints.setter
    def node_taints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "node_taints", value)

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[pulumi.Input[Union[builtins.str, 'OsType']]]:
        """
        OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
        """
        return pulumi.get(self, "os_type")

    @os_type.setter
    def os_type(self, value: Optional[pulumi.Input[Union[builtins.str, 'OsType']]]):
        pulumi.set(self, "os_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['AgentPoolProvisioningStatusStatusArgs']]:
        """
        HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['AgentPoolProvisioningStatusStatusArgs']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        VmSize - The size of the agent pool VMs.
        """
        return pulumi.get(self, "vm_size")

    @vm_size.setter
    def vm_size(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vm_size", value)


@pulumi.type_token("azure-native:hybridcontainerservice:AgentPool")
class AgentPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 cloud_provider_profile: Optional[pulumi.Input[Union['CloudProviderProfileArgs', 'CloudProviderProfileArgsDict']]] = None,
                 count: Optional[pulumi.Input[builtins.int]] = None,
                 extended_location: Optional[pulumi.Input[Union['AgentPoolExtendedLocationArgs', 'AgentPoolExtendedLocationArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 max_count: Optional[pulumi.Input[builtins.int]] = None,
                 max_pods: Optional[pulumi.Input[builtins.int]] = None,
                 min_count: Optional[pulumi.Input[builtins.int]] = None,
                 mode: Optional[pulumi.Input[Union[builtins.str, 'Mode']]] = None,
                 node_image_version: Optional[pulumi.Input[builtins.str]] = None,
                 node_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 node_taints: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 os_type: Optional[pulumi.Input[Union[builtins.str, 'OsType']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[Union['AgentPoolProvisioningStatusStatusArgs', 'AgentPoolProvisioningStatusStatusArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 vm_size: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The agentPool resource definition

        Uses Azure REST API version 2022-09-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] agent_pool_name: Parameter for the name of the agent pool in the provisioned cluster
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] availability_zones: AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
        :param pulumi.Input[Union['CloudProviderProfileArgs', 'CloudProviderProfileArgsDict']] cloud_provider_profile: The underlying cloud infra provider properties.
        :param pulumi.Input[builtins.int] count: Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        :param pulumi.Input[builtins.str] location: The resource location
        :param pulumi.Input[builtins.int] max_count: The maximum number of nodes for auto-scaling
        :param pulumi.Input[builtins.int] max_pods: The maximum number of pods that can run on a node.
        :param pulumi.Input[builtins.int] min_count: The minimum number of nodes for auto-scaling
        :param pulumi.Input[Union[builtins.str, 'Mode']] mode: Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
        :param pulumi.Input[builtins.str] node_image_version: The version of node image
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] node_labels: NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] node_taints: NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        :param pulumi.Input[Union[builtins.str, 'OsType']] os_type: OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_name_: Parameter for the name of the provisioned cluster
        :param pulumi.Input[Union['AgentPoolProvisioningStatusStatusArgs', 'AgentPoolProvisioningStatusStatusArgsDict']] status: HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags
        :param pulumi.Input[builtins.str] vm_size: VmSize - The size of the agent pool VMs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgentPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The agentPool resource definition

        Uses Azure REST API version 2022-09-01-preview.

        :param str resource_name: The name of the resource.
        :param AgentPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 cloud_provider_profile: Optional[pulumi.Input[Union['CloudProviderProfileArgs', 'CloudProviderProfileArgsDict']]] = None,
                 count: Optional[pulumi.Input[builtins.int]] = None,
                 extended_location: Optional[pulumi.Input[Union['AgentPoolExtendedLocationArgs', 'AgentPoolExtendedLocationArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 max_count: Optional[pulumi.Input[builtins.int]] = None,
                 max_pods: Optional[pulumi.Input[builtins.int]] = None,
                 min_count: Optional[pulumi.Input[builtins.int]] = None,
                 mode: Optional[pulumi.Input[Union[builtins.str, 'Mode']]] = None,
                 node_image_version: Optional[pulumi.Input[builtins.str]] = None,
                 node_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 node_taints: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 os_type: Optional[pulumi.Input[Union[builtins.str, 'OsType']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[Union['AgentPoolProvisioningStatusStatusArgs', 'AgentPoolProvisioningStatusStatusArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 vm_size: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AgentPoolArgs.__new__(AgentPoolArgs)

            __props__.__dict__["agent_pool_name"] = agent_pool_name
            __props__.__dict__["availability_zones"] = availability_zones
            __props__.__dict__["cloud_provider_profile"] = cloud_provider_profile
            if count is None:
                count = 1
            __props__.__dict__["count"] = count
            __props__.__dict__["extended_location"] = extended_location
            __props__.__dict__["location"] = location
            __props__.__dict__["max_count"] = max_count
            __props__.__dict__["max_pods"] = max_pods
            __props__.__dict__["min_count"] = min_count
            if mode is None:
                mode = 'User'
            __props__.__dict__["mode"] = mode
            __props__.__dict__["node_image_version"] = node_image_version
            __props__.__dict__["node_labels"] = node_labels
            __props__.__dict__["node_taints"] = node_taints
            __props__.__dict__["os_type"] = os_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__.__dict__["resource_name"] = resource_name_
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vm_size"] = vm_size
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:hybridcontainerservice/v20220501preview:AgentPool"), pulumi.Alias(type_="azure-native:hybridcontainerservice/v20220901preview:AgentPool"), pulumi.Alias(type_="azure-native:hybridcontainerservice/v20231115preview:AgentPool"), pulumi.Alias(type_="azure-native:hybridcontainerservice/v20240101:AgentPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AgentPool, __self__).__init__(
            'azure-native:hybridcontainerservice:AgentPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentPool':
        """
        Get an existing AgentPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentPoolArgs.__new__(AgentPoolArgs)

        __props__.__dict__["availability_zones"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["cloud_provider_profile"] = None
        __props__.__dict__["count"] = None
        __props__.__dict__["extended_location"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["max_count"] = None
        __props__.__dict__["max_pods"] = None
        __props__.__dict__["min_count"] = None
        __props__.__dict__["mode"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["node_image_version"] = None
        __props__.__dict__["node_labels"] = None
        __props__.__dict__["node_taints"] = None
        __props__.__dict__["os_type"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["vm_size"] = None
        return AgentPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="cloudProviderProfile")
    def cloud_provider_profile(self) -> pulumi.Output[Optional['outputs.CloudProviderProfileResponse']]:
        """
        The underlying cloud infra provider properties.
        """
        return pulumi.get(self, "cloud_provider_profile")

    @property
    @pulumi.getter
    def count(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> pulumi.Output[Optional['outputs.AgentPoolResponseExtendedLocation']]:
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The maximum number of nodes for auto-scaling
        """
        return pulumi.get(self, "max_count")

    @property
    @pulumi.getter(name="maxPods")
    def max_pods(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The maximum number of pods that can run on a node.
        """
        return pulumi.get(self, "max_pods")

    @property
    @pulumi.getter(name="minCount")
    def min_count(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The minimum number of nodes for auto-scaling
        """
        return pulumi.get(self, "min_count")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeImageVersion")
    def node_image_version(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The version of node image
        """
        return pulumi.get(self, "node_image_version")

    @property
    @pulumi.getter(name="nodeLabels")
    def node_labels(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
        """
        return pulumi.get(self, "node_labels")

    @property
    @pulumi.getter(name="nodeTaints")
    def node_taints(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        """
        return pulumi.get(self, "node_taints")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['outputs.AgentPoolProvisioningStatusResponseStatus']]:
        """
        HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource Type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        VmSize - The size of the agent pool VMs.
        """
        return pulumi.get(self, "vm_size")

