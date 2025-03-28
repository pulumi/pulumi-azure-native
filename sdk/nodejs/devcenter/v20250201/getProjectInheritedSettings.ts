// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets applicable inherited settings for this project.
 */
export function getProjectInheritedSettings(args: GetProjectInheritedSettingsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectInheritedSettingsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:devcenter/v20250201:getProjectInheritedSettings", {
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProjectInheritedSettingsArgs {
    /**
     * The name of the project.
     */
    projectName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Applicable inherited settings for a project.
 */
export interface GetProjectInheritedSettingsResult {
    /**
     * Network settings that will be enforced on this project.
     */
    readonly networkSettings: outputs.devcenter.v20250201.ProjectNetworkSettingsResponse;
    /**
     * Dev Center settings to be used when associating a project with a catalog.
     */
    readonly projectCatalogSettings: outputs.devcenter.v20250201.DevCenterProjectCatalogSettingsResponse;
}
/**
 * Gets applicable inherited settings for this project.
 */
export function getProjectInheritedSettingsOutput(args: GetProjectInheritedSettingsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProjectInheritedSettingsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:devcenter/v20250201:getProjectInheritedSettings", {
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProjectInheritedSettingsOutputArgs {
    /**
     * The name of the project.
     */
    projectName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
