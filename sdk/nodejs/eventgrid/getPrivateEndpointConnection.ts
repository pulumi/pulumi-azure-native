// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a specific private endpoint connection under a topic, domain, or partner namespace or namespace.
 *
 * Uses Azure REST API version 2025-02-15.
 *
 * Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPrivateEndpointConnection(args: GetPrivateEndpointConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateEndpointConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:eventgrid:getPrivateEndpointConnection", {
        "parentName": args.parentName,
        "parentType": args.parentType,
        "privateEndpointConnectionName": args.privateEndpointConnectionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateEndpointConnectionArgs {
    /**
     * The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name or namespace name).
     */
    parentName: string;
    /**
     * The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
     */
    parentType: string;
    /**
     * The name of the private endpoint connection connection.
     */
    privateEndpointConnectionName: string;
    /**
     * The name of the resource group within the user's subscription.
     */
    resourceGroupName: string;
}

export interface GetPrivateEndpointConnectionResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * GroupIds from the private link service resource.
     */
    readonly groupIds?: string[];
    /**
     * Fully qualified identifier of the resource.
     */
    readonly id: string;
    /**
     * Name of the resource.
     */
    readonly name: string;
    /**
     * The Private Endpoint resource for this Connection.
     */
    readonly privateEndpoint?: outputs.eventgrid.PrivateEndpointResponse;
    /**
     * Details about the state of the connection.
     */
    readonly privateLinkServiceConnectionState?: outputs.eventgrid.ConnectionStateResponse;
    /**
     * Provisioning state of the Private Endpoint Connection.
     */
    readonly provisioningState?: string;
    /**
     * Type of the resource.
     */
    readonly type: string;
}
/**
 * Get a specific private endpoint connection under a topic, domain, or partner namespace or namespace.
 *
 * Uses Azure REST API version 2025-02-15.
 *
 * Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPrivateEndpointConnectionOutput(args: GetPrivateEndpointConnectionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPrivateEndpointConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:eventgrid:getPrivateEndpointConnection", {
        "parentName": args.parentName,
        "parentType": args.parentType,
        "privateEndpointConnectionName": args.privateEndpointConnectionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateEndpointConnectionOutputArgs {
    /**
     * The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name or namespace name).
     */
    parentName: pulumi.Input<string>;
    /**
     * The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
     */
    parentType: pulumi.Input<string>;
    /**
     * The name of the private endpoint connection connection.
     */
    privateEndpointConnectionName: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription.
     */
    resourceGroupName: pulumi.Input<string>;
}
