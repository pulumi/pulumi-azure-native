// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ProvisioningState = {
    Unknown: "Unknown",
    Succeeded: "Succeeded",
    Failed: "Failed",
    Canceled: "Canceled",
    Accepted: "Accepted",
    Provisioning: "Provisioning",
    Deleting: "Deleting",
} as const;

/**
 * The resource provisioning state.
 */
export type ProvisioningState = (typeof ProvisioningState)[keyof typeof ProvisioningState];
