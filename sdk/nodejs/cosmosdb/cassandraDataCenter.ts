// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A managed Cassandra data center.
 *
 * Uses Azure REST API version 2024-11-15.
 *
 * Other available API versions: 2021-03-01-preview, 2021-04-01-preview, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class CassandraDataCenter extends pulumi.CustomResource {
    /**
     * Get an existing CassandraDataCenter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CassandraDataCenter {
        return new CassandraDataCenter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:cosmosdb:CassandraDataCenter';

    /**
     * Returns true if the given object is an instance of CassandraDataCenter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CassandraDataCenter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CassandraDataCenter.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The name of the database account.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of a managed Cassandra data center.
     */
    public readonly properties!: pulumi.Output<outputs.cosmosdb.DataCenterResourceResponseProperties>;
    /**
     * The type of Azure resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CassandraDataCenter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CassandraDataCenterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["dataCenterName"] = args ? args.dataCenterName : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:cosmosdb/v20210301preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20210401preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20210701preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20211015:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20211015preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20211115preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20220215preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20220515:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20220515preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20220815:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20220815preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20221115:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20221115preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20230301preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20230315:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20230315preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20230415:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20230915:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20230915preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20231115:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20231115preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20240215preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20240515:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20240515preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20240815:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20240901preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20241115:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20241201preview:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20250415:CassandraDataCenter" }, { type: "azure-native:cosmosdb/v20250501preview:CassandraDataCenter" }, { type: "azure-native:documentdb/v20230415:CassandraDataCenter" }, { type: "azure-native:documentdb/v20230915:CassandraDataCenter" }, { type: "azure-native:documentdb/v20230915preview:CassandraDataCenter" }, { type: "azure-native:documentdb/v20231115:CassandraDataCenter" }, { type: "azure-native:documentdb/v20231115preview:CassandraDataCenter" }, { type: "azure-native:documentdb/v20240215preview:CassandraDataCenter" }, { type: "azure-native:documentdb/v20240515:CassandraDataCenter" }, { type: "azure-native:documentdb/v20240515preview:CassandraDataCenter" }, { type: "azure-native:documentdb/v20240815:CassandraDataCenter" }, { type: "azure-native:documentdb/v20240901preview:CassandraDataCenter" }, { type: "azure-native:documentdb/v20241115:CassandraDataCenter" }, { type: "azure-native:documentdb/v20241201preview:CassandraDataCenter" }, { type: "azure-native:documentdb:CassandraDataCenter" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(CassandraDataCenter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CassandraDataCenter resource.
 */
export interface CassandraDataCenterArgs {
    /**
     * Managed Cassandra cluster name.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Data center name in a managed Cassandra cluster.
     */
    dataCenterName?: pulumi.Input<string>;
    /**
     * Properties of a managed Cassandra data center.
     */
    properties?: pulumi.Input<inputs.cosmosdb.DataCenterResourcePropertiesArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
