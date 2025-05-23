# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'AuthenticationType',
    'ConnectionType',
    'DigitalTwinsIdentityType',
    'EndpointType',
    'IdentityType',
    'PrivateLinkServiceConnectionStatus',
    'PublicNetworkAccess',
    'RecordPropertyAndItemRemovals',
]


@pulumi.type_token("azure-native:digitaltwins:AuthenticationType")
class AuthenticationType(builtins.str, Enum):
    """
    Specifies the authentication type being used for connecting to the endpoint. Defaults to 'KeyBased'. If 'KeyBased' is selected, a connection string must be specified (at least the primary connection string). If 'IdentityBased' is select, the endpointUri and entityPath properties must be specified.
    """
    KEY_BASED = "KeyBased"
    IDENTITY_BASED = "IdentityBased"


@pulumi.type_token("azure-native:digitaltwins:ConnectionType")
class ConnectionType(builtins.str, Enum):
    """
    The type of time series connection resource.
    """
    AZURE_DATA_EXPLORER = "AzureDataExplorer"


@pulumi.type_token("azure-native:digitaltwins:DigitalTwinsIdentityType")
class DigitalTwinsIdentityType(builtins.str, Enum):
    """
    The type of Managed Identity used by the DigitalTwinsInstance.
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned,UserAssigned"


@pulumi.type_token("azure-native:digitaltwins:EndpointType")
class EndpointType(builtins.str, Enum):
    """
    The type of Digital Twins endpoint
    """
    EVENT_HUB = "EventHub"
    EVENT_GRID = "EventGrid"
    SERVICE_BUS = "ServiceBus"


@pulumi.type_token("azure-native:digitaltwins:IdentityType")
class IdentityType(builtins.str, Enum):
    """
    The type of managed identity used.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"


@pulumi.type_token("azure-native:digitaltwins:PrivateLinkServiceConnectionStatus")
class PrivateLinkServiceConnectionStatus(builtins.str, Enum):
    """
    The status of a private endpoint connection.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"
    DISCONNECTED = "Disconnected"


@pulumi.type_token("azure-native:digitaltwins:PublicNetworkAccess")
class PublicNetworkAccess(builtins.str, Enum):
    """
    Public network access for the DigitalTwinsInstance.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


@pulumi.type_token("azure-native:digitaltwins:RecordPropertyAndItemRemovals")
class RecordPropertyAndItemRemovals(builtins.str, Enum):
    """
    Specifies whether or not to record twin / relationship property and item removals, including removals of indexed or keyed values (such as map entries, array elements, etc.). This feature is de-activated unless explicitly set to 'true'. Setting this property to 'true' will generate an additional column in the property events table in ADX.
    """
    TRUE = "true"
    FALSE = "false"
