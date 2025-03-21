// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DeploymentMode = {
    Incremental: "Incremental",
    Complete: "Complete",
} as const;

/**
 * The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode, resources are deployed without deleting existing resources that are not included in the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included in the template are deleted. Be careful when using Complete mode as you may unintentionally delete resources.
 */
export type DeploymentMode = (typeof DeploymentMode)[keyof typeof DeploymentMode];

export const ExpressionEvaluationOptionsScopeType = {
    NotSpecified: "NotSpecified",
    Outer: "Outer",
    Inner: "Inner",
} as const;

/**
 * The scope to be used for evaluation of parameters, variables and functions in a nested template.
 */
export type ExpressionEvaluationOptionsScopeType = (typeof ExpressionEvaluationOptionsScopeType)[keyof typeof ExpressionEvaluationOptionsScopeType];

export const ExtendedLocationType = {
    EdgeZone: "EdgeZone",
} as const;

/**
 * The extended location type.
 */
export type ExtendedLocationType = (typeof ExtendedLocationType)[keyof typeof ExtendedLocationType];

export const OnErrorDeploymentType = {
    LastSuccessful: "LastSuccessful",
    SpecificDeployment: "SpecificDeployment",
} as const;

/**
 * The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
 */
export type OnErrorDeploymentType = (typeof OnErrorDeploymentType)[keyof typeof OnErrorDeploymentType];

export const ResourceIdentityType = {
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned, UserAssigned",
    None: "None",
} as const;

/**
 * The identity type.
 */
export type ResourceIdentityType = (typeof ResourceIdentityType)[keyof typeof ResourceIdentityType];

export const ValidationLevel = {
    /**
     * Static analysis of the template is performed.
     */
    Template: "Template",
    /**
     * Static analysis of the template is performed and resource declarations are sent to resource providers for semantic validation. Validates that the caller has RBAC write permissions on each resource.
     */
    Provider: "Provider",
    /**
     * Static analysis of the template is performed and resource declarations are sent to resource providers for semantic validation. Skips validating that the caller has RBAC write permissions on each resource.
     */
    ProviderNoRbac: "ProviderNoRbac",
} as const;

/**
 * The validation level of the deployment
 */
export type ValidationLevel = (typeof ValidationLevel)[keyof typeof ValidationLevel];
