// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a specific governanceAssignment for the requested scope by AssignmentKey
 *
 * Uses Azure REST API version 2022-01-01-preview.
 */
export function getGovernanceAssignment(args: GetGovernanceAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetGovernanceAssignmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:security:getGovernanceAssignment", {
        "assessmentName": args.assessmentName,
        "assignmentKey": args.assignmentKey,
        "scope": args.scope,
    }, opts);
}

export interface GetGovernanceAssignmentArgs {
    /**
     * The Assessment Key - A unique key for the assessment type
     */
    assessmentName: string;
    /**
     * The governance assignment key - the assessment key of the required governance assignment
     */
    assignmentKey: string;
    /**
     * The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
     */
    scope: string;
}

/**
 * Governance assignment over a given scope
 */
export interface GetGovernanceAssignmentResult {
    /**
     * The additional data for the governance assignment - e.g. links to ticket (optional), see example
     */
    readonly additionalData?: outputs.security.GovernanceAssignmentAdditionalDataResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
     */
    readonly governanceEmailNotification?: outputs.security.GovernanceEmailNotificationResponse;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Defines whether there is a grace period on the governance assignment
     */
    readonly isGracePeriod?: boolean;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The Owner for the governance assignment - e.g. user@contoso.com - see example
     */
    readonly owner?: string;
    /**
     * The remediation due-date - after this date Secure Score will be affected (in case of  active grace-period)
     */
    readonly remediationDueDate: string;
    /**
     * The ETA (estimated time of arrival) for remediation (optional), see example
     */
    readonly remediationEta?: outputs.security.RemediationEtaResponse;
    /**
     * Resource type
     */
    readonly type: string;
}
/**
 * Get a specific governanceAssignment for the requested scope by AssignmentKey
 *
 * Uses Azure REST API version 2022-01-01-preview.
 */
export function getGovernanceAssignmentOutput(args: GetGovernanceAssignmentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGovernanceAssignmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:security:getGovernanceAssignment", {
        "assessmentName": args.assessmentName,
        "assignmentKey": args.assignmentKey,
        "scope": args.scope,
    }, opts);
}

export interface GetGovernanceAssignmentOutputArgs {
    /**
     * The Assessment Key - A unique key for the assessment type
     */
    assessmentName: pulumi.Input<string>;
    /**
     * The governance assignment key - the assessment key of the required governance assignment
     */
    assignmentKey: pulumi.Input<string>;
    /**
     * The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
     */
    scope: pulumi.Input<string>;
}
