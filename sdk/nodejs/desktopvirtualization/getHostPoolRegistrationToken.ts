// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Registration token of the host pool.
 *
 * Uses Azure REST API version 2024-04-03.
 *
 * Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getHostPoolRegistrationToken(args: GetHostPoolRegistrationTokenArgs, opts?: pulumi.InvokeOptions): Promise<GetHostPoolRegistrationTokenResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:desktopvirtualization:getHostPoolRegistrationToken", {
        "hostPoolName": args.hostPoolName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHostPoolRegistrationTokenArgs {
    /**
     * The name of the host pool within the specified resource group
     */
    hostPoolName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Represents a RegistrationInfo definition.
 */
export interface GetHostPoolRegistrationTokenResult {
    /**
     * Expiration time of registration token.
     */
    readonly expirationTime?: string;
    /**
     * The type of resetting the token.
     */
    readonly registrationTokenOperation?: string;
    /**
     * The registration token base64 encoded string.
     */
    readonly token?: string;
}
/**
 * Registration token of the host pool.
 *
 * Uses Azure REST API version 2024-04-03.
 *
 * Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getHostPoolRegistrationTokenOutput(args: GetHostPoolRegistrationTokenOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetHostPoolRegistrationTokenResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:desktopvirtualization:getHostPoolRegistrationToken", {
        "hostPoolName": args.hostPoolName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHostPoolRegistrationTokenOutputArgs {
    /**
     * The name of the host pool within the specified resource group
     */
    hostPoolName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
