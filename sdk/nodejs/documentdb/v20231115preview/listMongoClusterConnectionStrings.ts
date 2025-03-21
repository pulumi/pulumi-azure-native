// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * List mongo cluster connection strings. This includes the default connection string using SCRAM-SHA-256, as well as other connection strings supported by the cluster.
 */
export function listMongoClusterConnectionStrings(args: ListMongoClusterConnectionStringsArgs, opts?: pulumi.InvokeOptions): Promise<ListMongoClusterConnectionStringsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:documentdb/v20231115preview:listMongoClusterConnectionStrings", {
        "mongoClusterName": args.mongoClusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListMongoClusterConnectionStringsArgs {
    /**
     * The name of the mongo cluster.
     */
    mongoClusterName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * The connection strings for the given mongo cluster.
 */
export interface ListMongoClusterConnectionStringsResult {
    /**
     * An array that contains the connection strings for a mongo cluster.
     */
    readonly connectionStrings: outputs.documentdb.v20231115preview.ConnectionStringResponse[];
}
/**
 * List mongo cluster connection strings. This includes the default connection string using SCRAM-SHA-256, as well as other connection strings supported by the cluster.
 */
export function listMongoClusterConnectionStringsOutput(args: ListMongoClusterConnectionStringsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListMongoClusterConnectionStringsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:documentdb/v20231115preview:listMongoClusterConnectionStrings", {
        "mongoClusterName": args.mongoClusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListMongoClusterConnectionStringsOutputArgs {
    /**
     * The name of the mongo cluster.
     */
    mongoClusterName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
