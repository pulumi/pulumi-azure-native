// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class Workflow extends pulumi.CustomResource {
    /**
     * Get an existing Workflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workflow {
        return new Workflow(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:logic/v20150201preview:Workflow';

    /**
     * Returns true if the given object is an instance of Workflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workflow.__pulumiType;
    }

    /**
     * Gets the access endpoint.
     */
    public /*out*/ readonly accessEndpoint!: pulumi.Output<string>;
    /**
     * Gets the changed time.
     */
    public /*out*/ readonly changedTime!: pulumi.Output<string>;
    /**
     * Gets the created time.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * Gets or sets the definition.
     */
    public readonly definition!: pulumi.Output<any | undefined>;
    /**
     * Gets or sets the link to definition.
     */
    public readonly definitionLink!: pulumi.Output<outputs.logic.v20150201preview.ContentLinkResponse | undefined>;
    /**
     * Gets or sets the resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Gets the resource name.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the parameters.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: outputs.logic.v20150201preview.WorkflowParameterResponse} | undefined>;
    /**
     * Gets or sets the link to parameters.
     */
    public readonly parametersLink!: pulumi.Output<outputs.logic.v20150201preview.ContentLinkResponse | undefined>;
    /**
     * Gets the provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Gets or sets the sku.
     */
    public readonly sku!: pulumi.Output<outputs.logic.v20150201preview.SkuResponse | undefined>;
    /**
     * Gets or sets the state.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource type.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Gets the version.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Workflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["definition"] = args ? args.definition : undefined;
            resourceInputs["definitionLink"] = args ? args.definitionLink : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["parametersLink"] = args ? args.parametersLink : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["workflowName"] = args ? args.workflowName : undefined;
            resourceInputs["accessEndpoint"] = undefined /*out*/;
            resourceInputs["changedTime"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        } else {
            resourceInputs["accessEndpoint"] = undefined /*out*/;
            resourceInputs["changedTime"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["definition"] = undefined /*out*/;
            resourceInputs["definitionLink"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["parametersLink"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["sku"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:logic/v20160601:Workflow" }, { type: "azure-native:logic/v20180701preview:Workflow" }, { type: "azure-native:logic/v20190501:Workflow" }, { type: "azure-native:logic:Workflow" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Workflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workflow resource.
 */
export interface WorkflowArgs {
    /**
     * Gets or sets the definition.
     */
    definition?: any;
    /**
     * Gets or sets the link to definition.
     */
    definitionLink?: pulumi.Input<inputs.logic.v20150201preview.ContentLinkArgs>;
    /**
     * Gets or sets the resource id.
     */
    id?: pulumi.Input<string>;
    /**
     * Gets or sets the resource location.
     */
    location?: pulumi.Input<string>;
    /**
     * Gets the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Gets or sets the parameters.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<inputs.logic.v20150201preview.WorkflowParameterArgs>}>;
    /**
     * Gets or sets the link to parameters.
     */
    parametersLink?: pulumi.Input<inputs.logic.v20150201preview.ContentLinkArgs>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets the sku.
     */
    sku?: pulumi.Input<inputs.logic.v20150201preview.SkuArgs>;
    /**
     * Gets or sets the state.
     */
    state?: pulumi.Input<enums.logic.v20150201preview.WorkflowState>;
    /**
     * Gets or sets the resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Gets the resource type.
     */
    type?: pulumi.Input<string>;
    /**
     * The workflow name.
     */
    workflowName?: pulumi.Input<string>;
}
