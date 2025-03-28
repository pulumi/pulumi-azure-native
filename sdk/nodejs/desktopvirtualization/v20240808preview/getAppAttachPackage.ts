// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get an app attach package.
 */
export function getAppAttachPackage(args: GetAppAttachPackageArgs, opts?: pulumi.InvokeOptions): Promise<GetAppAttachPackageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:desktopvirtualization/v20240808preview:getAppAttachPackage", {
        "appAttachPackageName": args.appAttachPackageName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAppAttachPackageArgs {
    /**
     * The name of the App Attach package
     */
    appAttachPackageName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Schema for App Attach Package properties.
 */
export interface GetAppAttachPackageResult {
    /**
     * Field that can be populated with custom data and filtered on in list GET calls
     */
    readonly customData?: string;
    /**
     * Parameter indicating how the health check should behave if this package fails staging
     */
    readonly failHealthCheckOnStagingFailure?: string;
    /**
     * List of Hostpool resource Ids.
     */
    readonly hostPoolReferences?: string[];
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * Detailed properties for App Attach Package
     */
    readonly image?: outputs.desktopvirtualization.v20240808preview.AppAttachPackageInfoPropertiesResponse;
    /**
     * URL path to certificate name located in keyVault
     */
    readonly keyVaultURL?: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Lookback url to third party control plane, is null for native app attach packages
     */
    readonly packageLookbackUrl?: string;
    /**
     * Specific name of package owner, is "AppAttach" for native app attach packages
     */
    readonly packageOwnerName?: string;
    /**
     * The provisioning state of the App Attach Package.
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.desktopvirtualization.v20240808preview.SystemDataResponse;
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
 * Get an app attach package.
 */
export function getAppAttachPackageOutput(args: GetAppAttachPackageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAppAttachPackageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:desktopvirtualization/v20240808preview:getAppAttachPackage", {
        "appAttachPackageName": args.appAttachPackageName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAppAttachPackageOutputArgs {
    /**
     * The name of the App Attach package
     */
    appAttachPackageName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
