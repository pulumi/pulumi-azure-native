// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Class representing the config for the solution in the migrate project.
 *
 * Uses Azure REST API version 2018-09-01-preview.
 */
export function getSolutionConfig(args: GetSolutionConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetSolutionConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:migrate:getSolutionConfig", {
        "migrateProjectName": args.migrateProjectName,
        "resourceGroupName": args.resourceGroupName,
        "solutionName": args.solutionName,
    }, opts);
}

export interface GetSolutionConfigArgs {
    /**
     * Name of the Azure Migrate project.
     */
    migrateProjectName: string;
    /**
     * Name of the Azure Resource Group that migrate project is part of.
     */
    resourceGroupName: string;
    /**
     * Unique name of a migration solution within a migrate project.
     */
    solutionName: string;
}

/**
 * Class representing the config for the solution in the migrate project.
 */
export interface GetSolutionConfigResult {
    /**
     * Gets or sets the publisher sas uri for the solution.
     */
    readonly publisherSasUri?: string;
}
/**
 * Class representing the config for the solution in the migrate project.
 *
 * Uses Azure REST API version 2018-09-01-preview.
 */
export function getSolutionConfigOutput(args: GetSolutionConfigOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSolutionConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:migrate:getSolutionConfig", {
        "migrateProjectName": args.migrateProjectName,
        "resourceGroupName": args.resourceGroupName,
        "solutionName": args.solutionName,
    }, opts);
}

export interface GetSolutionConfigOutputArgs {
    /**
     * Name of the Azure Migrate project.
     */
    migrateProjectName: pulumi.Input<string>;
    /**
     * Name of the Azure Resource Group that migrate project is part of.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Unique name of a migration solution within a migrate project.
     */
    solutionName: pulumi.Input<string>;
}
