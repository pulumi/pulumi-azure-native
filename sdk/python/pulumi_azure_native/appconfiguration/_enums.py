# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'AuthenticationMode',
    'ConnectionStatus',
    'CreateMode',
    'IdentityType',
    'PrivateLinkDelegation',
    'PublicNetworkAccess',
]


@pulumi.type_token("azure-native:appconfiguration:AuthenticationMode")
class AuthenticationMode(builtins.str, Enum):
    """
    The data plane proxy authentication mode. This property manages the authentication mode of request to the data plane resources.
    """
    LOCAL = "Local"
    """
    The local authentication mode. Users are not required to have data plane permissions if local authentication is not disabled.
    """
    PASS_THROUGH = "Pass-through"
    """
    The pass-through authentication mode. User identity will be passed through from Azure Resource Manager (ARM), requiring user to have data plane action permissions (Available via App Configuration Data Owner/ App Configuration Data Reader).
    """


@pulumi.type_token("azure-native:appconfiguration:ConnectionStatus")
class ConnectionStatus(builtins.str, Enum):
    """
    The private link service connection status.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"
    DISCONNECTED = "Disconnected"


@pulumi.type_token("azure-native:appconfiguration:CreateMode")
class CreateMode(builtins.str, Enum):
    """
    Indicates whether the configuration store need to be recovered.
    """
    RECOVER = "Recover"
    DEFAULT = "Default"


@pulumi.type_token("azure-native:appconfiguration:IdentityType")
class IdentityType(builtins.str, Enum):
    """
    The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned, UserAssigned"


@pulumi.type_token("azure-native:appconfiguration:PrivateLinkDelegation")
class PrivateLinkDelegation(builtins.str, Enum):
    """
    The data plane proxy private link delegation. This property manages if a request from delegated Azure Resource Manager (ARM) private link is allowed when the data plane resource requires private link.
    """
    ENABLED = "Enabled"
    """
    Azure Resource Manager (ARM) private endpoint is required if the resource requires private link.
    """
    DISABLED = "Disabled"
    """
    Request is denied if the resource requires private link.
    """


@pulumi.type_token("azure-native:appconfiguration:PublicNetworkAccess")
class PublicNetworkAccess(builtins.str, Enum):
    """
    Control permission for data plane traffic coming from public networks while private endpoint is enabled.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"
