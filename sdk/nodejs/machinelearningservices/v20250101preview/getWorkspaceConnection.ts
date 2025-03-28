// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getWorkspaceConnection(args: GetWorkspaceConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:machinelearningservices/v20250101preview:getWorkspaceConnection", {
        "connectionName": args.connectionName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceConnectionArgs {
    /**
     * Friendly name of the workspace connection
     */
    connectionName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Azure Machine Learning Workspace Name
     */
    workspaceName: string;
}

export interface GetWorkspaceConnectionResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    readonly properties: outputs.machinelearningservices.v20250101preview.AADAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.AccessKeyAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.AccountKeyAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.ApiKeyAuthWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.CustomKeysWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.ManagedIdentityAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.NoneAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.OAuth2AuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.PATAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.SASAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.ServicePrincipalAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.v20250101preview.UsernamePasswordAuthTypeWorkspaceConnectionPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.machinelearningservices.v20250101preview.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
export function getWorkspaceConnectionOutput(args: GetWorkspaceConnectionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWorkspaceConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:machinelearningservices/v20250101preview:getWorkspaceConnection", {
        "connectionName": args.connectionName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceConnectionOutputArgs {
    /**
     * Friendly name of the workspace connection
     */
    connectionName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Azure Machine Learning Workspace Name
     */
    workspaceName: pulumi.Input<string>;
}
