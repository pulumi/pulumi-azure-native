// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const StandardSupportedClouds = {
    AWS: "AWS",
    GCP: "GCP",
} as const;

/**
 * The cloud that the standard is supported on.
 */
export type StandardSupportedClouds = (typeof StandardSupportedClouds)[keyof typeof StandardSupportedClouds];
