// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the peer ASN with the specified name under the given subscription.
 *
 * Uses Azure REST API version 2022-10-01.
 *
 * Other available API versions: 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native peering [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPeerAsn(args: GetPeerAsnArgs, opts?: pulumi.InvokeOptions): Promise<GetPeerAsnResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:peering:getPeerAsn", {
        "peerAsnName": args.peerAsnName,
    }, opts);
}

export interface GetPeerAsnArgs {
    /**
     * The peer ASN name.
     */
    peerAsnName: string;
}

/**
 * The essential information related to the peer's ASN.
 */
export interface GetPeerAsnResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The error message for the validation state
     */
    readonly errorMessage: string;
    /**
     * The ID of the resource.
     */
    readonly id: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The Autonomous System Number (ASN) of the peer.
     */
    readonly peerAsn?: number;
    /**
     * The contact details of the peer.
     */
    readonly peerContactDetail?: outputs.peering.ContactDetailResponse[];
    /**
     * The name of the peer.
     */
    readonly peerName?: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * The validation state of the ASN associated with the peer.
     */
    readonly validationState: string;
}
/**
 * Gets the peer ASN with the specified name under the given subscription.
 *
 * Uses Azure REST API version 2022-10-01.
 *
 * Other available API versions: 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native peering [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPeerAsnOutput(args: GetPeerAsnOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPeerAsnResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:peering:getPeerAsn", {
        "peerAsnName": args.peerAsnName,
    }, opts);
}

export interface GetPeerAsnOutputArgs {
    /**
     * The peer ASN name.
     */
    peerAsnName: pulumi.Input<string>;
}
