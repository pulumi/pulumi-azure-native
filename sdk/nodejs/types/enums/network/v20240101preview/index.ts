// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AddressPrefixType = {
    IPPrefix: "IPPrefix",
    ServiceTag: "ServiceTag",
    NetworkGroup: "NetworkGroup",
} as const;

/**
 * Address prefix type.
 */
export type AddressPrefixType = (typeof AddressPrefixType)[keyof typeof AddressPrefixType];

export const AddressSpaceAggregationOption = {
    None: "None",
    Manual: "Manual",
} as const;

/**
 * Determine update behavior for changes to network groups referenced within the rules in this configuration.
 */
export type AddressSpaceAggregationOption = (typeof AddressSpaceAggregationOption)[keyof typeof AddressSpaceAggregationOption];

export const AdminRuleKind = {
    Custom: "Custom",
    Default: "Default",
} as const;

/**
 * Whether the rule is custom or default.
 */
export type AdminRuleKind = (typeof AdminRuleKind)[keyof typeof AdminRuleKind];

export const ConfigurationType = {
    SecurityAdmin: "SecurityAdmin",
    Connectivity: "Connectivity",
} as const;

/**
 * Configuration Deployment Type.
 */
export type ConfigurationType = (typeof ConfigurationType)[keyof typeof ConfigurationType];

export const NetworkIntentPolicyBasedService = {
    None: "None",
    All: "All",
    AllowRulesOnly: "AllowRulesOnly",
} as const;

/**
 * Network intent policy based services.
 */
export type NetworkIntentPolicyBasedService = (typeof NetworkIntentPolicyBasedService)[keyof typeof NetworkIntentPolicyBasedService];

export const NetworkProtocol = {
    Any: "Any",
    TCP: "TCP",
    UDP: "UDP",
    ICMP: "ICMP",
} as const;

/**
 * Network protocol this resource applies to.
 */
export type NetworkProtocol = (typeof NetworkProtocol)[keyof typeof NetworkProtocol];

export const SecurityConfigurationRuleAccess = {
    Allow: "Allow",
    Deny: "Deny",
    AlwaysAllow: "AlwaysAllow",
} as const;

/**
 * Indicates the access allowed for this particular rule
 */
export type SecurityConfigurationRuleAccess = (typeof SecurityConfigurationRuleAccess)[keyof typeof SecurityConfigurationRuleAccess];

export const SecurityConfigurationRuleDirection = {
    Inbound: "Inbound",
    Outbound: "Outbound",
} as const;

/**
 * Indicates if the traffic matched against the rule in inbound or outbound.
 */
export type SecurityConfigurationRuleDirection = (typeof SecurityConfigurationRuleDirection)[keyof typeof SecurityConfigurationRuleDirection];

export const SecurityConfigurationRuleProtocol = {
    Tcp: "Tcp",
    Udp: "Udp",
    Icmp: "Icmp",
    Esp: "Esp",
    Any: "Any",
    Ah: "Ah",
} as const;

/**
 * Network protocol this rule applies to.
 */
export type SecurityConfigurationRuleProtocol = (typeof SecurityConfigurationRuleProtocol)[keyof typeof SecurityConfigurationRuleProtocol];
