# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'AccessRights',
    'AuthenticationType',
    'Capabilities',
    'DefaultAction',
    'IotHubSku',
    'IpFilterActionType',
    'NetworkRuleIPAction',
    'PrivateLinkServiceConnectionStatus',
    'PublicNetworkAccess',
    'ResourceIdentityType',
    'RoutingSource',
]


@pulumi.type_token("azure-native:iothub:AccessRights")
class AccessRights(builtins.str, Enum):
    """
    The permissions assigned to the shared access policy.
    """
    REGISTRY_READ = "RegistryRead"
    REGISTRY_WRITE = "RegistryWrite"
    SERVICE_CONNECT = "ServiceConnect"
    DEVICE_CONNECT = "DeviceConnect"
    REGISTRY_READ_REGISTRY_WRITE = "RegistryRead, RegistryWrite"
    REGISTRY_READ_SERVICE_CONNECT = "RegistryRead, ServiceConnect"
    REGISTRY_READ_DEVICE_CONNECT = "RegistryRead, DeviceConnect"
    REGISTRY_WRITE_SERVICE_CONNECT = "RegistryWrite, ServiceConnect"
    REGISTRY_WRITE_DEVICE_CONNECT = "RegistryWrite, DeviceConnect"
    SERVICE_CONNECT_DEVICE_CONNECT = "ServiceConnect, DeviceConnect"
    REGISTRY_READ_REGISTRY_WRITE_SERVICE_CONNECT = "RegistryRead, RegistryWrite, ServiceConnect"
    REGISTRY_READ_REGISTRY_WRITE_DEVICE_CONNECT = "RegistryRead, RegistryWrite, DeviceConnect"
    REGISTRY_READ_SERVICE_CONNECT_DEVICE_CONNECT = "RegistryRead, ServiceConnect, DeviceConnect"
    REGISTRY_WRITE_SERVICE_CONNECT_DEVICE_CONNECT = "RegistryWrite, ServiceConnect, DeviceConnect"
    REGISTRY_READ_REGISTRY_WRITE_SERVICE_CONNECT_DEVICE_CONNECT = "RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect"


@pulumi.type_token("azure-native:iothub:AuthenticationType")
class AuthenticationType(builtins.str, Enum):
    """
    Specifies authentication type being used for connecting to the storage account.
    """
    KEY_BASED = "keyBased"
    IDENTITY_BASED = "identityBased"


@pulumi.type_token("azure-native:iothub:Capabilities")
class Capabilities(builtins.str, Enum):
    """
    The capabilities and features enabled for the IoT hub.
    """
    NONE = "None"
    DEVICE_MANAGEMENT = "DeviceManagement"


@pulumi.type_token("azure-native:iothub:DefaultAction")
class DefaultAction(builtins.str, Enum):
    """
    Default Action for Network Rule Set
    """
    DENY = "Deny"
    ALLOW = "Allow"


@pulumi.type_token("azure-native:iothub:IotHubSku")
class IotHubSku(builtins.str, Enum):
    """
    The name of the SKU.
    """
    F1 = "F1"
    S1 = "S1"
    S2 = "S2"
    S3 = "S3"
    B1 = "B1"
    B2 = "B2"
    B3 = "B3"


@pulumi.type_token("azure-native:iothub:IpFilterActionType")
class IpFilterActionType(builtins.str, Enum):
    """
    The desired action for requests captured by this rule.
    """
    ACCEPT = "Accept"
    REJECT = "Reject"


@pulumi.type_token("azure-native:iothub:NetworkRuleIPAction")
class NetworkRuleIPAction(builtins.str, Enum):
    """
    IP Filter Action
    """
    ALLOW = "Allow"


@pulumi.type_token("azure-native:iothub:PrivateLinkServiceConnectionStatus")
class PrivateLinkServiceConnectionStatus(builtins.str, Enum):
    """
    The status of a private endpoint connection
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"
    DISCONNECTED = "Disconnected"


@pulumi.type_token("azure-native:iothub:PublicNetworkAccess")
class PublicNetworkAccess(builtins.str, Enum):
    """
    Whether requests from Public Network are allowed
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


@pulumi.type_token("azure-native:iothub:ResourceIdentityType")
class ResourceIdentityType(builtins.str, Enum):
    """
    The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned, UserAssigned"
    NONE = "None"


@pulumi.type_token("azure-native:iothub:RoutingSource")
class RoutingSource(builtins.str, Enum):
    """
    The source that the routing rule is to be applied to, such as DeviceMessages.
    """
    INVALID = "Invalid"
    DEVICE_MESSAGES = "DeviceMessages"
    TWIN_CHANGE_EVENTS = "TwinChangeEvents"
    DEVICE_LIFECYCLE_EVENTS = "DeviceLifecycleEvents"
    DEVICE_JOB_LIFECYCLE_EVENTS = "DeviceJobLifecycleEvents"
    DEVICE_CONNECTION_STATE_EVENTS = "DeviceConnectionStateEvents"
