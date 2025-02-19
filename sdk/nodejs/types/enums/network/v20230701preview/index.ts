// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AccessRuleDirection = {
    Inbound: "Inbound",
    Outbound: "Outbound",
} as const;

/**
 * Direction that specifies whether the access rules is inbound/outbound.
 */
export type AccessRuleDirection = (typeof AccessRuleDirection)[keyof typeof AccessRuleDirection];

export const ActionType = {
    Allow: "Allow",
    Alert: "Alert",
    Block: "Block",
} as const;

/**
 * The type of action to take.
 */
export type ActionType = (typeof ActionType)[keyof typeof ActionType];

export const AssociationAccessMode = {
    Learning: "Learning",
    Enforced: "Enforced",
    Audit: "Audit",
} as const;

/**
 * Access mode on the association.
 */
export type AssociationAccessMode = (typeof AssociationAccessMode)[keyof typeof AssociationAccessMode];

export const BlockResponseCode = {
    SERVFAIL: "SERVFAIL",
} as const;

/**
 * The response code for block actions.
 */
export type BlockResponseCode = (typeof BlockResponseCode)[keyof typeof BlockResponseCode];

export const DnsSecurityRuleState = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * The state of DNS security rule.
 */
export type DnsSecurityRuleState = (typeof DnsSecurityRuleState)[keyof typeof DnsSecurityRuleState];

export const ForwardingRuleState = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * The state of forwarding rule.
 */
export type ForwardingRuleState = (typeof ForwardingRuleState)[keyof typeof ForwardingRuleState];

export const IpAllocationMethod = {
    Static: "Static",
    Dynamic: "Dynamic",
} as const;

/**
 * Private IP address allocation method.
 */
export type IpAllocationMethod = (typeof IpAllocationMethod)[keyof typeof IpAllocationMethod];

export const ZoneType = {
    Public: "Public",
    Private: "Private",
} as const;

/**
 * The type of this DNS zone (Public or Private).
 */
export type ZoneType = (typeof ZoneType)[keyof typeof ZoneType];
