// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a ObservabilityServiceResource
 * Azure REST API version: 2023-10-15-preview.
 */
export function getObservabilityService(args: GetObservabilityServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetObservabilityServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:mobilepacketcore:getObservabilityService", {
        "observabilityServiceName": args.observabilityServiceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetObservabilityServiceArgs {
    /**
     * The name of the Observability Service
     */
    observabilityServiceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Azure for Operators 5G Core Observability Service Resource
 */
export interface GetObservabilityServiceResult {
    /**
     * Reference to cluster where the observability components are deployed
     */
    readonly clusterService: string;
    /**
     * Azure for Operators 5G Core Observability component parameters.  One set per component type
     */
    readonly componentParameters: outputs.mobilepacketcore.QualifiedComponentDeploymentParametersResponse[];
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
     * Operational status
     */
    readonly operationalStatus: outputs.mobilepacketcore.OperationalStatusResponse;
    /**
     * The status of the last operation.
     */
    readonly provisioningState: string;
    /**
     * Release version. This is inherited from the cluster
     */
    readonly releaseVersion: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.mobilepacketcore.SystemDataResponse;
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
 * Get a ObservabilityServiceResource
 * Azure REST API version: 2023-10-15-preview.
 */
export function getObservabilityServiceOutput(args: GetObservabilityServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetObservabilityServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:mobilepacketcore:getObservabilityService", {
        "observabilityServiceName": args.observabilityServiceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetObservabilityServiceOutputArgs {
    /**
     * The name of the Observability Service
     */
    observabilityServiceName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
