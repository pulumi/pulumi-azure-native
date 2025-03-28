// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a workspace instance.
 */
export function getWorkspace(args: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:operationalinsights/v20250201:getWorkspace", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceArgs {
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
 * The top level Workspace resource container.
 */
export interface GetWorkspaceResult {
    /**
     * Workspace creation date.
     */
    readonly createdDate: string;
    /**
     * This is a read-only property. Represents the ID associated with the workspace.
     */
    readonly customerId: string;
    /**
     * The resource ID of the default Data Collection Rule to use for this workspace. Expected format is - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dcrName}.
     */
    readonly defaultDataCollectionRuleResourceId?: string;
    /**
     * The etag of the workspace.
     */
    readonly etag?: string;
    /**
     * workspace failover properties.
     */
    readonly failover?: outputs.operationalinsights.v20250201.WorkspaceFailoverPropertiesResponse;
    /**
     * Workspace features.
     */
    readonly features?: outputs.operationalinsights.v20250201.WorkspaceFeaturesResponse;
    /**
     * Indicates whether customer managed storage is mandatory for query management.
     */
    readonly forceCmkForQuery?: boolean;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.operationalinsights.v20250201.IdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Workspace modification date.
     */
    readonly modifiedDate: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * List of linked private link scope resources.
     */
    readonly privateLinkScopedResources: outputs.operationalinsights.v20250201.PrivateLinkScopedResourceResponse[];
    /**
     * The provisioning state of the workspace.
     */
    readonly provisioningState: string;
    /**
     * The network access type for accessing Log Analytics ingestion.
     */
    readonly publicNetworkAccessForIngestion?: string;
    /**
     * The network access type for accessing Log Analytics query.
     */
    readonly publicNetworkAccessForQuery?: string;
    /**
     * workspace replication properties.
     */
    readonly replication?: outputs.operationalinsights.v20250201.WorkspaceReplicationPropertiesResponse;
    /**
     * The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers documentation for details.
     */
    readonly retentionInDays?: number;
    /**
     * The SKU of the workspace.
     */
    readonly sku?: outputs.operationalinsights.v20250201.WorkspaceSkuResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.operationalinsights.v20250201.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The daily volume cap for ingestion.
     */
    readonly workspaceCapping?: outputs.operationalinsights.v20250201.WorkspaceCappingResponse;
}
/**
 * Gets a workspace instance.
 */
export function getWorkspaceOutput(args: GetWorkspaceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWorkspaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:operationalinsights/v20250201:getWorkspace", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
