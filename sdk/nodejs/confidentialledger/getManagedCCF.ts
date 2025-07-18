// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieves the properties of a Managed CCF app.
 *
 * Uses Azure REST API version 2023-06-28-preview.
 *
 * Other available API versions: 2022-09-08-preview, 2023-01-26-preview, 2024-07-09-preview, 2024-09-19-preview, 2025-06-10-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native confidentialledger [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getManagedCCF(args: GetManagedCCFArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedCCFResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:confidentialledger:getManagedCCF", {
        "appName": args.appName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedCCFArgs {
    /**
     * Name of the Managed CCF
     */
    appName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Managed CCF. Contains the properties of Managed CCF Resource.
 */
export interface GetManagedCCFResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
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
     * Properties of Managed CCF Resource.
     */
    readonly properties: outputs.confidentialledger.ManagedCCFPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.confidentialledger.SystemDataResponse;
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
 * Retrieves the properties of a Managed CCF app.
 *
 * Uses Azure REST API version 2023-06-28-preview.
 *
 * Other available API versions: 2022-09-08-preview, 2023-01-26-preview, 2024-07-09-preview, 2024-09-19-preview, 2025-06-10-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native confidentialledger [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getManagedCCFOutput(args: GetManagedCCFOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetManagedCCFResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:confidentialledger:getManagedCCF", {
        "appName": args.appName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedCCFOutputArgs {
    /**
     * Name of the Managed CCF
     */
    appName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
