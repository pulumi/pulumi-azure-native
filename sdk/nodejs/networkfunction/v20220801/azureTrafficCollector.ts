// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Azure Traffic Collector resource.
 */
export class AzureTrafficCollector extends pulumi.CustomResource {
    /**
     * Get an existing AzureTrafficCollector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AzureTrafficCollector {
        return new AzureTrafficCollector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:networkfunction/v20220801:AzureTrafficCollector';

    /**
     * Returns true if the given object is an instance of AzureTrafficCollector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureTrafficCollector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureTrafficCollector.__pulumiType;
    }

    /**
     * Collector Policies for Azure Traffic Collector.
     */
    public readonly collectorPolicies!: pulumi.Output<outputs.networkfunction.v20220801.CollectorPolicyResponse[] | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the application rule collection resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.networkfunction.v20220801.TrackedResourceResponseSystemData>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The virtualHub to which the Azure Traffic Collector belongs.
     */
    public /*out*/ readonly virtualHub!: pulumi.Output<outputs.networkfunction.v20220801.ResourceReferenceResponse | undefined>;

    /**
     * Create a AzureTrafficCollector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureTrafficCollectorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["azureTrafficCollectorName"] = args ? args.azureTrafficCollectorName : undefined;
            resourceInputs["collectorPolicies"] = args ? args.collectorPolicies : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["virtualHub"] = undefined /*out*/;
        } else {
            resourceInputs["collectorPolicies"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["virtualHub"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:networkfunction/v20210901preview:AzureTrafficCollector" }, { type: "azure-native:networkfunction/v20220501:AzureTrafficCollector" }, { type: "azure-native:networkfunction/v20221101:AzureTrafficCollector" }, { type: "azure-native:networkfunction:AzureTrafficCollector" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(AzureTrafficCollector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AzureTrafficCollector resource.
 */
export interface AzureTrafficCollectorArgs {
    /**
     * Azure Traffic Collector name
     */
    azureTrafficCollectorName?: pulumi.Input<string>;
    /**
     * Collector Policies for Azure Traffic Collector.
     * These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
     */
    collectorPolicies?: pulumi.Input<pulumi.Input<inputs.networkfunction.v20220801.CollectorPolicyArgs>[]>;
    /**
     * Resource location.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
