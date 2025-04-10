// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the console for the user.
 *
 * Uses Azure REST API version 2018-10-01.
 */
export function getConsole(args: GetConsoleArgs, opts?: pulumi.InvokeOptions): Promise<GetConsoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:portal:getConsole", {
        "consoleName": args.consoleName,
    }, opts);
}

export interface GetConsoleArgs {
    /**
     * The name of the console
     */
    consoleName: string;
}

/**
 * Cloud shell console
 */
export interface GetConsoleResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Cloud shell console properties.
     */
    readonly properties: outputs.portal.ConsolePropertiesResponse;
}
/**
 * Gets the console for the user.
 *
 * Uses Azure REST API version 2018-10-01.
 */
export function getConsoleOutput(args: GetConsoleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConsoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:portal:getConsole", {
        "consoleName": args.consoleName,
    }, opts);
}

export interface GetConsoleOutputArgs {
    /**
     * The name of the console
     */
    consoleName: pulumi.Input<string>;
}
