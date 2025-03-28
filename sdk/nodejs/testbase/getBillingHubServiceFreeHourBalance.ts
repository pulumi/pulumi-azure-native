// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Uses Azure REST API version 2022-04-01-preview.
 *
 * Other available API versions: 2023-11-01-preview.
 */
export function getBillingHubServiceFreeHourBalance(args: GetBillingHubServiceFreeHourBalanceArgs, opts?: pulumi.InvokeOptions): Promise<GetBillingHubServiceFreeHourBalanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:testbase:getBillingHubServiceFreeHourBalance", {
        "resourceGroupName": args.resourceGroupName,
        "testBaseAccountName": args.testBaseAccountName,
    }, opts);
}

export interface GetBillingHubServiceFreeHourBalanceArgs {
    /**
     * The name of the resource group that contains the resource.
     */
    resourceGroupName: string;
    /**
     * The resource name of the Test Base Account.
     */
    testBaseAccountName: string;
}

export interface GetBillingHubServiceFreeHourBalanceResult {
    readonly incrementEntries?: outputs.testbase.BillingHubFreeHourIncrementEntryResponse[];
    readonly totalRemainingFreeHours?: number;
}
/**
 * Uses Azure REST API version 2022-04-01-preview.
 *
 * Other available API versions: 2023-11-01-preview.
 */
export function getBillingHubServiceFreeHourBalanceOutput(args: GetBillingHubServiceFreeHourBalanceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBillingHubServiceFreeHourBalanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:testbase:getBillingHubServiceFreeHourBalance", {
        "resourceGroupName": args.resourceGroupName,
        "testBaseAccountName": args.testBaseAccountName,
    }, opts);
}

export interface GetBillingHubServiceFreeHourBalanceOutputArgs {
    /**
     * The name of the resource group that contains the resource.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The resource name of the Test Base Account.
     */
    testBaseAccountName: pulumi.Input<string>;
}
