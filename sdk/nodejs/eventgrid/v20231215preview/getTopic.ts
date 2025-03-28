// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get properties of a topic.
 */
export function getTopic(args: GetTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:eventgrid/v20231215preview:getTopic", {
        "resourceGroupName": args.resourceGroupName,
        "topicName": args.topicName,
    }, opts);
}

export interface GetTopicArgs {
    /**
     * The name of the resource group within the user's subscription.
     */
    resourceGroupName: string;
    /**
     * Name of the topic.
     */
    topicName: string;
}

/**
 * EventGrid Topic
 */
export interface GetTopicResult {
    /**
     * Data Residency Boundary of the resource.
     */
    readonly dataResidencyBoundary?: string;
    /**
     * This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the topic.
     */
    readonly disableLocalAuth?: boolean;
    /**
     * Endpoint for the topic.
     */
    readonly endpoint: string;
    /**
     * Event Type Information for the user topic. This information is provided by the publisher and can be used by the 
     * subscriber to view different types of events that are published.
     */
    readonly eventTypeInfo?: outputs.eventgrid.v20231215preview.EventTypeInfoResponse;
    /**
     * Extended location of the resource.
     */
    readonly extendedLocation?: outputs.eventgrid.v20231215preview.ExtendedLocationResponse;
    /**
     * Fully qualified identifier of the resource.
     */
    readonly id: string;
    /**
     * Identity information for the resource.
     */
    readonly identity?: outputs.eventgrid.v20231215preview.IdentityInfoResponse;
    /**
     * This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
     */
    readonly inboundIpRules?: outputs.eventgrid.v20231215preview.InboundIpRuleResponse[];
    /**
     * This determines the format that Event Grid should expect for incoming events published to the topic.
     */
    readonly inputSchema?: string;
    /**
     * This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
     */
    readonly inputSchemaMapping?: outputs.eventgrid.v20231215preview.JsonInputSchemaMappingResponse;
    /**
     * Kind of the resource.
     */
    readonly kind?: string;
    /**
     * Location of the resource.
     */
    readonly location: string;
    /**
     * Metric resource id for the topic.
     */
    readonly metricResourceId: string;
    /**
     * Minimum TLS version of the publisher allowed to publish to this topic
     */
    readonly minimumTlsVersionAllowed?: string;
    /**
     * Name of the resource.
     */
    readonly name: string;
    readonly privateEndpointConnections: outputs.eventgrid.v20231215preview.PrivateEndpointConnectionResponse[];
    /**
     * Provisioning state of the topic.
     */
    readonly provisioningState: string;
    /**
     * This determines if traffic is allowed over public network. By default it is enabled. 
     * You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
     */
    readonly publicNetworkAccess?: string;
    /**
     * The Sku pricing tier for the topic.
     */
    readonly sku?: outputs.eventgrid.v20231215preview.ResourceSkuResponse;
    /**
     * The system metadata relating to Topic resource.
     */
    readonly systemData: outputs.eventgrid.v20231215preview.SystemDataResponse;
    /**
     * Tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Type of the resource.
     */
    readonly type: string;
}
/**
 * Get properties of a topic.
 */
export function getTopicOutput(args: GetTopicOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTopicResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:eventgrid/v20231215preview:getTopic", {
        "resourceGroupName": args.resourceGroupName,
        "topicName": args.topicName,
    }, opts);
}

export interface GetTopicOutputArgs {
    /**
     * The name of the resource group within the user's subscription.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the topic.
     */
    topicName: pulumi.Input<string>;
}
