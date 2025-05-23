// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Implements virtual machine template GET method.
 *
 * Uses Azure REST API version 2023-12-01.
 *
 * Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getVirtualMachineTemplate(args: GetVirtualMachineTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:connectedvmwarevsphere:getVirtualMachineTemplate", {
        "resourceGroupName": args.resourceGroupName,
        "virtualMachineTemplateName": args.virtualMachineTemplateName,
    }, opts);
}

export interface GetVirtualMachineTemplateArgs {
    /**
     * The Resource Group Name.
     */
    resourceGroupName: string;
    /**
     * Name of the virtual machine template resource.
     */
    virtualMachineTemplateName: string;
}

/**
 * Define the virtualMachineTemplate.
 */
export interface GetVirtualMachineTemplateResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Gets the name of the corresponding resource in Kubernetes.
     */
    readonly customResourceName: string;
    /**
     * Gets or sets the disks the template.
     */
    readonly disks: outputs.connectedvmwarevsphere.VirtualDiskResponse[];
    /**
     * Gets or sets the extended location.
     */
    readonly extendedLocation?: outputs.connectedvmwarevsphere.ExtendedLocationResponse;
    /**
     * Firmware type
     */
    readonly firmwareType: string;
    /**
     * Gets or sets the folder path of the template.
     */
    readonly folderPath: string;
    /**
     * Gets or sets the Id.
     */
    readonly id: string;
    /**
     * Gets or sets the inventory Item ID for the virtual machine template.
     */
    readonly inventoryItemId?: string;
    /**
     * Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
     */
    readonly kind?: string;
    /**
     * Gets or sets the location.
     */
    readonly location: string;
    /**
     * Gets or sets memory size in MBs for the template.
     */
    readonly memorySizeMB: number;
    /**
     * Gets or sets the vCenter Managed Object name for the virtual machine template.
     */
    readonly moName: string;
    /**
     * Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual machine
     * template.
     */
    readonly moRefId?: string;
    /**
     * Gets or sets the name.
     */
    readonly name: string;
    /**
     * Gets or sets the network interfaces of the template.
     */
    readonly networkInterfaces: outputs.connectedvmwarevsphere.NetworkInterfaceResponse[];
    /**
     * Gets or sets the number of vCPUs for the template.
     */
    readonly numCPUs: number;
    /**
     * Gets or sets the number of cores per socket for the template.
     * Defaults to 1 if unspecified.
     */
    readonly numCoresPerSocket: number;
    /**
     * Gets or sets os name.
     */
    readonly osName: string;
    /**
     * Gets or sets the type of the os.
     */
    readonly osType: string;
    /**
     * Gets the provisioning state.
     */
    readonly provisioningState: string;
    /**
     * The resource status information.
     */
    readonly statuses: outputs.connectedvmwarevsphere.ResourceStatusResponse[];
    /**
     * The system data.
     */
    readonly systemData: outputs.connectedvmwarevsphere.SystemDataResponse;
    /**
     * Gets or sets the Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Gets or sets the current version of VMware Tools.
     */
    readonly toolsVersion: string;
    /**
     * Gets or sets the current version status of VMware Tools installed in the guest operating system.
     */
    readonly toolsVersionStatus: string;
    /**
     * Gets or sets the type of the resource.
     */
    readonly type: string;
    /**
     * Gets or sets a unique identifier for this resource.
     */
    readonly uuid: string;
    /**
     * Gets or sets the ARM Id of the vCenter resource in which this template resides.
     */
    readonly vCenterId?: string;
}
/**
 * Implements virtual machine template GET method.
 *
 * Uses Azure REST API version 2023-12-01.
 *
 * Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getVirtualMachineTemplateOutput(args: GetVirtualMachineTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualMachineTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:connectedvmwarevsphere:getVirtualMachineTemplate", {
        "resourceGroupName": args.resourceGroupName,
        "virtualMachineTemplateName": args.virtualMachineTemplateName,
    }, opts);
}

export interface GetVirtualMachineTemplateOutputArgs {
    /**
     * The Resource Group Name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the virtual machine template resource.
     */
    virtualMachineTemplateName: pulumi.Input<string>;
}
