# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'AuthenticationMethod',
    'AutoUpgradeOptions',
    'AzureHybridBenefit',
    'ConnectedClusterKind',
    'PrivateLinkState',
    'ProvisioningState',
    'ResourceIdentityType',
]


@pulumi.type_token("azure-native:kubernetes:AuthenticationMethod")
class AuthenticationMethod(builtins.str, Enum):
    """
    The mode of client authentication.
    """
    TOKEN = "Token"
    AAD = "AAD"


@pulumi.type_token("azure-native:kubernetes:AutoUpgradeOptions")
class AutoUpgradeOptions(builtins.str, Enum):
    """
    Indicates whether the Arc agents on the be upgraded automatically to the latest version. Defaults to Enabled.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


@pulumi.type_token("azure-native:kubernetes:AzureHybridBenefit")
class AzureHybridBenefit(builtins.str, Enum):
    """
    Indicates whether Azure Hybrid Benefit is opted in
    """
    TRUE = "True"
    FALSE = "False"
    NOT_APPLICABLE = "NotApplicable"


@pulumi.type_token("azure-native:kubernetes:ConnectedClusterKind")
class ConnectedClusterKind(builtins.str, Enum):
    """
    The kind of connected cluster.
    """
    PROVISIONED_CLUSTER = "ProvisionedCluster"


@pulumi.type_token("azure-native:kubernetes:PrivateLinkState")
class PrivateLinkState(builtins.str, Enum):
    """
    Property which describes the state of private link on a connected cluster resource.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


@pulumi.type_token("azure-native:kubernetes:ProvisioningState")
class ProvisioningState(builtins.str, Enum):
    """
    Provisioning state of the connected cluster resource.
    """
    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    CANCELED = "Canceled"
    PROVISIONING = "Provisioning"
    UPDATING = "Updating"
    DELETING = "Deleting"
    ACCEPTED = "Accepted"


@pulumi.type_token("azure-native:kubernetes:ResourceIdentityType")
class ResourceIdentityType(builtins.str, Enum):
    """
    The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
