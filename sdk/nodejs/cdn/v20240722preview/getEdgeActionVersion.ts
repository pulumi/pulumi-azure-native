// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get EdgeActionVersion resource
 */
export function getEdgeActionVersion(args: GetEdgeActionVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetEdgeActionVersionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cdn/v20240722preview:getEdgeActionVersion", {
        "edgeActionName": args.edgeActionName,
        "resourceGroupName": args.resourceGroupName,
        "version": args.version,
    }, opts);
}

export interface GetEdgeActionVersionArgs {
    /**
     * The name of the Edge Action
     */
    edgeActionName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the Edge Action version
     */
    version: string;
}

/**
 * Concrete tracked resource types can be created by aliasing this type using a specific property type.
 */
export interface GetEdgeActionVersionResult {
    /**
     * The deployment type
     */
    readonly deploymentType: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The active state
     */
    readonly isDefaultVersion: string;
    /**
     * The last update time in UTC for package update
     */
    readonly lastPackageUpdateTime: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The provisioning state
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.cdn.v20240722preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The validation status
     */
    readonly validationStatus: string;
}
/**
 * Get EdgeActionVersion resource
 */
export function getEdgeActionVersionOutput(args: GetEdgeActionVersionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEdgeActionVersionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cdn/v20240722preview:getEdgeActionVersion", {
        "edgeActionName": args.edgeActionName,
        "resourceGroupName": args.resourceGroupName,
        "version": args.version,
    }, opts);
}

export interface GetEdgeActionVersionOutputArgs {
    /**
     * The name of the Edge Action
     */
    edgeActionName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Edge Action version
     */
    version: pulumi.Input<string>;
}
