// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getResourceGuard(args: GetResourceGuardArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceGuardResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:dataprotection/v20250201:getResourceGuard", {
        "resourceGroupName": args.resourceGroupName,
        "resourceGuardsName": args.resourceGuardsName,
    }, opts);
}

export interface GetResourceGuardArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of ResourceGuard
     */
    resourceGuardsName: string;
}

export interface GetResourceGuardResult {
    /**
     * Optional ETag.
     */
    readonly eTag?: string;
    /**
     * Resource Id represents the complete path to the resource.
     */
    readonly id: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name associated with the resource.
     */
    readonly name: string;
    /**
     * ResourceGuardResource properties
     */
    readonly properties: outputs.dataprotection.v20250201.ResourceGuardResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.dataprotection.v20250201.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    readonly type: string;
}
export function getResourceGuardOutput(args: GetResourceGuardOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetResourceGuardResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:dataprotection/v20250201:getResourceGuard", {
        "resourceGroupName": args.resourceGroupName,
        "resourceGuardsName": args.resourceGuardsName,
    }, opts);
}

export interface GetResourceGuardOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of ResourceGuard
     */
    resourceGuardsName: pulumi.Input<string>;
}
