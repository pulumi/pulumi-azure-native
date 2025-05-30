// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a specific security operator for the requested scope.
 *
 * Uses Azure REST API version 2023-01-01-preview.
 */
export function getSecurityOperator(args: GetSecurityOperatorArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityOperatorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:security:getSecurityOperator", {
        "pricingName": args.pricingName,
        "securityOperatorName": args.securityOperatorName,
    }, opts);
}

export interface GetSecurityOperatorArgs {
    /**
     * name of the pricing configuration
     */
    pricingName: string;
    /**
     * name of the securityOperator
     */
    securityOperatorName: string;
}

/**
 * Security operator under a given subscription and pricing
 */
export interface GetSecurityOperatorResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Identity for the resource.
     */
    readonly identity?: outputs.security.IdentityResponse;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Resource type
     */
    readonly type: string;
}
/**
 * Get a specific security operator for the requested scope.
 *
 * Uses Azure REST API version 2023-01-01-preview.
 */
export function getSecurityOperatorOutput(args: GetSecurityOperatorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSecurityOperatorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:security:getSecurityOperator", {
        "pricingName": args.pricingName,
        "securityOperatorName": args.securityOperatorName,
    }, opts);
}

export interface GetSecurityOperatorOutputArgs {
    /**
     * name of the pricing configuration
     */
    pricingName: pulumi.Input<string>;
    /**
     * name of the securityOperator
     */
    securityOperatorName: pulumi.Input<string>;
}
