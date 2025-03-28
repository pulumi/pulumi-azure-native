// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a Dev Box definition
 */
export function getDevBoxDefinition(args: GetDevBoxDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetDevBoxDefinitionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:devcenter/v20250201:getDevBoxDefinition", {
        "devBoxDefinitionName": args.devBoxDefinitionName,
        "devCenterName": args.devCenterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDevBoxDefinitionArgs {
    /**
     * The name of the Dev Box definition.
     */
    devBoxDefinitionName: string;
    /**
     * The name of the devcenter.
     */
    devCenterName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Represents a definition for a Developer Machine.
 */
export interface GetDevBoxDefinitionResult {
    /**
     * Image reference information for the currently active image (only populated during updates).
     */
    readonly activeImageReference: outputs.devcenter.v20250201.ImageReferenceResponse;
    /**
     * Indicates whether Dev Boxes created with this definition are capable of hibernation. Not all images are capable of supporting hibernation. To find out more see https://aka.ms/devbox/hibernate
     */
    readonly hibernateSupport?: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * Image reference information.
     */
    readonly imageReference: outputs.devcenter.v20250201.ImageReferenceResponse;
    /**
     * Details for image validator error. Populated when the image validation is not successful.
     */
    readonly imageValidationErrorDetails: outputs.devcenter.v20250201.ImageValidationErrorDetailsResponse;
    /**
     * Validation status of the configured image.
     */
    readonly imageValidationStatus: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The storage type used for the Operating System disk of Dev Boxes created using this definition.
     */
    readonly osStorageType?: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * The SKU for Dev Boxes created using this definition.
     */
    readonly sku: outputs.devcenter.v20250201.SkuResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.devcenter.v20250201.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Validation status for the Dev Box Definition.
     */
    readonly validationStatus: string;
}
/**
 * Gets a Dev Box definition
 */
export function getDevBoxDefinitionOutput(args: GetDevBoxDefinitionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDevBoxDefinitionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:devcenter/v20250201:getDevBoxDefinition", {
        "devBoxDefinitionName": args.devBoxDefinitionName,
        "devCenterName": args.devCenterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDevBoxDefinitionOutputArgs {
    /**
     * The name of the Dev Box definition.
     */
    devBoxDefinitionName: pulumi.Input<string>;
    /**
     * The name of the devcenter.
     */
    devCenterName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
