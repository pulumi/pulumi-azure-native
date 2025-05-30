// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Container App.
 *
 * Uses Azure REST API version 2024-03-01.
 *
 * Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getContainerApp(args: GetContainerAppArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerAppResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:app:getContainerApp", {
        "containerAppName": args.containerAppName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetContainerAppArgs {
    /**
     * Name of the Container App.
     */
    containerAppName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Container App.
 */
export interface GetContainerAppResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Non versioned Container App configuration properties.
     */
    readonly configuration?: outputs.app.ConfigurationResponse;
    /**
     * Id used to verify domain name ownership
     */
    readonly customDomainVerificationId: string;
    /**
     * Resource ID of environment.
     */
    readonly environmentId?: string;
    /**
     * The endpoint of the eventstream of the container app.
     */
    readonly eventStreamEndpoint: string;
    /**
     * The complex type of the extended location.
     */
    readonly extendedLocation?: outputs.app.ExtendedLocationResponse;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
     */
    readonly identity?: outputs.app.ManagedServiceIdentityResponse;
    /**
     * Name of the latest ready revision of the Container App.
     */
    readonly latestReadyRevisionName: string;
    /**
     * Fully Qualified Domain Name of the latest revision of the Container App.
     */
    readonly latestRevisionFqdn: string;
    /**
     * Name of the latest revision of the Container App.
     */
    readonly latestRevisionName: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
     */
    readonly managedBy?: string;
    /**
     * Deprecated. Resource ID of the Container App's environment.
     */
    readonly managedEnvironmentId?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Outbound IP Addresses for container app.
     */
    readonly outboundIpAddresses: string[];
    /**
     * Provisioning state of the Container App.
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.app.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Container App versioned application definition.
     */
    readonly template?: outputs.app.TemplateResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Workload profile name to pin for container app execution.
     */
    readonly workloadProfileName?: string;
}
/**
 * Container App.
 *
 * Uses Azure REST API version 2024-03-01.
 *
 * Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getContainerAppOutput(args: GetContainerAppOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetContainerAppResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:app:getContainerApp", {
        "containerAppName": args.containerAppName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetContainerAppOutputArgs {
    /**
     * Name of the Container App.
     */
    containerAppName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
