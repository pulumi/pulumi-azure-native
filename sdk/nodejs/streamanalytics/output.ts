// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
 *
 * Uses Azure REST API version 2020-03-01. In version 2.x of the Azure Native provider, it used API version 2020-03-01.
 *
 * Other available API versions: 2021-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native streamanalytics [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class Output extends pulumi.CustomResource {
    /**
     * Get an existing Output resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Output {
        return new Output(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:streamanalytics:Output';

    /**
     * Returns true if the given object is an instance of Output.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Output {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Output.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
     */
    public readonly datasource!: pulumi.Output<outputs.streamanalytics.AzureDataLakeStoreOutputDataSourceResponse | outputs.streamanalytics.AzureFunctionOutputDataSourceResponse | outputs.streamanalytics.AzureSqlDatabaseOutputDataSourceResponse | outputs.streamanalytics.AzureSynapseOutputDataSourceResponse | outputs.streamanalytics.AzureTableOutputDataSourceResponse | outputs.streamanalytics.BlobOutputDataSourceResponse | outputs.streamanalytics.DocumentDbOutputDataSourceResponse | outputs.streamanalytics.EventHubOutputDataSourceResponse | outputs.streamanalytics.EventHubV2OutputDataSourceResponse | outputs.streamanalytics.GatewayMessageBusOutputDataSourceResponse | outputs.streamanalytics.PowerBIOutputDataSourceResponse | outputs.streamanalytics.ServiceBusQueueOutputDataSourceResponse | outputs.streamanalytics.ServiceBusTopicOutputDataSourceResponse | undefined>;
    /**
     * Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
     */
    public /*out*/ readonly diagnostics!: pulumi.Output<outputs.streamanalytics.DiagnosticsResponse>;
    /**
     * The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
     */
    public readonly serialization!: pulumi.Output<outputs.streamanalytics.AvroSerializationResponse | outputs.streamanalytics.CsvSerializationResponse | outputs.streamanalytics.JsonSerializationResponse | outputs.streamanalytics.ParquetSerializationResponse | undefined>;
    /**
     * The size window to constrain a Stream Analytics output to.
     */
    public readonly sizeWindow!: pulumi.Output<number | undefined>;
    /**
     * The time frame for filtering Stream Analytics job outputs.
     */
    public readonly timeWindow!: pulumi.Output<string | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Output resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OutputArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.jobName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["datasource"] = args ? args.datasource : undefined;
            resourceInputs["jobName"] = args ? args.jobName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputName"] = args ? args.outputName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serialization"] = args ? args.serialization : undefined;
            resourceInputs["sizeWindow"] = args ? args.sizeWindow : undefined;
            resourceInputs["timeWindow"] = args ? args.timeWindow : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["diagnostics"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["datasource"] = undefined /*out*/;
            resourceInputs["diagnostics"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serialization"] = undefined /*out*/;
            resourceInputs["sizeWindow"] = undefined /*out*/;
            resourceInputs["timeWindow"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:streamanalytics/v20160301:Output" }, { type: "azure-native:streamanalytics/v20170401preview:Output" }, { type: "azure-native:streamanalytics/v20200301:Output" }, { type: "azure-native:streamanalytics/v20211001preview:Output" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Output.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Output resource.
 */
export interface OutputArgs {
    /**
     * Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
     */
    datasource?: pulumi.Input<inputs.streamanalytics.AzureDataLakeStoreOutputDataSourceArgs | inputs.streamanalytics.AzureFunctionOutputDataSourceArgs | inputs.streamanalytics.AzureSqlDatabaseOutputDataSourceArgs | inputs.streamanalytics.AzureSynapseOutputDataSourceArgs | inputs.streamanalytics.AzureTableOutputDataSourceArgs | inputs.streamanalytics.BlobOutputDataSourceArgs | inputs.streamanalytics.DocumentDbOutputDataSourceArgs | inputs.streamanalytics.EventHubOutputDataSourceArgs | inputs.streamanalytics.EventHubV2OutputDataSourceArgs | inputs.streamanalytics.GatewayMessageBusOutputDataSourceArgs | inputs.streamanalytics.PowerBIOutputDataSourceArgs | inputs.streamanalytics.ServiceBusQueueOutputDataSourceArgs | inputs.streamanalytics.ServiceBusTopicOutputDataSourceArgs>;
    /**
     * The name of the streaming job.
     */
    jobName: pulumi.Input<string>;
    /**
     * Resource name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the output.
     */
    outputName?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
     */
    serialization?: pulumi.Input<inputs.streamanalytics.AvroSerializationArgs | inputs.streamanalytics.CsvSerializationArgs | inputs.streamanalytics.JsonSerializationArgs | inputs.streamanalytics.ParquetSerializationArgs>;
    /**
     * The size window to constrain a Stream Analytics output to.
     */
    sizeWindow?: pulumi.Input<number>;
    /**
     * The time frame for filtering Stream Analytics job outputs.
     */
    timeWindow?: pulumi.Input<string>;
}
