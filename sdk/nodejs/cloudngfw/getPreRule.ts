// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a PreRulesResource
 *
 * Uses Azure REST API version 2025-02-06-preview.
 *
 * Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPreRule(args: GetPreRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetPreRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cloudngfw:getPreRule", {
        "globalRulestackName": args.globalRulestackName,
        "priority": args.priority,
    }, opts);
}

export interface GetPreRuleArgs {
    /**
     * GlobalRulestack resource name
     */
    globalRulestackName: string;
    /**
     * Pre Rule priority
     */
    priority: string;
}

/**
 * PreRulestack rule list
 */
export interface GetPreRuleResult {
    /**
     * rule action
     */
    readonly actionType?: string;
    /**
     * array of rule applications
     */
    readonly applications?: string[];
    /**
     * rule comment
     */
    readonly auditComment?: string;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * rule category
     */
    readonly category?: outputs.cloudngfw.CategoryResponse;
    /**
     * enable or disable decryption
     */
    readonly decryptionRuleType?: string;
    /**
     * rule description
     */
    readonly description?: string;
    /**
     * destination address
     */
    readonly destination?: outputs.cloudngfw.DestinationAddrResponse;
    /**
     * enable or disable logging
     */
    readonly enableLogging?: string;
    /**
     * etag info
     */
    readonly etag?: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * inbound Inspection Certificate
     */
    readonly inboundInspectionCertificate?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * cidr should not be 'any'
     */
    readonly negateDestination?: string;
    /**
     * cidr should not be 'any'
     */
    readonly negateSource?: string;
    readonly priority: number;
    /**
     * any, application-default, TCP:number, UDP:number
     */
    readonly protocol?: string;
    /**
     * prot port list
     */
    readonly protocolPortList?: string[];
    /**
     * Provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * rule name
     */
    readonly ruleName: string;
    /**
     * state of this rule
     */
    readonly ruleState?: string;
    /**
     * source address
     */
    readonly source?: outputs.cloudngfw.SourceAddrResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.cloudngfw.SystemDataResponse;
    /**
     * tag for rule
     */
    readonly tags?: outputs.cloudngfw.TagInfoResponse[];
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a PreRulesResource
 *
 * Uses Azure REST API version 2025-02-06-preview.
 *
 * Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPreRuleOutput(args: GetPreRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPreRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cloudngfw:getPreRule", {
        "globalRulestackName": args.globalRulestackName,
        "priority": args.priority,
    }, opts);
}

export interface GetPreRuleOutputArgs {
    /**
     * GlobalRulestack resource name
     */
    globalRulestackName: pulumi.Input<string>;
    /**
     * Pre Rule priority
     */
    priority: pulumi.Input<string>;
}
