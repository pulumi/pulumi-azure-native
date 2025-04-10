// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets all deployments for a remediation at management group scope.
 *
 * Uses Azure REST API version 2024-10-01.
 *
 * Other available API versions: 2021-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native policyinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listRemediationDeploymentsAtManagementGroup(args: ListRemediationDeploymentsAtManagementGroupArgs, opts?: pulumi.InvokeOptions): Promise<ListRemediationDeploymentsAtManagementGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:policyinsights:listRemediationDeploymentsAtManagementGroup", {
        "managementGroupId": args.managementGroupId,
        "managementGroupsNamespace": args.managementGroupsNamespace,
        "remediationName": args.remediationName,
        "top": args.top,
    }, opts);
}

export interface ListRemediationDeploymentsAtManagementGroupArgs {
    /**
     * Management group ID.
     */
    managementGroupId: string;
    /**
     * The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
     */
    managementGroupsNamespace: string;
    /**
     * The name of the remediation.
     */
    remediationName: string;
    /**
     * Maximum number of records to return.
     */
    top?: number;
}

/**
 * List of deployments for a remediation.
 */
export interface ListRemediationDeploymentsAtManagementGroupResult {
    /**
     * The URL to get the next set of results.
     */
    readonly nextLink: string;
    /**
     * Array of deployments for the remediation.
     */
    readonly value: outputs.policyinsights.RemediationDeploymentResponse[];
}
/**
 * Gets all deployments for a remediation at management group scope.
 *
 * Uses Azure REST API version 2024-10-01.
 *
 * Other available API versions: 2021-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native policyinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listRemediationDeploymentsAtManagementGroupOutput(args: ListRemediationDeploymentsAtManagementGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListRemediationDeploymentsAtManagementGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:policyinsights:listRemediationDeploymentsAtManagementGroup", {
        "managementGroupId": args.managementGroupId,
        "managementGroupsNamespace": args.managementGroupsNamespace,
        "remediationName": args.remediationName,
        "top": args.top,
    }, opts);
}

export interface ListRemediationDeploymentsAtManagementGroupOutputArgs {
    /**
     * Management group ID.
     */
    managementGroupId: pulumi.Input<string>;
    /**
     * The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
     */
    managementGroupsNamespace: pulumi.Input<string>;
    /**
     * The name of the remediation.
     */
    remediationName: pulumi.Input<string>;
    /**
     * Maximum number of records to return.
     */
    top?: pulumi.Input<number>;
}
