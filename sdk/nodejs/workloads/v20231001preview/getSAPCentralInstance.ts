// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the SAP Central Services Instance resource.
 */
export function getSAPCentralInstance(args: GetSAPCentralInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetSAPCentralInstanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:workloads/v20231001preview:getSAPCentralInstance", {
        "centralInstanceName": args.centralInstanceName,
        "resourceGroupName": args.resourceGroupName,
        "sapVirtualInstanceName": args.sapVirtualInstanceName,
    }, opts);
}

export interface GetSAPCentralInstanceArgs {
    /**
     * Central Services Instance resource name string modeled as parameter for auto generation to work correctly.
     */
    centralInstanceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the Virtual Instances for SAP solutions resource
     */
    sapVirtualInstanceName: string;
}

/**
 * Define the SAP Central Services Instance resource.
 */
export interface GetSAPCentralInstanceResult {
    /**
     * Defines the SAP Enqueue Replication Server (ERS) properties.
     */
    readonly enqueueReplicationServerProperties?: outputs.workloads.v20231001preview.EnqueueReplicationServerPropertiesResponse;
    /**
     * Defines the SAP Enqueue Server properties.
     */
    readonly enqueueServerProperties?: outputs.workloads.v20231001preview.EnqueueServerPropertiesResponse;
    /**
     * Defines the errors related to SAP Central Services Instance resource.
     */
    readonly errors: outputs.workloads.v20231001preview.SAPVirtualInstanceErrorResponse;
    /**
     * Defines the SAP Gateway Server properties.
     */
    readonly gatewayServerProperties?: outputs.workloads.v20231001preview.GatewayServerPropertiesResponse;
    /**
     * Defines the health of SAP Instances.
     */
    readonly health: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The central services instance number.
     */
    readonly instanceNo: string;
    /**
     * The central services instance Kernel Patch level.
     */
    readonly kernelPatch: string;
    /**
     * The central services instance Kernel Version.
     */
    readonly kernelVersion: string;
    /**
     * The Load Balancer details such as LoadBalancer ID attached to ASCS Virtual Machines
     */
    readonly loadBalancerDetails: outputs.workloads.v20231001preview.LoadBalancerDetailsResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Defines the SAP message server properties.
     */
    readonly messageServerProperties?: outputs.workloads.v20231001preview.MessageServerPropertiesResponse;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Defines the provisioning states.
     */
    readonly provisioningState: string;
    /**
     * Defines the SAP Instance status.
     */
    readonly status: string;
    /**
     * The central services instance subnet.
     */
    readonly subnet: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.workloads.v20231001preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The list of virtual machines corresponding to the Central Services instance.
     */
    readonly vmDetails: outputs.workloads.v20231001preview.CentralServerVmDetailsResponse[];
}
/**
 * Gets the SAP Central Services Instance resource.
 */
export function getSAPCentralInstanceOutput(args: GetSAPCentralInstanceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSAPCentralInstanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:workloads/v20231001preview:getSAPCentralInstance", {
        "centralInstanceName": args.centralInstanceName,
        "resourceGroupName": args.resourceGroupName,
        "sapVirtualInstanceName": args.sapVirtualInstanceName,
    }, opts);
}

export interface GetSAPCentralInstanceOutputArgs {
    /**
     * Central Services Instance resource name string modeled as parameter for auto generation to work correctly.
     */
    centralInstanceName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Virtual Instances for SAP solutions resource
     */
    sapVirtualInstanceName: pulumi.Input<string>;
}
