// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Defines the GuestAgent.
 */
export class GuestAgent extends pulumi.CustomResource {
    /**
     * Get an existing GuestAgent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GuestAgent {
        return new GuestAgent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:azurestackhci/v20240501preview:GuestAgent';

    /**
     * Returns true if the given object is an instance of GuestAgent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GuestAgent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GuestAgent.__pulumiType;
    }

    /**
     * Username / Password Credentials to provision guest agent.
     */
    public readonly credentials!: pulumi.Output<outputs.azurestackhci.v20240501preview.GuestCredentialResponse | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The guest agent provisioning action.
     */
    public readonly provisioningAction!: pulumi.Output<string | undefined>;
    /**
     * Provisioning state of the virtual machine instance.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The guest agent status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.azurestackhci.v20240501preview.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GuestAgent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GuestAgentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceUri'");
            }
            resourceInputs["credentials"] = args ? args.credentials : undefined;
            resourceInputs["provisioningAction"] = args ? args.provisioningAction : undefined;
            resourceInputs["resourceUri"] = args ? args.resourceUri : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["credentials"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningAction"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:azurestackhci/v20230701preview:GuestAgent" }, { type: "azure-native:azurestackhci/v20230901preview:GuestAgent" }, { type: "azure-native:azurestackhci/v20240101:GuestAgent" }, { type: "azure-native:azurestackhci/v20240201preview:GuestAgent" }, { type: "azure-native:azurestackhci/v20240715preview:GuestAgent" }, { type: "azure-native:azurestackhci/v20240801preview:GuestAgent" }, { type: "azure-native:azurestackhci/v20241001preview:GuestAgent" }, { type: "azure-native:azurestackhci/v20250201preview:GuestAgent" }, { type: "azure-native:azurestackhci/v20250401preview:GuestAgent" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GuestAgent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GuestAgent resource.
 */
export interface GuestAgentArgs {
    /**
     * Username / Password Credentials to provision guest agent.
     */
    credentials?: pulumi.Input<inputs.azurestackhci.v20240501preview.GuestCredentialArgs>;
    /**
     * The guest agent provisioning action.
     */
    provisioningAction?: pulumi.Input<string | enums.azurestackhci.v20240501preview.ProvisioningAction>;
    /**
     * The fully qualified Azure Resource manager identifier of the resource.
     */
    resourceUri: pulumi.Input<string>;
}
