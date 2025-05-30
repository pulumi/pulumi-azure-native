// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a DeploymentSetting
 *
 * Uses Azure REST API version 2024-04-01.
 *
 * Other available API versions: 2023-08-01-preview, 2023-11-01-preview, 2024-01-01, 2024-02-15-preview, 2024-09-01-preview, 2024-12-01-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getDeploymentSetting(args: GetDeploymentSettingArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentSettingResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:azurestackhci:getDeploymentSetting", {
        "clusterName": args.clusterName,
        "deploymentSettingsName": args.deploymentSettingsName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDeploymentSettingArgs {
    /**
     * The name of the cluster.
     */
    clusterName: string;
    /**
     * Name of Deployment Setting
     */
    deploymentSettingsName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Edge device resource
 */
export interface GetDeploymentSettingResult {
    /**
     * Azure resource ids of Arc machines to be part of cluster.
     */
    readonly arcNodeResourceIds: string[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Scale units will contains list of deployment data
     */
    readonly deploymentConfiguration: outputs.azurestackhci.DeploymentConfigurationResponse;
    /**
     * The deployment mode for cluster deployment.
     */
    readonly deploymentMode: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The intended operation for a cluster.
     */
    readonly operationType?: string;
    /**
     * DeploymentSetting provisioning state
     */
    readonly provisioningState: string;
    /**
     * Deployment Status reported from cluster.
     */
    readonly reportedProperties: outputs.azurestackhci.EceReportedPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.azurestackhci.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a DeploymentSetting
 *
 * Uses Azure REST API version 2024-04-01.
 *
 * Other available API versions: 2023-08-01-preview, 2023-11-01-preview, 2024-01-01, 2024-02-15-preview, 2024-09-01-preview, 2024-12-01-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getDeploymentSettingOutput(args: GetDeploymentSettingOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDeploymentSettingResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:azurestackhci:getDeploymentSetting", {
        "clusterName": args.clusterName,
        "deploymentSettingsName": args.deploymentSettingsName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDeploymentSettingOutputArgs {
    /**
     * The name of the cluster.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Name of Deployment Setting
     */
    deploymentSettingsName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
