// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get the callback url for a trigger of a workflow version.
 */
export function listWorkflowVersionTriggerCallbackUrl(args: ListWorkflowVersionTriggerCallbackUrlArgs, opts?: pulumi.InvokeOptions): Promise<ListWorkflowVersionTriggerCallbackUrlResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:logic/v20190501:listWorkflowVersionTriggerCallbackUrl", {
        "keyType": args.keyType,
        "notAfter": args.notAfter,
        "resourceGroupName": args.resourceGroupName,
        "triggerName": args.triggerName,
        "versionId": args.versionId,
        "workflowName": args.workflowName,
    }, opts);
}

export interface ListWorkflowVersionTriggerCallbackUrlArgs {
    /**
     * The key type.
     */
    keyType?: string | enums.logic.v20190501.KeyType;
    /**
     * The expiry time.
     */
    notAfter?: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
    /**
     * The workflow trigger name.
     */
    triggerName: string;
    /**
     * The workflow versionId.
     */
    versionId: string;
    /**
     * The workflow name.
     */
    workflowName: string;
}

/**
 * The workflow trigger callback URL.
 */
export interface ListWorkflowVersionTriggerCallbackUrlResult {
    /**
     * Gets the workflow trigger callback URL base path.
     */
    readonly basePath: string;
    /**
     * Gets the workflow trigger callback URL HTTP method.
     */
    readonly method: string;
    /**
     * Gets the workflow trigger callback URL query parameters.
     */
    readonly queries?: outputs.logic.v20190501.WorkflowTriggerListCallbackUrlQueriesResponse;
    /**
     * Gets the workflow trigger callback URL relative path.
     */
    readonly relativePath: string;
    /**
     * Gets the workflow trigger callback URL relative path parameters.
     */
    readonly relativePathParameters?: string[];
    /**
     * Gets the workflow trigger callback URL.
     */
    readonly value: string;
}
/**
 * Get the callback url for a trigger of a workflow version.
 */
export function listWorkflowVersionTriggerCallbackUrlOutput(args: ListWorkflowVersionTriggerCallbackUrlOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListWorkflowVersionTriggerCallbackUrlResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:logic/v20190501:listWorkflowVersionTriggerCallbackUrl", {
        "keyType": args.keyType,
        "notAfter": args.notAfter,
        "resourceGroupName": args.resourceGroupName,
        "triggerName": args.triggerName,
        "versionId": args.versionId,
        "workflowName": args.workflowName,
    }, opts);
}

export interface ListWorkflowVersionTriggerCallbackUrlOutputArgs {
    /**
     * The key type.
     */
    keyType?: pulumi.Input<string | enums.logic.v20190501.KeyType>;
    /**
     * The expiry time.
     */
    notAfter?: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The workflow trigger name.
     */
    triggerName: pulumi.Input<string>;
    /**
     * The workflow versionId.
     */
    versionId: pulumi.Input<string>;
    /**
     * The workflow name.
     */
    workflowName: pulumi.Input<string>;
}
