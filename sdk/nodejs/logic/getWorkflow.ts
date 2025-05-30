// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets a workflow.
 *
 * Uses Azure REST API version 2019-05-01.
 *
 * Other available API versions: 2015-02-01-preview, 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getWorkflow(args: GetWorkflowArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkflowResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:logic:getWorkflow", {
        "resourceGroupName": args.resourceGroupName,
        "workflowName": args.workflowName,
    }, opts);
}

export interface GetWorkflowArgs {
    /**
     * The resource group name.
     */
    resourceGroupName: string;
    /**
     * The workflow name.
     */
    workflowName: string;
}

/**
 * The workflow type.
 */
export interface GetWorkflowResult {
    /**
     * The access control configuration.
     */
    readonly accessControl?: outputs.logic.FlowAccessControlConfigurationResponse;
    /**
     * Gets the access endpoint.
     */
    readonly accessEndpoint: string;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Gets the changed time.
     */
    readonly changedTime: string;
    /**
     * Gets the created time.
     */
    readonly createdTime: string;
    /**
     * The definition.
     */
    readonly definition?: any;
    /**
     * The endpoints configuration.
     */
    readonly endpointsConfiguration?: outputs.logic.FlowEndpointsConfigurationResponse;
    /**
     * The resource id.
     */
    readonly id: string;
    /**
     * Managed service identity properties.
     */
    readonly identity?: outputs.logic.ManagedServiceIdentityResponse;
    /**
     * The integration account.
     */
    readonly integrationAccount?: outputs.logic.ResourceReferenceResponse;
    /**
     * The integration service environment.
     */
    readonly integrationServiceEnvironment?: outputs.logic.ResourceReferenceResponse;
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * Gets the resource name.
     */
    readonly name: string;
    /**
     * The parameters.
     */
    readonly parameters?: {[key: string]: outputs.logic.WorkflowParameterResponse};
    /**
     * Gets the provisioning state.
     */
    readonly provisioningState: string;
    /**
     * The sku.
     */
    readonly sku: outputs.logic.SkuResponse;
    /**
     * The state.
     */
    readonly state?: string;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Gets the resource type.
     */
    readonly type: string;
    /**
     * Gets the version.
     */
    readonly version: string;
}
/**
 * Gets a workflow.
 *
 * Uses Azure REST API version 2019-05-01.
 *
 * Other available API versions: 2015-02-01-preview, 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getWorkflowOutput(args: GetWorkflowOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWorkflowResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:logic:getWorkflow", {
        "resourceGroupName": args.resourceGroupName,
        "workflowName": args.workflowName,
    }, opts);
}

export interface GetWorkflowOutputArgs {
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The workflow name.
     */
    workflowName: pulumi.Input<string>;
}
