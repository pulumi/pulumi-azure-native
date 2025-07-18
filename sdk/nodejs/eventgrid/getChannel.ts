// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get properties of a channel.
 *
 * Uses Azure REST API version 2025-02-15.
 *
 * Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getChannel(args: GetChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetChannelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:eventgrid:getChannel", {
        "channelName": args.channelName,
        "partnerNamespaceName": args.partnerNamespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetChannelArgs {
    /**
     * Name of the channel.
     */
    channelName: string;
    /**
     * Name of the partner namespace.
     */
    partnerNamespaceName: string;
    /**
     * The name of the resource group within the partners subscription.
     */
    resourceGroupName: string;
}

/**
 * Channel info.
 */
export interface GetChannelResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The type of the event channel which represents the direction flow of events.
     */
    readonly channelType?: string;
    /**
     * Expiration time of the channel. If this timer expires while the corresponding partner topic is never activated,
     * the channel and corresponding partner topic are deleted.
     */
    readonly expirationTimeIfNotActivatedUtc?: string;
    /**
     * Fully qualified identifier of the resource.
     */
    readonly id: string;
    /**
     * Context or helpful message that can be used during the approval process by the subscriber.
     */
    readonly messageForActivation?: string;
    /**
     * Name of the resource.
     */
    readonly name: string;
    /**
     * This property should be populated when channelType is PartnerTopic and represents information about the partner topic resource corresponding to the channel.
     */
    readonly partnerTopicInfo?: outputs.eventgrid.PartnerTopicInfoResponse;
    /**
     * Provisioning state of the channel.
     */
    readonly provisioningState?: string;
    /**
     * The readiness state of the corresponding partner topic.
     */
    readonly readinessState?: string;
    /**
     * The system metadata relating to the Event Grid resource.
     */
    readonly systemData: outputs.eventgrid.SystemDataResponse;
    /**
     * Type of the resource.
     */
    readonly type: string;
}
/**
 * Get properties of a channel.
 *
 * Uses Azure REST API version 2025-02-15.
 *
 * Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getChannelOutput(args: GetChannelOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetChannelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:eventgrid:getChannel", {
        "channelName": args.channelName,
        "partnerNamespaceName": args.partnerNamespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetChannelOutputArgs {
    /**
     * Name of the channel.
     */
    channelName: pulumi.Input<string>;
    /**
     * Name of the partner namespace.
     */
    partnerNamespaceName: pulumi.Input<string>;
    /**
     * The name of the resource group within the partners subscription.
     */
    resourceGroupName: pulumi.Input<string>;
}
