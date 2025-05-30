// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Configuration information or intent on which to do the analysis on.
 *
 * Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2024-01-01-preview.
 *
 * Other available API versions: 2024-01-01-preview, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class ReachabilityAnalysisIntent extends pulumi.CustomResource {
    /**
     * Get an existing ReachabilityAnalysisIntent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReachabilityAnalysisIntent {
        return new ReachabilityAnalysisIntent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:network:ReachabilityAnalysisIntent';

    /**
     * Returns true if the given object is an instance of ReachabilityAnalysisIntent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReachabilityAnalysisIntent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReachabilityAnalysisIntent.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Represents the Reachability Analysis Intent properties.
     */
    public readonly properties!: pulumi.Output<outputs.network.ReachabilityAnalysisIntentPropertiesResponse>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.network.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ReachabilityAnalysisIntent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReachabilityAnalysisIntentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.networkManagerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkManagerName'");
            }
            if ((!args || args.properties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'properties'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["networkManagerName"] = args ? args.networkManagerName : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["reachabilityAnalysisIntentName"] = args ? args.reachabilityAnalysisIntentName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:network/v20240101preview:ReachabilityAnalysisIntent" }, { type: "azure-native:network/v20240501:ReachabilityAnalysisIntent" }, { type: "azure-native:network/v20240701:ReachabilityAnalysisIntent" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ReachabilityAnalysisIntent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReachabilityAnalysisIntent resource.
 */
export interface ReachabilityAnalysisIntentArgs {
    /**
     * The name of the network manager.
     */
    networkManagerName: pulumi.Input<string>;
    /**
     * Represents the Reachability Analysis Intent properties.
     */
    properties: pulumi.Input<inputs.network.ReachabilityAnalysisIntentPropertiesArgs>;
    /**
     * Reachability Analysis Intent name.
     */
    reachabilityAnalysisIntentName?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Workspace name.
     */
    workspaceName: pulumi.Input<string>;
}
