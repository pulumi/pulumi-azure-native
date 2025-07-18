// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents MicrosoftSecurityIncidentCreation rule.
 *
 * Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
 */
export class MicrosoftSecurityIncidentCreationAlertRule extends pulumi.CustomResource {
    /**
     * Get an existing MicrosoftSecurityIncidentCreationAlertRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MicrosoftSecurityIncidentCreationAlertRule {
        return new MicrosoftSecurityIncidentCreationAlertRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:securityinsights:MicrosoftSecurityIncidentCreationAlertRule';

    /**
     * Returns true if the given object is an instance of MicrosoftSecurityIncidentCreationAlertRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MicrosoftSecurityIncidentCreationAlertRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MicrosoftSecurityIncidentCreationAlertRule.__pulumiType;
    }

    /**
     * The Name of the alert rule template used to create this rule.
     */
    public readonly alertRuleTemplateName!: pulumi.Output<string | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The description of the alert rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name for alerts created by this alert rule.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * the alerts' displayNames on which the cases will not be generated
     */
    public readonly displayNamesExcludeFilter!: pulumi.Output<string[] | undefined>;
    /**
     * the alerts' displayNames on which the cases will be generated
     */
    public readonly displayNamesFilter!: pulumi.Output<string[] | undefined>;
    /**
     * Determines whether this alert rule is enabled or disabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Etag of the azure resource
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The kind of the alert rule
     * Expected value is 'MicrosoftSecurityIncidentCreation'.
     */
    public readonly kind!: pulumi.Output<"MicrosoftSecurityIncidentCreation">;
    /**
     * The last time that this alert has been modified.
     */
    public /*out*/ readonly lastModifiedUtc!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The alerts' productName on which the cases will be generated
     */
    public readonly productFilter!: pulumi.Output<string>;
    /**
     * the alerts' severities on which the cases will be generated
     */
    public readonly severitiesFilter!: pulumi.Output<string[] | undefined>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.securityinsights.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a MicrosoftSecurityIncidentCreationAlertRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MicrosoftSecurityIncidentCreationAlertRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.productFilter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productFilter'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["alertRuleTemplateName"] = args ? args.alertRuleTemplateName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["displayNamesExcludeFilter"] = args ? args.displayNamesExcludeFilter : undefined;
            resourceInputs["displayNamesFilter"] = args ? args.displayNamesFilter : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["kind"] = "MicrosoftSecurityIncidentCreation";
            resourceInputs["productFilter"] = args ? args.productFilter : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["severitiesFilter"] = args ? args.severitiesFilter : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["lastModifiedUtc"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["alertRuleTemplateName"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["displayNamesExcludeFilter"] = undefined /*out*/;
            resourceInputs["displayNamesFilter"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["lastModifiedUtc"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["productFilter"] = undefined /*out*/;
            resourceInputs["severitiesFilter"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:securityinsights/v20190101preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20200101:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20210301preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20210301preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20210301preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20210301preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20210301preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20210901preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20211001:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20211001preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20211001preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20211001preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20220101preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20220401preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20220501preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20220601preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20220701preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20220801:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20220801preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20220901preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20221001preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20221101:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20221101preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20221201preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230201:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230201preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230301preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230401preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230501preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230601preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20230601preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20230601preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230601preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20230601preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20230601preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20230701preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20230701preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20230701preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230701preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20230701preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20230701preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20230801preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20230801preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20230801preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230801preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20230801preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20230801preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20230901preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20230901preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20230901preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20230901preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20230901preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20230901preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20231001preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20231001preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20231001preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20231001preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20231001preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20231001preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20231101:FusionAlertRule" }, { type: "azure-native:securityinsights/v20231101:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20231101:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20231201preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20231201preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20231201preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20231201preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20231201preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20231201preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20240101preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20240101preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20240101preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20240101preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20240101preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20240101preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20240301:FusionAlertRule" }, { type: "azure-native:securityinsights/v20240301:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20240301:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20240401preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20240401preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20240401preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20240401preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20240401preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20240401preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20240901:FusionAlertRule" }, { type: "azure-native:securityinsights/v20240901:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20240901:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20241001preview:FusionAlertRule" }, { type: "azure-native:securityinsights/v20241001preview:MLBehaviorAnalyticsAlertRule" }, { type: "azure-native:securityinsights/v20241001preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20241001preview:NrtAlertRule" }, { type: "azure-native:securityinsights/v20241001preview:ScheduledAlertRule" }, { type: "azure-native:securityinsights/v20241001preview:ThreatIntelligenceAlertRule" }, { type: "azure-native:securityinsights/v20250101preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20250301:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20250401preview:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights/v20250601:MicrosoftSecurityIncidentCreationAlertRule" }, { type: "azure-native:securityinsights:FusionAlertRule" }, { type: "azure-native:securityinsights:ScheduledAlertRule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(MicrosoftSecurityIncidentCreationAlertRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MicrosoftSecurityIncidentCreationAlertRule resource.
 */
export interface MicrosoftSecurityIncidentCreationAlertRuleArgs {
    /**
     * The Name of the alert rule template used to create this rule.
     */
    alertRuleTemplateName?: pulumi.Input<string>;
    /**
     * The description of the alert rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name for alerts created by this alert rule.
     */
    displayName: pulumi.Input<string>;
    /**
     * the alerts' displayNames on which the cases will not be generated
     */
    displayNamesExcludeFilter?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * the alerts' displayNames on which the cases will be generated
     */
    displayNamesFilter?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines whether this alert rule is enabled or disabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The kind of the alert rule
     * Expected value is 'MicrosoftSecurityIncidentCreation'.
     */
    kind: pulumi.Input<"MicrosoftSecurityIncidentCreation">;
    /**
     * The alerts' productName on which the cases will be generated
     */
    productFilter: pulumi.Input<string | enums.securityinsights.MicrosoftSecurityProductName>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Alert rule ID
     */
    ruleId?: pulumi.Input<string>;
    /**
     * the alerts' severities on which the cases will be generated
     */
    severitiesFilter?: pulumi.Input<pulumi.Input<string | enums.securityinsights.AlertSeverity>[]>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
