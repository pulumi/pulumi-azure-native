// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets information about an azure ad administrator.
 *
 * Uses Azure REST API version 2023-12-30.
 *
 * Other available API versions: 2022-01-01, 2023-06-01-preview, 2023-06-30, 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbformysql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getAzureADAdministrator(args: GetAzureADAdministratorArgs, opts?: pulumi.InvokeOptions): Promise<GetAzureADAdministratorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:dbformysql:getAzureADAdministrator", {
        "administratorName": args.administratorName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetAzureADAdministratorArgs {
    /**
     * The name of the Azure AD Administrator.
     */
    administratorName: string;
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
 * Represents a Administrator.
 */
export interface GetAzureADAdministratorResult {
    /**
     * Type of the sever administrator.
     */
    readonly administratorType?: string;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The resource id of the identity used for AAD Authentication.
     */
    readonly identityResourceId?: string;
    /**
     * Login name of the server administrator.
     */
    readonly login?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * SID (object ID) of the server administrator.
     */
    readonly sid?: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.dbformysql.SystemDataResponse;
    /**
     * Tenant ID of the administrator.
     */
    readonly tenantId?: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets information about an azure ad administrator.
 *
 * Uses Azure REST API version 2023-12-30.
 *
 * Other available API versions: 2022-01-01, 2023-06-01-preview, 2023-06-30, 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbformysql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getAzureADAdministratorOutput(args: GetAzureADAdministratorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAzureADAdministratorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:dbformysql:getAzureADAdministrator", {
        "administratorName": args.administratorName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetAzureADAdministratorOutputArgs {
    /**
     * The name of the Azure AD Administrator.
     */
    administratorName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
}
