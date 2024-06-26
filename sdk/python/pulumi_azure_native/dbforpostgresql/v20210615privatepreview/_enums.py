# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CreateMode',
    'GeoRedundantBackupEnum',
    'HighAvailabilityMode',
    'ResourceIdentityType',
    'ServerVersion',
    'SkuTier',
]


class CreateMode(str, Enum):
    """
    The mode to create a new PostgreSQL server.
    """
    DEFAULT = "Default"
    CREATE = "Create"
    UPDATE = "Update"
    POINT_IN_TIME_RESTORE = "PointInTimeRestore"


class GeoRedundantBackupEnum(str, Enum):
    """
    A value indicating whether Geo-Redundant backup is enabled on the server.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class HighAvailabilityMode(str, Enum):
    """
    The HA mode for the server.
    """
    DISABLED = "Disabled"
    ZONE_REDUNDANT = "ZoneRedundant"


class ResourceIdentityType(str, Enum):
    """
    The identity type.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"


class ServerVersion(str, Enum):
    """
    PostgreSQL Server version.
    """
    SERVER_VERSION_13 = "13"
    SERVER_VERSION_12 = "12"
    SERVER_VERSION_11 = "11"


class SkuTier(str, Enum):
    """
    The tier of the particular SKU, e.g. Burstable.
    """
    BURSTABLE = "Burstable"
    GENERAL_PURPOSE = "GeneralPurpose"
    MEMORY_OPTIMIZED = "MemoryOptimized"
