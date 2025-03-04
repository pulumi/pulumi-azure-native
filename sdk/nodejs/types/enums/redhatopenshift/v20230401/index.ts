// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EncryptionAtHost = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Whether master virtual machines are encrypted at host.
 */
export type EncryptionAtHost = (typeof EncryptionAtHost)[keyof typeof EncryptionAtHost];

export const FipsValidatedModules = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * If FIPS validated crypto modules are used
 */
export type FipsValidatedModules = (typeof FipsValidatedModules)[keyof typeof FipsValidatedModules];

export const OutboundType = {
    Loadbalancer: "Loadbalancer",
    UserDefinedRouting: "UserDefinedRouting",
} as const;

/**
 * The OutboundType used for egress traffic.
 */
export type OutboundType = (typeof OutboundType)[keyof typeof OutboundType];

export const Visibility = {
    Private: "Private",
    Public: "Public",
} as const;

/**
 * Ingress visibility.
 */
export type Visibility = (typeof Visibility)[keyof typeof Visibility];
