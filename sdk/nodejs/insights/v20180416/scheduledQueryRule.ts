// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * The Log Search Rule resource.
 */
export class ScheduledQueryRule extends pulumi.CustomResource {
    /**
     * Get an existing ScheduledQueryRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ScheduledQueryRule {
        return new ScheduledQueryRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:insights/v20180416:ScheduledQueryRule';

    /**
     * Returns true if the given object is an instance of ScheduledQueryRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScheduledQueryRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScheduledQueryRule.__pulumiType;
    }

    /**
     * Action needs to be taken on rule execution.
     */
    public readonly action!: pulumi.Output<outputs.insights.v20180416.AlertingActionResponse | outputs.insights.v20180416.LogToMetricActionResponse>;
    /**
     * The flag that indicates whether the alert should be automatically resolved or not. The default is false.
     */
    public readonly autoMitigate!: pulumi.Output<boolean | undefined>;
    /**
     * The api-version used when creating this alert rule
     */
    public /*out*/ readonly createdWithApiVersion!: pulumi.Output<string>;
    /**
     * The description of the Log Search rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the alert rule
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The flag which indicates whether the Log Search rule is enabled. Value should be true or false
     */
    public readonly enabled!: pulumi.Output<string | undefined>;
    /**
     * The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. 
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * True if alert rule is legacy Log Analytic rule
     */
    public /*out*/ readonly isLegacyLogAnalyticsRule!: pulumi.Output<boolean>;
    /**
     * Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Last time the rule was updated in IS08601 format.
     */
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the scheduled query rule
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Schedule (Frequency, Time Window) for rule. Required for action type - AlertingAction
     */
    public readonly schedule!: pulumi.Output<outputs.insights.v20180416.ScheduleResponse | undefined>;
    /**
     * Data Source against which rule will Query Data
     */
    public readonly source!: pulumi.Output<outputs.insights.v20180416.SourceResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ScheduledQueryRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduledQueryRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["autoMitigate"] = (args ? args.autoMitigate : undefined) ?? false;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdWithApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["isLegacyLogAnalyticsRule"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["autoMitigate"] = undefined /*out*/;
            resourceInputs["createdWithApiVersion"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["isLegacyLogAnalyticsRule"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["schedule"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:insights/v20200501preview:ScheduledQueryRule" }, { type: "azure-native:insights/v20210201preview:ScheduledQueryRule" }, { type: "azure-native:insights/v20210801:ScheduledQueryRule" }, { type: "azure-native:insights/v20220615:ScheduledQueryRule" }, { type: "azure-native:insights/v20220801preview:ScheduledQueryRule" }, { type: "azure-native:insights/v20230315preview:ScheduledQueryRule" }, { type: "azure-native:insights/v20231201:ScheduledQueryRule" }, { type: "azure-native:insights/v20240101preview:ScheduledQueryRule" }, { type: "azure-native:insights/v20250101preview:ScheduledQueryRule" }, { type: "azure-native:insights:ScheduledQueryRule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ScheduledQueryRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScheduledQueryRule resource.
 */
export interface ScheduledQueryRuleArgs {
    /**
     * Action needs to be taken on rule execution.
     */
    action: pulumi.Input<inputs.insights.v20180416.AlertingActionArgs | inputs.insights.v20180416.LogToMetricActionArgs>;
    /**
     * The flag that indicates whether the alert should be automatically resolved or not. The default is false.
     */
    autoMitigate?: pulumi.Input<boolean>;
    /**
     * The description of the Log Search rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the alert rule
     */
    displayName?: pulumi.Input<string>;
    /**
     * The flag which indicates whether the Log Search rule is enabled. Value should be true or false
     */
    enabled?: pulumi.Input<string | enums.insights.v20180416.Enabled>;
    /**
     * Resource location
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the rule.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * Schedule (Frequency, Time Window) for rule. Required for action type - AlertingAction
     */
    schedule?: pulumi.Input<inputs.insights.v20180416.ScheduleArgs>;
    /**
     * Data Source against which rule will Query Data
     */
    source: pulumi.Input<inputs.insights.v20180416.SourceArgs>;
    /**
     * Resource tags
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
