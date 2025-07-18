// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets a Geo backup policy for the given database resource.
 *
 * Uses Azure REST API version 2023-08-01.
 *
 * Other available API versions: 2014-04-01, 2021-11-01, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGeoBackupPolicy(args: GetGeoBackupPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetGeoBackupPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:sql:getGeoBackupPolicy", {
        "databaseName": args.databaseName,
        "geoBackupPolicyName": args.geoBackupPolicyName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetGeoBackupPolicyArgs {
    /**
     * The name of the database.
     */
    databaseName: string;
    /**
     * The name of the Geo backup policy. This should always be 'Default'.
     */
    geoBackupPolicyName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: string;
    /**
     * The name of the server.
     */
    serverName: string;
}

/**
 * A Geo backup policy.
 */
export interface GetGeoBackupPolicyResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Kind of geo backup policy.  This is metadata used for the Azure portal experience.
     */
    readonly kind: string;
    /**
     * Backup policy location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The state of the geo backup policy.
     */
    readonly state: string;
    /**
     * The storage type of the geo backup policy.
     */
    readonly storageType: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Gets a Geo backup policy for the given database resource.
 *
 * Uses Azure REST API version 2023-08-01.
 *
 * Other available API versions: 2014-04-01, 2021-11-01, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGeoBackupPolicyOutput(args: GetGeoBackupPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGeoBackupPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:sql:getGeoBackupPolicy", {
        "databaseName": args.databaseName,
        "geoBackupPolicyName": args.geoBackupPolicyName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetGeoBackupPolicyOutputArgs {
    /**
     * The name of the database.
     */
    databaseName: pulumi.Input<string>;
    /**
     * The name of the Geo backup policy. This should always be 'Default'.
     */
    geoBackupPolicyName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
}
