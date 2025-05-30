// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a DynamoDBContinuousBackupsDescription
 *
 * Uses Azure REST API version 2024-12-01.
 */
export function getDynamoDbContinuousBackupsDescription(args: GetDynamoDbContinuousBackupsDescriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetDynamoDbContinuousBackupsDescriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:awsconnector:getDynamoDbContinuousBackupsDescription", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDynamoDbContinuousBackupsDescriptionArgs {
    /**
     * Name of DynamoDBContinuousBackupsDescription
     */
    name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A Microsoft.AwsConnector resource
 */
export interface GetDynamoDbContinuousBackupsDescriptionResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
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
     * The resource-specific properties for this resource.
     */
    readonly properties: outputs.awsconnector.DynamoDBContinuousBackupsDescriptionPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.awsconnector.SystemDataResponse;
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
 * Get a DynamoDBContinuousBackupsDescription
 *
 * Uses Azure REST API version 2024-12-01.
 */
export function getDynamoDbContinuousBackupsDescriptionOutput(args: GetDynamoDbContinuousBackupsDescriptionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDynamoDbContinuousBackupsDescriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:awsconnector:getDynamoDbContinuousBackupsDescription", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDynamoDbContinuousBackupsDescriptionOutputArgs {
    /**
     * Name of DynamoDBContinuousBackupsDescription
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
