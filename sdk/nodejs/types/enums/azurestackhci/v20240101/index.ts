// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AvailabilityType = {
    Local: "Local",
    Online: "Online",
    Notify: "Notify",
} as const;

/**
 * Indicates the way the update content can be downloaded.
 */
export type AvailabilityType = (typeof AvailabilityType)[keyof typeof AvailabilityType];

export const CloudInitDataSource = {
    NoCloud: "NoCloud",
    Azure: "Azure",
} as const;

/**
 * Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
 */
export type CloudInitDataSource = (typeof CloudInitDataSource)[keyof typeof CloudInitDataSource];

export const ComplianceAssignmentType = {
    /**
     * Report on the state of the machine, but don't make changes.
     */
    Audit: "Audit",
    /**
     * Applied to the machine. If it drifts, the local service inside the machine makes a correction at the next evaluation.
     */
    ApplyAndAutoCorrect: "ApplyAndAutoCorrect",
} as const;

/**
 * Secured Core Compliance Assignment
 */
export type ComplianceAssignmentType = (typeof ComplianceAssignmentType)[keyof typeof ComplianceAssignmentType];

export const DeploymentMode = {
    /**
     * Validate deployment settings for cluster.
     */
    Validate: "Validate",
    /**
     * Deploy cluster using deployment settings.
     */
    Deploy: "Deploy",
} as const;

/**
 * The deployment mode for cluster deployment.
 */
export type DeploymentMode = (typeof DeploymentMode)[keyof typeof DeploymentMode];

export const DiagnosticLevel = {
    Off: "Off",
    Basic: "Basic",
    Enhanced: "Enhanced",
} as const;

/**
 * Desired level of diagnostic data emitted by the cluster.
 */
export type DiagnosticLevel = (typeof DiagnosticLevel)[keyof typeof DiagnosticLevel];

export const DiskFileFormat = {
    Vhdx: "vhdx",
    Vhd: "vhd",
} as const;

/**
 * The format of the actual VHD file [vhd, vhdx]
 */
export type DiskFileFormat = (typeof DiskFileFormat)[keyof typeof DiskFileFormat];

export const ExtendedLocationTypes = {
    CustomLocation: "CustomLocation",
} as const;

/**
 * The type of the extended location.
 */
export type ExtendedLocationTypes = (typeof ExtendedLocationTypes)[keyof typeof ExtendedLocationTypes];

export const HyperVGeneration = {
    V1: "V1",
    V2: "V2",
} as const;

/**
 * The hypervisor generation of the Virtual Machine [V1, V2]
 */
export type HyperVGeneration = (typeof HyperVGeneration)[keyof typeof HyperVGeneration];

export const IPPoolTypeEnum = {
    Vm: "vm",
    Vippool: "vippool",
} as const;

/**
 * Type of the IP Pool [vm, vippool]
 */
export type IPPoolTypeEnum = (typeof IPPoolTypeEnum)[keyof typeof IPPoolTypeEnum];

export const IpAllocationMethodEnum = {
    Dynamic: "Dynamic",
    Static: "Static",
} as const;

/**
 * IPAllocationMethod - The IP address allocation method. Possible values include: 'Static', 'Dynamic'
 */
export type IpAllocationMethodEnum = (typeof IpAllocationMethodEnum)[keyof typeof IpAllocationMethodEnum];

export const ManagedServiceIdentityType = {
    None: "None",
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned, UserAssigned",
} as const;

/**
 * Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
 */
export type ManagedServiceIdentityType = (typeof ManagedServiceIdentityType)[keyof typeof ManagedServiceIdentityType];

export const OperatingSystemTypes = {
    Linux: "Linux",
    Windows: "Windows",
} as const;

/**
 * This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD. Possible values are: **Windows,** **Linux.**
 */
export type OperatingSystemTypes = (typeof OperatingSystemTypes)[keyof typeof OperatingSystemTypes];

export const ProvisioningAction = {
    Install: "install",
    Uninstall: "uninstall",
    Repair: "repair",
} as const;

/**
 * The guest agent provisioning action.
 */
export type ProvisioningAction = (typeof ProvisioningAction)[keyof typeof ProvisioningAction];

export const ResourceIdentityType = {
    SystemAssigned: "SystemAssigned",
} as const;

/**
 * The identity type.
 */
export type ResourceIdentityType = (typeof ResourceIdentityType)[keyof typeof ResourceIdentityType];

export const SecurityTypes = {
    TrustedLaunch: "TrustedLaunch",
    ConfidentialVM: "ConfidentialVM",
} as const;

/**
 * Specifies the SecurityType of the virtual machine. EnableTPM and SecureBootEnabled must be set to true for SecurityType to function.
 */
export type SecurityTypes = (typeof SecurityTypes)[keyof typeof SecurityTypes];

export const ServiceName = {
    WAC: "WAC",
} as const;

/**
 * Name of the service.
 */
export type ServiceName = (typeof ServiceName)[keyof typeof ServiceName];

export const SoftwareAssuranceIntent = {
    Enable: "Enable",
    Disable: "Disable",
} as const;

/**
 * Customer Intent for Software Assurance Benefit.
 */
export type SoftwareAssuranceIntent = (typeof SoftwareAssuranceIntent)[keyof typeof SoftwareAssuranceIntent];

export const State = {
    HasPrerequisite: "HasPrerequisite",
    Obsolete: "Obsolete",
    Ready: "Ready",
    NotApplicableBecauseAnotherUpdateIsInProgress: "NotApplicableBecauseAnotherUpdateIsInProgress",
    Preparing: "Preparing",
    Installing: "Installing",
    Installed: "Installed",
    PreparationFailed: "PreparationFailed",
    InstallationFailed: "InstallationFailed",
    Invalid: "Invalid",
    Recalled: "Recalled",
    Downloading: "Downloading",
    DownloadFailed: "DownloadFailed",
    HealthChecking: "HealthChecking",
    HealthCheckFailed: "HealthCheckFailed",
    ReadyToInstall: "ReadyToInstall",
    ScanInProgress: "ScanInProgress",
    ScanFailed: "ScanFailed",
    AdditionalContentRequired: "AdditionalContentRequired",
} as const;

/**
 * State of the update as it relates to this stamp.
 */
export type State = (typeof State)[keyof typeof State];

export const UpdateRunPropertiesState = {
    Unknown: "Unknown",
    Succeeded: "Succeeded",
    InProgress: "InProgress",
    Failed: "Failed",
} as const;

/**
 * State of the update run.
 */
export type UpdateRunPropertiesState = (typeof UpdateRunPropertiesState)[keyof typeof UpdateRunPropertiesState];

export const UpdateSummariesPropertiesState = {
    Unknown: "Unknown",
    AppliedSuccessfully: "AppliedSuccessfully",
    UpdateAvailable: "UpdateAvailable",
    UpdateInProgress: "UpdateInProgress",
    UpdateFailed: "UpdateFailed",
    NeedsAttention: "NeedsAttention",
    PreparationInProgress: "PreparationInProgress",
    PreparationFailed: "PreparationFailed",
} as const;

/**
 * Overall update state of the stamp.
 */
export type UpdateSummariesPropertiesState = (typeof UpdateSummariesPropertiesState)[keyof typeof UpdateSummariesPropertiesState];

export const VmSizeEnum = {
    Default: "Default",
    Standard_A2_v2: "Standard_A2_v2",
    Standard_A4_v2: "Standard_A4_v2",
    Standard_D2s_v3: "Standard_D2s_v3",
    Standard_D4s_v3: "Standard_D4s_v3",
    Standard_D8s_v3: "Standard_D8s_v3",
    Standard_D16s_v3: "Standard_D16s_v3",
    Standard_D32s_v3: "Standard_D32s_v3",
    Standard_DS2_v2: "Standard_DS2_v2",
    Standard_DS3_v2: "Standard_DS3_v2",
    Standard_DS4_v2: "Standard_DS4_v2",
    Standard_DS5_v2: "Standard_DS5_v2",
    Standard_DS13_v2: "Standard_DS13_v2",
    Standard_K8S_v1: "Standard_K8S_v1",
    Standard_K8S2_v1: "Standard_K8S2_v1",
    Standard_K8S3_v1: "Standard_K8S3_v1",
    Standard_K8S4_v1: "Standard_K8S4_v1",
    Standard_NK6: "Standard_NK6",
    Standard_NK12: "Standard_NK12",
    Standard_NV6: "Standard_NV6",
    Standard_NV12: "Standard_NV12",
    Standard_K8S5_v1: "Standard_K8S5_v1",
    Custom: "Custom",
} as const;

export type VmSizeEnum = (typeof VmSizeEnum)[keyof typeof VmSizeEnum];

export const WindowsServerSubscription = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Desired state of Windows Server Subscription.
 */
export type WindowsServerSubscription = (typeof WindowsServerSubscription)[keyof typeof WindowsServerSubscription];
