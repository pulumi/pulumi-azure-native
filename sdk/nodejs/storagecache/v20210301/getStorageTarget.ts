// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns a Storage Target from a Cache.
 */
export function getStorageTarget(args: GetStorageTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageTargetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:storagecache/v20210301:getStorageTarget", {
        "cacheName": args.cacheName,
        "resourceGroupName": args.resourceGroupName,
        "storageTargetName": args.storageTargetName,
    }, opts);
}

export interface GetStorageTargetArgs {
    /**
     * Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
     */
    cacheName: string;
    /**
     * Target resource group.
     */
    resourceGroupName: string;
    /**
     * Name of Storage Target.
     */
    storageTargetName: string;
}

/**
 * Type of the Storage Target.
 */
export interface GetStorageTargetResult {
    /**
     * Properties when targetType is blobNfs.
     */
    readonly blobNfs?: outputs.storagecache.v20210301.BlobNfsTargetResponse;
    /**
     * Properties when targetType is clfs.
     */
    readonly clfs?: outputs.storagecache.v20210301.ClfsTargetResponse;
    /**
     * Resource ID of the Storage Target.
     */
    readonly id: string;
    /**
     * List of Cache namespace junctions to target for namespace associations.
     */
    readonly junctions?: outputs.storagecache.v20210301.NamespaceJunctionResponse[];
    /**
     * Region name string.
     */
    readonly location: string;
    /**
     * Name of the Storage Target.
     */
    readonly name: string;
    /**
     * Properties when targetType is nfs3.
     */
    readonly nfs3?: outputs.storagecache.v20210301.Nfs3TargetResponse;
    /**
     * ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
     */
    readonly provisioningState?: string;
    /**
     * The system meta data relating to this resource.
     */
    readonly systemData: outputs.storagecache.v20210301.SystemDataResponse;
    /**
     * Type of the Storage Target.
     */
    readonly targetType: string;
    /**
     * Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
     */
    readonly type: string;
    /**
     * Properties when targetType is unknown.
     */
    readonly unknown?: outputs.storagecache.v20210301.UnknownTargetResponse;
}
/**
 * Returns a Storage Target from a Cache.
 */
export function getStorageTargetOutput(args: GetStorageTargetOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStorageTargetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:storagecache/v20210301:getStorageTarget", {
        "cacheName": args.cacheName,
        "resourceGroupName": args.resourceGroupName,
        "storageTargetName": args.storageTargetName,
    }, opts);
}

export interface GetStorageTargetOutputArgs {
    /**
     * Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
     */
    cacheName: pulumi.Input<string>;
    /**
     * Target resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of Storage Target.
     */
    storageTargetName: pulumi.Input<string>;
}
