# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ArcKindEnum',
    'AssessmentModeTypes',
    'GatewayType',
    'IdentityKeyStore',
    'LicenseAssignmentState',
    'LicenseCoreType',
    'LicenseEdition',
    'LicenseProfileProductType',
    'LicenseProfileSubscriptionStatus',
    'LicenseState',
    'LicenseTarget',
    'LicenseType',
    'PatchModeTypes',
    'ProgramYear',
    'PublicNetworkAccessType',
    'ResourceIdentityType',
    'StatusLevelTypes',
]


class ArcKindEnum(str, Enum):
    """
    Indicates which kind of Arc machine placement on-premises, such as HCI, SCVMM or VMware etc.
    """
    AVS = "AVS"
    HCI = "HCI"
    SCVMM = "SCVMM"
    V_MWARE = "VMware"
    EPS = "EPS"
    GCP = "GCP"
    AWS = "AWS"


class AssessmentModeTypes(str, Enum):
    """
    Specifies the assessment mode.
    """
    IMAGE_DEFAULT = "ImageDefault"
    AUTOMATIC_BY_PLATFORM = "AutomaticByPlatform"


class GatewayType(str, Enum):
    """
    The type of the Gateway resource.
    """
    PUBLIC = "Public"


class IdentityKeyStore(str, Enum):
    """
    Specifies the identity key store a machine is using.
    """
    TPM = "TPM"
    DEFAULT = "Default"


class LicenseAssignmentState(str, Enum):
    """
    Describes the license assignment state (Assigned or NotAssigned).
    """
    ASSIGNED = "Assigned"
    NOT_ASSIGNED = "NotAssigned"


class LicenseCoreType(str, Enum):
    """
    Describes the license core type (pCore or vCore).
    """
    P_CORE = "pCore"
    V_CORE = "vCore"


class LicenseEdition(str, Enum):
    """
    Describes the edition of the license. The values are either Standard or Datacenter.
    """
    STANDARD = "Standard"
    DATACENTER = "Datacenter"


class LicenseProfileProductType(str, Enum):
    """
    Indicates the product type of the license.
    """
    WINDOWS_SERVER = "WindowsServer"
    WINDOWS_IO_T_ENTERPRISE = "WindowsIoTEnterprise"


class LicenseProfileSubscriptionStatus(str, Enum):
    """
    Indicates the subscription status of the product.
    """
    UNKNOWN = "Unknown"
    ENABLING = "Enabling"
    ENABLED = "Enabled"
    DISABLED = "Disabled"
    DISABLING = "Disabling"
    FAILED = "Failed"


class LicenseState(str, Enum):
    """
    Describes the state of the license.
    """
    ACTIVATED = "Activated"
    DEACTIVATED = "Deactivated"


class LicenseTarget(str, Enum):
    """
    Describes the license target server.
    """
    WINDOWS_SERVER_2012 = "Windows Server 2012"
    WINDOWS_SERVER_2012_R2 = "Windows Server 2012 R2"


class LicenseType(str, Enum):
    """
    The type of the license resource.
    """
    ESU = "ESU"


class PatchModeTypes(str, Enum):
    """
    Specifies the patch mode.
    """
    IMAGE_DEFAULT = "ImageDefault"
    AUTOMATIC_BY_PLATFORM = "AutomaticByPlatform"
    AUTOMATIC_BY_OS = "AutomaticByOS"
    MANUAL = "Manual"


class ProgramYear(str, Enum):
    """
    Describes the program year the volume license is for.
    """
    YEAR_1 = "Year 1"
    YEAR_2 = "Year 2"
    YEAR_3 = "Year 3"


class PublicNetworkAccessType(str, Enum):
    """
    Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
    """
    ENABLED = "Enabled"
    """
    Allows Azure Arc agents to communicate with Azure Arc services over both public (internet) and private endpoints.
    """
    DISABLED = "Disabled"
    """
    Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet) endpoints. The agents must use the private link.
    """
    SECURED_BY_PERIMETER = "SecuredByPerimeter"
    """
    Azure Arc agent communication with Azure Arc services over public (internet) is enforced by Network Security Perimeter (NSP)
    """


class ResourceIdentityType(str, Enum):
    """
    The identity type.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"


class StatusLevelTypes(str, Enum):
    """
    The level code.
    """
    INFO = "Info"
    WARNING = "Warning"
    ERROR = "Error"
