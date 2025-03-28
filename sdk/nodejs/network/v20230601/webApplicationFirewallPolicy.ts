// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Defines web application firewall policy.
 */
export class WebApplicationFirewallPolicy extends pulumi.CustomResource {
    /**
     * Get an existing WebApplicationFirewallPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebApplicationFirewallPolicy {
        return new WebApplicationFirewallPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:network/v20230601:WebApplicationFirewallPolicy';

    /**
     * Returns true if the given object is an instance of WebApplicationFirewallPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebApplicationFirewallPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebApplicationFirewallPolicy.__pulumiType;
    }

    /**
     * A collection of references to application gateways.
     */
    public /*out*/ readonly applicationGateways!: pulumi.Output<outputs.network.v20230601.ApplicationGatewayResponse[]>;
    /**
     * The custom rules inside the policy.
     */
    public readonly customRules!: pulumi.Output<outputs.network.v20230601.WebApplicationFirewallCustomRuleResponse[] | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A collection of references to application gateway http listeners.
     */
    public /*out*/ readonly httpListeners!: pulumi.Output<outputs.network.v20230601.SubResourceResponse[]>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Describes the managedRules structure.
     */
    public readonly managedRules!: pulumi.Output<outputs.network.v20230601.ManagedRulesDefinitionResponse>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A collection of references to application gateway path rules.
     */
    public /*out*/ readonly pathBasedRules!: pulumi.Output<outputs.network.v20230601.SubResourceResponse[]>;
    /**
     * The PolicySettings for policy.
     */
    public readonly policySettings!: pulumi.Output<outputs.network.v20230601.PolicySettingsResponse | undefined>;
    /**
     * The provisioning state of the web application firewall policy resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource status of the policy.
     */
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebApplicationFirewallPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebApplicationFirewallPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.managedRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managedRules'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["customRules"] = args ? args.customRules : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managedRules"] = args ? args.managedRules : undefined;
            resourceInputs["policyName"] = args ? args.policyName : undefined;
            resourceInputs["policySettings"] = args ? (args.policySettings ? pulumi.output(args.policySettings).apply(inputs.network.v20230601.policySettingsArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["applicationGateways"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["httpListeners"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pathBasedRules"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["resourceState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["applicationGateways"] = undefined /*out*/;
            resourceInputs["customRules"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["httpListeners"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["managedRules"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pathBasedRules"] = undefined /*out*/;
            resourceInputs["policySettings"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["resourceState"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:network/v20181201:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20190201:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20190401:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20190601:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20190701:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20190801:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20190901:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20191101:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20191201:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20200301:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20200401:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20200501:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20200601:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20200701:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20200801:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20201101:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20210201:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20210301:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20210501:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20210801:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20220101:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20220501:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20220701:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20220901:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20221101:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20230201:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20230401:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20230501:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20230901:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20231101:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20240101:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20240301:WebApplicationFirewallPolicy" }, { type: "azure-native:network/v20240501:WebApplicationFirewallPolicy" }, { type: "azure-native:network:WebApplicationFirewallPolicy" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(WebApplicationFirewallPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebApplicationFirewallPolicy resource.
 */
export interface WebApplicationFirewallPolicyArgs {
    /**
     * The custom rules inside the policy.
     */
    customRules?: pulumi.Input<pulumi.Input<inputs.network.v20230601.WebApplicationFirewallCustomRuleArgs>[]>;
    /**
     * Resource ID.
     */
    id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    location?: pulumi.Input<string>;
    /**
     * Describes the managedRules structure.
     */
    managedRules: pulumi.Input<inputs.network.v20230601.ManagedRulesDefinitionArgs>;
    /**
     * The name of the policy.
     */
    policyName?: pulumi.Input<string>;
    /**
     * The PolicySettings for policy.
     */
    policySettings?: pulumi.Input<inputs.network.v20230601.PolicySettingsArgs>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
