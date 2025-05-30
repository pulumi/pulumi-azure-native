// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The Managed Network Group resource
 *
 * Uses Azure REST API version 2019-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2019-06-01-preview.
 */
export class ManagedNetworkGroup extends pulumi.CustomResource {
    /**
     * Get an existing ManagedNetworkGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagedNetworkGroup {
        return new ManagedNetworkGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:managednetwork:ManagedNetworkGroup';

    /**
     * Returns true if the given object is an instance of ManagedNetworkGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedNetworkGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedNetworkGroup.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Responsibility role under which this Managed Network Group will be created
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The collection of management groups covered by the Managed Network
     */
    public readonly managementGroups!: pulumi.Output<outputs.managednetwork.ResourceIdResponse[] | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the ManagedNetwork resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The collection of  subnets covered by the Managed Network
     */
    public readonly subnets!: pulumi.Output<outputs.managednetwork.ResourceIdResponse[] | undefined>;
    /**
     * The collection of subscriptions covered by the Managed Network
     */
    public readonly subscriptions!: pulumi.Output<outputs.managednetwork.ResourceIdResponse[] | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The collection of virtual nets covered by the Managed Network
     */
    public readonly virtualNetworks!: pulumi.Output<outputs.managednetwork.ResourceIdResponse[] | undefined>;

    /**
     * Create a ManagedNetworkGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedNetworkGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.managedNetworkName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managedNetworkName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managedNetworkGroupName"] = args ? args.managedNetworkGroupName : undefined;
            resourceInputs["managedNetworkName"] = args ? args.managedNetworkName : undefined;
            resourceInputs["managementGroups"] = args ? args.managementGroups : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["subnets"] = args ? args.subnets : undefined;
            resourceInputs["subscriptions"] = args ? args.subscriptions : undefined;
            resourceInputs["virtualNetworks"] = args ? args.virtualNetworks : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["managementGroups"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["subnets"] = undefined /*out*/;
            resourceInputs["subscriptions"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["virtualNetworks"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:managednetwork/v20190601preview:ManagedNetworkGroup" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ManagedNetworkGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedNetworkGroup resource.
 */
export interface ManagedNetworkGroupArgs {
    /**
     * Responsibility role under which this Managed Network Group will be created
     */
    kind?: pulumi.Input<string | enums.managednetwork.Kind>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Managed Network Group.
     */
    managedNetworkGroupName?: pulumi.Input<string>;
    /**
     * The name of the Managed Network.
     */
    managedNetworkName: pulumi.Input<string>;
    /**
     * The collection of management groups covered by the Managed Network
     */
    managementGroups?: pulumi.Input<pulumi.Input<inputs.managednetwork.ResourceIdArgs>[]>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The collection of  subnets covered by the Managed Network
     */
    subnets?: pulumi.Input<pulumi.Input<inputs.managednetwork.ResourceIdArgs>[]>;
    /**
     * The collection of subscriptions covered by the Managed Network
     */
    subscriptions?: pulumi.Input<pulumi.Input<inputs.managednetwork.ResourceIdArgs>[]>;
    /**
     * The collection of virtual nets covered by the Managed Network
     */
    virtualNetworks?: pulumi.Input<pulumi.Input<inputs.managednetwork.ResourceIdArgs>[]>;
}
