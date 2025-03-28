// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a data connector.
 */
export function getRestApiPollerDataConnector(args: GetRestApiPollerDataConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetRestApiPollerDataConnectorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:securityinsights/v20250301:getRestApiPollerDataConnector", {
        "dataConnectorId": args.dataConnectorId,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetRestApiPollerDataConnectorArgs {
    /**
     * Connector ID
     */
    dataConnectorId: string;
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
 * Represents Rest Api Poller data connector.
 */
export interface GetRestApiPollerDataConnectorResult {
    /**
     * The add on attributes. The key name will become attribute name (a column) and the value will become the attribute value in the payload.
     */
    readonly addOnAttributes?: {[key: string]: string};
    /**
     * The a authentication model.
     */
    readonly auth: outputs.securityinsights.v20250301.AWSAuthModelResponse | outputs.securityinsights.v20250301.ApiKeyAuthModelResponse | outputs.securityinsights.v20250301.BasicAuthModelResponse | outputs.securityinsights.v20250301.GCPAuthModelResponse | outputs.securityinsights.v20250301.GenericBlobSbsAuthModelResponse | outputs.securityinsights.v20250301.GitHubAuthModelResponse | outputs.securityinsights.v20250301.JwtAuthModelResponse | outputs.securityinsights.v20250301.NoneAuthModelResponse | outputs.securityinsights.v20250301.OAuthModelResponse | outputs.securityinsights.v20250301.OracleAuthModelResponse | outputs.securityinsights.v20250301.SessionAuthModelResponse;
    /**
     * The connector definition name (the dataConnectorDefinition resource id).
     */
    readonly connectorDefinitionName: string;
    /**
     * The Log Analytics table destination.
     */
    readonly dataType?: string;
    /**
     * The DCR related properties.
     */
    readonly dcrConfig?: outputs.securityinsights.v20250301.DCRConfigurationResponse;
    /**
     * Etag of the azure resource
     */
    readonly etag?: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * Indicates whether the connector is active or not.
     */
    readonly isActive?: boolean;
    /**
     * The kind of the data connector
     * Expected value is 'RestApiPoller'.
     */
    readonly kind: "RestApiPoller";
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The paging configuration.
     */
    readonly paging?: outputs.securityinsights.v20250301.RestApiPollerRequestPagingConfigResponse;
    /**
     * The request configuration.
     */
    readonly request: outputs.securityinsights.v20250301.RestApiPollerRequestConfigResponse;
    /**
     * The response configuration.
     */
    readonly response?: outputs.securityinsights.v20250301.CcpResponseConfigResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.securityinsights.v20250301.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets a data connector.
 */
export function getRestApiPollerDataConnectorOutput(args: GetRestApiPollerDataConnectorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRestApiPollerDataConnectorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:securityinsights/v20250301:getRestApiPollerDataConnector", {
        "dataConnectorId": args.dataConnectorId,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetRestApiPollerDataConnectorOutputArgs {
    /**
     * Connector ID
     */
    dataConnectorId: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
