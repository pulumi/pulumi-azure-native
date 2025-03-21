// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets details about the specified output.
 */
export function getOutput(args: GetOutputArgs, opts?: pulumi.InvokeOptions): Promise<GetOutputResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:streamanalytics/v20200301:getOutput", {
        "jobName": args.jobName,
        "outputName": args.outputName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOutputArgs {
    /**
     * The name of the streaming job.
     */
    jobName: string;
    /**
     * The name of the output.
     */
    outputName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
 */
export interface GetOutputResult {
    /**
     * Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
     */
    readonly datasource?: outputs.streamanalytics.v20200301.AzureDataLakeStoreOutputDataSourceResponse | outputs.streamanalytics.v20200301.AzureFunctionOutputDataSourceResponse | outputs.streamanalytics.v20200301.AzureSqlDatabaseOutputDataSourceResponse | outputs.streamanalytics.v20200301.AzureSynapseOutputDataSourceResponse | outputs.streamanalytics.v20200301.AzureTableOutputDataSourceResponse | outputs.streamanalytics.v20200301.BlobOutputDataSourceResponse | outputs.streamanalytics.v20200301.DocumentDbOutputDataSourceResponse | outputs.streamanalytics.v20200301.EventHubOutputDataSourceResponse | outputs.streamanalytics.v20200301.EventHubV2OutputDataSourceResponse | outputs.streamanalytics.v20200301.GatewayMessageBusOutputDataSourceResponse | outputs.streamanalytics.v20200301.PowerBIOutputDataSourceResponse | outputs.streamanalytics.v20200301.ServiceBusQueueOutputDataSourceResponse | outputs.streamanalytics.v20200301.ServiceBusTopicOutputDataSourceResponse;
    /**
     * Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
     */
    readonly diagnostics: outputs.streamanalytics.v20200301.DiagnosticsResponse;
    /**
     * The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
     */
    readonly etag: string;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Resource name
     */
    readonly name?: string;
    /**
     * Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
     */
    readonly serialization?: outputs.streamanalytics.v20200301.AvroSerializationResponse | outputs.streamanalytics.v20200301.CsvSerializationResponse | outputs.streamanalytics.v20200301.JsonSerializationResponse | outputs.streamanalytics.v20200301.ParquetSerializationResponse;
    /**
     * The size window to constrain a Stream Analytics output to.
     */
    readonly sizeWindow?: number;
    /**
     * The time frame for filtering Stream Analytics job outputs.
     */
    readonly timeWindow?: string;
    /**
     * Resource type
     */
    readonly type: string;
}
/**
 * Gets details about the specified output.
 */
export function getOutputOutput(args: GetOutputOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOutputResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:streamanalytics/v20200301:getOutput", {
        "jobName": args.jobName,
        "outputName": args.outputName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOutputOutputArgs {
    /**
     * The name of the streaming job.
     */
    jobName: pulumi.Input<string>;
    /**
     * The name of the output.
     */
    outputName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
