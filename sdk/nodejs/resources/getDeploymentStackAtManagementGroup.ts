// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets a Deployment stack with a given name at Management Group scope.
 *
 * Uses Azure REST API version 2024-03-01.
 *
 * Other available API versions: 2022-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getDeploymentStackAtManagementGroup(args: GetDeploymentStackAtManagementGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentStackAtManagementGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:resources:getDeploymentStackAtManagementGroup", {
        "deploymentStackName": args.deploymentStackName,
        "managementGroupId": args.managementGroupId,
    }, opts);
}

export interface GetDeploymentStackAtManagementGroupArgs {
    /**
     * Name of the deployment stack.
     */
    deploymentStackName: string;
    /**
     * Management Group id.
     */
    managementGroupId: string;
}

/**
 * Deployment stack object.
 */
export interface GetDeploymentStackAtManagementGroupResult {
    /**
     * Defines the behavior of resources that are no longer managed after the Deployment stack is updated or deleted.
     */
    readonly actionOnUnmanage: outputs.resources.ActionOnUnmanageResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The correlation id of the last Deployment stack upsert or delete operation. It is in GUID format and is used for tracing.
     */
    readonly correlationId: string;
    /**
     * The debug setting of the deployment.
     */
    readonly debugSetting?: outputs.resources.DeploymentStacksDebugSettingResponse;
    /**
     * An array of resources that were deleted during the most recent Deployment stack update. Deleted means that the resource was removed from the template and relevant deletion operations were specified.
     */
    readonly deletedResources: outputs.resources.ResourceReferenceResponse[];
    /**
     * Defines how resources deployed by the stack are locked.
     */
    readonly denySettings: outputs.resources.DenySettingsResponse;
    /**
     * The resourceId of the deployment resource created by the deployment stack.
     */
    readonly deploymentId: string;
    /**
     * The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
     */
    readonly deploymentScope?: string;
    /**
     * Deployment stack description. Max length of 4096 characters.
     */
    readonly description?: string;
    /**
     * An array of resources that were detached during the most recent Deployment stack update. Detached means that the resource was removed from the template, but no relevant deletion operations were specified. So, the resource still exists while no longer being associated with the stack.
     */
    readonly detachedResources: outputs.resources.ResourceReferenceResponse[];
    /**
     * The duration of the last successful Deployment stack update.
     */
    readonly duration: string;
    /**
     * The error detail.
     */
    readonly error?: outputs.resources.ErrorDetailResponse;
    /**
     * An array of resources that failed to reach goal state during the most recent update. Each resourceId is accompanied by an error message.
     */
    readonly failedResources: outputs.resources.ResourceReferenceExtendedResponse[];
    /**
     * String Id used to locate any resource on Azure.
     */
    readonly id: string;
    /**
     * The location of the Deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
     */
    readonly location?: string;
    /**
     * Name of this resource.
     */
    readonly name: string;
    /**
     * The outputs of the deployment resource created by the deployment stack.
     */
    readonly outputs: any;
    /**
     * Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both.
     */
    readonly parameters?: {[key: string]: outputs.resources.DeploymentParameterResponse};
    /**
     * The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
     */
    readonly parametersLink?: outputs.resources.DeploymentStacksParametersLinkResponse;
    /**
     * State of the deployment stack.
     */
    readonly provisioningState: string;
    /**
     * An array of resources currently managed by the deployment stack.
     */
    readonly resources: outputs.resources.ManagedResourceReferenceResponse[];
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.resources.SystemDataResponse;
    /**
     * Deployment stack resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Type of this resource.
     */
    readonly type: string;
}
/**
 * Gets a Deployment stack with a given name at Management Group scope.
 *
 * Uses Azure REST API version 2024-03-01.
 *
 * Other available API versions: 2022-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getDeploymentStackAtManagementGroupOutput(args: GetDeploymentStackAtManagementGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDeploymentStackAtManagementGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:resources:getDeploymentStackAtManagementGroup", {
        "deploymentStackName": args.deploymentStackName,
        "managementGroupId": args.managementGroupId,
    }, opts);
}

export interface GetDeploymentStackAtManagementGroupOutputArgs {
    /**
     * Name of the deployment stack.
     */
    deploymentStackName: pulumi.Input<string>;
    /**
     * Management Group id.
     */
    managementGroupId: pulumi.Input<string>;
}
