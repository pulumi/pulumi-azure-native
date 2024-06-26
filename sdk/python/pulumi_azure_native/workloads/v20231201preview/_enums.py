# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AppServicePlanTier',
    'ManagedServiceIdentityType',
    'RoutingPreference',
    'SslPreference',
]


class AppServicePlanTier(str, Enum):
    """
    The App Service plan tier.
    """
    ELASTIC_PREMIUM = "ElasticPremium"
    """
    Elastic Premium plan
    """
    PREMIUM_V3 = "PremiumV3"
    """
    Dedicated Premium V3 plan
    """


class ManagedServiceIdentityType(str, Enum):
    """
    Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned,UserAssigned"


class RoutingPreference(str, Enum):
    """
    Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
    """
    DEFAULT = "Default"
    """
    Default routing preference. Only RFC1918 traffic is routed to the customer VNET.
    """
    ROUTE_ALL = "RouteAll"
    """
    Route all traffic to the customer VNET.
    """


class SslPreference(str, Enum):
    """
    Gets or sets certificate preference if secure communication is enabled.
    """
    DISABLED = "Disabled"
    """
    Secure communication is disabled.
    """
    ROOT_CERTIFICATE = "RootCertificate"
    """
    Secure communication is enabled with root certificate.
    """
    SERVER_CERTIFICATE = "ServerCertificate"
    """
    Secure communication is enabled with server certificate.
    """
