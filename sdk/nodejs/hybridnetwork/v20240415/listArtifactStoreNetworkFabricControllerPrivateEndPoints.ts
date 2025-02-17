// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * List network fabric controllers to artifact stores
 */
export function listArtifactStoreNetworkFabricControllerPrivateEndPoints(args: ListArtifactStoreNetworkFabricControllerPrivateEndPointsArgs, opts?: pulumi.InvokeOptions): Promise<ListArtifactStoreNetworkFabricControllerPrivateEndPointsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:hybridnetwork/v20240415:listArtifactStoreNetworkFabricControllerPrivateEndPoints", {
        "artifactStoreName": args.artifactStoreName,
        "publisherName": args.publisherName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListArtifactStoreNetworkFabricControllerPrivateEndPointsArgs {
    /**
     * The name of the artifact store.
     */
    artifactStoreName: string;
    /**
     * The name of the publisher.
     */
    publisherName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * List of manual private endpoints.
 */
export interface ListArtifactStoreNetworkFabricControllerPrivateEndPointsResult {
    /**
     * The URI to get the next set of results.
     */
    readonly nextLink: string;
    /**
     * A list of network fabric controllers.
     */
    readonly value?: outputs.hybridnetwork.v20240415.ArtifactStoreNetworkFabricControllerEndPointsResponse[];
}
/**
 * List network fabric controllers to artifact stores
 */
export function listArtifactStoreNetworkFabricControllerPrivateEndPointsOutput(args: ListArtifactStoreNetworkFabricControllerPrivateEndPointsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListArtifactStoreNetworkFabricControllerPrivateEndPointsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:hybridnetwork/v20240415:listArtifactStoreNetworkFabricControllerPrivateEndPoints", {
        "artifactStoreName": args.artifactStoreName,
        "publisherName": args.publisherName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListArtifactStoreNetworkFabricControllerPrivateEndPointsOutputArgs {
    /**
     * The name of the artifact store.
     */
    artifactStoreName: pulumi.Input<string>;
    /**
     * The name of the publisher.
     */
    publisherName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
