# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'BackupType',
    'DayOfWeek',
    'EncryptionAlgorithm',
    'Kind',
    'ManagerSkuType',
    'ManagerType',
    'MonitoringStatus',
    'RecurrenceType',
    'ScheduleStatus',
    'SslStatus',
    'VolumeStatus',
    'VolumeType',
]


@pulumi.type_token("azure-native:storsimple:BackupType")
class BackupType(builtins.str, Enum):
    """
    The type of backup which needs to be taken.
    """
    LOCAL_SNAPSHOT = "LocalSnapshot"
    CLOUD_SNAPSHOT = "CloudSnapshot"


@pulumi.type_token("azure-native:storsimple:DayOfWeek")
class DayOfWeek(builtins.str, Enum):
    SUNDAY = "Sunday"
    MONDAY = "Monday"
    TUESDAY = "Tuesday"
    WEDNESDAY = "Wednesday"
    THURSDAY = "Thursday"
    FRIDAY = "Friday"
    SATURDAY = "Saturday"


@pulumi.type_token("azure-native:storsimple:EncryptionAlgorithm")
class EncryptionAlgorithm(builtins.str, Enum):
    """
    The algorithm used to encrypt "Value".
    """
    NONE = "None"
    AES256 = "AES256"
    RSAE_S_PKCS1_V_1_5 = "RSAES_PKCS1_v_1_5"


@pulumi.type_token("azure-native:storsimple:Kind")
class Kind(builtins.str, Enum):
    """
    The Kind of the object. Currently only Series8000 is supported
    """
    SERIES8000 = "Series8000"


@pulumi.type_token("azure-native:storsimple:ManagerSkuType")
class ManagerSkuType(builtins.str, Enum):
    """
    Refers to the sku name which should be "Standard"
    """
    STANDARD = "Standard"


@pulumi.type_token("azure-native:storsimple:ManagerType")
class ManagerType(builtins.str, Enum):
    """
    The type of StorSimple Manager.
    """
    GARDA_V1 = "GardaV1"
    HELSINKI_V1 = "HelsinkiV1"


@pulumi.type_token("azure-native:storsimple:MonitoringStatus")
class MonitoringStatus(builtins.str, Enum):
    """
    The monitoring status of the volume.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


@pulumi.type_token("azure-native:storsimple:RecurrenceType")
class RecurrenceType(builtins.str, Enum):
    """
    The recurrence type.
    """
    MINUTES = "Minutes"
    HOURLY = "Hourly"
    DAILY = "Daily"
    WEEKLY = "Weekly"


@pulumi.type_token("azure-native:storsimple:ScheduleStatus")
class ScheduleStatus(builtins.str, Enum):
    """
    The schedule status.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


@pulumi.type_token("azure-native:storsimple:SslStatus")
class SslStatus(builtins.str, Enum):
    """
    Signifies whether SSL needs to be enabled or not.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


@pulumi.type_token("azure-native:storsimple:VolumeStatus")
class VolumeStatus(builtins.str, Enum):
    """
    The volume status.
    """
    ONLINE = "Online"
    OFFLINE = "Offline"


@pulumi.type_token("azure-native:storsimple:VolumeType")
class VolumeType(builtins.str, Enum):
    """
    The type of the volume.
    """
    TIERED = "Tiered"
    ARCHIVAL = "Archival"
    LOCALLY_PINNED = "LocallyPinned"
