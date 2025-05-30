# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'DatabaseCreateMode',
    'IdentityType',
    'PrincipalType',
    'ZoneRedundancy',
]


@pulumi.type_token("azure-native:databasefleetmanager:DatabaseCreateMode")
class DatabaseCreateMode(builtins.str, Enum):
    """
    Create mode. Available options: Default - Create a database. Copy - Copy the source database (source database name must be specified) PointInTimeRestore - Create a database by restoring source database from a point in time (source database name and restore from time must be specified)
    """
    DEFAULT = "Default"
    """
    Create a database.
    """
    COPY = "Copy"
    """
    Copy the source database (source database name must be specified).
    """
    POINT_IN_TIME_RESTORE = "PointInTimeRestore"
    """
    Create a database by restoring source database from a point in time (source database name and restore from time must be specified).
    """


@pulumi.type_token("azure-native:databasefleetmanager:IdentityType")
class IdentityType(builtins.str, Enum):
    """
    Identity type of the main principal.
    """
    NONE = "None"
    """
    No identity.
    """
    USER_ASSIGNED = "UserAssigned"
    """
    User assigned identity.
    """


@pulumi.type_token("azure-native:databasefleetmanager:PrincipalType")
class PrincipalType(builtins.str, Enum):
    """
    Principal type of the main principal.
    """
    APPLICATION = "Application"
    """
    Application principal type.
    """
    USER = "User"
    """
    User principal type.
    """


@pulumi.type_token("azure-native:databasefleetmanager:ZoneRedundancy")
class ZoneRedundancy(builtins.str, Enum):
    """
    Enable zone redundancy for all databases in this tier.
    """
    ENABLED = "Enabled"
    """
    Zone redundancy enabled.
    """
    DISABLED = "Disabled"
    """
    Zone redundancy disabled.
    """
