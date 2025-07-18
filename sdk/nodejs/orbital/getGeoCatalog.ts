// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a GeoCatalog
 *
 * Uses Azure REST API version 2025-02-11-preview.
 *
 * Other available API versions: 2024-01-31-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGeoCatalog(args: GetGeoCatalogArgs, opts?: pulumi.InvokeOptions): Promise<GetGeoCatalogResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:orbital:getGeoCatalog", {
        "catalogName": args.catalogName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGeoCatalogArgs {
    /**
     * The name of the catalog
     */
    catalogName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A Microsoft Planetary Computer Pro GeoCatalog resource
 */
export interface GetGeoCatalogResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The managed service identities assigned to this resource.
     */
    readonly identity?: outputs.orbital.ManagedServiceIdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The resource-specific properties for this resource.
     */
    readonly properties: outputs.orbital.GeoCatalogPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.orbital.SystemDataResponse;
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
 * Get a GeoCatalog
 *
 * Uses Azure REST API version 2025-02-11-preview.
 *
 * Other available API versions: 2024-01-31-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGeoCatalogOutput(args: GetGeoCatalogOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGeoCatalogResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:orbital:getGeoCatalog", {
        "catalogName": args.catalogName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGeoCatalogOutputArgs {
    /**
     * The name of the catalog
     */
    catalogName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
