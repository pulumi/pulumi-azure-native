// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Execute Insights for an entity.
 *
 * Uses Azure REST API version 2025-01-01-preview.
 *
 * Other available API versions: 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getEntityInsights(args: GetEntityInsightsArgs, opts?: pulumi.InvokeOptions): Promise<GetEntityInsightsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:securityinsights:getEntityInsights", {
        "addDefaultExtendedTimeRange": args.addDefaultExtendedTimeRange,
        "endTime": args.endTime,
        "entityId": args.entityId,
        "insightQueryIds": args.insightQueryIds,
        "resourceGroupName": args.resourceGroupName,
        "startTime": args.startTime,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetEntityInsightsArgs {
    /**
     * Indicates if query time range should be extended with default time range of the query. Default value is false
     */
    addDefaultExtendedTimeRange?: boolean;
    /**
     * The end timeline date, so the results returned are before this date.
     */
    endTime: string;
    /**
     * entity ID
     */
    entityId: string;
    /**
     * List of Insights Query Id. If empty, default value is all insights of this entity
     */
    insightQueryIds?: string[];
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The start timeline date, so the results returned are after this date.
     */
    startTime: string;
    /**
     * The name of the workspace.
     */
    workspaceName: string;
}

/**
 * The Get Insights result operation response.
 */
export interface GetEntityInsightsResult {
    /**
     * The metadata from the get insights operation results.
     */
    readonly metaData?: outputs.securityinsights.GetInsightsResultsMetadataResponse;
    /**
     * The insights result values.
     */
    readonly value?: outputs.securityinsights.EntityInsightItemResponse[];
}
/**
 * Execute Insights for an entity.
 *
 * Uses Azure REST API version 2025-01-01-preview.
 *
 * Other available API versions: 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getEntityInsightsOutput(args: GetEntityInsightsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEntityInsightsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:securityinsights:getEntityInsights", {
        "addDefaultExtendedTimeRange": args.addDefaultExtendedTimeRange,
        "endTime": args.endTime,
        "entityId": args.entityId,
        "insightQueryIds": args.insightQueryIds,
        "resourceGroupName": args.resourceGroupName,
        "startTime": args.startTime,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetEntityInsightsOutputArgs {
    /**
     * Indicates if query time range should be extended with default time range of the query. Default value is false
     */
    addDefaultExtendedTimeRange?: pulumi.Input<boolean>;
    /**
     * The end timeline date, so the results returned are before this date.
     */
    endTime: pulumi.Input<string>;
    /**
     * entity ID
     */
    entityId: pulumi.Input<string>;
    /**
     * List of Insights Query Id. If empty, default value is all insights of this entity
     */
    insightQueryIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The start timeline date, so the results returned are after this date.
     */
    startTime: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
