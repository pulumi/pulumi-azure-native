// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets an existing origin group within an endpoint.
 *
 * Uses Azure REST API version 2024-09-01.
 *
 * Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getOriginGroup(args: GetOriginGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetOriginGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cdn:getOriginGroup", {
        "endpointName": args.endpointName,
        "originGroupName": args.originGroupName,
        "profileName": args.profileName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOriginGroupArgs {
    /**
     * Name of the endpoint under the profile which is unique globally.
     */
    endpointName: string;
    /**
     * Name of the origin group which is unique within the endpoint.
     */
    originGroupName: string;
    /**
     * Name of the CDN profile which is unique within the resource group.
     */
    profileName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    resourceGroupName: string;
}

/**
 * Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
 */
export interface GetOriginGroupResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Health probe settings to the origin that is used to determine the health of the origin.
     */
    readonly healthProbeSettings?: outputs.cdn.HealthProbeParametersResponse;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The source of the content being delivered via CDN within given origin group.
     */
    readonly origins: outputs.cdn.ResourceReferenceResponse[];
    /**
     * Provisioning status of the origin group.
     */
    readonly provisioningState: string;
    /**
     * Resource status of the origin group.
     */
    readonly resourceState: string;
    /**
     * The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
     */
    readonly responseBasedOriginErrorDetectionSettings?: outputs.cdn.ResponseBasedOriginErrorDetectionParametersResponse;
    /**
     * Read only system data
     */
    readonly systemData: outputs.cdn.SystemDataResponse;
    /**
     * Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
     */
    readonly trafficRestorationTimeToHealedOrNewEndpointsInMinutes?: number;
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Gets an existing origin group within an endpoint.
 *
 * Uses Azure REST API version 2024-09-01.
 *
 * Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getOriginGroupOutput(args: GetOriginGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOriginGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cdn:getOriginGroup", {
        "endpointName": args.endpointName,
        "originGroupName": args.originGroupName,
        "profileName": args.profileName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOriginGroupOutputArgs {
    /**
     * Name of the endpoint under the profile which is unique globally.
     */
    endpointName: pulumi.Input<string>;
    /**
     * Name of the origin group which is unique within the endpoint.
     */
    originGroupName: pulumi.Input<string>;
    /**
     * Name of the CDN profile which is unique within the resource group.
     */
    profileName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    resourceGroupName: pulumi.Input<string>;
}
