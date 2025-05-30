// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The details of subscription under management group.
 *
 * Uses Azure REST API version 2023-04-01. In version 2.x of the Azure Native provider, it used API version 2021-04-01.
 *
 * Other available API versions: 2021-04-01, 2024-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native management [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class ManagementGroupSubscription extends pulumi.CustomResource {
    /**
     * Get an existing ManagementGroupSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagementGroupSubscription {
        return new ManagementGroupSubscription(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:management:ManagementGroupSubscription';

    /**
     * Returns true if the given object is an instance of ManagementGroupSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagementGroupSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagementGroupSubscription.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The friendly name of the subscription.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the parent management group.
     */
    public /*out*/ readonly parent!: pulumi.Output<outputs.management.DescendantParentGroupInfoResponse | undefined>;
    /**
     * The state of the subscription.
     */
    public /*out*/ readonly state!: pulumi.Output<string | undefined>;
    /**
     * The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
     */
    public /*out*/ readonly tenant!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagementGroupSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagementGroupSubscriptionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["subscriptionId"] = args ? args.subscriptionId : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parent"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tenant"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parent"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tenant"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:management/v20200501:ManagementGroupSubscription" }, { type: "azure-native:management/v20201001:ManagementGroupSubscription" }, { type: "azure-native:management/v20210401:ManagementGroupSubscription" }, { type: "azure-native:management/v20230401:ManagementGroupSubscription" }, { type: "azure-native:management/v20240201preview:ManagementGroupSubscription" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ManagementGroupSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagementGroupSubscription resource.
 */
export interface ManagementGroupSubscriptionArgs {
    /**
     * Management Group ID.
     */
    groupId: pulumi.Input<string>;
    /**
     * Subscription ID.
     */
    subscriptionId?: pulumi.Input<string>;
}
