// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the details of a replication policy.
 */
export function getReplicationPolicy(args: GetReplicationPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:recoveryservices/v20241001:getReplicationPolicy", {
        "policyName": args.policyName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetReplicationPolicyArgs {
    /**
     * Replication policy name.
     */
    policyName: string;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    resourceGroupName: string;
    /**
     * The name of the recovery services vault.
     */
    resourceName: string;
}

/**
 * Protection profile details.
 */
export interface GetReplicationPolicyResult {
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Resource Location
     */
    readonly location?: string;
    /**
     * Resource Name
     */
    readonly name: string;
    /**
     * The custom data.
     */
    readonly properties: outputs.recoveryservices.v20241001.PolicyPropertiesResponse;
    /**
     * Resource Type
     */
    readonly type: string;
}
/**
 * Gets the details of a replication policy.
 */
export function getReplicationPolicyOutput(args: GetReplicationPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetReplicationPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:recoveryservices/v20241001:getReplicationPolicy", {
        "policyName": args.policyName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetReplicationPolicyOutputArgs {
    /**
     * Replication policy name.
     */
    policyName: pulumi.Input<string>;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the recovery services vault.
     */
    resourceName: pulumi.Input<string>;
}
