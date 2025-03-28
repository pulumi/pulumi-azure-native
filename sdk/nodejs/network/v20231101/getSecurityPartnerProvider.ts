// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the specified Security Partner Provider.
 */
export function getSecurityPartnerProvider(args: GetSecurityPartnerProviderArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityPartnerProviderResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network/v20231101:getSecurityPartnerProvider", {
        "resourceGroupName": args.resourceGroupName,
        "securityPartnerProviderName": args.securityPartnerProviderName,
    }, opts);
}

export interface GetSecurityPartnerProviderArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * The name of the Security Partner Provider.
     */
    securityPartnerProviderName: string;
}

/**
 * Security Partner Provider resource.
 */
export interface GetSecurityPartnerProviderResult {
    /**
     * The connection status with the Security Partner Provider.
     */
    readonly connectionStatus: string;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the Security Partner Provider resource.
     */
    readonly provisioningState: string;
    /**
     * The security provider name.
     */
    readonly securityProviderName?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The virtualHub to which the Security Partner Provider belongs.
     */
    readonly virtualHub?: outputs.network.v20231101.SubResourceResponse;
}
/**
 * Gets the specified Security Partner Provider.
 */
export function getSecurityPartnerProviderOutput(args: GetSecurityPartnerProviderOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSecurityPartnerProviderResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network/v20231101:getSecurityPartnerProvider", {
        "resourceGroupName": args.resourceGroupName,
        "securityPartnerProviderName": args.securityPartnerProviderName,
    }, opts);
}

export interface GetSecurityPartnerProviderOutputArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Security Partner Provider.
     */
    securityPartnerProviderName: pulumi.Input<string>;
}
