// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a host pool.
 *
 * Uses Azure REST API version 2024-04-03.
 *
 * Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getHostPool(args: GetHostPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetHostPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:desktopvirtualization:getHostPool", {
        "hostPoolName": args.hostPoolName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHostPoolArgs {
    /**
     * The name of the host pool within the specified resource group
     */
    hostPoolName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Represents a HostPool definition.
 */
export interface GetHostPoolResult {
    /**
     * The session host configuration for updating agent, monitoring agent, and stack component.
     */
    readonly agentUpdate?: outputs.desktopvirtualization.AgentUpdatePropertiesResponse;
    /**
     * List of App Attach Package links.
     */
    readonly appAttachPackageReferences: string[];
    /**
     * List of applicationGroup links.
     */
    readonly applicationGroupReferences: string[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Is cloud pc resource.
     */
    readonly cloudPcResource: boolean;
    /**
     * Custom rdp property of HostPool.
     */
    readonly customRdpProperty?: string;
    /**
     * Description of HostPool.
     */
    readonly description?: string;
    /**
     * The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. 
     */
    readonly etag: string;
    /**
     * Friendly name of HostPool.
     */
    readonly friendlyName?: string;
    /**
     * HostPool type for desktop.
     */
    readonly hostPoolType: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    readonly identity?: outputs.desktopvirtualization.ResourceModelWithAllowedPropertySetResponseIdentity;
    /**
     * Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type. E.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
     */
    readonly kind?: string;
    /**
     * The type of the load balancer.
     */
    readonly loadBalancerType: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
     */
    readonly managedBy?: string;
    /**
     * The max session limit of HostPool.
     */
    readonly maxSessionLimit?: number;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * ObjectId of HostPool. (internal use)
     */
    readonly objectId: string;
    /**
     * PersonalDesktopAssignment type for HostPool.
     */
    readonly personalDesktopAssignmentType?: string;
    readonly plan?: outputs.desktopvirtualization.ResourceModelWithAllowedPropertySetResponsePlan;
    /**
     * The type of preferred application group type, default to Desktop Application Group
     */
    readonly preferredAppGroupType: string;
    /**
     * List of private endpoint connection associated with the specified resource
     */
    readonly privateEndpointConnections: outputs.desktopvirtualization.PrivateEndpointConnectionResponse[];
    /**
     * Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
     */
    readonly publicNetworkAccess?: string;
    /**
     * The registration info of HostPool.
     */
    readonly registrationInfo?: outputs.desktopvirtualization.RegistrationInfoResponse;
    /**
     * The ring number of HostPool.
     */
    readonly ring?: number;
    readonly sku?: outputs.desktopvirtualization.ResourceModelWithAllowedPropertySetResponseSku;
    /**
     * ClientId for the registered Relying Party used to issue WVD SSO certificates.
     */
    readonly ssoClientId?: string;
    /**
     * Path to Azure KeyVault storing the secret used for communication to ADFS.
     */
    readonly ssoClientSecretKeyVaultPath?: string;
    /**
     * The type of single sign on Secret Type.
     */
    readonly ssoSecretType?: string;
    /**
     * URL to customer ADFS server for signing WVD SSO certificates.
     */
    readonly ssoadfsAuthority?: string;
    /**
     * The flag to turn on/off StartVMOnConnect feature.
     */
    readonly startVMOnConnect?: boolean;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.desktopvirtualization.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Is validation environment.
     */
    readonly validationEnvironment?: boolean;
    /**
     * VM template for sessionhosts configuration within hostpool.
     */
    readonly vmTemplate?: string;
}
/**
 * Get a host pool.
 *
 * Uses Azure REST API version 2024-04-03.
 *
 * Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getHostPoolOutput(args: GetHostPoolOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetHostPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:desktopvirtualization:getHostPool", {
        "hostPoolName": args.hostPoolName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHostPoolOutputArgs {
    /**
     * The name of the host pool within the specified resource group
     */
    hostPoolName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
