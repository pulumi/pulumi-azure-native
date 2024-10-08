// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20210815 from "./v20210815";
import * as v20210831preview from "./v20210831preview";

export {
    v20210815,
    v20210831preview,
};

export const HostType = {
    Kubernetes: "Kubernetes",
} as const;

/**
 * Type of host the Custom Locations is referencing (Kubernetes, etc...).
 */
export type HostType = (typeof HostType)[keyof typeof HostType];

export const ResourceIdentityType = {
    SystemAssigned: "SystemAssigned",
    None: "None",
} as const;

/**
 * The identity type.
 */
export type ResourceIdentityType = (typeof ResourceIdentityType)[keyof typeof ResourceIdentityType];
