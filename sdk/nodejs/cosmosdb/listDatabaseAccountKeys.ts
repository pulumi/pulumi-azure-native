// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Lists the access keys for the specified Azure Cosmos DB database account.
 *
 * Uses Azure REST API version 2024-11-15.
 *
 * Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19, 2016-03-31, 2019-08-01, 2019-12-12, 2020-03-01, 2020-04-01, 2020-06-01-preview, 2020-09-01, 2021-01-15, 2021-03-01-preview, 2021-03-15, 2021-04-01-preview, 2021-04-15, 2021-05-15, 2021-06-15, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listDatabaseAccountKeys(args: ListDatabaseAccountKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListDatabaseAccountKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cosmosdb:listDatabaseAccountKeys", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListDatabaseAccountKeysArgs {
    /**
     * Cosmos DB database account name.
     */
    accountName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * The access keys for the given database account.
 */
export interface ListDatabaseAccountKeysResult {
    /**
     * Base 64 encoded value of the primary read-write key.
     */
    readonly primaryMasterKey: string;
    /**
     * Base 64 encoded value of the primary read-only key.
     */
    readonly primaryReadonlyMasterKey: string;
    /**
     * Base 64 encoded value of the secondary read-write key.
     */
    readonly secondaryMasterKey: string;
    /**
     * Base 64 encoded value of the secondary read-only key.
     */
    readonly secondaryReadonlyMasterKey: string;
}
/**
 * Lists the access keys for the specified Azure Cosmos DB database account.
 *
 * Uses Azure REST API version 2024-11-15.
 *
 * Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19, 2016-03-31, 2019-08-01, 2019-12-12, 2020-03-01, 2020-04-01, 2020-06-01-preview, 2020-09-01, 2021-01-15, 2021-03-01-preview, 2021-03-15, 2021-04-01-preview, 2021-04-15, 2021-05-15, 2021-06-15, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listDatabaseAccountKeysOutput(args: ListDatabaseAccountKeysOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListDatabaseAccountKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cosmosdb:listDatabaseAccountKeys", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListDatabaseAccountKeysOutputArgs {
    /**
     * Cosmos DB database account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
