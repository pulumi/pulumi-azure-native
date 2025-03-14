# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CreatedByType',
    'DiagnosticLevel',
    'WindowsServerSubscription',
]


class CreatedByType(str, Enum):
    """
    The type of identity that last modified the resource.
    """
    USER = "User"
    APPLICATION = "Application"
    MANAGED_IDENTITY = "ManagedIdentity"
    KEY = "Key"


class DiagnosticLevel(str, Enum):
    """
    Desired level of diagnostic data emitted by the cluster.
    """
    OFF = "Off"
    BASIC = "Basic"
    ENHANCED = "Enhanced"


class WindowsServerSubscription(str, Enum):
    """
    Desired state of Windows Server Subscription.
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"
