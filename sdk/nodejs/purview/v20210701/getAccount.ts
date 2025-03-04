// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get an account
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:purview/v20210701:getAccount", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAccountArgs {
    /**
     * The name of the account.
     */
    accountName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
}

/**
 * Account resource
 */
export interface GetAccountResult {
    /**
     * Cloud connectors.
     * External cloud identifier used as part of scanning configuration.
     */
    readonly cloudConnectors?: outputs.purview.v20210701.CloudConnectorsResponse;
    /**
     * Gets the time at which the entity was created.
     */
    readonly createdAt: string;
    /**
     * Gets the creator of the entity.
     */
    readonly createdBy: string;
    /**
     * Gets the creators of the entity's object id.
     */
    readonly createdByObjectId: string;
    /**
     * The URIs that are the public endpoints of the account.
     */
    readonly endpoints: outputs.purview.v20210701.AccountPropertiesResponseEndpoints;
    /**
     * Gets or sets the friendly name.
     */
    readonly friendlyName: string;
    /**
     * Gets or sets the identifier.
     */
    readonly id: string;
    /**
     * Identity Info on the tracked resource
     */
    readonly identity?: outputs.purview.v20210701.IdentityResponse;
    /**
     * Gets or sets the location.
     */
    readonly location?: string;
    /**
     * Gets or sets the managed resource group name
     */
    readonly managedResourceGroupName?: string;
    /**
     * Gets the resource identifiers of the managed resources.
     */
    readonly managedResources: outputs.purview.v20210701.AccountPropertiesResponseManagedResources;
    /**
     * Gets or sets the name.
     */
    readonly name: string;
    /**
     * Gets the private endpoint connections information.
     */
    readonly privateEndpointConnections: outputs.purview.v20210701.PrivateEndpointConnectionResponse[];
    /**
     * Gets or sets the state of the provisioning.
     */
    readonly provisioningState: string;
    /**
     * Gets or sets the public network access.
     */
    readonly publicNetworkAccess?: string;
    /**
     * Gets or sets the Sku.
     */
    readonly sku: outputs.purview.v20210701.AccountResponseSku;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.purview.v20210701.TrackedResourceResponseSystemData;
    /**
     * Tags on the azure resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Gets or sets the type.
     */
    readonly type: string;
}
/**
 * Get an account
 */
export function getAccountOutput(args: GetAccountOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccountResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:purview/v20210701:getAccount", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAccountOutputArgs {
    /**
     * The name of the account.
     */
    accountName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
}
