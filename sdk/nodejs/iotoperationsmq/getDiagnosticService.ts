// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a DiagnosticServiceResource
 *
 * Uses Azure REST API version 2023-10-04-preview.
 */
export function getDiagnosticService(args: GetDiagnosticServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetDiagnosticServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:iotoperationsmq:getDiagnosticService", {
        "diagnosticServiceName": args.diagnosticServiceName,
        "mqName": args.mqName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDiagnosticServiceArgs {
    /**
     * Name of MQ diagnostic resource
     */
    diagnosticServiceName: string;
    /**
     * Name of MQ resource
     */
    mqName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * MQ diagnostic services resource
 */
export interface GetDiagnosticServiceResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The frequency at which the data will be exported.
     */
    readonly dataExportFrequencySeconds?: number;
    /**
     * Extended Location
     */
    readonly extendedLocation: outputs.iotoperationsmq.ExtendedLocationPropertyResponse;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The details of Diagnostic Service Docker Image.
     */
    readonly image: outputs.iotoperationsmq.ContainerImageResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The format for the logs generated.
     */
    readonly logFormat?: string;
    /**
     * The format for the logs generated.
     */
    readonly logLevel?: string;
    /**
     * The maximum data stored in MiB.
     */
    readonly maxDataStorageSize?: number;
    /**
     * The port at which metrics is exposed.
     */
    readonly metricsPort?: number;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The destination to collect traces. Diagnostic service will push traces to this endpoint
     */
    readonly openTelemetryTracesCollectorAddr?: string;
    /**
     * The status of the last operation.
     */
    readonly provisioningState: string;
    /**
     * Metric inactivity timeout.
     */
    readonly staleDataTimeoutSeconds?: number;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.iotoperationsmq.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a DiagnosticServiceResource
 *
 * Uses Azure REST API version 2023-10-04-preview.
 */
export function getDiagnosticServiceOutput(args: GetDiagnosticServiceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDiagnosticServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:iotoperationsmq:getDiagnosticService", {
        "diagnosticServiceName": args.diagnosticServiceName,
        "mqName": args.mqName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDiagnosticServiceOutputArgs {
    /**
     * Name of MQ diagnostic resource
     */
    diagnosticServiceName: pulumi.Input<string>;
    /**
     * Name of MQ resource
     */
    mqName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
