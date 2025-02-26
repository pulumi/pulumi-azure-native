// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DataBoxEdgeDeviceKind = {
    AzureDataBoxGateway: "AzureDataBoxGateway",
    AzureStackEdge: "AzureStackEdge",
    AzureStackHub: "AzureStackHub",
    AzureModularDataCentre: "AzureModularDataCentre",
} as const;

/**
 * The kind of the device.
 */
export type DataBoxEdgeDeviceKind = (typeof DataBoxEdgeDeviceKind)[keyof typeof DataBoxEdgeDeviceKind];

export const DataBoxEdgeDeviceStatus = {
    ReadyToSetup: "ReadyToSetup",
    Online: "Online",
    Offline: "Offline",
    NeedsAttention: "NeedsAttention",
    Disconnected: "Disconnected",
    PartiallyDisconnected: "PartiallyDisconnected",
    Maintenance: "Maintenance",
} as const;

/**
 * The status of the Data Box Edge/Gateway device.
 */
export type DataBoxEdgeDeviceStatus = (typeof DataBoxEdgeDeviceStatus)[keyof typeof DataBoxEdgeDeviceStatus];

export const DataResidencyType = {
    GeoZoneReplication: "GeoZoneReplication",
    ZoneReplication: "ZoneReplication",
} as const;

/**
 * DataResidencyType enum
 */
export type DataResidencyType = (typeof DataResidencyType)[keyof typeof DataResidencyType];

export const MsiIdentityType = {
    None: "None",
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
} as const;

/**
 * Identity type
 */
export type MsiIdentityType = (typeof MsiIdentityType)[keyof typeof MsiIdentityType];

export const SkuName = {
    Gateway: "Gateway",
    Edge: "Edge",
    TEA_1Node: "TEA_1Node",
    TEA_1Node_UPS: "TEA_1Node_UPS",
    TEA_1Node_Heater: "TEA_1Node_Heater",
    TEA_1Node_UPS_Heater: "TEA_1Node_UPS_Heater",
    TEA_4Node_Heater: "TEA_4Node_Heater",
    TEA_4Node_UPS_Heater: "TEA_4Node_UPS_Heater",
    TMA: "TMA",
    TDC: "TDC",
    TCA_Small: "TCA_Small",
    GPU: "GPU",
    TCA_Large: "TCA_Large",
    EdgeP_Base: "EdgeP_Base",
    EdgeP_High: "EdgeP_High",
    EdgePR_Base: "EdgePR_Base",
    EdgePR_Base_UPS: "EdgePR_Base_UPS",
    EP2_64_1VPU_W: "EP2_64_1VPU_W",
    EP2_128_1T4_Mx1_W: "EP2_128_1T4_Mx1_W",
    EP2_256_2T4_W: "EP2_256_2T4_W",
    EdgeMR_Mini: "EdgeMR_Mini",
    RCA_Small: "RCA_Small",
    RCA_Large: "RCA_Large",
    RDC: "RDC",
    Management: "Management",
} as const;

/**
 * SKU name.
 */
export type SkuName = (typeof SkuName)[keyof typeof SkuName];

export const SkuTier = {
    Standard: "Standard",
} as const;

/**
 * The SKU tier. This is based on the SKU name.
 */
export type SkuTier = (typeof SkuTier)[keyof typeof SkuTier];
