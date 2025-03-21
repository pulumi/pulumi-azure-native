// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get the API portal custom domain.
 */
export function getApiPortalCustomDomain(args: GetApiPortalCustomDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetApiPortalCustomDomainResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:appplatform/v20230701preview:getApiPortalCustomDomain", {
        "apiPortalName": args.apiPortalName,
        "domainName": args.domainName,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetApiPortalCustomDomainArgs {
    /**
     * The name of API portal.
     */
    apiPortalName: string;
    /**
     * The name of the API portal custom domain.
     */
    domainName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: string;
    /**
     * The name of the Service resource.
     */
    serviceName: string;
}

/**
 * Custom domain of the API portal
 */
export interface GetApiPortalCustomDomainResult {
    /**
     * Fully qualified resource Id for the resource.
     */
    readonly id: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The properties of custom domain for API portal
     */
    readonly properties: outputs.appplatform.v20230701preview.ApiPortalCustomDomainPropertiesResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.appplatform.v20230701preview.SystemDataResponse;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
/**
 * Get the API portal custom domain.
 */
export function getApiPortalCustomDomainOutput(args: GetApiPortalCustomDomainOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApiPortalCustomDomainResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:appplatform/v20230701preview:getApiPortalCustomDomain", {
        "apiPortalName": args.apiPortalName,
        "domainName": args.domainName,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetApiPortalCustomDomainOutputArgs {
    /**
     * The name of API portal.
     */
    apiPortalName: pulumi.Input<string>;
    /**
     * The name of the API portal custom domain.
     */
    domainName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Service resource.
     */
    serviceName: pulumi.Input<string>;
}
