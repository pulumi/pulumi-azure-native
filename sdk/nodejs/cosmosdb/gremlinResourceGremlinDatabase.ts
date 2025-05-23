// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An Azure Cosmos DB Gremlin database.
 *
 * Uses Azure REST API version 2024-11-15.
 *
 * Other available API versions: 2019-08-01, 2019-12-12, 2020-03-01, 2020-04-01, 2020-06-01-preview, 2020-09-01, 2021-01-15, 2021-03-01-preview, 2021-03-15, 2021-04-01-preview, 2021-04-15, 2021-05-15, 2021-06-15, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class GremlinResourceGremlinDatabase extends pulumi.CustomResource {
    /**
     * Get an existing GremlinResourceGremlinDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GremlinResourceGremlinDatabase {
        return new GremlinResourceGremlinDatabase(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:cosmosdb:GremlinResourceGremlinDatabase';

    /**
     * Returns true if the given object is an instance of GremlinResourceGremlinDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GremlinResourceGremlinDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GremlinResourceGremlinDatabase.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the ARM resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly options!: pulumi.Output<outputs.cosmosdb.GremlinDatabaseGetPropertiesResponseOptions | undefined>;
    public readonly resource!: pulumi.Output<outputs.cosmosdb.GremlinDatabaseGetPropertiesResponseResource | undefined>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of Azure resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GremlinResourceGremlinDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GremlinResourceGremlinDatabaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["resource"] = args ? args.resource : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["options"] = undefined /*out*/;
            resourceInputs["resource"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:cosmosdb/v20150401:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20150408:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20151106:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20160319:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20160331:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20190801:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20191212:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20200301:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20200401:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20200601preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20200901:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20210115:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20210301preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20210315:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20210401preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20210415:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20210515:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20210615:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20210701preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20211015:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20211015preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20211115preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20220215preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20220515:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20220515preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20220815:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20220815preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20221115:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20221115preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20230301preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20230315:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20230315preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20230415:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20230915:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20230915preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20231115:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20231115preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20240215preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20240515:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20240515preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20240815:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20240901preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20241115:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20241201preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20250415:GremlinResourceGremlinDatabase" }, { type: "azure-native:cosmosdb/v20250501preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20230315preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20230415:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20230915:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20230915preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20231115:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20231115preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20240215preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20240515:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20240515preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20240815:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20240901preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20241115:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb/v20241201preview:GremlinResourceGremlinDatabase" }, { type: "azure-native:documentdb:GremlinResourceGremlinDatabase" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GremlinResourceGremlinDatabase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GremlinResourceGremlinDatabase resource.
 */
export interface GremlinResourceGremlinDatabaseArgs {
    /**
     * Cosmos DB database account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * Cosmos DB database name.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    location?: pulumi.Input<string>;
    /**
     * A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
     */
    options?: pulumi.Input<inputs.cosmosdb.CreateUpdateOptionsArgs>;
    /**
     * The standard JSON format of a Gremlin database
     */
    resource: pulumi.Input<inputs.cosmosdb.GremlinDatabaseResourceArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
