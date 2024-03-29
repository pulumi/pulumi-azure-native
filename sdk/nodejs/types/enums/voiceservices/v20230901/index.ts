// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApiBridgeActivationState = {
    /**
     * API Bridge is enabled
     */
    Enabled: "enabled",
    /**
     * API Bridge is disabled
     */
    Disabled: "disabled",
} as const;

/**
 * The activation state of the API Bridge for this Communications Gateway
 */
export type ApiBridgeActivationState = (typeof ApiBridgeActivationState)[keyof typeof ApiBridgeActivationState];

export const AutoGeneratedDomainNameLabelScope = {
    /**
     * Generated domain name label depends on resource name and tenant ID.
     */
    TenantReuse: "TenantReuse",
    /**
     * Generated domain name label depends on resource name, tenant ID and subscription ID.
     */
    SubscriptionReuse: "SubscriptionReuse",
    /**
     * Generated domain name label depends on resource name, tenant ID, subscription ID and resource group name.
     */
    ResourceGroupReuse: "ResourceGroupReuse",
    /**
     * Generated domain name label is always unique.
     */
    NoReuse: "NoReuse",
} as const;

/**
 * The scope at which the auto-generated domain name can be re-used
 */
export type AutoGeneratedDomainNameLabelScope = (typeof AutoGeneratedDomainNameLabelScope)[keyof typeof AutoGeneratedDomainNameLabelScope];

export const CommunicationsPlatform = {
    /**
     * Operator Connect
     */
    OperatorConnect: "OperatorConnect",
    /**
     * Teams Phone Mobile
     */
    TeamsPhoneMobile: "TeamsPhoneMobile",
    /**
     * Teams Direct Routing
     */
    TeamsDirectRouting: "TeamsDirectRouting",
} as const;

/**
 * Available platform types.
 */
export type CommunicationsPlatform = (typeof CommunicationsPlatform)[keyof typeof CommunicationsPlatform];

export const Connectivity = {
    /**
     * This deployment connects to the operator network using a Public IP address, e.g. when using MAPS
     */
    PublicAddress: "PublicAddress",
} as const;

/**
 * How to connect back to the operator network, e.g. MAPS
 */
export type Connectivity = (typeof Connectivity)[keyof typeof Connectivity];

export const E911Type = {
    /**
     * Emergency calls are not handled different from other calls
     */
    Standard: "Standard",
    /**
     * Emergency calls are routed directly to the ESRP
     */
    DirectToEsrp: "DirectToEsrp",
} as const;

/**
 * How to handle 911 calls
 */
export type E911Type = (typeof E911Type)[keyof typeof E911Type];

export const ManagedServiceIdentityType = {
    None: "None",
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned, UserAssigned",
} as const;

/**
 * Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
 */
export type ManagedServiceIdentityType = (typeof ManagedServiceIdentityType)[keyof typeof ManagedServiceIdentityType];

export const SkuTier = {
    Free: "Free",
    Basic: "Basic",
    Standard: "Standard",
    Premium: "Premium",
} as const;

/**
 * This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
 */
export type SkuTier = (typeof SkuTier)[keyof typeof SkuTier];

export const TeamsCodecs = {
    /**
     * Pulse code modulation(PCM) U-law narrowband audio codec(G.711u)
     */
    PCMA: "PCMA",
    /**
     * Pulse code modulation(PCM) U-law narrowband audio codec(G.711u)
     */
    PCMU: "PCMU",
    /**
     * G.722 wideband audio codec
     */
    G722: "G722",
    /**
     * G.722.2 wideband audio codec
     */
    G722_2: "G722_2",
    /**
     * SILK/8000 narrowband audio codec
     */
    SILK_8: "SILK_8",
    /**
     * SILK/16000 wideband audio codec
     */
    SILK_16: "SILK_16",
} as const;

/**
 * The voice codecs expected for communication with Teams.
 */
export type TeamsCodecs = (typeof TeamsCodecs)[keyof typeof TeamsCodecs];

export const TestLinePurpose = {
    /**
     * The test line is used for manual testing
     */
    Manual: "Manual",
    /**
     * The test line is used for automated testing
     */
    Automated: "Automated",
} as const;

/**
 * Purpose of this test line, e.g. automated or manual testing
 */
export type TestLinePurpose = (typeof TestLinePurpose)[keyof typeof TestLinePurpose];
