// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * MAK key details.
 *
 * Uses Azure REST API version 2019-09-16-preview. In version 2.x of the Azure Native provider, it used API version 2019-09-16-preview.
 */
export class MultipleActivationKey extends pulumi.CustomResource {
    /**
     * Get an existing MultipleActivationKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MultipleActivationKey {
        return new MultipleActivationKey(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:windowsesu:MultipleActivationKey';

    /**
     * Returns true if the given object is an instance of MultipleActivationKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MultipleActivationKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MultipleActivationKey.__pulumiType;
    }

    /**
     * Agreement number under which the key is requested.
     */
    public readonly agreementNumber!: pulumi.Output<string | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * End of support of security updates activated by the MAK key.
     */
    public /*out*/ readonly expirationDate!: pulumi.Output<string>;
    /**
     * Number of activations/servers using the MAK key.
     */
    public readonly installedServerNumber!: pulumi.Output<number | undefined>;
    /**
     * <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
     */
    public readonly isEligible!: pulumi.Output<boolean | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * MAK 5x5 key.
     */
    public /*out*/ readonly multipleActivationKey!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Type of OS for which the key is requested.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Type of support
     */
    public readonly supportType!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a MultipleActivationKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MultipleActivationKeyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["agreementNumber"] = args ? args.agreementNumber : undefined;
            resourceInputs["installedServerNumber"] = args ? args.installedServerNumber : undefined;
            resourceInputs["isEligible"] = args ? args.isEligible : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["multipleActivationKeyName"] = args ? args.multipleActivationKeyName : undefined;
            resourceInputs["osType"] = args ? args.osType : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["supportType"] = (args ? args.supportType : undefined) ?? "SupplementalServicing";
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["expirationDate"] = undefined /*out*/;
            resourceInputs["multipleActivationKey"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["agreementNumber"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["expirationDate"] = undefined /*out*/;
            resourceInputs["installedServerNumber"] = undefined /*out*/;
            resourceInputs["isEligible"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["multipleActivationKey"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["osType"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["supportType"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:windowsesu/v20190916preview:MultipleActivationKey" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(MultipleActivationKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MultipleActivationKey resource.
 */
export interface MultipleActivationKeyArgs {
    /**
     * Agreement number under which the key is requested.
     */
    agreementNumber?: pulumi.Input<string>;
    /**
     * Number of activations/servers using the MAK key.
     */
    installedServerNumber?: pulumi.Input<number>;
    /**
     * <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
     */
    isEligible?: pulumi.Input<boolean>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the MAK key.
     */
    multipleActivationKeyName?: pulumi.Input<string>;
    /**
     * Type of OS for which the key is requested.
     */
    osType?: pulumi.Input<string | enums.windowsesu.OsType>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Type of support
     */
    supportType?: pulumi.Input<string | enums.windowsesu.SupportType>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
