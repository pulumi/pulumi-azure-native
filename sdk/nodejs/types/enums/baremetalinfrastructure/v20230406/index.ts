// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ProvisioningState = {
    Accepted: "Accepted",
    Creating: "Creating",
    Updating: "Updating",
    Failed: "Failed",
    Succeeded: "Succeeded",
    Deleting: "Deleting",
    Canceled: "Canceled",
    Migrating: "Migrating",
} as const;

/**
 * State of provisioning of the AzureBareMetalStorageInstance
 */
export type ProvisioningState = (typeof ProvisioningState)[keyof typeof ProvisioningState];
