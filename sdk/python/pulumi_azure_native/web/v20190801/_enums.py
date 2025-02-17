# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AccessControlEntryAction',
    'ComputeModeOptions',
    'InternalLoadBalancingMode',
    'LogLevel',
]


class AccessControlEntryAction(str, Enum):
    """
    Action object.
    """
    PERMIT = "Permit"
    DENY = "Deny"


class ComputeModeOptions(str, Enum):
    """
    Shared or dedicated app hosting.
    """
    SHARED = "Shared"
    DEDICATED = "Dedicated"
    DYNAMIC = "Dynamic"


class InternalLoadBalancingMode(str, Enum):
    """
    Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
    """
    NONE = "None"
    WEB = "Web"
    PUBLISHING = "Publishing"


class LogLevel(str, Enum):
    """
    Log level.
    """
    OFF = "Off"
    VERBOSE = "Verbose"
    INFORMATION = "Information"
    WARNING = "Warning"
    ERROR = "Error"
