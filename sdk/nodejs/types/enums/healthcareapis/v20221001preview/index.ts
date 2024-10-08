// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AnalyticsConnectorDataDestinationType = {
    Datalake: "datalake",
} as const;

/**
 * Type of data destination.
 */
export type AnalyticsConnectorDataDestinationType = (typeof AnalyticsConnectorDataDestinationType)[keyof typeof AnalyticsConnectorDataDestinationType];

export const AnalyticsConnectorDataSourceType = {
    Fhirservice: "fhirservice",
} as const;

/**
 * Type of data source.
 */
export type AnalyticsConnectorDataSourceType = (typeof AnalyticsConnectorDataSourceType)[keyof typeof AnalyticsConnectorDataSourceType];

export const AnalyticsConnectorMappingType = {
    FhirToParquet: "fhirToParquet",
} as const;

/**
 * Type of data mapping.
 */
export type AnalyticsConnectorMappingType = (typeof AnalyticsConnectorMappingType)[keyof typeof AnalyticsConnectorMappingType];

export const FhirServiceVersion = {
    STU3: "STU3",
    R4: "R4",
} as const;

/**
 * The kind of FHIR Service.
 */
export type FhirServiceVersion = (typeof FhirServiceVersion)[keyof typeof FhirServiceVersion];

export const ServiceManagedIdentityType = {
    None: "None",
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned,UserAssigned",
} as const;

/**
 * Type of identity being specified, currently SystemAssigned and None are allowed.
 */
export type ServiceManagedIdentityType = (typeof ServiceManagedIdentityType)[keyof typeof ServiceManagedIdentityType];
