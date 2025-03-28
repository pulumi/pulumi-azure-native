// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The services resource is the top-level resource that represents the Database Migration Service. The GET method retrieves information about a service instance.
 *
 * Uses Azure REST API version 2021-06-30.
 *
 * Other available API versions: 2022-03-30-preview, 2023-07-15-preview.
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:datamigration:getService", {
        "groupName": args.groupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetServiceArgs {
    /**
     * Name of the resource group
     */
    groupName: string;
    /**
     * Name of the service
     */
    serviceName: string;
}

/**
 * A Database Migration Service resource
 */
export interface GetServiceResult {
    /**
     * HTTP strong entity tag value. Ignored if submitted
     */
    readonly etag?: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * The resource kind. Only 'vm' (the default) is supported.
     */
    readonly kind?: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The resource's provisioning state
     */
    readonly provisioningState: string;
    /**
     * The public key of the service, used to encrypt secrets sent to the service
     */
    readonly publicKey?: string;
    /**
     * Service SKU
     */
    readonly sku?: outputs.datamigration.ServiceSkuResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.datamigration.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The ID of the Microsoft.Network/networkInterfaces resource which the service have
     */
    readonly virtualNicId?: string;
    /**
     * The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
     */
    readonly virtualSubnetId: string;
}
/**
 * The services resource is the top-level resource that represents the Database Migration Service. The GET method retrieves information about a service instance.
 *
 * Uses Azure REST API version 2021-06-30.
 *
 * Other available API versions: 2022-03-30-preview, 2023-07-15-preview.
 */
export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:datamigration:getService", {
        "groupName": args.groupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetServiceOutputArgs {
    /**
     * Name of the resource group
     */
    groupName: pulumi.Input<string>;
    /**
     * Name of the service
     */
    serviceName: pulumi.Input<string>;
}
