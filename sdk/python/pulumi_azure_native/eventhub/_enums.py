# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AccessRights',
    'ApplicationGroupPolicyType',
    'CaptureIdentityType',
    'CleanupPolicyRetentionDescription',
    'ClusterSkuName',
    'DefaultAction',
    'EncodingCaptureDescription',
    'EndPointProvisioningState',
    'EntityStatus',
    'IPAction',
    'KeySource',
    'ManagedServiceIdentityType',
    'MetricId',
    'NetworkRuleIPAction',
    'PrivateLinkConnectionStatus',
    'PublicNetworkAccess',
    'PublicNetworkAccessFlag',
    'SchemaCompatibility',
    'SchemaType',
    'SkuName',
    'SkuTier',
    'TlsVersion',
]


class AccessRights(str, Enum):
    MANAGE = "Manage"
    SEND = "Send"
    LISTEN = "Listen"


class ApplicationGroupPolicyType(str, Enum):
    """
    Application Group Policy types
    """
    THROTTLING_POLICY = "ThrottlingPolicy"


class CaptureIdentityType(str, Enum):
    """
    Type of Azure Active Directory Managed Identity.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"


class CleanupPolicyRetentionDescription(str, Enum):
    """
    Enumerates the possible values for cleanup policy
    """
    DELETE = "Delete"
    COMPACT = "Compact"


class ClusterSkuName(str, Enum):
    """
    Name of this SKU.
    """
    DEDICATED = "Dedicated"


class DefaultAction(str, Enum):
    """
    Default Action for Network Rule Set
    """
    ALLOW = "Allow"
    DENY = "Deny"


class EncodingCaptureDescription(str, Enum):
    """
    Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be deprecated in New API Version
    """
    AVRO = "Avro"
    AVRO_DEFLATE = "AvroDeflate"


class EndPointProvisioningState(str, Enum):
    """
    Provisioning state of the Private Endpoint Connection.
    """
    CREATING = "Creating"
    UPDATING = "Updating"
    DELETING = "Deleting"
    SUCCEEDED = "Succeeded"
    CANCELED = "Canceled"
    FAILED = "Failed"


class EntityStatus(str, Enum):
    """
    Enumerates the possible values for the status of the Event Hub.
    """
    ACTIVE = "Active"
    DISABLED = "Disabled"
    RESTORING = "Restoring"
    SEND_DISABLED = "SendDisabled"
    RECEIVE_DISABLED = "ReceiveDisabled"
    CREATING = "Creating"
    DELETING = "Deleting"
    RENAMING = "Renaming"
    UNKNOWN = "Unknown"


class IPAction(str, Enum):
    """
    The IP Filter Action
    """
    ACCEPT = "Accept"
    REJECT = "Reject"


class KeySource(str, Enum):
    """
    Enumerates the possible value of keySource for Encryption
    """
    MICROSOFT_KEY_VAULT = "Microsoft.KeyVault"


class ManagedServiceIdentityType(str, Enum):
    """
    Type of managed service identity.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned, UserAssigned"
    NONE = "None"


class MetricId(str, Enum):
    """
    Metric Id on which the throttle limit should be set, MetricId can be discovered by hovering over Metric in the Metrics section of Event Hub Namespace inside Azure Portal
    """
    INCOMING_BYTES = "IncomingBytes"
    OUTGOING_BYTES = "OutgoingBytes"
    INCOMING_MESSAGES = "IncomingMessages"
    OUTGOING_MESSAGES = "OutgoingMessages"


class NetworkRuleIPAction(str, Enum):
    """
    The IP Filter Action
    """
    ALLOW = "Allow"


class PrivateLinkConnectionStatus(str, Enum):
    """
    Status of the connection.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"
    DISCONNECTED = "Disconnected"


class PublicNetworkAccess(str, Enum):
    """
    This determines if traffic is allowed over public network. By default it is enabled.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"
    SECURED_BY_PERIMETER = "SecuredByPerimeter"


class PublicNetworkAccessFlag(str, Enum):
    """
    This determines if traffic is allowed over public network. By default it is enabled. If value is SecuredByPerimeter then Inbound and Outbound communication is controlled by the network security perimeter and profile's access rules. 
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"
    SECURED_BY_PERIMETER = "SecuredByPerimeter"


class SchemaCompatibility(str, Enum):
    NONE = "None"
    BACKWARD = "Backward"
    FORWARD = "Forward"


class SchemaType(str, Enum):
    UNKNOWN = "Unknown"
    AVRO = "Avro"


class SkuName(str, Enum):
    """
    Name of this SKU.
    """
    BASIC = "Basic"
    STANDARD = "Standard"
    PREMIUM = "Premium"


class SkuTier(str, Enum):
    """
    The billing tier of this particular SKU.
    """
    BASIC = "Basic"
    STANDARD = "Standard"
    PREMIUM = "Premium"


class TlsVersion(str, Enum):
    """
    The minimum TLS version for the cluster to support, e.g. '1.2'
    """
    TLS_VERSION_1_0 = "1.0"
    TLS_VERSION_1_1 = "1.1"
    TLS_VERSION_1_2 = "1.2"
