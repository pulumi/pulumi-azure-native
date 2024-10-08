# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AccessKeysAuthentication',
    'AofFrequency',
    'ClusteringPolicy',
    'CmkIdentityType',
    'DeferUpgradeSetting',
    'EvictionPolicy',
    'HighAvailability',
    'ManagedServiceIdentityType',
    'PrivateEndpointServiceConnectionStatus',
    'Protocol',
    'RdbFrequency',
    'SkuName',
    'TlsVersion',
]


class AccessKeysAuthentication(str, Enum):
    """
    This property can be Enabled/Disabled to allow or deny access with the current access keys. Can be updated even after database is created.
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"


class AofFrequency(str, Enum):
    """
    Sets the frequency at which data is written to disk. Defaults to '1s', meaning 'every second'. Note that the 'always' setting is deprecated, because of its performance impact.
    """
    AOF_FREQUENCY_1S = "1s"
    ALWAYS = "always"


class ClusteringPolicy(str, Enum):
    """
    Clustering policy - default is OSSCluster. This property must be chosen at create time, and cannot be changed without deleting the database.
    """
    ENTERPRISE_CLUSTER = "EnterpriseCluster"
    """
    Enterprise clustering policy uses only the classic redis protocol, which does not support redis cluster commands.
    """
    OSS_CLUSTER = "OSSCluster"
    """
    OSS clustering policy follows the redis cluster specification, and requires all clients to support redis clustering.
    """


class CmkIdentityType(str, Enum):
    """
    Only userAssignedIdentity is supported in this API version; other types may be supported in the future
    """
    SYSTEM_ASSIGNED_IDENTITY = "systemAssignedIdentity"
    USER_ASSIGNED_IDENTITY = "userAssignedIdentity"


class DeferUpgradeSetting(str, Enum):
    """
    Option to defer upgrade when newest version is released - default is NotDeferred. Learn more: https://aka.ms/redisversionupgrade
    """
    DEFERRED = "Deferred"
    NOT_DEFERRED = "NotDeferred"


class EvictionPolicy(str, Enum):
    """
    Redis eviction policy - default is VolatileLRU
    """
    ALL_KEYS_LFU = "AllKeysLFU"
    ALL_KEYS_LRU = "AllKeysLRU"
    ALL_KEYS_RANDOM = "AllKeysRandom"
    VOLATILE_LRU = "VolatileLRU"
    VOLATILE_LFU = "VolatileLFU"
    VOLATILE_TTL = "VolatileTTL"
    VOLATILE_RANDOM = "VolatileRandom"
    NO_EVICTION = "NoEviction"


class HighAvailability(str, Enum):
    """
    Enabled by default. If highAvailability is disabled, the data set is not replicated. This affects the availability SLA, and increases the risk of data loss.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class ManagedServiceIdentityType(str, Enum):
    """
    Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned, UserAssigned"


class PrivateEndpointServiceConnectionStatus(str, Enum):
    """
    Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"


class Protocol(str, Enum):
    """
    Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted.
    """
    ENCRYPTED = "Encrypted"
    PLAINTEXT = "Plaintext"


class RdbFrequency(str, Enum):
    """
    Sets the frequency at which a snapshot of the database is created.
    """
    RDB_FREQUENCY_1H = "1h"
    RDB_FREQUENCY_6H = "6h"
    RDB_FREQUENCY_12H = "12h"


class SkuName(str, Enum):
    """
    The level of Redis Enterprise cluster to deploy. Possible values: ('Balanced_B5', 'MemoryOptimized_M10', 'ComputeOptimized_X5', etc.). For more information on SKUs see the latest pricing documentation. Note that additional SKUs may become supported in the future.
    """
    ENTERPRISE_E1 = "Enterprise_E1"
    ENTERPRISE_E5 = "Enterprise_E5"
    ENTERPRISE_E10 = "Enterprise_E10"
    ENTERPRISE_E20 = "Enterprise_E20"
    ENTERPRISE_E50 = "Enterprise_E50"
    ENTERPRISE_E100 = "Enterprise_E100"
    ENTERPRISE_E200 = "Enterprise_E200"
    ENTERPRISE_E400 = "Enterprise_E400"
    ENTERPRISE_FLASH_F300 = "EnterpriseFlash_F300"
    ENTERPRISE_FLASH_F700 = "EnterpriseFlash_F700"
    ENTERPRISE_FLASH_F1500 = "EnterpriseFlash_F1500"
    BALANCED_B0 = "Balanced_B0"
    BALANCED_B1 = "Balanced_B1"
    BALANCED_B3 = "Balanced_B3"
    BALANCED_B5 = "Balanced_B5"
    BALANCED_B10 = "Balanced_B10"
    BALANCED_B20 = "Balanced_B20"
    BALANCED_B50 = "Balanced_B50"
    BALANCED_B100 = "Balanced_B100"
    BALANCED_B150 = "Balanced_B150"
    BALANCED_B250 = "Balanced_B250"
    BALANCED_B350 = "Balanced_B350"
    BALANCED_B500 = "Balanced_B500"
    BALANCED_B700 = "Balanced_B700"
    BALANCED_B1000 = "Balanced_B1000"
    MEMORY_OPTIMIZED_M10 = "MemoryOptimized_M10"
    MEMORY_OPTIMIZED_M20 = "MemoryOptimized_M20"
    MEMORY_OPTIMIZED_M50 = "MemoryOptimized_M50"
    MEMORY_OPTIMIZED_M100 = "MemoryOptimized_M100"
    MEMORY_OPTIMIZED_M150 = "MemoryOptimized_M150"
    MEMORY_OPTIMIZED_M250 = "MemoryOptimized_M250"
    MEMORY_OPTIMIZED_M350 = "MemoryOptimized_M350"
    MEMORY_OPTIMIZED_M500 = "MemoryOptimized_M500"
    MEMORY_OPTIMIZED_M700 = "MemoryOptimized_M700"
    MEMORY_OPTIMIZED_M1000 = "MemoryOptimized_M1000"
    MEMORY_OPTIMIZED_M1500 = "MemoryOptimized_M1500"
    MEMORY_OPTIMIZED_M2000 = "MemoryOptimized_M2000"
    COMPUTE_OPTIMIZED_X3 = "ComputeOptimized_X3"
    COMPUTE_OPTIMIZED_X5 = "ComputeOptimized_X5"
    COMPUTE_OPTIMIZED_X10 = "ComputeOptimized_X10"
    COMPUTE_OPTIMIZED_X20 = "ComputeOptimized_X20"
    COMPUTE_OPTIMIZED_X50 = "ComputeOptimized_X50"
    COMPUTE_OPTIMIZED_X100 = "ComputeOptimized_X100"
    COMPUTE_OPTIMIZED_X150 = "ComputeOptimized_X150"
    COMPUTE_OPTIMIZED_X250 = "ComputeOptimized_X250"
    COMPUTE_OPTIMIZED_X350 = "ComputeOptimized_X350"
    COMPUTE_OPTIMIZED_X500 = "ComputeOptimized_X500"
    COMPUTE_OPTIMIZED_X700 = "ComputeOptimized_X700"
    FLASH_OPTIMIZED_A250 = "FlashOptimized_A250"
    FLASH_OPTIMIZED_A500 = "FlashOptimized_A500"
    FLASH_OPTIMIZED_A700 = "FlashOptimized_A700"
    FLASH_OPTIMIZED_A1000 = "FlashOptimized_A1000"
    FLASH_OPTIMIZED_A1500 = "FlashOptimized_A1500"
    FLASH_OPTIMIZED_A2000 = "FlashOptimized_A2000"
    FLASH_OPTIMIZED_A4500 = "FlashOptimized_A4500"


class TlsVersion(str, Enum):
    """
    The minimum TLS version for the cluster to support, e.g. '1.2'. Newer versions can be added in the future. Note that TLS 1.0 and TLS 1.1 are now completely obsolete -- you cannot use them. They are mentioned only for the sake of consistency with old API versions.
    """
    TLS_VERSION_1_0 = "1.0"
    TLS_VERSION_1_1 = "1.1"
    TLS_VERSION_1_2 = "1.2"
