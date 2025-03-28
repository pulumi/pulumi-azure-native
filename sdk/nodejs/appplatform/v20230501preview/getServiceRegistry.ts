// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get the Service Registry and its properties.
 */
export function getServiceRegistry(args: GetServiceRegistryArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceRegistryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:appplatform/v20230501preview:getServiceRegistry", {
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "serviceRegistryName": args.serviceRegistryName,
    }, opts);
}

export interface GetServiceRegistryArgs {
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: string;
    /**
     * The name of the Service resource.
     */
    serviceName: string;
    /**
     * The name of Service Registry.
     */
    serviceRegistryName: string;
}

/**
 * Service Registry resource
 */
export interface GetServiceRegistryResult {
    /**
     * Fully qualified resource Id for the resource.
     */
    readonly id: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * Service Registry properties payload
     */
    readonly properties: outputs.appplatform.v20230501preview.ServiceRegistryPropertiesResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.appplatform.v20230501preview.SystemDataResponse;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
/**
 * Get the Service Registry and its properties.
 */
export function getServiceRegistryOutput(args: GetServiceRegistryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceRegistryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:appplatform/v20230501preview:getServiceRegistry", {
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "serviceRegistryName": args.serviceRegistryName,
    }, opts);
}

export interface GetServiceRegistryOutputArgs {
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Service resource.
     */
    serviceName: pulumi.Input<string>;
    /**
     * The name of Service Registry.
     */
    serviceRegistryName: pulumi.Input<string>;
}
