// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets properties of a DNS security rule for a DNS resolver policy.
 *
 * Uses Azure REST API version 2023-07-01-preview.
 *
 * Other available API versions: 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getDnsSecurityRule(args: GetDnsSecurityRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetDnsSecurityRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:dnsresolver:getDnsSecurityRule", {
        "dnsResolverPolicyName": args.dnsResolverPolicyName,
        "dnsSecurityRuleName": args.dnsSecurityRuleName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDnsSecurityRuleArgs {
    /**
     * The name of the DNS resolver policy.
     */
    dnsResolverPolicyName: string;
    /**
     * The name of the DNS security rule.
     */
    dnsSecurityRuleName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Describes a DNS security rule.
 */
export interface GetDnsSecurityRuleResult {
    /**
     * The action to take on DNS requests that match the DNS security rule.
     */
    readonly action: outputs.dnsresolver.DnsSecurityRuleActionResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * DNS resolver policy domains lists that the DNS security rule applies to.
     */
    readonly dnsResolverDomainLists: outputs.dnsresolver.SubResourceResponse[];
    /**
     * The state of DNS security rule.
     */
    readonly dnsSecurityRuleState?: string;
    /**
     * ETag of the DNS security rule.
     */
    readonly etag: string;
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
     * The priority of the DNS security rule.
     */
    readonly priority: number;
    /**
     * The current provisioning state of the DNS security rule. This is a read-only property and any attempt to set this value will be ignored.
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.dnsresolver.SystemDataResponse;
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
 * Gets properties of a DNS security rule for a DNS resolver policy.
 *
 * Uses Azure REST API version 2023-07-01-preview.
 *
 * Other available API versions: 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getDnsSecurityRuleOutput(args: GetDnsSecurityRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDnsSecurityRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:dnsresolver:getDnsSecurityRule", {
        "dnsResolverPolicyName": args.dnsResolverPolicyName,
        "dnsSecurityRuleName": args.dnsSecurityRuleName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDnsSecurityRuleOutputArgs {
    /**
     * The name of the DNS resolver policy.
     */
    dnsResolverPolicyName: pulumi.Input<string>;
    /**
     * The name of the DNS security rule.
     */
    dnsSecurityRuleName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
