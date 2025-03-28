// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20200701preview from "./v20200701preview";
import * as v20211101preview from "./v20211101preview";
import * as v20220101preview from "./v20220101preview";
import * as v20220402preview from "./v20220402preview";
import * as v20220701 from "./v20220701";
import * as v20230501 from "./v20230501";
import * as v20240401preview from "./v20240401preview";
import * as v20241101 from "./v20241101";
import * as v20241101preview from "./v20241101preview";

export {
    v20200701preview,
    v20211101preview,
    v20220101preview,
    v20220402preview,
    v20220701,
    v20230501,
    v20240401preview,
    v20241101,
    v20241101preview,
};

export const AKSIdentityType = {
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
} as const;

/**
 * The identity type.
 */
export type AKSIdentityType = (typeof AKSIdentityType)[keyof typeof AKSIdentityType];

export const LevelType = {
    Error: "Error",
    Warning: "Warning",
    Information: "Information",
} as const;

/**
 * Level of the status.
 */
export type LevelType = (typeof LevelType)[keyof typeof LevelType];

export const OperatorScopeType = {
    Cluster: "cluster",
    Namespace: "namespace",
} as const;

/**
 * Scope at which the operator will be installed.
 */
export type OperatorScopeType = (typeof OperatorScopeType)[keyof typeof OperatorScopeType];

export const OperatorType = {
    Flux: "Flux",
} as const;

/**
 * Type of the operator
 */
export type OperatorType = (typeof OperatorType)[keyof typeof OperatorType];

export const PrivateEndpointServiceConnectionStatus = {
    Pending: "Pending",
    Approved: "Approved",
    Rejected: "Rejected",
} as const;

/**
 * Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
 */
export type PrivateEndpointServiceConnectionStatus = (typeof PrivateEndpointServiceConnectionStatus)[keyof typeof PrivateEndpointServiceConnectionStatus];

export const PublicNetworkAccessType = {
    /**
     * Allows Azure Arc agents to communicate with Azure Arc services over both public (internet) and private endpoints.
     */
    Enabled: "Enabled",
    /**
     * Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet) endpoints. The agents must use the private link.
     */
    Disabled: "Disabled",
} as const;

/**
 * Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
 */
export type PublicNetworkAccessType = (typeof PublicNetworkAccessType)[keyof typeof PublicNetworkAccessType];

export const ResourceIdentityType = {
    SystemAssigned: "SystemAssigned",
} as const;

/**
 * The identity type.
 */
export type ResourceIdentityType = (typeof ResourceIdentityType)[keyof typeof ResourceIdentityType];

export const ScopeType = {
    Cluster: "cluster",
    Namespace: "namespace",
} as const;

/**
 * Scope at which the operator will be installed.
 */
export type ScopeType = (typeof ScopeType)[keyof typeof ScopeType];

export const SourceKindType = {
    GitRepository: "GitRepository",
    Bucket: "Bucket",
    AzureBlob: "AzureBlob",
} as const;

/**
 * Source Kind to pull the configuration data from.
 */
export type SourceKindType = (typeof SourceKindType)[keyof typeof SourceKindType];
