// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Implements VirtualMachineTemplate GET method.
 *
 * Uses Azure REST API version 2023-04-01-preview.
 *
 * Other available API versions: 2022-05-21-preview, 2023-10-07, 2024-06-01, 2025-03-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getVirtualMachineTemplate(args: GetVirtualMachineTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:scvmm:getVirtualMachineTemplate", {
        "resourceGroupName": args.resourceGroupName,
        "virtualMachineTemplateName": args.virtualMachineTemplateName,
    }, opts);
}

export interface GetVirtualMachineTemplateArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * Name of the VirtualMachineTemplate.
     */
    virtualMachineTemplateName: string;
}

/**
 * The VirtualMachineTemplates resource definition.
 */
export interface GetVirtualMachineTemplateResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Gets or sets computer name.
     */
    readonly computerName: string;
    /**
     * Gets or sets the desired number of vCPUs for the vm.
     */
    readonly cpuCount: number;
    /**
     * Gets or sets the disks of the template.
     */
    readonly disks: outputs.scvmm.VirtualDiskResponse[];
    /**
     * Gets or sets a value indicating whether to enable dynamic memory or not.
     */
    readonly dynamicMemoryEnabled: string;
    /**
     * Gets or sets the max dynamic memory for the vm.
     */
    readonly dynamicMemoryMaxMB: number;
    /**
     * Gets or sets the min dynamic memory for the vm.
     */
    readonly dynamicMemoryMinMB: number;
    /**
     * The extended location.
     */
    readonly extendedLocation: outputs.scvmm.ExtendedLocationResponse;
    /**
     * Gets or sets the generation for the vm.
     */
    readonly generation: number;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Gets or sets the inventory Item ID for the resource.
     */
    readonly inventoryItemId?: string;
    /**
     * Gets or sets a value indicating whether the vm template is customizable or not.
     */
    readonly isCustomizable: string;
    /**
     * Gets highly available property.
     */
    readonly isHighlyAvailable: string;
    /**
     * Gets or sets a value indicating whether to enable processor compatibility mode for live migration of VMs.
     */
    readonly limitCpuForMigration: string;
    /**
     * Gets or sets the location.
     */
    readonly location: string;
    /**
     * MemoryMB is the desired size of a virtual machine's memory, in MB.
     */
    readonly memoryMB: number;
    /**
     * Resource Name
     */
    readonly name: string;
    /**
     * Gets or sets the network interfaces of the template.
     */
    readonly networkInterfaces: outputs.scvmm.NetworkInterfacesResponse[];
    /**
     * Gets or sets os name.
     */
    readonly osName: string;
    /**
     * Gets or sets the type of the os.
     */
    readonly osType: string;
    /**
     * Gets or sets the provisioning state.
     */
    readonly provisioningState: string;
    /**
     * The system data.
     */
    readonly systemData: outputs.scvmm.SystemDataResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource Type
     */
    readonly type: string;
    /**
     * Unique ID of the virtual machine template.
     */
    readonly uuid?: string;
    /**
     * ARM Id of the vmmServer resource in which this resource resides.
     */
    readonly vmmServerId?: string;
}
/**
 * Implements VirtualMachineTemplate GET method.
 *
 * Uses Azure REST API version 2023-04-01-preview.
 *
 * Other available API versions: 2022-05-21-preview, 2023-10-07, 2024-06-01, 2025-03-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getVirtualMachineTemplateOutput(args: GetVirtualMachineTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualMachineTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:scvmm:getVirtualMachineTemplate", {
        "resourceGroupName": args.resourceGroupName,
        "virtualMachineTemplateName": args.virtualMachineTemplateName,
    }, opts);
}

export interface GetVirtualMachineTemplateOutputArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the VirtualMachineTemplate.
     */
    virtualMachineTemplateName: pulumi.Input<string>;
}
