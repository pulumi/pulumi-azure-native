# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AnalyticsConnectorDataDestinationType',
    'AnalyticsConnectorDataSourceType',
    'AnalyticsConnectorMappingType',
    'FhirResourceVersionPolicy',
    'FhirServiceKind',
    'FhirServiceVersion',
    'IotIdentityResolutionType',
    'Kind',
    'ManagedServiceIdentityType',
    'PrivateEndpointServiceConnectionStatus',
    'PublicNetworkAccess',
    'ServiceManagedIdentityType',
    'SmartDataActions',
]


class AnalyticsConnectorDataDestinationType(str, Enum):
    """
    Type of data destination.
    """
    DATALAKE = "datalake"


class AnalyticsConnectorDataSourceType(str, Enum):
    """
    Type of data source.
    """
    FHIRSERVICE = "fhirservice"


class AnalyticsConnectorMappingType(str, Enum):
    """
    Type of data mapping.
    """
    FHIR_TO_PARQUET = "fhirToParquet"


class FhirResourceVersionPolicy(str, Enum):
    """
    Controls how resources are versioned on the FHIR service
    """
    NO_VERSION = "no-version"
    VERSIONED = "versioned"
    VERSIONED_UPDATE = "versioned-update"


class FhirServiceKind(str, Enum):
    """
    The kind of the service.
    """
    FHIR_STU3 = "fhir-Stu3"
    FHIR_R4 = "fhir-R4"


class FhirServiceVersion(str, Enum):
    """
    The kind of FHIR Service.
    """
    STU3 = "STU3"
    R4 = "R4"


class IotIdentityResolutionType(str, Enum):
    """
    Determines how resource identity is resolved on the destination.
    """
    CREATE = "Create"
    LOOKUP = "Lookup"


class Kind(str, Enum):
    """
    The kind of the service.
    """
    FHIR = "fhir"
    FHIR_STU3 = "fhir-Stu3"
    FHIR_R4 = "fhir-R4"


class ManagedServiceIdentityType(str, Enum):
    """
    Type of identity being specified, currently SystemAssigned and None are allowed.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    NONE = "None"


class PrivateEndpointServiceConnectionStatus(str, Enum):
    """
    Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"


class PublicNetworkAccess(str, Enum):
    """
    Control permission for data plane traffic coming from public networks while private endpoint is enabled.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class ServiceManagedIdentityType(str, Enum):
    """
    Type of identity being specified, currently SystemAssigned and None are allowed.
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned,UserAssigned"


class SmartDataActions(str, Enum):
    """
    The Data Actions that can be enabled for a Smart Identity Provider Application.
    """
    READ = "Read"
