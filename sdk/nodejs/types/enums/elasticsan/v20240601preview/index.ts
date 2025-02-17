// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const Action = {
    Allow: "Allow",
} as const;

/**
 * The action of virtual network rule.
 */
export type Action = (typeof Action)[keyof typeof Action];

export const AutoScalePolicyEnforcement = {
    None: "None",
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Enable or Disable scale up setting on Elastic San Appliance.
 */
export type AutoScalePolicyEnforcement = (typeof AutoScalePolicyEnforcement)[keyof typeof AutoScalePolicyEnforcement];

export const EncryptionType = {
    /**
     * Volume is encrypted at rest with Platform managed key. It is the default encryption type.
     */
    EncryptionAtRestWithPlatformKey: "EncryptionAtRestWithPlatformKey",
    /**
     * Volume is encrypted at rest with Customer managed key that can be changed and revoked by a customer.
     */
    EncryptionAtRestWithCustomerManagedKey: "EncryptionAtRestWithCustomerManagedKey",
} as const;

/**
 * Type of encryption
 */
export type EncryptionType = (typeof EncryptionType)[keyof typeof EncryptionType];

export const IdentityType = {
    None: "None",
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
} as const;

/**
 * The identity type.
 */
export type IdentityType = (typeof IdentityType)[keyof typeof IdentityType];

export const PrivateEndpointServiceConnectionStatus = {
    Pending: "Pending",
    Approved: "Approved",
    Failed: "Failed",
    Rejected: "Rejected",
} as const;

/**
 * Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
 */
export type PrivateEndpointServiceConnectionStatus = (typeof PrivateEndpointServiceConnectionStatus)[keyof typeof PrivateEndpointServiceConnectionStatus];

export const PublicNetworkAccess = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Allow or disallow public network access to ElasticSan. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
 */
export type PublicNetworkAccess = (typeof PublicNetworkAccess)[keyof typeof PublicNetworkAccess];

export const SkuName = {
    /**
     * Premium locally redundant storage
     */
    Premium_LRS: "Premium_LRS",
    /**
     * Premium zone redundant storage
     */
    Premium_ZRS: "Premium_ZRS",
} as const;

/**
 * The sku name.
 */
export type SkuName = (typeof SkuName)[keyof typeof SkuName];

export const SkuTier = {
    /**
     * Premium Tier
     */
    Premium: "Premium",
} as const;

/**
 * The sku tier.
 */
export type SkuTier = (typeof SkuTier)[keyof typeof SkuTier];

export const StorageTargetType = {
    Iscsi: "Iscsi",
    None: "None",
} as const;

/**
 * Type of storage target
 */
export type StorageTargetType = (typeof StorageTargetType)[keyof typeof StorageTargetType];

export const VolumeCreateOption = {
    None: "None",
    VolumeSnapshot: "VolumeSnapshot",
    DiskSnapshot: "DiskSnapshot",
    Disk: "Disk",
    DiskRestorePoint: "DiskRestorePoint",
} as const;

/**
 * This enumerates the possible sources of a volume creation.
 */
export type VolumeCreateOption = (typeof VolumeCreateOption)[keyof typeof VolumeCreateOption];
