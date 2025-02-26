// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20171201 from "./v20171201";
import * as v20180601privatepreview from "./v20180601privatepreview";
import * as v20200701preview from "./v20200701preview";
import * as v20200701privatepreview from "./v20200701privatepreview";
import * as v20220101 from "./v20220101";
import * as v20220930preview from "./v20220930preview";
import * as v20230601preview from "./v20230601preview";
import * as v20230630 from "./v20230630";
import * as v20231001preview from "./v20231001preview";
import * as v20231201preview from "./v20231201preview";
import * as v20231230 from "./v20231230";
import * as v20240201preview from "./v20240201preview";
import * as v20240601preview from "./v20240601preview";
import * as v20241001preview from "./v20241001preview";

export {
    v20171201,
    v20180601privatepreview,
    v20200701preview,
    v20200701privatepreview,
    v20220101,
    v20220930preview,
    v20230601preview,
    v20230630,
    v20231001preview,
    v20231201preview,
    v20231230,
    v20240201preview,
    v20240601preview,
    v20241001preview,
};

export const AdministratorType = {
    ActiveDirectory: "ActiveDirectory",
} as const;

/**
 * Type of the sever administrator.
 */
export type AdministratorType = (typeof AdministratorType)[keyof typeof AdministratorType];

export const ConfigurationSource = {
    System_default: "system-default",
    User_override: "user-override",
} as const;

/**
 * Source of the configuration.
 */
export type ConfigurationSource = (typeof ConfigurationSource)[keyof typeof ConfigurationSource];

export const CreateMode = {
    Default: "Default",
    PointInTimeRestore: "PointInTimeRestore",
    Replica: "Replica",
    GeoRestore: "GeoRestore",
} as const;

/**
 * The mode to create a new MySQL server.
 */
export type CreateMode = (typeof CreateMode)[keyof typeof CreateMode];

export const DataEncryptionType = {
    AzureKeyVault: "AzureKeyVault",
    SystemManaged: "SystemManaged",
} as const;

/**
 * The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
 */
export type DataEncryptionType = (typeof DataEncryptionType)[keyof typeof DataEncryptionType];

export const EnableStatusEnum = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Enable Log On Disk or not.
 */
export type EnableStatusEnum = (typeof EnableStatusEnum)[keyof typeof EnableStatusEnum];

export const HighAvailabilityMode = {
    Disabled: "Disabled",
    ZoneRedundant: "ZoneRedundant",
    SameZone: "SameZone",
} as const;

/**
 * High availability mode for a server.
 */
export type HighAvailabilityMode = (typeof HighAvailabilityMode)[keyof typeof HighAvailabilityMode];

export const ManagedServiceIdentityType = {
    UserAssigned: "UserAssigned",
} as const;

/**
 * Type of managed service identity.
 */
export type ManagedServiceIdentityType = (typeof ManagedServiceIdentityType)[keyof typeof ManagedServiceIdentityType];

export const PrivateEndpointServiceConnectionStatus = {
    Pending: "Pending",
    Approved: "Approved",
    Rejected: "Rejected",
} as const;

/**
 * Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
 */
export type PrivateEndpointServiceConnectionStatus = (typeof PrivateEndpointServiceConnectionStatus)[keyof typeof PrivateEndpointServiceConnectionStatus];

export const ReplicationRole = {
    None: "None",
    Source: "Source",
    Replica: "Replica",
} as const;

/**
 * The replication role.
 */
export type ReplicationRole = (typeof ReplicationRole)[keyof typeof ReplicationRole];

export const ServerVersion = {
    ServerVersion_5_7: "5.7",
    ServerVersion_8_0_21: "8.0.21",
} as const;

/**
 * Server version.
 */
export type ServerVersion = (typeof ServerVersion)[keyof typeof ServerVersion];

export const SkuTier = {
    Burstable: "Burstable",
    GeneralPurpose: "GeneralPurpose",
    MemoryOptimized: "MemoryOptimized",
} as const;

/**
 * The tier of the particular SKU, e.g. GeneralPurpose.
 */
export type SkuTier = (typeof SkuTier)[keyof typeof SkuTier];
