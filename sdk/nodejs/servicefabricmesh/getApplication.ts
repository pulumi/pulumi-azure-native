// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the information about the application resource with the given name. The information include the description and other properties of the application.
 *
 * Uses Azure REST API version 2018-09-01-preview.
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:servicefabricmesh:getApplication", {
        "applicationResourceName": args.applicationResourceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The identity of the application.
     */
    applicationResourceName: string;
    /**
     * Azure resource group name
     */
    resourceGroupName: string;
}

/**
 * This type describes an application resource.
 */
export interface GetApplicationResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Internal - used by Visual Studio to setup the debugging session on the local development environment.
     */
    readonly debugParams?: string;
    /**
     * User readable description of the application.
     */
    readonly description?: string;
    /**
     * Describes the diagnostics definition and usage for an application resource.
     */
    readonly diagnostics?: outputs.servicefabricmesh.DiagnosticsDescriptionResponse;
    /**
     * Describes the health state of an application resource.
     */
    readonly healthState: string;
    /**
     * Fully qualified identifier for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * State of the resource.
     */
    readonly provisioningState: string;
    /**
     * Names of the services in the application.
     */
    readonly serviceNames: string[];
    /**
     * Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
     */
    readonly services?: outputs.servicefabricmesh.ServiceResourceDescriptionResponse[];
    /**
     * Status of the application.
     */
    readonly status: string;
    /**
     * Gives additional information about the current status of the application.
     */
    readonly statusDetails: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
    /**
     * When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
     */
    readonly unhealthyEvaluation: string;
}
/**
 * Gets the information about the application resource with the given name. The information include the description and other properties of the application.
 *
 * Uses Azure REST API version 2018-09-01-preview.
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:servicefabricmesh:getApplication", {
        "applicationResourceName": args.applicationResourceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationOutputArgs {
    /**
     * The identity of the application.
     */
    applicationResourceName: pulumi.Input<string>;
    /**
     * Azure resource group name
     */
    resourceGroupName: pulumi.Input<string>;
}
