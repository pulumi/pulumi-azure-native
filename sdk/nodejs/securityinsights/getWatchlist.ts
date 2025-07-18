// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a watchlist, without its watchlist items.
 *
 * Uses Azure REST API version 2024-09-01.
 *
 * Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-11-01, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getWatchlist(args: GetWatchlistArgs, opts?: pulumi.InvokeOptions): Promise<GetWatchlistResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:securityinsights:getWatchlist", {
        "resourceGroupName": args.resourceGroupName,
        "watchlistAlias": args.watchlistAlias,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWatchlistArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The watchlist alias
     */
    watchlistAlias: string;
    /**
     * The name of the workspace.
     */
    workspaceName: string;
}

/**
 * Represents a Watchlist in Azure Security Insights.
 */
export interface GetWatchlistResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The content type of the raw content. Example : text/csv or text/tsv
     */
    readonly contentType?: string;
    /**
     * The time the watchlist was created
     */
    readonly created?: string;
    /**
     * Describes a user that created the watchlist
     */
    readonly createdBy?: outputs.securityinsights.WatchlistUserInfoResponse;
    /**
     * The default duration of a watchlist (in ISO 8601 duration format)
     */
    readonly defaultDuration?: string;
    /**
     * A description of the watchlist
     */
    readonly description?: string;
    /**
     * The display name of the watchlist
     */
    readonly displayName: string;
    /**
     * Etag of the azure resource
     */
    readonly etag?: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * A flag that indicates if the watchlist is deleted or not
     */
    readonly isDeleted?: boolean;
    /**
     * The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
     */
    readonly itemsSearchKey: string;
    /**
     * List of labels relevant to this watchlist
     */
    readonly labels?: string[];
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The number of lines in a csv/tsv content to skip before the header
     */
    readonly numberOfLinesToSkip?: number;
    /**
     * The provider of the watchlist
     */
    readonly provider: string;
    /**
     * Describes provisioning state
     */
    readonly provisioningState: string;
    /**
     * The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
     */
    readonly rawContent?: string;
    /**
     * The filename of the watchlist, called 'source'
     */
    readonly source?: string;
    /**
     * The sourceType of the watchlist
     */
    readonly sourceType?: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.securityinsights.SystemDataResponse;
    /**
     * The tenantId where the watchlist belongs to
     */
    readonly tenantId?: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The last time the watchlist was updated
     */
    readonly updated?: string;
    /**
     * Describes a user that updated the watchlist
     */
    readonly updatedBy?: outputs.securityinsights.WatchlistUserInfoResponse;
    /**
     * The status of the Watchlist upload : New, InProgress or Complete. **Note** : When a Watchlist upload status is InProgress, the Watchlist cannot be deleted
     */
    readonly uploadStatus?: string;
    /**
     * The alias of the watchlist
     */
    readonly watchlistAlias?: string;
    /**
     * The id (a Guid) of the watchlist
     */
    readonly watchlistId?: string;
    /**
     * The type of the watchlist
     */
    readonly watchlistType?: string;
}
/**
 * Get a watchlist, without its watchlist items.
 *
 * Uses Azure REST API version 2024-09-01.
 *
 * Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-11-01, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getWatchlistOutput(args: GetWatchlistOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWatchlistResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:securityinsights:getWatchlist", {
        "resourceGroupName": args.resourceGroupName,
        "watchlistAlias": args.watchlistAlias,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWatchlistOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The watchlist alias
     */
    watchlistAlias: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
