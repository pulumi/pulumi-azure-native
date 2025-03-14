// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AssociationType = {
    Subnets: "subnets",
} as const;

/**
 * Association Type
 */
export type AssociationType = (typeof AssociationType)[keyof typeof AssociationType];

export const FrontendIPAddressVersion = {
    IPv4: "IPv4",
    IPv6: "IPv6",
} as const;

/**
 * Frontend IP Address Version (Optional).
 */
export type FrontendIPAddressVersion = (typeof FrontendIPAddressVersion)[keyof typeof FrontendIPAddressVersion];

export const FrontendMode = {
    Public: "public",
} as const;

/**
 * Frontend Mode (Optional).
 */
export type FrontendMode = (typeof FrontendMode)[keyof typeof FrontendMode];
