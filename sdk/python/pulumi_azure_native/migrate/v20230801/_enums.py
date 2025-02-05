# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'MoveType',
    'ResourceIdentityType',
    'TargetAvailabilityZone',
    'ZoneRedundant',
]


class MoveType(str, Enum):
    """
    Defines the MoveType.
    """
    REGION_TO_REGION = "RegionToRegion"
    REGION_TO_ZONE = "RegionToZone"


class ResourceIdentityType(str, Enum):
    """
    The type of identity used for the resource mover service.
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"


class TargetAvailabilityZone(str, Enum):
    """
    Gets or sets the target availability zone.
    """
    ONE = "1"
    TWO = "2"
    THREE = "3"
    NA = "NA"


class ZoneRedundant(str, Enum):
    """
    Defines the zone redundant resource setting.
    """
    ENABLE = "Enable"
    DISABLE = "Disable"
