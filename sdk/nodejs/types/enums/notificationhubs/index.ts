// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20170401 from "./v20170401";
import * as v20230101preview from "./v20230101preview";
import * as v20230901 from "./v20230901";
import * as v20231001preview from "./v20231001preview";

export {
    v20170401,
    v20230101preview,
    v20230901,
    v20231001preview,
};

export const AccessRights = {
    Manage: "Manage",
    Send: "Send",
    Listen: "Listen",
} as const;

/**
 * Defines values for AccessRights.
 */
export type AccessRights = (typeof AccessRights)[keyof typeof AccessRights];

export const NamespaceStatus = {
    Created: "Created",
    Creating: "Creating",
    Suspended: "Suspended",
    Deleting: "Deleting",
} as const;

/**
 * Namespace status.
 */
export type NamespaceStatus = (typeof NamespaceStatus)[keyof typeof NamespaceStatus];

export const NamespaceType = {
    Messaging: "Messaging",
    NotificationHub: "NotificationHub",
} as const;

/**
 * Defines values for NamespaceType.
 */
export type NamespaceType = (typeof NamespaceType)[keyof typeof NamespaceType];

export const OperationProvisioningState = {
    Unknown: "Unknown",
    InProgress: "InProgress",
    Succeeded: "Succeeded",
    Failed: "Failed",
    Canceled: "Canceled",
    Pending: "Pending",
    Disabled: "Disabled",
} as const;

/**
 * Defines values for OperationProvisioningState.
 */
export type OperationProvisioningState = (typeof OperationProvisioningState)[keyof typeof OperationProvisioningState];

export const PrivateEndpointConnectionProvisioningState = {
    Unknown: "Unknown",
    Succeeded: "Succeeded",
    Creating: "Creating",
    Updating: "Updating",
    UpdatingByProxy: "UpdatingByProxy",
    Deleting: "Deleting",
    DeletingByProxy: "DeletingByProxy",
    Deleted: "Deleted",
} as const;

/**
 * State of Private Endpoint Connection.
 */
export type PrivateEndpointConnectionProvisioningState = (typeof PrivateEndpointConnectionProvisioningState)[keyof typeof PrivateEndpointConnectionProvisioningState];

export const PrivateLinkConnectionStatus = {
    Disconnected: "Disconnected",
    Pending: "Pending",
    Approved: "Approved",
    Rejected: "Rejected",
} as const;

/**
 * State of Private Link Connection.
 */
export type PrivateLinkConnectionStatus = (typeof PrivateLinkConnectionStatus)[keyof typeof PrivateLinkConnectionStatus];

export const PublicNetworkAccess = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Type of public network access.
 */
export type PublicNetworkAccess = (typeof PublicNetworkAccess)[keyof typeof PublicNetworkAccess];

export const ReplicationRegion = {
    Default: "Default",
    WestUs2: "WestUs2",
    NorthEurope: "NorthEurope",
    AustraliaEast: "AustraliaEast",
    BrazilSouth: "BrazilSouth",
    SouthEastAsia: "SouthEastAsia",
    SouthAfricaNorth: "SouthAfricaNorth",
    None: "None",
} as const;

/**
 * Allowed replication region
 */
export type ReplicationRegion = (typeof ReplicationRegion)[keyof typeof ReplicationRegion];

export const SkuName = {
    Free: "Free",
    Basic: "Basic",
    Standard: "Standard",
} as const;

/**
 * Namespace SKU name.
 */
export type SkuName = (typeof SkuName)[keyof typeof SkuName];

export const ZoneRedundancyPreference = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Namespace SKU name.
 */
export type ZoneRedundancyPreference = (typeof ZoneRedundancyPreference)[keyof typeof ZoneRedundancyPreference];
