// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Get the verification id of a subscription used for verifying custom domains
 */
export function getCustomDomainVerificationId(args?: GetCustomDomainVerificationIdArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomDomainVerificationIdResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:app/v20250101:getCustomDomainVerificationId", {
    }, opts);
}

export interface GetCustomDomainVerificationIdArgs {
}

/**
 * Custom domain verification Id of a subscription
 */
export interface GetCustomDomainVerificationIdResult {
    readonly value?: string;
}
/**
 * Get the verification id of a subscription used for verifying custom domains
 */
export function getCustomDomainVerificationIdOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCustomDomainVerificationIdResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:app/v20250101:getCustomDomainVerificationId", {
    }, opts);
}

