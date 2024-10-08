# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ApiKey',
    'AutoGeneratedDomainNameLabelScope',
    'DeterministicOutboundIP',
    'ManagedServiceIdentityType',
    'MarketplaceAutoRenew',
    'PrivateEndpointServiceConnectionStatus',
    'PublicNetworkAccess',
    'StartTLSPolicy',
    'ZoneRedundancy',
]


class ApiKey(str, Enum):
    """
    The api key setting of the Grafana instance.
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"


class AutoGeneratedDomainNameLabelScope(str, Enum):
    """
    Scope for dns deterministic name hash calculation.
    """
    TENANT_REUSE = "TenantReuse"


class DeterministicOutboundIP(str, Enum):
    """
    Whether a Grafana instance uses deterministic outbound IPs.
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"


class ManagedServiceIdentityType(str, Enum):
    """
    Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned,UserAssigned"


class MarketplaceAutoRenew(str, Enum):
    """
    The AutoRenew setting of the Enterprise subscription
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"


class PrivateEndpointServiceConnectionStatus(str, Enum):
    """
    Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"


class PublicNetworkAccess(str, Enum):
    """
    Indicate the state for enable or disable traffic over the public interface.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class StartTLSPolicy(str, Enum):
    """
    The StartTLSPolicy setting of the SMTP configuration
    https://pkg.go.dev/github.com/go-mail/mail#StartTLSPolicy
    """
    OPPORTUNISTIC_START_TLS = "OpportunisticStartTLS"
    MANDATORY_START_TLS = "MandatoryStartTLS"
    NO_START_TLS = "NoStartTLS"


class ZoneRedundancy(str, Enum):
    """
    The zone redundancy setting of the Grafana instance.
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"
