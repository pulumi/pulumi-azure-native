// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * cloud provider regions available for creating Schema Registry clusters.
 *
 * Uses Azure REST API version 2024-07-01.
 *
 * Other available API versions: 2024-02-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native confluent [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listOrganizationRegions(args: ListOrganizationRegionsArgs, opts?: pulumi.InvokeOptions): Promise<ListOrganizationRegionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:confluent:listOrganizationRegions", {
        "organizationName": args.organizationName,
        "resourceGroupName": args.resourceGroupName,
        "searchFilters": args.searchFilters,
    }, opts);
}

export interface ListOrganizationRegionsArgs {
    /**
     * Organization resource name
     */
    organizationName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Search filters for the request
     */
    searchFilters?: {[key: string]: string};
}

/**
 * Result of POST request to list regions supported by confluent
 */
export interface ListOrganizationRegionsResult {
    /**
     * List of regions supported by confluent
     */
    readonly data?: outputs.confluent.RegionRecordResponse[];
}
/**
 * cloud provider regions available for creating Schema Registry clusters.
 *
 * Uses Azure REST API version 2024-07-01.
 *
 * Other available API versions: 2024-02-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native confluent [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listOrganizationRegionsOutput(args: ListOrganizationRegionsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListOrganizationRegionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:confluent:listOrganizationRegions", {
        "organizationName": args.organizationName,
        "resourceGroupName": args.resourceGroupName,
        "searchFilters": args.searchFilters,
    }, opts);
}

export interface ListOrganizationRegionsOutputArgs {
    /**
     * Organization resource name
     */
    organizationName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Search filters for the request
     */
    searchFilters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
