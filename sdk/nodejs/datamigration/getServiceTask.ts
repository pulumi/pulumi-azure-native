// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The service tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The GET method retrieves information about a service task.
 *
 * Uses Azure REST API version 2021-06-30.
 *
 * Other available API versions: 2022-03-30-preview, 2023-07-15-preview.
 */
export function getServiceTask(args: GetServiceTaskArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceTaskResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:datamigration:getServiceTask", {
        "expand": args.expand,
        "groupName": args.groupName,
        "serviceName": args.serviceName,
        "taskName": args.taskName,
    }, opts);
}

export interface GetServiceTaskArgs {
    /**
     * Expand the response
     */
    expand?: string;
    /**
     * Name of the resource group
     */
    groupName: string;
    /**
     * Name of the service
     */
    serviceName: string;
    /**
     * Name of the Task
     */
    taskName: string;
}

/**
 * A task resource
 */
export interface GetServiceTaskResult {
    /**
     * HTTP strong entity tag value. This is ignored if submitted.
     */
    readonly etag?: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Custom task properties
     */
    readonly properties: outputs.datamigration.ConnectToMongoDbTaskPropertiesResponse | outputs.datamigration.ConnectToSourceMySqlTaskPropertiesResponse | outputs.datamigration.ConnectToSourceOracleSyncTaskPropertiesResponse | outputs.datamigration.ConnectToSourcePostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.ConnectToSourceSqlServerSyncTaskPropertiesResponse | outputs.datamigration.ConnectToSourceSqlServerTaskPropertiesResponse | outputs.datamigration.ConnectToTargetAzureDbForMySqlTaskPropertiesResponse | outputs.datamigration.ConnectToTargetAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.ConnectToTargetSqlDbTaskPropertiesResponse | outputs.datamigration.ConnectToTargetSqlMISyncTaskPropertiesResponse | outputs.datamigration.ConnectToTargetSqlMITaskPropertiesResponse | outputs.datamigration.ConnectToTargetSqlSqlDbSyncTaskPropertiesResponse | outputs.datamigration.GetTdeCertificatesSqlTaskPropertiesResponse | outputs.datamigration.GetUserTablesMySqlTaskPropertiesResponse | outputs.datamigration.GetUserTablesOracleTaskPropertiesResponse | outputs.datamigration.GetUserTablesPostgreSqlTaskPropertiesResponse | outputs.datamigration.GetUserTablesSqlSyncTaskPropertiesResponse | outputs.datamigration.GetUserTablesSqlTaskPropertiesResponse | outputs.datamigration.MigrateMongoDbTaskPropertiesResponse | outputs.datamigration.MigrateMySqlAzureDbForMySqlOfflineTaskPropertiesResponse | outputs.datamigration.MigrateMySqlAzureDbForMySqlSyncTaskPropertiesResponse | outputs.datamigration.MigrateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.MigratePostgreSqlAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.MigrateSqlServerSqlDbSyncTaskPropertiesResponse | outputs.datamigration.MigrateSqlServerSqlDbTaskPropertiesResponse | outputs.datamigration.MigrateSqlServerSqlMISyncTaskPropertiesResponse | outputs.datamigration.MigrateSqlServerSqlMITaskPropertiesResponse | outputs.datamigration.MigrateSsisTaskPropertiesResponse | outputs.datamigration.ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponse | outputs.datamigration.ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponse | outputs.datamigration.ValidateMigrationInputSqlServerSqlMITaskPropertiesResponse | outputs.datamigration.ValidateMongoDbTaskPropertiesResponse | outputs.datamigration.ValidateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.datamigration.SystemDataResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * The service tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The GET method retrieves information about a service task.
 *
 * Uses Azure REST API version 2021-06-30.
 *
 * Other available API versions: 2022-03-30-preview, 2023-07-15-preview.
 */
export function getServiceTaskOutput(args: GetServiceTaskOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceTaskResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:datamigration:getServiceTask", {
        "expand": args.expand,
        "groupName": args.groupName,
        "serviceName": args.serviceName,
        "taskName": args.taskName,
    }, opts);
}

export interface GetServiceTaskOutputArgs {
    /**
     * Expand the response
     */
    expand?: pulumi.Input<string>;
    /**
     * Name of the resource group
     */
    groupName: pulumi.Input<string>;
    /**
     * Name of the service
     */
    serviceName: pulumi.Input<string>;
    /**
     * Name of the Task
     */
    taskName: pulumi.Input<string>;
}
