# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AadAuthFailureMode',
    'HostingMode',
    'IdentityType',
    'PublicNetworkAccess',
    'SearchBypass',
    'SearchDisabledDataExfiltrationOption',
    'SearchEncryptionWithCmk',
    'SearchSemanticSearch',
    'SkuName',
]


class AadAuthFailureMode(str, Enum):
    """
    Describes what response the data plane API of a Search service would send for requests that failed authentication.
    """
    HTTP403 = "http403"
    """
    Indicates that requests that failed authentication should be presented with an HTTP status code of 403 (Forbidden).
    """
    HTTP401_WITH_BEARER_CHALLENGE = "http401WithBearerChallenge"
    """
    Indicates that requests that failed authentication should be presented with an HTTP status code of 401 (Unauthorized) and present a Bearer Challenge.
    """


class HostingMode(str, Enum):
    """
    Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
    """
    DEFAULT = "default"
    """
    The limit on number of indexes is determined by the default limits for the SKU.
    """
    HIGH_DENSITY = "highDensity"
    """
    Only application for standard3 SKU, where the search service can have up to 1000 indexes.
    """


class IdentityType(str, Enum):
    """
    The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an identity created by the system and a set of user assigned identities. The type 'None' will remove all identities from the service.
    """
    NONE = "None"
    """
    Indicates that any identity associated with the search service needs to be removed.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    """
    Indicates that system-assigned identity for the search service will be enabled.
    """
    USER_ASSIGNED = "UserAssigned"
    """
    Indicates that one or more user assigned identities will be assigned to the search service.
    """
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned, UserAssigned"
    """
    Indicates that system-assigned identity for the search service will be enabled along with the assignment of one or more user assigned identities.
    """


class PublicNetworkAccess(str, Enum):
    """
    This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
    """
    ENABLED = "enabled"
    """
    The search service is accessible from traffic originating from the public internet.
    """
    DISABLED = "disabled"
    """
    The search service is not accessible from traffic originating from the public internet. Access is only permitted over approved private endpoint connections.
    """


class SearchBypass(str, Enum):
    """
    Possible origins of inbound traffic that can bypass the rules defined in the 'ipRules' section.
    """
    NONE = "None"
    """
    Indicates that no origin can bypass the rules defined in the 'ipRules' section. This is the default.
    """
    AZURE_PORTAL = "AzurePortal"
    """
    Indicates that requests originating from the Azure portal can bypass the rules defined in the 'ipRules' section.
    """


class SearchDisabledDataExfiltrationOption(str, Enum):
    """
    A specific data exfiltration scenario that is disabled for the service.
    """
    ALL = "All"
    """
    Indicates that all data exfiltration scenarios are disabled.
    """


class SearchEncryptionWithCmk(str, Enum):
    """
    Describes how a search service should enforce having one or more non customer encrypted resources.
    """
    DISABLED = "Disabled"
    """
    No enforcement will be made and the search service can have non customer encrypted resources.
    """
    ENABLED = "Enabled"
    """
    Search service will be marked as non-compliant if there are one or more non customer encrypted resources.
    """
    UNSPECIFIED = "Unspecified"
    """
    Enforcement policy is not explicitly specified, with the behavior being the same as if it were set to 'Disabled'.
    """


class SearchSemanticSearch(str, Enum):
    """
    Sets options that control the availability of semantic search. This configuration is only possible for certain Azure Cognitive Search SKUs in certain locations.
    """
    DISABLED = "disabled"
    """
    Indicates that semantic search is disabled for the search service. This is the default.
    """
    FREE = "free"
    """
    Enables semantic search on a search service and indicates that it is to be used within the limits of the free tier. This would cap the volume of semantic search requests and is offered at no extra charge.
    """
    STANDARD = "standard"
    """
    Enables semantic search on a search service as a billable feature, with higher throughput and volume of semantic search queries.
    """


class SkuName(str, Enum):
    """
    The SKU of the search service. Valid values include: 'free': Shared service. 'basic': Dedicated service with up to 3 replicas. 'standard': Dedicated service with up to 12 partitions and 12 replicas. 'standard2': Similar to standard, but with more capacity per search unit. 'standard3': The largest Standard offering with up to 12 partitions and 12 replicas (or up to 3 partitions with more indexes if you also set the hostingMode property to 'highDensity'). 'storage_optimized_l1': Supports 1TB per partition, up to 12 partitions. 'storage_optimized_l2': Supports 2TB per partition, up to 12 partitions.'
    """
    FREE = "free"
    """
    Free tier, with no SLA guarantees and a subset of features offered to paid tiers.
    """
    BASIC = "basic"
    """
    Paid tier dedicated service with up to 3 replicas.
    """
    STANDARD = "standard"
    """
    Paid tier dedicated service with up to 12 partitions and 12 replicas.
    """
    STANDARD2 = "standard2"
    """
    Similar to 'standard', but with more capacity per search unit.
    """
    STANDARD3 = "standard3"
    """
     The largest Standard offering with up to 12 partitions and 12 replicas (or up to 3 partitions with more indexes if you also set the hostingMode property to 'highDensity').
    """
    STORAGE_OPTIMIZED_L1 = "storage_optimized_l1"
    """
    Paid tier dedicated service that supports 1TB per partition, up to 12 partitions.
    """
    STORAGE_OPTIMIZED_L2 = "storage_optimized_l2"
    """
    Paid tier dedicated service that supports 2TB per partition, up to 12 partitions.
    """
