// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets Business Application Agent.
 *
 * Uses Azure REST API version 2025-01-01-preview.
 *
 * Other available API versions: 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getBusinessApplicationAgent(args: GetBusinessApplicationAgentArgs, opts?: pulumi.InvokeOptions): Promise<GetBusinessApplicationAgentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:securityinsights:getBusinessApplicationAgent", {
        "agentResourceName": args.agentResourceName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetBusinessApplicationAgentArgs {
    /**
     * Business Application Agent Name
     */
    agentResourceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the workspace.
     */
    workspaceName: string;
}

/**
 * Describes the configuration of a Business Application Agent.
 */
export interface GetBusinessApplicationAgentResult {
    readonly agentSystems: outputs.securityinsights.AgentSystemResponse[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Describes the configuration of a Business Application Agent.
     */
    readonly configuration: outputs.securityinsights.SapAgentConfigurationResponse;
    readonly displayName: string;
    /**
     * Etag of the azure resource
     */
    readonly etag?: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    readonly lastModifiedTimeUtc: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.securityinsights.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets Business Application Agent.
 *
 * Uses Azure REST API version 2025-01-01-preview.
 *
 * Other available API versions: 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getBusinessApplicationAgentOutput(args: GetBusinessApplicationAgentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBusinessApplicationAgentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:securityinsights:getBusinessApplicationAgent", {
        "agentResourceName": args.agentResourceName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetBusinessApplicationAgentOutputArgs {
    /**
     * Business Application Agent Name
     */
    agentResourceName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
