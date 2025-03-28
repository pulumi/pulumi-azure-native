// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * The top level Workspace resource container.
 */
export class Workspace extends pulumi.CustomResource {
    /**
     * Get an existing Workspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workspace {
        return new Workspace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:operationalinsights/v20151101preview:Workspace';

    /**
     * Returns true if the given object is an instance of Workspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workspace.__pulumiType;
    }

    /**
     * This is a read-only property. Represents the ID associated with the workspace.
     */
    public /*out*/ readonly customerId!: pulumi.Output<string>;
    /**
     * The ETag of the workspace.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * This is a legacy property and is not used anymore. Kept here for backward compatibility.
     */
    public /*out*/ readonly portalUrl!: pulumi.Output<string>;
    /**
     * The provisioning state of the workspace.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus. 
     */
    public readonly retentionInDays!: pulumi.Output<number | undefined>;
    /**
     * The SKU of the workspace.
     */
    public readonly sku!: pulumi.Output<outputs.operationalinsights.v20151101preview.SkuResponse | undefined>;
    /**
     * This is a read-only legacy property. It is always set to 'Azure' by the service. Kept here for backward compatibility.
     */
    public /*out*/ readonly source!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["eTag"] = args ? args.eTag : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["provisioningState"] = args ? args.provisioningState : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["customerId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["portalUrl"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["customerId"] = undefined /*out*/;
            resourceInputs["eTag"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["portalUrl"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["retentionInDays"] = undefined /*out*/;
            resourceInputs["sku"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:operationalinsights/v20200301preview:Workspace" }, { type: "azure-native:operationalinsights/v20200801:Workspace" }, { type: "azure-native:operationalinsights/v20201001:Workspace" }, { type: "azure-native:operationalinsights/v20210601:Workspace" }, { type: "azure-native:operationalinsights/v20211201preview:Workspace" }, { type: "azure-native:operationalinsights/v20221001:Workspace" }, { type: "azure-native:operationalinsights/v20230901:Workspace" }, { type: "azure-native:operationalinsights/v20250201:Workspace" }, { type: "azure-native:operationalinsights:Workspace" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Workspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * The ETag of the workspace.
     */
    eTag?: pulumi.Input<string>;
    /**
     * Resource location
     */
    location?: pulumi.Input<string>;
    /**
     * The provisioning state of the workspace.
     */
    provisioningState?: pulumi.Input<string | enums.operationalinsights.v20151101preview.EntityStatus>;
    /**
     * The resource group name of the workspace.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus. 
     */
    retentionInDays?: pulumi.Input<number>;
    /**
     * The SKU of the workspace.
     */
    sku?: pulumi.Input<inputs.operationalinsights.v20151101preview.SkuArgs>;
    /**
     * Resource tags
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the workspace.
     */
    workspaceName?: pulumi.Input<string>;
}
