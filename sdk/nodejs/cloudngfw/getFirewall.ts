// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a FirewallResource
 *
 * Uses Azure REST API version 2025-02-06-preview.
 *
 * Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getFirewall(args: GetFirewallArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cloudngfw:getFirewall", {
        "firewallName": args.firewallName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFirewallArgs {
    /**
     * Firewall resource name
     */
    firewallName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * PaloAltoNetworks Firewall
 */
export interface GetFirewallResult {
    /**
     * Associated Rulestack
     */
    readonly associatedRulestack?: outputs.cloudngfw.RulestackDetailsResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * DNS settings for Firewall
     */
    readonly dnsSettings: outputs.cloudngfw.DNSSettingsResponse;
    /**
     * Frontend settings for Firewall
     */
    readonly frontEndSettings?: outputs.cloudngfw.FrontendSettingResponse[];
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The managed service identities assigned to this resource.
     */
    readonly identity?: outputs.cloudngfw.AzureResourceManagerManagedIdentityPropertiesResponse;
    /**
     * Panorama Managed: Default is False. Default will be CloudSec managed
     */
    readonly isPanoramaManaged?: string;
    /**
     * Strata Cloud Managed: Default is False. Default will be CloudSec managed
     */
    readonly isStrataCloudManaged?: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Marketplace details
     */
    readonly marketplaceDetails: outputs.cloudngfw.MarketplaceDetailsResponse;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Network settings
     */
    readonly networkProfile: outputs.cloudngfw.NetworkProfileResponse;
    /**
     * panEtag info
     */
    readonly panEtag?: string;
    /**
     * Panorama Configuration
     */
    readonly panoramaConfig?: outputs.cloudngfw.PanoramaConfigResponse;
    /**
     * Billing plan information.
     */
    readonly planData: outputs.cloudngfw.PlanDataResponse;
    /**
     * Provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * Strata Cloud Manager Configuration, only applicable if Strata Cloud Manager is selected.
     */
    readonly strataCloudManagerConfig?: outputs.cloudngfw.StrataCloudManagerConfigResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.cloudngfw.SystemDataResponse;
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
 * Get a FirewallResource
 *
 * Uses Azure REST API version 2025-02-06-preview.
 *
 * Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getFirewallOutput(args: GetFirewallOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFirewallResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cloudngfw:getFirewall", {
        "firewallName": args.firewallName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFirewallOutputArgs {
    /**
     * Firewall resource name
     */
    firewallName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
