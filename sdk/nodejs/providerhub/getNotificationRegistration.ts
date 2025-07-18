// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the notification registration details.
 *
 * Uses Azure REST API version 2024-09-01.
 *
 * Other available API versions: 2021-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native providerhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getNotificationRegistration(args: GetNotificationRegistrationArgs, opts?: pulumi.InvokeOptions): Promise<GetNotificationRegistrationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:providerhub:getNotificationRegistration", {
        "notificationRegistrationName": args.notificationRegistrationName,
        "providerNamespace": args.providerNamespace,
    }, opts);
}

export interface GetNotificationRegistrationArgs {
    /**
     * The notification registration.
     */
    notificationRegistrationName: string;
    /**
     * The name of the resource provider hosted within ProviderHub.
     */
    providerNamespace: string;
}

export interface GetNotificationRegistrationResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    readonly properties: outputs.providerhub.NotificationRegistrationPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.providerhub.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets the notification registration details.
 *
 * Uses Azure REST API version 2024-09-01.
 *
 * Other available API versions: 2021-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native providerhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getNotificationRegistrationOutput(args: GetNotificationRegistrationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNotificationRegistrationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:providerhub:getNotificationRegistration", {
        "notificationRegistrationName": args.notificationRegistrationName,
        "providerNamespace": args.providerNamespace,
    }, opts);
}

export interface GetNotificationRegistrationOutputArgs {
    /**
     * The notification registration.
     */
    notificationRegistrationName: pulumi.Input<string>;
    /**
     * The name of the resource provider hosted within ProviderHub.
     */
    providerNamespace: pulumi.Input<string>;
}
