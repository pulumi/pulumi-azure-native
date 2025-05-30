// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AmlFilesystemIdentityType = {
    UserAssigned: "UserAssigned",
    None: "None",
} as const;

/**
 * The type of identity used for the resource.
 */
export type AmlFilesystemIdentityType = (typeof AmlFilesystemIdentityType)[keyof typeof AmlFilesystemIdentityType];

export const AmlFilesystemSquashMode = {
    None: "None",
    RootOnly: "RootOnly",
    All: "All",
} as const;

/**
 * Squash mode of the AML file system. 'All': User and Group IDs on files will be squashed to the provided values for all users on non-trusted systems. 'RootOnly': User and Group IDs on files will be squashed to provided values for solely the root user on non-trusted systems. 'None': No squashing of User and Group IDs is performed for any users on any systems.
 */
export type AmlFilesystemSquashMode = (typeof AmlFilesystemSquashMode)[keyof typeof AmlFilesystemSquashMode];

export const AutoExportJobAdminStatus = {
    Enable: "Enable",
    Disable: "Disable",
} as const;

/**
 * The administrative status of the auto export job. Possible values: 'Enable', 'Disable'. Passing in a value of 'Disable' will disable the current active auto export job. By default it is set to 'Enable'.
 */
export type AutoExportJobAdminStatus = (typeof AutoExportJobAdminStatus)[keyof typeof AutoExportJobAdminStatus];

export const AutoExportStatusType = {
    InProgress: "InProgress",
    Disabling: "Disabling",
    Disabled: "Disabled",
    DisableFailed: "DisableFailed",
    Failed: "Failed",
} as const;

/**
 * The operational state of auto export. InProgress indicates the export is running.  Disabling indicates the user has requested to disable the export but the disabling is still in progress. Disabled indicates auto export has been disabled.  DisableFailed indicates the disabling has failed.  Failed means the export was unable to continue, due to a fatal error.
 */
export type AutoExportStatusType = (typeof AutoExportStatusType)[keyof typeof AutoExportStatusType];

export const CacheIdentityType = {
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned, UserAssigned",
    None: "None",
} as const;

/**
 * The type of identity used for the cache
 */
export type CacheIdentityType = (typeof CacheIdentityType)[keyof typeof CacheIdentityType];

export const ConflictResolutionMode = {
    Fail: "Fail",
    Skip: "Skip",
    OverwriteIfDirty: "OverwriteIfDirty",
    OverwriteAlways: "OverwriteAlways",
} as const;

/**
 * How the import job will handle conflicts. For example, if the import job is trying to bring in a directory, but a file is at that path, how it handles it. Fail indicates that the import job should stop immediately and not do anything with the conflict. Skip indicates that it should pass over the conflict. OverwriteIfDirty causes the import job to delete and re-import the file or directory if it is a conflicting type, is dirty, or was not previously imported. OverwriteAlways extends OverwriteIfDirty to include releasing files that had been restored but were not dirty. Please reference https://learn.microsoft.com/en-us/azure/azure-managed-lustre/ for a thorough explanation of these resolution modes.
 */
export type ConflictResolutionMode = (typeof ConflictResolutionMode)[keyof typeof ConflictResolutionMode];

export const MaintenanceDayOfWeekType = {
    Monday: "Monday",
    Tuesday: "Tuesday",
    Wednesday: "Wednesday",
    Thursday: "Thursday",
    Friday: "Friday",
    Saturday: "Saturday",
    Sunday: "Sunday",
} as const;

/**
 * Day of the week on which the maintenance window will occur.
 */
export type MaintenanceDayOfWeekType = (typeof MaintenanceDayOfWeekType)[keyof typeof MaintenanceDayOfWeekType];

export const NfsAccessRuleAccess = {
    No: "no",
    Ro: "ro",
    Rw: "rw",
} as const;

/**
 * Access allowed by this rule.
 */
export type NfsAccessRuleAccess = (typeof NfsAccessRuleAccess)[keyof typeof NfsAccessRuleAccess];

export const NfsAccessRuleScope = {
    Default: "default",
    Network: "network",
    Host: "host",
} as const;

/**
 * Scope for this rule. The scope and filter determine which clients match the rule.
 */
export type NfsAccessRuleScope = (typeof NfsAccessRuleScope)[keyof typeof NfsAccessRuleScope];

export const OperationalStateType = {
    Ready: "Ready",
    Busy: "Busy",
    Suspended: "Suspended",
    Flushing: "Flushing",
} as const;

/**
 * Storage target operational state.
 */
export type OperationalStateType = (typeof OperationalStateType)[keyof typeof OperationalStateType];

export const StorageTargetType = {
    Nfs3: "nfs3",
    Clfs: "clfs",
    Unknown: "unknown",
    BlobNfs: "blobNfs",
} as const;

/**
 * Type of the Storage Target.
 */
export type StorageTargetType = (typeof StorageTargetType)[keyof typeof StorageTargetType];

export const UsernameSource = {
    AD: "AD",
    LDAP: "LDAP",
    File: "File",
    None: "None",
} as const;

/**
 * This setting determines how the cache gets username and group names for clients.
 */
export type UsernameSource = (typeof UsernameSource)[keyof typeof UsernameSource];
