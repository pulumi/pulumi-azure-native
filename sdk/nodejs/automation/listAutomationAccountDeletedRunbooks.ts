// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieve the deleted runbooks for an automation account.
 *
 * Uses Azure REST API version 2023-05-15-preview.
 *
 * Other available API versions: 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listAutomationAccountDeletedRunbooks(args: ListAutomationAccountDeletedRunbooksArgs, opts?: pulumi.InvokeOptions): Promise<ListAutomationAccountDeletedRunbooksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:automation:listAutomationAccountDeletedRunbooks", {
        "automationAccountName": args.automationAccountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListAutomationAccountDeletedRunbooksArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: string;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: string;
}

/**
 * The response model for the list deleted runbook.
 */
export interface ListAutomationAccountDeletedRunbooksResult {
    /**
     * Gets or sets the next link.
     */
    readonly nextLink?: string;
    /**
     * List of deleted runbooks in automation account.
     */
    readonly value?: outputs.automation.DeletedRunbookResponse[];
}
/**
 * Retrieve the deleted runbooks for an automation account.
 *
 * Uses Azure REST API version 2023-05-15-preview.
 *
 * Other available API versions: 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listAutomationAccountDeletedRunbooksOutput(args: ListAutomationAccountDeletedRunbooksOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListAutomationAccountDeletedRunbooksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:automation:listAutomationAccountDeletedRunbooks", {
        "automationAccountName": args.automationAccountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListAutomationAccountDeletedRunbooksOutputArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
