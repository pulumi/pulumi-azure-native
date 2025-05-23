// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get an Operator API Connection.
 *
 * Uses Azure REST API version 2024-01-15-preview.
 *
 * Other available API versions: 2025-03-30-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native programmableconnectivity [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getOperatorApiConnection(args: GetOperatorApiConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetOperatorApiConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:programmableconnectivity:getOperatorApiConnection", {
        "operatorApiConnectionName": args.operatorApiConnectionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOperatorApiConnectionArgs {
    /**
     * Azure Programmable Connectivity (APC) Operator API Connection Name.
     */
    operatorApiConnectionName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A Programmable Connectivity Operator API Connection resource
 */
export interface GetOperatorApiConnectionResult {
    /**
     * Type of the account the user has with the Operator's Network API infrastructure. AzureManaged | UserManaged.
     */
    readonly accountType: string;
    /**
     * Application ID of the App Developer that is registered with the Operator in a specific country/region.
     */
    readonly appId?: string;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The Network API for the current operator in the country/region provided in the linked Operator API Plan.
     */
    readonly camaraApiName: string;
    /**
     * Details about the Application that would use the Operator's Network APIs.
     */
    readonly configuredApplication?: outputs.programmableconnectivity.ApplicationPropertiesResponse;
    /**
     * Reference to the APC Gateway resource ID.
     */
    readonly gatewayId: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Reference to the Operator API Plan Resource ID.
     */
    readonly operatorApiPlanId: string;
    /**
     * Name of the Operator in the linked Operator API Plan belongs to.
     */
    readonly operatorName: string;
    /**
     * The status of the last operation.
     */
    readonly provisioningState: string;
    /**
     * Details about the SaaS offer purchased from the marketplace.
     */
    readonly saasProperties?: outputs.programmableconnectivity.SaasPropertiesResponse;
    /**
     * The status of the OperatorApiConnection resource.
     */
    readonly status: outputs.programmableconnectivity.StatusResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.programmableconnectivity.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get an Operator API Connection.
 *
 * Uses Azure REST API version 2024-01-15-preview.
 *
 * Other available API versions: 2025-03-30-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native programmableconnectivity [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getOperatorApiConnectionOutput(args: GetOperatorApiConnectionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOperatorApiConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:programmableconnectivity:getOperatorApiConnection", {
        "operatorApiConnectionName": args.operatorApiConnectionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOperatorApiConnectionOutputArgs {
    /**
     * Azure Programmable Connectivity (APC) Operator API Connection Name.
     */
    operatorApiConnectionName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
