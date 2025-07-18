// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Uses Azure REST API version 2025-04-01-preview.
 *
 * Other available API versions: 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getOutboundRule(args: GetOutboundRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetOutboundRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:machinelearningservices:getOutboundRule", {
        "managedNetworkName": args.managedNetworkName,
        "resourceGroupName": args.resourceGroupName,
        "ruleName": args.ruleName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetOutboundRuleArgs {
    /**
     * Name of the managedNetwork associated with the workspace. Only 'default' is supported.
     */
    managedNetworkName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Name of the workspace managed network outbound rule
     */
    ruleName: string;
    /**
     * Azure Machine Learning Workspace Name
     */
    workspaceName: string;
}

export interface GetOutboundRuleResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Outbound Rule for the managed network of a machine learning workspace.
     */
    readonly properties: outputs.machinelearningservices.FqdnOutboundRuleResponse | outputs.machinelearningservices.PrivateEndpointOutboundRuleResponse | outputs.machinelearningservices.ServiceTagOutboundRuleResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.machinelearningservices.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Uses Azure REST API version 2025-04-01-preview.
 *
 * Other available API versions: 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getOutboundRuleOutput(args: GetOutboundRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOutboundRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:machinelearningservices:getOutboundRule", {
        "managedNetworkName": args.managedNetworkName,
        "resourceGroupName": args.resourceGroupName,
        "ruleName": args.ruleName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetOutboundRuleOutputArgs {
    /**
     * Name of the managedNetwork associated with the workspace. Only 'default' is supported.
     */
    managedNetworkName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the workspace managed network outbound rule
     */
    ruleName: pulumi.Input<string>;
    /**
     * Azure Machine Learning Workspace Name
     */
    workspaceName: pulumi.Input<string>;
}
