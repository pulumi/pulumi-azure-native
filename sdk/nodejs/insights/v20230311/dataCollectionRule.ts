// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Definition of ARM tracked top level resource.
 */
export class DataCollectionRule extends pulumi.CustomResource {
    /**
     * Get an existing DataCollectionRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataCollectionRule {
        return new DataCollectionRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:insights/v20230311:DataCollectionRule';

    /**
     * Returns true if the given object is an instance of DataCollectionRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataCollectionRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataCollectionRule.__pulumiType;
    }

    /**
     * Agent settings used to modify agent behavior on a given host
     */
    public readonly agentSettings!: pulumi.Output<outputs.insights.v20230311.DataCollectionRuleResponseAgentSettings | undefined>;
    /**
     * The resource ID of the data collection endpoint that this rule can be used with.
     */
    public readonly dataCollectionEndpointId!: pulumi.Output<string | undefined>;
    /**
     * The specification of data flows.
     */
    public readonly dataFlows!: pulumi.Output<outputs.insights.v20230311.DataFlowResponse[] | undefined>;
    /**
     * The specification of data sources. 
     * This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
     */
    public readonly dataSources!: pulumi.Output<outputs.insights.v20230311.DataCollectionRuleResponseDataSources | undefined>;
    /**
     * Description of the data collection rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The specification of destinations.
     */
    public readonly destinations!: pulumi.Output<outputs.insights.v20230311.DataCollectionRuleResponseDestinations | undefined>;
    /**
     * Defines the ingestion endpoints to send data to via this rule.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.insights.v20230311.DataCollectionRuleResponseEndpoints>;
    /**
     * Resource entity tag (ETag).
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Managed service identity of the resource.
     */
    public readonly identity!: pulumi.Output<outputs.insights.v20230311.DataCollectionRuleResourceResponseIdentity | undefined>;
    /**
     * The immutable ID of this data collection rule. This property is READ-ONLY.
     */
    public /*out*/ readonly immutableId!: pulumi.Output<string>;
    /**
     * The kind of the resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Metadata about the resource
     */
    public /*out*/ readonly metadata!: pulumi.Output<outputs.insights.v20230311.DataCollectionRuleResponseMetadata>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Defines all the references that may be used in other sections of the DCR
     */
    public readonly references!: pulumi.Output<outputs.insights.v20230311.DataCollectionRuleResponseReferences | undefined>;
    /**
     * Declaration of custom streams used in this rule.
     */
    public readonly streamDeclarations!: pulumi.Output<{[key: string]: outputs.insights.v20230311.StreamDeclarationResponse} | undefined>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.insights.v20230311.DataCollectionRuleResourceResponseSystemData>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DataCollectionRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataCollectionRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["agentSettings"] = args ? args.agentSettings : undefined;
            resourceInputs["dataCollectionEndpointId"] = args ? args.dataCollectionEndpointId : undefined;
            resourceInputs["dataCollectionRuleName"] = args ? args.dataCollectionRuleName : undefined;
            resourceInputs["dataFlows"] = args ? args.dataFlows : undefined;
            resourceInputs["dataSources"] = args ? args.dataSources : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinations"] = args ? args.destinations : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["references"] = args ? args.references : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["streamDeclarations"] = args ? args.streamDeclarations : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["immutableId"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["agentSettings"] = undefined /*out*/;
            resourceInputs["dataCollectionEndpointId"] = undefined /*out*/;
            resourceInputs["dataFlows"] = undefined /*out*/;
            resourceInputs["dataSources"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["destinations"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["identity"] = undefined /*out*/;
            resourceInputs["immutableId"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["references"] = undefined /*out*/;
            resourceInputs["streamDeclarations"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:insights/v20191101preview:DataCollectionRule" }, { type: "azure-native:insights/v20210401:DataCollectionRule" }, { type: "azure-native:insights/v20210901preview:DataCollectionRule" }, { type: "azure-native:insights/v20220601:DataCollectionRule" }, { type: "azure-native:insights:DataCollectionRule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DataCollectionRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataCollectionRule resource.
 */
export interface DataCollectionRuleArgs {
    /**
     * Agent settings used to modify agent behavior on a given host
     */
    agentSettings?: pulumi.Input<inputs.insights.v20230311.DataCollectionRuleAgentSettingsArgs>;
    /**
     * The resource ID of the data collection endpoint that this rule can be used with.
     */
    dataCollectionEndpointId?: pulumi.Input<string>;
    /**
     * The name of the data collection rule. The name is case insensitive.
     */
    dataCollectionRuleName?: pulumi.Input<string>;
    /**
     * The specification of data flows.
     */
    dataFlows?: pulumi.Input<pulumi.Input<inputs.insights.v20230311.DataFlowArgs>[]>;
    /**
     * The specification of data sources. 
     * This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
     */
    dataSources?: pulumi.Input<inputs.insights.v20230311.DataCollectionRuleDataSourcesArgs>;
    /**
     * Description of the data collection rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The specification of destinations.
     */
    destinations?: pulumi.Input<inputs.insights.v20230311.DataCollectionRuleDestinationsArgs>;
    /**
     * Managed service identity of the resource.
     */
    identity?: pulumi.Input<inputs.insights.v20230311.DataCollectionRuleResourceIdentityArgs>;
    /**
     * The kind of the resource.
     */
    kind?: pulumi.Input<string | enums.insights.v20230311.KnownDataCollectionRuleResourceKind>;
    /**
     * The geo-location where the resource lives.
     */
    location?: pulumi.Input<string>;
    /**
     * Defines all the references that may be used in other sections of the DCR
     */
    references?: pulumi.Input<inputs.insights.v20230311.DataCollectionRuleReferencesArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Declaration of custom streams used in this rule.
     */
    streamDeclarations?: pulumi.Input<{[key: string]: pulumi.Input<inputs.insights.v20230311.StreamDeclarationArgs>}>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
