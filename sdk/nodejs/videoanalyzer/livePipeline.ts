// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Live pipeline represents a unique instance of a live topology, used for real-time ingestion, archiving and publishing of content for a unique RTSP camera.
 *
 * Uses Azure REST API version 2021-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-11-01-preview.
 */
export class LivePipeline extends pulumi.CustomResource {
    /**
     * Get an existing LivePipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LivePipeline {
        return new LivePipeline(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:videoanalyzer:LivePipeline';

    /**
     * Returns true if the given object is an instance of LivePipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LivePipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LivePipeline.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Maximum bitrate capacity in Kbps reserved for the live pipeline. The allowed range is from 500 to 3000 Kbps in increments of 100 Kbps. If the RTSP camera exceeds this capacity, then the service will disconnect temporarily from the camera. It will retry to re-establish connection (with exponential backoff), checking to see if the camera bitrate is now below the reserved capacity. Doing so will ensure that one 'noisy neighbor' does not affect other live pipelines in your account.
     */
    public readonly bitrateKbps!: pulumi.Output<number>;
    /**
     * An optional description for the pipeline.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
     */
    public readonly parameters!: pulumi.Output<outputs.videoanalyzer.ParameterDefinitionResponse[] | undefined>;
    /**
     * Current state of the pipeline (read-only).
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.videoanalyzer.SystemDataResponse>;
    /**
     * The reference to an existing pipeline topology defined for real-time content processing. When activated, this live pipeline will process content according to the pipeline topology definition.
     */
    public readonly topologyName!: pulumi.Output<string>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LivePipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LivePipelineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.bitrateKbps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bitrateKbps'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.topologyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topologyName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["bitrateKbps"] = args ? args.bitrateKbps : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["livePipelineName"] = args ? args.livePipelineName : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["topologyName"] = args ? args.topologyName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["bitrateKbps"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["topologyName"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:videoanalyzer/v20211101preview:LivePipeline" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(LivePipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LivePipeline resource.
 */
export interface LivePipelineArgs {
    /**
     * The Azure Video Analyzer account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * Maximum bitrate capacity in Kbps reserved for the live pipeline. The allowed range is from 500 to 3000 Kbps in increments of 100 Kbps. If the RTSP camera exceeds this capacity, then the service will disconnect temporarily from the camera. It will retry to re-establish connection (with exponential backoff), checking to see if the camera bitrate is now below the reserved capacity. Doing so will ensure that one 'noisy neighbor' does not affect other live pipelines in your account.
     */
    bitrateKbps: pulumi.Input<number>;
    /**
     * An optional description for the pipeline.
     */
    description?: pulumi.Input<string>;
    /**
     * Live pipeline unique identifier.
     */
    livePipelineName?: pulumi.Input<string>;
    /**
     * List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.videoanalyzer.ParameterDefinitionArgs>[]>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The reference to an existing pipeline topology defined for real-time content processing. When activated, this live pipeline will process content according to the pipeline topology definition.
     */
    topologyName: pulumi.Input<string>;
}
