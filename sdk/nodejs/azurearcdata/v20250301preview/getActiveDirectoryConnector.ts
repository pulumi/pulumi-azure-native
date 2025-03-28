// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves an Active Directory connector resource
 */
export function getActiveDirectoryConnector(args: GetActiveDirectoryConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetActiveDirectoryConnectorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:azurearcdata/v20250301preview:getActiveDirectoryConnector", {
        "activeDirectoryConnectorName": args.activeDirectoryConnectorName,
        "dataControllerName": args.dataControllerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetActiveDirectoryConnectorArgs {
    /**
     * The name of the Active Directory connector instance
     */
    activeDirectoryConnectorName: string;
    /**
     * The name of the data controller
     */
    dataControllerName: string;
    /**
     * The name of the Azure resource group
     */
    resourceGroupName: string;
}

/**
 * Active directory connector resource
 */
export interface GetActiveDirectoryConnectorResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * null
     */
    readonly properties: outputs.azurearcdata.v20250301preview.ActiveDirectoryConnectorPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.azurearcdata.v20250301preview.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Retrieves an Active Directory connector resource
 */
export function getActiveDirectoryConnectorOutput(args: GetActiveDirectoryConnectorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetActiveDirectoryConnectorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:azurearcdata/v20250301preview:getActiveDirectoryConnector", {
        "activeDirectoryConnectorName": args.activeDirectoryConnectorName,
        "dataControllerName": args.dataControllerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetActiveDirectoryConnectorOutputArgs {
    /**
     * The name of the Active Directory connector instance
     */
    activeDirectoryConnectorName: pulumi.Input<string>;
    /**
     * The name of the data controller
     */
    dataControllerName: pulumi.Input<string>;
    /**
     * The name of the Azure resource group
     */
    resourceGroupName: pulumi.Input<string>;
}
