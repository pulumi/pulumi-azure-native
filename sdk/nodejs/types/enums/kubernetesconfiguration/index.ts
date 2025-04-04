// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


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
