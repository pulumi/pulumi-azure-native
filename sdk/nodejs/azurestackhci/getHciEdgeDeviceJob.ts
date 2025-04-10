// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a EdgeDeviceJob
 *
 * Uses Azure REST API version 2024-12-01-preview.
 */
export function getHciEdgeDeviceJob(args: GetHciEdgeDeviceJobArgs, opts?: pulumi.InvokeOptions): Promise<GetHciEdgeDeviceJobResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:azurestackhci:getHciEdgeDeviceJob", {
        "edgeDeviceName": args.edgeDeviceName,
        "jobsName": args.jobsName,
        "resourceUri": args.resourceUri,
    }, opts);
}

export interface GetHciEdgeDeviceJobArgs {
    /**
     * The name of the EdgeDevice
     */
    edgeDeviceName: string;
    /**
     * Name of EdgeDevice Job
     */
    jobsName: string;
    /**
     * The fully qualified Azure Resource manager identifier of the resource.
     */
    resourceUri: string;
}

/**
 * Edge device job for Azure Stack HCI solution.
 */
export interface GetHciEdgeDeviceJobResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * Edge device kind.
     * Expected value is 'HCI'.
     */
    readonly kind: "HCI";
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * HCI Edge device job properties
     */
    readonly properties: outputs.azurestackhci.HciCollectLogJobPropertiesResponse | outputs.azurestackhci.HciRemoteSupportJobPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.azurestackhci.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a EdgeDeviceJob
 *
 * Uses Azure REST API version 2024-12-01-preview.
 */
export function getHciEdgeDeviceJobOutput(args: GetHciEdgeDeviceJobOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetHciEdgeDeviceJobResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:azurestackhci:getHciEdgeDeviceJob", {
        "edgeDeviceName": args.edgeDeviceName,
        "jobsName": args.jobsName,
        "resourceUri": args.resourceUri,
    }, opts);
}

export interface GetHciEdgeDeviceJobOutputArgs {
    /**
     * The name of the EdgeDevice
     */
    edgeDeviceName: pulumi.Input<string>;
    /**
     * Name of EdgeDevice Job
     */
    jobsName: pulumi.Input<string>;
    /**
     * The fully qualified Azure Resource manager identifier of the resource.
     */
    resourceUri: pulumi.Input<string>;
}
