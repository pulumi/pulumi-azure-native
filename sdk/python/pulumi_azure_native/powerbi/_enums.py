# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'AzureSkuName',
    'AzureSkuTier',
    'PersistedConnectionStatus',
    'ResourceProvisioningState',
]


@pulumi.type_token("azure-native:powerbi:AzureSkuName")
class AzureSkuName(builtins.str, Enum):
    """
    SKU name
    """
    S1 = "S1"


@pulumi.type_token("azure-native:powerbi:AzureSkuTier")
class AzureSkuTier(builtins.str, Enum):
    """
    SKU tier
    """
    STANDARD = "Standard"


@pulumi.type_token("azure-native:powerbi:PersistedConnectionStatus")
class PersistedConnectionStatus(builtins.str, Enum):
    """
    Status of the connection.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"
    DISCONNECTED = "Disconnected"


@pulumi.type_token("azure-native:powerbi:ResourceProvisioningState")
class ResourceProvisioningState(builtins.str, Enum):
    """
    Provisioning state of the Private Endpoint Connection.
    """
    CREATING = "Creating"
    UPDATING = "Updating"
    DELETING = "Deleting"
    SUCCEEDED = "Succeeded"
    CANCELED = "Canceled"
    FAILED = "Failed"
