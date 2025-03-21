// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the properties of the connected registry.
 */
export function getConnectedRegistry(args: GetConnectedRegistryArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectedRegistryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:containerregistry/v20230101preview:getConnectedRegistry", {
        "connectedRegistryName": args.connectedRegistryName,
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConnectedRegistryArgs {
    /**
     * The name of the connected registry.
     */
    connectedRegistryName: string;
    /**
     * The name of the container registry.
     */
    registryName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * An object that represents a connected registry for a container registry.
 */
export interface GetConnectedRegistryResult {
    /**
     * The activation properties of the connected registry.
     */
    readonly activation: outputs.containerregistry.v20230101preview.ActivationPropertiesResponse;
    /**
     * The list of the ACR token resource IDs used to authenticate clients to the connected registry.
     */
    readonly clientTokenIds?: string[];
    /**
     * The current connection state of the connected registry.
     */
    readonly connectionState: string;
    /**
     * The resource ID.
     */
    readonly id: string;
    /**
     * The last activity time of the connected registry.
     */
    readonly lastActivityTime: string;
    /**
     * The logging properties of the connected registry.
     */
    readonly logging?: outputs.containerregistry.v20230101preview.LoggingPropertiesResponse;
    /**
     * The login server properties of the connected registry.
     */
    readonly loginServer?: outputs.containerregistry.v20230101preview.LoginServerPropertiesResponse;
    /**
     * The mode of the connected registry resource that indicates the permissions of the registry.
     */
    readonly mode: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The list of notifications subscription information for the connected registry.
     */
    readonly notificationsList?: string[];
    /**
     * The parent of the connected registry.
     */
    readonly parent: outputs.containerregistry.v20230101preview.ParentPropertiesResponse;
    /**
     * Provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * The list of current statuses of the connected registry.
     */
    readonly statusDetails: outputs.containerregistry.v20230101preview.StatusDetailPropertiesResponse[];
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.containerregistry.v20230101preview.SystemDataResponse;
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * The current version of ACR runtime on the connected registry.
     */
    readonly version: string;
}
/**
 * Gets the properties of the connected registry.
 */
export function getConnectedRegistryOutput(args: GetConnectedRegistryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConnectedRegistryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:containerregistry/v20230101preview:getConnectedRegistry", {
        "connectedRegistryName": args.connectedRegistryName,
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConnectedRegistryOutputArgs {
    /**
     * The name of the connected registry.
     */
    connectedRegistryName: pulumi.Input<string>;
    /**
     * The name of the container registry.
     */
    registryName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
