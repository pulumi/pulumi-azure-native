// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const IscsiTargetAclMode = {
    Dynamic: "Dynamic",
    Static: "Static",
} as const;

/**
 * Mode for Target connectivity.
 */
export type IscsiTargetAclMode = (typeof IscsiTargetAclMode)[keyof typeof IscsiTargetAclMode];
