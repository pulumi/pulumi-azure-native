// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This operation lists all the policy set definition versions for all policy set definitions within a subscription.
 *
 * Uses Azure REST API version 2023-04-01.
 *
 * Other available API versions: 2024-05-01, 2025-01-01, 2025-03-01.
 */
export function listPolicySetDefinitionVersionAll(args?: ListPolicySetDefinitionVersionAllArgs, opts?: pulumi.InvokeOptions): Promise<ListPolicySetDefinitionVersionAllResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:authorization:listPolicySetDefinitionVersionAll", {
    }, opts);
}

export interface ListPolicySetDefinitionVersionAllArgs {
}

/**
 * List of policy set definition versions.
 */
export interface ListPolicySetDefinitionVersionAllResult {
    /**
     * The URL to use for getting the next set of results.
     */
    readonly nextLink?: string;
    /**
     * An array of policy set definition versions.
     */
    readonly value?: outputs.authorization.PolicySetDefinitionVersionResponse[];
}
/**
 * This operation lists all the policy set definition versions for all policy set definitions within a subscription.
 *
 * Uses Azure REST API version 2023-04-01.
 *
 * Other available API versions: 2024-05-01, 2025-01-01, 2025-03-01.
 */
export function listPolicySetDefinitionVersionAllOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListPolicySetDefinitionVersionAllResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:authorization:listPolicySetDefinitionVersionAll", {
    }, opts);
}

