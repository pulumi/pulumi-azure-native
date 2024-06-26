// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AzureDevOpsPermissionType = {
    /**
     * Pool will inherit permissions from the project or organization.
     */
    Inherit: "Inherit",
    /**
     * Only the pool creator will be an admin of the pool.
     */
    CreatorOnly: "CreatorOnly",
    /**
     * Only the specified accounts will be admins of the pool.
     */
    SpecificAccounts: "SpecificAccounts",
} as const;

/**
 * Determines who has admin permissions to the Azure DevOps pool.
 */
export type AzureDevOpsPermissionType = (typeof AzureDevOpsPermissionType)[keyof typeof AzureDevOpsPermissionType];

export const CachingType = {
    /**
     * Don't use host caching.
     */
    None: "None",
    /**
     * For workloads that only do read operations.
     */
    ReadOnly: "ReadOnly",
    /**
     * For workloads that do a balance of read and write operations.
     */
    ReadWrite: "ReadWrite",
} as const;

/**
 * The type of caching to be enabled for the data disks. The default value for caching is readwrite. For information about the caching options see: https://blogs.msdn.microsoft.com/windowsazurestorage/2012/06/27/exploring-windows-azure-drives-disks-and-images/.
 */
export type CachingType = (typeof CachingType)[keyof typeof CachingType];

export const LogonType = {
    /**
     * Run as a service.
     */
    Service: "Service",
    /**
     * Run in interactive mode.
     */
    Interactive: "Interactive",
} as const;

/**
 * Determines how the service should be run. By default, this will be set to Service.
 */
export type LogonType = (typeof LogonType)[keyof typeof LogonType];

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

export const OsDiskStorageAccountType = {
    /**
     * Standard OS disk type.
     */
    Standard: "Standard",
    /**
     * Premium OS disk type.
     */
    Premium: "Premium",
    /**
     * Standard SSD OS disk type.
     */
    StandardSSD: "StandardSSD",
} as const;

/**
 * The Azure SKU name of the machines in the pool.
 */
export type OsDiskStorageAccountType = (typeof OsDiskStorageAccountType)[keyof typeof OsDiskStorageAccountType];

export const ProvisioningState = {
    /**
     * Represents a succeeded operation.
     */
    Succeeded: "Succeeded",
    /**
     * Represents a failed operation.
     */
    Failed: "Failed",
    /**
     * Represents a canceled operation.
     */
    Canceled: "Canceled",
    /**
     * Represents a pending operation.
     */
    Provisioning: "Provisioning",
    /**
     * Represents a pending operation.
     */
    Updating: "Updating",
    /**
     * Represents an operation under deletion.
     */
    Deleting: "Deleting",
    /**
     * Represents an accepted operation.
     */
    Accepted: "Accepted",
} as const;

/**
 * The status of the current operation.
 */
export type ProvisioningState = (typeof ProvisioningState)[keyof typeof ProvisioningState];

export const StorageAccountType = {
    /**
     * The data disk should use standard locally redundant storage.
     */
    StandardLRS: "standard_lrs",
    /**
     * The data disk should use premium locally redundant storage.
     */
    PremiumLRS: "premium_lrs",
    /**
     * The data disk should use standard SSD locally redundant storage.
     */
    StandardSSDLRS: "standardssd_lrs",
    /**
     * The data disk should use premium SSD zonal redundant storage.
     */
    PremiumZRS: "premium_zrs",
    /**
     * The data disk should use standard SSD zonal redundant storage.
     */
    StandardSSDZRS: "standardssd_zrs",
} as const;

/**
 * The storage Account type to be used for the data disk. If omitted, the default is "standard_lrs".
 */
export type StorageAccountType = (typeof StorageAccountType)[keyof typeof StorageAccountType];
