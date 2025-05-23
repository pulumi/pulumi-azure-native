// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * List all effective security admin rules applied on a virtual network.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listNetworkManagerEffectiveSecurityAdminRules(args: ListNetworkManagerEffectiveSecurityAdminRulesArgs, opts?: pulumi.InvokeOptions): Promise<ListNetworkManagerEffectiveSecurityAdminRulesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network:listNetworkManagerEffectiveSecurityAdminRules", {
        "resourceGroupName": args.resourceGroupName,
        "skipToken": args.skipToken,
        "top": args.top,
        "virtualNetworkName": args.virtualNetworkName,
    }, opts);
}

export interface ListNetworkManagerEffectiveSecurityAdminRulesArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
     */
    skipToken?: string;
    /**
     * An optional query parameter which specifies the maximum number of records to be returned by the server.
     */
    top?: number;
    /**
     * The name of the virtual network.
     */
    virtualNetworkName: string;
}

/**
 * Result of the request to list networkManagerEffectiveSecurityAdminRules. It contains a list of groups and a skiptoken to get the next set of results.
 */
export interface ListNetworkManagerEffectiveSecurityAdminRulesResult {
    /**
     * When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
     */
    readonly skipToken?: string;
    /**
     * Gets a page of NetworkManagerEffectiveSecurityAdminRules
     */
    readonly value?: (outputs.network.EffectiveDefaultSecurityAdminRuleResponse | outputs.network.EffectiveSecurityAdminRuleResponse)[];
}
/**
 * List all effective security admin rules applied on a virtual network.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function listNetworkManagerEffectiveSecurityAdminRulesOutput(args: ListNetworkManagerEffectiveSecurityAdminRulesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListNetworkManagerEffectiveSecurityAdminRulesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network:listNetworkManagerEffectiveSecurityAdminRules", {
        "resourceGroupName": args.resourceGroupName,
        "skipToken": args.skipToken,
        "top": args.top,
        "virtualNetworkName": args.virtualNetworkName,
    }, opts);
}

export interface ListNetworkManagerEffectiveSecurityAdminRulesOutputArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
     */
    skipToken?: pulumi.Input<string>;
    /**
     * An optional query parameter which specifies the maximum number of records to be returned by the server.
     */
    top?: pulumi.Input<number>;
    /**
     * The name of the virtual network.
     */
    virtualNetworkName: pulumi.Input<string>;
}
