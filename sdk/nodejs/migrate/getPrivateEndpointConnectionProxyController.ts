// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get the of private link proxy resources from a migrate project and private link proxy resource.
 *
 * Uses Azure REST API version 2023-01-01.
 */
export function getPrivateEndpointConnectionProxyController(args: GetPrivateEndpointConnectionProxyControllerArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateEndpointConnectionProxyControllerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:migrate:getPrivateEndpointConnectionProxyController", {
        "migrateProjectName": args.migrateProjectName,
        "pecProxyName": args.pecProxyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateEndpointConnectionProxyControllerArgs {
    /**
     * Name of the Azure Migrate project.
     */
    migrateProjectName: string;
    /**
     * Private link proxy name.
     */
    pecProxyName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Defines Private endpoint proxy resource.
 */
export interface GetPrivateEndpointConnectionProxyControllerResult {
    readonly eTag?: string;
    readonly id: string;
    readonly name: string;
    /**
     * Properties of a private endpoint connection proxy.
     */
    readonly properties: outputs.migrate.PrivateEndpointConnectionProxyPropertiesResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.migrate.SystemDataResponse;
    readonly type: string;
}
/**
 * Get the of private link proxy resources from a migrate project and private link proxy resource.
 *
 * Uses Azure REST API version 2023-01-01.
 */
export function getPrivateEndpointConnectionProxyControllerOutput(args: GetPrivateEndpointConnectionProxyControllerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPrivateEndpointConnectionProxyControllerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:migrate:getPrivateEndpointConnectionProxyController", {
        "migrateProjectName": args.migrateProjectName,
        "pecProxyName": args.pecProxyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateEndpointConnectionProxyControllerOutputArgs {
    /**
     * Name of the Azure Migrate project.
     */
    migrateProjectName: pulumi.Input<string>;
    /**
     * Private link proxy name.
     */
    pecProxyName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
