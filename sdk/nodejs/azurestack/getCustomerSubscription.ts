// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Returns the specified product.
 *
 * Uses Azure REST API version 2022-06-01.
 *
 * Other available API versions: 2020-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestack [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getCustomerSubscription(args: GetCustomerSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomerSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:azurestack:getCustomerSubscription", {
        "customerSubscriptionName": args.customerSubscriptionName,
        "registrationName": args.registrationName,
        "resourceGroup": args.resourceGroup,
    }, opts);
}

export interface GetCustomerSubscriptionArgs {
    /**
     * Name of the product.
     */
    customerSubscriptionName: string;
    /**
     * Name of the Azure Stack registration.
     */
    registrationName: string;
    /**
     * Name of the resource group.
     */
    resourceGroup: string;
}

/**
 * Customer subscription.
 */
export interface GetCustomerSubscriptionResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The entity tag used for optimistic concurrency when modifying the resource.
     */
    readonly etag?: string;
    /**
     * ID of the resource.
     */
    readonly id: string;
    /**
     * Name of the resource.
     */
    readonly name: string;
    /**
     * Tenant Id.
     */
    readonly tenantId?: string;
    /**
     * Type of Resource.
     */
    readonly type: string;
}
/**
 * Returns the specified product.
 *
 * Uses Azure REST API version 2022-06-01.
 *
 * Other available API versions: 2020-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestack [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getCustomerSubscriptionOutput(args: GetCustomerSubscriptionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCustomerSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:azurestack:getCustomerSubscription", {
        "customerSubscriptionName": args.customerSubscriptionName,
        "registrationName": args.registrationName,
        "resourceGroup": args.resourceGroup,
    }, opts);
}

export interface GetCustomerSubscriptionOutputArgs {
    /**
     * Name of the product.
     */
    customerSubscriptionName: pulumi.Input<string>;
    /**
     * Name of the Azure Stack registration.
     */
    registrationName: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    resourceGroup: pulumi.Input<string>;
}
