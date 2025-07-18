// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get the access keys of the CommunicationService resource.
 *
 * Uses Azure REST API version 2023-06-01-preview.
 *
 * Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listCommunicationServiceKeys(args: ListCommunicationServiceKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListCommunicationServiceKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:communication:listCommunicationServiceKeys", {
        "communicationServiceName": args.communicationServiceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListCommunicationServiceKeysArgs {
    /**
     * The name of the CommunicationService resource.
     */
    communicationServiceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A class representing the access keys of a CommunicationService.
 */
export interface ListCommunicationServiceKeysResult {
    /**
     * CommunicationService connection string constructed via the primaryKey
     */
    readonly primaryConnectionString?: string;
    /**
     * The primary access key.
     */
    readonly primaryKey?: string;
    /**
     * CommunicationService connection string constructed via the secondaryKey
     */
    readonly secondaryConnectionString?: string;
    /**
     * The secondary access key.
     */
    readonly secondaryKey?: string;
}
/**
 * Get the access keys of the CommunicationService resource.
 *
 * Uses Azure REST API version 2023-06-01-preview.
 *
 * Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listCommunicationServiceKeysOutput(args: ListCommunicationServiceKeysOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListCommunicationServiceKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:communication:listCommunicationServiceKeys", {
        "communicationServiceName": args.communicationServiceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListCommunicationServiceKeysOutputArgs {
    /**
     * The name of the CommunicationService resource.
     */
    communicationServiceName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
