// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const KeyEncryptionKeyIdentityType = {
    /**
     * System assigned identity
     */
    SystemAssignedIdentity: "SystemAssignedIdentity",
    /**
     * User assigned identity
     */
    UserAssignedIdentity: "UserAssignedIdentity",
} as const;

/**
 * The type of identity to use. Values can be systemAssignedIdentity, userAssignedIdentity, or delegatedResourceIdentity.
 */
export type KeyEncryptionKeyIdentityType = (typeof KeyEncryptionKeyIdentityType)[keyof typeof KeyEncryptionKeyIdentityType];

export const ManagedServiceIdentityType = {
    None: "None",
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned,UserAssigned",
} as const;

/**
 * Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
 */
export type ManagedServiceIdentityType = (typeof ManagedServiceIdentityType)[keyof typeof ManagedServiceIdentityType];

export const OnlineExperimentationWorkspaceSkuName = {
    /**
     * The Free service sku name.
     */
    F0: "F0",
    /**
     * The Standard service sku name.
     */
    S0: "S0",
    /**
     * The Premium service sku name.
     */
    P0: "P0",
    /**
     * The Developer service sku name.
     */
    D0: "D0",
} as const;

/**
 * The name of the SKU. Ex - F0, P0. It is typically a letter+number code
 */
export type OnlineExperimentationWorkspaceSkuName = (typeof OnlineExperimentationWorkspaceSkuName)[keyof typeof OnlineExperimentationWorkspaceSkuName];
