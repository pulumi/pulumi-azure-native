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
     * Validate ECE action deployment for a cluster.
     */
    Validate: "Validate",
    /**
     * Deploy ECE action deployment for a cluster.
     */
    Deploy: "Deploy",
} as const;

/**
 * The deployment mode for cluster deployment.
 */
export type DeploymentMode = (typeof DeploymentMode)[keyof typeof DeploymentMode];

export const DeviceKind = {
    /**
     * Arc-enabled edge device with HCI OS.
     */
    HCI: "HCI",
} as const;

/**
 * Device kind to support polymorphic resource.
 */
export type DeviceKind = (typeof DeviceKind)[keyof typeof DeviceKind];

export const DiagnosticLevel = {
    Off: "Off",
    Basic: "Basic",
    Enhanced: "Enhanced",
} as const;

/**
 * Desired level of diagnostic data emitted by the cluster.
 */
export type DiagnosticLevel = (typeof DiagnosticLevel)[keyof typeof DiagnosticLevel];

export const EceSecrets = {
    /**
     * AzureStackLCMUserCredential used for LCM operations for AzureStackHCI cluster.
     */
    AzureStackLCMUserCredential: "AzureStackLCMUserCredential",
    /**
     * DefaultARBApplication used to manage Azure Arc resource bridge (ARB) for AzureStackHCI cluster.
     */
    DefaultARBApplication: "DefaultARBApplication",
    /**
     * LocalAdminCredential used for admin operations for AzureStackHCI cluster.
     */
    LocalAdminCredential: "LocalAdminCredential",
    /**
     * WitnessStorageKey used for setting up a cloud witness for AzureStackHCI cluster.
     */
    WitnessStorageKey: "WitnessStorageKey",
} as const;

/**
 * Secret name expected for Enterprise Cloud Engine (ECE) deployment.
 */
export type EceSecrets = (typeof EceSecrets)[keyof typeof EceSecrets];

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

export const WindowsServerSubscription = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Desired state of Windows Server Subscription.
 */
export type WindowsServerSubscription = (typeof WindowsServerSubscription)[keyof typeof WindowsServerSubscription];
