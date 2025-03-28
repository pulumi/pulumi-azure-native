// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * This operation lists all the policy definition versions for all policy definitions within a subscription.
 */
export function listPolicyDefinitionVersionAll(args?: ListPolicyDefinitionVersionAllArgs, opts?: pulumi.InvokeOptions): Promise<ListPolicyDefinitionVersionAllResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:authorization/v20250301:listPolicyDefinitionVersionAll", {
    }, opts);
}

export interface ListPolicyDefinitionVersionAllArgs {
}

/**
 * List of policy definition versions.
 */
export interface ListPolicyDefinitionVersionAllResult {
    /**
     * The URL to use for getting the next set of results.
     */
    readonly nextLink?: string;
    /**
     * An array of policy definitions versions.
     */
    readonly value?: outputs.authorization.v20250301.PolicyDefinitionVersionResponse[];
}
/**
 * This operation lists all the policy definition versions for all policy definitions within a subscription.
 */
export function listPolicyDefinitionVersionAllOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListPolicyDefinitionVersionAllResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:authorization/v20250301:listPolicyDefinitionVersionAll", {
    }, opts);
}

