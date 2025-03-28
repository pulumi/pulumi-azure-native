// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a server.
 */
export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:sql/v20240501preview:getServer", {
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerArgs {
    /**
     * The child resources to include in the response.
     */
    expand?: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: string;
    /**
     * The name of the server.
     */
    serverName: string;
}

/**
 * An Azure SQL Database server.
 */
export interface GetServerResult {
    /**
     * Administrator username for the server. Once created it cannot be changed.
     */
    readonly administratorLogin?: string;
    /**
     * The Azure Active Directory administrator of the server. This can only be used at server create time. If used for server update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
     */
    readonly administrators?: outputs.sql.v20240501preview.ServerExternalAdministratorResponse;
    /**
     * Status of external governance.
     */
    readonly externalGovernanceStatus: string;
    /**
     * The Client id used for cross tenant CMK scenario
     */
    readonly federatedClientId?: string;
    /**
     * The fully qualified domain name of the server.
     */
    readonly fullyQualifiedDomainName: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * The Azure Active Directory identity of the server.
     */
    readonly identity?: outputs.sql.v20240501preview.ResourceIdentityResponse;
    /**
     * Whether or not to enable IPv6 support for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
     */
    readonly isIPv6Enabled?: string;
    /**
     * A CMK URI of the key to use for encryption.
     */
    readonly keyId?: string;
    /**
     * Kind of sql server. This is metadata used for the Azure portal experience.
     */
    readonly kind: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Minimal TLS version. Allowed values: 'None', 1.0', '1.1', '1.2', '1.3'
     */
    readonly minimalTlsVersion?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The resource id of a user assigned identity to be used by default.
     */
    readonly primaryUserAssignedIdentityId?: string;
    /**
     * List of private endpoint connections on a server
     */
    readonly privateEndpointConnections: outputs.sql.v20240501preview.ServerPrivateEndpointConnectionResponse[];
    /**
     * Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled' or 'SecuredByPerimeter'
     */
    readonly publicNetworkAccess?: string;
    /**
     * Whether or not to restrict outbound network access for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
     */
    readonly restrictOutboundNetworkAccess?: string;
    /**
     * The state of the server.
     */
    readonly state: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The version of the server.
     */
    readonly version?: string;
    /**
     * Whether or not existing server has a workspace created and if it allows connection from workspace
     */
    readonly workspaceFeature: string;
}
/**
 * Gets a server.
 */
export function getServerOutput(args: GetServerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:sql/v20240501preview:getServer", {
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerOutputArgs {
    /**
     * The child resources to include in the response.
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
}
