// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Response for list of user's role for Logz.io account.
 * Azure REST API version: 2022-01-01-preview.
 */
export function listMonitorUserRoles(args: ListMonitorUserRolesArgs, opts?: pulumi.InvokeOptions): Promise<ListMonitorUserRolesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:logz:listMonitorUserRoles", {
        "emailAddress": args.emailAddress,
        "monitorName": args.monitorName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListMonitorUserRolesArgs {
    /**
     * Email of the user used by Logz for contacting them if needed
     */
    emailAddress?: string;
    /**
     * Monitor resource name
     */
    monitorName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Response for list of user's role for Logz.io account.
 */
export interface ListMonitorUserRolesResult {
    /**
     * Link to the next set of results, if any.
     */
    readonly nextLink?: string;
    /**
     * List of user roles for Logz.io account.
     */
    readonly value?: outputs.logz.UserRoleResponseResponse[];
}
/**
 * Response for list of user's role for Logz.io account.
 * Azure REST API version: 2022-01-01-preview.
 */
export function listMonitorUserRolesOutput(args: ListMonitorUserRolesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListMonitorUserRolesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:logz:listMonitorUserRoles", {
        "emailAddress": args.emailAddress,
        "monitorName": args.monitorName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListMonitorUserRolesOutputArgs {
    /**
     * Email of the user used by Logz for contacting them if needed
     */
    emailAddress?: pulumi.Input<string>;
    /**
     * Monitor resource name
     */
    monitorName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
