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
from ._enums import *

__all__ = [
    'AddressPrefixItemArgs',
    'AddressPrefixItemArgsDict',
    'IPTrafficArgs',
    'IPTrafficArgsDict',
    'IpamPoolPropertiesArgs',
    'IpamPoolPropertiesArgsDict',
    'NetworkManagerPropertiesNetworkManagerScopesArgs',
    'NetworkManagerPropertiesNetworkManagerScopesArgsDict',
    'NetworkManagerSecurityGroupItemArgs',
    'NetworkManagerSecurityGroupItemArgsDict',
    'ReachabilityAnalysisIntentPropertiesArgs',
    'ReachabilityAnalysisIntentPropertiesArgsDict',
    'ReachabilityAnalysisRunPropertiesArgs',
    'ReachabilityAnalysisRunPropertiesArgsDict',
    'StaticCidrPropertiesArgs',
    'StaticCidrPropertiesArgsDict',
    'VerifierWorkspacePropertiesArgs',
    'VerifierWorkspacePropertiesArgsDict',
]

MYPY = False

if not MYPY:
    class AddressPrefixItemArgsDict(TypedDict):
        """
        Address prefix item.
        """
        address_prefix: NotRequired[pulumi.Input[str]]
        """
        Address prefix.
        """
        address_prefix_type: NotRequired[pulumi.Input[Union[str, 'AddressPrefixType']]]
        """
        Address prefix type.
        """
elif False:
    AddressPrefixItemArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AddressPrefixItemArgs:
    def __init__(__self__, *,
                 address_prefix: Optional[pulumi.Input[str]] = None,
                 address_prefix_type: Optional[pulumi.Input[Union[str, 'AddressPrefixType']]] = None):
        """
        Address prefix item.
        :param pulumi.Input[str] address_prefix: Address prefix.
        :param pulumi.Input[Union[str, 'AddressPrefixType']] address_prefix_type: Address prefix type.
        """
        if address_prefix is not None:
            pulumi.set(__self__, "address_prefix", address_prefix)
        if address_prefix_type is not None:
            pulumi.set(__self__, "address_prefix_type", address_prefix_type)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Address prefix.
        """
        return pulumi.get(self, "address_prefix")

    @address_prefix.setter
    def address_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_prefix", value)

    @property
    @pulumi.getter(name="addressPrefixType")
    def address_prefix_type(self) -> Optional[pulumi.Input[Union[str, 'AddressPrefixType']]]:
        """
        Address prefix type.
        """
        return pulumi.get(self, "address_prefix_type")

    @address_prefix_type.setter
    def address_prefix_type(self, value: Optional[pulumi.Input[Union[str, 'AddressPrefixType']]]):
        pulumi.set(self, "address_prefix_type", value)


if not MYPY:
    class IPTrafficArgsDict(TypedDict):
        """
        IP traffic information.
        """
        destination_ips: pulumi.Input[Sequence[pulumi.Input[str]]]
        """
        List of destination IP addresses of the traffic..
        """
        destination_ports: pulumi.Input[Sequence[pulumi.Input[str]]]
        """
        The destination ports of the traffic.
        """
        protocols: pulumi.Input[Sequence[pulumi.Input[Union[str, 'NetworkProtocol']]]]
        source_ips: pulumi.Input[Sequence[pulumi.Input[str]]]
        """
        List of source IP addresses of the traffic..
        """
        source_ports: pulumi.Input[Sequence[pulumi.Input[str]]]
        """
        The source ports of the traffic.
        """
elif False:
    IPTrafficArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IPTrafficArgs:
    def __init__(__self__, *,
                 destination_ips: pulumi.Input[Sequence[pulumi.Input[str]]],
                 destination_ports: pulumi.Input[Sequence[pulumi.Input[str]]],
                 protocols: pulumi.Input[Sequence[pulumi.Input[Union[str, 'NetworkProtocol']]]],
                 source_ips: pulumi.Input[Sequence[pulumi.Input[str]]],
                 source_ports: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        IP traffic information.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] destination_ips: List of destination IP addresses of the traffic..
        :param pulumi.Input[Sequence[pulumi.Input[str]]] destination_ports: The destination ports of the traffic.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_ips: List of source IP addresses of the traffic..
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_ports: The source ports of the traffic.
        """
        pulumi.set(__self__, "destination_ips", destination_ips)
        pulumi.set(__self__, "destination_ports", destination_ports)
        pulumi.set(__self__, "protocols", protocols)
        pulumi.set(__self__, "source_ips", source_ips)
        pulumi.set(__self__, "source_ports", source_ports)

    @property
    @pulumi.getter(name="destinationIps")
    def destination_ips(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of destination IP addresses of the traffic..
        """
        return pulumi.get(self, "destination_ips")

    @destination_ips.setter
    def destination_ips(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "destination_ips", value)

    @property
    @pulumi.getter(name="destinationPorts")
    def destination_ports(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The destination ports of the traffic.
        """
        return pulumi.get(self, "destination_ports")

    @destination_ports.setter
    def destination_ports(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "destination_ports", value)

    @property
    @pulumi.getter
    def protocols(self) -> pulumi.Input[Sequence[pulumi.Input[Union[str, 'NetworkProtocol']]]]:
        return pulumi.get(self, "protocols")

    @protocols.setter
    def protocols(self, value: pulumi.Input[Sequence[pulumi.Input[Union[str, 'NetworkProtocol']]]]):
        pulumi.set(self, "protocols", value)

    @property
    @pulumi.getter(name="sourceIps")
    def source_ips(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of source IP addresses of the traffic..
        """
        return pulumi.get(self, "source_ips")

    @source_ips.setter
    def source_ips(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "source_ips", value)

    @property
    @pulumi.getter(name="sourcePorts")
    def source_ports(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The source ports of the traffic.
        """
        return pulumi.get(self, "source_ports")

    @source_ports.setter
    def source_ports(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "source_ports", value)


if not MYPY:
    class IpamPoolPropertiesArgsDict(TypedDict):
        """
        Properties of IpamPool resource properties which are specific to the Pool resource.
        """
        address_prefixes: pulumi.Input[Sequence[pulumi.Input[str]]]
        """
        List of IP address prefixes of the resource.
        """
        description: NotRequired[pulumi.Input[str]]
        display_name: NotRequired[pulumi.Input[str]]
        """
        String representing a friendly name for the resource.
        """
        parent_pool_name: NotRequired[pulumi.Input[str]]
        """
        String representing parent IpamPool resource name. If empty the IpamPool will be a root pool.
        """
elif False:
    IpamPoolPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpamPoolPropertiesArgs:
    def __init__(__self__, *,
                 address_prefixes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 parent_pool_name: Optional[pulumi.Input[str]] = None):
        """
        Properties of IpamPool resource properties which are specific to the Pool resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_prefixes: List of IP address prefixes of the resource.
        :param pulumi.Input[str] display_name: String representing a friendly name for the resource.
        :param pulumi.Input[str] parent_pool_name: String representing parent IpamPool resource name. If empty the IpamPool will be a root pool.
        """
        pulumi.set(__self__, "address_prefixes", address_prefixes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if parent_pool_name is not None:
            pulumi.set(__self__, "parent_pool_name", parent_pool_name)

    @property
    @pulumi.getter(name="addressPrefixes")
    def address_prefixes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of IP address prefixes of the resource.
        """
        return pulumi.get(self, "address_prefixes")

    @address_prefixes.setter
    def address_prefixes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "address_prefixes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        String representing a friendly name for the resource.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="parentPoolName")
    def parent_pool_name(self) -> Optional[pulumi.Input[str]]:
        """
        String representing parent IpamPool resource name. If empty the IpamPool will be a root pool.
        """
        return pulumi.get(self, "parent_pool_name")

    @parent_pool_name.setter
    def parent_pool_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_pool_name", value)


if not MYPY:
    class NetworkManagerPropertiesNetworkManagerScopesArgsDict(TypedDict):
        """
        Scope of Network Manager.
        """
        management_groups: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        List of management groups.
        """
        subscriptions: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        List of subscriptions.
        """
elif False:
    NetworkManagerPropertiesNetworkManagerScopesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class NetworkManagerPropertiesNetworkManagerScopesArgs:
    def __init__(__self__, *,
                 management_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Scope of Network Manager.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] management_groups: List of management groups.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subscriptions: List of subscriptions.
        """
        if management_groups is not None:
            pulumi.set(__self__, "management_groups", management_groups)
        if subscriptions is not None:
            pulumi.set(__self__, "subscriptions", subscriptions)

    @property
    @pulumi.getter(name="managementGroups")
    def management_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of management groups.
        """
        return pulumi.get(self, "management_groups")

    @management_groups.setter
    def management_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "management_groups", value)

    @property
    @pulumi.getter
    def subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of subscriptions.
        """
        return pulumi.get(self, "subscriptions")

    @subscriptions.setter
    def subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subscriptions", value)


if not MYPY:
    class NetworkManagerSecurityGroupItemArgsDict(TypedDict):
        """
        Network manager security group item.
        """
        network_group_id: pulumi.Input[str]
        """
        Network manager group Id.
        """
elif False:
    NetworkManagerSecurityGroupItemArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class NetworkManagerSecurityGroupItemArgs:
    def __init__(__self__, *,
                 network_group_id: pulumi.Input[str]):
        """
        Network manager security group item.
        :param pulumi.Input[str] network_group_id: Network manager group Id.
        """
        pulumi.set(__self__, "network_group_id", network_group_id)

    @property
    @pulumi.getter(name="networkGroupId")
    def network_group_id(self) -> pulumi.Input[str]:
        """
        Network manager group Id.
        """
        return pulumi.get(self, "network_group_id")

    @network_group_id.setter
    def network_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_group_id", value)


if not MYPY:
    class ReachabilityAnalysisIntentPropertiesArgsDict(TypedDict):
        """
        Represents the Reachability Analysis Intent properties.
        """
        destination_resource_id: pulumi.Input[str]
        """
        Destination resource id to verify the reachability path of.
        """
        ip_traffic: pulumi.Input['IPTrafficArgsDict']
        """
        IP traffic information.
        """
        source_resource_id: pulumi.Input[str]
        """
        Source resource id to verify the reachability path of.
        """
        description: NotRequired[pulumi.Input[str]]
elif False:
    ReachabilityAnalysisIntentPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ReachabilityAnalysisIntentPropertiesArgs:
    def __init__(__self__, *,
                 destination_resource_id: pulumi.Input[str],
                 ip_traffic: pulumi.Input['IPTrafficArgs'],
                 source_resource_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        Represents the Reachability Analysis Intent properties.
        :param pulumi.Input[str] destination_resource_id: Destination resource id to verify the reachability path of.
        :param pulumi.Input['IPTrafficArgs'] ip_traffic: IP traffic information.
        :param pulumi.Input[str] source_resource_id: Source resource id to verify the reachability path of.
        """
        pulumi.set(__self__, "destination_resource_id", destination_resource_id)
        pulumi.set(__self__, "ip_traffic", ip_traffic)
        pulumi.set(__self__, "source_resource_id", source_resource_id)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="destinationResourceId")
    def destination_resource_id(self) -> pulumi.Input[str]:
        """
        Destination resource id to verify the reachability path of.
        """
        return pulumi.get(self, "destination_resource_id")

    @destination_resource_id.setter
    def destination_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_resource_id", value)

    @property
    @pulumi.getter(name="ipTraffic")
    def ip_traffic(self) -> pulumi.Input['IPTrafficArgs']:
        """
        IP traffic information.
        """
        return pulumi.get(self, "ip_traffic")

    @ip_traffic.setter
    def ip_traffic(self, value: pulumi.Input['IPTrafficArgs']):
        pulumi.set(self, "ip_traffic", value)

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> pulumi.Input[str]:
        """
        Source resource id to verify the reachability path of.
        """
        return pulumi.get(self, "source_resource_id")

    @source_resource_id.setter
    def source_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_resource_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


if not MYPY:
    class ReachabilityAnalysisRunPropertiesArgsDict(TypedDict):
        """
        Represents the Reachability Analysis Run properties.
        """
        intent_id: pulumi.Input[str]
        """
        Id of the intent resource to run analysis on.
        """
        description: NotRequired[pulumi.Input[str]]
elif False:
    ReachabilityAnalysisRunPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ReachabilityAnalysisRunPropertiesArgs:
    def __init__(__self__, *,
                 intent_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        Represents the Reachability Analysis Run properties.
        :param pulumi.Input[str] intent_id: Id of the intent resource to run analysis on.
        """
        pulumi.set(__self__, "intent_id", intent_id)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="intentId")
    def intent_id(self) -> pulumi.Input[str]:
        """
        Id of the intent resource to run analysis on.
        """
        return pulumi.get(self, "intent_id")

    @intent_id.setter
    def intent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "intent_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


if not MYPY:
    class StaticCidrPropertiesArgsDict(TypedDict):
        """
        Properties of static CIDR resource.
        """
        address_prefixes: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        List of IP address prefixes of the resource.
        """
        description: NotRequired[pulumi.Input[str]]
        number_of_ip_addresses_to_allocate: NotRequired[pulumi.Input[str]]
        """
        Number of IP addresses to allocate for a static CIDR resource. The IP addresses will be assigned based on IpamPools available space.
        """
elif False:
    StaticCidrPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StaticCidrPropertiesArgs:
    def __init__(__self__, *,
                 address_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 number_of_ip_addresses_to_allocate: Optional[pulumi.Input[str]] = None):
        """
        Properties of static CIDR resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_prefixes: List of IP address prefixes of the resource.
        :param pulumi.Input[str] number_of_ip_addresses_to_allocate: Number of IP addresses to allocate for a static CIDR resource. The IP addresses will be assigned based on IpamPools available space.
        """
        if address_prefixes is not None:
            pulumi.set(__self__, "address_prefixes", address_prefixes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if number_of_ip_addresses_to_allocate is not None:
            pulumi.set(__self__, "number_of_ip_addresses_to_allocate", number_of_ip_addresses_to_allocate)

    @property
    @pulumi.getter(name="addressPrefixes")
    def address_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP address prefixes of the resource.
        """
        return pulumi.get(self, "address_prefixes")

    @address_prefixes.setter
    def address_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "address_prefixes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="numberOfIPAddressesToAllocate")
    def number_of_ip_addresses_to_allocate(self) -> Optional[pulumi.Input[str]]:
        """
        Number of IP addresses to allocate for a static CIDR resource. The IP addresses will be assigned based on IpamPools available space.
        """
        return pulumi.get(self, "number_of_ip_addresses_to_allocate")

    @number_of_ip_addresses_to_allocate.setter
    def number_of_ip_addresses_to_allocate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "number_of_ip_addresses_to_allocate", value)


if not MYPY:
    class VerifierWorkspacePropertiesArgsDict(TypedDict):
        """
        Properties of Verifier Workspace resource.
        """
        description: NotRequired[pulumi.Input[str]]
elif False:
    VerifierWorkspacePropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class VerifierWorkspacePropertiesArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None):
        """
        Properties of Verifier Workspace resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


