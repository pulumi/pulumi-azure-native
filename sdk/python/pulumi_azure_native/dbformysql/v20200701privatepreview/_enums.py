# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CreateMode',
    'HaEnabledEnum',
    'InfrastructureEncryptionEnum',
    'ResourceIdentityType',
    'ServerKeyType',
    'ServerVersion',
    'SkuTier',
    'SslEnforcementEnum',
    'StorageAutogrow',
]


class CreateMode(str, Enum):
    """
    The mode to create a new MySQL server.
    """
    DEFAULT = "Default"
    POINT_IN_TIME_RESTORE = "PointInTimeRestore"
    REPLICA = "Replica"


class HaEnabledEnum(str, Enum):
    """
    Enable HA or not for a server.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class InfrastructureEncryptionEnum(str, Enum):
    """
    Status showing whether the server enabled infrastructure encryption.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class ResourceIdentityType(str, Enum):
    """
    The identity type.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"


class ServerKeyType(str, Enum):
    """
    The key type like 'AzureKeyVault'.
    """
    AZURE_KEY_VAULT = "AzureKeyVault"


class ServerVersion(str, Enum):
    """
    Server version.
    """
    SERVER_VERSION_5_7 = "5.7"


class SkuTier(str, Enum):
    """
    The tier of the particular SKU, e.g. GeneralPurpose.
    """
    BURSTABLE = "Burstable"
    GENERAL_PURPOSE = "GeneralPurpose"
    MEMORY_OPTIMIZED = "MemoryOptimized"


class SslEnforcementEnum(str, Enum):
    """
    Enable ssl enforcement or not when connect to server.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class StorageAutogrow(str, Enum):
    """
    Enable Storage Auto Grow.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"
