// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
 */
export function getManagedCertificate(args: GetManagedCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedCertificateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:app/v20250101:getManagedCertificate", {
        "environmentName": args.environmentName,
        "managedCertificateName": args.managedCertificateName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedCertificateArgs {
    /**
     * Name of the Managed Environment.
     */
    environmentName: string;
    /**
     * Name of the Managed Certificate.
     */
    managedCertificateName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
 */
export interface GetManagedCertificateResult {
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
     * Certificate resource specific properties
     */
    readonly properties: outputs.app.v20250101.ManagedCertificateResponseProperties;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.app.v20250101.SystemDataResponse;
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
 * Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
 */
export function getManagedCertificateOutput(args: GetManagedCertificateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetManagedCertificateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:app/v20250101:getManagedCertificate", {
        "environmentName": args.environmentName,
        "managedCertificateName": args.managedCertificateName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedCertificateOutputArgs {
    /**
     * Name of the Managed Environment.
     */
    environmentName: pulumi.Input<string>;
    /**
     * Name of the Managed Certificate.
     */
    managedCertificateName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
