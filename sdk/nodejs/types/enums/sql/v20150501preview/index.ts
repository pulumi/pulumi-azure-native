// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ServerKeyType = {
    ServiceManaged: "ServiceManaged",
    AzureKeyVault: "AzureKeyVault",
} as const;

/**
 * The server key type like 'ServiceManaged', 'AzureKeyVault'.
 */
export type ServerKeyType = (typeof ServerKeyType)[keyof typeof ServerKeyType];
