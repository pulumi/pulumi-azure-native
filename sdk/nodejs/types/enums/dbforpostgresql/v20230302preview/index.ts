// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ActiveDirectoryAuth = {
    Enabled: "enabled",
    Disabled: "disabled",
} as const;

export type ActiveDirectoryAuth = (typeof ActiveDirectoryAuth)[keyof typeof ActiveDirectoryAuth];

export const DataEncryptionType = {
    AzureKeyVault: "AzureKeyVault",
    SystemAssigned: "SystemAssigned",
} as const;

export type DataEncryptionType = (typeof DataEncryptionType)[keyof typeof DataEncryptionType];

export const IdentityType = {
    UserAssigned: "UserAssigned",
    SystemAssigned: "SystemAssigned",
} as const;

export type IdentityType = (typeof IdentityType)[keyof typeof IdentityType];

export const PasswordAuth = {
    Enabled: "enabled",
    Disabled: "disabled",
} as const;

export type PasswordAuth = (typeof PasswordAuth)[keyof typeof PasswordAuth];

export const PrincipalType = {
    User: "user",
    ServicePrincipal: "servicePrincipal",
    Group: "group",
} as const;

export type PrincipalType = (typeof PrincipalType)[keyof typeof PrincipalType];

export const PrivateEndpointServiceConnectionStatus = {
    Pending: "Pending",
    Approved: "Approved",
    Rejected: "Rejected",
} as const;

/**
 * Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
 */
export type PrivateEndpointServiceConnectionStatus = (typeof PrivateEndpointServiceConnectionStatus)[keyof typeof PrivateEndpointServiceConnectionStatus];

export const RoleType = {
    User: "user",
    Admin: "admin",
} as const;

export type RoleType = (typeof RoleType)[keyof typeof RoleType];
