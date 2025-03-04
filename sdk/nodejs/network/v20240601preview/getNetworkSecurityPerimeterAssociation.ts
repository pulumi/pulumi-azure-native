// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the specified NSP association by name.
 */
export function getNetworkSecurityPerimeterAssociation(args: GetNetworkSecurityPerimeterAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkSecurityPerimeterAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network/v20240601preview:getNetworkSecurityPerimeterAssociation", {
        "associationName": args.associationName,
        "networkSecurityPerimeterName": args.networkSecurityPerimeterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNetworkSecurityPerimeterAssociationArgs {
    /**
     * The name of the NSP association.
     */
    associationName: string;
    /**
     * The name of the network security perimeter.
     */
    networkSecurityPerimeterName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * The NSP resource association resource
 */
export interface GetNetworkSecurityPerimeterAssociationResult {
    /**
     * Access mode on the association.
     */
    readonly accessMode?: string;
    /**
     * Specifies if there are provisioning issues
     */
    readonly hasProvisioningIssues: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The PaaS resource to be associated.
     */
    readonly privateLinkResource?: outputs.network.v20240601preview.SubResourceResponse;
    /**
     * Profile id to which the PaaS resource is associated.
     */
    readonly profile?: outputs.network.v20240601preview.SubResourceResponse;
    /**
     * The provisioning state of the resource  association resource.
     */
    readonly provisioningState: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Gets the specified NSP association by name.
 */
export function getNetworkSecurityPerimeterAssociationOutput(args: GetNetworkSecurityPerimeterAssociationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNetworkSecurityPerimeterAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network/v20240601preview:getNetworkSecurityPerimeterAssociation", {
        "associationName": args.associationName,
        "networkSecurityPerimeterName": args.networkSecurityPerimeterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNetworkSecurityPerimeterAssociationOutputArgs {
    /**
     * The name of the NSP association.
     */
    associationName: pulumi.Input<string>;
    /**
     * The name of the network security perimeter.
     */
    networkSecurityPerimeterName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
