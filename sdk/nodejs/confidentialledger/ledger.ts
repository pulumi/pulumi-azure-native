// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Confidential Ledger. Contains the properties of Confidential Ledger Resource.
 *
 * Uses Azure REST API version 2023-06-28-preview. In version 2.x of the Azure Native provider, it used API version 2022-05-13.
 *
 * Other available API versions: 2022-05-13, 2022-09-08-preview, 2023-01-26-preview, 2024-07-09-preview, 2024-09-19-preview, 2025-06-10-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native confidentialledger [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class Ledger extends pulumi.CustomResource {
    /**
     * Get an existing Ledger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Ledger {
        return new Ledger(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:confidentialledger:Ledger';

    /**
     * Returns true if the given object is an instance of Ledger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ledger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ledger.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of Confidential Ledger Resource.
     */
    public readonly properties!: pulumi.Output<outputs.confidentialledger.LedgerPropertiesResponse>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.confidentialledger.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Ledger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LedgerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["ledgerName"] = args ? args.ledgerName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:confidentialledger/v20201201preview:Ledger" }, { type: "azure-native:confidentialledger/v20210513preview:Ledger" }, { type: "azure-native:confidentialledger/v20220513:Ledger" }, { type: "azure-native:confidentialledger/v20220908preview:Ledger" }, { type: "azure-native:confidentialledger/v20230126preview:Ledger" }, { type: "azure-native:confidentialledger/v20230628preview:Ledger" }, { type: "azure-native:confidentialledger/v20240709preview:Ledger" }, { type: "azure-native:confidentialledger/v20240919preview:Ledger" }, { type: "azure-native:confidentialledger/v20250610preview:Ledger" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Ledger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Ledger resource.
 */
export interface LedgerArgs {
    /**
     * Name of the Confidential Ledger
     */
    ledgerName?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * Properties of Confidential Ledger Resource.
     */
    properties?: pulumi.Input<inputs.confidentialledger.LedgerPropertiesArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
