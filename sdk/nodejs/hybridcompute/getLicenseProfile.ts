// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieves information about the view of a license profile.
 *
 * Uses Azure REST API version 2023-06-20-preview.
 *
 * Other available API versions: 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-07-10, 2024-07-31-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13.
 */
export function getLicenseProfile(args: GetLicenseProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetLicenseProfileResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:hybridcompute:getLicenseProfile", {
        "licenseProfileName": args.licenseProfileName,
        "machineName": args.machineName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetLicenseProfileArgs {
    /**
     * The name of the license profile.
     */
    licenseProfileName: string;
    /**
     * The name of the hybrid machine.
     */
    machineName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Describes a license profile in a hybrid machine.
 */
export interface GetLicenseProfileResult {
    /**
     * The resource id of the license.
     */
    readonly assignedLicense?: string;
    /**
     * The guid id of the license.
     */
    readonly assignedLicenseImmutableId: string;
    /**
     * Indicates the eligibility state of Esu.
     */
    readonly esuEligibility: string;
    /**
     * Indicates whether there is an ESU Key currently active for the machine.
     */
    readonly esuKeyState: string;
    /**
     * The list of ESU keys.
     */
    readonly esuKeys: outputs.hybridcompute.EsuKeyResponse[];
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
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * The type of the Esu servers.
     */
    readonly serverType: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.hybridcompute.SystemDataResponse;
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
 * Retrieves information about the view of a license profile.
 *
 * Uses Azure REST API version 2023-06-20-preview.
 *
 * Other available API versions: 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-07-10, 2024-07-31-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13.
 */
export function getLicenseProfileOutput(args: GetLicenseProfileOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLicenseProfileResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:hybridcompute:getLicenseProfile", {
        "licenseProfileName": args.licenseProfileName,
        "machineName": args.machineName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetLicenseProfileOutputArgs {
    /**
     * The name of the license profile.
     */
    licenseProfileName: pulumi.Input<string>;
    /**
     * The name of the hybrid machine.
     */
    machineName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
