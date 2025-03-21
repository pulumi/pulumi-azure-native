// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the properties of the specified container group in the specified subscription and resource group. The operation returns the properties of each container group including containers, image registry credentials, restart policy, IP address type, OS type, state, and volumes.
 */
export function getContainerGroup(args: GetContainerGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:containerinstance/v20240501preview:getContainerGroup", {
        "containerGroupName": args.containerGroupName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetContainerGroupArgs {
    /**
     * The name of the container group.
     */
    containerGroupName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A container group.
 */
export interface GetContainerGroupResult {
    /**
     * The properties for confidential container group
     */
    readonly confidentialComputeProperties?: outputs.containerinstance.v20240501preview.ConfidentialComputePropertiesResponse;
    /**
     * The reference container group profile properties.
     */
    readonly containerGroupProfile?: outputs.containerinstance.v20240501preview.ContainerGroupProfileReferenceDefinitionResponse;
    /**
     * The containers within the container group.
     */
    readonly containers: outputs.containerinstance.v20240501preview.ContainerResponse[];
    /**
     * The diagnostic information for a container group.
     */
    readonly diagnostics?: outputs.containerinstance.v20240501preview.ContainerGroupDiagnosticsResponse;
    /**
     * The DNS config information for a container group.
     */
    readonly dnsConfig?: outputs.containerinstance.v20240501preview.DnsConfigurationResponse;
    /**
     * The encryption properties for a container group.
     */
    readonly encryptionProperties?: outputs.containerinstance.v20240501preview.EncryptionPropertiesResponse;
    /**
     * extensions used by virtual kubelet
     */
    readonly extensions?: outputs.containerinstance.v20240501preview.DeploymentExtensionSpecResponse[];
    /**
     * The resource id.
     */
    readonly id: string;
    /**
     * The identity of the container group, if configured.
     */
    readonly identity?: outputs.containerinstance.v20240501preview.ContainerGroupIdentityResponse;
    /**
     * The image registry credentials by which the container group is created from.
     */
    readonly imageRegistryCredentials?: outputs.containerinstance.v20240501preview.ImageRegistryCredentialResponse[];
    /**
     * The init containers for a container group.
     */
    readonly initContainers?: outputs.containerinstance.v20240501preview.InitContainerDefinitionResponse[];
    /**
     * The instance view of the container group. Only valid in response.
     */
    readonly instanceView: outputs.containerinstance.v20240501preview.ContainerGroupPropertiesResponseInstanceView;
    /**
     * The IP address type of the container group.
     */
    readonly ipAddress?: outputs.containerinstance.v20240501preview.IpAddressResponse;
    /**
     * The flag indicating whether the container group is created by standby pool.
     */
    readonly isCreatedFromStandbyPool: boolean;
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * The operating system type required by the containers in the container group.
     */
    readonly osType?: string;
    /**
     * The priority of the container group.
     */
    readonly priority?: string;
    /**
     * The provisioning state of the container group. This only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Restart policy for all containers within the container group. 
     * - `Always` Always restart
     * - `OnFailure` Restart on failure
     * - `Never` Never restart
     */
    readonly restartPolicy?: string;
    /**
     * The SKU for a container group.
     */
    readonly sku?: string;
    /**
     * The reference standby pool profile properties.
     */
    readonly standbyPoolProfile?: outputs.containerinstance.v20240501preview.StandbyPoolProfileDefinitionResponse;
    /**
     * The subnet resource IDs for a container group.
     */
    readonly subnetIds?: outputs.containerinstance.v20240501preview.ContainerGroupSubnetIdResponse[];
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The resource type.
     */
    readonly type: string;
    /**
     * The list of volumes that can be mounted by containers in this container group.
     */
    readonly volumes?: outputs.containerinstance.v20240501preview.VolumeResponse[];
    /**
     * The zones for the container group.
     */
    readonly zones?: string[];
}
/**
 * Gets the properties of the specified container group in the specified subscription and resource group. The operation returns the properties of each container group including containers, image registry credentials, restart policy, IP address type, OS type, state, and volumes.
 */
export function getContainerGroupOutput(args: GetContainerGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetContainerGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:containerinstance/v20240501preview:getContainerGroup", {
        "containerGroupName": args.containerGroupName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetContainerGroupOutputArgs {
    /**
     * The name of the container group.
     */
    containerGroupName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
