// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets information about a server.
 */
export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:dbformysql/v20231001preview:getServer", {
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the server.
     */
    serverName: string;
}

/**
 * Represents a server.
 */
export interface GetServerResult {
    /**
     * The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
     */
    readonly administratorLogin?: string;
    /**
     * availability Zone information of the server.
     */
    readonly availabilityZone?: string;
    /**
     * Backup related properties of a server.
     */
    readonly backup?: outputs.dbformysql.v20231001preview.BackupResponse;
    /**
     * The Data Encryption for CMK.
     */
    readonly dataEncryption?: outputs.dbformysql.v20231001preview.DataEncryptionResponse;
    /**
     * The fully qualified domain name of a server.
     */
    readonly fullyQualifiedDomainName: string;
    /**
     * High availability related properties of a server.
     */
    readonly highAvailability?: outputs.dbformysql.v20231001preview.HighAvailabilityResponse;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The cmk identity for the server.
     */
    readonly identity?: outputs.dbformysql.v20231001preview.MySQLServerIdentityResponse;
    /**
     * Source properties for import from storage.
     */
    readonly importSourceProperties?: outputs.dbformysql.v20231001preview.ImportSourcePropertiesResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Maintenance window of a server.
     */
    readonly maintenanceWindow?: outputs.dbformysql.v20231001preview.MaintenanceWindowResponse;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Network related properties of a server.
     */
    readonly network?: outputs.dbformysql.v20231001preview.NetworkResponse;
    /**
     * PrivateEndpointConnections related properties of a server.
     */
    readonly privateEndpointConnections: outputs.dbformysql.v20231001preview.PrivateEndpointConnectionResponse[];
    /**
     * The maximum number of replicas that a primary server can have.
     */
    readonly replicaCapacity: number;
    /**
     * The replication role.
     */
    readonly replicationRole?: string;
    /**
     * The SKU (pricing tier) of the server.
     */
    readonly sku?: outputs.dbformysql.v20231001preview.MySQLServerSkuResponse;
    /**
     * The source MySQL server id.
     */
    readonly sourceServerResourceId?: string;
    /**
     * The state of a server.
     */
    readonly state: string;
    /**
     * Storage related properties of a server.
     */
    readonly storage?: outputs.dbformysql.v20231001preview.StorageResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.dbformysql.v20231001preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Server version.
     */
    readonly version?: string;
}
/**
 * Gets information about a server.
 */
export function getServerOutput(args: GetServerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:dbformysql/v20231001preview:getServer", {
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
}
