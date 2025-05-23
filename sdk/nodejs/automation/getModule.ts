// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieve the module identified by module name.
 *
 * Uses Azure REST API version 2023-11-01.
 *
 * Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2022-08-08, 2023-05-15-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getModule(args: GetModuleArgs, opts?: pulumi.InvokeOptions): Promise<GetModuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:automation:getModule", {
        "automationAccountName": args.automationAccountName,
        "moduleName": args.moduleName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetModuleArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: string;
    /**
     * The module name.
     */
    moduleName: string;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: string;
}

/**
 * Definition of the module type.
 */
export interface GetModuleResult {
    /**
     * Gets the activity count of the module.
     */
    readonly activityCount?: number;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Gets the creation time.
     */
    readonly creationTime?: string;
    /**
     * Gets or sets the description.
     */
    readonly description?: string;
    /**
     * Gets the error info of the module.
     */
    readonly error?: outputs.automation.ModuleErrorInfoResponse;
    /**
     * Gets the etag of the resource.
     */
    readonly etag?: string;
    /**
     * Fully qualified resource Id for the resource
     */
    readonly id: string;
    /**
     * Gets type of module, if its composite or not.
     */
    readonly isComposite?: boolean;
    /**
     * Gets the isGlobal flag of the module.
     */
    readonly isGlobal?: boolean;
    /**
     * Gets the last modified time.
     */
    readonly lastModifiedTime?: string;
    /**
     * The Azure Region where the resource lives
     */
    readonly location?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Gets the provisioning state of the module.
     */
    readonly provisioningState?: string;
    /**
     * Gets the size in bytes of the module.
     */
    readonly sizeInBytes?: number;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * Gets the version of the module.
     */
    readonly version?: string;
}
/**
 * Retrieve the module identified by module name.
 *
 * Uses Azure REST API version 2023-11-01.
 *
 * Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2022-08-08, 2023-05-15-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getModuleOutput(args: GetModuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetModuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:automation:getModule", {
        "automationAccountName": args.automationAccountName,
        "moduleName": args.moduleName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetModuleOutputArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: pulumi.Input<string>;
    /**
     * The module name.
     */
    moduleName: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
