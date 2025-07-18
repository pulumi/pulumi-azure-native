// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * List SAS credentials of a storage account.
 *
 * Uses Azure REST API version 2024-01-01.
 *
 * Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listStorageAccountSAS(args: ListStorageAccountSASArgs, opts?: pulumi.InvokeOptions): Promise<ListStorageAccountSASResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:storage:listStorageAccountSAS", {
        "accountName": args.accountName,
        "iPAddressOrRange": args.iPAddressOrRange,
        "keyToSign": args.keyToSign,
        "permissions": args.permissions,
        "protocols": args.protocols,
        "resourceGroupName": args.resourceGroupName,
        "resourceTypes": args.resourceTypes,
        "services": args.services,
        "sharedAccessExpiryTime": args.sharedAccessExpiryTime,
        "sharedAccessStartTime": args.sharedAccessStartTime,
    }, opts);
}

export interface ListStorageAccountSASArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    accountName: string;
    /**
     * An IP address or a range of IP addresses from which to accept requests.
     */
    iPAddressOrRange?: string;
    /**
     * The key to sign the account SAS token with.
     */
    keyToSign?: string;
    /**
     * The signed permissions for the account SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update (u) and Process (p).
     */
    permissions: string | enums.storage.Permissions;
    /**
     * The protocol permitted for a request made with the account SAS.
     */
    protocols?: enums.storage.HttpProtocol;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The signed resource types that are accessible with the account SAS. Service (s): Access to service-level APIs; Container (c): Access to container-level APIs; Object (o): Access to object-level APIs for blobs, queue messages, table entities, and files.
     */
    resourceTypes: string | enums.storage.SignedResourceTypes;
    /**
     * The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t), File (f).
     */
    services: string | enums.storage.Services;
    /**
     * The time at which the shared access signature becomes invalid.
     */
    sharedAccessExpiryTime: string;
    /**
     * The time at which the SAS becomes valid.
     */
    sharedAccessStartTime?: string;
}

/**
 * The List SAS credentials operation response.
 */
export interface ListStorageAccountSASResult {
    /**
     * List SAS credentials of storage account.
     */
    readonly accountSasToken: string;
}
/**
 * List SAS credentials of a storage account.
 *
 * Uses Azure REST API version 2024-01-01.
 *
 * Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listStorageAccountSASOutput(args: ListStorageAccountSASOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListStorageAccountSASResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:storage:listStorageAccountSAS", {
        "accountName": args.accountName,
        "iPAddressOrRange": args.iPAddressOrRange,
        "keyToSign": args.keyToSign,
        "permissions": args.permissions,
        "protocols": args.protocols,
        "resourceGroupName": args.resourceGroupName,
        "resourceTypes": args.resourceTypes,
        "services": args.services,
        "sharedAccessExpiryTime": args.sharedAccessExpiryTime,
        "sharedAccessStartTime": args.sharedAccessStartTime,
    }, opts);
}

export interface ListStorageAccountSASOutputArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    accountName: pulumi.Input<string>;
    /**
     * An IP address or a range of IP addresses from which to accept requests.
     */
    iPAddressOrRange?: pulumi.Input<string>;
    /**
     * The key to sign the account SAS token with.
     */
    keyToSign?: pulumi.Input<string>;
    /**
     * The signed permissions for the account SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update (u) and Process (p).
     */
    permissions: pulumi.Input<string | enums.storage.Permissions>;
    /**
     * The protocol permitted for a request made with the account SAS.
     */
    protocols?: pulumi.Input<enums.storage.HttpProtocol>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The signed resource types that are accessible with the account SAS. Service (s): Access to service-level APIs; Container (c): Access to container-level APIs; Object (o): Access to object-level APIs for blobs, queue messages, table entities, and files.
     */
    resourceTypes: pulumi.Input<string | enums.storage.SignedResourceTypes>;
    /**
     * The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t), File (f).
     */
    services: pulumi.Input<string | enums.storage.Services>;
    /**
     * The time at which the shared access signature becomes invalid.
     */
    sharedAccessExpiryTime: pulumi.Input<string>;
    /**
     * The time at which the SAS becomes valid.
     */
    sharedAccessStartTime?: pulumi.Input<string>;
}
