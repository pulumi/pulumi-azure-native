// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Lists the PNS Credentials associated with a notification hub.
 *
 * Uses Azure REST API version 2023-10-01-preview.
 *
 * Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getNotificationHubPnsCredentials(args: GetNotificationHubPnsCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetNotificationHubPnsCredentialsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:notificationhubs:getNotificationHubPnsCredentials", {
        "namespaceName": args.namespaceName,
        "notificationHubName": args.notificationHubName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNotificationHubPnsCredentialsArgs {
    /**
     * Namespace name
     */
    namespaceName: string;
    /**
     * Notification Hub name
     */
    notificationHubName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Description of a NotificationHub PNS Credentials. This is a response of the POST requests that return namespace or hubs
 * PNS credentials.
 */
export interface GetNotificationHubPnsCredentialsResult {
    /**
     * Description of a NotificationHub AdmCredential.
     */
    readonly admCredential?: outputs.notificationhubs.AdmCredentialResponse;
    /**
     * Description of a NotificationHub ApnsCredential.
     */
    readonly apnsCredential?: outputs.notificationhubs.ApnsCredentialResponse;
    /**
     * Description of a NotificationHub BaiduCredential.
     */
    readonly baiduCredential?: outputs.notificationhubs.BaiduCredentialResponse;
    /**
     * Description of a NotificationHub BrowserCredential.
     */
    readonly browserCredential?: outputs.notificationhubs.BrowserCredentialResponse;
    /**
     * Description of a NotificationHub FcmV1Credential.
     */
    readonly fcmV1Credential?: outputs.notificationhubs.FcmV1CredentialResponse;
    /**
     * Description of a NotificationHub GcmCredential.
     */
    readonly gcmCredential?: outputs.notificationhubs.GcmCredentialResponse;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * Deprecated - only for compatibility.
     */
    readonly location?: string;
    /**
     * Description of a NotificationHub MpnsCredential.
     */
    readonly mpnsCredential?: outputs.notificationhubs.MpnsCredentialResponse;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.notificationhubs.SystemDataResponse;
    /**
     * Deprecated - only for compatibility.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Description of a NotificationHub WnsCredential.
     */
    readonly wnsCredential?: outputs.notificationhubs.WnsCredentialResponse;
    /**
     * Description of a NotificationHub XiaomiCredential.
     */
    readonly xiaomiCredential?: outputs.notificationhubs.XiaomiCredentialResponse;
}
/**
 * Lists the PNS Credentials associated with a notification hub.
 *
 * Uses Azure REST API version 2023-10-01-preview.
 *
 * Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getNotificationHubPnsCredentialsOutput(args: GetNotificationHubPnsCredentialsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNotificationHubPnsCredentialsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:notificationhubs:getNotificationHubPnsCredentials", {
        "namespaceName": args.namespaceName,
        "notificationHubName": args.notificationHubName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNotificationHubPnsCredentialsOutputArgs {
    /**
     * Namespace name
     */
    namespaceName: pulumi.Input<string>;
    /**
     * Notification Hub name
     */
    notificationHubName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
