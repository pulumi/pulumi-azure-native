// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * GlobalRulestack prefixList
 *
 * Uses Azure REST API version 2025-02-06-preview. In version 2.x of the Azure Native provider, it used API version 2023-09-01.
 *
 * Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class PrefixListGlobalRulestack extends pulumi.CustomResource {
    /**
     * Get an existing PrefixListGlobalRulestack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrefixListGlobalRulestack {
        return new PrefixListGlobalRulestack(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:cloudngfw:PrefixListGlobalRulestack';

    /**
     * Returns true if the given object is an instance of PrefixListGlobalRulestack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrefixListGlobalRulestack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrefixListGlobalRulestack.__pulumiType;
    }

    /**
     * comment for this object
     */
    public readonly auditComment!: pulumi.Output<string | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * prefix description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * etag info
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * prefix list
     */
    public readonly prefixList!: pulumi.Output<string[]>;
    /**
     * Provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.cloudngfw.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PrefixListGlobalRulestack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrefixListGlobalRulestackArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.globalRulestackName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'globalRulestackName'");
            }
            if ((!args || args.prefixList === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prefixList'");
            }
            resourceInputs["auditComment"] = args ? args.auditComment : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["globalRulestackName"] = args ? args.globalRulestackName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["prefixList"] = args ? args.prefixList : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["auditComment"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["prefixList"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:cloudngfw/v20220829:PrefixListGlobalRulestack" }, { type: "azure-native:cloudngfw/v20220829preview:PrefixListGlobalRulestack" }, { type: "azure-native:cloudngfw/v20230901:PrefixListGlobalRulestack" }, { type: "azure-native:cloudngfw/v20230901preview:PrefixListGlobalRulestack" }, { type: "azure-native:cloudngfw/v20231010preview:PrefixListGlobalRulestack" }, { type: "azure-native:cloudngfw/v20240119preview:PrefixListGlobalRulestack" }, { type: "azure-native:cloudngfw/v20240207preview:PrefixListGlobalRulestack" }, { type: "azure-native:cloudngfw/v20250206preview:PrefixListGlobalRulestack" }, { type: "azure-native:cloudngfw/v20250523:PrefixListGlobalRulestack" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PrefixListGlobalRulestack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrefixListGlobalRulestack resource.
 */
export interface PrefixListGlobalRulestackArgs {
    /**
     * comment for this object
     */
    auditComment?: pulumi.Input<string>;
    /**
     * prefix description
     */
    description?: pulumi.Input<string>;
    /**
     * GlobalRulestack resource name
     */
    globalRulestackName: pulumi.Input<string>;
    /**
     * Local Rule priority
     */
    name?: pulumi.Input<string>;
    /**
     * prefix list
     */
    prefixList: pulumi.Input<pulumi.Input<string>[]>;
}
