// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Implements VirtualMachine GET method.
 *
 * Uses Azure REST API version 2023-04-01-preview.
 *
 * Other available API versions: 2022-05-21-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getVirtualMachine(args: GetVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:scvmm:getVirtualMachine", {
        "resourceGroupName": args.resourceGroupName,
        "virtualMachineName": args.virtualMachineName,
    }, opts);
}

export interface GetVirtualMachineArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * Name of the VirtualMachine.
     */
    virtualMachineName: string;
}

/**
 * The VirtualMachines resource definition.
 */
export interface GetVirtualMachineResult {
    /**
     * Availability Sets in vm.
     */
    readonly availabilitySets?: outputs.scvmm.VirtualMachinePropertiesResponseAvailabilitySets[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Type of checkpoint supported for the vm.
     */
    readonly checkpointType?: string;
    /**
     * Checkpoints in the vm.
     */
    readonly checkpoints?: outputs.scvmm.CheckpointResponse[];
    /**
     * ARM Id of the cloud resource to use for deploying the vm.
     */
    readonly cloudId?: string;
    /**
     * The extended location.
     */
    readonly extendedLocation: outputs.scvmm.ExtendedLocationResponse;
    /**
     * Gets or sets the generation for the vm.
     */
    readonly generation?: number;
    /**
     * Guest agent status properties.
     */
    readonly guestAgentProfile?: outputs.scvmm.GuestAgentProfileResponse;
    /**
     * Hardware properties.
     */
    readonly hardwareProfile?: outputs.scvmm.HardwareProfileResponse;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.scvmm.IdentityResponse;
    /**
     * Gets or sets the inventory Item ID for the resource.
     */
    readonly inventoryItemId?: string;
    /**
     * Last restored checkpoint in the vm.
     */
    readonly lastRestoredVMCheckpoint: outputs.scvmm.CheckpointResponse;
    /**
     * Gets or sets the location.
     */
    readonly location: string;
    /**
     * Resource Name
     */
    readonly name: string;
    /**
     * Network properties.
     */
    readonly networkProfile?: outputs.scvmm.NetworkProfileResponse;
    /**
     * OS properties.
     */
    readonly osProfile?: outputs.scvmm.OsProfileResponse;
    /**
     * Gets the power state of the virtual machine.
     */
    readonly powerState: string;
    /**
     * Gets or sets the provisioning state.
     */
    readonly provisioningState: string;
    /**
     * Storage properties.
     */
    readonly storageProfile?: outputs.scvmm.StorageProfileResponse;
    /**
     * The system data.
     */
    readonly systemData: outputs.scvmm.SystemDataResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * ARM Id of the template resource to use for deploying the vm.
     */
    readonly templateId?: string;
    /**
     * Resource Type
     */
    readonly type: string;
    /**
     * Unique ID of the virtual machine.
     */
    readonly uuid?: string;
    /**
     * VMName is the name of VM on the SCVMM server.
     */
    readonly vmName?: string;
    /**
     * ARM Id of the vmmServer resource in which this resource resides.
     */
    readonly vmmServerId?: string;
}
/**
 * Implements VirtualMachine GET method.
 *
 * Uses Azure REST API version 2023-04-01-preview.
 *
 * Other available API versions: 2022-05-21-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getVirtualMachineOutput(args: GetVirtualMachineOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualMachineResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:scvmm:getVirtualMachine", {
        "resourceGroupName": args.resourceGroupName,
        "virtualMachineName": args.virtualMachineName,
    }, opts);
}

export interface GetVirtualMachineOutputArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the VirtualMachine.
     */
    virtualMachineName: pulumi.Input<string>;
}
