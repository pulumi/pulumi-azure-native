// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * This operation retrieves the policy definition in the given subscription with the given name.
 */
export function getPolicyDefinition(args: GetPolicyDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyDefinitionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:authorization/v20190601:getPolicyDefinition", {
        "policyDefinitionName": args.policyDefinitionName,
    }, opts);
}

export interface GetPolicyDefinitionArgs {
    /**
     * The name of the policy definition to get.
     */
    policyDefinitionName: string;
}

/**
 * The policy definition.
 */
export interface GetPolicyDefinitionResult {
    /**
     * The policy definition description.
     */
    readonly description?: string;
    /**
     * The display name of the policy definition.
     */
    readonly displayName?: string;
    /**
     * The ID of the policy definition.
     */
    readonly id: string;
    /**
     * The policy definition metadata.
     */
    readonly metadata?: any;
    /**
     * The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
     */
    readonly mode?: string;
    /**
     * The name of the policy definition.
     */
    readonly name: string;
    /**
     * Required if a parameter is used in policy rule.
     */
    readonly parameters?: any;
    /**
     * The policy rule.
     */
    readonly policyRule?: any;
    /**
     * The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
     */
    readonly policyType?: string;
    /**
     * The type of the resource (Microsoft.Authorization/policyDefinitions).
     */
    readonly type: string;
}
/**
 * This operation retrieves the policy definition in the given subscription with the given name.
 */
export function getPolicyDefinitionOutput(args: GetPolicyDefinitionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPolicyDefinitionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:authorization/v20190601:getPolicyDefinition", {
        "policyDefinitionName": args.policyDefinitionName,
    }, opts);
}

export interface GetPolicyDefinitionOutputArgs {
    /**
     * The name of the policy definition to get.
     */
    policyDefinitionName: pulumi.Input<string>;
}
