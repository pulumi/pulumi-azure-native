// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets information about the specified DDoS custom policy.
 */
export function getDdosCustomPolicy(args: GetDdosCustomPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetDdosCustomPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network/v20231101:getDdosCustomPolicy", {
        "ddosCustomPolicyName": args.ddosCustomPolicyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDdosCustomPolicyArgs {
    /**
     * The name of the DDoS custom policy.
     */
    ddosCustomPolicyName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * A DDoS custom policy in a resource group.
 */
export interface GetDdosCustomPolicyResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the DDoS custom policy resource.
     */
    readonly provisioningState: string;
    /**
     * The resource GUID property of the DDoS custom policy resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
     */
    readonly resourceGuid: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Gets information about the specified DDoS custom policy.
 */
export function getDdosCustomPolicyOutput(args: GetDdosCustomPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDdosCustomPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network/v20231101:getDdosCustomPolicy", {
        "ddosCustomPolicyName": args.ddosCustomPolicyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDdosCustomPolicyOutputArgs {
    /**
     * The name of the DDoS custom policy.
     */
    ddosCustomPolicyName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
