// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2021-09-01-preview.
 *
 * Other available API versions: 2021-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native providerhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class SkusNestedResourceTypeThird extends pulumi.CustomResource {
    /**
     * Get an existing SkusNestedResourceTypeThird resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SkusNestedResourceTypeThird {
        return new SkusNestedResourceTypeThird(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:providerhub:SkusNestedResourceTypeThird';

    /**
     * Returns true if the given object is an instance of SkusNestedResourceTypeThird.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SkusNestedResourceTypeThird {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SkusNestedResourceTypeThird.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly properties!: pulumi.Output<outputs.providerhub.SkuResourceResponseProperties>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.providerhub.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SkusNestedResourceTypeThird resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SkusNestedResourceTypeThirdArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.nestedResourceTypeFirst === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nestedResourceTypeFirst'");
            }
            if ((!args || args.nestedResourceTypeSecond === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nestedResourceTypeSecond'");
            }
            if ((!args || args.nestedResourceTypeThird === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nestedResourceTypeThird'");
            }
            if ((!args || args.providerNamespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerNamespace'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            resourceInputs["nestedResourceTypeFirst"] = args ? args.nestedResourceTypeFirst : undefined;
            resourceInputs["nestedResourceTypeSecond"] = args ? args.nestedResourceTypeSecond : undefined;
            resourceInputs["nestedResourceTypeThird"] = args ? args.nestedResourceTypeThird : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["providerNamespace"] = args ? args.providerNamespace : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
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
        const aliasOpts = { aliases: [{ type: "azure-native:providerhub/v20201120:SkusNestedResourceTypeThird" }, { type: "azure-native:providerhub/v20210501preview:SkusNestedResourceTypeThird" }, { type: "azure-native:providerhub/v20210601preview:SkusNestedResourceTypeThird" }, { type: "azure-native:providerhub/v20210901preview:SkusNestedResourceTypeThird" }, { type: "azure-native:providerhub/v20240901:SkusNestedResourceTypeThird" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SkusNestedResourceTypeThird.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SkusNestedResourceTypeThird resource.
 */
export interface SkusNestedResourceTypeThirdArgs {
    /**
     * The first child resource type.
     */
    nestedResourceTypeFirst: pulumi.Input<string>;
    /**
     * The second child resource type.
     */
    nestedResourceTypeSecond: pulumi.Input<string>;
    /**
     * The third child resource type.
     */
    nestedResourceTypeThird: pulumi.Input<string>;
    properties?: pulumi.Input<inputs.providerhub.SkuResourcePropertiesArgs>;
    /**
     * The name of the resource provider hosted within ProviderHub.
     */
    providerNamespace: pulumi.Input<string>;
    /**
     * The resource type.
     */
    resourceType: pulumi.Input<string>;
    /**
     * The SKU.
     */
    sku?: pulumi.Input<string>;
}
