// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get DigitalTwinsInstances Endpoint.
 */
export function getDigitalTwinsEndpoint(args: GetDigitalTwinsEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetDigitalTwinsEndpointResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:digitaltwins/v20230131:getDigitalTwinsEndpoint", {
        "endpointName": args.endpointName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetDigitalTwinsEndpointArgs {
    /**
     * Name of Endpoint Resource.
     */
    endpointName: string;
    /**
     * The name of the resource group that contains the DigitalTwinsInstance.
     */
    resourceGroupName: string;
    /**
     * The name of the DigitalTwinsInstance.
     */
    resourceName: string;
}

/**
 * DigitalTwinsInstance endpoint resource.
 */
export interface GetDigitalTwinsEndpointResult {
    /**
     * The resource identifier.
     */
    readonly id: string;
    /**
     * Extension resource name.
     */
    readonly name: string;
    /**
     * DigitalTwinsInstance endpoint resource properties.
     */
    readonly properties: outputs.digitaltwins.v20230131.EventGridResponse | outputs.digitaltwins.v20230131.EventHubResponse | outputs.digitaltwins.v20230131.ServiceBusResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.digitaltwins.v20230131.SystemDataResponse;
    /**
     * The resource type.
     */
    readonly type: string;
}
/**
 * Get DigitalTwinsInstances Endpoint.
 */
export function getDigitalTwinsEndpointOutput(args: GetDigitalTwinsEndpointOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDigitalTwinsEndpointResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:digitaltwins/v20230131:getDigitalTwinsEndpoint", {
        "endpointName": args.endpointName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetDigitalTwinsEndpointOutputArgs {
    /**
     * Name of Endpoint Resource.
     */
    endpointName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the DigitalTwinsInstance.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the DigitalTwinsInstance.
     */
    resourceName: pulumi.Input<string>;
}
