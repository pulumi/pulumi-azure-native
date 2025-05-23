// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieves a SQL Server ESU license resource
 *
 * Uses Azure REST API version 2024-05-01-preview.
 *
 * Other available API versions: 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getSqlServerEsuLicense(args: GetSqlServerEsuLicenseArgs, opts?: pulumi.InvokeOptions): Promise<GetSqlServerEsuLicenseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:azurearcdata:getSqlServerEsuLicense", {
        "resourceGroupName": args.resourceGroupName,
        "sqlServerEsuLicenseName": args.sqlServerEsuLicenseName,
    }, opts);
}

export interface GetSqlServerEsuLicenseArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Name of SQL Server ESU License
     */
    sqlServerEsuLicenseName: string;
}

/**
 * Describe SQL Server ESU license resource.
 */
export interface GetSqlServerEsuLicenseResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * SQL Server ESU license properties
     */
    readonly properties: outputs.azurearcdata.SqlServerEsuLicensePropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.azurearcdata.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Retrieves a SQL Server ESU license resource
 *
 * Uses Azure REST API version 2024-05-01-preview.
 *
 * Other available API versions: 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getSqlServerEsuLicenseOutput(args: GetSqlServerEsuLicenseOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSqlServerEsuLicenseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:azurearcdata:getSqlServerEsuLicense", {
        "resourceGroupName": args.resourceGroupName,
        "sqlServerEsuLicenseName": args.sqlServerEsuLicenseName,
    }, opts);
}

export interface GetSqlServerEsuLicenseOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of SQL Server ESU License
     */
    sqlServerEsuLicenseName: pulumi.Input<string>;
}
