// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieves the details of a P2SVpnServerConfiguration.
 *
 * Uses Azure REST API version 2019-07-01.
 */
export function getP2sVpnServerConfiguration(args: GetP2sVpnServerConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetP2sVpnServerConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network:getP2sVpnServerConfiguration", {
        "p2SVpnServerConfigurationName": args.p2SVpnServerConfigurationName,
        "resourceGroupName": args.resourceGroupName,
        "virtualWanName": args.virtualWanName,
    }, opts);
}

export interface GetP2sVpnServerConfigurationArgs {
    /**
     * The name of the P2SVpnServerConfiguration.
     */
    p2SVpnServerConfigurationName: string;
    /**
     * The resource group name of the P2SVpnServerConfiguration.
     */
    resourceGroupName: string;
    /**
     * The name of the VirtualWan.
     */
    virtualWanName: string;
}

/**
 * P2SVpnServerConfiguration Resource.
 */
export interface GetP2sVpnServerConfigurationResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
     */
    readonly name?: string;
    /**
     * List of references to P2SVpnGateways.
     */
    readonly p2SVpnGateways: outputs.network.SubResourceResponse[];
    /**
     * Radius client root certificate of P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigRadiusClientRootCertificates?: outputs.network.P2SVpnServerConfigRadiusClientRootCertificateResponse[];
    /**
     * Radius Server root certificate of P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigRadiusServerRootCertificates?: outputs.network.P2SVpnServerConfigRadiusServerRootCertificateResponse[];
    /**
     * VPN client revoked certificate of P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigVpnClientRevokedCertificates?: outputs.network.P2SVpnServerConfigVpnClientRevokedCertificateResponse[];
    /**
     * VPN client root certificate of P2SVpnServerConfiguration.
     */
    readonly p2SVpnServerConfigVpnClientRootCertificates?: outputs.network.P2SVpnServerConfigVpnClientRootCertificateResponse[];
    /**
     * The provisioning state of the P2S VPN server configuration resource.
     */
    readonly provisioningState: string;
    /**
     * The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
     */
    readonly radiusServerAddress?: string;
    /**
     * The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
     */
    readonly radiusServerSecret?: string;
    /**
     * VpnClientIpsecPolicies for P2SVpnServerConfiguration.
     */
    readonly vpnClientIpsecPolicies?: outputs.network.IpsecPolicyResponse[];
    /**
     * VPN protocols for the P2SVpnServerConfiguration.
     */
    readonly vpnProtocols?: string[];
}
/**
 * Retrieves the details of a P2SVpnServerConfiguration.
 *
 * Uses Azure REST API version 2019-07-01.
 */
export function getP2sVpnServerConfigurationOutput(args: GetP2sVpnServerConfigurationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetP2sVpnServerConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network:getP2sVpnServerConfiguration", {
        "p2SVpnServerConfigurationName": args.p2SVpnServerConfigurationName,
        "resourceGroupName": args.resourceGroupName,
        "virtualWanName": args.virtualWanName,
    }, opts);
}

export interface GetP2sVpnServerConfigurationOutputArgs {
    /**
     * The name of the P2SVpnServerConfiguration.
     */
    p2SVpnServerConfigurationName: pulumi.Input<string>;
    /**
     * The resource group name of the P2SVpnServerConfiguration.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the VirtualWan.
     */
    virtualWanName: pulumi.Input<string>;
}
