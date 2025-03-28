// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20200815preview from "./v20200815preview";
import * as v20220510preview from "./v20220510preview";
import * as v20221227 from "./v20221227";
import * as v20230620preview from "./v20230620preview";
import * as v20231003preview from "./v20231003preview";
import * as v20240331preview from "./v20240331preview";
import * as v20240520preview from "./v20240520preview";
import * as v20240710 from "./v20240710";
import * as v20240731preview from "./v20240731preview";
import * as v20240910preview from "./v20240910preview";
import * as v20241110preview from "./v20241110preview";
import * as v20250113 from "./v20250113";

export {
    v20200815preview,
    v20220510preview,
    v20221227,
    v20230620preview,
    v20231003preview,
    v20240331preview,
    v20240520preview,
    v20240710,
    v20240731preview,
    v20240910preview,
    v20241110preview,
    v20250113,
};

export const AssessmentModeTypes = {
    ImageDefault: "ImageDefault",
    AutomaticByPlatform: "AutomaticByPlatform",
} as const;

/**
 * Specifies the assessment mode.
 */
export type AssessmentModeTypes = (typeof AssessmentModeTypes)[keyof typeof AssessmentModeTypes];

export const GatewayType = {
    Public: "Public",
} as const;

/**
 * The type of the Gateway resource.
 */
export type GatewayType = (typeof GatewayType)[keyof typeof GatewayType];

export const LicenseCoreType = {
    PCore: "pCore",
    VCore: "vCore",
} as const;

/**
 * Describes the license core type (pCore or vCore).
 */
export type LicenseCoreType = (typeof LicenseCoreType)[keyof typeof LicenseCoreType];

export const LicenseEdition = {
    Standard: "Standard",
    Datacenter: "Datacenter",
} as const;

/**
 * Describes the edition of the license. The values are either Standard or Datacenter.
 */
export type LicenseEdition = (typeof LicenseEdition)[keyof typeof LicenseEdition];

export const LicenseState = {
    Activated: "Activated",
    Deactivated: "Deactivated",
} as const;

/**
 * Describes the state of the license.
 */
export type LicenseState = (typeof LicenseState)[keyof typeof LicenseState];

export const LicenseTarget = {
    Windows_Server_2012: "Windows Server 2012",
    Windows_Server_2012_R2: "Windows Server 2012 R2",
} as const;

/**
 * Describes the license target server.
 */
export type LicenseTarget = (typeof LicenseTarget)[keyof typeof LicenseTarget];

export const LicenseType = {
    ESU: "ESU",
} as const;

/**
 * The type of the license resource.
 */
export type LicenseType = (typeof LicenseType)[keyof typeof LicenseType];

export const PatchModeTypes = {
    ImageDefault: "ImageDefault",
    AutomaticByPlatform: "AutomaticByPlatform",
    AutomaticByOS: "AutomaticByOS",
    Manual: "Manual",
} as const;

/**
 * Specifies the patch mode.
 */
export type PatchModeTypes = (typeof PatchModeTypes)[keyof typeof PatchModeTypes];

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

export const StatusLevelTypes = {
    Info: "Info",
    Warning: "Warning",
    Error: "Error",
} as const;

/**
 * The level code.
 */
export type StatusLevelTypes = (typeof StatusLevelTypes)[keyof typeof StatusLevelTypes];
