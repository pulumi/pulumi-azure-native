// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a Watchlist in Azure Security Insights.
 *
 * Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
 *
 * Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-11-01, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class Watchlist extends pulumi.CustomResource {
    /**
     * Get an existing Watchlist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Watchlist {
        return new Watchlist(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:securityinsights:Watchlist';

    /**
     * Returns true if the given object is an instance of Watchlist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Watchlist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Watchlist.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The content type of the raw content. Example : text/csv or text/tsv
     */
    public readonly contentType!: pulumi.Output<string | undefined>;
    /**
     * The time the watchlist was created
     */
    public readonly created!: pulumi.Output<string | undefined>;
    /**
     * Describes a user that created the watchlist
     */
    public readonly createdBy!: pulumi.Output<outputs.securityinsights.WatchlistUserInfoResponse | undefined>;
    /**
     * The default duration of a watchlist (in ISO 8601 duration format)
     */
    public readonly defaultDuration!: pulumi.Output<string | undefined>;
    /**
     * A description of the watchlist
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the watchlist
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Etag of the azure resource
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * A flag that indicates if the watchlist is deleted or not
     */
    public readonly isDeleted!: pulumi.Output<boolean | undefined>;
    /**
     * The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
     */
    public readonly itemsSearchKey!: pulumi.Output<string>;
    /**
     * List of labels relevant to this watchlist
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The number of lines in a csv/tsv content to skip before the header
     */
    public readonly numberOfLinesToSkip!: pulumi.Output<number | undefined>;
    /**
     * The provider of the watchlist
     */
    public readonly provider!: pulumi.Output<string>;
    /**
     * Describes provisioning state
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
     */
    public readonly rawContent!: pulumi.Output<string | undefined>;
    /**
     * The filename of the watchlist, called 'source'
     */
    public readonly source!: pulumi.Output<string | undefined>;
    /**
     * The sourceType of the watchlist
     */
    public readonly sourceType!: pulumi.Output<string | undefined>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.securityinsights.SystemDataResponse>;
    /**
     * The tenantId where the watchlist belongs to
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The last time the watchlist was updated
     */
    public readonly updated!: pulumi.Output<string | undefined>;
    /**
     * Describes a user that updated the watchlist
     */
    public readonly updatedBy!: pulumi.Output<outputs.securityinsights.WatchlistUserInfoResponse | undefined>;
    /**
     * The status of the Watchlist upload : New, InProgress or Complete. **Note** : When a Watchlist upload status is InProgress, the Watchlist cannot be deleted
     */
    public readonly uploadStatus!: pulumi.Output<string | undefined>;
    /**
     * The alias of the watchlist
     */
    public readonly watchlistAlias!: pulumi.Output<string | undefined>;
    /**
     * The id (a Guid) of the watchlist
     */
    public readonly watchlistId!: pulumi.Output<string | undefined>;
    /**
     * The type of the watchlist
     */
    public readonly watchlistType!: pulumi.Output<string | undefined>;

    /**
     * Create a Watchlist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WatchlistArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.itemsSearchKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'itemsSearchKey'");
            }
            if ((!args || args.provider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'provider'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["created"] = args ? args.created : undefined;
            resourceInputs["createdBy"] = args ? args.createdBy : undefined;
            resourceInputs["defaultDuration"] = args ? args.defaultDuration : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["isDeleted"] = args ? args.isDeleted : undefined;
            resourceInputs["itemsSearchKey"] = args ? args.itemsSearchKey : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["numberOfLinesToSkip"] = args ? args.numberOfLinesToSkip : undefined;
            resourceInputs["provider"] = args ? args.provider : undefined;
            resourceInputs["rawContent"] = args ? args.rawContent : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["updated"] = args ? args.updated : undefined;
            resourceInputs["updatedBy"] = args ? args.updatedBy : undefined;
            resourceInputs["uploadStatus"] = args ? args.uploadStatus : undefined;
            resourceInputs["watchlistAlias"] = args ? args.watchlistAlias : undefined;
            resourceInputs["watchlistId"] = args ? args.watchlistId : undefined;
            resourceInputs["watchlistType"] = args ? args.watchlistType : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["contentType"] = undefined /*out*/;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["defaultDuration"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["isDeleted"] = undefined /*out*/;
            resourceInputs["itemsSearchKey"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["numberOfLinesToSkip"] = undefined /*out*/;
            resourceInputs["provider"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["rawContent"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
            resourceInputs["sourceType"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tenantId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
            resourceInputs["updatedBy"] = undefined /*out*/;
            resourceInputs["uploadStatus"] = undefined /*out*/;
            resourceInputs["watchlistAlias"] = undefined /*out*/;
            resourceInputs["watchlistId"] = undefined /*out*/;
            resourceInputs["watchlistType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:securityinsights/v20190101preview:Watchlist" }, { type: "azure-native:securityinsights/v20210301preview:Watchlist" }, { type: "azure-native:securityinsights/v20210401:Watchlist" }, { type: "azure-native:securityinsights/v20210901preview:Watchlist" }, { type: "azure-native:securityinsights/v20211001:Watchlist" }, { type: "azure-native:securityinsights/v20211001preview:Watchlist" }, { type: "azure-native:securityinsights/v20220101preview:Watchlist" }, { type: "azure-native:securityinsights/v20220401preview:Watchlist" }, { type: "azure-native:securityinsights/v20220501preview:Watchlist" }, { type: "azure-native:securityinsights/v20220601preview:Watchlist" }, { type: "azure-native:securityinsights/v20220701preview:Watchlist" }, { type: "azure-native:securityinsights/v20220801:Watchlist" }, { type: "azure-native:securityinsights/v20220801preview:Watchlist" }, { type: "azure-native:securityinsights/v20220901preview:Watchlist" }, { type: "azure-native:securityinsights/v20221001preview:Watchlist" }, { type: "azure-native:securityinsights/v20221101:Watchlist" }, { type: "azure-native:securityinsights/v20221101preview:Watchlist" }, { type: "azure-native:securityinsights/v20221201preview:Watchlist" }, { type: "azure-native:securityinsights/v20230201:Watchlist" }, { type: "azure-native:securityinsights/v20230201preview:Watchlist" }, { type: "azure-native:securityinsights/v20230301preview:Watchlist" }, { type: "azure-native:securityinsights/v20230401preview:Watchlist" }, { type: "azure-native:securityinsights/v20230501preview:Watchlist" }, { type: "azure-native:securityinsights/v20230601preview:Watchlist" }, { type: "azure-native:securityinsights/v20230701preview:Watchlist" }, { type: "azure-native:securityinsights/v20230801preview:Watchlist" }, { type: "azure-native:securityinsights/v20230901preview:Watchlist" }, { type: "azure-native:securityinsights/v20231001preview:Watchlist" }, { type: "azure-native:securityinsights/v20231101:Watchlist" }, { type: "azure-native:securityinsights/v20231201preview:Watchlist" }, { type: "azure-native:securityinsights/v20240101preview:Watchlist" }, { type: "azure-native:securityinsights/v20240301:Watchlist" }, { type: "azure-native:securityinsights/v20240401preview:Watchlist" }, { type: "azure-native:securityinsights/v20240901:Watchlist" }, { type: "azure-native:securityinsights/v20241001preview:Watchlist" }, { type: "azure-native:securityinsights/v20250101preview:Watchlist" }, { type: "azure-native:securityinsights/v20250301:Watchlist" }, { type: "azure-native:securityinsights/v20250401preview:Watchlist" }, { type: "azure-native:securityinsights/v20250601:Watchlist" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Watchlist.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Watchlist resource.
 */
export interface WatchlistArgs {
    /**
     * The content type of the raw content. Example : text/csv or text/tsv
     */
    contentType?: pulumi.Input<string>;
    /**
     * The time the watchlist was created
     */
    created?: pulumi.Input<string>;
    /**
     * Describes a user that created the watchlist
     */
    createdBy?: pulumi.Input<inputs.securityinsights.WatchlistUserInfoArgs>;
    /**
     * The default duration of a watchlist (in ISO 8601 duration format)
     */
    defaultDuration?: pulumi.Input<string>;
    /**
     * A description of the watchlist
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the watchlist
     */
    displayName: pulumi.Input<string>;
    /**
     * A flag that indicates if the watchlist is deleted or not
     */
    isDeleted?: pulumi.Input<boolean>;
    /**
     * The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
     */
    itemsSearchKey: pulumi.Input<string>;
    /**
     * List of labels relevant to this watchlist
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of lines in a csv/tsv content to skip before the header
     */
    numberOfLinesToSkip?: pulumi.Input<number>;
    /**
     * The provider of the watchlist
     */
    provider: pulumi.Input<string>;
    /**
     * The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
     */
    rawContent?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The filename of the watchlist, called 'source'
     */
    source?: pulumi.Input<string>;
    /**
     * The sourceType of the watchlist
     */
    sourceType?: pulumi.Input<string | enums.securityinsights.SourceType>;
    /**
     * The tenantId where the watchlist belongs to
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The last time the watchlist was updated
     */
    updated?: pulumi.Input<string>;
    /**
     * Describes a user that updated the watchlist
     */
    updatedBy?: pulumi.Input<inputs.securityinsights.WatchlistUserInfoArgs>;
    /**
     * The status of the Watchlist upload : New, InProgress or Complete. **Note** : When a Watchlist upload status is InProgress, the Watchlist cannot be deleted
     */
    uploadStatus?: pulumi.Input<string>;
    /**
     * The alias of the watchlist
     */
    watchlistAlias?: pulumi.Input<string>;
    /**
     * The id (a Guid) of the watchlist
     */
    watchlistId?: pulumi.Input<string>;
    /**
     * The type of the watchlist
     */
    watchlistType?: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
