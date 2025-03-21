# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'PublicNetworkAccess',
    'SkuFamily',
    'SkuName',
    'TlsVersion',
]


class PublicNetworkAccess(str, Enum):
    """
    Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is 'Enabled'
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class SkuFamily(str, Enum):
    """
    The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
    """
    C = "C"
    P = "P"


class SkuName(str, Enum):
    """
    The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
    """
    BASIC = "Basic"
    STANDARD = "Standard"
    PREMIUM = "Premium"


class TlsVersion(str, Enum):
    """
    Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
    """
    TLS_VERSION_1_0 = "1.0"
    TLS_VERSION_1_1 = "1.1"
    TLS_VERSION_1_2 = "1.2"
