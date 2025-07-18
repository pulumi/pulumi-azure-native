// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The storage account blob inventory policy.
 *
 * Uses Azure REST API version 2024-01-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
 *
 * Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class BlobInventoryPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BlobInventoryPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BlobInventoryPolicy {
        return new BlobInventoryPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:storage:BlobInventoryPolicy';

    /**
     * Returns true if the given object is an instance of BlobInventoryPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlobInventoryPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlobInventoryPolicy.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Returns the last modified date and time of the blob inventory policy.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The storage account blob inventory policy object. It is composed of policy rules.
     */
    public readonly policy!: pulumi.Output<outputs.storage.BlobInventoryPolicySchemaResponse>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.storage.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a BlobInventoryPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlobInventoryPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["blobInventoryPolicyName"] = args ? args.blobInventoryPolicyName : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["policy"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:storage/v20190601:BlobInventoryPolicy" }, { type: "azure-native:storage/v20200801preview:BlobInventoryPolicy" }, { type: "azure-native:storage/v20210101:BlobInventoryPolicy" }, { type: "azure-native:storage/v20210201:BlobInventoryPolicy" }, { type: "azure-native:storage/v20210401:BlobInventoryPolicy" }, { type: "azure-native:storage/v20210601:BlobInventoryPolicy" }, { type: "azure-native:storage/v20210801:BlobInventoryPolicy" }, { type: "azure-native:storage/v20210901:BlobInventoryPolicy" }, { type: "azure-native:storage/v20220501:BlobInventoryPolicy" }, { type: "azure-native:storage/v20220901:BlobInventoryPolicy" }, { type: "azure-native:storage/v20230101:BlobInventoryPolicy" }, { type: "azure-native:storage/v20230401:BlobInventoryPolicy" }, { type: "azure-native:storage/v20230501:BlobInventoryPolicy" }, { type: "azure-native:storage/v20240101:BlobInventoryPolicy" }, { type: "azure-native:storage/v20250101:BlobInventoryPolicy" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(BlobInventoryPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BlobInventoryPolicy resource.
 */
export interface BlobInventoryPolicyArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name of the storage account blob inventory policy. It should always be 'default'
     */
    blobInventoryPolicyName?: pulumi.Input<string>;
    /**
     * The storage account blob inventory policy object. It is composed of policy rules.
     */
    policy: pulumi.Input<inputs.storage.BlobInventoryPolicySchemaArgs>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
