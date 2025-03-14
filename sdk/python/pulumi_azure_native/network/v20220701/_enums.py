# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ForwardingRuleState',
    'IpAllocationMethod',
    'RouteNextHopType',
    'SecurityRuleAccess',
    'SecurityRuleDirection',
    'SecurityRuleProtocol',
    'VirtualNetworkPrivateEndpointNetworkPolicies',
    'VirtualNetworkPrivateLinkServiceNetworkPolicies',
]


class ForwardingRuleState(str, Enum):
    """
    The state of forwarding rule.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class IpAllocationMethod(str, Enum):
    """
    Private IP address allocation method.
    """
    STATIC = "Static"
    DYNAMIC = "Dynamic"


class RouteNextHopType(str, Enum):
    """
    The type of Azure hop the packet should be sent to.
    """
    VIRTUAL_NETWORK_GATEWAY = "VirtualNetworkGateway"
    VNET_LOCAL = "VnetLocal"
    INTERNET = "Internet"
    VIRTUAL_APPLIANCE = "VirtualAppliance"
    NONE = "None"


class SecurityRuleAccess(str, Enum):
    """
    The network traffic is allowed or denied.
    """
    ALLOW = "Allow"
    DENY = "Deny"


class SecurityRuleDirection(str, Enum):
    """
    The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
    """
    INBOUND = "Inbound"
    OUTBOUND = "Outbound"


class SecurityRuleProtocol(str, Enum):
    """
    Network protocol this rule applies to.
    """
    TCP = "Tcp"
    UDP = "Udp"
    ICMP = "Icmp"
    ESP = "Esp"
    ASTERISK = "*"
    AH = "Ah"


class VirtualNetworkPrivateEndpointNetworkPolicies(str, Enum):
    """
    Enable or Disable apply network policies on private end point in the subnet.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class VirtualNetworkPrivateLinkServiceNetworkPolicies(str, Enum):
    """
    Enable or Disable apply network policies on private link service in the subnet.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"
