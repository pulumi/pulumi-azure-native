// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Azure Resource Manager resource envelope.
 *
 * Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-10-01-preview.
 *
 * Other available API versions: 2024-10-01-preview, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class CapabilityHost extends pulumi.CustomResource {
    /**
     * Get an existing CapabilityHost resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CapabilityHost {
        return new CapabilityHost(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:machinelearningservices:CapabilityHost';

    /**
     * Returns true if the given object is an instance of CapabilityHost.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CapabilityHost {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CapabilityHost.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * [Required] Additional attributes of the entity.
     */
    public readonly capabilityHostProperties!: pulumi.Output<outputs.machinelearningservices.CapabilityHostResponse>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.machinelearningservices.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CapabilityHost resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CapabilityHostArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.capabilityHostProperties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capabilityHostProperties'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["capabilityHostProperties"] = args ? (args.capabilityHostProperties ? pulumi.output(args.capabilityHostProperties).apply(inputs.machinelearningservices.capabilityHostArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["capabilityHostProperties"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:machinelearningservices/v20241001preview:CapabilityHost" }, { type: "azure-native:machinelearningservices/v20250101preview:CapabilityHost" }, { type: "azure-native:machinelearningservices/v20250401preview:CapabilityHost" }, { type: "azure-native:machinelearningservices/v20250601:CapabilityHost" }, { type: "azure-native:machinelearningservices/v20250701preview:CapabilityHost" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(CapabilityHost.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CapabilityHost resource.
 */
export interface CapabilityHostArgs {
    /**
     * [Required] Additional attributes of the entity.
     */
    capabilityHostProperties: pulumi.Input<inputs.machinelearningservices.CapabilityHostArgs>;
    /**
     * CapabilityHost name.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    workspaceName: pulumi.Input<string>;
}
