// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets an existing delivery rule within a rule set.
 *
 * Uses Azure REST API version 2024-09-01.
 *
 * Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getRule(args: GetRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cdn:getRule", {
        "profileName": args.profileName,
        "resourceGroupName": args.resourceGroupName,
        "ruleName": args.ruleName,
        "ruleSetName": args.ruleSetName,
    }, opts);
}

export interface GetRuleArgs {
    /**
     * Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
     */
    profileName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    resourceGroupName: string;
    /**
     * Name of the delivery rule which is unique within the endpoint.
     */
    ruleName: string;
    /**
     * Name of the rule set under the profile.
     */
    ruleSetName: string;
}

/**
 * Friendly Rules name mapping to the any Rules or secret related information.
 */
export interface GetRuleResult {
    /**
     * A list of actions that are executed when all the conditions of a rule are satisfied.
     */
    readonly actions: (outputs.cdn.DeliveryRuleCacheExpirationActionResponse | outputs.cdn.DeliveryRuleCacheKeyQueryStringActionResponse | outputs.cdn.DeliveryRuleRequestHeaderActionResponse | outputs.cdn.DeliveryRuleResponseHeaderActionResponse | outputs.cdn.DeliveryRuleRouteConfigurationOverrideActionResponse | outputs.cdn.OriginGroupOverrideActionResponse | outputs.cdn.UrlRedirectActionResponse | outputs.cdn.UrlRewriteActionResponse | outputs.cdn.UrlSigningActionResponse)[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * A list of conditions that must be matched for the actions to be executed
     */
    readonly conditions?: (outputs.cdn.DeliveryRuleClientPortConditionResponse | outputs.cdn.DeliveryRuleCookiesConditionResponse | outputs.cdn.DeliveryRuleHostNameConditionResponse | outputs.cdn.DeliveryRuleHttpVersionConditionResponse | outputs.cdn.DeliveryRuleIsDeviceConditionResponse | outputs.cdn.DeliveryRulePostArgsConditionResponse | outputs.cdn.DeliveryRuleQueryStringConditionResponse | outputs.cdn.DeliveryRuleRemoteAddressConditionResponse | outputs.cdn.DeliveryRuleRequestBodyConditionResponse | outputs.cdn.DeliveryRuleRequestHeaderConditionResponse | outputs.cdn.DeliveryRuleRequestMethodConditionResponse | outputs.cdn.DeliveryRuleRequestSchemeConditionResponse | outputs.cdn.DeliveryRuleRequestUriConditionResponse | outputs.cdn.DeliveryRuleServerPortConditionResponse | outputs.cdn.DeliveryRuleSocketAddrConditionResponse | outputs.cdn.DeliveryRuleSslProtocolConditionResponse | outputs.cdn.DeliveryRuleUrlFileExtensionConditionResponse | outputs.cdn.DeliveryRuleUrlFileNameConditionResponse | outputs.cdn.DeliveryRuleUrlPathConditionResponse)[];
    readonly deploymentStatus: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
     */
    readonly matchProcessingBehavior?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
     */
    readonly order: number;
    /**
     * Provisioning status
     */
    readonly provisioningState: string;
    /**
     * The name of the rule set containing the rule.
     */
    readonly ruleSetName: string;
    /**
     * Read only system data
     */
    readonly systemData: outputs.cdn.SystemDataResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Gets an existing delivery rule within a rule set.
 *
 * Uses Azure REST API version 2024-09-01.
 *
 * Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getRuleOutput(args: GetRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cdn:getRule", {
        "profileName": args.profileName,
        "resourceGroupName": args.resourceGroupName,
        "ruleName": args.ruleName,
        "ruleSetName": args.ruleSetName,
    }, opts);
}

export interface GetRuleOutputArgs {
    /**
     * Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
     */
    profileName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the delivery rule which is unique within the endpoint.
     */
    ruleName: pulumi.Input<string>;
    /**
     * Name of the rule set under the profile.
     */
    ruleSetName: pulumi.Input<string>;
}
