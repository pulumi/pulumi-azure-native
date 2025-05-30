# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'ManagedServiceIdentityType',
    'Type',
]


@pulumi.type_token("azure-native:loadtestservice:ManagedServiceIdentityType")
class ManagedServiceIdentityType(builtins.str, Enum):
    """
    Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned,UserAssigned"


@pulumi.type_token("azure-native:loadtestservice:Type")
class Type(builtins.str, Enum):
    """
    Managed identity type to use for accessing encryption key Url.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    """
    System assigned identity.
    """
    USER_ASSIGNED = "UserAssigned"
    """
    User assigned identity.
    """
