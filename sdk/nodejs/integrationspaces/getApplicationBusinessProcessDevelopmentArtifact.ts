// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The get business process development artifact action.
 *
 * Uses Azure REST API version 2023-11-14-preview.
 */
export function getApplicationBusinessProcessDevelopmentArtifact(args: GetApplicationBusinessProcessDevelopmentArtifactArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationBusinessProcessDevelopmentArtifactResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:integrationspaces:getApplicationBusinessProcessDevelopmentArtifact", {
        "applicationName": args.applicationName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetApplicationBusinessProcessDevelopmentArtifactArgs {
    /**
     * The name of the Application
     */
    applicationName: string;
    /**
     * The name of the business process development artifact.
     */
    name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the space
     */
    spaceName: string;
}

/**
 * The business process development artifact save or get response.
 */
export interface GetApplicationBusinessProcessDevelopmentArtifactResult {
    /**
     * The name of the business process development artifact.
     */
    readonly name: string;
    /**
     * The properties of the business process development artifact.
     */
    readonly properties: outputs.integrationspaces.BusinessProcessDevelopmentArtifactPropertiesResponse;
    /**
     * The system data of the business process development artifact.
     */
    readonly systemData: outputs.integrationspaces.SystemDataResponse;
}
/**
 * The get business process development artifact action.
 *
 * Uses Azure REST API version 2023-11-14-preview.
 */
export function getApplicationBusinessProcessDevelopmentArtifactOutput(args: GetApplicationBusinessProcessDevelopmentArtifactOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationBusinessProcessDevelopmentArtifactResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:integrationspaces:getApplicationBusinessProcessDevelopmentArtifact", {
        "applicationName": args.applicationName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetApplicationBusinessProcessDevelopmentArtifactOutputArgs {
    /**
     * The name of the Application
     */
    applicationName: pulumi.Input<string>;
    /**
     * The name of the business process development artifact.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the space
     */
    spaceName: pulumi.Input<string>;
}
