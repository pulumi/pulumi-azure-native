// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the managed application definition.
 *
 * Uses Azure REST API version 2021-07-01.
 *
 * Other available API versions: 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native solutions [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getApplicationDefinition(args: GetApplicationDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationDefinitionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:solutions:getApplicationDefinition", {
        "applicationDefinitionName": args.applicationDefinitionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationDefinitionArgs {
    /**
     * The name of the managed application definition.
     */
    applicationDefinitionName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Information about managed application definition.
 */
export interface GetApplicationDefinitionResult {
    /**
     * The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
     */
    readonly artifacts?: outputs.solutions.ApplicationDefinitionArtifactResponse[];
    /**
     * The managed application provider authorizations.
     */
    readonly authorizations?: outputs.solutions.ApplicationAuthorizationResponse[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
     */
    readonly createUiDefinition?: any;
    /**
     * The managed application deployment policy.
     */
    readonly deploymentPolicy?: outputs.solutions.ApplicationDeploymentPolicyResponse;
    /**
     * The managed application definition description.
     */
    readonly description?: string;
    /**
     * The managed application definition display name.
     */
    readonly displayName?: string;
    /**
     * Resource ID
     */
    readonly id: string;
    /**
     * A value indicating whether the package is enabled or not.
     */
    readonly isEnabled?: boolean;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * The managed application lock level.
     */
    readonly lockLevel: string;
    /**
     * The managed application locking policy.
     */
    readonly lockingPolicy?: outputs.solutions.ApplicationPackageLockingPolicyDefinitionResponse;
    /**
     * The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
     */
    readonly mainTemplate?: any;
    /**
     * ID of the resource that manages this resource.
     */
    readonly managedBy?: string;
    /**
     * The managed application management policy that determines publisher's access to the managed resource group.
     */
    readonly managementPolicy?: outputs.solutions.ApplicationManagementPolicyResponse;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The managed application notification policy.
     */
    readonly notificationPolicy?: outputs.solutions.ApplicationNotificationPolicyResponse;
    /**
     * The managed application definition package file Uri. Use this element
     */
    readonly packageFileUri?: string;
    /**
     * The managed application provider policies.
     */
    readonly policies?: outputs.solutions.ApplicationPolicyResponse[];
    /**
     * The SKU of the resource.
     */
    readonly sku?: outputs.solutions.SkuResponse;
    /**
     * The storage account id for bring your own storage scenario.
     */
    readonly storageAccountId?: string;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.solutions.SystemDataResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
/**
 * Gets the managed application definition.
 *
 * Uses Azure REST API version 2021-07-01.
 *
 * Other available API versions: 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native solutions [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getApplicationDefinitionOutput(args: GetApplicationDefinitionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationDefinitionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:solutions:getApplicationDefinition", {
        "applicationDefinitionName": args.applicationDefinitionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationDefinitionOutputArgs {
    /**
     * The name of the managed application definition.
     */
    applicationDefinitionName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
