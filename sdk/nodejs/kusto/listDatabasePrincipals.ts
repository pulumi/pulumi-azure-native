// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Returns a list of database principals of the given Kusto cluster and database.
 *
 * Uses Azure REST API version 2022-12-29.
 *
 * Other available API versions: 2022-07-07, 2023-05-02, 2023-08-15, 2024-04-13.
 */
export function listDatabasePrincipals(args: ListDatabasePrincipalsArgs, opts?: pulumi.InvokeOptions): Promise<ListDatabasePrincipalsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:kusto:listDatabasePrincipals", {
        "clusterName": args.clusterName,
        "databaseName": args.databaseName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListDatabasePrincipalsArgs {
    /**
     * The name of the Kusto cluster.
     */
    clusterName: string;
    /**
     * The name of the database in the Kusto cluster.
     */
    databaseName: string;
    /**
     * The name of the resource group containing the Kusto cluster.
     */
    resourceGroupName: string;
}

/**
 * The list Kusto database principals operation response.
 */
export interface ListDatabasePrincipalsResult {
    /**
     * The list of Kusto database principals.
     */
    readonly value?: outputs.kusto.DatabasePrincipalResponse[];
}
/**
 * Returns a list of database principals of the given Kusto cluster and database.
 *
 * Uses Azure REST API version 2022-12-29.
 *
 * Other available API versions: 2022-07-07, 2023-05-02, 2023-08-15, 2024-04-13.
 */
export function listDatabasePrincipalsOutput(args: ListDatabasePrincipalsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListDatabasePrincipalsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:kusto:listDatabasePrincipals", {
        "clusterName": args.clusterName,
        "databaseName": args.databaseName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListDatabasePrincipalsOutputArgs {
    /**
     * The name of the Kusto cluster.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The name of the database in the Kusto cluster.
     */
    databaseName: pulumi.Input<string>;
    /**
     * The name of the resource group containing the Kusto cluster.
     */
    resourceGroupName: pulumi.Input<string>;
}
