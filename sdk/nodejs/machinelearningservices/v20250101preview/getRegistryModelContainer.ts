// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Azure Resource Manager resource envelope.
 */
export function getRegistryModelContainer(args: GetRegistryModelContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryModelContainerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:machinelearningservices/v20250101preview:getRegistryModelContainer", {
        "modelName": args.modelName,
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRegistryModelContainerArgs {
    /**
     * Container name. This is case-sensitive.
     */
    modelName: string;
    /**
     * Name of Azure Machine Learning registry. This is case-insensitive
     */
    registryName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Azure Resource Manager resource envelope.
 */
export interface GetRegistryModelContainerResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * [Required] Additional attributes of the entity.
     */
    readonly modelContainerProperties: outputs.machinelearningservices.v20250101preview.ModelContainerResponse;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.machinelearningservices.v20250101preview.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Azure Resource Manager resource envelope.
 */
export function getRegistryModelContainerOutput(args: GetRegistryModelContainerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRegistryModelContainerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:machinelearningservices/v20250101preview:getRegistryModelContainer", {
        "modelName": args.modelName,
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRegistryModelContainerOutputArgs {
    /**
     * Container name. This is case-sensitive.
     */
    modelName: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning registry. This is case-insensitive
     */
    registryName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
