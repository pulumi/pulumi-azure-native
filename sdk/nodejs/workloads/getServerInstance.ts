// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the Server Instance resource.
 *
 * Uses Azure REST API version 2023-10-01-preview.
 */
export function getServerInstance(args: GetServerInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetServerInstanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:workloads:getServerInstance", {
        "resourceGroupName": args.resourceGroupName,
        "sapDiscoverySiteName": args.sapDiscoverySiteName,
        "sapInstanceName": args.sapInstanceName,
        "serverInstanceName": args.serverInstanceName,
    }, opts);
}

export interface GetServerInstanceArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the discovery site resource for SAP Migration.
     */
    sapDiscoverySiteName: string;
    /**
     * The name of SAP Instance resource for SAP Migration.
     */
    sapInstanceName: string;
    /**
     * The name of the Server instance resource for SAP Migration.
     */
    serverInstanceName: string;
}

/**
 * Define the Server Instance resource.
 */
export interface GetServerInstanceResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Configuration data for this server instance.
     */
    readonly configurationData: outputs.workloads.ConfigurationDataResponse;
    /**
     * Defines the errors related to SAP Instance resource.
     */
    readonly errors: outputs.workloads.SAPMigrateErrorResponse;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * This is the Instance SID for ASCS/AP/DB instance.  An SAP system with HANA database for example could have a different SID for database Instance than that of ASCS instance.
     */
    readonly instanceSid: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * This is Operating System on which the host server is running.
     */
    readonly operatingSystem: string;
    /**
     * Configuration data for this server instance.
     */
    readonly performanceData: outputs.workloads.ExcelPerformanceDataResponse | outputs.workloads.NativePerformanceDataResponse;
    /**
     * Defines the provisioning states.
     */
    readonly provisioningState: string;
    /**
     * Defines the type SAP instance on this server instance.
     */
    readonly sapInstanceType: string;
    /**
     * This is the SAP Application Component; e.g. SAP S/4HANA 2022, SAP ERP ENHANCE PACKAGE.
     */
    readonly sapProduct: string;
    /**
     * Provide the product version of the SAP product.
     */
    readonly sapProductVersion: string;
    /**
     * This is the Virtual Machine Name of the SAP system. Add all the virtual machines attached to an SAP system which you wish to migrate to Azure. Keeping this not equal to ID as for single tier all InstanceTypes would be on same server, leading to multiple resources with same servername.
     */
    readonly serverName: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.workloads.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets the Server Instance resource.
 *
 * Uses Azure REST API version 2023-10-01-preview.
 */
export function getServerInstanceOutput(args: GetServerInstanceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServerInstanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:workloads:getServerInstance", {
        "resourceGroupName": args.resourceGroupName,
        "sapDiscoverySiteName": args.sapDiscoverySiteName,
        "sapInstanceName": args.sapInstanceName,
        "serverInstanceName": args.serverInstanceName,
    }, opts);
}

export interface GetServerInstanceOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the discovery site resource for SAP Migration.
     */
    sapDiscoverySiteName: pulumi.Input<string>;
    /**
     * The name of SAP Instance resource for SAP Migration.
     */
    sapInstanceName: pulumi.Input<string>;
    /**
     * The name of the Server instance resource for SAP Migration.
     */
    serverInstanceName: pulumi.Input<string>;
}
