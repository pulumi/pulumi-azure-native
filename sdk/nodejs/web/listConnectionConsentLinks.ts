// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Lists the consent links of a connection
 *
 * Uses Azure REST API version 2016-06-01.
 *
 * Other available API versions: 2015-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listConnectionConsentLinks(args: ListConnectionConsentLinksArgs, opts?: pulumi.InvokeOptions): Promise<ListConnectionConsentLinksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:web:listConnectionConsentLinks", {
        "connectionName": args.connectionName,
        "parameters": args.parameters,
        "resourceGroupName": args.resourceGroupName,
        "subscriptionId": args.subscriptionId,
    }, opts);
}

export interface ListConnectionConsentLinksArgs {
    /**
     * Connection name
     */
    connectionName: string;
    /**
     * Collection of resources
     */
    parameters?: inputs.web.ConsentLinkParameterDefinition[];
    /**
     * The resource group
     */
    resourceGroupName: string;
    /**
     * Subscription Id
     */
    subscriptionId?: string;
}

/**
 * Collection of consent links
 */
export interface ListConnectionConsentLinksResult {
    /**
     * Collection of resources
     */
    readonly value?: outputs.web.ConsentLinkDefinitionResponse[];
}
/**
 * Lists the consent links of a connection
 *
 * Uses Azure REST API version 2016-06-01.
 *
 * Other available API versions: 2015-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listConnectionConsentLinksOutput(args: ListConnectionConsentLinksOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListConnectionConsentLinksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:web:listConnectionConsentLinks", {
        "connectionName": args.connectionName,
        "parameters": args.parameters,
        "resourceGroupName": args.resourceGroupName,
        "subscriptionId": args.subscriptionId,
    }, opts);
}

export interface ListConnectionConsentLinksOutputArgs {
    /**
     * Connection name
     */
    connectionName: pulumi.Input<string>;
    /**
     * Collection of resources
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.web.ConsentLinkParameterDefinitionArgs>[]>;
    /**
     * The resource group
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Subscription Id
     */
    subscriptionId?: pulumi.Input<string>;
}
